package workbench

import (
	"fmt"
	"strings"
	"testing"

	"github.com/vivaui/viva/ui"
	"motty/internal/sshconn"
)

func TestSessionsRestoreWithoutDialingAndReconnectOnEnter(t *testing.T) {
	world, options := newWorld(t)
	first := newTestState(t, options)
	openHost(first)
	first.Click(hitHosts)
	first.Click(hitHostList.At(2))
	openHost(first)
	first.Click(hitTabList.At(int(first.sessions[0].ID)))
	activeAddress := first.sessions[0].Address
	if len(world.conns) != 2 {
		t.Fatalf("initial dials = %d", len(world.conns))
	}

	second := newTestState(t, options)
	if len(world.conns) != 2 {
		t.Fatalf("restoring dialed %d extra connections", len(world.conns)-2)
	}
	if len(second.sessions) != 2 || second.activeSection != sectionTerminal {
		t.Fatalf("restored sessions=%d section=%d", len(second.sessions), second.activeSection)
	}
	active, ok := second.sessionByIDPointer(second.activeSessionID)
	if !ok || active.Address != activeAddress || active.conn != nil || active.phase != sshconn.PhaseClosed {
		t.Fatalf("restored active session = %+v", active)
	}
	if active.term != nil || !active.restored || !second.canReconnect(active) {
		t.Fatalf("restored session = %+v", active)
	}
	canvas, _ := paintedTexts(second)
	retry, close, log := hitConnRetry, hitConnClose, hitConnLog
	if !hasHit(canvas, retry) || !hasHit(canvas, close) || hasHit(canvas, log) {
		t.Fatalf("restored page hits = %+v", canvas.Hits())
	}
	var actionTop float32
	for _, hit := range canvas.Hits() {
		if hit.ID == close || hit.ID == retry {
			if actionTop == 0 || hit.Bounds.Y < actionTop {
				actionTop = hit.Bounds.Y
			}
		}
	}
	// The explanation may wrap onto several lines; it reads as the title,
	// then the body, all above the buttons and in order.
	var title ui.Text
	var body strings.Builder
	var bodyBottom float32
	for _, text := range canvas.Texts() {
		switch {
		case text.Value == second.messages.SessionRestoredTitle:
			title = text
		case title.Value != "" && text.Bounds.Y > title.Bounds.Y && text.Bounds.Y < actionTop && text.Bounds.Y > windowBarHeight:
			body.WriteString(text.Value)
			bodyBottom = max(bodyBottom, text.Bounds.Y+text.Bounds.Height)
		}
	}
	if title.Value == "" {
		t.Fatalf("restored page has no %q", second.messages.SessionRestoredTitle)
	}
	if strings.ReplaceAll(body.String(), " ", "") != strings.ReplaceAll(second.messages.SessionRestoredBody, " ", "") {
		t.Fatalf("restored body = %q, want %q", body.String(), second.messages.SessionRestoredBody)
	}
	if actionTop == 0 || bodyBottom > actionTop {
		t.Fatalf("restored body bottom=%v action top=%v", bodyBottom, actionTop)
	}
	// The divider above the buttons must not run through the explanation.
	for _, rectangle := range canvas.Rectangles() {
		if rectangle.Bounds.Height == 1 && rectangle.Bounds.Y > title.Bounds.Y && rectangle.Bounds.Y < bodyBottom && rectangle.Bounds.Width > 100 {
			t.Fatalf("a rule at y=%v crosses the explanation (title y=%v, bottom=%v)", rectangle.Bounds.Y, title.Bounds.Y, bodyBottom)
		}
	}

	before := len(world.conns)
	second.Click(retry)
	if len(world.conns) != before+1 {
		t.Fatalf("reconnect dials=%d", len(world.conns)-before)
	}
	if len(second.sessions) != 2 || second.activeSessionID != active.ID || active.conn == nil || active.phase != sshconn.PhaseConnecting || active.restored {
		t.Fatalf("reconnected session = %+v", active)
	}
}

func TestRestoreNeverDialsEvenWithALegacyRestorePolicy(t *testing.T) {
	// Reopening the app must not touch the network: a saved tab is metadata until
	// the user connects it. State files from older versions may still carry the
	// removed restore_connect choice, so every value is tried.
	for _, policy := range []string{"", "active", "all", "off"} {
		t.Run("policy_"+policy, func(t *testing.T) {
			world, options := newWorld(t)
			first := newTestState(t, options)
			openHost(first)
			first.Click(hitHosts)
			first.Click(hitHostList.At(2))
			openHost(first)
			if len(world.conns) != 2 {
				t.Fatalf("setup dialed %d connections, want 2", len(world.conns))
			}
			first.save()

			saved, _, err := loadState(options.StatePath)
			if err != nil {
				t.Fatal(err)
			}
			saved.RestoreConnect = policy
			if err := saveState(options.StatePath, saved); err != nil {
				t.Fatal(err)
			}
			if reloaded, _, err := loadState(options.StatePath); err != nil || reloaded.RestoreConnect != policy {
				t.Fatalf("legacy restore_connect=%q did not survive the fixture: %q, %v", policy, reloaded.RestoreConnect, err)
			}

			restored := newTestState(t, options)
			restored.Wake()
			restored.Wake()
			if len(world.conns) != 2 {
				t.Fatalf("restart with restore_connect=%q dialed %d connections", policy, len(world.conns)-2)
			}
			for index := range restored.sessions {
				session := &restored.sessions[index]
				if session.conn != nil || !session.restored || session.phase != sshconn.PhaseClosed {
					t.Fatalf("restored tab %d was connected: %+v", session.ID, session)
				}
			}
		})
	}
}

func TestWorkspaceAndClosedTabsPersist(t *testing.T) {
	_, options := newWorld(t)
	first := newTestState(t, options)
	openHost(first)
	first.Click(hitHosts)
	first.Click(hitHostList.At(2))
	openHost(first)
	closed := first.sessions[0].Address
	kept := first.sessions[1].Address
	first.Click(hitTabCloseList.At(int(first.sessions[0].ID)))
	first.Click(hitKeychain)

	second := newTestState(t, options)
	if second.activeSection != sectionKeys {
		t.Fatalf("restored section = %d", second.activeSection)
	}
	if len(second.sessions) != 1 || second.sessions[0].Address != kept || second.sessions[0].Address == closed {
		t.Fatalf("restored sessions = %+v", second.sessions)
	}
	second.closeWorkspace(second.deck.ActiveWorkspace())
	third := newTestState(t, options)
	if len(third.sessions) != 0 || third.activeSection != sectionHosts {
		t.Fatalf("closed restored tab returned: sessions=%+v section=%d", third.sessions, third.activeSection)
	}
}

func TestTypedSessionRestoresAsAnEphemeralTarget(t *testing.T) {
	world, options := newWorld(t)
	first := newTestState(t, options)
	first.search.Insert("guest@typed.example.com:2222")
	first.Click(hitConnect)
	if len(world.conns) != 1 || len(first.sessions) != 1 || first.sessions[0].HostID != 0 {
		t.Fatalf("typed session = %+v dials=%d", first.sessions, len(world.conns))
	}

	second := newTestState(t, options)
	if len(world.conns) != 1 {
		t.Fatalf("restoring typed target dialed again: %d", len(world.conns))
	}
	if len(second.sessions) != 1 {
		t.Fatalf("restored typed sessions = %+v", second.sessions)
	}
	session := second.sessions[0]
	if session.HostID != 0 || session.host.User != "guest" || session.host.Address != "typed.example.com" || session.host.Port != 2222 {
		t.Fatalf("restored typed target = %+v", session)
	}
}

func TestRestoreSessionsSkipsInvalidStaleDuplicateAndExcessEntries(t *testing.T) {
	options := freshOptions(t)
	entries := []savedSession{
		{Address: ""},
		{Label: "Stale", Address: "stale.example.com", Port: 22, SavedHost: true},
		{Label: "One", Address: "one.example.com", Port: 22},
		{Label: "Duplicate", Address: "one.example.com", Port: 22},
	}
	for index := 0; index < maxRestoredSessions+4; index++ {
		entries = append(entries, savedSession{Label: fmt.Sprintf("Host %d", index), Address: fmt.Sprintf("host-%d.example.com", index), Port: 22})
	}
	if err := saveState(options.withDefaults().StatePath, savedState{Sessions: entries, Workspace: "terminal", ActiveSession: 1}); err != nil {
		t.Fatal(err)
	}
	state := newTestState(t, options)
	if len(state.sessions) != maxRestoredSessions {
		t.Fatalf("restored session count = %d", len(state.sessions))
	}
	if state.sessions[0].Address != state.systemUser+"@one.example.com:22" {
		t.Fatalf("first restored session = %+v", state.sessions[0])
	}
	if state.activeSessionID != state.sessions[0].ID || state.activeSection != sectionTerminal {
		t.Fatalf("active=%d section=%d", state.activeSessionID, state.activeSection)
	}
}

// A restored tab's page has no terminal to hold the keyboard: Enter on it
// reconnects, Command+W closes the tab, and both keys are shown on the page.
func TestRestoredPageTakesEnterAndCommandWAndShowsTheKeys(t *testing.T) {
	_, options := newWorld(t)
	first := newTestState(t, options)
	openHost(first)
	first.save()

	state := newTestState(t, options)
	session, _ := state.sessionByIDPointer(state.activeSessionID)
	if !session.restored {
		t.Fatal("setup: the session was not restored")
	}
	canvas, texts := paintedTexts(state)
	if !contains(texts, "↵") || !contains(texts, "W") {
		t.Fatalf("the page does not show its keys: %q", texts)
	}
	if _, ok := hitByID(canvas, hitConnRetry); !ok {
		t.Fatal("no Connect button")
	}

	if result := state.Key(ui.KeyEvent{Key: ui.KeyEnter}); !result.Handled || session.conn == nil {
		t.Fatalf("Enter on the restored page: %+v conn=%v", result, session.conn != nil)
	}

	second := newTestState(t, options)
	tabs := len(second.deck.Workspaces(nil))
	if result := commandW(second); !result.Handled || len(second.deck.Workspaces(nil)) != tabs-1 {
		t.Fatalf("Command+W on the restored page: %+v tabs %d→%d", result, tabs, len(second.deck.Workspaces(nil)))
	}
}

// Picking a restored tab leaves the keyboard on that tab, not on the page, and
// Enter there must still connect: it used to need a click on the page first.
func TestEnterOnTheActiveTabConnectsARestoredPage(t *testing.T) {
	_, options := newWorld(t)
	first := newTestState(t, options)
	openHost(first)
	first.save()

	state := newTestState(t, options)
	session, _ := state.sessionByIDPointer(state.activeSessionID)
	if !session.restored {
		t.Fatal("setup: the session was not restored")
	}
	other := hitTabList.At(int(state.deck.ActiveWorkspace()) + 1)
	if result := state.Key(ui.KeyEvent{Key: ui.KeyEnter, Target: other}); result.Handled && session.conn != nil {
		t.Fatal("Enter on some other tab must not connect this page")
	}
	active := hitTabList.At(int(state.deck.ActiveWorkspace()))
	if result := state.Key(ui.KeyEvent{Key: ui.KeyEnter, Target: active}); !result.Handled || session.conn == nil {
		t.Fatalf("Enter on the active tab: %+v conn=%v", result, session.conn != nil)
	}
}

func TestRestoredSessionIdlePathsDoNotAllocate(t *testing.T) {
	options := freshOptions(t)
	options = options.withDefaults()
	if err := saveState(options.StatePath, savedState{
		Workspace: "terminal",
		Sessions:  []savedSession{{Label: "Restored", Address: "restored.example.com", User: "user", Port: 22}},
	}); err != nil {
		t.Fatal(err)
	}
	state := newTestState(t, options)
	var canvas ui.Canvas
	size := ui.Size{Width: 1400, Height: 900}
	state.Paint(&canvas, size)
	if allocations := testing.AllocsPerRun(100, state.Wake); allocations != 0 {
		t.Fatalf("restored Wake allocations = %v", allocations)
	}
	if allocations := testing.AllocsPerRun(100, func() {
		canvas.Reset()
		state.Paint(&canvas, size)
	}); allocations != 0 {
		t.Fatalf("restored paint allocations = %v", allocations)
	}
}

func BenchmarkRestoreWorkspace(b *testing.B) {
	for _, count := range []int{0, 8} {
		b.Run(fmt.Sprintf("sessions_%d", count), func(b *testing.B) {
			options := freshOptions(b).withDefaults()
			saved := savedState{Workspace: "terminal", Sessions: make([]savedSession, count)}
			for index := range saved.Sessions {
				saved.Sessions[index] = savedSession{
					Label: fmt.Sprintf("Host %d", index), Address: fmt.Sprintf("host-%d.example.com", index),
					User: "deploy", Port: 22,
				}
			}
			if err := saveState(options.StatePath, saved); err != nil {
				b.Fatal(err)
			}
			b.ReportAllocs()
			b.ResetTimer()
			for b.Loop() {
				_ = newTestState(b, options)
			}
		})
	}
}
