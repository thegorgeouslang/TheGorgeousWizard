// Author: James Mallon <jamesmallondev@gmail.com>
// project package -
package resources

import (
	"fmt"
)

// Struct type MddwMkr -
type MddwMkr struct {
	options map[string]string
}

// CreateMiddleware method -
func (this *MddwMkr) CrtMddw(name string) {
	fmt.Println("Creating Middleware: " + name)
}
