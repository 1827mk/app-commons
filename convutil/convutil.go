package convutil

import (
	"encoding/json"
	"fmt"
)

func Obj2Map(obj interface{}) map[string]any {
	// Marshal the object into JSON
	jsonData, err := json.Marshal(obj)
	if err != nil {
		fmt.Printf("Error marshalling object to JSON: %v\n", err)
		return nil
	}

	// Unmarshal JSON into a map
	var result map[string]any
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		fmt.Printf("Error unmarshalling JSON to map: %v\n", err)
		return nil
	}

	return result
}
