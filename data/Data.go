package data

import (
	"fmt"
	"io/ioutil"
	"os"

	// "net/http"
	// "strconv"

	"encoding/json"

	"github.com/bedrock17/router"
)

//Get : get data
func Get(c *router.Context) {

	filePath := fmt.Sprintf("test/%s.json", c.Param["data"])

	// if _, err := os.Stat(filePath); err == nil {
	// path/to/whatever exists
	// fmt.Fprintf(c.ResponseWriter, "%s found", filePath)

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Fprint(c.ResponseWriter, err)
	} else {
		fmt.Fprintf(c.ResponseWriter, "%s", data)
	}

	// } else {
	// 	fmt.Fprintf(c.ResponseWriter, "%s not found", filePath)
	// }
	// open

	// file, err := os.Open("") // For read access.
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// //닫기
	// defer file.close()

	// fileInfo, fierr := file.Stat()

	// if fierr != nil {
	// 	return fierr
	// }

	// //read
	// data := make(fileInfo.Size())

	//

	// fmt.Fprint(c.ResponseWriter, c.Param["data"])
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

	err = os.Mkdir("test", 0644)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprint(c.ResponseWriter, string(b))

	filePath := fmt.Sprintf("test/%s.json", c.Param["data"])
	file, err := os.Create(filePath)
	if err == nil {
		defer file.Close()
		fmt.Fprint(file, string(b))
		fmt.Println("File create")
	} else {
		fmt.Println("post memo err", err)
	}

}
