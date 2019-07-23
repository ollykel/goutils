/**
 * <DESCRIPTION>
 * @author Oliver Kelton, oakelton@gmail.com
 * @date Jul 05, 2019
 */
package goutils

import (
	"io"
	"log"
	"strings"
	"encoding/json"
)

func PrintJSON (dest io.Writer, src interface{}) {
	enc := json.NewEncoder(dest)
	enc.SetIndent("", "  ")
	if err := enc.Encode(src); err != nil { log.Panic(err) }
}

// orig: original string with varnames included
// demarc: character that demarcs the beginning of a variable name
// vars: mapping of variable names to values
func ParseVars (demarc byte, vars map[string]string, orig ...string) (parsed []string) {
	pairs := make([]string, 0, len(vars) * 2)
	for k, v := range vars {
		varName := make([]rune, len(k) + 1)
		varName[0] = rune(demarc)
		copy(varName[1:], []rune(k))
		pairs = append(pairs, string(varName), v)
	}//-- end for k, v
	rep := strings.NewReplacer(pairs...)
	parsed = make([]string, len(orig))
	for i, s := range orig {
		parsed[i] = rep.Replace(s)
	}//-- end for range orig
	return
}

