package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/timer"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)
const useHighPerformanceRenderer = false


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
		return lipgloss.JoinHorizontal(lipgloss.Left,lipgloss.JoinVertical(lipgloss.Center,focusedStyling.Render(timerView),okButton,),lipgloss.JoinVertical(lipgloss.Left,m.headerView(), m.viewport.View(), m.footerView(),m.Styles.InputField.Render(m.answer.View())),)//columnStyling.Render(timerView),columnStyling.Render(notesView))
}}
return "Loading..."}

//update function to update values
func (m Model) Update(msg tea.Msg) (tea.Model,tea.Cmd){
		var (cmd tea.Cmd
			 cmds []tea.Cmd)
		switch msg :=msg.(type){
		case timer.TickMsg:
		m.timer,cmd=m.timer.Update(msg)
		return m,cmd
		case tea.WindowSizeMsg:
			headerHeight :=lipgloss.Height(m.headerView())
			footerHeight := lipgloss.Height(m.footerView())
			verticalMarginHeight := headerHeight + footerHeight
		if !m.loaded{
			//for loading viewport
			m.viewport = viewport.New(msg.Width, msg.Height-verticalMarginHeight-6)
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
			} else {
			m.viewport.Width = msg.Width
			m.viewport.Height = msg.Height - verticalMarginHeight
		}
		if useHighPerformanceRenderer {
			cmds = append(cmds, viewport.Sync(m.viewport))
		}
		case tea.KeyMsg:
			switch msg.String(){
				case "ctrl+c","q": m.quitting=true
				return m,tea.Quit
			case "enter":
				ans:=m.answer.Value()
				logToFile(ans)
				m.content=updateLogger()
				m.viewport.SetContent(m.content)
				m.answer.SetValue("")
				return m,nil
			}
	}
	m.viewport, cmd = m.viewport.Update(msg)
	cmds = append(cmds, cmd)
	m.tasks[m.focused],cmd=m.tasks[m.focused].Update(msg)
	m.answer,cmd=m.answer.Update(msg)
	cmds=append(cmds, cmd)
	return m,tea.Batch(cmds...)
}

func main(){
	WORKDIR:=getWorkingDirectory()
	content, err := os.ReadFile(WORKDIR+"/log4j.log")
	if err != nil {
		fmt.Println("could not load file:", err)
		os.Exit(1)
	}
	m:=New()
	m.quitting=false
	m.content=string(content)
	m.content=updateLogger()
	p:= tea.NewProgram(m,tea.WithAltScreen(),tea.WithMouseCellMotion(),)
	if _,err:=p.Run();err!=nil{
	os.Exit(1)}
}
