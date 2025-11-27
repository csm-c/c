package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

const hello_world_c_program = `
#include <stdio.h>

int main(int argc,char* argv[]){
	printf("hello!\n");
	return 0;
}
`

// c tool version
const VERSION = "0.0.4"

func main() {
	//cli args
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("c", VERSION)
		fmt.Println("A build tool for c (like cargo for rust).")
		fmt.Println()
		fmt.Println("On Usage:")
		fmt.Println(" c help")
	} else if len(args) == 1 {
		if args[0] == "version" || args[0] == "-v" || args[0] == "--version" {
			fmt.Println("c", VERSION)
		} else if args[0] == "init" {
			fmt.Println("Scaffolding new project..")
			fmt.Println()

			//check for existing files and folders structure
			if fileExists(".gitignore") {
				fmt.Println("\nError: .gitignore file already exists!")
				os.Exit(1)
			}

			if fileExists("c.toml") {
				fmt.Println("\nError: c.toml file already exists!")
				os.Exit(1)
			}

			if folderExist("src") {
				fmt.Println("\nError: src/ folder already exists!")
				os.Exit(1)
			}

			if folderExist("deps") {
				fmt.Println("\nError: deps/ folder already exists!")
				os.Exit(1)
			}

			//create project structure
			createFileWithText(".gitignore", "deps/* \n\n!deps/.gitkeep")
			createFileWithText("c.toml", fmt.Sprintf("[project]\nname = %v\nversion = 0.0.1\nstd = c11\n\n[dependencies]", getCWD()))
			createFolder("src")
			createFileWithText("./src/main.c", hello_world_c_program)
			createFolder("deps")
			createFileWithText("./deps/.gitkeep", "")

			fmt.Println("created src/main.c")
			fmt.Println("created deps/.gitkeep")
			fmt.Println("created c.toml")
			fmt.Println("created .gitignore")
		}
	}
}

func fileExists(name string) bool {
	_, err := os.Stat(name)

	if err == nil && !errors.Is(err, os.ErrNotExist) {
		return true
	}

	return false
}

func folderExist(name string) bool {
	info, err := os.Stat(name)

	if os.IsNotExist(err) {
		return false
	}

	if err != nil {
		return false
	}

	return info.IsDir()
}

func createFileWithText(filename string, text string) {
	err := os.WriteFile(filename, []byte(text), 0644)
	if err != nil {
		fmt.Printf("Error while creating '%v'!", filename)
		os.Exit(1)
	}
}

func getCWD() string {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Internal Error!")
		os.Exit(1)
	}

	dirName := filepath.Base(cwd)
	return dirName
}

func createFolder(name string) {
	err := os.Mkdir(name, 0755)
	if err != nil {
		fmt.Printf("Error while creating '%v'!\n", name)
		os.Exit(1)
	}
}
