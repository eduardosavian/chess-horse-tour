package main

import (
    "errors"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/widget"
    "fyne.io/fyne/v2/dialog"
    "strconv"
    "fmt"
)

func menu() {
    // Create a new application instance
    myApp := app.New()

    // Create a new window
    myWindow := myApp.NewWindow("Menu")

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

    // Create a submit button
    submitButton := widget.NewButton("Submit", func() {
        // Retrieve input values
        startXStr := input1.Text
        startYStr := input2.Text
        concurrencyStr := input3.Text
        boardSizeStr := input4.Text

        // Validate input
        startX, err := strconv.Atoi(startXStr)
        if err != nil {
            // Display error message for startX
            dialog.ShowError(errors.New("Invalid startX: " + err.Error()), myWindow)
            return
        }

        startY, err := strconv.Atoi(startYStr)
        if err != nil {
            // Display error message for startY
            dialog.ShowError(errors.New("Invalid startY: " + err.Error()), myWindow)
            return
        }

        concurrency, err := strconv.Atoi(concurrencyStr)
        if err != nil || concurrency <= 0 {
            // Display error message for concurrency
            dialog.ShowError(errors.New("Invalid number of threads. Must be a positive integer."), myWindow)
            return
        }

        boardSize, err := strconv.Atoi(boardSizeStr)
        if err != nil || boardSize <= 0 {
            // Display error message for boardSize
            dialog.ShowError(errors.New("Invalid board size. Must be a positive integer."), myWindow)
            return
        }

        // If all inputs are valid, proceed with the validated values
        fmt.Println("Validated Input:")
        fmt.Println("StartX:", startX)
        fmt.Println("StartY:", startY)
        fmt.Println("Concurrency:", concurrency)
        fmt.Println("BoardSize:", boardSize)

        // Now you can proceed with using these validated values for your application logic
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
