package heapprofiler

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

type EventAddHeapSnapshotChunk struct {
	Chunk string `json:"chunk,omitempty"`
}

type EventResetProfiles struct{}

type EventReportHeapSnapshotProgress struct {
	Done     int64 `json:"done,omitempty"`
	Total    int64 `json:"total,omitempty"`
	Finished bool  `json:"finished,omitempty"`
}

// EventLastSeenObjectID if heap objects tracking has been started then
// backend regulary sends a current value for last seen object id and
// corresponding timestamp. If the were changes in the heap since last event
// then one or more heapStatsUpdate events will be sent before a new
// lastSeenObjectId event.
type EventLastSeenObjectID struct {
	LastSeenObjectID int64     `json:"lastSeenObjectId,omitempty"`
	Timestamp        Timestamp `json:"timestamp,omitempty"`
}

// EventHeapStatsUpdate if heap objects tracking has been started then
// backend may send update for one or more fragments.
type EventHeapStatsUpdate struct {
	StatsUpdate []int64 `json:"statsUpdate,omitempty"` // An array of triplets. Each triplet describes a fragment. The first integer is the fragment index, the second integer is a total count of objects for the fragment, the third integer is a total size of the objects for the fragment.
}

// EventTypes is all event types in the domain.
var EventTypes = []MethodType{
	EventHeapProfilerAddHeapSnapshotChunk,
	EventHeapProfilerResetProfiles,
	EventHeapProfilerReportHeapSnapshotProgress,
	EventHeapProfilerLastSeenObjectID,
	EventHeapProfilerHeapStatsUpdate,
}
