package main

import (
	"fmt"

	"./common"

	"github.com/bedrock17/myMemoServer/requesthandle"
)

func main() {
	fmt.Println("my server start")

	common.GlobalConfig.Load()

	requesthandle.Run(common.GlobalConfig.ServicePort)
}
