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
	if len(args) == 1 {
		//TODO
		fmt.Println("// Show Help")
		return nil
	}

	switch args[1] {
	case "h":
		fmt.Println("Show some help")
		break
	case "get":
		if len(args) == 3 {
			name := args[2]

			qls, err := GetQLs()
			if err != nil {
				return err
			}

			for _, ql := range qls {
				if ql.Name == name {
					fmt.Printf("%s", ql.Path)
					return nil
				}
			}

			fmt.Printf("%s", "UNDEFINED")
		} else {
			fmt.Println("Invalid arguments. Use: qm <name>")
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
			fmt.Println("Invalid arguments. Use: qm add <name> <path>")
			fmt.Println("Name must only be one word")
		}
		break
	// LIST COMMAND
	case "list":
		qls, err := GetQLs()
		if err != nil {
			return err
		}
		if len(qls) == 0 {
			fmt.Println("There are no quicklinks yet")
			fmt.Println("Use: qm add <name> <path>")
			return nil
		}

		var data [][]string

		for _, ql := range qls {
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
			fmt.Println("Invalid arguments. Use: qm remove|rm <name>")
			fmt.Println("Name must only be one word")
		}
		break
	}

	return nil
}
