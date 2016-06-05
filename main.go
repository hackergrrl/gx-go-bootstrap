package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	cli "github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "gx-go-bootstrap"
	app.Usage = "Bootstrap a new Go package for building/installation with gx."
	app.Flags = []cli.Flag{}

	app.Action = func(c *cli.Context) error {
		var prefix string = ""

		// Fail early if some files already exist.
		var someExist bool = false
		for _, name := range AssetNames() {
			_, err := os.Stat(prefix + name)
			if err == nil {
				someExist = true
				fmt.Fprintln(os.Stderr, prefix+name+" already exists.")
			}
		}
		if someExist {
			fmt.Fprintln(os.Stderr, "Failing to write anything.")
			os.Exit(1)
		}

		// Write the assets.
		for _, name := range AssetNames() {
			writeName := prefix + name

			fmt.Println("Writing", name)

			// Create directories (if needed)
			os.MkdirAll(filepath.Dir(writeName), os.ModeDir|os.ModePerm)

			// Get the asset bytes
			bytes, err := Asset(name)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}

			// Write the bytes
			err = ioutil.WriteFile(writeName, bytes, 0744)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
		}

		return nil
	}

	app.Run(os.Args)
}
