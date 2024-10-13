package main

import (
	"fmt"

	"github.com/JunNishimura/jsop/cmd"
)

func main() {
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
	}
}
