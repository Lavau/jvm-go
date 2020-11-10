/**
 * 在 Go 中，main 是一个特殊的包，该包所在目录会被编译成可执行文件
 * Go 程序执行入口是 main() 函数，不接受参数，也无返回值
 */

package main

import "fmt"

func main()  {
	cmd := parseCmd()
	if cmd.versionFlog {
		fmt.Println("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd)  {
	fmt.Printf("classpath:%s class:%s args:%v\n", cmd.cpOption, cmd.class, cmd.args)
}