package types

import (
	"encoding/json"
	"strings"
	"time"
)

type LBTime time.Time

func (d *LBTime) String() string {
	if time.Time(*d).IsZero() {
		return ""
	}
	return time.Time(*d).Format(time.DateTime)
}

func (c *LBTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.String())
}

func (c *LBTime) UnmarshalJSON(b []byte) error {

	// nil if format is incorrect
	if len(b) != 21 && len(b) != 12 {
		return nil
	}

	date := strings.Trim(string(b), "\"")

	if date == "0000-00-00 00:00:00" || date == "0000-00-00" {
		return nil
	}

	var dateFormat = time.DateTime

	if len(date) == 10 {
		dateFormat = time.DateOnly
	}

	t, err := time.Parse(dateFormat, date)
	if err != nil {
		return err
	}

	*c = LBTime(t)
	return nil
}
