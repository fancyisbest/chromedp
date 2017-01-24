package systeminfo

// AUTOGENERATED. DO NOT EDIT.

import (
	. "github.com/knq/chromedp/cdp"
	"github.com/mailru/easyjson"
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

// GPUDevice describes a single graphics processor (GPU).
type GPUDevice struct {
	VendorID     float64 `json:"vendorId,omitempty"`     // PCI ID of the GPU vendor, if available; 0 otherwise.
	DeviceID     float64 `json:"deviceId,omitempty"`     // PCI ID of the GPU device, if available; 0 otherwise.
	VendorString string  `json:"vendorString,omitempty"` // String description of the GPU vendor, if the PCI ID is not available.
	DeviceString string  `json:"deviceString,omitempty"` // String description of the GPU device, if the PCI ID is not available.
}

// GPUInfo provides information about the GPU(s) on the system.
type GPUInfo struct {
	Devices              []*GPUDevice        `json:"devices,omitempty"` // The graphics devices on the system. Element 0 is the primary GPU.
	AuxAttributes        easyjson.RawMessage `json:"auxAttributes,omitempty"`
	FeatureStatus        easyjson.RawMessage `json:"featureStatus,omitempty"`
	DriverBugWorkarounds []string            `json:"driverBugWorkarounds,omitempty"` // An optional array of GPU driver bug workarounds.
}
