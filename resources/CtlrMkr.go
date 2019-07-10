// Author: James Mallon <jamesmallondev@gmail.com>
// project package -
package resources

import (
	"fmt"
)

// Struct type CtlrMkr -
type CtlrMkr struct {
	options map[string]string
}

// Create method -
func (this *CtlrMkr) CrtCtrlr(name string) {
	fmt.Println("Creating controller: " + name)
}
