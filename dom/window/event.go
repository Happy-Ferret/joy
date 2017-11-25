package window

// js:"Event,omit"
type Event interface {

	// InitEvent
	// js:"initEvent"
	InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool)

	// PreventDefault
	// js:"preventDefault"
	PreventDefault()

	// StopImmediatePropagation
	// js:"stopImmediatePropagation"
	StopImmediatePropagation()

	// StopPropagation
	// js:"stopPropagation"
	StopPropagation()

	// Bubbles prop
	// js:"bubbles"
	Bubbles() (bubbles bool)

	// Cancelable prop
	// js:"cancelable"
	Cancelable() (cancelable bool)

	// CancelBubble prop
	// js:"cancelBubble"
	CancelBubble() (cancelBubble bool)

	// CancelBubble prop
	// js:"setcancelBubble"
	SetCancelBubble(cancelBubble bool)

	// CurrentTarget prop
	// js:"currentTarget"
	CurrentTarget() (currentTarget EventTarget)

	// DefaultPrevented prop
	// js:"defaultPrevented"
	DefaultPrevented() (defaultPrevented bool)

	// EventPhase prop
	// js:"eventPhase"
	EventPhase() (eventPhase uint8)

	// IsTrusted prop
	// js:"isTrusted"
	IsTrusted() (isTrusted bool)

	// ReturnValue prop
	// js:"returnValue"
	ReturnValue() (returnValue bool)

	// ReturnValue prop
	// js:"setreturnValue"
	SetReturnValue(returnValue bool)

	// SrcElement prop
	// js:"srcElement"
	SrcElement() (srcElement Element)

	// Target prop
	// js:"target"
	Target() (target EventTarget)

	// TimeStamp prop
	// js:"timeStamp"
	TimeStamp() (timeStamp uint64)

	// Type prop
	// js:"type"
	Type() (kind string)
}
