// Author: James Mallon <jamesmallondev@gmail.com>
// project package -
package resources

import (
	"TheGorgeousWizard/helpers"
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Struct type Project -
type PrjMkr struct {
	proj     map[string]string
	projOpts map[string][][]string
}

var reader = bufio.NewReader(os.Stdin) // initialize the bufio

// Creat function -
func (this *PrjMkr) CrtPrj(name string) {
	fmt.Println("Let's launch that ship")
	fmt.Println("---------------------")

	this.proj = map[string]string{"path": name}
	helpers.Parse("proj-opts.json", &this.projOpts)

	for k, _ := range this.projOpts {
		this.load(k)
	}
	fmt.Println(this.proj)
}

// createSession function -
func (this *PrjMkr) load(rcr string) {
	fmt.Print(this.projOpts[rcr][0][0])

	entry, _ := reader.ReadString('\n')
	entry = strings.ToLower(strings.Replace(entry, "\n", "", -1)) // convert CRLF to LF

	for _, v := range this.projOpts[rcr][1] {
		if entry == v {
			this.proj[rcr] = entry
			return
		}
	}

	if !(len(entry) > 0) {
		this.proj[rcr] = this.projOpts[rcr][1][0]
	} else {
		fmt.Print("Invalid option")
		this.load(rcr)
	}
}
