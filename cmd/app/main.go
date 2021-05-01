package main

import (
	"flag"
)

var cfgPath = flag.String("c", "configs/app.yaml", "set config file path")

func init() {
	flag.Parse()
}

// @title TodoList list API
// @version 0.0.1
// @description TodoList list API

// @contact.name Sean Cheng
// @contact.email blackhorseya@gmail.com
// @contact.url https://blog.seancheng.space

// @license.name GPL-3.0
// @license.url https://spdx.org/licenses/GPL-3.0-only.html

// @BasePath /api
func main() {
}
