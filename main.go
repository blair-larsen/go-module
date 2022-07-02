package main

import (
	"fmt"
	"mymodule/mypackage"
	"mymodule/mypackage/extrapkg"

	//	morestrings "github.com/blair-larsen/go-hello-mod"
	"github.com/blair-larsen/hellomod"
	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello Moduels!")

			mypackage.PrintHello()
		},
	}

	fmt.Println("calling cmd.Execute()!")
	cmd.Execute()

	//	fmt.Println("hello world")

	//	mypackage.PrintHello()

	extrapkg.PrintExtra()

	//	fmt.Println(hellomod.ReverseRunes("larsen"))
	//hellomod.PrintHello()
	hellomod.PrintTest()

}
