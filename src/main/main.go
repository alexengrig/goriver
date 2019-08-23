/*
 * Copyright 2019 Goriver contributors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"fmt"
	"os"
)

func main() {
	argsLength := len(os.Args)
	if argsLength == 1 {
		help()
	} else {
		command := os.Args[1]
		switch command {
		case HelpCommandName:
			help()
			break
		case VersionCommandName:
			version()
			break
		default:
			fmt.Printf("%s: '%s' is not a %s command. See '%s %s'.",
				RootCommandName, command, RootCommandName, RootCommandName, HelpCommandName)
		}
	}
}

const (
	Version            string = "0.0.1"
	RootCommandName    string = "gori"
	HelpCommandName    string = "help"
	VersionCommandName string = "version"
)

func help() {
	fmt.Printf("usage: %s <command> [<args>]\n", RootCommandName)
}

func version() {
	fmt.Printf("%s version %s", RootCommandName, Version)
}
