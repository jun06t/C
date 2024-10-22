package C

import (
	"fmt"

	"github.com/jun06t/D"
)

const version = "1.2.0"

func Func() string {
	fmt.Println(D.Func())
	return version
}
