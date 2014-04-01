package main

import (
  "fmt"
)

func main(){
  usage();
}

func usage() {

  option :=`Usage: %s [SHORT-OPTION]... [STRING]...
  or:  %s LONG-OPTION
  Echo the STRING(s) to standard output.

  -n  do not output the trailing newline
  -e  enable interpretation of backslash escapes (default)
  -E  disable interpretation of backslash escapes
  --help display this help and exit
  --version output version information and exit
`

  description :=`  If -e is in effect, the following sequences are recognized:
  \\ backslash
  a alert (BEL)
  b backspace
  c produce no further output
  e escape
  f form feed
  n new line
  r carriage return
  t horizontal tab
  v vertical tab
  0NNN   byte with octal value NNN (1 to 3 digits)
  xHH    byte with hexadecimal value HH (1 to 2 digits)
`

	fmt.Print(option+"\n"+description)
}

