/*
Copyright (c) 2023. Vili and contributors.

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
)

var version = "0.1-dev"

func printBanner() {
	banner := fmt.Sprintf(`
 ___  ________  ________     ___    ___ _________  ________  ________  ___       ________      
|\  \|\   ____\|\   __  \   |\  \  /  /|\___   ___\\   __  \|\   __  \|\  \     |\   ____\     
\ \  \ \  \___|\ \  \|\  \  \ \  \/  / ||___ \  \_\ \  \|\  \ \  \|\  \ \  \    \ \  \___|_    
 \ \  \ \  \  __\ \  \\\  \  \ \    / /     \ \  \ \ \  \\\  \ \  \\\  \ \  \    \ \_____  \   
  \/  /\ \  \|\  \ \  \\\  \  /     \/       \ \  \ \ \  \\\  \ \  \\\  \ \  \____\|____|\  \  
  /  // \ \_______\ \_______\/  /\   \        \ \__\ \ \_______\ \_______\ \_______\____\_\  \ 
 /_ //   \|_______|\|_______/__/ /\ __\        \|__|  \|_______|\|_______|\|_______|\_________\
|__|/                       |__|/ \|__|                                            \|_________|

v%s ~~by Vili (https://vili.dev)
	`, version)
	color.Cyan(banner)
}

func printAbout() {
	fmt.Println(color.GreenString("GoXTools, port of H4X-Tools and its successor written in Go."))
	time.Sleep(2 * time.Second)
}

func printMenu() {
	menuOptions := map[string]string{
		"1": "Example",
		"2": "About",
		"3": "Exit",
	}

	fmt.Println()
	for key, value := range menuOptions {
		fmt.Printf("[%s] %s\n", key, strings.ToTitle(strings.ReplaceAll(value, "_", " ")))
	}
	fmt.Println()
}

func handleExit() {
	color.Yellow("Quitting!")
	time.Sleep(1 * time.Second)
	os.Exit(0)
}

func main() {
	for {
		printBanner()
		time.Sleep(1 * time.Second)
		printMenu()
		fmt.Print("[$] Select your option ~> \t")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		option := strings.TrimSpace(scanner.Text())

		switch option {
		case "1":
			fmt.Println(color.GreenString("Test."))
		case "2":
			printAbout()
		case "3":
			handleExit()
		default:
			color.Red("Invalid option!")
			time.Sleep(2 * time.Second)
		}
	}
}
