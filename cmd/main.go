package main

import (
	"flag"
	"log/slog"
	"os"

	"plutus.io/plutus/pkg/library/resource"
)

var (
	execType string
)

func init() {
	// 2. 注册需要解析的命令行参: 参数名、默认值、参数说明
	flag.StringVar(&execType, "exectype", "local", "执行方式：web/local")
	// 3. 解析命令行参数
	flag.Parse()

	resource.Slog = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	if resource.Slog == nil {
		panic("slog init error")
	}
}

func main() {
	// 根context
	//ctx := context.Background()
}
