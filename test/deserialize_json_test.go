package test

import (
	"encoding/json"
	"testing"

	"github.com/digiseg-labs/api-client-go/openapi"
	"github.com/stretchr/testify/assert"
)

func TestDeserializeJsonWithAdditionalFields(t *testing.T) {
	data := `{
		"id": "id",
		"start_date": "2024-01-01T00:00:00.000Z",
		"end_date": "2024-03-01T00:00:00.000Z",
		"created_at": "2024-01-01T00:00:00.000Z",
		"created_by": "test",
		"new_field_that_doesnt_exist_in_schema": true
	}`
	study := &openapi.StudyFull{}
	err := json.Unmarshal([]byte(data), study)
	assert.NoError(t, err)
	assert.Equal(t, "id", study.Id)
}
