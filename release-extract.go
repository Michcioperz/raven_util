package raven_util

import (
	"log"
	"strings"
	"path"
	"io/ioutil"
)

func ExtractCurrentRelease(repo_root string, verbose_logger *log.Logger) string {
	currentPath := path.Join(repo_root, "legit-fake-path")
	var reference string = "ref: .git/HEAD"
	var err error = nil
	for strings.HasPrefix(reference, "ref:") {
		newSubpath := strings.TrimSpace(strings.TrimPrefix(reference, "ref:"))
		oldPath := path.Dir(currentPath)
		currentPath = path.Join(oldPath, newSubpath)
		if verbose_logger != nil {
			verbose_logger.Print(currentPath)
		}
		var refbytes []byte
		refbytes, err = ioutil.ReadFile(currentPath)
		if err != nil {
			return "undefined"
		}
		reference = string(refbytes)
	}
	return reference
}
