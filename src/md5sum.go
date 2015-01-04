package main

import (
	"crypto/md5"
  "fmt"
	"flag"
	"io/ioutil"
	"os"
)

func main(){
  help_flag := flag.Bool("help", false, "help flag")
	version_flag := flag.Bool("version", false, "version")
	check_flag := flag.Bool("check", false, "check")
//	quiet_flag := flag.Bool("quiet", false, "quiet")
//	warn_flag := flag.Bool("warn", false, "warn")
//	status_flag := flag.Bool("status", false, "status")
	
	flag.Parse()
	
	if *help_flag {
		usage()
		os.Exit(0)
	}
	if *version_flag {
		version()
		os.Exit(0)
	}

	if !*check_flag {
		checksum()
	}
	
}

func checksum() {
	if flag.NArg() > 0 {
		for _, file := range flag.Args() {
			buf, err := ioutil.ReadFile(file)
			if err != nil {
				fmt.Printf("md5sum: cannot read '%s': %s\n", file, err)
			}
			fmt.Printf("%x %s\n", md5.Sum(buf), file)
		}
	} else {
		buf, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			fmt.Printf("md5sum: cannot read 'STDIN': %s\n", err)
		}
		fmt.Printf("%x -\n", md5.Sum(buf))
	}
}

func usage() {
  description :=`Usage: md5sum [OPTION]... [FILE]...
  Print or check MD5 (128-bit) checksums.
  With no FILE, or when FILE is -, read standard input.

  -b, --binary            read in binary mode
  -c, --check             read MD5 sums from the FILEs and check them
  -t, --text              read in text mode (default)
  Note: There is no difference between binary and text mode option on GNU system.

The following three options are useful only when verifying checksums:
     --quiet             don't print OK for each successfully verified file
     --status            don't output anything, status code shows success
 -w, --warn              warn about improperly formatted checksum lines

     --help     display this help and exit
     --version  output version information and exit

The sums are computed as described in RFC 1321.  When checking, the input
should be a former output of this program.  The default mode is to print
a line with checksum, a character indicating type (*' for binary,  ' for
text), and name for each FILE
`
  fmt.Print(description)
  }

func version() {
	fmt.Print("md5sum (go-coreutils) 0.01\n")
}
