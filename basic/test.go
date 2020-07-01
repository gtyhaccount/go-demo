package main

import (
	"fmt"
	"github.com/pkg/errors"
)

func main() {
    err := errors.New("whoops")
    //fmt.Println(err)
	//
    fmt.Printf("%+v\n",err)
	//
    //fmt.Println(errors.WithStack(err))
    fmt.Printf("%+v\n",errors.WithStack(err))
}
