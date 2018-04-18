/* webpage_save.go */
/*
modification history
--------------------
2017/07/21, by Xiongmin LIN, create
*/
/*
DESCRIPTION
*/
package web_package

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"path"
)

const (
	OutputFileMode = 0644
)

/*
generate root directory to save web page

Params:
	- rootPath: root path for saving file
Returns:
	- error: any failure
*/
func genRootDir(rootPath string) error{
	if _, err := os.Stat(rootPath); os.IsNotExist(err) {
		if os.MkdirAll(rootPath, 0777) != nil {
			return fmt.Errorf("os.MkdirAll(%s):%s", rootPath, err.Error())
		}
	}

	return nil
}

/*
generate file path for given url

Params:
	- url: url to crawl
	- rootPath: root path for saving file

Returns:
	- file path
*/
func genFilePath(urlStr, rootPath string) string {
	filePath := url.QueryEscape(urlStr)
	filePath = path.Join(rootPath, filePath)
	return filePath
}

/*
save web page of given url to file

Params:
	- rootPath: root path for saving file
	- url: url to crawl
	- data: data to save
Returns:
	- error: any failure
*/
func SaveWebPage(rootPath string, url string, data []byte) error {
	// create root dir, if not exist
	if err := genRootDir(rootPath); err != nil {
		return fmt.Errorf("genRootDir(%s):%s", rootPath, err.Error())
	}

	// generate full file path
	filePath := genFilePath(url, rootPath)

	// save to file
	err := ioutil.WriteFile(filePath, data, OutputFileMode)
	if err != nil {
		return fmt.Errorf("ioutil.WriteFile(%s):%s", filePath, err.Error())
	}

	return nil
}
