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
	"time"
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

type Replacer struct {
	demarc		byte
	defs		map[string]string
	rep			*strings.Replacer
}//-- end type Replacer

func (rep *Replacer) refresh () {
	pairs := make([]string, 0, len(rep.defs) << 2)
	for k, v := range rep.defs {
		varName := make([]rune, len(k) + 1)
		varName[0] = rune(rep.demarc)
		copy(varName[1:], []rune(k))
		pairs = append(pairs, string(varName), v)
	}
	rep.rep = strings.NewReplacer(pairs...)
}//-- end func Replacer.refresh

func NewReplacer (demarc byte, defs map[string]string) (rep Replacer) {
	rep.demarc, rep.defs = demarc, make(map[string]string)
	for k, v := range defs {
		rep.defs[k] = v
	}
	rep.refresh()
	return
}//-- end func NewReplacer

func (rep *Replacer) Get (key string) (val string, exists bool) {
	val, exists = rep.defs[key]
	return
}//-- end func Replacer.Get

func (rep *Replacer) Set (key, val string) {
	rep.defs[key] = val
	rep.refresh()
}//-- end func Replacer.Set

func (rep *Replacer) Remap (mapper func(string)string) (self *Replacer) {
	for k, v := range rep.defs {
		rep.defs[k] = mapper(v)
	}//-- end for range rep.defs
	rep.refresh()
	return rep
}//-- end func Replacer.Remap

func (rep *Replacer) Replace (orig string) string {
	return rep.rep.Replace(orig)
}//-- end func Replacer.Replace

func DisplayMsg (dest io.Writer, msg string, msecs int) {
	dest.Write([]byte(msg))
	time.Sleep(time.Duration(msecs) * time.Millisecond)
}

