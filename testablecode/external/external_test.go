package external_test

import (
	"fmt"

	"github.com/bketelsen/testingclass/testablecode/external"
)

func ExampleMultiply() {
	res := external.Multiply(1, 2)
	fmt.Println(res)
	// Output: 2
}
