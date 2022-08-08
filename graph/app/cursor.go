package app

import (
	"encoding/base64"
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

var validCursor = regexp.MustCompile(`.+,(\d+)`)

func decodeCursor(c *string) (int, error) {
	if c == nil {
		return 0, errors.New("nil cursor")
	}

	decoded, err := base64.StdEncoding.DecodeString(*c)
	if err != nil {
		return 0, errors.New("error decode cursor")
	}

	arr := validCursor.FindStringSubmatch(string(decoded))

	id, err := strconv.Atoi(arr[1])
	if err != nil {
		return 0, errors.New("error id not number")
	}
	return id, nil

}
func encodeCursor(table string, id uint) *string {
	c := fmt.Sprintf("%s,%d", table, id)
	e := base64.StdEncoding.EncodeToString([]byte(c))
	return &e
}
