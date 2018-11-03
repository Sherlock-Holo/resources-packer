package main

var tplString = `package {{.module}}

import (
	"golang.org/x/tools/godoc/vfs"
	"golang.org/x/tools/godoc/vfs/mapfs"
)

var m = map[string]string{
{{range .resources}}
    "{{.Path}}": {{.Content}},
{{end}}
}

func NewFS() vfs.FileSystem {
	return mapfs.New(m)
}
`
