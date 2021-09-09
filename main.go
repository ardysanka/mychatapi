package main

import "os"

func main() {
	a := App{}
	a.Initialize(
		os.Getenv("postgres"),
		os.Getenv("qwerty"),
		os.Getenv("postgres"))

	a.Run(":8010")
}
