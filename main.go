package main

import (
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"

	_ "voichatter/internal/packed"

	_ "voichatter/internal/logic"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/os/genv"

	"github.com/gogf/gf/v2/os/gctx"

	"voichatter/internal/cmd"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
)

func main() {
	gdb.SetConfig(gdb.Config{
		"default": gdb.ConfigGroup{
			gdb.ConfigNode{
				Type:  "mysql",
				Link:  genv.Get("MYSQL_LINK").String(),
				Debug: true,
			},
		},
	})
	cmd.Main.Run(gctx.GetInitCtx())
}
