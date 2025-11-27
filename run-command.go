package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func Run(c_std string, output_name string, cli_args []string) {
	//build the project
	var c_source_files []string
	err := filepath.WalkDir("src", func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() && filepath.Ext(path) == ".c" {
			c_source_files = append(c_source_files, path)
		}
		return nil
	})
	if err != nil || len(c_source_files) == 0 {
		fmt.Println("Internal error!")
		os.Exit(1)
	}

	build_args := []string{"-Wall", "-Wextra", fmt.Sprintf("-std=%v", c_std), "-Isrc"}
	build_args = append(build_args, c_source_files...)
	build_args = append(build_args, "-o")
	build_args = append(build_args, output_name)

	build_cmd := exec.Command("gcc", build_args...)
	build_cmd.Stdout = os.Stdout
	build_cmd.Stderr = os.Stderr
	build_err := build_cmd.Run()
	if build_err != nil {
		fmt.Println(build_err)
		os.Exit(1)
	}

	//run the project
	run_args := cli_args[1:]
	run_cmd := exec.Command("./"+output_name, run_args...)
	run_cmd.Stdout = os.Stdout
	run_cmd.Stderr = os.Stderr
	run_err := run_cmd.Run()
	if run_err != nil {
		fmt.Println(run_err)
		os.Exit(1)
	}

}
