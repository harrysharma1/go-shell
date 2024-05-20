package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func execute_command(input string) error {
	input = strings.TrimSuffix(input, "\n")
	args := strings.Split(input, " ")
	switch args[0]{
		case "cd":
			if len(args) <2{
				return os.Chdir("/Users/harrysharma/")
			}
			if args[1] == "~/"{
				return os.Chdir("/Users/harrysharma/")
			}
			return os.Chdir(args[1])
		case "exit":
			os.Exit(0)
	}

	command := exec.Command(args[0],args[1:]...)
	
	command.Stderr = os.Stderr
	command.Stdout = os.Stdout
	
	
	return command.Run()
}

func main()  {
	reader := bufio.NewReader(os.Stdin)
	for{
		wd, err := os.Getwd()
		if err != nil{
			log.Fatal(err)
		}
		fmt.Print(wd)
		fmt.Print(" =>")
		input,err := reader.ReadString('\n')
		if err != nil{
			fmt.Fprintln(os.Stderr,err)
		}
		
		if err = execute_command(input); err != nil {
			fmt.Fprintln(os.Stderr,err)
		}
	}
}	
