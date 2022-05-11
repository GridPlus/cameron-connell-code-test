package util

type Target struct {
	AppCode       string `json:"AppCode,omitempty"`
	TargetVersion string `json:"TargetVersion,omitempty"`
	DownloadURL   string `json:"downloadURL,omitempty"`
	DownloadSize  int    `json:"DownloadSize,omitempty"`
	ArtifactSig   string `json:"ArtifactSig,omitempty"`
	*PrereqUpdate `json:"PrereqUpdate,omitempty"`
}

// appCode is a getter for the appCode field.
func (t *Target) appCode() string {
	return t.AppCode
}

// currentVersion is a getter for the currentVersion field.
func (t *Target) currentVersion() string {
	return t.TargetVersion
}

// hasPrereq returns true if the target has a prereq update.
func (t *Target) hasPrereq() bool {
	return t.PrereqUpdate != nil
}
