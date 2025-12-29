package util

import (
	"fmt"
	"net/http"
	"strings"
)

var basePath string

// SetBasePath replaces basePath with a new one
func SetBasePath(newPath string) {
	basePath = newPath
}

// ReplaceBasePath replaces a pattern in a string, injecting basePath into it
func ReplaceBasePath(str, old, newPattern string) string {
	return strings.Replace(str, old, fmt.Sprintf(newPattern, basePath), 1)
}

// GetFullPath preprends basePath to a string
func GetFullPath(path string) string {
	return fmt.Sprintf("%s%s", basePath, path)
}

// TrimBasePath removes basePath from the beginning of a string
func TrimBasePath(r *http.Request, prefix string) string {
	return strings.TrimPrefix(r.URL.Path, GetFullPath(prefix))
}
