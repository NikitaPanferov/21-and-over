package main

import (
	"fmt"
	"log"

	tea "github.com/charmbracelet/bubbletea"
	tcpclient "github.com/NikitaPanferov/21-and-over/client/pkg/tcp-client"
)

type model struct {
	client    *tcpclient.Client
	name      string
	connected bool
	err       error
	stage     string
	message   string
}

func initialModel() model {
	return model{
		stage: "enter_name", // Стадия ввода имени
	}
}

// Init инициализирует начальное состояние программы.
func (m model) Init() tea.Cmd {
	return nil
}

// Update обновляет модель на основе сообщений.
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch m.stage {
		case "enter_name":
			if msg.Type == tea.KeyEnter && m.name != "" {
				m.stage = "connect"
				return m, connectToServer(m.name)
			}
			if msg.Type == tea.KeyBackspace && len(m.name) > 0 {
				m.name = m.name[:len(m.name)-1]
			} else {
				m.name += msg.String()
			}

		case "connect":
			if msg.Type == tea.KeyCtrlC {
				return m, tea.Quit
			}

		default:
			if msg.Type == tea.KeyCtrlC {
				return m, tea.Quit
			}
		}

	case connectSuccess:
		m.connected = true
		m.stage = "play"
		m.message = "Подключение успешно! Ожидание игры..."
		return m, nil

	case connectFailed:
		m.err = msg.err
		m.stage = "error"
		return m, nil
	}

	return m, nil
}

// View возвращает текущее представление модели.
func (m model) View() string {
	switch m.stage {
	case "enter_name":
		return fmt.Sprintf("Введите ваше имя: %s", m.name)

	case "connect":
		return "Подключение к серверу..."

	case "play":
		return m.message + "\nНажмите Ctrl+C для выхода."

	case "error":
		return fmt.Sprintf("Ошибка: %v\nНажмите Ctrl+C для выхода.", m.err)

	default:
		return "Неизвестная стадия"
	}
}

// Команды и сообщения для работы с сервером.
type connectSuccess struct{}
type connectFailed struct {
	err error
}

func connectToServer(name string) tea.Cmd {
	return func() tea.Msg {
		client, err := tcpclient.NewClient("localhost:9000")
		if err != nil {
			return connectFailed{err}
		}

		err = client.Join(name)
		if err != nil {
			return connectFailed{err}
		}

		return connectSuccess{}
	}
}

func main() {
	// Запускаем Bubble Tea.
	if err := tea.NewProgram(initialModel()).Start(); err != nil {
		log.Fatalf("Ошибка запуска программы: %v", err)
	}
}
