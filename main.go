package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"runtime"
	"strings"

	"github.com/go-git/go-git/v5"
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

func print_user_dir(){
		home_dir, err := os.UserHomeDir()
		if err != nil{
			log.Fatal(err)
		}
		working_absolute_dir, err := os.Getwd()
		if err != nil{
			log.Fatal(err)
		}
		working_relative_dir := strings.Replace(working_absolute_dir,home_dir,"~",1)
		current_user,err := user.Current()
		if err != nil{
			log.Fatal(err)
		}
		host_name, err := os.Hostname()
		if err != nil {
			log.Fatal(err)
		}
		working_tree, err := git.PlainOpen(working_absolute_dir)
		if err != nil {
			log.Fatal(err)	
		}
		working_tree_name, err := working_tree.Head()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("(%s@%s)-",current_user.Username,host_name)
		fmt.Printf("[%s][%s]",working_relative_dir,working_tree_name)
		fmt.Print(" $ ")
}

func main()  {
	reader := bufio.NewReader(os.Stdin)
	for{
		print_user_dir()
		input,err := reader.ReadString('\n')
		if err != nil{
			fmt.Fprintln(os.Stderr,err)
		}
		
		if err = execute_command(input); err != nil {
			fmt.Fprintln(os.Stderr,err)
		}

	 
		
	}
}	
