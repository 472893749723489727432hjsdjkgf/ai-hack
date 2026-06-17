package main

import (
	"fmt"

	"github.com/472893749723489727432hjsdjkgf/ai-hack/configs"
)

func main() {
	cfg := configs.GetConfig()
	dbUrl := configs.GetDbUrl()
	fmt.Println(cfg)
	fmt.Println(dbUrl)
}
