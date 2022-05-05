package main

import (
	"os"
)

var rcURL = os.Getenv("RC_URL")

func main() {
	a := App{}
	a.Initialize(rcURL)
	a.Run(":8000")
}
