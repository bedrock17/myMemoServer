package main

import (
	"fmt"

	"github.com/bedrock17/myMemoServer/common"
	"github.com/bedrock17/myMemoServer/requesthandle"
)

func main() {
	fmt.Println("my server start")

	common.GlobalConfig.Load()

	requesthandle.Run(common.GlobalConfig.ServicePort)

}
