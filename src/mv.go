package main

import (
	"fmt"
	"flag"
	"os"
)

func main() {
	help_flag := flag.Bool("help", false, "help flag")
	version_flag := flag.Bool("version", false, "version")
	flag.Parse()

	if *help_flag {
		usage()
		os.Exit(0)
	}
	if *version_flag {
		os.Exit(0)
	}
		

}

func usage() {
	description :=`Usage: mv [OPTION]... [-T] SOURCE DEST
	or:  mv [OPTION]... SOURCE... DIRECTORY
	or:  mv [OPTION]... -t DIRECTORY SOURCE...
		Rename SOURCE to DEST, or move SOURCE(s) to DIRECTORY.

		Mandatory arguments to long options are mandatory for short options too.
		--backup[=CONTROL]       make a backup of each existing destination file
	-b                           like --backup but does not accept an argument
	-f, --force                  do not prompt before overwriting
	-i, --interactive            prompt before overwrite
	-n, --no-clobber             do not overwrite an existing file
	If you specify more than one of -i, -f, -n, only the final one takes effect.
		--strip-trailing-slashes  remove any trailing slashes from each SOURCE
	argument
	-S, --suffix=SUFFIX          override the usual backup suffix
	-t, --target-directory=DIRECTORY  move all SOURCE arguments into DIRECTORY
	-T, --no-target-directory    treat DEST as a normal file
	-u, --update                 move only when the SOURCE file is newer
	than the destination file or when the
	destination file is missing
	-v, --verbose                explain what is being done
	--help     display this help and exit
	--version  output version information and exit

	The backup suffix is '~', unless set with --suffix or SIMPLE_BACKUP_SUFFIX.
The version control method may be selected via the --backup option or through
the VERSION_CONTROL environment variable.  Here are the values:

  none, off       never make backups (even if --backup is given)
  numbered, t     make numbered backups
  existing, nil   numbered if numbered backups exist, simple otherwise
  simple, never   always make simple backups
`
  fmt.Print(description)
}
