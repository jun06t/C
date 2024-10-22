package C

import (
	"fmt"

	"github.com/jun06t/D"
)

const version = "1."

func Func() string {
	fmt.Println(D.Func())
	return version
}
