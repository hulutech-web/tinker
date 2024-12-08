package view

import (
	"fmt"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/pterm/pterm"
	"os"
	"os/exec"
)

// Define a style for the table
var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

// Define a model to hold the state of our application
type Model struct {
	table table.Model
}

var selectedOption string

func NewModel() Model {
	//先清理屏幕clear命令
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	_ = cmd.Run()

	// 设置标题
	fmt.Println("                                                                                                        ")
	fmt.Println("------------------------------------------ Welcome to tinker -------------------------------------------")
	title := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Padding(0, 0, 0, 0).
		// 宽度
		Width(100).
		// 字体大小
		Background(lipgloss.Color("#099268")).
		Foreground(lipgloss.Color("231")).
		Render(
			"--SELECTIONS: 1. database 2. cache 3. command 4. Go Interpreter 5. Goravel Orm --\n" +
				"--ACTIONS:  you can select item to handle different logic !--\n" +
				"----------------------------------------------------------------------------------------- enjoy it ! \n" +
				"                                                                                --by hulutech-web--",
		)
	fmt.Println(title)
	tip := "------------------------------  you can use up ⬆️  and down ⬇️ to choose one as you like !--------------------"
	fmt.Println(tip)
	// Define table columns and rows
	columns := []table.Column{
		{Title: "index", Width: 10},
		{Title: "name", Width: 20},
		{Title: "description", Width: 30},
		{Title: "eg", Width: 100},
	}

	rows := []table.Row{
		{"1", "database", "DB table", `enter shebang to Handle data with tablename,eg:db.Find("depts")`},
		{"2", "cache", "Handle cache", `Handle Cache,eg:cache.Store("company","hulutech-web") \n cache.Get("company")`},
		{"3", "command", "App key && Jwt key", "Params Generate.eg:jwt or appkey or command.Call(command string)"},
		{"4", "Go Interpreter", "Go environment", "Enter interpreter mode,input yaegi in the terminal to enter the Go environment"},
		{"5", "Goravel Orm", "Orm Model", `enter shebang to Handle data with model,eg:users := []models.User{} query.Query().Find(&users) enter command [:save] to Execute`},
	}

	// Initialize the table model
	tb := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(7),
	)

	// Define table styles, 设置整个表格背景色为#18dcff
	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(true)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(true)
	tb.SetStyles(s)

	return Model{
		table: tb,
	}
}

func Call() {
	// Initialize the model
	m := NewModel()
	// Create and run the Bubble Tea program
	p := tea.NewProgram(m)
	_, err := p.Run()
	if err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
	switch selectedOption {
	case "1":
		StartYaegiDatabase()
	case "2":
		StartYaegiCache()
	case "3":
		StartCommandlineMode()
	case "4":
		StartRuntimeMode()
	case "5":
		StartYaegiModel()
	}
}

// Initialize the model
func (m Model) Init() tea.Cmd { return nil }

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			if m.table.Focused() {
				m.table.Blur()
			} else {
				m.table.Focus()
			}
		case "q", "ctrl+c":
			return m, tea.Quit
		case "enter":
			if m.table.Focused() {
				row := m.table.SelectedRow()
				switch row[0] {
				case "1":
					selectedOption = "1"
					return m, tea.Quit
				case "2":
					selectedOption = "2"

					return m, tea.Quit
				case "3":
					selectedOption = "3"

					return m, tea.Quit
				case "4":
					selectedOption = "4"

					return m, tea.Quit
				case "5":
					selectedOption = "5"

					return m, tea.Quit
				default:
					return m, tea.Batch(
						tea.Printf("You selected: %s (index: %s)\n", row[1], m.table.SelectedRow()[0]),
					)
				}
			}
		}
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

// View renders the table and options as a string
func (m Model) View() string {
	tableView := baseStyle.Render(m.table.View())
	footer := "\nPress q to quit.\n"
	if m.table.Focused() {
		footer += "Press enter to select a row."
	} else {
		footer += "Use j/k to navigate rows."
	}

	// Send the UI for rendering
	return tableView + "\n\n" + footer
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
