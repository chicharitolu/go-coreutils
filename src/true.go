package main

import (
	"fmt"
	"flag"
	"os"
)

func main() {
	help_flag := flag.Bool("help", false, "help flag")
	version_flag := flag.Bool("version", false, "version flag")
	
	flag.Parse()
	if flag.NFlag() > 1 {
		os.Exit(0)
	}

	if *help_flag {
		usage()
	}

	if *version_flag {
		version()
	}
	os.Exit(0)
}

func usage () {
	description :=`Usage: true [ignored command line arguments]
  or:  true OPTION
Exit with a status code indicating success.

      --help     display this help and exit
      --version  output version information and exit
`
	fmt.Print(description)
}

func version (){
	fmt.Print("true (go-coreutils) 0.01\n")
}
