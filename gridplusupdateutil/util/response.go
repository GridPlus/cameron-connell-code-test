package util

import (
	"encoding/json"
	"fmt"
)

type Response map[string]Target

// GetResponse returns the raw bytes of the targets available for update (if any).
func GetResponse(data []byte) (*Response, error) {
	var result Response
	err := json.Unmarshal(data, &result)
	return &result, err
}

// ProcessUpdates checks the response struct for updates and updates the target if necessary.
func (res *Response) ProcessUpdates() {
	if entry, exists := (*res)["HSM"]; exists {
		res.updateOne(entry)
	}

	if entry, exists := (*res)["GCE"]; exists {
		res.updateOne(entry)
	}
}

// UpdateOne processes a single target, sending the data to the update function.
func (res *Response) updateOne(entry Target) {
	if entry.hasPrereq() {
		entry.processPrereq()
		writeVersion(entry.PrereqUpdate)
	}

	fmt.Printf("Updating %v to %v\n", entry.appCode(), entry.currentVersion())
	writeVersion(&entry)
}
