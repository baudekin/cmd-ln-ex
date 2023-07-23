package main

import (
	top "github.com/baudekin/cmd-ln-ex"
	suba "github.com/baudekin/cmd-ln-ex/subpacka"
	subb "github.com/baudekin/cmd-ln-ex/subpackb"
)

func main() {
	top.HelloTop()
	suba.Hello()
	subb.Hello()
}
