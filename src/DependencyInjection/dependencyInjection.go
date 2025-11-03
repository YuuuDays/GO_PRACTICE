package dependencyinjection

import (
	"bytes"
	"fmt"
)

func Greet(wirter *bytes.Buffer, name string) {
	fmt.Println("Hello, %s", name)
}
