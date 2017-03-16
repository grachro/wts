package main

import (
 "fmt"
 "strings"
 "os"
)

func main() {

 str := os.Args[1]

 // 最初の\\を置換
 resStr1 := strings.Replace(str, `\\`, "smb://", 1)

 // 残りの\を全部置換
 resStr2 := strings.Replace(resStr1, `\`, "/", -1)
 fmt.Println(resStr2)
}
