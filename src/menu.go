package main

import (
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/widget"
)

func gui() {
    // Create a new application instance
    myApp := app.New()

    // Create a new window
    myWindow := myApp.NewWindow("Input GUI")

    // Create input widgets
    input1 := widget.NewEntry()
    input2 := widget.NewEntry()
    input3 := widget.NewEntry()
    input4 := widget.NewEntry()

    // Create a form container to hold the input widgets
    form := &widget.Form{
        Items: []*widget.FormItem{
            {Text: "Input 1", Widget: input1},
            {Text: "Input 2", Widget: input2},
            {Text: "Input 3", Widget: input3},
            {Text: "Input 4", Widget: input4},
        },
    }

    // Create a submit button (optional)
    submitButton := widget.NewButton("Submit", func() {
        // Handle form submission here
        // You can access input values using input1.Text(), input2.Text(), etc.
        // Example:
        // fmt.Println("Input 1:", input1.Text())
        // fmt.Println("Input 2:", input2.Text())
        // fmt.Println("Input 3:", input3.Text())
        // fmt.Println("Input 4:", input4.Text())
    })

    // Create a container for the form and button
    content := container.NewVBox(
        form,
        submitButton,
    )

    // Set the window content
    myWindow.SetContent(content)

    // Show the window and run the application
    myWindow.ShowAndRun()
}
