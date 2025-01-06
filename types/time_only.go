package types

import (
	"encoding/json"
	"fmt"
	"time"
)

type TimeOnly time.Duration

func (d *TimeOnly) String() string {
	hours := int(time.Duration(*d).Hours())
	minutes := int(time.Duration(*d).Minutes()) % 60
	seconds := int(time.Duration(*d).Seconds()) % 60
	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
}

func (d *TimeOnly) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}

func (d *TimeOnly) UnmarshalJSON(data []byte) error {

	var timeStr string
	if err := json.Unmarshal(data, &timeStr); err != nil {
		return err
	}

	t, err := time.Parse("15:04:05", timeStr)
	if err != nil {
		return err
	}

	*d = TimeOnly(time.Duration(t.Hour())*time.Hour +
		time.Duration(t.Minute())*time.Minute +
		time.Duration(t.Second())*time.Second)

	return nil
}
