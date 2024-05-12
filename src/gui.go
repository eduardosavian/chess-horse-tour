package main

import (
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/widget"
)

func gui() {
    // Create a new application
    myApp := app.New()

    // Create a new window
    myWindow := myApp.NewWindow("Simple GUI")

    // Create a widget
    myLabel := widget.NewLabel("Hello, Go GUI!")

    // Create a container to hold the widget
    myContainer := container.NewVBox(myLabel)

    // Set the window content to the container
    myWindow.SetContent(myContainer)

    // Show the window
    myWindow.ShowAndRun()
}
