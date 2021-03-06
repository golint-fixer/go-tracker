package elastic

import (
	"github.com/index0h/go-tracker/share/types"
)

type elasticVisit struct {
	VisitID   string   `json:"_id"`
	Timestamp string   `json:"@timestamp"`
	SessionID string   `json:"sessionId"`
	ClientID  string   `json:"clientId"`
	Fields    []keyVal `json:"fields"`
}

type keyVal struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func keyValFromHash(data types.Hash) []keyVal {
	result := make([]keyVal, len(data))

	var i uint
	for key, value := range data {
		result[i] = keyVal{Key: key, Value: value}

		i++
	}

	return result
}

func hashFromKeyVal(data []keyVal) types.Hash {
	result := make(types.Hash, len(data))

	for _, element := range data {
		result[element.Key] = element.Value
	}

	return result
}
