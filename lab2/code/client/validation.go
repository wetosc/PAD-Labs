package main

import (
	"io/ioutil"

	"pad.com/lab2/code/eugddc"

	libxml2 "github.com/lestrrat/go-libxml2"
	"github.com/lestrrat/go-libxml2/xsd"
)

func validateXML(xmlData []byte) bool {
	data, err := ioutil.ReadFile("schema.xsd")
	eugddc.CheckError(err, "Failed to read XML Schema file")
	s, err := xsd.Parse(data)
	eugddc.CheckError(err, "Failed to parse XML Schema")
	defer s.Free()

	d, err := libxml2.Parse(xmlData)
	eugddc.CheckError(err, "Failed to parse XML Schema")

	errors := s.Validate(d)
	if errors != nil {
		for _, e := range errors.(xsd.SchemaValidationError).Errors() {
			eugddc.CheckError(e, "[XSD Schema] Error in your XML")
		}
		return false
	}
	return true
}

func validateJSON(jsonData []byte) bool {
	return true
}
