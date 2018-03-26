package service

import (
	"io/ioutil"
	"path"
	"runtime"

	"github.com/xeipuuv/gojsonschema"
	yaml "gopkg.in/yaml.v2"
)

func schemaFilePath() (filepath string) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("Cannot retrieve the path for the JSON schema")
	}
	filepath = path.Join(path.Dir(filename), "./schema.json")
	return
}

func convert(i interface{}) interface{} {
	switch x := i.(type) {
	case map[interface{}]interface{}:
		m2 := map[string]interface{}{}
		for k, v := range x {
			m2[k.(string)] = convert(v)
		}
		return m2
	}
	return i
}

// ValidServiceData returns true if the file is a valid service, false otherwise
// The all validation can be found in https://github.com/mesg-foundation/application/tree/dev/service/schema.json
func ValidServiceData(body interface{}) (result *gojsonschema.Result, err error) {
	schema := gojsonschema.NewReferenceLoader("file://" + schemaFilePath())
	data := gojsonschema.NewGoLoader(body)
	result, err = gojsonschema.Validate(schema, data)
	return
}

// ValidServiceFile returns true is the file is a valid service, false otherwise
func ValidServiceFile(filepath string) (result *gojsonschema.Result, err error) {
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		return
	}
	var body interface{}
	if err = yaml.Unmarshal([]byte(file), &body); err != nil {
		return
	}

	body = convert(body)
	result, err = ValidServiceData(body)
	return
}

// IsValid returns true if the service is valid, false otherwise
// The all validation can be found in https://github.com/mesg-foundation/application/tree/dev/service/schema.json
func (service *Service) IsValid() (result *gojsonschema.Result, err error) {
	result, err = ValidServiceData(service)
	return
}