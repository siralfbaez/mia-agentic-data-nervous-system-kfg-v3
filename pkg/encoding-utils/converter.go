package encodingutils

import (
	"encoding/json"
	"fmt"
)

// ToKFGJSON converts raw signal payloads into the standardized KFG JSONB format
func ToKFGJSON(payload interface{}) ([]byte, error) {
	// Staff Tip: In production, use a fast-json library or Protobuf-to-JSON
	data, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("encoding failure: %w", err)
	}
	return data, nil
}
