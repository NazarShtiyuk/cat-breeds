package util

import (
	"encoding/json"
	"fmt"
	"os"
)

func WriteJSONFile(data interface{}, fileName string) error {
	dataJSON, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("json marshal error: %v", err)
	}

	err = os.WriteFile(fileName, dataJSON, 0644)
	if err != nil {
		return fmt.Errorf("write file error: %v", err)
	}

	return nil
}
