package util

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
)

type versionList []struct {
	AppType        string `json:"AppCode,omitempty"`
	CurrentVersion string `json:"CurrentVersion,omitempty"`
}

type VersionAble interface {
	appCode() string
	currentVersion() string
}

func GetCurrent() (io.Reader, error) {

	data, err := ioutil.ReadFile("current.json")
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(data), err
}

func loadVersions() (versionList, error) {
	data, err := GetCurrent()
	if err != nil {
		return nil, err
	}
	var versions versionList
	err = json.NewDecoder(data).Decode(&versions)
	return versions, err
}

func writeVersion(v VersionAble) error {
	versions, _ := loadVersions()
	for i, version := range versions {
		if version.AppType == v.appCode() && version.CurrentVersion < v.currentVersion() {
			versions[i].CurrentVersion = v.currentVersion()
		}
	}

	file, err := os.Create("current.json")
	if err != nil {
		return err
	}
	defer file.Close()
	err = json.NewEncoder(file).Encode(versions)
	return err
}
