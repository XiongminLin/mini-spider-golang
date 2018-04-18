/* seedfile_load.go: load seed file*/
/*
modification history
--------------------
2017/07/16, by Xiongmin LIN, create
*/
/*
DESCRIPTION
*/

package mini_spider

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

/*
load seed data from file

Params:
	- filePath: full path of seed file

Returns:
	- urls
	- error
*/
func LoadSeedFile(filePath string) ([]string, error) {
	// read json data from file
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("ioutil.ReadFile(%s): %s", filePath, err.Error())
	}

	// decode json data
	var seeds []string
	err = json.Unmarshal(data, &seeds)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal(): %s", err.Error())
	}

	return seeds, nil
}