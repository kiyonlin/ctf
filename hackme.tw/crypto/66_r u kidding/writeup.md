凯撒加密
golang代码
```
package main

import (
	"fmt"
)

func main() {
	str := "EKZF{Hs'r snnn dzrx, itrs bzdrzq bhogdq}"

	for i := 1; i <= 25; i++ {
		result := ""

		for _, v := range str {
			if v >= 65 && v <= 90 {
				v = rune((int(v)-65+i)%26 + 65)
			}

			if v >= 97 && v <= 122 {
				v = rune((int(v)-97+i)%26 + 97)
			}

			result += fmt.Sprintf("%c", v)
		}

		fmt.Println(result)

		if result[0] == 'F' {
			break
		}
	}

}
```