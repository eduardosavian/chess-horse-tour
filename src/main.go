package main

import (
	"fmt"
	"os"
  "strconv"
)

func main() {
  fmt.Println("Hello, world")

  if length := len(os.Args);  length < 3 || length > 3 {
    panic("Format invalid, use go run x y")
  }

  x, err := strconv.Atoi(os.Args[1])
  if err != nil {
      panic(err)
  }

  y, err := strconv.Atoi(os.Args[2])
  if err != nil {
      panic(err)
  }

  var chess [8][8]int;
  chess[x][y] = 1

  print(chess[x][y])
}