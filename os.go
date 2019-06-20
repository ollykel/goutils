package goutils

/**
 * @author Oliver Kelton, oakelton@gmail.com
 * @date Jun 20, 2019
 * os-based utils
 */

import (
	"os"
	"strings"
	fpath "path/filepath"
)

// Searches the PATH env variable and returns the first path to match the
// executable name
// Returns "" if no executable found
func Which (executable string) (path string) {
	paths := strings.Split(os.Getenv("PATH"), string(os.PathListSeparator))
	for _, p := range paths {
		path = fpath.Join(p, executable)
		_, err := os.Stat(path)
		if err == nil { return }
	}//-- end for range paths
	return ""
}//-- end func Which

