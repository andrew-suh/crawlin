package main
import (
  "os"
  "fmt"
  "crawlin/helper"
)

func main() {

  if len(os.Args) < 2 {
    fmt.Println("Input a url")
    os.Exit(1)
  }
  get_url := os.Args[1]

  stringed_body := helper.GetBodyUrl(get_url)

  //Use custom func to write to filename
  end := helper.CreateWrite("dump.txt",stringed_body)

  fmt.Println(end)
}
