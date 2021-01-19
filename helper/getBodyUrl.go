package helper

import (
  "net/http"
  "io/ioutil"
  "strings"
  "fmt"
  "github.com/PuerkitoBio/goquery"
)

func GetBodyUrl(get_url string) ([]string, error) {
  //Get request
  resp, err := http.Get(get_url)

  //check for an error with the request
  if err != nil {
    fmt.Println("Error: ", err)
    return nil, err
  }

  //Close response request
  defer resp.Body.Close()

  //Read the response body
  body, err := ioutil.ReadAll(resp.Body)

  return string(body)
}
