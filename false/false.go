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
		os.Exit(1)
	}

	if *help_flag {
		usage()
	}

	if *version_flag {
		version()
	}
	os.Exit(1)
}

func usage () {
	description :=`Usage: false [ignored command line arguments]
  or:  false OPTION
Exit with a status code indicating failure.

      --help     display this help and exit
      --version  output version information and exit
`
	fmt.Print(description)
}

func version (){
	fmt.Print("false (go-coreutils) 0.01\n")
}
