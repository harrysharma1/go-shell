package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"os/user"
	"strings"
	"syscall"

)



func execute_command(input string) error {
	input = strings.TrimSuffix(input, "\n")
	args := strings.Split(input, " ")
	switch args[0]{
		case "cd":
			homedir, err := os.UserHomeDir()
			if err != nil{
				log.Fatal(err)
			}
			if len(args) <2{
				return os.Chdir(homedir)
			}
			if args[1] == "~/"{
				return os.Chdir(homedir)
			}
			if args[1] == "~"{
				return os.Chdir(homedir)
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
		fmt.Printf("(%s@%s)-",current_user.Username,host_name)
		fmt.Printf("[%s]",working_relative_dir)
		fmt.Print(" $ ")
}

func print_introduction(){
		hello := `
		 Welcome to the Bat Shell:
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣀⣀⣀⣀⣤⣤⣤⣤⣤⣤⣤⣤⣤⣤⣀⣀⣀⣀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣀⣠⣤⣶⠶⠿⠛⠛⠛⠋⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠙⠛⠛⠛⠿⠶⣶⣤⣄⣀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⢀⣠⣴⠾⠟⠛⠉⣀⣀⡤⠀⠀⠀⠀⠀⠀⠀⢰⣆⠀⠀⠀⠀⣰⡇⠀⠀⠀⠀⠀⠀⠀⢤⣄⣀⡉⠛⠻⠷⣦⣄⡀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⢀⣠⡾⠟⠋⢀⣠⣴⣾⣿⣿⡟⠀⠀⠀⠀⠀⠀⠀⠀⢸⣿⣧⣤⣤⣴⣿⡇⠀⠀⠀⠀⠀⠀⠀⠀⢹⣿⣿⣷⣦⣄⡀⠙⠻⢷⣤⡀⠀⠀⠀⠀⠀
⠀⠀⠀⢀⣴⠟⠉⢀⣴⣾⣿⣿⣿⣿⣿⣿⡇⠀⠀⠀⠀⠀⠀⠀⠀⢸⣿⣿⣿⣿⣿⣿⡇⠀⠀⠀⠀⠀⠀⠀⠀⢸⣿⣿⣿⣿⣿⣿⣷⣦⡀⠉⠻⣦⡀⠀⠀⠀
⠀⠀⣴⡿⠁⢠⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣄⠀⠀⠀⠀⠀⠀⢀⣾⣿⣿⣿⣿⣿⣿⣷⡀⠀⠀⠀⠀⠀⠀⣠⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⣄⠈⢻⣦⠀⠀
⠀⣼⠏⠀⣴⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⣦⣤⣤⣤⣶⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣶⣤⣤⣤⣴⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣦⠀⠹⣧⠀
⢸⡿⠀⢰⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡇⠀⢿⡇
⢸⡇⠀⢸⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡇⠀⢸⡇
⠸⣿⠀⠈⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠁⠀⣾⡇
⠀⠹⣧⠀⠈⢿⣿⣿⣿⣿⣿⣿⣿⣿⠿⠿⢿⣿⣿⣿⣿⠿⠿⠿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠿⠿⠿⣿⣿⣿⣿⡿⠿⠿⣿⣿⣿⣿⣿⣿⣿⣿⡿⠃⠀⣼⡟⠀
⠀⠀⠙⣷⣄⠀⠙⢿⣿⣿⣿⣿⠋⠀⠀⠀⠀⠈⠻⡿⠁⠀⠀⠀⠀⠙⢿⣿⣿⣿⣿⡿⠋⠀⠀⠀⠀⠀⢻⠟⠁⠀⠀⠀⠀⠙⣿⣿⣿⣿⡿⠋⠀⣠⣾⠏⠀⠀
⠀⠀⠀⠈⠻⢷⣄⡀⠙⠻⣿⣿⡄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠻⣿⣿⡟⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢠⣿⣿⠟⠋⠀⣠⣾⠟⠁⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠙⠻⣶⣤⡀⠉⠛⠦⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢹⡟⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠴⠛⠉⢀⣤⣶⠟⠋⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠉⠛⠻⢶⣦⣄⣀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣀⣠⣤⡶⠿⠛⠉⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠉⠙⠛⠿⠶⢶⣦⣤⣤⣤⣤⣀⣀⣀⣀⣀⣀⣀⣀⣤⣤⣤⣤⣴⣶⠶⠿⠛⠛⠉⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
`
	fmt.Println(hello)
}

func main()  {
	reader := bufio.NewReader(os.Stdin)
	print_introduction()
	channel_signal := make(chan os.Signal,1)
	signal.Notify(channel_signal,syscall.SIGINT)

	go func(){
		signal := <- channel_signal
		switch signal{
			case syscall.SIGINT:
				os.Exit(0)
			
		}
	}()
	
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
