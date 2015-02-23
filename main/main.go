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
	discovered := make(chan FileData, 1000)

	copy.InitFilter(ignore)
	copy.Discover(dirs, discovered)
	copy.Save(discovered, destPtr)




}