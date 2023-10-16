package vin

import (
	"fmt"

	"guthib.com/sayphp/cheese"
)

func Hello() string {
	return fmt.Sprintf("vin %s", cheese.Hello())
}
