package main

import (
	"fmt"
  "flag"
	"os"
)

func main() {
	help_flag := flag.Bool("help", false, "help flag")
	version_flag := flag.Bool("version", false, "version flag")
	logical_flag := flag.Bool("logical", false , "logical flag")
	l_flag := flag.Bool("L", false , "logical flag")
	physical_flag := flag.Bool("physical", false, "physical flag")
	p_flag := flag.Bool("P", false, "physical flag")

	flag.Parse()
	
	if *help_flag {
		usage()
		os.Exit(0)
	}

	if *version_flag {
		version()
		os.Exit(0)
	}
	if *p_flag || *physical_flag {
		physical_path()
	} else if *l_flag || *logical_flag {
		logical_path()
	}

	physical_path()
}

func logical_path() {
	pwd := os.Getenv("PWD")
	fmt.Print(pwd + "\n")
	os.Exit(0)
	}
func physical_path() {
	pwd, err := os.Getwd()
	if err != nil {
		usage()
		os.Exit(1)
	}
	fmt.Print(pwd + "\n")
	os.Exit(0)
}

func usage() {
	description :=`Usage: pwd [OPTION]...
  Print the full filename of the current working directory.

	-L, --logical   use PWD from environment, even if it contains symlinks
	-P, --physical  avoid all symlinks
	--help     display this help and exit
	--version  output version information and exit

If no option is specified, -P is assumed.

NOTE: your shell may have its own version of pwd, which usually supersedes
the version described here.  Please refer to your shell's documentation
for details about the options it supports.`
	fmt.Print(description + "\n")
}

func version(){
	fmt.Print("pwd (go-coreutils) 0.01\n")
}
