package xkbcommon

// KeyCodeInvalid is a special value indicating that Key Code is invalid
const KeyCodeInvalid = 0xffffffff

// LayoutInvalid is a special value indicating that Layout is invalid
const LayoutInvalid = 0xffffffff

// LevelInvalid is a special value indicating that Level is invalid
const LevelInvalid = 0xffffffff

// ModInvalid is a special value indicating that Modifier is invalid
const ModInvalid = 0xffffffff

// LedInvalid is a special value indicating that Led is invalid
const LedInvalid = 0xffffffff

// KeyCodeMax is the maximum KeyCode possible
const KeyCodeMax = 0xffffffff - 1

// KeymapFormatTextV1 is the current/classic XKB text format, as generated by xkbcomp -xkb.
// This keymap consists of a single top-level `xkb_keymap` block, underwhich are nested sections.
const KeymapFormatTextV1 = 1

// ComposeFormatTextV1 is the classic libX11 Compose text format, described in Compose(5).
const ComposeFormatTextV1 = 1

// ModShiftMask is the Shift modifier mask - provided for convenience only
const ModShiftMask = 0x01

// ModAltMask is the Alt modifier mask - provided for convenience only
const ModAltMask = 0x02

// ModControlMask is the Control modifier mask - provided for convenience only
const ModControlMask = 0x04

// ContextNoFlags provides no flags for the context.
const ContextNoFlags = 0

// ComposeCompileNoFlags provides no flags for compose compile.
const ComposeCompileNoFlags = 0

// ComposeStateNoFlags provides no flags for compose state.
const ComposeStateNoFlags = 0

// ComposeStatus is the status of the Compose sequence state machine.
type ComposeStatus uint8

const (
	// ComposeNothing is The initial state; no sequence has started yet.
	ComposeNothing ComposeStatus = 0
	// ComposeNothing is In the middle of a sequence.
	ComposeComposing ComposeStatus = 1
	// ComposeComposed is A complete sequence has been matched.
	ComposeComposed ComposeStatus = 2
	// ComposeCancelled is The last sequence was cancelled due to an unmatched keysym.
	ComposeCancelled ComposeStatus = 3
)

// ComposeFeedResult is The effect of a keysym fed to ComposeStateFeed().
type ComposeFeedResult uint8

const (
	// ComposeFeedIgnored is The keysym had no effect - it did not affect the status.
	ComposeFeedIgnored ComposeFeedResult = 0
	// ComposeFeedAccepted is The keysym started, advanced or cancelled a sequence.
	ComposeFeedAccepted ComposeFeedResult = 1
)
