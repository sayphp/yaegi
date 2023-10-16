package fromage

import (
	"fmt"

	"guthib.com/sayphp/cheese"
	"guthib.com/sayphp/fromage/couteau"
)

func Hello() string {
	return fmt.Sprintf("Fromage %s %s", couteau.Hello(), cheese.Hello())
}
