package custom_json

import "time"

type CustomTime struct {
	time.Time
}

const customTimeLayout = "2006-01-02 15:04:05"

func (ct CustomTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(customTimeLayout)+2)
	b = append(b, '"')
	b = ct.Time.AppendFormat(b, customTimeLayout)
	b = append(b, '"')
	return b, nil
}

func (ct *CustomTime) UnmarshalJSON(bytes []byte) error {
	if string(bytes) == "null" {
		return nil
	}

	t, err := time.Parse(`"`+customTimeLayout+`"`, string(bytes))
	if err != nil {
		return err
	}

	ct.Time = t
	return nil
}
