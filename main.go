package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gitlawr/mergeyaml/merger"
)

var (
	formerFile string
	latterFile string
	outputFile string
)

func init() {

	flag.Usage = func() {
		fmt.Fprintf(os.Stdout, "Usage:\n%s <formerYamlFile> <latterFile>\n", os.Args[0])
		flag.PrintDefaults()
	}

	// parse flags
	flag.StringVar(&outputFile, "output", "merge_output.yml", "output file")
	flag.StringVar(&outputFile, "o", "merge_output.yml", "output file (shorthand)")
	flag.Parse()

	args := flag.Args()
	if len(args) != 2 {
		flag.Usage()
		os.Exit(0)
	}
	formerFile = args[0]
	latterFile = args[1]
}

func main() {

	fdat, err := ioutil.ReadFile(formerFile)
	check(err)
	ldat, err := ioutil.ReadFile(latterFile)
	check(err)
	mdat, err := merger.MergeYaml(fdat, ldat)
	check(err)
	err = ioutil.WriteFile(outputFile, mdat, 0644)
	check(err)
}

func check(e error) {
	if e != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", e)
		os.Exit(1)
	}
}
