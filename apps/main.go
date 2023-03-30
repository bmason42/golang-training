package main

import (
	"fmt"
	"github.com/bmason42/golang-training/pkg"
)

func main() {
	x := pkg.NewSampleStruct("internal", "external")
	fmt.Println(x.ExportData)
	fmt.Println(x.GetInternal())
	fmt.Println(x.Version)
}
