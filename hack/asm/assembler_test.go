package asm

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

var testdata = "testdata"

func TestAssemble(t *testing.T) {
	list, err := ioutil.ReadDir(testdata)
	if err != nil {
		t.Fatal(err)
	}

	for _, fi := range list {
		filename := fi.Name()
		if !fi.IsDir() && !strings.HasPrefix(filename, ".") && strings.HasSuffix(filename, ".asm") {
			file, err := os.Open(filename)
			if err != nil {
				t.Fatal(err)
			}
			Assemble(file)
		}
	}
}
