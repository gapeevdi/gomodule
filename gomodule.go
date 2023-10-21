package gomodule

import "fmt"

// git commit -a -m "v1.1.0"
// git push
// git tag v1.1.0
// git push -q origin v1.1.0

func Version() {
	fmt.Println("Version 4.1.0")
}
