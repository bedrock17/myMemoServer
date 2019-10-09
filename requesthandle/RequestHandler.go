package requesthandle

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/bedrock17/myMemoServer/data"
	"github.com/bedrock17/router"
)

func index(c *router.Context) {
	fmt.Fprintf(c.ResponseWriter, "Welcome!")
}

func staticHandle(c *router.Context) {
	http.NotFound(c.ResponseWriter, c.Request) //정적파일을 찾지 못한경우
}

//Run : 핸들러를 등록하고 http 서버를 시작한다.
func Run(port int) {

	data.MemoInit() //load config

	server := router.NewServer()
	server.DevMode = true
	server.AppendMiidleWare(router.LogHandler)
	server.AppendMiidleWare(router.RecoverHandler)
	server.AppendMiidleWare(router.StaticHandler)

	server.HandleFunc("GET", "/", index)
	server.HandleFunc("GET", "/memo", data.GetList)
	server.HandleFunc("OPTIONS", "/memo/:data", data.Options)
	server.HandleFunc("GET", "/memo/:data", data.Get)
	server.HandleFunc("POST", "/memo/:data", data.Post)
	server.HandleFunc("PUT", "/memo/:data", data.Update)
	server.HandleFunc("DELETE", "/memo/:data", data.Delete)

	server.HandleFunc("GET", "/static/*", staticHandle)

	addr := "localhost:" + strconv.Itoa(int(port))
	server.Run(addr)

}
