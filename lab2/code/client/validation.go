package main

import (
	"io/ioutil"

	"github.com/lestrrat/go-jsschema"
	"github.com/lestrrat/go-jsval/builder"

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

func validateJSON(jsonData interface{}) bool {
	s, err := schema.ReadFile("schema.json")
	eugddc.CheckError(err, "Failed to read JSON Schema file")

	b := builder.New()
	v, err := b.Build(s)
	eugddc.CheckError(err, "Failed to build JSON Schema validator")

	err = v.Validate(jsonData)
	eugddc.CheckError(err, "[JSON Schema] Error in your JSON")
	return err == nil
}
