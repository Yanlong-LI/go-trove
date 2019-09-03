package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func Load(filePath string) (TrovePackage, error) {
	var trovePackage TrovePackage
	file, err := os.Open(filePath)
	if err != nil {
		//log.Fatal("Profile does not exist", err)
		return trovePackage, err
	}

	fileByte, err := ioutil.ReadAll(file)
	if err != nil {
		//log.Fatal("Read Packet Configuration Error", err)
		return trovePackage, err
	}
	defer file.Close()

	err = json.Unmarshal(fileByte, &trovePackage)
	if err != nil {
		//log.Fatal("Failure to configure switching structure", err)
		return trovePackage, err
	}

	return trovePackage, nil
}
