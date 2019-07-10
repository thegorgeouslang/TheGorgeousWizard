// Author: James Mallon <jamesmallondev@gmail.com>
// files package -
package helpers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// Parse a json file and fill the received object
func Parse(fname string, obj interface{}) {
	fpath := os.Getenv("GOPATH") + "/src/TheGorgeousWizard/json/" + fname
	jsonFile, e := os.Open(fpath)
	if e != nil {
		log.Fatal(e.Error())
	}
	jsonData, e := ioutil.ReadAll(jsonFile)
	if e != nil {
		log.Fatal(e.Error())
	}
	json.Unmarshal(jsonData, &obj)
	jsonFile.Close()
}
