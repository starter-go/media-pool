package main

import (
	"os"

	mediapool "github.com/starter-go/media-pool"
	"github.com/starter-go/starter"
)

func main() {

	a := os.Args
	m := mediapool.ModuleForTest()
	i := starter.Init(a)

	i.MainModule(m)

	i.WithPanic(true).Run()
}
