package main
import(

	"github.com/charmbracelet/bubbles/textinput"
)
func New() *Model{
	styles:=DefaultStyle()
	answer:=textinput.New()
	answer.Focus()
	answer.Placeholder="log4j here!!!"
	return &Model{answer: answer,Styles:styles}}


