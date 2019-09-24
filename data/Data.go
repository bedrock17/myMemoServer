package data

import (
	"fmt"
	// "net/http"
	// "strconv"

	"encoding/json"

	"github.com/bedrock17/router"
)

//Get : get data
func Get(c *router.Context) {
	fmt.Fprint(c.ResponseWriter, c.Param["data"])
}

//Post : porc post data
func Post(c *router.Context) {

	// reader, err := c.Request.GetBody()
	body := make([]byte, c.Request.ContentLength)
	c.Request.Body.Read(body)

	var data memo
	err := json.Unmarshal(body, &data)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println(data)

	b, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Fprint(c.ResponseWriter, string(b))
}
