package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
)

func openGraphFile(options CliOptions) (io.Reader, error) {
	if options.ModulePath != "" {
		options.RunGoMod = true
	}

	if !options.RunGoMod && options.GraphFile == "" {
		stat, _ := os.Stdin.Stat()
		if (stat.Mode() & os.ModeCharDevice) != 0 {
			// command is run with no stdin piped to it (stdin looks like a terminal)
			// run 'go mod graph'
			options.RunGoMod = true
		}
	}

	if options.RunGoMod {
		return goModGraph(options.ModulePath)
	}

	if options.GraphFile == "" ||
		options.GraphFile == "-" {
		return os.Stdin, nil
	}

	ctnt, err := os.ReadFile(options.GraphFile)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(ctnt), nil
}

func goModGraph(path string) (io.Reader, error) {
	if path != "" {
		st, err := os.Stat(path)
		if st != nil && !st.IsDir() {
			path = filepath.Dir(path)
			_, err = os.Stat(path)
		}

		if err != nil {
			return nil, err
		}
	}

	cmd := exec.Command("go", "mod", "graph")
	cmd.Dir = path

	var b bytes.Buffer
	cmd.Stdout = &b
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		if path == "" {
			path = "current directory"
		}
		fmt.Fprintf(os.Stderr, "failed to run 'go mod graph' from %s\n", path)
		return nil, err
	}

	return &b, nil
}
