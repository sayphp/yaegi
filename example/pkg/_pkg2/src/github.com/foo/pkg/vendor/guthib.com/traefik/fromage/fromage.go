package fromage

import (
	"fmt"

	"guthib.com/sayphp/cheese"
)

func Hello() string {
	return fmt.Sprintf("Fromage %s", cheese.Hello())
}
