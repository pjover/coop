package main

import "github.com/pjover/coop/cmd"

func main() {
	err := cmd.Execute()
	if err != nil {
		panic(err)
	}
}
