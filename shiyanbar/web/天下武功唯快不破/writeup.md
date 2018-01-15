查看response headers，发现有FLAG字段，是用base64加密的，解开以后是这样的
`P0ST_THIS_T0_CH4NGE_FL4G:R5PeqkX1P`

再结合网页的返回信息：

> There is no martial art is indefectible, while the fastest speed is the only way for long success.
  \>\>\>\>\>\>----You must do it as fast as you can!----<<<<<< <br>
> \<!-- please post what you find with parameter:key -->

所以，解法就是以解密后的值，发送post请求，来获取flag

golang代码如下：
```
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

	resPost, _ := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(data))

	defer resPost.Body.Close()

	body, _ := ioutil.ReadAll(resPost.Body)

	fmt.Println(string(body))
}
```