package main

import (
	"github.com/ak9024/okr-generator/cmd"
)

// @title OKR Generator API
// @description This is Official API for OKR Generator API
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	cmd.Execute()
}
