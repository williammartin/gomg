package ui

import (
	"fmt"
	"io"
	"text/template"

	"github.com/fatih/color"
)

type UI struct {
	Out io.Writer
	Err io.Writer
}

func (ui *UI) DisplayText(text string, data ...map[string]interface{}) {
	var keys interface{}
	if len(data) > 0 {
		keys = data[0]
	}

	formattedTemplate := template.Must(template.New("Display Text").Parse(text))
	formattedTemplate.Execute(ui.Out, keys)
}

func (ui *UI) DisplayError(text string) {
	fmt.Fprintln(ui.Err, text)
}

func (ui *UI) DisplayFailed() {
	style := color.New(color.FgRed, color.Bold)
	style.Fprintln(ui.Err, "FAILED")
}

func (ui *UI) DisplaySuccess() {
	style := color.New(color.FgGreen, color.Bold)
	style.Fprintln(ui.Out, "SUCCESS")
}

func (ui *UI) DisplayNewline() {
	fmt.Fprintln(ui.Out)
}
