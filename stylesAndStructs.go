package main

import (
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/timer"
	"github.com/charmbracelet/bubbles/viewport"
	"github.com/charmbracelet/lipgloss"
)



type Model struct{
focused status
loaded bool
addNew bool
answer textinput.Model
tasks []list.Model
err error
Styles *Styles
timer timer.Model
viewport viewport.Model
quitting bool 
content string 
}

type Styles struct{
BorderColor lipgloss.Color
InputField lipgloss.Style }

func DefaultStyle() *Styles{
s:=new(Styles)
s.BorderColor=lipgloss.Color("36")
s.InputField=lipgloss.NewStyle().BorderForeground(s.BorderColor).BorderStyle(lipgloss.NormalBorder()).Padding(1).Width(80)
return s
}
//styling

var(
	titleStyle =func() lipgloss.Style{
		b:= lipgloss.RoundedBorder()
		b.Right="|-"
	return lipgloss.NewStyle().BorderStyle(b).Padding(0,1).BorderForeground(lipgloss.Color("69"))}()
	infoStyle = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		b.Left = "┤"
		return titleStyle.Copy().BorderStyle(b)
	}()
	columnStyling=lipgloss.NewStyle().Padding(1,2)
	activeButtonStyle=lipgloss.NewStyle().Foreground(lipgloss.Color("#000000")).Background(lipgloss.Color("#F25D94")).Padding(1,24).Border(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("69"))
	focusedStyling=lipgloss.NewStyle().Padding(1,2).Border(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("69"))
)


const timeout=time.Second*3
//constants
type status int 
const(
	running status=iota
	notes
	info
)

// model for the program
type TimeCell struct{
status status
title string
description string
}


