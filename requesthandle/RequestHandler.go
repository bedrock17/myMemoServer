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

//Run : 핸들러를 등록하고 http 서버를 시작한다.
func Run(port int) {

	mainRouter := &router.Router{Handlers: make(map[string]map[string]router.HandlerFunc), DevMode: true}
	mainRouter.HandleFunc("GET", "/", index)
	mainRouter.HandleFunc("GET", "/memo/:data", data.Get)
	mainRouter.HandleFunc("POST", "/memo/:data", data.Post)
	mainRouter.HandleFunc("PUT", "/memo/:data", data.Update)

	portst := "localhost:" + strconv.Itoa(int(port))

	// http.Handle("/public/", new(staticHandler))
	err := http.ListenAndServe(portst, mainRouter)

	if err != nil {
		fmt.Println("RUN ERROR ", err)
	}

}
