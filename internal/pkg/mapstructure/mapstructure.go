package mapstructure

import (
	"github.com/mitchellh/mapstructure"
)

func Decode(input interface{}, result interface{}) error {
	return mapstructure.Decode(input, result)
}
