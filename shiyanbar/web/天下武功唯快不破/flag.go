package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	url := "http://ctf5.shiyanbar.com/web/10/10.php"

	res, _ := http.Get(url)

	sDec, _ := base64.StdEncoding.DecodeString(res.Header["Flag"][0])

	data := fmt.Sprintf("key=%s", strings.Split(string(sDec), ":")[1])

	// fmt.Println(data)

	resPost, _ := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(data))

	defer resPost.Body.Close()

	body, _ := ioutil.ReadAll(resPost.Body)
	// fmt.Println(resPost)
	fmt.Println(string(body))
}
