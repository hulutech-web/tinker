package command

import (
	"bufio"
	"fmt"
	"github.com/hulutech-web/goravel-tinker/symbol"
	_ "github.com/hulutech-web/symbols/view/db"
	"github.com/pterm/pterm"
	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
	"os"
	"os/exec"
	"strings"
	"time"
)

func previewCommandStr() string {
	str := `
import "github.com/goravel/framework/facades"

func JwtGenerate() {
	facades.Artisan().Call("jwt:generate")
}

func AppkeyGenerate() {
	facades.Artisan().Call("key:generate")
}

func Call(command string) {
	facades.Artisan().Call(command)
}

`
	return str
}

func StartYaegiCommand() {
	//睡500ms
	time.Sleep(500 * time.Millisecond)
	//执行os清屏，执行exec.Command clear指令
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout

	fmt.Println("startYaegiRepl start...")

	// 创建一个新的解释器实例
	i := interp.New(interp.Options{
		GoPath: "/dev/null",
	})

	// 导入标准库
	i.Use(stdlib.Symbols)
	//导入自定义的符号
	i.Use(symbol.Symbols)
	fmt.Println("Entering Yaegi REPL. Type your Go code below (type 'exit' or 'quit' to return to menu):")
	fmt.Println("----------------------------------------  input command to start !-------------------------------------")
	envStr := previewCommandStr()
	_, err3 := i.Eval(envStr)
	if err3 != nil {
		fmt.Println("Error evaluating expression:", err3)
		panic(err3)
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		//从startColor到endColor渐变输出
		beatifyPrint("tinker>")
		//如果直接输入一个回车，继续

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input == "" {
			continue
		}
		if input == "exit" || input == "quit" {
			fmt.Println("Exiting Yaegi REPL. Returning to shell...")
			break
		}
		res, err := i.Eval(input)
		if err != nil {
			fmt.Println("Error evaluating expression:", err)
		} else {
			// 输出结果
			fmt.Printf("Result: %v\n", res)
		}
	}
}

func beatifyPrint(input string) {
	startColor := pterm.NewRGB(0, 136, 94) // 蓝色
	endColor := pterm.NewRGB(67, 53, 255)  // 红色

	str := "Tinker> "
	strs := []rune(str)

	for i := 0; i < len(strs); i++ {
		color := startColor.Fade(0, float32(len(strs)), float32(i), endColor)
		color.Print(string(strs[i]))
	}
}
