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
		Id    string
		Age   int `tb:"user's Age'"`
		Class int
	}

	array := []*St{
		{Name: "wwj", Id: "502948676638566431", Age: 15, Class: 1},
		{Name: "bartholomew", Id: "501925768954674307", Age: 100, Class: 2},
		{Name: "bobo", Id: "522586759406295906", Age: 31, Class: 3},
	}
	tab, _ := dogtb.Create(array)

	fmt.Println(tab.String())  // style 0
	fmt.Println(tab.String(1)) // style 1
	fmt.Println(tab.String(2)) // style 2
	fmt.Println(tab.String(3)) // style 3
}
```
here is output of this code:

style 0
```
┌──────────────┬────────────────────┬──────────────┬────────┐
│     Name     │         Id         │  user's Age' │  Class │
├──────────────┼────────────────────┼──────────────┼────────┤
│      wwj     │ 502948676638566431 │      15      │    1   │
│  bartholomew │ 501925768954674307 │      100     │    2   │
│     bobo     │ 522586759406295906 │      31      │    3   │
└──────────────┴────────────────────┴──────────────┴────────┘
```
style 1
```
┏━━━━━━━━━━━━━━┳━━━━━━━━━━━━━━━━━━━━┳━━━━━━━━━━━━━━┳━━━━━━━━┓
┃     Name     ┃         Id         ┃  user's Age' ┃  Class ┃
┣━━━━━━━━━━━━━━╋━━━━━━━━━━━━━━━━━━━━╋━━━━━━━━━━━━━━╋━━━━━━━━┫
┃      wwj     ┃ 502948676638566431 ┃      15      ┃    1   ┃
┃  bartholomew ┃ 501925768954674307 ┃      100     ┃    2   ┃
┃     bobo     ┃ 522586759406295906 ┃      31      ┃    3   ┃
┗━━━━━━━━━━━━━━┻━━━━━━━━━━━━━━━━━━━━┻━━━━━━━━━━━━━━┻━━━━━━━━┛
```
style 2
```
╔══════════════╦════════════════════╦══════════════╦════════╗
║     Name     ║         Id         ║  user's Age' ║  Class ║
╠══════════════╬════════════════════╬══════════════╬════════╣
║      wwj     ║ 502948676638566431 ║      15      ║    1   ║
║  bartholomew ║ 501925768954674307 ║      100     ║    2   ║
║     bobo     ║ 522586759406295906 ║      31      ║    3   ║
╚══════════════╩════════════════════╩══════════════╩════════╝
```
stype 3
```
+--------------+--------------------+--------------+--------+
|     Name     |         Id         |  user's Age' |  Class |
+--------------+--------------------+--------------+--------+
|      wwj     | 502948676638566431 |      15      |    1   |
|  bartholomew | 501925768954674307 |      100     |    2   |
|     bobo     | 522586759406295906 |      31      |    3   |
+--------------+--------------------+--------------+--------+
```
