paramstostruct
===
[![GoDoc](https://godoc.org/github.com/Iwark/paramstostruct?status.svg)](https://godoc.org/github.com/Iwark/paramstostruct)

Package ``paramstostruct`` convert a map to a struct.

## Installation

```
$ go get github.com/Iwark/paramstostruct
```

## Example

```go
package main

import (
  "fmt"
  "github.com/Iwark/paramstostruct"
)

type Person struct {
  Name   string `json:"-"`
  Age    int    `json:"age"`
  Gender string `json:"gender"`
}

func httpHandler(res http.ResponseWriter, req *http.Request) {
  req.ParseForm()

  formValue := req.Form()
  // map[string][]string{
  //   "person[name]":   []string{"Iwark"},
  //   "person[age]":    []string{"24"},
  //   "person[gender]": []string{"man"},
  //   "person[hobby]":  []string{"playing the piano"},
  // }

  result := &Person{}
  err := paramstostruct.Convert(formValue, result)
  if err != nil {
      fmt.Println(err)
  }
  fmt.Printf("%#v", result)
}
```

## License

paramstostruct is released under the MIT License.