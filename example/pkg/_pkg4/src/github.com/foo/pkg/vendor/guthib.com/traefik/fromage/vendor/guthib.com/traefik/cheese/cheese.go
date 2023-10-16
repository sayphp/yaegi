package cheese

import (
	"fmt"

	"guthib.com/sayphp/cheese/vin"
)

func Hello() string {
	return fmt.Sprintf("Cheese %s", vin.Hello())
}
