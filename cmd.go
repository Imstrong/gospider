package main

import (
	"flag"
	"fmt"
)

type Cmd struct {
	helpFlag bool
	versionFlag bool
	//process number
	ps int
	//idle time millis
	idle int
	conf string
}
func ParseCmd() *Cmd{
	var cmd=&Cmd{}
	flag.Usage=printUsage
	flag.BoolVar(&cmd.helpFlag,"help",false,"print usage")
	flag.BoolVar(&cmd.helpFlag,"h",false,"print usage")
	flag.BoolVar(&cmd.helpFlag,"?",false,"print usage")
	flag.BoolVar(&cmd.versionFlag,"v",false,"print version")
	flag.BoolVar(&cmd.versionFlag,"version",false,"print version")
	flag.IntVar(&cmd.ps,"ps",0,"set process number")
	flag.IntVar(&cmd.idle,"idle",1000,"set idle time millis")

	flag.Parse()
	args:=flag.Args()
	if(len(args)==1){
		cmd.conf=args[0]
	}else{
		printUsage()
	}
	return cmd
}
func printUsage(){
	fmt.Print("usage:gospider.exe [-v|version|help|h|?|ps] <params>")
}