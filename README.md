# cmd-ln-ex
Example on how to use SPF13 to run command line application:


```
├── LICENSE
├── README.md
├── app
│   └── main.go
├── go.mod
├── subpacka
│   └── example1.go
├── subpackb
│   └── example2.go
└── toppack.go``
```

We have one module called cmd-ln-ex. It contains to sub packages subpacka and subpackb.  All dependancies are mangaged by 
cmd-ln-ex/mod.go.

To setup spf13 cobra:
# Install Cobra Helper: ```shell go install github.com/spf13/cobra-cli@latest```
# Switch to Application directory: ```cd app```
# Initialize command line: ```shell cobra-cli init```


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
