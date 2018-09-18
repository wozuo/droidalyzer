package droidalyzer

import (
	"fmt"
	"os"
	"path/filepath"
)

// Project represents an Android project
type Project struct {
	Name      string
	Path      string
	Libraries map[string]struct{}
	APFiles   []APFile
}

// NewProject creates and returns a new Project object
func NewProject(name string, path string) *Project {
	p := &Project{}

	p.Name = name
	p.Path = path
	p.Libraries = make(map[string]struct{})

	return p
}

// Scan scans all files of the project and loads them into APFiles
func (p *Project) Scan() error {
	err := filepath.Walk(p.Path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		f := NewAPFile(info.Name(), path, filepath.Ext(path))

		p.APFiles = append(p.APFiles, *f)

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

// PrintLibs prints libraries contained in Project struct
func (p *Project) PrintLibs() {
	for libs := range p.Libraries {
		fmt.Println("Library:", libs)
	}
}
