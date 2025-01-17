package pkg

import (
	"fmt"

	"guthib.com/sayphp/vin"
)

func Here() string {
	return "root"
}

func NewSample() func() string {
	return func() string {
		return fmt.Sprintf("%s %s", Here(), vin.Hello())
	}
}
