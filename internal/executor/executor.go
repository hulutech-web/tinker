package executor

import (
	"fmt"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/hulutech-web/goravel-tinker/view/cache"
	"github.com/hulutech-web/goravel-tinker/view/command"
	"github.com/hulutech-web/goravel-tinker/view/db"
	"github.com/hulutech-web/goravel-tinker/view/query"
	"github.com/hulutech-web/goravel-tinker/view/shebang"
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
			"--SELECTIONS: 1. database 2. model 3. cache 4. auth 5. command 6. go interpreter --\n" +
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
		{"1", "database", "Orm table", `Handle data with tablename,eg:db.Find("depts")`},
		{"2", "orm", "Orm Model", `Handle data with models,eg:db.Find("depts"),enter :save to exec`},
		{"3", "cache", "Handle cache", `Handle Cache,eg:cache.Store("company","hulutech-web") \n cache.Get("company")`},
		{"4", "command", "App key && Jwt key", "Params Generate.eg:command.JwtGenerate() or command.AppKeyGenerate() or command.Call(command string)"},
		{"5", "Go Interpreter", "Go environment", "Enter interpreter mode,input yaegi in the terminal to enter the Go environment"},
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

func Call() error {
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
		db.StartYaegiDatabase()
	case "2":
		query.StartYaegiModel()
	case "3":
		cache.StartYaegiCache()
	case "4":
		command.StartYaegiCommand()
	case "5":
		shebang.StartCommandlineMode()
	}
	return nil
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
