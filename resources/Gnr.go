// Author: James Mallon <jamesmallondev@gmail.com>
// project package -
package resources

import (
	"fmt"
)

// Struct type DAOMkr -
type Gnr struct {
	options map[string]string
}

// ListResources method -
func (this *Gnr) LstRsr(name ...string) {
	fmt.Println("Listing Resources...")
}
