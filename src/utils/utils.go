package utils

import (
	"io/ioutil"
	"path"
	"runtime"
)

// Check if error and panic
func Check(e error) {
  if e != nil {
    panic(e)
  }
}

// LoadInputFile to load file with name from input
func LoadFile (fileName string) string {
  _, dirName, _, _ := runtime.Caller(1)
  filePath := path.Dir(dirName) + "/" + fileName

  data, err := ioutil.ReadFile(filePath)
  Check(err)
  return string(data)
}
