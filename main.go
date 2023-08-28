package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/timer"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)
const useHighPerformanceRenderer = false

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
//logger function 
func logToFile(input string) error {
	fileName := "dummy.md"

	// Open the file in append mode or create it if it doesn't exist
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Set log output to the file
	log.SetOutput(file)

	// Log the input
	log.Println(input)

	return nil
}

// model for the program
type TimeCell struct{
status status
title string
description string
}

//functions for structures
func (t TimeCell) FilterValue() string{
	return t.title}

func (t TimeCell) Title() string{
	return t.title 
}
func (t TimeCell) Description() string{
	return t.description
}


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


//initializing tea init
func (m Model) Init() tea.Cmd{return nil}

//defining tasks 
func (m *Model) initLists(width,height int){
	defaultList :=list.New([]list.Item{},list.NewDefaultDelegate(),width/4,height/2)
	m.tasks=[]list.Model{defaultList}
	m.tasks[running].Title="Tasks"
	m.tasks[running].SetItems([]list.Item{
	TimeCell{status: running,title: "Task1",description: "lame things" },
	TimeCell{status: running,title: "Task2",description: "nothing" },
	})


}

//function to return view 
func (m Model) View() string{
	if m.quitting{
	return "nothing here"}
	if m.loaded{
		timerView:=m.tasks[running].View()
		okButton:=activeButtonStyle.Render("Add New")
		switch m.focused{
		case notes:
		return lipgloss.JoinHorizontal(lipgloss.Left,focusedStyling.Render(timerView),)//columnStyling.Render(timerView),columnStyling.Render(notesView))
		default:
		return lipgloss.JoinHorizontal(lipgloss.Left,lipgloss.JoinVertical(lipgloss.Center,focusedStyling.Render(timerView),okButton,),lipgloss.JoinVertical(lipgloss.Center,m.headerView(), m.viewport.View(), m.footerView(),m.Styles.InputField.Render(m.answer.View())),)//columnStyling.Render(timerView),columnStyling.Render(notesView))
}}
return "Loading..."}

//update function to update values
func (m Model) Update(msg tea.Msg) (tea.Model,tea.Cmd){
	switch msg :=msg.(type){
	case timer.TickMsg:
		var cmd tea.Cmd
		m.timer,cmd=m.timer.Update(msg)
		return m,cmd
		case tea.WindowSizeMsg:
			headerHeight :=lipgloss.Height(m.headerView())
			footerHeight := lipgloss.Height(m.footerView())
			verticalMarginHeight := headerHeight + footerHeight
		if !m.loaded{
			//for loading viewport
			m.viewport = viewport.New(msg.Width, msg.Height-verticalMarginHeight-25)
			m.viewport.YPosition = headerHeight
			m.viewport.HighPerformanceRendering = useHighPerformanceRenderer
			m.viewport.SetContent(m.content)
			m.viewport.YPosition = headerHeight + 1
			//for viewing list 
			m.initLists(msg.Width,msg.Height)
			columnStyling.Width(msg.Width/3)
			focusedStyling.Width(msg.Width/3)
			columnStyling.Height(msg.Height-8)
			focusedStyling.Height(msg.Height-8)
			m.loaded=true
			}
		case tea.KeyMsg:
			switch msg.String(){
				case "ctrl+c","q": m.quitting=true
				return m,tea.Quit
			case "enter":
				ans:=m.answer.Value()
				m.answer.SetValue("")
				logToFile(ans)
				m.content=m.updateLogger()
				return m,nil
			}
	}
	var cmd tea.Cmd
	m.tasks[m.focused],cmd=m.tasks[m.focused].Update(msg)
	m.answer,cmd=m.answer.Update(msg)
	return m,cmd
}

//creating a instance of model 
func New() *Model{
	styles:=DefaultStyle()
	answer:=textinput.New()
	answer.Focus()
	answer.Placeholder="log4j here!!!"
	return &Model{answer: answer,Styles:styles}}

//for viewport
func (m Model) headerView() string {
	title := titleStyle.Render("logger")
	line := strings.Repeat("\033[34;1;1m─\033[0m", max(0, m.viewport.Width/2-5*lipgloss.Width(title)))
	return lipgloss.JoinHorizontal(lipgloss.Center, title, line)
}

func (m Model) footerView() string {
	info := infoStyle.Render(fmt.Sprintf("%3.f%%", m.viewport.ScrollPercent()*100))
	line := strings.Repeat("\033[34;1;1m─\033[0m", max(0, m.viewport.Width/2-5*lipgloss.Width(info)))
	return lipgloss.JoinHorizontal(lipgloss.Center, line, info)
}

func (m Model) updateLogger() string{
	content, err := os.ReadFile("dummy.md")
	if err != nil {
		fmt.Println("could not load file:", err)
		os.Exit(1)
		time.Sleep(8 * time.Second)
		fmt.Print(string(content))
	return string(content)
	}
	return ""
}
func main(){
	m:=New()
	m.quitting=false
	m.updateLogger()
	p:= tea.NewProgram(m,tea.WithAltScreen())
	if _,err:=p.Run();err!=nil{
	os.Exit(1)}
}
