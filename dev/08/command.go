package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

const (
	// CdCommand - команда "cd"
	CdCommand = "cd"
	// PwdCommand - команда "pwd"
	PwdCommand = "pwd"
	// EchoCommand - команда "echo"
	EchoCommand = "echo"
	// KillCommand - команда "kill"
	KillCommand = "kill"
	// PsCommand - команда "ps"
	PsCommand = "ps"
	// QuitCommand - команда "quit"
	QuitCommand = "quit"
)

// Commander - интерфейс команд, с методом execute
type Commander interface {
	execute(args ...string) ([]byte, error)
}

// Ниже конкретные реализации команд интерфейса Commander

// CommandPwd - структура команды "pwd"
type CommandPwd struct{}

func (c CommandPwd) execute(args ...string) ([]byte, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	return []byte(dir), nil
}

// CommandCd - структура команды "cd"
type CommandCd struct{}

func (c CommandCd) execute(args ...string) ([]byte, error) {
	dir := args[0]
	err := os.Chdir(dir)
	if err != nil {
		return nil, err
	}

	dir, err = os.Getwd()
	if err != nil {
		return nil, err
	}
	return []byte("changed directory to " + dir), nil
}

// CommandEcho - структура команды "echo"
type CommandEcho struct{}

func (c CommandEcho) execute(args ...string) ([]byte, error) {
	prefix := strings.Split("/c echo", " ")
	args = append(prefix, args...)
	cmd := exec.Command("cmd.exe", args...)
	return cmd.Output()
}

// CommandKill - структура команды "kill"
type CommandKill struct{}

func (c CommandKill) execute(args ...string) ([]byte, error) {
	pid, err := strconv.Atoi(args[0])
	if err != nil {
		return nil, err
	}
	process, err := os.FindProcess(pid)
	if err != nil {
		return nil, err
	}
	err = process.Kill()
	if err != nil {
		return nil, err
	}
	return []byte("process was killed"), nil
}

// CommandPs - структура команды "ps"
type CommandPs struct{}

func (c CommandPs) execute(args ...string) ([]byte, error) {
	cmd := exec.Command("cmd.exe", "/c tasklist")
	return cmd.Output()
}

// UnixShell - структура хранящая текущую команду и вывод
type UnixShell struct {
	command Commander
	output  io.Writer
}

// SetCommand - устанавливает текущую команду
func (s *UnixShell) SetCommand(command Commander) {
	s.command = command
}

// ExecuteCommand - выполняет команду
func (s *UnixShell) ExecuteCommand(args ...string) {
	bytes, err := s.command.execute(args...)
	if err != nil {
		fmt.Println("error:", err.Error())
		return
	}
	_, err = fmt.Fprintf(s.output, "%s\n", string(bytes))
	if err != nil {
		fmt.Println("error:", err.Error())
		return
	}
}

// ExecuteCommands - вызывает слайс команд
func (s *UnixShell) ExecuteCommands(commands []string) {
	for _, command := range commands {
		args := strings.Split(command, " ")
		com := args[0]
		if len(args) > 1 {
			args = args[1:]
		}

		switch com {
		case PwdCommand:
			cmd := &CommandPwd{}
			s.SetCommand(cmd)
		case CdCommand:
			cmd := &CommandCd{}
			s.SetCommand(cmd)
		case EchoCommand:
			cmd := &CommandEcho{}
			s.SetCommand(cmd)
		case KillCommand:
			cmd := &CommandKill{}
			s.SetCommand(cmd)
		case PsCommand:
			cmd := &CommandPs{}
			s.SetCommand(cmd)
		case QuitCommand:
			_, err := fmt.Fprintln(s.output, "exit...")
			if err != nil {
				fmt.Println("error:", err.Error())
				return
			}
			os.Exit(0)
		default:
			fmt.Println("Неизвестная команда")
			continue
		}
		s.ExecuteCommand(args...)
	}

}
