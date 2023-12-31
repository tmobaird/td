package main

import (
	"errors"
)

var NoTodoSpecifedErr = errors.New("No todo specified (need id or index).")
var NoIndexOrNameErr = errors.New("Not enough arguments (need id or index and new name).")
var NoIndexOrRankErr = errors.New("Not enough arguments (need id or index and new rank).")
var NoSuchCommandErr = errors.New("No such command")
var NoConfigKeyOrValErr = errors.New("No config key or value specified")

type Cli struct {
	Command   string
	Args      []string
	Commander Commander
	Verbose   bool
	Config    Config
	PrintFunc func(a ...interface{}) (n int, err error)
}

func (c Cli) verifyArguments(argsNeeded int, ifErr error) error {
	if len(c.Args) < argsNeeded {
		return ifErr
	}
	return nil
}

func (c *Cli) runCommand() ([]Todo, error) {
	var todos []Todo
	var err error

	switch c.Command {
	case "list", "ls", "status", "st":
		todos, err = c.Commander.List(c.Commander.ReaderWriter)
	case "delete", "d", "rm":
		todos, err = c.Commander.Delete(c.Args[0:], c.Commander.ReaderWriter)
	case "add", "a":
		todos, err = c.Commander.Add(c.Args, c.Commander.ReaderWriter)
	case "done", "do":
		err = c.verifyArguments(1, NoTodoSpecifedErr)
		if err != nil {
			return todos, err
		}
		todos, err = c.Commander.Done(c.Args[0], c.Commander.ReaderWriter)
	case "undo", "un", "reset":
		err = c.verifyArguments(1, NoTodoSpecifedErr)
		if err != nil {
			return todos, err
		}
		todos, err = c.Commander.Undo(c.Args[0], c.Commander.ReaderWriter)
	case "edit", "e":
		err = c.verifyArguments(2, NoIndexOrNameErr)
		if err != nil {
			return todos, err
		}
		todos, err = c.Commander.Edit(c.Args[0], c.Args[1], c.Commander.ReaderWriter)
	case "rank", "r":
		err = c.verifyArguments(2, NoIndexOrRankErr)
		if err != nil {
			return todos, err
		}
		todos, err = c.Commander.Rank(c.Args[0], c.Args[1], c.Commander.ReaderWriter)
	case "config":
		err = c.verifyArguments(1, NoConfigKeyOrValErr)
		if err != nil {
			return todos, err
		}
		c.Config, err = c.Commander.Config(c.Args[0], c.Args[1:], c.Config, c.Commander.ReaderWriter)
	case "version", "v":
		c.PrintFunc("td version: " + Version() + "\n")
	default:
		err = NoSuchCommandErr
	}

	return todos, err
}

func (c *Cli) Run() {
	todos, err := c.runCommand()

	if err != nil {
		c.PrintFunc(ReportError(err, c.Command))
	} else if c.Command == "config" {
		c.PrintFunc(ReportConfig(c.Config))
	} else {
		c.PrintFunc(ReportTodos(todos, c.Verbose, c.Config.HideCompleted))
	}
}
