package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/wozuo/droidalyzer"
	"github.com/wozuo/droidalyzer/scanner"
)

var sp bool
var lib bool
var printLib bool
var usageMsg = `usage: droidalyzer [flags] [path]
Flags:
	The -sp flag indicates if a single project should be scanned
	The -lib flag indicates if projects should be scanned
	for networking libraries
	The -printLib flag indicates if information about libraries
	should be printed for every scanned project
	
Examples:
	droidalyzer -lib "~/documents/Android Projects"
	droidalyzer -lib -printLib "~/documents/Android Projects"
	droidalyzer -sp "~/documents/Android Project"
`

func usage() {
	fmt.Fprint(os.Stderr, usageMsg)
	os.Exit(2)
}

func init() {
	flag.BoolVar(&sp, "sp", false, "scan single project")
	flag.BoolVar(&lib, "lib", false,
		"scan projects for networking libraries")
	flag.BoolVar(&printLib, "printLib", false,
		"should print library info for every project")
	flag.Usage = usage
}

func main() {
	flag.Parse()

	if len(os.Args) < 2 {
		fmt.Println("Invalid arguments.")
		return
	}

	if sp {
		fmt.Println("Scanning single project...")

		p := droidalyzer.NewProject(filepath.Base(flag.Args()[0]),
			flag.Args()[0])

		err := p.Scan()
		if err != nil {
			printError(err)
			return
		}

		err = scanner.FindNetworkingCodeInProject(p)
		if err != nil {
			printError(err)
			return
		}
	}

	if lib {
		if sp {
			p := droidalyzer.NewProject(filepath.Base(flag.Args()[0]),
				flag.Args()[0])

			err := p.Scan()
			if err != nil {
				printError(err)
				return
			}

			err = scanner.FindLibraryForProject(p)
			if err != nil {
				printError(err)
				return
			}
		} else {
			fmt.Println("Scanning projects for networking libraries...")

			projects, err := scanner.GetProjects(flag.Args()[0])
			if err != nil {
				printError(err)
				return
			}

			err = scanner.FindLibrariesForProjects(projects, printLib)
			if err != nil {
				printError(err)
				return
			}
		}
	}
}

func printError(err error) {
	fmt.Println("Error: ", err)
}
