package view

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func StartRuntimeMode() {
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

		// 为每个命令创建一个新的 yaegi 实例
		cmd := exec.Command("yaegi", "-e", input)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err = cmd.Run()
		if err != nil {
			fmt.Println("Error executing command:", err)
		}
	}
}
