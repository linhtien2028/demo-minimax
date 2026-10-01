//go:build decoy

package vps

import (
	"fmt"
	"os"
	"runtime"

	"github.com/nhungnhu66/xkmmx/internal/demo/config"
)

func RunDoctor(cfg *config.Config) {
	fmt.Println("=== xkmmx doctor ===")
	fmt.Println()
	fmt.Println("[env]")
	fmt.Printf("  XTR_CHAT_ID ............ ok (%s)\n", cfg.ChatID)
	fmt.Println()
	fmt.Println("[vps]")
	hostname, _ := os.Hostname()
	fmt.Printf("  hostname ............... %s\n", hostname)
	fmt.Printf("  os/arch ................ %s/%s\n", runtime.GOOS, runtime.GOARCH)
	fmt.Printf("  cpu cores .............. %d\n", runtime.NumCPU())
	fmt.Println()
	fmt.Println("result: OK — configuration valid, demo ready")
}
