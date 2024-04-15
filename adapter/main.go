package main

import "github.com/blackhorseya/ekko/adapter/cmd"

// @title Ekko
// @version 0.1.0
// @description Ekko open api swagger
//
// @contact.name Sean Zheng
// @contact.email blackhorseya@gmail.com
// @contact.url https://blog.seancheng.space
//
// @license.name GPL-3.0
// @license.url https://spdx.org/licenses/GPL-3.0-only.html
//
// @BasePath /api
//
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
func main() {
	cmd.Execute()
}
