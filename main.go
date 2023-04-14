package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		os.Exit(1)
	}
}

func run() (err error) {
	if len(os.Args) != 2 {
		return fmt.Errorf("invalid parameter length. this command requires only one parameters")
	}
	filename := os.Args[1]
	dir := filepath.Dir(filename)
	if _, err = os.Stat(dir); os.IsNotExist(err) {
		if err = os.MkdirAll(dir, os.ModePerm); err != nil {
			return
		}
	}
	_, err = os.Stat(filename)
	if os.IsNotExist(err) {
		file, _err := os.Create(filename)
		if _err != nil {
			err = _err
			return
		}

		defer func(file *os.File) {
			err = file.Close()
			if err != nil {
				return
			}
		}(file)
	} else {
		currenTime := time.Now().Local()
		if err = os.Chtimes(filename, currenTime, currenTime); err != nil {
			return
		}
	}
	return
}
