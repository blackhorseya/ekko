package main

import "github.com/blackhorseya/ekko/adapter/ekko/cmd"

// @title Ekko API
// @version 0.0.1
// @description ekko service api doc
//
// @contact.name blackhorseya
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
