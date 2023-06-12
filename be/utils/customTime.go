package utils

import (
	"time"
)

type Unmarshaler interface {
	UnmarshalJSON([]byte) error
}

type Date string

func (d *Date) UnmarshalJSON(bytes []byte) error {
	dd, err := time.Parse(`"2006-01-02T15:04:05.000+0000"`, string(bytes))
	if err != nil {
		return err
	}
	*d = Date(dd.Format("01/02/2006"))

	return nil
}
