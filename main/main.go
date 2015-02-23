package main

import (
	"flag"
	"strings"
	"backup/copy"
)

func main() {
	// Get command line arguments
	dirsPtr := flag.String("src", nil, "CSV list of dirs to copy from")
	ignorePtr := flag.String("ignore", nil, "CSV list of patterns to ignore. Default: .,..,bin")
	destPtr := flag.String("dest", nil, "Path to destination")

	if dirsPtr == nil || destPtr == nil {
		log.fatal("You must provide src and dest directories")
	}
	
	//parse csv
	dirs := strings.Split(&dirsPtr, ",")
	ignore := strings.Split(&ignorePtr, ",")

	// set up the channels
	discovered := make(chan *File, 1000)
	filtered := make(chan *File, 100)

	copy.Discover(dirs, discovered)
	copy.Filter(discovered, filtered, ignore)
	copy.Save(filtered, destPtr)




}