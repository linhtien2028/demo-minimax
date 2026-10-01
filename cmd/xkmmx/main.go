//go:build decoy

package main

import (
	"fmt"
	"os"

	"github.com/nhungnhu66/xkmmx/internal/demo/config"
	"github.com/nhungnhu66/xkmmx/internal/demo/gateway"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "doctor" {
		fmt.Println("ok")
		return
	}

	cfg, _ := config.Load()
	if err := gateway.Run(cfg); err != nil {
		os.Exit(1)
	}
}
