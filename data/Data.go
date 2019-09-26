package data

import (
	"fmt"
	"io/ioutil"
	"os"

	// "net/http"
	// "strconv"

	"encoding/json"

	"github.com/bedrock17/myMemoServer/common"
	"github.com/bedrock17/router"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

//Get : get data
func Get(c *router.Context) {

	filePath := fmt.Sprintf("test/%s.json", c.Param["data"])

	if common.FileExists(filePath) {
		data, err := ioutil.ReadFile(filePath)
		check(err)

		fmt.Fprintf(c.ResponseWriter, "%s", data)
	} else {
		fmt.Fprintf(c.ResponseWriter, "%s %s", c.Param["data"], "is not exist")
	}

}

//Post : porc post data
func Post(c *router.Context) {

	body := make([]byte, c.Request.ContentLength)
	c.Request.Body.Read(body)

	var data memo
	err := json.Unmarshal(body, &data)
	check(err)

	b, err := json.Marshal(data)
	check(err)

	dirName := "test"
	err = os.Mkdir(dirName, 0644)

	filePath := fmt.Sprintf("%s/%s.json", dirName, c.Param["data"])

	if common.FileExists(filePath) {
		fmt.Fprintf(c.ResponseWriter, "%s %s", c.Param["data"], "is exist")
	} else {
		err = ioutil.WriteFile(filePath, b, 0644)
		check(err)
		fmt.Fprint(c.ResponseWriter, string(b))
	}

}

//Update : 작성된 내용 업데이트
func Update(c *router.Context) {
	body := make([]byte, c.Request.ContentLength)
	c.Request.Body.Read(body)

	var data memo
	err := json.Unmarshal(body, &data)
	check(err)

	b, err := json.Marshal(data)
	check(err)

	dirName := "test"
	err = os.Mkdir(dirName, 0644)

	filePath := fmt.Sprintf("%s/%s.json", dirName, c.Param["data"])

	if common.FileExists(filePath) {
		err = ioutil.WriteFile(filePath, b, 0644)
		check(err)
		fmt.Fprint(c.ResponseWriter, string(b))
	} else {
		fmt.Fprintf(c.ResponseWriter, "%s %s", c.Param["data"], "is not exist")
	}
}
