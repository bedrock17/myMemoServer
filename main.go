package main

import (
	"fmt"

	"github.com/bedrock17/myMemoServer/requesthandle"
)

func main() {
	fmt.Println("my server start")

	requesthandle.Run(8888)
}
