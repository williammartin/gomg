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

	formattedTemplate := template.Must(template.New("Display Text").Parse(text + "\n"))
	formattedTemplate.Execute(ui.Out, keys)
}

func (ui *UI) DisplayError(err error) {
	fmt.Fprintln(ui.Err, err.Error())
}

func (ui *UI) DisplayErrorAndFailed(err error) {
	fmt.Fprintln(ui.Err, err.Error())
	ui.DisplayFailed()
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

func (ui *UI) DisplayStream(reader io.Reader) (int64, error) {
	return io.Copy(ui.Out, reader)
}
