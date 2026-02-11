package generator

import (
	"embed"
	"io/fs"
	"os"
	"path/filepath"
)

var templates embed.FS

func Create(name string) error {
	return fs.WalkDir(templates, "templates/server", func(path string, d fs.DirEntry, err error) error {

		target := filepath.Join(name, path[len("templates/server/"):])

		if d.IsDir() {
			return os.MkdirAll(target, 0755)
		}

		data, _ := templates.ReadFile(path)
		return os.WriteFile(target, data, 0644)
	})
}
