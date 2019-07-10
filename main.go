// Author: James Mallon <jamesmallondev@gmail.com>
// Main package
package main

import (
	"TheGorgeousWizard/helpers"
	. "TheGorgeousWizard/resources"
	"flag"
	"reflect"
)

type AppMux struct {
	CtlrMkr
	PrjMkr
	DAOMkr
	MddwMkr
	MdlMkr
	Gnr
}

// main function - Garden of Eden
func main() {
	wolv := map[string]*string{}

	var cmds map[string][]string
	helpers.Parse("cmds.json", &cmds)

	for k, v := range cmds {
		wolv[k] = flag.String(v[0], v[1], v[2])
	}

	flag.Parse()
	app := AppMux{}

	for k, v := range wolv {
		if len(*v) > 0 {
			_ = reflect.ValueOf(&app).MethodByName(cmds[k][3]).
				Call([]reflect.Value{
					reflect.ValueOf(*v),
				})
		}
	}
}
