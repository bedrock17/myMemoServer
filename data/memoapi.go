package data

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	// "net/http"
	// "strconv"

	"encoding/json"

	"github.com/bedrock17/myMemoServer/common"
	"github.com/bedrock17/router"
)

//MemoInit : memo경로 설정
func MemoInit() {
	common.GlobalConfig.Load()
}

//GetList : 저장된 글 목록
func GetList(c *router.Context) {
	var memos memoList
	dirName := common.GlobalConfig.DataPath

	err := filepath.Walk(dirName, func(path string, info os.FileInfo, err error) error {
		if path != dirName {

			old := path[:len(dirName)+1]
			path = strings.Replace(path, old, "", 1)
			path = path[:len(path)-len(".json")]

			memos.Memos = append(memos.Memos, path)
		}
		return nil
	})
	common.Check(err)

	data, err := json.Marshal(memos)

	common.Check(err)

	fmt.Fprintf(c.ResponseWriter, "%s", data)

}

//Get : 읽기
func Get(c *router.Context) {
	fmt.Println(common.GlobalConfig)
	dirName := common.GlobalConfig.DataPath
	filePath := fmt.Sprintf("%s/%s.json", dirName, c.Param["data"])

	if common.FileExists(filePath) {
		data, err := ioutil.ReadFile(filePath)
		common.Check(err)

		fmt.Fprintf(c.ResponseWriter, "%s", data)
	} else {
		fmt.Fprintf(c.ResponseWriter, "%s %s", c.Param["data"], "is not exist")
	}

}

//Post : 생성
func Post(c *router.Context) {

	body := make([]byte, c.Request.ContentLength)
	c.Request.Body.Read(body)

	var data memo
	err := json.Unmarshal(body, &data)
	common.Check(err)

	b, err := json.Marshal(data)
	common.Check(err)

	dirName := common.GlobalConfig.DataPath
	err = os.Mkdir(dirName, 0644)

	filePath := fmt.Sprintf("%s/%s.json", dirName, c.Param["data"])

	if common.FileExists(filePath) {
		fmt.Fprintf(c.ResponseWriter, "%s %s", c.Param["data"], "is exist")
	} else {
		err = ioutil.WriteFile(filePath, b, 0644)
		common.Check(err)
		fmt.Fprint(c.ResponseWriter, string(b))
	}

}

//Update : 작성된 내용 업데이트
func Update(c *router.Context) {
	body := make([]byte, c.Request.ContentLength)
	c.Request.Body.Read(body)

	var data memo
	err := json.Unmarshal(body, &data)
	common.Check(err)

	b, err := json.Marshal(data)
	common.Check(err)

	dirName := common.GlobalConfig.DataPath
	filePath := fmt.Sprintf("%s/%s.json", dirName, c.Param["data"])

	if common.FileExists(filePath) {
		err = ioutil.WriteFile(filePath, b, 0644)
		common.Check(err)
		fmt.Fprint(c.ResponseWriter, string(b))
	} else {
		fmt.Fprintf(c.ResponseWriter, "%s %s", c.Param["data"], "is not exist")
	}
}

//Delete : 작성된 내용 삭제
func Delete(c *router.Context) {
	body := make([]byte, c.Request.ContentLength)
	c.Request.Body.Read(body)

	dirName := common.GlobalConfig.DataPath
	filePath := fmt.Sprintf("%s/%s.json", dirName, c.Param["data"])

	if common.FileExists(filePath) {
		os.Remove(filePath)
		fmt.Fprint(c.ResponseWriter, c.Param["data"], " is deleted")
	} else {
		fmt.Fprintf(c.ResponseWriter, "%s %s", c.Param["data"], "is not exist")
	}
}

//Options : 이 url이 어떤 method가 가능한지 응답
func Options(c *router.Context) {
	c.ResponseWriter.Header().Set("Access-Control-Allow-Headers", "*")
	c.ResponseWriter.Header().Set("Allow", "OPTIONS, GET, POST, PUT")
}
