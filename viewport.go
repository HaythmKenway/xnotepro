package main
import (
	"fmt"
	"os"
	"strings"

	
	"github.com/charmbracelet/lipgloss"
)

func (m Model) headerView() string {
	title := titleStyle.Render("logger")
	line := strings.Repeat("\033[34;1;1m─\033[0m", max(0, m.viewport.Width/2-5*lipgloss.Width(title)))
	return lipgloss.JoinHorizontal(lipgloss.Center, title, line)
}

func (m Model) footerView() string {
	info := infoStyle.Render(fmt.Sprintf("%3.f%%", 100-m.viewport.ScrollPercent()*100))
	line := strings.Repeat("\033[34;1;1m─\033[0m", max(0, m.viewport.Width/2-5*lipgloss.Width(info)))
	return lipgloss.JoinHorizontal(lipgloss.Center, line, info)
}

/*func  updateLogger() string {
    content, err := os.ReadFile("log4j.log")
    if err != nil {
        return ""
    }
    
    return string(content)
}
*/

func updateLogger() string {
	content, err := os.ReadFile(getWorkingDirectory()+"/log4j.log")
	if err != nil {
		return ""
	}

	lines := strings.Split(string(content), "\n")
	reversedContent := ""
	for i := len(lines) - 1; i >= 0; i-- {
		line := strings.TrimSpace(lines[i])
		if line != "" {
			reversedContent += line + "\n"
		}
	}

	return reversedContent
}
