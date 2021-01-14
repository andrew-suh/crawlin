//Package file_writer simplifies file creation and writing process
package file_helpers
import (
  "os"
  "fmt"
)

//CreatWrite creates a file with an imported name 
//and dumps the string body into created file

func CreateWrite(filename,body string) string {
  //Create file
  f, err:= os.Create(filename)
  //Stop if file cannot be made
  if err != nil {
    fmt.Println("Error: ", err)
    os.Exit(1)
  }

  //Create file descriptor
  w, err := f.WriteString(body)
  //Stop if write fails
  if err != nil {
    fmt.Println("Writing failed")
    fmt.Println("Error: ", err)
    //Close fd
    f.Close()
    os.Exit(1)
  }

  fmt.Println(w, "bytes written successfully")
  //Close open file descriptor
  err = f.Close()
  if err != nil {
    fmt.Println("Closing created file fd failed")
    fmt.Println("Error: ", err)
    os.Exit(1)
  }
  return "Successfully finished writing to: " + filename
}
