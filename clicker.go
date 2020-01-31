package main


import (
	"github.com/go-vgo/robotgo"
	"math/rand"
	"time"
	"fmt"
)

func main() {
  //curX, curY := robotgo.GetMousePos()
  //fmt.Println("pos:", x, y)
  for true {
    sleepTime := time.Duration(rand.Intn(1000000000) + 3000000000)
    robotgo.MouseClick("left", true)
    fmt.Print("Sleeping for: ")
    fmt.Println(sleepTime)
    time.Sleep(sleepTime)
  }
}