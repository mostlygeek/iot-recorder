package eagle

import "strconv"

type hexNumber string

func (h hexNumber) Uint64() uint64 {
	if v, err := strconv.ParseUint(string(h), 0, 64); err != nil {
		return 0
	} else {
		return v
	}
}
