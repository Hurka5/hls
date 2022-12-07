package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

/* Global Varibles  */
var isTerminal = true

/* Check for error  */
func check(err error) {
	if err != nil {
		panic(err)
	}
}

/* Check if runned in terminal and not as pipe */
func stdoutMode() bool {
	o, _ := os.Stdout.Stat()
	return (o.Mode() & os.ModeCharDevice) == os.ModeCharDevice
}

func main() {

	handleArguments()

	isTerminal = stdoutMode()

	dirs := []string{}
	files := []Item{}

	/* Group by folders and files */
	for _, path := range args.Paths {

		fileStat, err := os.Stat(path)

		if err != nil && strings.Contains(err.Error(), "no such file or directory") {
			fmt.Fprintln(os.Stderr, path + ": No such file or directory")
			continue
		} else {
			check(err)
		}

		if fileStat.IsDir() {
			dirs = append(dirs, path)
		} else {
			files = append(files, Item{info: fileStat, path: path})
		}
	}

	/* List files */
	if len(files) > 0 {
		files = filterItems(files)
		files = sortItems(files)

		listItems(files)

		if len(dirs) > 0 {
			fmt.Fprintln(os.Stdout, "")
		}
	}

	/* List dirs */
	if len(dirs) > 0 {
		for i, dir := range dirs {

			files, err := ioutil.ReadDir(dir)

			if err != nil && strings.Contains(err.Error(), "permission denied") {
				fmt.Fprintln(os.Stderr, dir + ": Permission denied")
				continue
			} else {
				check(err)
			}

			items := newItemArray(files, dir) // New Item Array
			items = filterItems(items)        // Filter
			items = sortItems(items)          // Sort

			if len(args.Paths) != 1 {
				printDir(dir)
			}

			listItems(items)

			if i != len(dirs)-1 && len(files) > 0 {
				fmt.Fprintln(os.Stdout, "")
			}

		}
	}

}
