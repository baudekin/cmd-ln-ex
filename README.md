# cmd-ln-ex
Example on how to use SPF13 to run command line application:


```
├── LICENSE
├── README.md
├── app
│   ├── LICENSE
│   ├── cmd
│   │   ├── create.go
│   │   ├── root.go
│   │   ├── serve.go
│   │   └── version.go
│   └── main.go
├── go.mod
├── go.sum
├── subpacka
│   └── example1.go
├── subpackb
│   └── example2.go
└── toppack.go
```

We have one module called cmd-ln-ex. It contains to sub packages subpacka and subpackb.  All dependancies are mangaged by 
cmd-ln-ex/mod.go.

To setup spf13 cobra:
1. Install Cobra Helper: ```go install github.com/spf13/cobra-cli@latest```
2. Switch to Application directory: ```cd app```
3. Add serve command: ```cobra-cli add serve```
4. Add version command: ```cobra-cli add version```
5. Add configCmd command: ```cobra-cli add -p 'configCmd```


```go
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
```
