package main

import (
	"fmt"
	"ko44d/notion-api-golang-sample/pkg"
)

func main() {
	di := pkg.DI{}
	panic(di.Init().Run(fmt.Sprintf(":%s", "8080")))
}
