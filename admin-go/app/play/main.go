package main

import (
	_ "gbaseadmin/app/play/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"gbaseadmin/app/play/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
