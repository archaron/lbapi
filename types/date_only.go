package types

import (
	"encoding/json"
	"time"
)

type DateOnly time.Time

func (d *DateOnly) String() string {
	if time.Time(*d).IsZero() {
		return ""
	}
	return time.Time(*d).Format(time.DateOnly)
}

func (d *DateOnly) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}

func (d *DateOnly) UnmarshalJSON(data []byte) error {

	var timeStr string
	if err := json.Unmarshal(data, &timeStr); err != nil {
		return err
	}

	t, err := time.Parse(time.DateOnly, timeStr)
	if err != nil {
		return err
	}

	*d = DateOnly(t)

	return nil
}
