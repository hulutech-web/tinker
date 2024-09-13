package styles

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
)

func Tip(message string) {
	style := lipgloss.NewStyle().
		Background(lipgloss.Color("yellow")).
		Foreground(lipgloss.Color("white")).
		Padding(1, 2).
		Bold(true).
		Align(lipgloss.Center)
	fmt.Println(style.Render(message))
}

func Success(message string) {
	style := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFFFFF")). // 白色字体
		Background(lipgloss.Color("#00885e")). // 绿色背景
		Padding(0, 2).                         // 上下左右内边距
		Bold(true)                             // 加粗字体
	// 使用样式打印提示信息
	fmt.Println(style.Render(message))
}

func Warning(message string) {
	style := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFFFFF")). // 白色字体
		Background(lipgloss.Color("#af8600")). // 绿色背景
		Padding(0, 2).                         // 上下左右内边距
		Bold(true)                             // 加粗字体
	// 使用样式打印提示信息
	fmt.Println(style.Render(message))
}
