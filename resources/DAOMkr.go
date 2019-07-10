// Author: James Mallon <jamesmallondev@gmail.com>
// project package -
package resources

import (
	"fmt"
)

// Struct type DAOMkr -
type DAOMkr struct {
	options map[string]string
}

// CreateDAO method -
func (this *DAOMkr) CrtDAO(name string) {
	fmt.Println("Creating DAO: " + name)
}
