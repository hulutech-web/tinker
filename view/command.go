package view

import (
	"bufio"
	"fmt"
	"goravel/packages/tinker/funcs/command"
	"os"
	"regexp"
	"strings"
)

func StartCommandlineMode() {
	fmt.Println("Entering command-line mode. Type your command below (type 'exit' or 'quit' to return to menu):")
	for {
		beatifyPrint("tinker>")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}
		// 移除换行符
		input = strings.TrimSpace(input)

		if input == "exit" || input == "quit" {
			fmt.Println("Exiting command-line mode. Quit ...")
			break
		} else if input == "" {
			// 忽略空输入
			continue
		}
		if input == "jwt" {
			command.JwtGenerate()
		}
		if input == "appkey" {
			command.AppkeyGenerate()
		}

		if strings.Contains(input, "command.Call(") {
			//如：command.Call(“abc"),使用正则查找字符串abc
			reg := regexp.MustCompile(`\bcommand\.Call$"([^"]+)"$`)
			matches := reg.FindStringSubmatch(input)
			if len(matches) > 1 {
				command.Call(matches[1])
			}
		}

	}
}
