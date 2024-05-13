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
    myApp := app.New()

    myWindow := myApp.NewWindow("Menu")

    input1 := widget.NewEntry()
    input2 := widget.NewEntry()
    input3 := widget.NewEntry()
    input4 := widget.NewEntry()

    form := &widget.Form{
        Items: []*widget.FormItem{
            {Text: "Input 1", Widget: input1},
            {Text: "Input 2", Widget: input2},
            {Text: "Input 3", Widget: input3},
            {Text: "Input 4", Widget: input4},
        },
    }

    submitButton := widget.NewButton("Submit", func() {
        startXStr := input1.Text
        startYStr := input2.Text
        concurrencyStr := input3.Text
        boardSizeStr := input4.Text

        startX, err := strconv.Atoi(startXStr)
        if err != nil {
            dialog.ShowError(errors.New("Invalid startX: " + err.Error()), myWindow)
            return
        }

        startY, err := strconv.Atoi(startYStr)
        if err != nil {
            dialog.ShowError(errors.New("Invalid startY: " + err.Error()), myWindow)
            return
        }

        concurrency, err := strconv.Atoi(concurrencyStr)
        if err != nil || concurrency <= 0 {
            dialog.ShowError(errors.New("Invalid number of threads. Must be a positive integer."), myWindow)
            return
        }

        boardSize, err := strconv.Atoi(boardSizeStr)
        if err != nil || boardSize <= 0 {
            dialog.ShowError(errors.New("Invalid board size. Must be a positive integer."), myWindow)
            return
        }

        fmt.Println("Validated Input:")
        fmt.Println("StartX:", startX)
        fmt.Println("StartY:", startY)
        fmt.Println("Concurrency:", concurrency)
        fmt.Println("BoardSize:", boardSize)
    })

    content := container.NewVBox(
        form,
        submitButton,
    )

    myWindow.SetContent(content)

    myWindow.ShowAndRun()
}
