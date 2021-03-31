package yamlschema_test

import (

	"github.com/stretchr/testify/assert"
	"github.com/zefhemel/yamlschema"
	"testing"
	_ "embed"
)

//go:embed test/test.schema.yml
var testSchema string

func TestValidateStrings(t *testing.T) {
	assert.NoError(t, yamlschema.ValidateStrings(testSchema, `
url: http://localhost
token: abcdef
events:
  bla:
  - trigger1
  - trigger2
`))
	assert.Contains(t, yamlschema.ValidateStrings(testSchema, `
url: http://localhost
`).Error(), "token is required")
	assert.Contains(t, yamlschema.ValidateStrings(testSchema, `
url: http://localhost
token: 1234
`).Error(), "Invalid type")
}
