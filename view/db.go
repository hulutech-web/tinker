package view

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	tinker "github.com/hulutech-web/goravel-tinker"
	"github.com/hulutech-web/goravel-tinker/symbols"
	"github.com/pterm/pterm"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
	"os"
	"os/exec"
	"reflect"
	"strings"
	"time"
)

func StartYaegiDatabase() {

	dbFileContent := tinker.GetDBFileContent()
	//睡500ms
	time.Sleep(500 * time.Millisecond)
	//执行os清屏，执行exec.Command clear指令
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout

	fmt.Println("StartYaegiDatabase start...")

	// 创建一个新的解释器实例
	i := interp.New(interp.Options{
		GoPath: "/dev/null",
	})

	// 导入标准库
	i.Use(stdlib.Symbols)
	//导入自定义的符号
	i.Use(symbols.Symbols)
	fmt.Println("Entering Yaegi REPL. Type your Go code below (type 'exit' or 'quit' to return to menu):")
	fmt.Println("----------------------------------------  input command to start !-------------------------------------")
	pflag.String("db", dbFileContent, "db file path")
	if err := viper.BindPFlags(pflag.CommandLine); err != nil {
		panic(err)
	}
	pflag.Parse()
	file, err2 := os.ReadFile(viper.GetString("db"))
	if err2 != nil {
		fmt.Println("Error reading cache file:", err2)
	}
	_, err3 := i.Eval(string(file))
	if err3 != nil {
		fmt.Println("Error evaluating expression:", err3)
		panic(err3)
	}

	reader := bufio.NewReader(os.Stdin)

	for {

		//从startColor到endColor渐变输出
		beatifyPrint("tinker>")
		//修复：有功能缺陷，输入文字后不能箭头左右移动在原有字符中插入新字符

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input == "exit" || input == "quit" {
			fmt.Println("Exiting Yaegi REPL. Returning to shell...")
			break
		}
		//如果直接输入一个回车，继续
		if input == "" {
			continue
		}
		fmt.Println()
		res, err := i.Eval(input)
		count := 0
		if err != nil {
			fmt.Println("Error evaluating expression:", err)
		} else {
			fmt.Println(">>>")
			t := reflect.TypeOf(res.Interface())

			if t.Kind() == reflect.Map {
				resMap := res.Interface().(map[string]interface{})
				PrettyPrint(resMap)
			} else if t.Kind() == reflect.Slice {
				resSlice := res.Interface().([]map[string]interface{})
				for _, item := range resSlice {
					PrettyPrint(item)
				}
				count = len(resSlice)
			} else if t.Kind() == reflect.Struct {
				resStruct := res.Interface().(map[string]interface{})
				PrettyPrint(resStruct)
			} else if t.Kind() == reflect.Int64 {
				count = int(res.Interface().(int64))
			} else {
				// 其他类型
				fmt.Println("Unknown type:", t)
			}
			// 处理返回值

			fmt.Println(">>> end, total:", count, " rows")

		}
	}
}

func PrettyPrint(v interface{}) {
	//pterm给出更好的格式化输出，同时给与一些颜色
	b, err := json.Marshal(v)
	if err != nil {
		fmt.Println(v)
	}
	var out bytes.Buffer
	err = json.Indent(&out, b, "", "  ")
	if err != nil {
		fmt.Println(v)
	}
	pterm.NewRGB(219, 155, 52).Println(out.String())

}
