package main

import (
	"flag"
	"fmt"
	"os"
)


type Cmd struct {
	helpFlag bool		// 帮助标志
	versionFlag bool	// 版本标志
	cpOption string		// classpath选项
	XjreOption string   // 启动路径
	class string		// 主类名
	args []string		// 参数
}

func parseCmd() *Cmd {
	cmd := &Cmd{}
	// 调用flag包提供的Var函数设置需要解析的选项
	flag.Usage = printUsage
	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.helpFlag, "?", false, "print help message")
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version and exit")
	flag.StringVar(&cmd.cpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.cpOption, "cp", "", "classpath")
	flag.StringVar(&cmd.XjreOption, "Xjre", "", "path to jre")
	// 解析上述选项
	flag.Parse()
	// 如果解析成功，调用下面函数和获取其他没有被解析的参数
	args := flag.Args()
	if len(args) >0 {
		cmd.class = args[0]		// 主类名
		cmd.args =  args[1:]	// 传递给主类的参数
	}
	return cmd
}

func printUsage() {
	fmt.Print("Usage:%s [-options] class[args...]\n", os.Args[0])
}