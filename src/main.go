package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// Create a new application
	myApp := app.New()

	// Create a new window
	myWindow := myApp.NewWindow("Hello")

	// Create a label widget
	helloLabel := widget.NewLabel("Hello, Fyne!")

	// Create a button widget
	myButton := widget.NewButton("Click me!", func() {
		helloLabel.SetText("Button Clicked!")
	})

	// Create a container for the label and button
	content := container.NewVBox(
		helloLabel,
		myButton,
	)

	// Set the window content
	myWindow.SetContent(content)

	// Show the window
	myWindow.ShowAndRun()
}