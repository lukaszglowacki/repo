package repo

import (
	"encoding/xml"
	"io/ioutil"
	"os"
)

type DefaultXmlElement struct {
	Revision string `xml:"revision,attr"`
	Remote   string `xml:"remote,attr"`
}
type RemoteXmlElement struct {
	Name  string `xml:"name,attr"`
	Fetch string `xml:"fetch,attr"`
}

type RemoteXmlElements []RemoteXmlElement

func (r RemoteXmlElements) ByName(name string) *RemoteXmlElement {
	for _, rm := range r {
		if rm.Name == name {
			return &rm
		}
	}
	return nil
}

type ProjectXmlElement struct {
	Name   string `xml:"name,attr"`
	Path   string `xml:"path,attr"`
	Remote string `xml:"remote,attr"`
}

type Manifest struct {
	XMLName xml.Name            `xml:manifest"`
	Default DefaultXmlElement   `xml:"default"`
	Remote  RemoteXmlElements   `xml:"remote"`
	Project []ProjectXmlElement `xml:"project"`
}

func NewManifestFromFile(path string) (*Manifest, error) {
	xmlFile, err := os.Open(path)
	defer xmlFile.Close()
	if err != nil {
		return nil, err
	}

	byteValue, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		return nil, err
	}

	var m Manifest
	err = xml.Unmarshal(byteValue, &m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}
