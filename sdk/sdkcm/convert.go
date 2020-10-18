package sdkcm

import "strconv"

// convert string to uint32
func StrToUint(s string) (uint32, error) {
	id, err := strconv.ParseUint(s, 10, 16)

	if err != nil {
		return 0, err
	}

	return uint32(id), nil
}
