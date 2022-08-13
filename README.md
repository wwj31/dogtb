# dogtb
`dogtb` is formatting tools about tabular string
## primary module
```shell
 go get -u github.com/wwj31/dogtb
```

# Quick Start
```go
package main

import (
	"fmt"
	"github.com/wwj31/dogtb"
)

func main() {
	type St struct {
		Name  string
		Age   int `dogtg:"user's Age'"`
		Class int
	}

	st1 := &St{Name: "wwj", Age: 15, Class: 1}
	st2 := &St{Name: "nibgy", Age: 100, Class: 2}
	st3 := &St{Name: "zhgyru", Age: 31, Class: 3}
	array := []*St{st1, st2, st3}

	tab, _ := dogtb.Create(array)
	fmt.Println(tab.String())  // style 0
	fmt.Println(tab.String(1)) // style 1
	fmt.Println(tab.String(2)) // style 2
	fmt.Println(tab.String(3)) // style 3
}
```
here is output of this code:

```
┏━━━━━━━━┳━━━━━━━━━━━━━━┳━━━━━━━━┓
┃  Name  ┃  user's Age' ┃  Class ┃
┣━━━━━━━━╋━━━━━━━━━━━━━━╋━━━━━━━━┫
┃  530µs ┃      15      ┃    1   ┃
┃  nibgy ┃      100     ┃    2   ┃
┃ zhgyru ┃      31      ┃    3   ┃
┗━━━━━━━━┻━━━━━━━━━━━━━━┻━━━━━━━━┛

╔════════╦══════════════╦════════╗
║  Name  ║  user's Age' ║  Class ║
╠════════╬══════════════╬════════╣
║  530µs ║      15      ║    1   ║
║  nibgy ║      100     ║    2   ║
║ zhgyru ║      31      ║    3   ║
╚════════╩══════════════╩════════╝

+--------+--------------+--------+
|  Name  |  user's Age' |  Class |
+--------+--------------+--------+
|  530µs |      15      |    1   |
|  nibgy |      100     |    2   |
| zhgyru |      31      |    3   |
+--------+--------------+--------+
```
