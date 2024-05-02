package main

import (
	"os"
	"fmt"
	"strings"
	"bufio"
	"encoding/base64"
	"github.com/pborman/getopt/v2"
	"github.com/getlantern/systray"
	_ "embed"
)

const help string =
`Usage: sicon CMDS... < CMDS > IDS

Create a system tray icon as defined by the commands received from
the standard input and print the selected menu item IDs to the
standard output.

Options:

  -n : Do not read commands from standard input.

Commands:

  title=TITLE                : Set the icon title.
  tooltip=TOOLTIP            : Set the icon tooltip.
  icon=(b64:BASE64_ICO|FILE) : Set the icon image.

  menu_add=ID|LABEL|TOOLTIP          : Add a menu item.
  menu_add_quit=LABEL|TOOLTIP        : Add a quit menu item (quit ID).
  menu_add_separator                 : Add a menu separator.
  menu_hide=ID                       : Hide a menu item.
  menu_show=ID                       : Show a menu item.
  menu_icon=ID|(b64:BASE64_ICO|FILE) : Add icon to menu item.

Copyright (c) 2024 - Harkaitz Agirre - All rights reserved.`

var menuItems map[string]*systray.MenuItem = map[string]*systray.MenuItem {}
//go:embed sicon.ico
var defaultIcon []byte
var cmdline []string


func main() {
	var err   error
	
	// Error manager.
	defer func() {
		if err != nil {
			fmt.Fprintf(os.Stderr, "sicon: error: %v\n", err.Error())
			os.Exit(1)
		}
	}()
	
	// Parse command line arguments.
	hFlag := getopt.BoolLong("help", 'h')
	nFlag := getopt.BoolLong("no-stdin", 'n')
	getopt.SetUsage(func() { fmt.Println(help) })
	getopt.Parse()
	if *hFlag { getopt.Usage(); return }
	
	// Read commands from command line.
	cmdline = getopt.Args()
	
	// Read commands from standard input.
	if !*nFlag {
		go func() {
			var line     string
			var err      error
			var scanner *bufio.Scanner
			
			scanner = bufio.NewScanner(os.Stdin)
			for scanner.Scan() {
				line = scanner.Text()
				err = runCommand(line)
				if err != nil {
					fmt.Fprintf(os.Stderr, "sicon: error: %v\n", err.Error())
					continue
				}
			}
		}()
	}
	systray.Run(onReady, onExit)
}

func onReady() {
	var err error
	systray.SetIcon(defaultIcon)
	for _, arg := range cmdline {
		err = runCommand(arg)
		if err != nil { fmt.Fprintf(os.Stderr, "sicon: error: %v\n", err.Error()) }
	}
}

func onExit() {
	// clean up here
}

func runCommand(line string) (err error) {
	var cmd			[]string
	var args		[]string
	var menuItem	 *systray.MenuItem
	var cmd0, cmd1	  string
	var found		  bool
	var icon		[]byte
	
	// Split command and arguments.
	cmd = strings.SplitN(line, "=", 2)
	if len(cmd) < 2 {
		return fmt.Errorf("Invalid command: %s", line)
	}
	cmd0 = cmd[0]
	cmd1 = cmd[1]
	
	// Process the command.
	switch cmd0 {
	case "title":
		systray.SetTitle(cmd1)
	case "tooltip":
		systray.SetTooltip(cmd1)
	case "icon":
		icon, err = readIcon(cmd1)
		if err != nil { return err }
		systray.SetIcon(icon)
		
	case "menu_add":
		args = strings.SplitN(cmd1, "|", 3)
		switch len(args) {
		case 1:  menuItem = systray.AddMenuItem(args[0], "")
		case 2:  menuItem = systray.AddMenuItem(args[1], "")
		case 3:  menuItem = systray.AddMenuItem(args[1], args[2])
		default: return fmt.Errorf("Invalid add-menu command: %s", line)
		}
		go func() {
			for range menuItem.ClickedCh {
				fmt.Println(args[0])
			}
		}()
		menuItems[args[0]] = menuItem
	case "menu_add_quit":
		args = strings.SplitN(cmd1, "|", 2)
		switch len(args) {
		case 1: menuItem = systray.AddMenuItem(args[0], "")
		case 2: menuItem = systray.AddMenuItem(args[0], args[1])
		default: return fmt.Errorf("Invalid add-quit command: %s", line)
		}
		go func() {
			for range menuItem.ClickedCh {
				systray.Quit()
			}
		}()
		menuItems["quit"] = menuItem
	case "menu_add_separator":
		systray.AddSeparator()
	case "menu_hide":
		menuItem, found = menuItems[cmd1]
		if found { menuItem.Hide() }
	case "menu_show":
		menuItem, found = menuItems[cmd1]
		if found { menuItem.Show() }
	case "menu_icon":
		args = strings.SplitN(cmd1, "|", 2)
		if len(args) < 2 {
			return fmt.Errorf("Invalid number of arguments: %s", line)
		}
		menuItem, found = menuItems[args[0]]
		if !found {
			return fmt.Errorf("Menu item not found: %s", args[0])
		}
		icon, err = readIcon(args[1])
		if err != nil { return err }
		menuItem.SetIcon(icon)
	default:
		return fmt.Errorf("Invalid command: %s", line)
	}
	
	return nil
}

func readIcon(val string) (b []byte, err error) {
	var icon []byte
	
	// Check if the icon is a base64 encoded string.
	if strings.HasPrefix(val, "b64:") {
		icon, err = base64.StdEncoding.DecodeString(val[4:])
		if err != nil { return nil, err }
		return icon, nil
	}
	
	// Check if the icon is a file.
	icon, err = os.ReadFile(val)
	if err != nil { return nil, err }
	return icon, nil
}

