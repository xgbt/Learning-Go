package main

import (
	"flag"
	"fmt"
	"os"
)

func cat(f *os.File) {
	const NBUF = 512
	var buf [NBUF]byte

	for {
		nr, err := f.Read(buf[:])
		switch {
		case nr < 0:
			fmt.Fprintf(os.Stderr, "cat: error reading: %s\n", err.Error())
			os.Exit(0)
		case nr == 0:
			return
		case nr > 0:
			if nw, ew := os.Stdout.Write(buf[0:nr]); nw != nr {
				fmt.Fprintf(os.Stderr, "cat: error writing: %s\b", ew.Error())
			}
		}
	}
}

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		cat(os.Stdin)
	}

	for i := 0; i < flag.NArg(); i++ {
		f, err := os.Open(flag.Arg(i))
		if err != nil {
			fmt.Fprintf(os.Stderr, "cat: can't open %s: error %s\n", flag.Arg(i), err)
			os.Exit(1)
		}
		cat(f)
		f.Close()
	}
}
