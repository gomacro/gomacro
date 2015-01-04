package main

import (
	"fmt"
	. "github.com/anlhord/goerr"
	"go/ast"
)

func errA(f *ast.File, err error) (*ast.File) {
	if err != nil {
		XQZ(nil)

		fmt.Println("No file to preprocess.\n")
		Return()
	}
	return f
}
