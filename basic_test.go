package hexdump_test

import (
	"strings"

	yup "github.com/gloo-foo/framework"
	. "github.com/yupsh/hexdump"
)

func ExampleHexdump_basic() {
	// echo "Hello" | hexdump
	yup.MustRun(
		Hexdump(strings.NewReader("Hello")),
	)
	// Output:
	// 00000000  48 65 6c 6c 6f                                    |Hello|
	// 00000005
}

