/**
 * <DESCRIPTION>
 * @author Oliver Kelton, oakelton@gmail.com
 * @date Jul 05, 2019
 */
package goutils

import (
	"io"
	"log"

	"encoding/json"
)

func PrintJSON (dest io.Writer, src interface{}) {
	enc := json.NewEncoder(dest)
	enc.SetIndent("", "  ")
	if err := enc.Encode(dest); err != nil { log.Panic(err) }
}

