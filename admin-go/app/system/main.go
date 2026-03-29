package main

import (
	_ "gbaseadmin/app/system/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"gbaseadmin/app/system/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
