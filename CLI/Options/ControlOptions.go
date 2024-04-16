package Options

import "github.com/c-bata/go-prompt"

var ControlOptionsSuggests = []prompt.Suggest{
	{Text: "cd", Description: "切换目录"},
	{Text: "ls", Description: "查看当前目录文件"},
	{Text: "pwd", Description: "查看当前工作目录"},
	{Text: "down", Description: "下载文件"},
	{Text: "proclist", Description: "获取进程列表"},
	{Text: "shell", Description: "执行shell命令"},
	{Text: "shellcode", Description: "执行shellcode"},
	{Text: "proclist", Description: "获取进程列表"},
	{Text: "help", Description: "查看帮助"},
	//{Text: "setRHOST", Description: "Listen Port"},
}
