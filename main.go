package main

import (
	"fmt"
	fs "go-commons/io"
)

func main() {
	fmt.Println(fs.CopyDir("D:\\autotest", "D:\\11111", true))
}
