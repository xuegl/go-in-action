package custom_json

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCustomTime_MarshalJSON(t *testing.T) {
	ct := CustomTime{time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local)}
	b, err := json.Marshal(ct)
	assert.Nil(t, err)
	assert.Equal(t, `"2023-01-01 00:00:00"`, string(b))
}

func TestCustomTime_UnmarshalJSON(t *testing.T) {
	b := []byte(`"2023-01-01 00:00:00"`)
	ct := CustomTime{}
	err := json.Unmarshal(b, &ct)
	assert.Nil(t, err)
	assert.Equal(t, 2023, ct.Year())
	assert.Equal(t, time.January, ct.Month())
	assert.Equal(t, 01, ct.Day())
	assert.Equal(t, 0, ct.Hour())
	assert.Equal(t, 0, ct.Minute())
	assert.Equal(t, 0, ct.Second())
}
