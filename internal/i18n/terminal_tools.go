package i18n

// ToolsMessages are the texts of the terminal's own tools: color schemes,
// keyword highlighting, snippets and broadcast input.
type ToolsMessages struct {
	// Color schemes.
	SchemeGroup       string
	SchemeDescription string
	SchemeDefault     string

	// Keyword highlighting.
	HighlightGroup              string
	HighlightBuiltin            string
	HighlightBuiltinDescription string
	HighlightCustom             string
	HighlightCustomDescription  string
	HighlightAdd                string
	HighlightDelete             string
	HighlightKindText           string
	HighlightKindRegex          string
	HighlightKindTextCase       string
	HighlightKindRegexCase      string
	HighlightDialogTitle        string
	HighlightPattern            string
	HighlightPatternPlaceholder string
	HighlightRegex              string
	HighlightCase               string
	HighlightColor              string
	HighlightErrorEmpty         string
	HighlightErrorTooLong       string
	HighlightErrorEverywhere    string
	HighlightErrorInvalidFmt    string
	HighlightErrorTooMany       string
	HighlightSave               string
	ColorRed                    string
	ColorYellow                 string
	ColorGreen                  string
	ColorCyan                   string
	ColorBlue                   string
	ColorMagenta                string

	// Snippets.
	SnippetsGroup             string
	SnippetsRow               string
	SnippetsDescription       string
	SnippetAdd                string
	SnippetEdit               string
	SnippetDelete             string
	SnippetDialogAddTitle     string
	SnippetDialogEditTitle    string
	SnippetName               string
	SnippetNamePlaceholder    string
	SnippetCommand            string
	SnippetCommandPlaceholder string
	SnippetRun                string
	SnippetErrorName          string
	SnippetErrorCommand       string
	SnippetErrorTooMany       string
	SnippetSave               string
	SnippetPaletteFmt         string
	StatusSnippetSentFmt      string
	StatusSnippetNoTerminal   string
	StatusSnippetNotConnected string

	// The categories of the settings page, and what each holds.
	CatGeneral    string
	CatTerminal   string
	CatConnection string
	CatSecurity   string

	// Terminal zoom and padding.
	ZoomIn            string
	ZoomOut           string
	ZoomReset         string
	StatusFontSizeFmt string
	PaddingRow        string
	PaddingCompact    string
	PaddingNormal     string
	PaddingRoomy      string
	GutterRow         string
	GutterDescription string
	GutterOff         string
	GutterNumbers     string
	GutterTime        string
	GutterBoth        string
	CursorRow         string
	CursorDescription string

	// The terminal's font.
	FontChoose              string
	FontCurrentFmt          string
	FontMissingFmt          string
	FontPickerTitle         string
	FontPickerDescription   string
	FontPickerPlaceholder   string
	FontPickerLoading       string
	FontPickerEmpty         string
	FontPickerSample        string
	FontPickerSampleSymbols string

	// Notifications.
	NotifyRow         string
	NotifyDescription string

	// The outbound proxy.
	ProxyRow                 string
	ProxyConfigure           string
	ProxyDescriptionNone     string
	ProxyDescriptionFmt      string // the kind, then the address
	ProxyNone                string
	ProxySOCKS5              string
	ProxyHTTP                string
	ProxyDialogTitle         string
	ProxyDialogDescription   string
	ProxyAddress             string
	ProxyUser                string
	ProxyPassword            string
	ProxyPasswordPlaceholder string
	ProxyBypass              string
	ProxySave                string
	ProxyErrorAddress        string
	ProxyErrorLogin          string
	ProxyErrorBypassFmt      string
	ProxyErrorNoVault        string
	ProxyPasswordLocked      string

	// Searching the settings.
	SearchPlaceholder string
	SearchTitle       string
	SearchEmptyFmt    string

	// Importing sessions from other clients.
	ImportSessions            string
	ImportChooserTitle        string
	ImportChooserBody         string
	ImportPickerFmt           string // the client's name
	ImportCanceled            string
	ImportFailedFmt           string
	ImportSummaryFmt          string // hosts, client, duplicates
	ImportWarningKeysFmt      string
	ImportWarningSkippedFmt   string
	ImportWarningPasswords    string
	ImportWarningCleartextFmt string // sessions whose file held readable passwords or keys

	// Broadcast input.
	BroadcastStart            string
	BroadcastStop             string
	StatusBroadcastOnFmt      string
	StatusBroadcastOff        string
	StatusBroadcastNeedsSplit string

	// Telnet and serial hosts.
	ProtocolSSH            string
	ProtocolTelnet         string
	ProtocolSerial         string
	FieldDevice            string
	FieldDevicePlaceholder string
	FieldLine              string
	FieldLinePlaceholder   string
	FieldFlow              string
	FlowNone               string
	FlowHardware           string
	FlowSoftware           string
	PortsPicker            string
	PortsNone              string
	ErrSerialLine          string
	ErrSerialDevice        string
	TelnetPlaintext        string
	SerialOpenedFmt        string

	// File transfer over the terminal (ZMODEM).
	TransferSentFmt           string
	TransferReceivedFmt       string
	TransferCancelled         string
	TransferFailedFmt         string
	TransferDeclined          string
	TransferChooseFolderTitle string
	TransferNoDownloads       string
	TransferChooseFilesTitle  string
	TransferTooManyFmt        string
	TransferNotFileFmt        string
	TransferTooBigFmt         string
	TransferReceiveTitle      string
	TransferSendTitle         string
	TransferReceiveOfferFmt   string
	TransferSendOfferFmt      string
	TransferDecline           string
	TransferChooseFiles       string
	TransferChooseFolder      string
	TransferToDownloads       string

	// The file panel.
	CatFiles                        string
	FilesParallel                   string
	FilesKeepTimes                  string
	FilesKeepTimesDescription       string
	FilesFollow                     string
	FilesFollowDescription          string
	FilesFollowSetup                string
	FilesFollowSetupDescription     string
	FilesFollowSetupCopy            string
	FilesFollowSetupCopied          string
	TipNewLink                      string
	LinkTitle                       string
	LinkName                        string
	LinkTarget                      string
	LinkErrorName                   string
	LinkErrorTarget                 string
	LinkCreatingFmt                 string
	LinkCreatedFmt                  string
	LinkBroken                      string
	Bookmarks                       string
	BookmarksEmpty                  string
	BookmarkAdd                     string
	BookmarkRemove                  string
	BookmarkLimitFmt                string
	Extract                         string
	ExtractTitle                    string
	ExtractDescriptionFmt           string
	ExtractFolder                   string
	ExtractConfirm                  string
	ExtractErrorFolder              string
	ExtractRunningFmt               string
	ExtractDoneFmt                  string
	ExtractFailedFmt                string
	TransferPause                   string
	TransferResume                  string
	TransferPausedFmt               string
	SkipExisting                    string
	TransferSpeedFmt                string
	LinkCreate                      string
	PaneLocal                       string
	PaneRemote                      string
	LocalFiles                      string
	SendToServer                    string
	DownloadHere                    string
	XYMenu                          string
	XYTitle                         string
	XYDescription                   string
	XYProtocol                      string
	XYDirection                     string
	XYSend                          string
	XYReceive                       string
	XYStart                         string
	XYSaveTitle                     string
	DropUploadingFmt                string
	LaunchTitle                     string
	LaunchWarning                   string
	LaunchDescriptionFmt            string
	LaunchConnect                   string
	GeneralLogin                    string
	GeneralLoginDescription         string
	GeneralUpdate                   string
	GeneralUpdateDescription        string
	GeneralUpdateCheck              string
	GeneralUpdateDownload           string
	GeneralUpdateStartup            string
	GeneralUpdateStartupDescription string
	UpdateChecking                  string
	UpdateLatestFmt                 string
	UpdateAvailableFmt              string
	UpdateNoReleaseText             string
	ActionSSHFmt                    string
	ActionBrowserFmt                string
	ActionFilesFmt                  string
	ActionExtractFmt                string
	ActionCopyFmt                   string
	DiagnosticsRow                  string
	DiagnosticsDescription          string
	DiagnosticsExport               string
	DiagnosticsSaveTitle            string
	DiagnosticsSavedFmt             string
	DiagnosticsFailedFmt            string
	CatBackup                       string
	CatMCP                          string
	BackupGroupFile                 string
	BackupGroupSync                 string
	BackupExportRow                 string
	BackupExportDescription         string
	BackupExport                    string
	BackupRestore                   string
	BackupExportTitle               string
	BackupExportDialogDescription   string
	BackupPassword                  string
	BackupPasswordConfirm           string
	BackupPasswordPlaceholder       string
	BackupExportSave                string
	BackupSaveTitle                 string
	BackupSavedFmt                  string
	BackupFailedFmt                 string
	BackupRestoreTitle              string
	BackupRestoreDescription        string
	BackupRestoreConfirm            string
	BackupPickerTitle               string
	BackupRestoredFmt               string
	BackupErrPasswordShort          string
	BackupErrPasswordMismatch       string
	BackupErrWrong                  string
	BackupErrNotBackup              string
	BackupErrTooNew                 string
	BackupErrSessions               string
	BackupErrVault                  string
	BackupSecretsFailedFmt          string
	SyncRow                         string
	SyncRowNone                     string
	SyncRowFmt                      string
	SyncConfigure                   string
	SyncDialogTitle                 string
	SyncDialogDescription           string
	SyncURL                         string
	SyncUser                        string
	SyncPassword                    string
	SyncPath                        string
	SyncSave                        string
	SyncRunRow                      string
	SyncRunNever                    string
	SyncRunLastFmt                  string
	SyncNow                         string
	SyncAuto                        string
	SyncAutoDescription             string
	SyncRunFailedFmt                string
	SyncWorking                     string
	SyncDone                        string
	SyncErrNoStore                  string
	SyncTestOK                      string
	SyncFailedFmt                   string
	SyncErrUnauthorized             string
	SyncErrNotFound                 string
	SyncErrInsecure                 string
	SyncErrNeedsVault               string
	SyncErrFields                   string
	SyncErrAddress                  string
	SyncConfigureS3                 string
	SyncConfigureGist               string
	SyncErrTooLarge                 string
	S3DialogTitle                   string
	S3DialogDescription             string
	S3Endpoint                      string
	S3Region                        string
	S3Bucket                        string
	S3Key                           string
	S3PathStyle                     string
	S3LoginTitle                    string
	S3LoginDescription              string
	S3AccessKey                     string
	S3Secret                        string
	S3ErrFields                     string
	GistDialogTitle                 string
	GistDialogDescription           string
	GistToken                       string
	GistID                          string
	GistFile                        string
	GistAPI                         string
	GistErrFields                   string
	GistErrTooLarge                 string
	MenuSaveOutput                  string
	MenuSearchSelection             string
	OutputSaveTitle                 string
	OutputSavedFmt                  string
	OutputSaveFailedFmt             string
	OutputEmpty                     string
	CloseAppTitle                   string
	CloseAppDescriptionFmt          string
	CloseAppConfirm                 string
	RemoteMenu                      string
	RemoteTitle                     string
	RemoteDescriptionFmt            string
	RemoteTabProcesses              string
	RemoteTabContainers             string
	RemoteTabImages                 string
	RemoteTabVolumes                string
	RemoteTabNetworks               string
	RemoteTabCompose                string
	RemoteColConfig                 string
	RemoteEmptyCompose              string
	RemoteConfirmDownTitle          string
	RemoteConfirmDownFmt            string
	RemoteDown                      string
	RemoteErrNoCompose              string
	RemoteReniceLower               string
	RemoteReniceNormal              string
	RemoteReniceHigher              string
	RemoteReniceDoneFmt             string
	RemoteColSize                   string
	RemoteColCreated                string
	RemoteColDriver                 string
	RemoteColScope                  string
	RemoteEmptyImages               string
	RemoteEmptyVolumes              string
	RemoteEmptyNetworks             string
	RemoteConfirmRemoveImageTitle   string
	RemoteConfirmRemoveImageFmt     string
	RemoteConfirmRemoveVolumeTitle  string
	RemoteConfirmRemoveVolumeFmt    string
	RemoteConfirmRemoveNetworkTitle string
	RemoteConfirmRemoveNetworkFmt   string
	RemoteSortCPU                   string
	RemoteSortMemory                string
	RemoteColPID                    string
	RemoteColUser                   string
	RemoteColCPU                    string
	RemoteColMemory                 string
	RemoteColTime                   string
	RemoteColCommand                string
	RemoteColName                   string
	RemoteColImage                  string
	RemoteColState                  string
	RemoteColStatus                 string
	RemoteColID                     string
	RemoteRefresh                   string
	RemoteTerminate                 string
	RemoteForceKill                 string
	RemoteStart                     string
	RemoteStop                      string
	RemoteRestart                   string
	RemoteLogs                      string
	RemoteRemove                    string
	RemotePull                      string
	RemotePullTitle                 string
	RemotePullDescription           string
	RemoteImageField                string
	RemoteImagePlaceholder          string
	RemoteBadImage                  string
	RemoteCreate                    string
	RemoteCreateVolumeTitle         string
	RemoteCreateVolumeDescription   string
	RemoteCreateNetworkTitle        string
	RemoteCreateNetworkDescription  string
	RemoteNameField                 string
	RemoteBadName                   string
	RemoteLoading                   string
	RemoteLoadedFmt                 string
	RemoteEmptyContainers           string
	RemoteErrUnsupportedFmt         string
	RemoteErrNoDocker               string
	RemoteErrFailedFmt              string
	RemoteNeedSelection             string
	RemoteSignalSentFmt             string
	RemoteActionDoneFmt             string
	RemoteConfirmTerminateTitle     string
	RemoteConfirmKillTitle          string
	RemoteConfirmProcessFmt         string
	RemoteConfirmRemoveTitle        string
	RemoteConfirmRemoveFmt          string
	RemoteLogsFmt                   string
	RemoteClose                     string
	ExtEditButton                   string
	ExtEditOpenedFmt                string
	ExtEditSentFmt                  string
	ExtEditSendFailedFmt            string
	ExtEditBadFmt                   string
	ExtEditRefused                  string
	ExtEditOpenFailedFmt            string
	ExtEditConflictTitle            string
	ExtEditConflictDescriptionFmt   string
	ExtEditReplace                  string
	HostViewList                    string
	HostViewGrid                    string
	BgGroup                         string
	BgRow                           string
	BgRowNone                       string
	BgRowSetFmt                     string
	BgChoose                        string
	BgRemove                        string
	BgPickerTitle                   string
	BgVisibility                    string
	BgSubtle                        string
	BgMedium                        string
	BgStrong                        string
	BgLoading                       string
	BgErrFmt                        string
	BgErrTooLarge                   string
	BgErrTooManyPixels              string
	BgErrFormat                     string
	TabMoveBusy                     string
	OTPField                        string
	OTPPlaceholder                  string
	OTPUncheckedPlaceholder         string
	OTPSavedPlaceholder             string
	OTPRemovingPlaceholder          string
	OTPCheck                        string
	OTPRemove                       string
	OTPScan                         string
	OTPPaste                        string
	ErrOTPPasteNoImage              string
	OTPScanTitle                    string
	OTPScanFilter                   string
	OTPScannedFmt                   string
	ErrOTPScanTooLarge              string
	ErrOTPScanDimension             string
	ErrOTPScanNotImage              string
	ErrOTPScanNoCode                string
	ErrOTPScanNotOTP                string
	ErrOTPScanMigration             string
	ErrOTPScanManyFmt               string
	ErrOTPScanReadFmt               string
	X11Forward                      string
	X11UnavailableFmt               string
	AgentForward                    string
	AgentForwardUnavailableFmt      string
	OTPKeep                         string
	ErrOTPInvalid                   string
	ErrOTPNoVault                   string
	ErrOTPLocked                    string
	ErrOTPCheckFmt                  string
	OTPSavedFmt                     string
	OTPRemovedFmt                   string
	AlgoField                       string
	AdvancedOptions                 string
	TransferField                   string
	TransferAuto                    string
	AlgoDefault                     string
	AlgoCompatible                  string
	AlgoCustom                      string
	AlgoCustomField                 string
	AlgoCustomPlaceholder           string
	ErrAlgoRequired                 string
	ErrAlgoUnknownFmt               string
	AlgoWeakFmt                     string
	AlgoNegotiatedFmt               string
	UpdateFailedFmt                 string
	LoginFailedFmt                  string
}

var englishTools = ToolsMessages{
	SchemeGroup:       "Terminal colors",
	SchemeDescription: "The terminal's background, text and the sixteen named colors. The window around it keeps its own theme.",
	SchemeDefault:     "Default",

	HighlightGroup:              "Keyword highlighting",
	HighlightBuiltin:            "Common words",
	HighlightBuiltinDescription: "Errors, warnings, successes and IP addresses get a color. Text a program already colors is left alone.",
	HighlightCustom:             "Your rules",
	HighlightCustomDescription:  "Words or patterns of your own, colored wherever they appear in the terminal.",
	HighlightAdd:                "Add rule",
	HighlightDelete:             "Delete",
	HighlightKindText:           "text",
	HighlightKindRegex:          "regular expression",
	HighlightKindTextCase:       "text, match case",
	HighlightKindRegexCase:      "regular expression, match case",
	HighlightDialogTitle:        "Add a highlight rule",
	HighlightPattern:            "Text to find",
	HighlightPatternPlaceholder: "for example: ERROR or 5\\d\\d",
	HighlightRegex:              "Regular expression",
	HighlightCase:               "Match case",
	HighlightColor:              "Color",
	HighlightErrorEmpty:         "Enter the text to find.",
	HighlightErrorTooLong:       "That is too long for a rule.",
	HighlightErrorEverywhere:    "That would match every line; use something more specific.",
	HighlightErrorInvalidFmt:    "Not a valid pattern: %s",
	HighlightErrorTooMany:       "There are already as many rules as can be used. Delete one first.",
	HighlightSave:               "Add",
	ColorRed:                    "Red",
	ColorYellow:                 "Yellow",
	ColorGreen:                  "Green",
	ColorCyan:                   "Cyan",
	ColorBlue:                   "Blue",
	ColorMagenta:                "Magenta",

	SnippetsGroup:             "Snippets",
	SnippetsRow:               "Quick commands",
	SnippetsDescription:       "Sent to the active terminal from the command palette.",
	SnippetAdd:                "Add snippet",
	SnippetEdit:               "Edit",
	SnippetDelete:             "Delete",
	SnippetDialogAddTitle:     "Add a snippet",
	SnippetDialogEditTitle:    "Edit snippet",
	SnippetName:               "Name",
	SnippetNamePlaceholder:    "for example: Restart nginx",
	SnippetCommand:            "Command",
	SnippetCommandPlaceholder: "systemctl restart nginx",
	SnippetRun:                "Press Enter after sending it",
	SnippetErrorName:          "Give the snippet a name.",
	SnippetErrorCommand:       "Enter the command to send.",
	SnippetErrorTooMany:       "There are already as many snippets as can be kept. Delete one first.",
	SnippetSave:               "Save",
	SnippetPaletteFmt:         "Snippet: %s",
	StatusSnippetSentFmt:      "Sent “%s”.",
	StatusSnippetNoTerminal:   "Open a terminal to send a snippet to.",
	StatusSnippetNotConnected: "That terminal is not connected.",

	CatGeneral:    "General",
	CatTerminal:   "Terminal",
	CatConnection: "Connection",
	CatSecurity:   "Security",

	ZoomIn:            "Zoom in terminal text",
	ZoomOut:           "Zoom out terminal text",
	ZoomReset:         "Reset terminal text size",
	StatusFontSizeFmt: "Terminal font size %d.",
	PaddingRow:        "Terminal padding",
	PaddingCompact:    "Compact",
	PaddingNormal:     "Normal",
	PaddingRoomy:      "Roomy",
	GutterRow:         "Line gutter",
	GutterDescription: "Shows line numbers or write times.",
	GutterOff:         "Off",
	GutterNumbers:     "Numbers",
	GutterTime:        "Time",
	GutterBoth:        "Both",
	CursorRow:         "Cursor blink",
	CursorDescription: "Solid again when idle, or when a program asks for a steady cursor.",

	FontChoose:              "Choose",
	FontCurrentFmt:          "Now: %s. Any installed fixed-pitch font can be used.",
	FontMissingFmt:          "The font “%s” is no longer installed; using the system font.",
	FontPickerTitle:         "Terminal font",
	FontPickerDescription:   "Choose the system fonts or any installed fixed-pitch font.",
	FontPickerPlaceholder:   "Search fonts",
	FontPickerLoading:       "Reading the installed fonts…",
	FontPickerEmpty:         "No font matches.",
	FontPickerSample:        "The quick brown fox jumps over the lazy dog",
	FontPickerSampleSymbols: "0123456789 {}[]() <> -> == != ~ $ # | iIl1 oO0",

	NotifyRow:         "Desktop notifications",
	NotifyDescription: "Show a system notification when a program in the terminal asks for one (OSC 9, 777 and 99). At most one a second per terminal.",

	ProxyRow:                 "Outbound proxy",
	ProxyConfigure:           "Configure",
	ProxyDescriptionNone:     "Connect directly. A SOCKS5 or HTTP proxy can carry the first connection of every session.",
	ProxyDescriptionFmt:      "%s at %s carries the first connection of every session.",
	ProxyNone:                "None",
	ProxySOCKS5:              "SOCKS5",
	ProxyHTTP:                "HTTP CONNECT",
	ProxyDialogTitle:         "Outbound proxy",
	ProxyDialogDescription:   "The proxy resolves host names. Destinations on the direct list skip it.",
	ProxyAddress:             "Proxy address (host:port)",
	ProxyUser:                "User name (optional)",
	ProxyPassword:            "Password (optional)",
	ProxyPasswordPlaceholder: "Leave empty to keep the saved one",
	ProxyBypass:              "Connect directly to",
	ProxySave:                "Save",
	ProxyErrorAddress:        "Enter the proxy as host:port, with a port from 1 to 65535.",
	ProxyErrorLogin:          "The user name and password may be 255 characters at most.",
	ProxyErrorBypassFmt:      "The direct list takes up to %d entries of 100 characters.",
	ProxyErrorNoVault:        "Passwords cannot be saved here. Leave the password empty to connect without one.",
	ProxyPasswordLocked:      "The saved proxy password is locked: unlock the saved passwords in Settings › Security first.",

	SearchPlaceholder: "Search settings",
	SearchTitle:       "Search results",
	SearchEmptyFmt:    "No setting matches “%s”.",

	ImportSessions:            "Import from",
	ImportChooserTitle:        "Import sessions",
	ImportChooserBody:         "Choose a client, then pick its saved session files.",
	ImportPickerFmt:           "Import %s sessions",
	ImportCanceled:            "Import canceled.",
	ImportFailedFmt:           "Could not import: %v",
	ImportSummaryFmt:          "Imported %d hosts from %s; skipped %d already saved.",
	ImportWarningKeysFmt:      "%d sessions use a private key file. Keys are not imported; add them under Keychain.",
	ImportWarningSkippedFmt:   "%d sessions are of a kind that is not imported (remote desktop, file transfer, serial ports from this client…) and were skipped.",
	ImportWarningPasswords:    "Saved passwords are not imported: each program encrypts them its own way.",
	ImportWarningCleartextFmt: "%d sessions keep a password or private key in this file in clear text. None of it was imported; delete the file when you are done.",

	BroadcastStart:            "Type into every pane of this tab",
	BroadcastStop:             "Stop typing into every pane",
	StatusBroadcastOnFmt:      "What you type now goes to all %d panes of this tab.",
	StatusBroadcastOff:        "Broadcast is off.",
	StatusBroadcastNeedsSplit: "Split the tab first: broadcast types into every pane of it.",

	ProtocolSSH:            "SSH",
	ProtocolTelnet:         "Telnet",
	ProtocolSerial:         "Serial",
	FieldDevice:            "Device",
	FieldDevicePlaceholder: "/dev/ttyUSB0 or COM3",
	FieldLine:              "Baud rate and frame",
	FieldLinePlaceholder:   "115200 8N1",
	FieldFlow:              "Flow control",
	FlowNone:               "None",
	FlowHardware:           "RTS/CTS",
	FlowSoftware:           "XON/XOFF",
	PortsPicker:            "Detected ports",
	PortsNone:              "No serial ports found",
	ErrSerialLine:          "Write the line as a baud rate and a frame, like 115200 8N1 (5 to 8 data bits, parity N, O, E, M or S, 1 or 2 stop bits).",
	ErrSerialDevice:        "The device name has characters a serial port cannot have.",
	TelnetPlaintext:        "Telnet is not encrypted: what you type, passwords included, crosses the network in the clear.",
	SerialOpenedFmt:        "Opened %s at %s.",

	TransferSentFmt:           "Sent %d files.",
	TransferReceivedFmt:       "Received %d files into %s.",
	TransferCancelled:         "File transfer cancelled.",
	TransferFailedFmt:         "File transfer failed: %v",
	TransferDeclined:          "File transfer declined.",
	TransferChooseFolderTitle: "Save received files in",
	TransferNoDownloads:       "There is no Downloads folder. Choose a folder instead.",
	TransferChooseFilesTitle:  "Files to send",
	TransferTooManyFmt:        "At most %d files can be sent at once.",
	TransferNotFileFmt:        "%s is not a regular file.",
	TransferTooBigFmt:         "%s is larger than 4 GiB, the most this protocol can send.",
	TransferReceiveTitle:      "Receive files",
	TransferSendTitle:         "Send files",
	TransferReceiveOfferFmt:   "%s wants to send you files (ZMODEM). Nothing is saved until you choose where.",
	TransferSendOfferFmt:      "%s is ready to receive files (ZMODEM). Choose the files to send.",
	TransferDecline:           "Decline",
	TransferChooseFiles:       "Choose files",
	TransferChooseFolder:      "Choose folder",
	TransferToDownloads:       "Save to Downloads",

	CatFiles:                        "Files",
	FilesParallel:                   "Parallel transfers",
	FilesKeepTimes:                  "Keep modification times",
	FilesKeepTimesDescription:       "A copied file gets the time of the original. Permissions are always kept.",
	FilesFollow:                     "Follow the terminal's folder",
	FilesFollowDescription:          "The file panel goes to the folder the shell reports (OSC 7). The shell must be set up to send it.",
	FilesFollowSetup:                "Folder reporting for bash and zsh",
	FilesFollowSetupDescription:     "Makes the prompt report its folder. Paste it into the remote shell, or its startup file.",
	FilesFollowSetupCopy:            "Copy command",
	FilesFollowSetupCopied:          "Copied. Paste it into the remote shell.",
	TipNewLink:                      "New symbolic link",
	LinkTitle:                       "New symbolic link",
	LinkName:                        "Link name",
	LinkTarget:                      "Points to",
	LinkErrorName:                   "Enter a name for the link.",
	LinkErrorTarget:                 "Enter what the link points to.",
	LinkCreatingFmt:                 "Linking %s to %s…",
	LinkCreatedFmt:                  "Linked %s to %s.",
	LinkBroken:                      "missing",
	Bookmarks:                       "Bookmarks",
	BookmarksEmpty:                  "No bookmarked folders yet. The star bookmarks the folder you are in.",
	BookmarkAdd:                     "Bookmark this folder",
	BookmarkRemove:                  "Remove bookmark",
	BookmarkLimitFmt:                "A server can have %d bookmarks.",
	Extract:                         "Extract here",
	ExtractTitle:                    "Extract archive",
	ExtractDescriptionFmt:           "Unpack %s into a new folder. Nothing is overwritten.",
	ExtractFolder:                   "Folder",
	ExtractConfirm:                  "Extract",
	ExtractErrorFolder:              "Enter a folder name without slashes.",
	ExtractRunningFmt:               "Extracting %s…",
	ExtractDoneFmt:                  "Extracted %s into %s.",
	ExtractFailedFmt:                "Could not extract %s: %s",
	TransferPause:                   "Pause",
	TransferResume:                  "Resume",
	TransferPausedFmt:               "Paused · %s of %s",
	SkipExisting:                    "Skip existing",
	TransferSpeedFmt:                "%s/s",
	LinkCreate:                      "Create link",
	PaneLocal:                       "Show this computer's files",
	PaneRemote:                      "Show the server's files",
	LocalFiles:                      "This computer",
	SendToServer:                    "Upload to the server's folder",
	DownloadHere:                    "Download to the folder of this computer on show",
	XYMenu:                          "File Transfer",
	XYTitle:                         "File transfer",
	XYDescription:                   "Start the program on the other side first (rx or rb to receive, sx or sb to send), then press Start.",
	XYProtocol:                      "Protocol",
	XYDirection:                     "Direction",
	XYSend:                          "Send",
	XYReceive:                       "Receive",
	XYStart:                         "Start",
	XYSaveTitle:                     "Save the received file as",
	DropUploadingFmt:                "Uploading %d dropped items to %s",
	LaunchTitle:                     "Open connection",
	LaunchDescriptionFmt:            "Connect to %[2]s over %[1]s?",
	LaunchWarning:                   "Another program asked for this connection. Connect only if you expect it.",
	LaunchConnect:                   "Connect",
	GeneralLogin:                    "Start at login",
	GeneralLoginDescription:         "Start Motty in the background after you sign in.",
	GeneralUpdate:                   "Check for updates",
	GeneralUpdateDescription:        "Looks for a newer version on GitHub and opens its download page. Nothing is installed for you.",
	GeneralUpdateCheck:              "Check now",
	GeneralUpdateDownload:           "Download",
	GeneralUpdateStartup:            "Check at startup",
	GeneralUpdateStartupDescription: "Look for a newer version each time Motty starts. A check lets GitHub see your address, so it is off by default.",
	UpdateChecking:                  "Checking for updates…",
	UpdateLatestFmt:                 "Motty %s is the newest version.",
	UpdateAvailableFmt:              "Motty %s is available.",
	UpdateNoReleaseText:             "No published release was found yet.",
	ActionSSHFmt:                    "SSH to %s",
	ActionBrowserFmt:                "Open %s in browser",
	ActionFilesFmt:                  "Show %s in Files",
	ActionExtractFmt:                "Unpack %s",
	ActionCopyFmt:                   "Copy %s",
	DiagnosticsRow:                  "Diagnostics report",
	DiagnosticsDescription:          "Save the version, settings and recent log as text for a bug report. Hosts and accounts are hidden.",
	DiagnosticsExport:               "Export",
	DiagnosticsSaveTitle:            "Save diagnostics report",
	DiagnosticsSavedFmt:             "Saved diagnostics report to %s",
	DiagnosticsFailedFmt:            "Could not save the diagnostics report: %v",
	CatBackup:                       "Backup & sync",
	CatMCP:                          "MCP",
	BackupGroupFile:                 "Backup file",
	BackupGroupSync:                 "Cloud sync",
	BackupExportRow:                 "Back up everything",
	BackupExportDescription:         "Hosts, settings, trusted host keys and saved passwords in one file, encrypted with a password you choose.",
	BackupExport:                    "Export",
	BackupRestore:                   "Restore",
	BackupExportTitle:               "Back up to a file",
	BackupExportDialogDescription:   "Choose the password that opens the backup. It cannot be recovered if you lose it.",
	BackupPassword:                  "Backup password",
	BackupPasswordConfirm:           "Confirm password",
	BackupPasswordPlaceholder:       "At least 8 characters",
	BackupExportSave:                "Choose file",
	BackupSaveTitle:                 "Save backup",
	BackupSavedFmt:                  "Backup saved to %s.",
	BackupFailedFmt:                 "Backup failed: %s",
	BackupRestoreTitle:              "Restore a backup",
	BackupRestoreDescription:        "This replaces the hosts and settings on this computer with the backup's. The current ones are kept in config.db.before-restore.",
	BackupRestoreConfirm:            "Restore",
	BackupPickerTitle:               "Choose a backup",
	BackupRestoredFmt:               "Restored the backup (%d hosts).",
	BackupErrPasswordShort:          "Use at least 8 characters.",
	BackupErrPasswordMismatch:       "The two passwords differ.",
	BackupErrWrong:                  "Wrong password, or the file was changed.",
	BackupErrNotBackup:              "That is not a Motty backup.",
	BackupErrTooNew:                 "The backup was made by a newer version of Motty.",
	BackupErrSessions:               "Close all sessions first: restoring replaces the hosts they belong to.",
	BackupErrVault:                  "Unlock the saved passwords first.",
	BackupSecretsFailedFmt:          "Saved passwords could not be restored: %s",
	SyncRow:                         "Server",
	SyncRowNone:                     "No sync service set. Keep the backup on WebDAV, S3 or a GitHub Gist.",
	SyncRowFmt:                      "The backup is kept at %s.",
	SyncConfigure:                   "WebDAV",
	SyncDialogTitle:                 "WebDAV server",
	SyncDialogDescription:           "The backup is encrypted before it leaves this computer; the server only sees the encrypted file. Leave the address empty to stop syncing.",
	SyncURL:                         "Server address",
	SyncUser:                        "User",
	SyncPassword:                    "Password",
	SyncPath:                        "File",
	SyncSave:                        "Save",
	SyncRunRow:                      "Sync",
	SyncRunNever:                    "Not synced yet. Syncing merges this computer's hosts and settings with the server's.",
	SyncRunLastFmt:                  "Last synced %s. Syncing merges this computer's hosts and settings with the server's.",
	SyncNow:                         "Sync now",
	SyncAuto:                        "Sync automatically",
	SyncAutoDescription:             "A few seconds after hosts, settings or saved passwords change here, and when the application starts. Other computers' changes arrive the next time this one syncs.",
	SyncRunFailedFmt:                "The last sync failed: %s",
	SyncWorking:                     "Talking to the server…",
	SyncDone:                        "Synced with the server.",
	SyncErrNoStore:                  "The settings database could not be opened, so there is nothing to sync.",
	SyncTestOK:                      "Connected to the server.",
	SyncFailedFmt:                   "Sync failed: %s",
	SyncErrUnauthorized:             "The service refused the login, or the login may not do this.",
	SyncErrNotFound:                 "There is no backup on the server yet.",
	SyncErrInsecure:                 "Use an https:// address (http:// only for this computer).",
	SyncErrNeedsVault:               "Saved passwords are needed to keep the sync login, and they are not available.",
	SyncErrFields:                   "Enter the address, user, password, file and backup password.",
	SyncErrAddress:                  "That is not a server address.",
	SyncConfigureS3:                 "S3",
	SyncConfigureGist:               "Gist",
	SyncErrTooLarge:                 "The backup is too large for this service.",
	S3DialogTitle:                   "S3-compatible storage (1/2)",
	S3DialogDescription:             "AWS S3, Cloudflare R2, MinIO, Backblaze B2 and others. The backup is encrypted before it leaves this computer. Leave the endpoint empty to stop syncing.",
	S3Endpoint:                      "Endpoint",
	S3Region:                        "Region",
	S3Bucket:                        "Bucket",
	S3Key:                           "Object key",
	S3PathStyle:                     "Path-style addressing (yes or no)",
	S3LoginTitle:                    "S3 access key (2/2)",
	S3LoginDescription:              "Use a key that may read and write this one object. The secret is kept in the saved passwords, never in the settings file. Leave it empty to keep the saved one.",
	S3AccessKey:                     "Access key ID",
	S3Secret:                        "Secret access key",
	S3ErrFields:                     "Enter the endpoint, bucket, access key, secret and backup password.",
	GistDialogTitle:                 "GitHub Gist",
	GistDialogDescription:           "The encrypted backup is kept in a secret gist. Use a token with only the Gists permission (read and write); it is kept in the saved passwords. Leave the gist empty to create one on the first upload.",
	GistToken:                       "Access token",
	GistID:                          "Gist ID (optional)",
	GistFile:                        "File name",
	GistAPI:                         "API address (optional)",
	GistErrFields:                   "Enter the access token and the backup password.",
	GistErrTooLarge:                 "A gist file holds at most 1 MB, and this backup is too large for it.",
	MenuSaveOutput:                  "Save output as text",
	MenuSearchSelection:             "Search the web for selection",
	OutputSaveTitle:                 "Save terminal output",
	OutputSavedFmt:                  "Saved the terminal output to %s.",
	OutputSaveFailedFmt:             "Could not save the terminal output: %s",
	OutputEmpty:                     "The terminal has no output to save.",
	CloseAppTitle:                   "Close Motty?",
	CloseAppDescriptionFmt:          "%d sessions are open and will be disconnected.",
	CloseAppConfirm:                 "Close",
	RemoteMenu:                      "Processes & containers",
	RemoteTitle:                     "Remote tools",
	RemoteDescriptionFmt:            "Processes and Docker containers on %s. Nothing refreshes by itself.",
	RemoteTabProcesses:              "Processes",
	RemoteTabImages:                 "Images",
	RemoteTabVolumes:                "Volumes",
	RemoteTabNetworks:               "Networks",
	RemoteTabCompose:                "Compose",
	RemoteColConfig:                 "Config files",
	RemoteEmptyCompose:              "There are no compose projects.",
	RemoteConfirmDownTitle:          "Take this project down?",
	RemoteConfirmDownFmt:            "Project %s will be stopped and its containers and networks removed (volumes are kept).",
	RemoteDown:                      "Down",
	RemoteErrNoCompose:              "This server has no docker compose plugin.",
	RemoteReniceLower:               "Nice +10",
	RemoteReniceNormal:              "Nice 0",
	RemoteReniceHigher:              "Nice -5",
	RemoteReniceDoneFmt:             "Set the niceness of process %d to %d.",
	RemoteColSize:                   "Size",
	RemoteColCreated:                "Created",
	RemoteColDriver:                 "Driver",
	RemoteColScope:                  "Scope",
	RemoteEmptyImages:               "There are no images.",
	RemoteEmptyVolumes:              "There are no volumes.",
	RemoteEmptyNetworks:             "There are no networks.",
	RemoteConfirmRemoveImageTitle:   "Remove this image?",
	RemoteConfirmRemoveImageFmt:     "%s will be removed. Docker refuses it while a container uses it.",
	RemoteConfirmRemoveVolumeTitle:  "Remove this volume?",
	RemoteConfirmRemoveVolumeFmt:    "Volume %s and the data in it will be removed. Docker refuses it while a container uses it.",
	RemoteConfirmRemoveNetworkTitle: "Remove this network?",
	RemoteConfirmRemoveNetworkFmt:   "Network %s will be removed. Docker refuses it while a container uses it.",
	RemoteTabContainers:             "Containers",
	RemoteSortCPU:                   "By CPU",
	RemoteSortMemory:                "By memory",
	RemoteColPID:                    "PID",
	RemoteColUser:                   "User",
	RemoteColCPU:                    "CPU %",
	RemoteColMemory:                 "Memory %",
	RemoteColTime:                   "Running",
	RemoteColCommand:                "Command",
	RemoteColName:                   "Name",
	RemoteColImage:                  "Image",
	RemoteColState:                  "State",
	RemoteColStatus:                 "Status",
	RemoteColID:                     "ID",
	RemoteRefresh:                   "Refresh",
	RemoteTerminate:                 "Terminate",
	RemoteForceKill:                 "Force kill",
	RemoteStart:                     "Start",
	RemoteStop:                      "Stop",
	RemoteRestart:                   "Restart",
	RemoteLogs:                      "Logs",
	RemoteRemove:                    "Remove",
	RemotePull:                      "Pull image",
	RemotePullTitle:                 "Pull an image",
	RemotePullDescription:           "Runs docker pull in the terminal, where its progress shows and Ctrl+C stops it.",
	RemoteImageField:                "Image",
	RemoteImagePlaceholder:          "e.g. nginx:latest",
	RemoteBadImage:                  "That is not an image reference. Use a name such as nginx:latest or ghcr.io/org/app:1.2.",
	RemoteCreate:                    "New",
	RemoteCreateVolumeTitle:         "New volume",
	RemoteCreateVolumeDescription:   "Creates an empty volume on the server.",
	RemoteCreateNetworkTitle:        "New network",
	RemoteCreateNetworkDescription:  "Creates a bridge network on the server.",
	RemoteNameField:                 "Name",
	RemoteBadName:                   "Use letters, digits, underscore, dot and hyphen (up to 128), starting with a letter or digit.",
	RemoteLoading:                   "Reading from the server…",
	RemoteLoadedFmt:                 "%d entries.",
	RemoteEmptyContainers:           "There are no containers.",
	RemoteErrUnsupportedFmt:         "This server's ps is not supported: %s",
	RemoteErrNoDocker:               "Docker is not installed on this server.",
	RemoteErrFailedFmt:              "Failed: %s",
	RemoteNeedSelection:             "Select a row first.",
	RemoteSignalSentFmt:             "Sent %[1]s to process %[2]d.",
	RemoteActionDoneFmt:             "Done: %s.",
	RemoteConfirmTerminateTitle:     "Terminate this process?",
	RemoteConfirmKillTitle:          "Force kill this process?",
	RemoteConfirmProcessFmt:         "Process %d (%s). Unsaved work in it is lost.",
	RemoteConfirmRemoveTitle:        "Remove this container?",
	RemoteConfirmRemoveFmt:          "%s will be removed, with the data in its writable layer.",
	RemoteLogsFmt:                   "Following the log of %s in the terminal.",
	RemoteClose:                     "Close",
	ExtEditButton:                   "Edit in the system editor",
	ExtEditOpenedFmt:                "Opened %s in the system editor. Save it there and it is sent back.",
	ExtEditSentFmt:                  "Sent %s back to the server.",
	ExtEditSendFailedFmt:            "Could not send %s back: %s",
	ExtEditBadFmt:                   "%s is not UTF-8 text now, so it was not sent back.",
	ExtEditRefused:                  "This kind of file would be run, not edited. Use the built-in editor.",
	ExtEditOpenFailedFmt:            "Could not open %s in the system editor: %s",
	ExtEditConflictTitle:            "The server's file changed",
	ExtEditConflictDescriptionFmt:   "%s changed on the server after you opened it. Replace the server's version with yours?",
	ExtEditReplace:                  "Replace",
	HostViewList:                    "List view",
	HostViewGrid:                    "Card view",
	BgGroup:                         "Terminal background",
	BgRow:                           "Terminal background image",
	BgRowNone:                       "A picture behind the terminal text. Only the terminals turn translucent.",
	BgRowSetFmt:                     "%s, %d×%d. Shown cover-fit behind the terminals.",
	BgChoose:                        "Choose",
	BgRemove:                        "Remove",
	BgPickerTitle:                   "Choose a background image",
	BgVisibility:                    "Image visibility",
	BgSubtle:                        "Subtle",
	BgMedium:                        "Medium",
	BgStrong:                        "Strong",
	BgLoading:                       "Loading the background image…",
	BgErrFmt:                        "Could not use that image: %s",
	BgErrTooLarge:                   "The image file is larger than 30 MB.",
	BgErrTooManyPixels:              "The image has more than 64 million pixels.",
	BgErrFormat:                     "Only PNG and JPEG images are supported.",
	TabMoveBusy:                     "A file transfer is running in that tab, so it cannot move to another window yet.",
	OTPField:                        "One-time code secret",
	OTPPlaceholder:                  "Optional: otpauth:// address or base32 secret",
	OTPUncheckedPlaceholder:         "Leave blank to keep any saved secret",
	OTPSavedPlaceholder:             "Saved. Type a new secret to replace it",
	OTPRemovingPlaceholder:          "It will be removed when you save",
	OTPCheck:                        "Check",
	OTPRemove:                       "Remove",
	OTPScan:                         "Scan QR image",
	OTPPaste:                        "Paste image",
	ErrOTPPasteNoImage:              "No image on the clipboard. Copy the QR code picture first",
	OTPScanTitle:                    "Choose an image of the QR code",
	OTPScanFilter:                   "Images (PNG, JPEG, GIF)",
	OTPScannedFmt:                   "QR code read: %s. Save the host to keep it.",
	ErrOTPScanTooLarge:              "The image is larger than 10 MB. Crop it to the QR code and try again",
	ErrOTPScanDimension:             "The image has too many pixels. Scale it down or crop it to the QR code",
	ErrOTPScanNotImage:              "That file is not a PNG, JPEG or GIF image",
	ErrOTPScanNoCode:                "No QR code found in the image. Use a sharper, larger picture of the whole code",
	ErrOTPScanNotOTP:                "The QR code is not a one-time code (otpauth://) address",
	ErrOTPScanMigration:             "The authenticator export QR code is damaged or uses an algorithm that is not supported (MD5). Export again or add the account by its own QR code",
	ErrOTPScanManyFmt:               "The export QR code holds %d accounts and none is named like this host's user. Export a single account from your authenticator and scan that",
	ErrOTPScanReadFmt:               "Could not read the image: %v",
	X11Forward:                      "Forward X11 (session-only cookie)",
	X11UnavailableFmt:               "X11 forwarding is off for this session: %s",
	AgentForward:                    "Forward ssh-agent",
	AgentForwardUnavailableFmt:      "ssh-agent forwarding is off for this session: %s",
	OTPKeep:                         "Keep",
	ErrOTPInvalid:                   "Not a valid one-time code secret",
	ErrOTPNoVault:                   "Saved passwords are not available, so a secret cannot be kept",
	ErrOTPLocked:                    "Unlock the saved passwords in Settings first",
	ErrOTPCheckFmt:                  "Could not check saved one-time code secrets: %v",
	OTPSavedFmt:                     "One-time code secret saved for %s",
	OTPRemovedFmt:                   "One-time code secret removed for %s",
	AlgoField:                       "SSH algorithms",
	AdvancedOptions:                 "Advanced",
	TransferField:                   "File transfer",
	TransferAuto:                    "Auto",
	AlgoDefault:                     "Default",
	AlgoCompatible:                  "Compatible",
	AlgoCustom:                      "Custom",
	AlgoCustomField:                 "Allowed algorithms",
	AlgoCustomPlaceholder:           "aes256-gcm@openssh.com, hmac-sha2-256, …",
	ErrAlgoRequired:                 "Custom needs a list of algorithms",
	ErrAlgoUnknownFmt:               "Not an algorithm this program implements: %s",
	AlgoWeakFmt:                     "This connection negotiated weak algorithms: %s",
	AlgoNegotiatedFmt:               "algorithms %s",
	UpdateFailedFmt:                 "Could not check for updates: %v",
	LoginFailedFmt:                  "Could not change starting at login: %v",
}

var chineseTools = ToolsMessages{
	SchemeGroup:       "终端配色",
	SchemeDescription: "终端的背景、文字和十六种命名颜色。窗口其余部分仍使用自己的主题。",
	SchemeDefault:     "默认",

	HighlightGroup:              "关键字高亮",
	HighlightBuiltin:            "常见词",
	HighlightBuiltinDescription: "错误、警告、成功字样和 IP 地址会着色；程序自己已经上色的文字保持不变。",
	HighlightCustom:             "自定义规则",
	HighlightCustomDescription:  "你自己的词或模式，在终端里出现的地方都会着色。",
	HighlightAdd:                "添加规则",
	HighlightDelete:             "删除",
	HighlightKindText:           "文本",
	HighlightKindRegex:          "正则表达式",
	HighlightKindTextCase:       "文本，区分大小写",
	HighlightKindRegexCase:      "正则表达式，区分大小写",
	HighlightDialogTitle:        "添加高亮规则",
	HighlightPattern:            "要查找的文字",
	HighlightPatternPlaceholder: "例如：ERROR 或 5\\d\\d",
	HighlightRegex:              "正则表达式",
	HighlightCase:               "区分大小写",
	HighlightColor:              "颜色",
	HighlightErrorEmpty:         "请输入要查找的文字。",
	HighlightErrorTooLong:       "规则太长了。",
	HighlightErrorEverywhere:    "这会匹配每一行，请写得更具体一些。",
	HighlightErrorInvalidFmt:    "不是有效的表达式：%s",
	HighlightErrorTooMany:       "规则已达上限，请先删除一条。",
	HighlightSave:               "添加",
	ColorRed:                    "红",
	ColorYellow:                 "黄",
	ColorGreen:                  "绿",
	ColorCyan:                   "青",
	ColorBlue:                   "蓝",
	ColorMagenta:                "紫",

	SnippetsGroup:             "快捷命令",
	SnippetsRow:               "常用命令",
	SnippetsDescription:       "从命令面板发送到当前终端。",
	SnippetAdd:                "添加命令",
	SnippetEdit:               "编辑",
	SnippetDelete:             "删除",
	SnippetDialogAddTitle:     "添加快捷命令",
	SnippetDialogEditTitle:    "编辑快捷命令",
	SnippetName:               "名称",
	SnippetNamePlaceholder:    "例如：重启 nginx",
	SnippetCommand:            "命令",
	SnippetCommandPlaceholder: "例如：systemctl restart nginx",
	SnippetRun:                "发送后自动按回车",
	SnippetErrorName:          "请给快捷命令起个名字。",
	SnippetErrorCommand:       "请输入要发送的命令。",
	SnippetErrorTooMany:       "快捷命令已达上限，请先删除一条。",
	SnippetSave:               "保存",
	SnippetPaletteFmt:         "快捷命令：%s",
	StatusSnippetSentFmt:      "已发送“%s”。",
	StatusSnippetNoTerminal:   "请先打开一个终端再发送快捷命令。",
	StatusSnippetNotConnected: "该终端尚未连接。",

	CatGeneral:    "通用",
	CatTerminal:   "终端",
	CatConnection: "连接",
	CatSecurity:   "安全",

	ZoomIn:            "放大终端文字",
	ZoomOut:           "缩小终端文字",
	ZoomReset:         "恢复终端文字大小",
	StatusFontSizeFmt: "终端字号 %d。",
	PaddingRow:        "终端边距",
	PaddingCompact:    "紧凑",
	PaddingNormal:     "标准",
	PaddingRoomy:      "宽松",
	GutterRow:         "行号栏",
	GutterDescription: "显示行号或写入时间。",
	GutterOff:         "关闭",
	GutterNumbers:     "行号",
	GutterTime:        "时间",
	GutterBoth:        "行号+时间",
	CursorRow:         "光标闪烁",
	CursorDescription: "闲置后恢复常亮；程序要求时不闪烁。",

	FontChoose:              "选择",
	FontCurrentFmt:          "当前：%s。可以使用任何已安装的等宽字体。",
	FontMissingFmt:          "字体“%s”已不在系统中，已改用系统字体。",
	FontPickerTitle:         "终端字体",
	FontPickerDescription:   "选择系统字体，或任何已安装的等宽字体。",
	FontPickerPlaceholder:   "搜索字体",
	FontPickerLoading:       "正在读取已安装的字体…",
	FontPickerEmpty:         "没有匹配的字体。",
	FontPickerSample:        "The quick brown fox 敏捷的棕色狐狸 0123",
	FontPickerSampleSymbols: "0123456789 {}[]() <> -> == != ~ $ # | iIl1 oO0",

	NotifyRow:         "桌面通知",
	NotifyDescription: "终端里的程序请求通知时（OSC 9、777、99），显示系统通知。每个终端每秒最多一条。",

	ProxyRow:                 "出站代理",
	ProxyConfigure:           "配置",
	ProxyDescriptionNone:     "直接连接。可以让 SOCKS5 或 HTTP 代理承载每个会话的第一条连接。",
	ProxyDescriptionFmt:      "%s（%s）承载每个会话的第一条连接。",
	ProxyNone:                "无",
	ProxySOCKS5:              "SOCKS5",
	ProxyHTTP:                "HTTP CONNECT",
	ProxyDialogTitle:         "出站代理",
	ProxyDialogDescription:   "主机名由代理解析。直连列表中的目标不走代理。",
	ProxyAddress:             "代理地址（主机:端口）",
	ProxyUser:                "用户名（可选）",
	ProxyPassword:            "密码（可选）",
	ProxyPasswordPlaceholder: "留空则保留已保存的密码",
	ProxyBypass:              "直接连接的目标",
	ProxySave:                "保存",
	ProxyErrorAddress:        "请按 主机:端口 填写代理，端口为 1 到 65535。",
	ProxyErrorLogin:          "用户名和密码最多 255 个字符。",
	ProxyErrorBypassFmt:      "直连列表最多 %d 项，每项不超过 100 个字符。",
	ProxyErrorNoVault:        "这里无法保存密码。密码留空则不带密码连接。",
	ProxyPasswordLocked:      "已保存的代理密码被锁定：请先在 设置 › 安全 中解锁。",

	SearchPlaceholder: "搜索设置",
	SearchTitle:       "搜索结果",
	SearchEmptyFmt:    "没有匹配“%s”的设置。",

	ImportSessions:            "从其他软件导入",
	ImportChooserTitle:        "导入会话",
	ImportChooserBody:         "选择导入来源，再选取它保存会话的文件。",
	ImportPickerFmt:           "导入 %s 会话",
	ImportCanceled:            "已取消导入。",
	ImportFailedFmt:           "无法导入：%v",
	ImportSummaryFmt:          "已从 %[2]s 导入 %[1]d 台主机；跳过 %[3]d 台已保存的主机。",
	ImportWarningKeysFmt:      "有 %d 个会话使用私钥文件。密钥不会被导入，请在“密钥”页面添加。",
	ImportWarningSkippedFmt:   "有 %d 个会话的类型不支持导入（远程桌面、文件传输、该软件的串口等），已跳过。",
	ImportWarningPasswords:    "不会导入已保存的密码：各个软件都用自己的方式加密了它们。",
	ImportWarningCleartextFmt: "有 %d 个会话在此文件里以明文保存了密码或私钥。这些内容没有被导入；导入完成后请删除该文件。",

	BroadcastStart:            "同时输入到本标签页的所有窗格",
	BroadcastStop:             "停止同时输入",
	StatusBroadcastOnFmt:      "现在输入的内容会发送到本标签页的全部 %d 个窗格。",
	StatusBroadcastOff:        "广播已关闭。",
	StatusBroadcastNeedsSplit: "请先拆分标签页：广播会把输入发送到其中的每个窗格。",

	ProtocolSSH:            "SSH",
	ProtocolTelnet:         "Telnet",
	ProtocolSerial:         "串口",
	FieldDevice:            "设备",
	FieldDevicePlaceholder: "/dev/ttyUSB0 或 COM3",
	FieldLine:              "波特率与帧格式",
	FieldLinePlaceholder:   "115200 8N1",
	FieldFlow:              "流控",
	FlowNone:               "无",
	FlowHardware:           "RTS/CTS",
	FlowSoftware:           "XON/XOFF",
	PortsPicker:            "已发现的串口",
	PortsNone:              "没有发现串口",
	ErrSerialLine:          "请按 波特率 帧格式 填写，例如 115200 8N1（数据位 5 到 8，校验 N、O、E、M 或 S，停止位 1 或 2）。",
	ErrSerialDevice:        "设备名含有串口不可能有的字符。",
	TelnetPlaintext:        "Telnet 不加密：你输入的内容（包括密码）会以明文经过网络。",
	SerialOpenedFmt:        "已用 %[2]s 打开 %[1]s。",

	TransferSentFmt:           "已发送 %d 个文件。",
	TransferReceivedFmt:       "已接收 %d 个文件，保存在 %s。",
	TransferCancelled:         "文件传输已取消。",
	TransferFailedFmt:         "文件传输失败：%v",
	TransferDeclined:          "已拒绝文件传输。",
	TransferChooseFolderTitle: "选择接收文件的保存位置",
	TransferNoDownloads:       "没有“下载”文件夹，请改选一个文件夹。",
	TransferChooseFilesTitle:  "选择要发送的文件",
	TransferTooManyFmt:        "一次最多发送 %d 个文件。",
	TransferNotFileFmt:        "%s 不是普通文件。",
	TransferTooBigFmt:         "%s 超过 4 GiB，这是该协议能发送的上限。",
	TransferReceiveTitle:      "接收文件",
	TransferSendTitle:         "发送文件",
	TransferReceiveOfferFmt:   "%s 想向你发送文件（ZMODEM）。选择保存位置之前不会保存任何内容。",
	TransferSendOfferFmt:      "%s 已准备好接收文件（ZMODEM）。请选择要发送的文件。",
	TransferDecline:           "拒绝",
	TransferChooseFiles:       "选择文件",
	TransferChooseFolder:      "选择文件夹",
	TransferToDownloads:       "保存到“下载”",

	CatFiles:                        "文件",
	FilesParallel:                   "并行传输数",
	FilesKeepTimes:                  "保留修改时间",
	FilesKeepTimesDescription:       "传输后的文件沿用原文件的修改时间。权限始终保留。",
	FilesFollow:                     "跟随终端所在目录",
	FilesFollowDescription:          "文件面板会跳到 shell 上报的目录（OSC 7）。需要把 shell 配置为上报。",
	FilesFollowSetup:                "bash / zsh 目录上报",
	FilesFollowSetupDescription:     "让提示符上报当前目录。粘贴到远程 shell 或其启动文件中。",
	FilesFollowSetupCopy:            "复制命令",
	FilesFollowSetupCopied:          "已复制，请粘贴到远程 shell 中。",
	TipNewLink:                      "新建符号链接",
	LinkTitle:                       "新建符号链接",
	LinkName:                        "链接名称",
	LinkTarget:                      "指向",
	LinkErrorName:                   "请输入链接名称。",
	LinkErrorTarget:                 "请输入链接指向的路径。",
	LinkCreatingFmt:                 "正在把 %s 链接到 %s…",
	LinkCreatedFmt:                  "已把 %s 链接到 %s。",
	LinkBroken:                      "目标不存在",
	Bookmarks:                       "书签",
	BookmarksEmpty:                  "还没有书签。点星标可收藏当前所在目录。",
	BookmarkAdd:                     "收藏此目录",
	BookmarkRemove:                  "移除书签",
	BookmarkLimitFmt:                "每台服务器最多 %d 个书签。",
	Extract:                         "在此解压",
	ExtractTitle:                    "解压压缩包",
	ExtractDescriptionFmt:           "把 %s 解压到一个新目录，不会覆盖已有内容。",
	ExtractFolder:                   "目录名",
	ExtractConfirm:                  "解压",
	ExtractErrorFolder:              "请输入不含斜杠的目录名。",
	ExtractRunningFmt:               "正在解压 %s…",
	ExtractDoneFmt:                  "已把 %s 解压到 %s。",
	ExtractFailedFmt:                "无法解压 %s：%s",
	TransferPause:                   "暂停",
	TransferResume:                  "继续",
	TransferPausedFmt:               "已暂停 · %s / %s",
	SkipExisting:                    "跳过已有",
	TransferSpeedFmt:                "%s/秒",
	LinkCreate:                      "创建链接",
	PaneLocal:                       "显示本机文件",
	PaneRemote:                      "显示服务器文件",
	LocalFiles:                      "本机",
	SendToServer:                    "上传到服务器当前目录",
	DownloadHere:                    "下载到本机当前显示的目录",
	XYMenu:                          "文件传输",
	XYTitle:                         "文件传输",
	XYDescription:                   "请先在对端启动程序（接收用 rx 或 rb，发送用 sx 或 sb），再点“开始”。",
	XYProtocol:                      "协议",
	XYDirection:                     "方向",
	XYSend:                          "发送",
	XYReceive:                       "接收",
	XYStart:                         "开始",
	XYSaveTitle:                     "将接收的文件保存为",
	DropUploadingFmt:                "正在把拖入的 %d 项上传到 %s",
	LaunchTitle:                     "打开连接",
	LaunchDescriptionFmt:            "通过 %[1]s 连接到 %[2]s?",
	LaunchWarning:                   "其他程序请求了这次连接。仅在你预期它时才连接。",
	LaunchConnect:                   "连接",
	GeneralLogin:                    "开机自启",
	GeneralLoginDescription:         "登录后在后台启动。",
	GeneralUpdate:                   "检查更新",
	GeneralUpdateDescription:        "到 GitHub 查找新版本并打开下载页,不会替你安装任何东西。",
	GeneralUpdateCheck:              "立即检查",
	GeneralUpdateDownload:           "下载",
	GeneralUpdateStartup:            "启动时检查",
	GeneralUpdateStartupDescription: "每次启动 Motty 时查找新版本。检查会让 GitHub 看到你的地址，所以默认关闭。",
	UpdateChecking:                  "正在检查更新…",
	UpdateLatestFmt:                 "Motty %s 已是最新版本。",
	UpdateAvailableFmt:              "Motty %s 可更新。",
	UpdateNoReleaseText:             "还没有找到已发布的版本。",
	ActionSSHFmt:                    "SSH 连接 %s",
	ActionBrowserFmt:                "在浏览器中打开 %s",
	ActionFilesFmt:                  "在文件面板中显示 %s",
	ActionExtractFmt:                "解压 %s",
	ActionCopyFmt:                   "复制 %s",
	DiagnosticsRow:                  "诊断报告",
	DiagnosticsDescription:          "把版本、设置和最近日志存成文本，用于问题反馈。主机和账号已隐藏。",
	DiagnosticsExport:               "导出",
	DiagnosticsSaveTitle:            "保存诊断报告",
	DiagnosticsSavedFmt:             "诊断报告已保存到 %s",
	DiagnosticsFailedFmt:            "无法保存诊断报告:%v",
	CatBackup:                       "备份与同步",
	CatMCP:                          "MCP",
	BackupGroupFile:                 "备份文件",
	BackupGroupSync:                 "云同步",
	BackupExportRow:                 "备份全部数据",
	BackupExportDescription:         "主机、设置、已信任的主机密钥和已保存的密码，用你设定的密码加密成一个文件。",
	BackupExport:                    "导出",
	BackupRestore:                   "恢复",
	BackupExportTitle:               "备份到文件",
	BackupExportDialogDescription:   "设置打开备份用的密码，丢失后无法找回。",
	BackupPassword:                  "备份密码",
	BackupPasswordConfirm:           "确认密码",
	BackupPasswordPlaceholder:       "至少 8 个字符",
	BackupExportSave:                "选择文件",
	BackupSaveTitle:                 "保存备份",
	BackupSavedFmt:                  "备份已保存到 %s。",
	BackupFailedFmt:                 "备份失败：%s",
	BackupRestoreTitle:              "恢复备份",
	BackupRestoreDescription:        "这会用备份中的主机和设置替换本机的内容；当前的内容保存在 config.db.before-restore。",
	BackupRestoreConfirm:            "恢复",
	BackupPickerTitle:               "选择备份文件",
	BackupRestoredFmt:               "已恢复备份（%d 台主机）。",
	BackupErrPasswordShort:          "密码至少 8 个字符。",
	BackupErrPasswordMismatch:       "两次输入的密码不一致。",
	BackupErrWrong:                  "密码错误，或文件已被改动。",
	BackupErrNotBackup:              "这不是 Motty 的备份文件。",
	BackupErrTooNew:                 "这份备份由更新版本的 Motty 创建。",
	BackupErrSessions:               "请先关闭所有会话：恢复会替换会话所属的主机。",
	BackupErrVault:                  "请先解锁已保存的密码。",
	BackupSecretsFailedFmt:          "已保存的密码未能恢复：%s",
	SyncRow:                         "服务器",
	SyncRowNone:                     "未设置同步服务。把备份放在 WebDAV、S3 或 GitHub Gist 上。",
	SyncRowFmt:                      "备份保存在 %s。",
	SyncConfigure:                   "WebDAV",
	SyncDialogTitle:                 "WebDAV 服务器",
	SyncDialogDescription:           "备份在离开本机前已加密，服务器只能看到密文。地址留空即停止同步。",
	SyncURL:                         "服务器地址",
	SyncUser:                        "用户",
	SyncPassword:                    "密码",
	SyncPath:                        "文件",
	SyncSave:                        "保存",
	SyncRunRow:                      "同步",
	SyncRunNever:                    "尚未同步。同步会把本机的主机和设置与服务器上的合并。",
	SyncRunLastFmt:                  "上次同步于 %s。同步会把本机的主机和设置与服务器上的合并。",
	SyncNow:                         "立即同步",
	SyncAuto:                        "自动同步",
	SyncAutoDescription:             "主机、设置或已保存的密码在本机改动后几秒内同步，启动应用时也会同步一次。其他电脑的改动会在本机下次同步时到达。",
	SyncRunFailedFmt:                "上次同步失败：%s",
	SyncWorking:                     "正在与服务器通信…",
	SyncDone:                        "已与服务器同步。",
	SyncErrNoStore:                  "无法打开设置数据库，没有可同步的内容。",
	SyncTestOK:                      "已连接到服务器。",
	SyncFailedFmt:                   "同步出错：%s",
	SyncErrUnauthorized:             "服务拒绝了该登录，或该登录没有这个权限。",
	SyncErrNotFound:                 "服务器上还没有备份。",
	SyncErrInsecure:                 "请使用 https:// 地址（http:// 仅限本机）。",
	SyncErrNeedsVault:               "保存同步登录需要用到「已保存的密码」，但它当前不可用。",
	SyncErrFields:                   "请填写地址、用户、密码、文件和备份密码。",
	SyncErrAddress:                  "这不是有效的服务器地址。",
	SyncConfigureS3:                 "S3",
	SyncConfigureGist:               "Gist",
	SyncErrTooLarge:                 "备份对这个服务来说太大了。",
	S3DialogTitle:                   "S3 兼容存储（1/2）",
	S3DialogDescription:             "AWS S3、Cloudflare R2、MinIO、Backblaze B2 等。备份在离开本机前已加密。端点留空即停止同步。",
	S3Endpoint:                      "端点",
	S3Region:                        "区域",
	S3Bucket:                        "存储桶",
	S3Key:                           "对象键",
	S3PathStyle:                     "路径式寻址（yes 或 no）",
	S3LoginTitle:                    "S3 访问密钥（2/2）",
	S3LoginDescription:              "请使用只能读写这一个对象的密钥。密钥保存在「已保存的密码」里，不会写进设置文件。留空则沿用已保存的。",
	S3AccessKey:                     "访问密钥 ID",
	S3Secret:                        "Secret Access Key（密钥）",
	S3ErrFields:                     "请填写端点、存储桶、访问密钥、密钥和备份密码。",
	GistDialogTitle:                 "GitHub Gist 代码片段",
	GistDialogDescription:           "加密后的备份保存在一个私密 Gist 里。令牌只需要 Gists 权限（读写），保存在「已保存的密码」中。Gist 留空则在第一次上传时创建。",
	GistToken:                       "访问令牌",
	GistID:                          "Gist ID（可选）",
	GistFile:                        "文件名",
	GistAPI:                         "API 地址（可选）",
	GistErrFields:                   "请填写访问令牌和备份密码。",
	GistErrTooLarge:                 "Gist 单个文件最大 1 MB，这份备份太大了。",
	MenuSaveOutput:                  "将输出保存为文本",
	MenuSearchSelection:             "用网页搜索所选内容",
	OutputSaveTitle:                 "保存终端输出",
	OutputSavedFmt:                  "终端输出已保存到 %s。",
	OutputSaveFailedFmt:             "无法保存终端输出：%s",
	OutputEmpty:                     "终端里没有可保存的输出。",
	CloseAppTitle:                   "关闭 Motty？",
	CloseAppDescriptionFmt:          "有 %d 个会话仍然打开，关闭后会断开连接。",
	CloseAppConfirm:                 "关闭",
	RemoteMenu:                      "进程与容器",
	RemoteTitle:                     "远程工具",
	RemoteDescriptionFmt:            "%s 上的进程与 Docker 容器。内容不会自动刷新。",
	RemoteTabProcesses:              "进程",
	RemoteTabImages:                 "镜像",
	RemoteTabVolumes:                "卷",
	RemoteTabNetworks:               "网络",
	RemoteTabCompose:                "Compose",
	RemoteColConfig:                 "配置文件",
	RemoteEmptyCompose:              "没有 Compose 项目。",
	RemoteConfirmDownTitle:          "停用并移除此项目？",
	RemoteConfirmDownFmt:            "项目 %s 将被停止，其容器和网络会被移除（卷保留）。",
	RemoteDown:                      "停用",
	RemoteErrNoCompose:              "此服务器没有 docker compose 插件。",
	RemoteReniceLower:               "降低优先级 +10",
	RemoteReniceNormal:              "恢复 0",
	RemoteReniceHigher:              "提高优先级 -5",
	RemoteReniceDoneFmt:             "已把进程 %d 的 nice 值设为 %d。",
	RemoteColSize:                   "大小",
	RemoteColCreated:                "创建于",
	RemoteColDriver:                 "驱动",
	RemoteColScope:                  "范围",
	RemoteEmptyImages:               "没有镜像。",
	RemoteEmptyVolumes:              "没有卷。",
	RemoteEmptyNetworks:             "没有网络。",
	RemoteConfirmRemoveImageTitle:   "删除该镜像？",
	RemoteConfirmRemoveImageFmt:     "镜像 %s 将被删除；有容器在用它时 Docker 会拒绝。",
	RemoteConfirmRemoveVolumeTitle:  "删除该卷？",
	RemoteConfirmRemoveVolumeFmt:    "卷 %s 及其中的数据将被删除；有容器在用它时 Docker 会拒绝。",
	RemoteConfirmRemoveNetworkTitle: "删除该网络？",
	RemoteConfirmRemoveNetworkFmt:   "网络 %s 将被删除；有容器在用它时 Docker 会拒绝。",
	RemoteTabContainers:             "容器",
	RemoteSortCPU:                   "按 CPU",
	RemoteSortMemory:                "按内存",
	RemoteColPID:                    "PID",
	RemoteColUser:                   "用户",
	RemoteColCPU:                    "CPU %",
	RemoteColMemory:                 "内存 %",
	RemoteColTime:                   "运行时长",
	RemoteColCommand:                "命令",
	RemoteColName:                   "名称",
	RemoteColImage:                  "镜像",
	RemoteColState:                  "状态",
	RemoteColStatus:                 "运行情况",
	RemoteColID:                     "ID",
	RemoteRefresh:                   "刷新",
	RemoteTerminate:                 "终止",
	RemoteForceKill:                 "强制结束",
	RemoteStart:                     "启动",
	RemoteStop:                      "停止",
	RemoteRestart:                   "重启",
	RemoteLogs:                      "日志",
	RemoteRemove:                    "删除",
	RemotePull:                      "拉取镜像",
	RemotePullTitle:                 "拉取镜像",
	RemotePullDescription:           "在终端里运行 docker pull,进度可见,Ctrl+C 可中断。",
	RemoteImageField:                "镜像",
	RemoteImagePlaceholder:          "例如 nginx:latest",
	RemoteBadImage:                  "这不是有效的镜像名。例如 nginx:latest 或 ghcr.io/org/app:1.2。",
	RemoteCreate:                    "新建",
	RemoteCreateVolumeTitle:         "新建卷",
	RemoteCreateVolumeDescription:   "在服务器上创建一个空卷。",
	RemoteCreateNetworkTitle:        "新建网络",
	RemoteCreateNetworkDescription:  "在服务器上创建一个 bridge 网络。",
	RemoteNameField:                 "名称",
	RemoteBadName:                   "只能用字母、数字、下划线、点和连字符(最多 128 个),且以字母或数字开头。",
	RemoteLoading:                   "正在从服务器读取…",
	RemoteLoadedFmt:                 "共 %d 项。",
	RemoteEmptyContainers:           "没有容器。",
	RemoteErrUnsupportedFmt:         "不支持这台服务器上的 ps：%s",
	RemoteErrNoDocker:               "这台服务器没有安装 Docker。",
	RemoteErrFailedFmt:              "失败：%s",
	RemoteNeedSelection:             "请先选择一行。",
	RemoteSignalSentFmt:             "已向进程 %[2]d 发送 %[1]s。",
	RemoteActionDoneFmt:             "已完成：%s。",
	RemoteConfirmTerminateTitle:     "终止该进程？",
	RemoteConfirmKillTitle:          "强制结束该进程？",
	RemoteConfirmProcessFmt:         "进程 %d（%s）。其中未保存的内容会丢失。",
	RemoteConfirmRemoveTitle:        "删除该容器？",
	RemoteConfirmRemoveFmt:          "容器 %s 将被删除，其可写层中的数据一并丢失。",
	RemoteLogsFmt:                   "在终端中查看 %s 的日志。",
	RemoteClose:                     "关闭",
	ExtEditButton:                   "用系统编辑器打开",
	ExtEditOpenedFmt:                "已用系统编辑器打开 %s。在那里保存后会自动传回服务器。",
	ExtEditSentFmt:                  "已把 %s 传回服务器。",
	ExtEditSendFailedFmt:            "无法把 %s 传回服务器：%s",
	ExtEditBadFmt:                   "%s 现在不是 UTF-8 文本，没有传回服务器。",
	ExtEditRefused:                  "这类文件会被直接运行而不是编辑，请使用内置编辑器。",
	ExtEditOpenFailedFmt:            "无法用系统编辑器打开 %s：%s",
	ExtEditConflictTitle:            "服务器上的文件已改变",
	ExtEditConflictDescriptionFmt:   "%s 在你打开之后又在服务器上被改动。要用你的版本替换服务器上的吗？",
	ExtEditReplace:                  "替换",
	HostViewList:                    "列表视图",
	HostViewGrid:                    "卡片视图",
	BgGroup:                         "终端背景",
	BgRow:                           "终端背景图片",
	BgRowNone:                       "把一张图片放在终端文字后面，只有终端区域变为半透明。",
	BgRowSetFmt:                     "%s，%d×%d。按铺满显示在终端后面。",
	BgChoose:                        "选择",
	BgRemove:                        "移除",
	BgPickerTitle:                   "选择背景图片",
	BgVisibility:                    "图片可见度",
	BgSubtle:                        "淡",
	BgMedium:                        "中",
	BgStrong:                        "浓",
	BgLoading:                       "正在载入背景图片…",
	BgErrFmt:                        "无法使用该图片：%s",
	BgErrTooLarge:                   "图片文件超过 30 MB。",
	BgErrTooManyPixels:              "图片像素超过 6400 万。",
	BgErrFormat:                     "只支持 PNG 和 JPEG 图片。",
	TabMoveBusy:                     "该标签页正在传输文件，暂时不能移到另一个窗口。",
	OTPField:                        "一次性验证码密钥",
	OTPPlaceholder:                  "可选:otpauth:// 地址或 base32 密钥",
	OTPUncheckedPlaceholder:         "留空则保留已有密钥",
	OTPSavedPlaceholder:             "已保存。输入新密钥可替换",
	OTPRemovingPlaceholder:          "保存时将被删除",
	OTPCheck:                        "检查",
	OTPRemove:                       "删除",
	OTPScan:                         "扫描二维码图片",
	OTPPaste:                        "粘贴图片",
	ErrOTPPasteNoImage:              "剪贴板里没有图片,请先复制二维码图片",
	OTPScanTitle:                    "选择二维码图片",
	OTPScanFilter:                   "图片 (PNG, JPEG, GIF)",
	OTPScannedFmt:                   "已读取二维码:%s。保存主机后生效",
	ErrOTPScanTooLarge:              "图片超过 10 MB,请裁剪到二维码范围后重试",
	ErrOTPScanDimension:             "图片像素过多,请缩小或裁剪到二维码范围",
	ErrOTPScanNotImage:              "该文件不是 PNG、JPEG 或 GIF 图片",
	ErrOTPScanNoCode:                "图片中没有找到二维码,请使用更清晰、完整的二维码图片",
	ErrOTPScanNotOTP:                "该二维码不是一次性验证码(otpauth://)地址",
	ErrOTPScanMigration:             "验证器导出二维码已损坏,或使用了不支持的算法(MD5),请重新导出或使用单个账号的二维码",
	ErrOTPScanManyFmt:               "导出二维码包含 %d 个账号,且没有与此主机用户名相符的。请在验证器中只导出一个账号再扫描",
	ErrOTPScanReadFmt:               "无法读取图片:%v",
	X11Forward:                      "转发 X11(仅本会话有效的 cookie)",
	X11UnavailableFmt:               "本次会话未启用 X11 转发:%s",
	AgentForward:                    "转发 ssh-agent",
	AgentForwardUnavailableFmt:      "本次会话未启用 ssh-agent 转发:%s",
	OTPKeep:                         "保留",
	ErrOTPInvalid:                   "不是有效的一次性验证码密钥",
	ErrOTPNoVault:                   "无法使用保存密码的保险库,密钥保存不了",
	ErrOTPLocked:                    "请先在设置里解锁保存的密码",
	ErrOTPCheckFmt:                  "无法检查已保存的一次性验证码密钥:%v",
	OTPSavedFmt:                     "已为 %s 保存一次性验证码密钥",
	OTPRemovedFmt:                   "已删除 %s 的一次性验证码密钥",
	AlgoField:                       "SSH 算法",
	AdvancedOptions:                 "高级选项",
	TransferField:                   "文件传输协议",
	TransferAuto:                    "自动",
	AlgoDefault:                     "默认",
	AlgoCompatible:                  "兼容",
	AlgoCustom:                      "自定义",
	AlgoCustomField:                 "允许的算法",
	AlgoCustomPlaceholder:           "aes256-gcm@openssh.com, hmac-sha2-256, …",
	ErrAlgoRequired:                 "自定义需要填写算法列表",
	ErrAlgoUnknownFmt:               "本程序不支持的算法:%s",
	AlgoWeakFmt:                     "这条连接协商到了较弱的算法:%s",
	AlgoNegotiatedFmt:               "算法 %s",
	UpdateFailedFmt:                 "无法检查更新:%v",
	LoginFailedFmt:                  "无法更改开机自启:%v",
}
