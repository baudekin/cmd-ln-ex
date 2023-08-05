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

## Setup module and packages
1. Setup module creating go.mod file: ```go mod init github.com/baudekin/cmd-ln-ex```
2. Add the top leve package toppackag.go referenced as "github.com/baudekin/cmd-ln-ex"
3. Add sub package a subpacka/example1.go referneced as "github.com/baudekin/cmd-ln-ex/subpacka"
4. Add sub package b subpacka/example1.go referneced as "github.com/baudekin/cmd-ln-ex/subpackb"
5. Create application directory for running the code called app. 


## To setup spf13 cobra:
1. Install Cobra Helper: ```go install github.com/spf13/cobra-cli@latest```
2. Switch to Application directory: ```cd app```
3. Create root.go: ```cobra-cli init```
4. Add serve command: ```cobra-cli add serve```
5. Add version command: ```cobra-cli add version```
6. Add the logic you want to run to go run main.go go routine. See serve.go below 


## main.go 
```go
package main

import "github.com/baudekin/cmd-ln-ex/app/cmd"

func main() {
	cmd.Execute()
}


```
## serve.go - The command to run.
```go
package cmd

import (
	"fmt"

	top "github.com/baudekin/cmd-ln-ex"
	suba "github.com/baudekin/cmd-ln-ex/subpacka"
	subb "github.com/baudekin/cmd-ln-ex/subpackb"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start Service",
	Long: `"start" service aka run hello commands.
and usage of using your command. For example:

go run main.go service`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("serve called")
		top.HelloTop()
		suba.Hello()
		subb.Hello()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	serveCmd.PersistentFlags().String("port", "-p", "Service Port")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
```
