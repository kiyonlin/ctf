16进制 + base64
golang代码
```
package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	src := "526b78425233745561476c7a49476c7a4947566863336b7349484a705a3268305033303d"

	var dst []byte

	fmt.Sscanf(src, "%X", &dst)

	sDec, _ := base64.StdEncoding.DecodeString(string(dst))

	fmt.Println(string(sDec))
}

```