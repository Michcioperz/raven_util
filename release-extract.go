package raven_util

import (
	"os"
	"log"
	"strings"
	"path"
	"io/ioutil"
)

func ExtractCurrentRelease(repo_root string) string {
	e := log.New(os.Stderr, "researching: ", log.LstdFlags)
	currentPath := repo_root
	var reference string = "ref: .git/HEAD"
	var err error = nil
	for strings.HasPrefix(reference, "ref:") {
		newSubpath := strings.TrimPrefix(reference, "ref:")
		oldPath := path.Dir(currentPath)
		currentPath := path.Join(oldPath, newSubpath)
		e.Print(currentPath)
		var refbytes []byte
		refbytes, err = ioutil.ReadFile(currentPath)
		if err != nil {
			return "undefined"
		}
		reference = string(refbytes)
	}
	return reference
}
