package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/cmh2166/elag18apis/taquito/datautils"
)

func TestMarshal(t *testing.T) {
	input := jsonData()
	result, err := MarshalMap(input)
	assert.Nil(t, err)
	// Empty lists should remain a list.
	assert.NotNil(t, result["structural"].M["hasMember"].L)
	assert.Equal(t, 0, len(result["structural"].M["hasMember"].L))

	var output datautils.JSONObject
	err = UnmarshalMap(result, &output)
	assert.Nil(t, err)

	structural := output.GetObj("structural")
	// Empty lists should remain a list (not changed to nil)
	assert.NotNil(t, (*structural)["hasMember"])
	assert.Equal(t, input, output)
}
