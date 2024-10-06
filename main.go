package main

import "github.com/JunNishimura/jsop/cmd"

func main() {
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}
