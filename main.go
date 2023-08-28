package main

import (
	"fmt"
	"log"
	"os"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/timer"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)
const useHighPerformanceRenderer = false


//logger function 
func logToFile(input string) error {
	fileName := "log4j.log"

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

//functions for structures
func (t TimeCell) FilterValue() string{
	return t.title}

func (t TimeCell) Title() string{
	return t.title 
}
func (t TimeCell) Description() string{
	return t.description
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
			m.viewport.YPosition = headerHeight + 1
			m.viewport.SetContent(m.content)
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
				m.loaded=false
				m.content=m.updateLogger()
				return m,nil
			}
	}
	var cmd tea.Cmd
	m.tasks[m.focused],cmd=m.tasks[m.focused].Update(msg)
	m.answer,cmd=m.answer.Update(msg)
	return m,cmd
}

func main(){
	content, err := os.ReadFile("log4j.log")
	if err != nil {
		fmt.Println("could not load file:", err)
		os.Exit(1)
	}
	m:=New()
	m.quitting=false
	m.content=string(content)
	m.updateLogger()
	p:= tea.NewProgram(m,tea.WithAltScreen(),tea.WithMouseCellMotion(),)
	if _,err:=p.Run();err!=nil{
	os.Exit(1)}
}
