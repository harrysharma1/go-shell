package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"strings"
	"errors"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"

)

var error_no_villain = errors.New("Forgot to input name")

func search_db(moniker string) string {
	db,err := sql.Open("sqlite3","villains.db")
	if err != nil {
		log.Fatal(err)
	}
}

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
		case "villain":
			if len(args)<2{
				return error_no_villain;
			}
			name := strings.Join(args[1:]," ")
			switch name{
			case "Joker","joker":
				data := `
			Name: Unkown
			Age: 34
			Height: 6ft 8"
			Weight: 195lbs
			`
				return fmt.Errorf(data)
			case "Bane","bane":
				data := `
			Name: Eduardo Dorrance
			Age: 62
			Height: 6ft 8"
			Weight: 350lbs
			`
				return fmt.Errorf(data)
			case "Harley Quinn","harley quinn","harley Quinn","Harley quinn":
				data:=`
			Name: Dr. Harleen Quinzel
			Age: 29
			Height: 5ft 6"
			Weight: 140lbs
			`
				return fmt.Errorf(data)
			default:
				return fmt.Errorf("Villain not found")
			}
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
