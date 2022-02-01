package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/ikorihn/xron"
	"github.com/urfave/cli/v2"
)

func main() {

	var xmlFilePath string

	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "file",
				Usage:       "xml file to parse",
				Aliases:     []string{"f"},
				Destination: &xmlFilePath,
			},
		},
		Action: func(c *cli.Context) error {
			if len(xmlFilePath) == 0 {
				return errors.New("please specify file")
			}
			f, err := os.Open(xmlFilePath)
			if err != nil {
				return err
			}
			xpaths := xron.ConvertXmlToXpath(f)
			for _, v := range xpaths {
				fmt.Println(v)
			}
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
