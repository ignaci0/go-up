package decorator

import (
	"github.com/stretchr/testify/assert"
	"github.com/ignaci0/go-up/reader"
	"testing"
)

func Test_ShouldLowerCaseTheKeys(t *testing.T) {

	properties := reader.NewProgrammaticReader()
	properties.AddProperty("KEY_one", reader.Property{"ENV_ONE", false})
	properties.Add("KEY.TWO", "VALUE_TWO.")
	properties.Add("key.three", "value.three")

	lower := &ToLowerCaseKeyDecoratorReader{properties}
	lowerProp, _ := lower.Read()

	assert.Equal(t, "ENV_ONE", lowerProp["key_one"].Value)
	assert.Equal(t, "VALUE_TWO.", lowerProp["key.two"].Value)
	assert.Equal(t, "value.three", lowerProp["key.three"].Value)

}
