package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	p := fmt.Println

	path := filepath.Join("dir1", "dir2", "filename")

	p(filepath.Join("dir1//", "filename"))
	p(filepath.Join("dir1/../dir1", "filename"))

	p("Dir:", filepath.Dir(path))
	p("Base:", filepath.Base(path))
	p(filepath.Split(path))

	p(filepath.IsAbs("dir/file"))
	p(filepath.IsAbs("/dir/file"))

	filename := "config.json"

	ext := filepath.Ext(filename)
	p(ext)
	p(strings.TrimSuffix(filename, ext))

	if rel, err := filepath.Rel("a/b", "a/c/t/file"); err != nil {
		panic(err)
	} else {
		p(rel)
	}

}
