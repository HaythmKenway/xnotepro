package main
import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
)

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
//there is some problem over here
func (m Model) updateLogger() string{
	content, err := os.ReadFile("log4j.log")
	if err != nil {
		fmt.Println("could not load file:", err)
		os.Exit(1)
		time.Sleep(8 * time.Second)
		fmt.Print(string(content))
	return string(content)
	}
	return ""
}

