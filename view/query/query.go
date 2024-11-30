package query

import (
	"bufio"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/goravel/framework/support/path"
	"github.com/hulutech-web/goravel-tinker/contracts"
	"github.com/hulutech-web/goravel-tinker/styles"
	"github.com/hulutech-web/goravel-tinker/symbols"
	"github.com/pkg/errors"
	"github.com/pterm/pterm"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"regexp"
	"sort"
	"strings"
	"sync"
	"time"
)

type Engine struct {
	handle func(str string) string
	mutex  sync.RWMutex

	modelPath   string
	modelDigest string
}

func NewEngine(ctx context.Context, modelPath string) (*Engine, error) {
	ret := &Engine{
		modelPath: modelPath,
	}
	if err := ret.ReloadModels(ctx); err != nil {
		return nil, errors.Wrap(err, "failed to reload")
	}
	go ret.proc(ctx)
	return ret, nil
}

func (e *Engine) proc(ctx context.Context) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			if err := e.ReloadModels(ctx); err != nil {
				logrus.WithError(err).Error("failed to reload")
				continue
			}
		}
	}

}

func (e *Engine) readDirRecursive(dir string) []byte {

	var content []byte
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
		return content
	}

	// 对文件和目录排序，确保每次读取顺序相同
	sort.Slice(files, func(i, j int) bool {
		return files[i].Name() < files[j].Name()
	})

	for _, file := range files {
		if file.IsDir() {
			content = append(content, e.readDirRecursive(filepath.Join(dir, file.Name()))...)
		} else {
			data, err := os.ReadFile(filepath.Join(dir, file.Name()))
			if err != nil {
				fmt.Println(err)
				continue
			}
			content = append(content, data...)
		}
	}
	return content
}

// 当模型变化时，重新加载模型
func (e *Engine) ReloadModels(ctx context.Context) error {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	modelPath := e.modelPath

	// 读取模型内容并标准化
	modelContent := e.readDirRecursive(modelPath)

	// 合并Go文件，处理import语句
	importLines := MergeGoFiles(string(modelContent))
	modelContentStrByte := []byte(importLines)

	// 计算模型内容的哈希值
	modelSumBin := sha256.Sum256(modelContent)
	modelDigest := hex.EncodeToString(modelSumBin[:])
	//fmt.Println("judge", "e->", e.modelDigest, "c->", modelDigest)
	if e.modelDigest != "" && e.modelDigest != modelDigest {
		go Rebuild()
	}
	modelDigestOld := e.modelDigest

	if modelDigest == modelDigestOld {
		return nil
	}
	i := interp.New(interp.Options{
		GoPath: "/dev/null",
	})
	if err := i.Use(stdlib.Symbols); err != nil {
		return errors.Wrap(err, "failed to use stdlib symbols")
	}
	if err := i.Use(symbols.Symbols); err != nil {
		return errors.Wrap(err, "failed to use symbol symbols")
	}

	if _, err := i.EvalWithContext(ctx, string(modelContentStrByte)); err != nil {
		return errors.Wrap(err, "failed to eval model")
	}
	ret, err := i.EvalWithContext(ctx, "models.WorkingTest")
	if err != nil {
		return errors.Wrap(err, "failed to eval Handle")
	}
	handle, ok := ret.Interface().(func(str string) string)
	if !ok {
		return errors.New("failed to convert Handle to func(http.ResponseWriter, *http.Request)")
	}

	if modelDigest == e.modelDigest {
		return nil
	}
	e.handle = handle
	e.modelDigest = modelDigest
	// 打印模型重新加载的信息
	fmt.Println()
	logrus.WithField("digest", modelDigest).WithField("timestamp", time.Now().Format(time.RFC3339)).Info("Folder Or Model Reloaded")
	// 定义样式
	fmt.Println()
	//这里需要执行重建
	return nil
}

func Rebuild() {
	//lipgloss创建一个背景黄色，文字白色的提示信息the models fold or file has changed, please wait a few minutes,rebuild...
	fmt.Println()
	styles.Tip("app models fold or file has changed, please wait a few seconds ...")
	// 创建一个用于停止旋转的 channel
	stopChan := make(chan struct{})

	start_time := time.Now()
	// 启动旋转动画
	go Rotating(stopChan)

	// 定义要执行的命令（`make`）和目标（`all`）
	cmd := exec.Command("make", "all")
	cmd.Dir = "packages/tinker" // 设置工作目录为包含 Makefile 的目录

	// 执行命令并获取输出
	output, err := cmd.CombinedOutput()

	if err != nil {
		log.Fatalf("Error executing command: %v\nOutput: %s", err, output)
	}

	// 停止旋转动画
	stopChan <- struct{}{}
	fmt.Println()
	// 打印命令输出
	styles.Success(fmt.Sprintf("Command executed successfully!!!  Elapsed time: %f s", time.Since(start_time).Seconds()))
}

// 合并多个Go文件为一个符合语法要求的文件
func MergeGoFiles(input string) string {
	// 正则表达式匹配 package 和 import 声明
	packageRegex := regexp.MustCompile(`(?m)^package\s+\w+`)
	importRegex := regexp.MustCompile(`(?ms)^import\s+\([\s\S]+?\)`)

	// 提取所有的 import 声明
	importMatches := importRegex.FindAllString(input, -1)

	// 合并所有 import 声明中的包，并去重
	importsSet := make(map[string]struct{})
	for _, match := range importMatches {
		importLines := strings.Split(match, "\n")
		for _, line := range importLines {
			line = strings.TrimSpace(line)
			if strings.HasPrefix(line, "\"") && strings.HasSuffix(line, "\"") {
				importsSet[line] = struct{}{}
			}
		}
	}

	// 构建去重后的 import 块
	var finalImports []string
	finalImports = append(finalImports, "import (")
	for imp := range importsSet {
		finalImports = append(finalImports, "\t"+imp)
	}
	finalImports = append(finalImports, ")")

	// 移除原始的 package 和 import 声明
	cleanedInput := packageRegex.ReplaceAllString(input, "")
	cleanedInput = importRegex.ReplaceAllString(cleanedInput, "")

	// 去除多余的空行
	cleanedInput = strings.TrimSpace(cleanedInput)

	// 拼接最终的 package、import 和内容，import 必须在类型声明之前
	packageDecl := "package models" // 假设最终的包名是 models
	finalOutput := fmt.Sprintf("%s\n\n%s\n\n%s", packageDecl, strings.Join(finalImports, "\n"), cleanedInput)
	return finalOutput
}

func StartYaegiModel() {
	fmt.Println("Entering Yaegi REPL. Type your Go code below (type 'exit' or 'quit' to return to menu):")
	//观察models目录下文件的变化
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var baseTinker contracts.Tinker
	queryfile := baseTinker.GetQueryFileContent()
	pflag.String("model", queryfile, "model file path")
	pflag.Parse()
	if err := viper.BindPFlags(pflag.CommandLine); err != nil {
		panic(err)
	}
	// 获取当前工作目录
	cwd, err_ := os.Getwd()
	if err_ != nil {
		panic(err_)
	}
	// 获取模型路径
	relativePath := viper.GetString("model")
	modelPath := filepath.Join(cwd, relativePath)
	fmt.Println("Model path:", modelPath)
	_, err := NewEngine(ctx, modelPath)
	if err != nil {
		panic(err)
	}
	//观察models目录下文件的变化

	// 创建一个新的解释器实例
	i := interp.New(interp.Options{
		GoPath: "/dev/null",
	})

	// 导入标准库
	i.Use(stdlib.Symbols)
	//导入自定义的符号
	i.Use(symbols.Symbols)
	pflag.String("query", "./funcs/query/query.go", "query file path")
	if err := viper.BindPFlags(pflag.CommandLine); err != nil {
		panic(err)
	}
	pflag.Parse()
	file, err2 := os.ReadFile(viper.GetString("query"))
	if err2 != nil {
		fmt.Println("Error reading query file:", err2)
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
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input == "exit" || input == "quit" {
			ctx.Done()
			fmt.Println("Exiting Yaegi REPL. Returning to shell...")
			break
		}
		if input == "shebang" {
			fmt.Println("enter shebang runtime ...")
			var inputshebang string
			for {
				beatifyPrint("tinker>")
				readerExpr := bufio.NewReader(os.Stdin)
				str, _ := readerExpr.ReadString('\n')
				if str == "exit" || str == "quit" {
					fmt.Println("Exiting Yaegi REPL. Returning to shell...")
					break
				}
				if strings.Contains(str, ":save") {
					//	保存inputshebang到path.Storage("/tmp/test")
					os.WriteFile(path.Storage("/tmp/test"), []byte(inputshebang), 777)
					fmt.Println("file saved...") // 打开文件
					os.Chmod(path.Storage("/tmp/test"), 0777)
					// 打开文件
					file, err_ := os.Open(path.Storage("/tmp/test"))
					if err_ != nil {
						log.Fatalf("无法打开文件: %s", err_)
					}
					defer file.Close()
					// 使用Scanner逐行读取文件
					scanner := bufio.NewScanner(file)
					lineNumber := 0
					for scanner.Scan() {
						lineNumber++ // 打印行号
						// 获取当前行
						fmt.Println(">>> line:", lineNumber)
						line := scanner.Text()
						res, err_res := i.Eval(line)
						if err_res != nil {
							fmt.Println("Error evaluating expression:", err_res)
						}
						value := reflect.ValueOf(res)
						if value.IsValid() {
							fmt.Println("=>", value)
						} else {
							fmt.Println("Invalid value, skipping...")
						}
					}
					//将/tmp/test
					break
				} else {
					inputshebang += strings.TrimSpace(str) + "\n"
					fmt.Println("Your input is:\n", inputshebang)
				}
			}
		}

	}
}

// RotatingSpinner 实现旋转效果的函数
func Rotating(stop chan struct{}) {
	go styles.Rotating(stop)
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
