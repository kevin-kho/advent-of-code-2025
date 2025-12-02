package common

import "os"

func ReadInput(filePath string) ([]byte, error) {
	var data []byte

	data, err := os.ReadFile(filePath)
	if err != nil {
		return data, err
	}

	return data, nil

}
