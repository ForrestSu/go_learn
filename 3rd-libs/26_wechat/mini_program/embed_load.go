package main

import (
	_ "embed"
)

//go:embed app_secret.txt
var appSecret string
