// Author: James Mallon <jamesmallondev@gmail.com>
// project package -
package resources

import (
	"flag"
	"fmt"
)

// Struct type Project -
type PrjMkr struct {
	proj map[string]string
}

// Create function -
func (this *PrjMkr) CrtPrj(name string) {
	db := flag.Lookup("db").Value
	sess := flag.Lookup("sess").Value

	fmt.Println("Project was created in the path", name)
	fmt.Println("Storing option:", db)
	fmt.Println("Session approach:", sess)
}
