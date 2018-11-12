package scanner

import (
	"fmt"
	"io/ioutil"
	"math"
	"path/filepath"
	"runtime"
	"sync"

	"github.com/wozuo/droidalyzer"
)

var libraries = []string{
	// JSON & XML libraries
	"com.google.code.gson",
	"com.squareup.moshi",
	"converter-jaxb",
	"converter-simplexml",
	"converter-jackson",
	"com.fasterxml.jackson",
	"javax.xml.stream:stax-api",
	"com.stanfy:gson-xml-java",
	// Networking libraries
	"com.squareup.retrofit",
	"com.squareup.okhttp",
	"com.github.bumptech.glide",
	"org.apache.httpcomponents",
	"com.android.volley",
	"com.mcxiaoke.volley", // Deprecated library implementation
	"android-async-http",
	"cz.msebera.android:httpclient", // dep from android async
	"com.koushikdutta.ion",
	"com.koushikdutta.async",
}

var sPrint = true
var wg sync.WaitGroup

func worker(pChan chan *droidalyzer.Project) {
	for p := range pChan {
		err := FindLibraryForProject(p)
		if err != nil {
			panic(err)
		}

		wg.Done()
	}
}

// FindLibrariesForProjects scans Android projects for
// networking libraries.
// The shouldPrint parameter indicates if found libraries
// are printed for every project.
func FindLibrariesForProjects(ps *[]droidalyzer.Project,
	shouldPrint bool) error {
	sPrint = shouldPrint

	pChan := make(chan *droidalyzer.Project, len(*ps))

	for i := 0; i < runtime.NumCPU(); i++ {
		go worker(pChan)
	}

	for i := range *ps {
		wg.Add(1)
		pChan <- &(*ps)[i]
	}

	wg.Wait()
	close(pChan)

	libOcc := make(map[string]int)

	for _, lib := range libraries {
		libOcc[lib] = 0
	}

	noLibs := 0

	for _, project := range *ps {
		if len(project.Libraries) == 0 {
			noLibs++
		}
		for lib := range project.Libraries {
			libOcc[lib] = libOcc[lib] + 1
		}
	}

	fmt.Println("Stats for", len(*ps), "projects:")
	for l, occ := range libOcc {
		fmt.Println("Library", l, "is used in", occ, "projects")
	}

	usingLibs := len(*ps) - noLibs
	p := math.Floor(float64(usingLibs)/float64(len(*ps))*100 + .5)
	fmt.Println(p, "% (", usingLibs, "out of", len(*ps),
		") use at least one (known) networking library")

	return nil
}

// FindLibraryForProject scans a single Android project for
// networking libraries
func FindLibraryForProject(project *droidalyzer.Project) error {
	err := project.Scan()
	if err != nil {
		return err
	}

	for _, apf := range project.APFiles {
		if apf.Extension == ".gradle" {
			err := apf.Scan(&libraries)
			if err != nil {
				return err
			}

			if len(apf.NCode) > 0 {
				for _, lib := range apf.NCode {
					project.Libraries[lib.Keyword] = struct{}{}
				}
			}
		}
	}

	if sPrint && len(project.Libraries) > 0 {
		fmt.Println("Found libraries in project:",
			project.Name)
		project.PrintLibs()
		fmt.Println("=============================")
	}

	return nil
}

// GetProjects returns all projects at a path
func GetProjects(path string) (*[]droidalyzer.Project, error) {
	files, err := ioutil.ReadDir(path)
	var projects []droidalyzer.Project
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if file.IsDir() {
			p := droidalyzer.NewProject(file.Name(),
				filepath.Join(path, file.Name()))

			projects = append(projects, *p)
		}
	}

	return &projects, nil
}
