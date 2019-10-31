package main

import (
	"ch02/classpath"
	"fmt"
	"strings"
)

func main(){
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Print("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	//fmt.Print("classpath:%s class:%s args:%v\n", cmd.cpOption, cmd.class, cmd.args)

	cp := classpath.Parse(cmd.XjreOption,  cmd.cpOption)
	fmt.Print("classpath:%s class:%s args:%v\n", cp, cmd.class, cmd.args)
	className := strings.Replace(cmd.class, ".", "/", -1)
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		fmt.Printf("Could not find or load main class %s\n", cmd.class)
		return
	}
	fmt.Printf("class data:%v\n", classData)
}