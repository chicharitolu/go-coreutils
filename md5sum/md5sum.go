package main

import (
  "fmt"
	"flag"
)

func main(){
  help_flag := flag.Bool("help", false, "help flag")
	version_flag := flag.Bool("version", false, "version")
	flag.Parse()
	
	if *help_flag {
		usage()
	}
	if *version_flag {
		version()
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
