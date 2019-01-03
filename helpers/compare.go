package helpers

import (
	"bytes"
	"encoding/json"
)

func Compare(items ...interface{}) (bool, []int, error) {
	listOfByteLengths := []int{}

	if len(items) > 1 {
		current, err := json.Marshal(items[0])

		if err != nil {
			return false, listOfByteLengths, err
		}

		for _, val := range items {
			mar, err := json.Marshal(val)

			listOfByteLengths = append(listOfByteLengths, len(mar))

			if !bytes.Equal(mar, current) {
				return false, listOfByteLengths, err
			}
		}
	} else {
		return false, listOfByteLengths, nil
	}
	return true, listOfByteLengths, nil
}
