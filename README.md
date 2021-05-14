# START
```
go get github.com/alvin1007/go-bf
```
# USE

```
gobf.BrainFuck("{your bf code}")
```
# EXAMPLE
```
package main

import gobf "github.com/alvin1007/go-bf"

func main() {
	fmt.Print(gobf.BrainFuck("+[>[<-[]>+[>+++>[+++++++++++>][>]-[<]>-]]++++++++++<]>>>>>>----.<<+++.<-..+++.<-.>>>.<<.+++.------.>-.<<+.<."))
}
```
