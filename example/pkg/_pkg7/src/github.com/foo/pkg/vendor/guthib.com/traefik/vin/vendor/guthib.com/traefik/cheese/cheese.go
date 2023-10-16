package cheese

import (
	"fmt"

	"guthib.com/sayphp/fromage"
)

func Hello() string {
	return fmt.Sprintf("cheese %s", fromage.Hello())
}
