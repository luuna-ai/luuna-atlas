package main

import (
	"github.com/luuna-ai/luuna-atlas/internal/atlas"
	"os"
)

func main() { os.Exit(atlas.NewApp().Run(os.Args[1:])) }
