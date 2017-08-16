package main

import (
  "flag"
  "fmt"
  "net/http"
  "time"
  "io/ioutil"
  "encoding/json"
)

type character struct {
  Died string `json:"died"`
}

func main() {
  firstName := flag.String("firstName", "", "first name of character")
  lastName := flag.String("lastName", "", "last name of character")
  flag.Parse()
  url := fmt.Sprintf("https://www.anapioficeandfire.com/api/characters/?name=%s%%20%s", *firstName, *lastName)
  client := http.Client {
    Timeout: time.Second * 5,
  }

  req, reqErr := http.NewRequest(http.MethodGet, url, nil)
  if reqErr != nil {
    fmt.Println(reqErr)
    return
  }

  res, resErr := client.Do(req)
  if resErr != nil {
    fmt.Println(resErr)
    return
  }

  read, readErr := ioutil.ReadAll(res.Body)
  if resErr != nil {
    fmt.Println(readErr)
    return
  }

  var chars []character
  jsonErr := json.Unmarshal(read, &chars)
  if jsonErr != nil {
    fmt.Println(jsonErr)
  } else if len(chars) > 0{
    char := chars[0]
    died := char.Died
    if died == "" {
      fmt.Printf("%s %s is alive!\n", *firstName, *lastName)
    } else {
      fmt.Printf("%s %s died %s\n", *firstName, *lastName, died)
    }
  }
}
