/**
 * Test editor
 * @author Oliver Kelton, oakelton@gmail.com
 * @date Jul 19, 2019
 */
package .test

import (
	"os"
	"fmt"
	"log"

	"github.com/ollykel/goutils"
)

func main () {
	cfg := goutils.EditorConfig{
		Name: "emacs",
		Flags: []string{"-f", "python-mode"}}
	output, err := goutils.Edit([]byte("\n# This is a python script"), &cfg)
	if err != nil { log.Fatal(err) }
	fmt.Printf("Output:\n`%s`\n", string(output))
}//-- end func main

