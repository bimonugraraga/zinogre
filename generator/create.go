package generator

import (
	"embed"
	"io/fs"
	"os"
	"path/filepath"
)

// embed everything inside generator/templates
//
//go:embed templates/**
var templates embed.FS

func Create(name string) error {
	root := "templates/server"

	return fs.WalkDir(templates, root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// get safe relative path
		rel, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}

		target := filepath.Join(name, rel)

		if d.IsDir() {
			return os.MkdirAll(target, 0755)
		}

		data, err := templates.ReadFile(path)
		if err != nil {
			return err
		}

		return os.WriteFile(target, data, 0644)
	})
}
