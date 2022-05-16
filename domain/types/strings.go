package types

import (
	"database/sql/driver"
	"encoding/json"
	"strings"
)

type Strings []string

func (s *Strings) Scan(value interface{}) error {
	vi := value.([]uint8)
	v := string(vi)
	v = strings.Trim(v, `[`)
	v = strings.Trim(v, `]`)
	v = strings.Replace(v, `"`, "", -1)
	*s = strings.Split(v, ",")

	return nil
}

func (s Strings) Value() (driver.Value, error) {
	b, _ := json.Marshal(s)

	return string(b), nil
}
