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
  //headerPrefixes: "┌","╭","╭","▛","▗","┏","╔"," " - You can add more
  headerPrefix = "╭"
	// Pipe Style
	defaultStyle = lipgloss.NewStyle().MarginRight(2)
	// Default Terminal Style
	normalStyle = defaultStyle.Copy().
		  Border(lipgloss.RoundedBorder(), false, false, false, true).
      Foreground(lipgloss.AdaptiveColor{Light: "0",Dark: "7"})
	// Bold Sample Style
  boldStyle = normalStyle.Copy().Bold(true)
	//Directory Style
	dirStyle = boldStyle.Copy().
			Foreground(lipgloss.Color("4")).
			BorderForeground(lipgloss.Color("4"))
	//Executable Style
	execStyle = boldStyle.Copy().
			Foreground(lipgloss.Color("2")).
			BorderForeground(lipgloss.Color("2"))
	//Link Style
	linkStyle = boldStyle.Copy().
			Foreground(lipgloss.Color("6")).
			BorderForeground(lipgloss.Color("6"))
	//Broken Link Style
	blinkStyle = boldStyle.Copy().
			Foreground(lipgloss.Color("1")).
			BorderForeground(lipgloss.Color("1"))
	
  // Long Format Details Style Sample
	detailsStyle = lipgloss.NewStyle().MarginRight(1)
  //Permission Style
	permStyle = detailsStyle.Copy()
	//Size style
	sizeStyle = detailsStyle.Copy().
			Foreground(lipgloss.Color("2"))
	//User style
	userStyle = detailsStyle.Copy().
			Foreground(lipgloss.Color("5"))
	//Group style
	groupStyle = detailsStyle.Copy().
			Foreground(lipgloss.Color("5"))
	//Date style
	dateStyle = detailsStyle.Copy().
			Foreground(lipgloss.Color("3"))


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
