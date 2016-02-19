fmap
===
[![GoDoc](https://godoc.org/github.com/Iwark/fmap?status.svg)](https://godoc.org/github.com/Iwark/fmap)

Package ``fmap`` converts ``req.Form()`` to a struct.

## Installation

```
$ go get github.com/Iwark/fmap
```

## Example

```go
package main

import (
  "fmt"
  "github.com/Iwark/fmap"
)

type Person struct {
  Name   string `fmap:"-"`
  Age    int    `fmap:"age"`
  Gender string `fmap:"gender"`
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
  err := fmap.ConvertToStruct(formValue, result)
  if err != nil {
      fmt.Println(err)
  }
  fmt.Printf("%#v", result)
}
```

## License

fmap is released under the MIT License.