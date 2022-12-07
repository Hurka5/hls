package main

import (
	"fmt"
	"os"
)

type Item struct {
	info os.FileInfo
	path string
}

func (item Item) Perm() string {
	perm := ""
	mode := item.info.Mode()

	// File type
	filetype := mode.String()[0]

	if filetype == 'L' {

		if _, isBroken := item.isLink(); isBroken {
			perm += typeIcons['l']
		} else {
			perm += typeIcons[filetype]
		}

	} else {
		perm += typeIcons[filetype]
	}

	// File permissions
	for i := 1; i < 10; i++ {
		perm += permIcons[mode.String()[i]]
	}

	return perm
}

func (item Item) Size() string {
	size := item.info.Size()
	const unit = 1000
	if size < unit {
		return fmt.Sprintf("%d", size)
	}
	div, exp := int64(unit), 0
	for n := size / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f%c", float64(size)/float64(div), "kMGTPE"[exp])
}

// Convert
func newItemArray(files []os.FileInfo, p string) []Item {
	var items []Item
	for _, file := range files {
		var fitem = Item{info: file, path: p + "/" + file.Name()}
		items = append(items, fitem)
	}
	return items
}
