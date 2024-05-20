package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func execute_command(input string) error {
	input = strings.TrimSuffix(input, "\n")
	
	command := exec.Command(input)

	command.Stderr = os.Stderr
	command.Stdout = os.Stdout

	return command.Run()
}

func main()  {
	reader := bufio.NewReader(os.Stdin)
	for{
		fmt.Print("=>")
		input,err := reader.ReadString('\n')
		if err != nil{
			fmt.Fprintln(os.Stderr,err)
		}
		
		if err = execute_command(input); err != nil {
			fmt.Fprintln(os.Stderr,err)
		}
	}
}	
