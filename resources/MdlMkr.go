// Author: James Mallon <jamesmallondev@gmail.com>
// project package -
package resources

import (
	"fmt"
)

// Struct type MdlMkr -
type MdlMkr struct {
	options map[string]string
}

// CreateModel method -
func (this *MdlMkr) CrtMdl(name string) {
	fmt.Println("Creating Model: " + name)
}
