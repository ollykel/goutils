package goutils

/**
 * @author Oliver Kelton, oakelton@gmail.com
 * @date Jun 20, 2019
 * os-based utils
 */

import (
	"os"
	"fmt"
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

func FileExists (fname string) bool {
	_, err := os.Stat(fname)
	return err == nil
}

func Call (exec string, args []string, attr *os.ProcAttr) (proc *os.Process, err error) {
	path := Which(exec); if path == "" {
		return nil, fmt.Errorf("program '%s' not found", exec)
	}
	if attr == nil {
		defAttr := os.ProcAttr{Files: []*os.File{os.Stdin, os.Stdout, os.Stderr}}
		attr = &defAttr
	}
	argv := make([]string, len(args) + 1)
	argv[0] = path
	copy(argv[1:], args)
	return os.StartProcess(path, argv, attr)
}//-- end func Call

