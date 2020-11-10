package main

// 在 Go 中 API 以 package 的形式提供
import "flag"  // 帮助处理命令行的选项
import "fmt"
import "os"    // 定义了一个 Args 变量，存放传递给命令行的全部参数

type Cmd struct {
	helpFlag      bool
	versionFlog   bool
	cpOption      string
	class         string
	args          []string
}

func parseCmd() *Cmd {
	cmd := &Cmd{}
	flag.Usage = printUsage
	// 调用 flag 包中的各种 Var() 函数设置要解析的选项
	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.helpFlag, "?", false, "print help message")
	flag.StringVar(&cmd.cpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.cpOption, "cp", "", "classpath")
	// 调用 flag 包中的 Parse() 函数解析选项
	flag.Parse()

	// 解析成功，调用 flag 包中的 Args() 函数捕获未被解析的参数
	args := flag.Args()
	if len(args) > 0 {
		// 第一个参数是类名
		cmd.class = args[0]
		// 其他参数是传递给主类的参数
		cmd.args = args[1:]
	}
	return cmd
}

func printUsage()  {
	fmt.Printf("Usage: %s [-options] class [args...]\n", os.Args)
}
