package main
import (
  "net/http"
  "io/ioutil"
  "fmt"
  "crawlin/file_helpers"
)

func main() {
  //Get request
  resp, err := http.Get("https://golang.org/pkg/fmt/")

  //check for an error
  if err != nil {
    panic(err)
  }

  //close response body when finished request
  defer resp.Body.Close()

  //Read the Body content from the request
  body, err := ioutil.ReadAll(resp.Body)

  //delete line later
  fmt.Printf("type: %T\n",body)

  //Use custom func to write to filename
  end := file_helpers.CreateWrite("dump.txt",string(body))

  fmt.Println(end)
}
