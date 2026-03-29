package main

import (
	_ "gbaseadmin/app/job-template/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"gbaseadmin/app/job-template/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
