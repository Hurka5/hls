package main

import (
	"github.com/charmbracelet/lipgloss"
)

/*
|-------------------------------|
|             ICONS             |
|-------------------------------|
*/
var typeIcons = map[byte]string{
	'-': "\uebb5 ", //regular file
	'd': "\uf413 ", //directory
	'c': "CC",      //character device file
	'b': "BB",      //block device file
	's': "SS",      //local socket file
	'p': "PP",      //named pipe
	'L': "\uf0c1 ", //symbolic link
	'l': "\uf127 ", //broken symbolic link
}
var permIcons = map[byte]string{
	'-': "\uebb5 ", //no permission
	'r': "\uf707 ", //read
	'w': "\uf040 ", //write
	'x': "\uf423 ", //execute
	's': "SS",      //uses group privilage
	't': "TT",      //can only be deleted by owner
}

/*
|-------------------------------|
|             STYLES            |
|-------------------------------|
*/
var (
	// Pipe Style
	defaultStyle = lipgloss.NewStyle().MarginRight(2)
	detailsStyle = lipgloss.NewStyle().MarginRight(1)
	// Default Terminal Style
	normalStyle = defaultStyle.Copy().
			Border(lipgloss.RoundedBorder() /*.ThickBorder()*/, false, false, false, true)
	//Directory Style
	dirStyle = normalStyle.Copy().
			Foreground(lipgloss.Color("#0078ff")).
			BorderForeground(lipgloss.Color("#0078ff")).
			Bold(true)
	//Executable Style
	execStyle = normalStyle.Copy().
			Foreground(lipgloss.Color("#5aeb7e")).
			BorderForeground(lipgloss.Color("#5aeb7e")).
			Bold(true)
	//Link Style
	linkStyle = normalStyle.Copy().
			Foreground(lipgloss.Color("#00ddff")).
			BorderForeground(lipgloss.Color("#00ddff")).
			Bold(true)
	//Broken Link Style
	blinkStyle = normalStyle.Copy().
			Foreground(lipgloss.Color("#aa0000")).
			BorderForeground(lipgloss.Color("#aa0000")).
			Bold(true)
	//Permission Style
	permStyle = detailsStyle.Copy()
	//Size style
	sizeStyle = detailsStyle.Copy().
			Foreground(lipgloss.Color("#5aeb7e"))
	//User style
	userStyle = detailsStyle.Copy().
			Foreground(lipgloss.Color("#ff64b8"))
	//Group style
	groupStyle = detailsStyle.Copy().
			Foreground(lipgloss.Color("#ec3aaa"))
	//Date style
	dateStyle = detailsStyle.Copy().
			Foreground(lipgloss.Color("#ffff66"))
)

func styleItem(item Item) string {

	if !isTerminal {
		return defaultStyle.Render(item.info.Name())
	}

	// Check if dir
	if item.info.IsDir() {
		return dirStyle.Render(item.info.Name())
	}

	// Check if link
	isLink, isBroken := item.isLink()
	if isLink {
		if isBroken {
			return blinkStyle.Render(item.info.Name())
		} else {
			return linkStyle.Render(item.info.Name())
		}
	}

	// Check if executable
	if item.isExecutable() {
		return execStyle.Render(item.info.Name())
	}

	return normalStyle.Render(item.info.Name())
}
