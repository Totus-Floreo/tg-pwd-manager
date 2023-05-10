package application

import "github.com/mymmrac/telego"

func NewCommandParams() *telego.SetMyCommandsParams {

	commandsParams := telego.SetMyCommandsParams{}

	setCommand := NewCommand("set", "Add login and password for service send is like a service login password")
	getCommand := NewCommand("get", "Get login and password by name service")
	delCommand := NewCommand("del", "Delete login and password by name service")

	return commandsParams.WithCommands(setCommand, getCommand, delCommand)
}

func NewCommand(cmd string, desc string) telego.BotCommand {
	return telego.BotCommand{
		Command:     cmd,
		Description: desc,
	}
}
