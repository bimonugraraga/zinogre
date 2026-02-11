package generator

import (
	"embed"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

//go:embed templates/**
var templates embed.FS

func Create(name string, modulePath string) error {
	root := "templates/server"

	return fs.WalkDir(templates, root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

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

		content := string(data)

		if strings.HasSuffix(path, ".go") {
			tmpl, err := template.New(rel).Parse(content)
			if err != nil {
				return err
			}

			outFile, err := os.Create(target)
			if err != nil {
				return err
			}
			defer outFile.Close()

			return tmpl.Execute(outFile, struct {
				ModulePath string
			}{
				ModulePath: modulePath,
			})
		}

		return os.WriteFile(target, data, 0644)
	})
}
