package droidalyzer

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// NCode holds information about networking code occurrences
type NCode struct {
	Line    uint
	Keyword string
}

// APFile represents an Android project file
type APFile struct {
	Name      string
	Path      string
	Extension string
	NCode     []NCode
}

// Scan scans a file for keywords (case insensitive)
func (apf *APFile) Scan(keywords *[]string) error {
	apfile, err := os.Open(apf.Path)
	if err != nil {
		return err
	}
	defer func() {
		err := apfile.Close()
		if err != nil {
			panic(err)
		}
	}()

	var lineNbr uint = 1

	scanner := bufio.NewScanner(apfile)
	for scanner.Scan() {
		for _, keyword := range *keywords {
			if strings.Contains(strings.ToLower(scanner.Text()),
				strings.ToLower(keyword)) {
				nc := NCode{lineNbr, keyword}

				apf.NCode = append(apf.NCode, nc)
			}
		}
		lineNbr++
	}

	return nil
}

// PrintNCSResults prints results of file scan
func (apf *APFile) PrintNCSResults() {
	for _, ncode := range apf.NCode {
		fmt.Println("Line:", ncode.Line, ", keyword:", ncode.Keyword)
	}
}

// NewAPFile creates and returns a new APFile object
func NewAPFile(name string, path string, extension string) *APFile {
	apf := &APFile{}

	apf.Name = name
	apf.Path = path
	apf.Extension = extension
	apf.NCode = []NCode{}

	return apf
}
