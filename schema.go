package yamlschema

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/xeipuuv/gojsonschema"
	"gopkg.in/yaml.v3"
	"strings"
)

func ValidateObjects(yamlSchema map[string]interface{}, obj interface{}) error {
	// Marshal schema object to JSON and load it
	schemaJsonBuf, err := json.Marshal(yamlSchema)
	if err != nil {
		return errors.Wrap(err, "schema json marshal")
	}
	schemaLoader := gojsonschema.NewStringLoader(string(schemaJsonBuf))

	// Marshal object to validate to JSON and load it
	objJsonBuf, err := json.Marshal(obj)
	if err != nil {
		return errors.Wrap(err, "obj json marshal")
	}

	jsonObjectLoader := gojsonschema.NewStringLoader(string(objJsonBuf))

	// Perform the actual validation
	result, err := gojsonschema.Validate(schemaLoader, jsonObjectLoader)
	if err != nil {
		return errors.Wrap(err, "validation")
	}
	if result.Valid() {
		return nil
	} else {
		errorItems := []string{}
		for _, err := range result.Errors() {
			errorItems = append(errorItems, fmt.Sprintf("- %s", err.String()))
		}
		return fmt.Errorf("Validation errors:\n\n%s", strings.Join(errorItems, "\n"))
	}
}

func ValidateStrings(schemaYaml string, objYaml string) error {
	var schemaObj map[string]interface{}
	if err := yaml.Unmarshal([]byte(schemaYaml), &schemaObj); err != nil {
		return errors.Wrap(err, "unmarshal schema")
	}

	var obj interface{}
	if err := yaml.Unmarshal([]byte(objYaml), &obj); err != nil {
		return errors.Wrap(err, "unmarshal obj")
	}

	return ValidateObjects(schemaObj, obj)
}
