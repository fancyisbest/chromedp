package database

// AUTOGENERATED. DO NOT EDIT.

import (
	. "github.com/knq/chromedp/cdp"
)

var (
	_ BackendNode
	_ BackendNodeID
	_ ComputedProperty
	_ ErrorType
	_ Frame
	_ FrameID
	_ LoaderID
	_ Message
	_ MessageError
	_ MethodType
	_ Node
	_ NodeID
	_ NodeType
	_ PseudoType
	_ RGBA
	_ ShadowRootType
	_ Timestamp
)

type EventAddDatabase struct {
	Database *Database `json:"database,omitempty"`
}

// EventTypes is all event types in the domain.
var EventTypes = []MethodType{
	EventDatabaseAddDatabase,
}
