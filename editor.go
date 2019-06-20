package goutils

/**
 * @author Oliver Kelton, oakelton@gmail.com
 * @date Jun 20, 2019
 * Allows calling of the user's text editor (specified by $EDITOR) for reading
 * user input.
 */

import (
	"os"
	"log"
	"io/ioutil"
)

func getEditor (def string) (path string) {
	editor := os.Getenv("EDITOR"); if editor == "" {
		editor = def
	}
	path = Which(editor)
	return
}

func CallEditor (defEditor string, template []byte) (output []byte) {
	tmp, err := ioutil.TempFile("", "goedit_*.md"); if err != nil {
		log.Panic(err)
	}
	fName := tmp.Name()
	defer tmp.Close(); defer os.Remove(fName)
	tmp.Write(template)
	editor := getEditor(defEditor)
	proc, err := os.StartProcess(editor, []string{editor, "+set backupcopy=yes" , fName},
		&os.ProcAttr{Files: []*os.File{os.Stdin, os.Stdout, os.Stderr}})
	if err != nil { log.Panic(err) }
	_, err = proc.Wait(); if err != nil { log.Panic(err) }
	tmp.Seek(0, os.SEEK_SET)
	output, err = ioutil.ReadAll(tmp); if err != nil { log.Panic(err) }
	return output
}//-- end func CallEditor

