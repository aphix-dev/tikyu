package main

import (
	"io"
	"os"
)

func main() {
	println("tikyu engine")

	if os.Args[1] == "compile" {
		fn := os.Args[2]
		file, err := os.Open(fn)
		if err != nil {
			panic(err)
		}

		fc := make([]byte, 0)
		for {
			b := make([]byte, 1024)
			_, err = file.Read(b)
			fc = append(fc, b...)
			if err == io.EOF {
				println("file read complete")
				break
			} else if err != nil {
				panic(err)
			}
		}

		println(string(fc))
	}
}
