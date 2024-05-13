package main

import (
    "strconv"
    "fmt"

    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/dialog"
    "fyne.io/fyne/v2/widget"
)

var myApp fyne.App

func main() {
    myApp = app.New()

    myWindow := myApp.NewWindow("Knight's Tour")
    myWindow.Resize(fyne.NewSize(400, 300))

    input1 := widget.NewEntry()
    input1.SetPlaceHolder("Start X")
    input2 := widget.NewEntry()
    input2.SetPlaceHolder("Start Y")
    input3 := widget.NewEntry()
    input3.SetPlaceHolder("Board Size")

    submitButton := widget.NewButton("Submit", func() {
        startXStr := input1.Text
        startYStr := input2.Text
        boardSizeStr := input3.Text

        startX, err := strconv.Atoi(startXStr)
        if err != nil {
            dialog.ShowError(fmt.Errorf("Invalid startX: %s", err.Error()), myWindow)
            return
        }

        startY, err := strconv.Atoi(startYStr)
        if err != nil {
            dialog.ShowError(fmt.Errorf("Invalid startY: %s", err.Error()), myWindow)
            return
        }

        boardSize, err := strconv.Atoi(boardSizeStr)
        if err != nil || boardSize <= 0 {
            dialog.ShowError(fmt.Errorf("Invalid board size. Must be a positive integer."), myWindow)
            return
        }

        handleSubmission(startX, startY, boardSize, myWindow)
    })

    inputs := container.NewVBox(
        input1,
        input2,
        input3,
        submitButton,
    )

    myWindow.SetContent(inputs)

    myWindow.ShowAndRun()
}

func handleSubmission(startX, startY, boardSize int, myWindow fyne.Window) {
    board := make([][]int, boardSize)
    for i := range board {
        board[i] = make([]int, boardSize)
    }

    method := "warnsdorff"

    if !backtrack(board, 1, startX, startY, boardSize, method) {
        dialog.ShowInformation("Knight's Tour", "No valid Knight's Tour found after exhausting all attempts.", myWindow)
        return
    }

    boardTable := widget.NewTable(
        func() (int, int) {
            return boardSize, boardSize
        },

        func() fyne.CanvasObject {
            return widget.NewLabel("")
        },

        func(i widget.TableCellID, cell fyne.CanvasObject) {
            row := i.Row
            col := i.Col
            label := cell.(*widget.Label)
            label.SetText(strconv.Itoa(board[row][col]))
        },
    )


    tableContainer := container.NewScroll(boardTable)

    myWindow.SetContent(tableContainer)
}
