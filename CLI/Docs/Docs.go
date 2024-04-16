package Docs

import (
	"fmt"
	"github.com/jedib0t/go-pretty/table"
)

type CommandDoc struct {
	Cmd     string
	Desc    string
	Usage   string
	Example string
}
type CommandDocs struct {
	CommandDocs []CommandDoc
}

var MainDocs CommandDocs = CommandDocs{[]CommandDoc{
	{Cmd: "listeners", Desc: "进入监听器视图", Usage: "Listeners", Example: "Listeners"},
	{Cmd: "control", Desc: "进入控制试图", Usage: "control uuid", Example: "control abcdefg12344"},
	{Cmd: "help", Desc: "帮助", Usage: "Help", Example: "Help"},
}}
var ControlDocs CommandDocs = CommandDocs{[]CommandDoc{
	{Cmd: "shellcode", Desc: "执行shellcode", Usage: "shellcode path", Example: "shellcode c:\\1.bin"},
	{Cmd: "shell", Desc: "执行shell命令", Usage: "shell cmd", Example: "shell whoami"},
	{Cmd: "proclist", Desc: "获取进程列表", Usage: "proclist", Example: "proclist"},
	{Cmd: "cd", Desc: "切换工作目录", Usage: "cd xx", Example: "cd .."},
	{Cmd: "ls", Desc: "查看目录下的文件", Usage: "ls", Example: "ls"},
	{Cmd: "pwd", Desc: "查看工作目录", Usage: "pwd", Example: "pwd"},
	{Cmd: "down", Desc: "下载文件", Usage: "down xxx", Example: "down 1.txt"},
}}
var ListenerDocs CommandDocs = CommandDocs{[]CommandDoc{
	{Cmd: "add", Desc: "添加监听器", Usage: "add name ip port TCP/HTTP/HTTPS/WEBSOCKET", Example: "add xxx 192.168.0.1 8888 TCP"},
	{Cmd: "remove", Desc: "删除监听器", Usage: "remove name", Example: "remove xxx"},
	{Cmd: "run", Desc: "启动监听器", Usage: "run name", Example: "run xxx"},
	{Cmd: "stop", Desc: "停止监听器", Usage: "stop name", Example: "stop xxx"},
	{Cmd: "show", Desc: "查看监听器列表", Usage: "show", Example: "show"},
	{Cmd: "help", Desc: "帮助", Usage: "Help", Example: "Help"},
}}

func (d *CommandDocs) PrintTable() {
	t := table.Table{}
	header := table.Row{"COMMAND", "DESCRIPTION", "USAGE", "EXAMPLE"}
	t.AppendHeader(header)
	t.Style().Options.SeparateRows = true
	for _, each := range d.CommandDocs {
		t.AppendRow(table.Row{each.Cmd, each.Desc, each.Usage, each.Example})

	}
	fmt.Println(t.Render())
}
