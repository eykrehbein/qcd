package utils

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/olekukonko/tablewriter"
)

// Commander Function that handles incoming CLI commands
func Commander(args []string) error {

	version := "1.0"

	switch args[1] {
	case "help":
		fallthrough
	case "h":
		fmt.Println(`
NAME:
  quickcd - Change directories using shortcuts

TO CHANGE DIRECTORY USE:
  qcd <shortcut-name>

VERSION:
  ` + version + `
	
COMMANDS:
  add <name> <rel_or_abs_path> 	Add a shortcut
  remove|rm <name>		Remove a shortcut
  list			Shows a list of all available shortcuts

  help, h		Shows a list of commands or help for one command
  version		Shows the version of this program
`)
		break

	case "version":
		fmt.Println("Version " + version)
		break
	case "get":
		if len(args) == 4 {
			name := args[2]
			currentDir := args[3]

			qls, err := GetQLs()
			if err != nil {
				return err
			}

			for _, ql := range qls {
				if ql.Name == name {

					// store the current directory as "back" QuickLink
					err := AddQL(QuickLink{Name: "back", Path: currentDir})
					if err != nil {
						return err
					}
					// print the path. The bash script will change directories afterwards
					fmt.Printf("%s", ql.Path)
					return nil
				}
			}

			fmt.Printf("%s", "UNDEFINED")
		} else {
			fmt.Println("Invalid arguments. Use: qcd <name>")
		}
		break
	// ADD COMMAND
	case "add":
		if len(args) == 4 {
			name := args[2]
			if name == "add" || name == "list" || name == "remove" || name == "help" {
				return errors.New("Sorry, you are not allowed to create a QuickLink with this name, because it's a internal keyword")
			}

			absFilePath, err := filepath.Abs(args[3])
			if err != nil {
				return err
			}
			err = AddQL(QuickLink{Name: name, Path: absFilePath})
			if err != nil {
				return err
			}

			fmt.Println("Successfully added QuickLink '" + name + "'")
		} else {
			fmt.Println(args)
			fmt.Println("Invalid arguments. Use: qcd add <name> <path>")
			fmt.Println("Name must only be one word")
		}
		break
	// LIST COMMAND
	case "list":
		qls, err := GetQLs()
		if err != nil {
			return errors.New("Couldn't read the store file. Try to create a new QuickLink: qcd add <name> <path>")
		}
		if len(qls) == 0 {
			fmt.Println("There are no quicklinks yet")
			fmt.Println("Use: qcd add <name> <path>")
			return nil
		}

		var data [][]string

		for _, ql := range qls {
			// ignore the "back" link because it's automatically generated
			if ql.Name == "back" {
				continue
			}
			data = append(data, []string{ql.Name, ql.Path})
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Name", "Path"})

		for _, v := range data {
			table.Append(v)
		}
		table.Render()
		break

	// REMOVE COMMAND
	case "remove":
		fallthrough
	case "rm":
		if len(args) == 3 {
			name := args[2]
			err := RemoveQL(name)
			if err != nil {
				return err
			}
			fmt.Println("Successfully removed QuickLink '" + name + "'")
		} else {
			fmt.Println("Invalid arguments. Use: qcd remove|rm <name>")
			fmt.Println("Name must only be one word")
		}
		break
	}

	return nil
}
