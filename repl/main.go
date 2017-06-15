package main

import (
	"fmt"
	"go.rls.moe/fidgetlang"
	"gopkg.in/urfave/cli.v1"
	"io/ioutil"
	"os"
	"time"
)

func main() {
	app := cli.NewApp()
	app.Name = "FidgetLang REPL"
	app.Usage = "I'm being honest, this is not one of my greatest moments"
	app.HideHelp = true
	app.HideVersion = true
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "file, f",
			Usage: "File to execute",
			Value: "fidget.fl",
		},
		cli.BoolFlag{
			Name:  "debug, d",
			Usage: "Enable Debug Language",
		},
		cli.BoolFlag{
			Name:  "dump, q",
			Usage: "Dump Compiled Program",
		},
		cli.BoolFlag{
			Name:  "hex, h",
			Usage: "Print output in hex",
		},
	}

	app.Action = func(c *cli.Context) error {
		s := fidgetlang.NewState()
		f, err := ioutil.ReadFile(c.String("file"))
		if err != nil {
			fmt.Printf("Err: %s\n", err)
			return nil
		}
		s.DebugOut = c.Bool("hex")

		if c.Bool("debug") {
			s.CompileDebug(string(f))
		} else {
			s.Compile(string(f))
		}

		if c.Bool("dump") {
			fmt.Println(s.GetDebugProgram())
			return nil
		}

		for s.Step() {
			time.Sleep(1 * time.Millisecond)
		}

		fmt.Println("\nFinished")
		return nil
	}

	app.Run(os.Args)
}
