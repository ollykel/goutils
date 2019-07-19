/**
 * Manual testing for editor.go
 * @author Oliver Kelton, oakelton@gmail.com
 * @date Jul 19, 2019
 */
package main

import (
	"fmt"
	"log"

	"github.com/ollykel/goutils"
)

func unitTestEdit (name string, content []byte, cfg *goutils.EditorConfig) {
	fmt.Printf("Testing Edit %s:\n", name)
	output, err := goutils.Edit(content, cfg)
	if err != nil { log.Panic(err) }
	fmt.Printf("Output:\n`%s`\n", string(output))
}//-- end func unitTestEdit

func main () {
	cfg := goutils.EditorConfig{
		Name: "emacs",
		Flags: []string{"-f", "python-mode"}}
	unitTestEdit("emacs", []byte("\n# This is a python script"), &cfg)
	cfg.Name, cfg.Flags = "vim", []string{"+set backupcopy=yes", "+setf python"}
	unitTestEdit("vim", []byte("\n# This is also a python script"), &cfg)
}//-- end func main

