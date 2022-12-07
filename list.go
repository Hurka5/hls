package main

import (
	"fmt"
	"math"
	"os"
	"time"

	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

const GAP = 2 + 2 // Gap between items + left border

func printDir(dir string) {
	fmt.Fprintln(os.Stdout, "╭"+dir)
}

func listItems(items []Item) {

	if len(items) == 0 {
		return
	}

	if args.Long {
		detailedList(items)
	} else if !isTerminal || args.Oneline {
		oneList(items)
	} else {
		simpleList(items)
	}

}

/* Listing in one column */
func oneList(items []Item) {
	for _, item := range items {
		if isTerminal {
			fmt.Fprintln(os.Stdout, styleItem(item))
		} else {
			fmt.Fprintln(os.Stdout, item.info.Name())
		}
	}
}

// Simple Listing
func simpleList(items []Item) {

	// Get term width
	width, _, err := term.GetSize(0)
	check(err)

	// Search for file with longest name
	longest := 0
	for _, item := range items {
		nameLen := len(item.info.Name())
		if nameLen > longest {
			longest = nameLen
		}
	}

	longest += GAP

	// Calculate cols and rows
	var cols int = int(math.Round(float64(width) / float64(longest)))
	if cols == 0 {
		cols = 1
	}
	var rows int = int(math.Ceil(float64(len(items)) / float64(cols)))

	/*
	  Gather styled item names to array
	*/
	names := make([][]string, cols)
	index := 0
	for _, item := range items {
		names[index] = append(names[index], styleItem(item))
		if len(names[index]) == rows {
			index++
		}
	}

	/*
	  Assembles table with lipgloss
	  1. Creates cols from item names
	  2. Puts cols together horiontally
	*/
	var itemCols []string
	for _, name := range names {
		itemCols = append(itemCols, lipgloss.JoinVertical(lipgloss.Left, name...))
	}
	table := lipgloss.JoinHorizontal(lipgloss.Top, itemCols...)

	// Print table
	fmt.Fprint(os.Stdout, table+"\n")
}

/* Detailed Listing */
func detailedList(items []Item) {
	//TODO: Oprional Columns
	// Permissions | Size | User | Group | Date | Name

	datas := make([][]string, 6)
	for _, item := range items {

		// Add Permissions
		datas[0] = append(datas[0], permStyle.Render(item.Perm()))

		// Add Size
		if item.info.IsDir() {
			datas[1] = append(datas[1], sizeStyle.Render("-"))
		} else {
			datas[1] = append(datas[1], sizeStyle.Render(item.Size()))
		}

		// Add Owners
		user, group := item.Owner()
		datas[2] = append(datas[2], userStyle.Render(user))
		datas[3] = append(datas[3], groupStyle.Render(group))

		// Add Last Modification Date
		date := item.info.ModTime()
		if date.Year() < time.Now().Year() {
			datas[4] = append(datas[4], dateStyle.Render(date.Format("02 Jan 2006")))
		} else {
			datas[4] = append(datas[4], dateStyle.Render(date.Format("02 Jan 15:04")))
		}

		// Add File Name
		datas[5] = append(datas[5], styleItem(item))

	}

	// Assemble Table
	var cols []string
	for _, data := range datas {
		cols = append(cols, lipgloss.JoinVertical(lipgloss.Left, data...))
	}
	table := lipgloss.JoinHorizontal(lipgloss.Top, cols...)
	fmt.Fprint(os.Stdout, table+"\n")

}
