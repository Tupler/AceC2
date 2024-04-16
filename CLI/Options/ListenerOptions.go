package Options

import "github.com/c-bata/go-prompt"

var ListenerOptionsSuggests = []prompt.Suggest{
	{Text: "add", Description: "添加监听器"},
	{Text: "remove", Description: "删除监听器"},
	{Text: "run", Description: "启动监听器"},
	{Text: "stop", Description: "停止监听器"},
	{Text: "show", Description: "查看监听器列表"},
	{Text: "help", Description: "查看帮助"},
	//{Text: "setRHOST", Description: "Listen Port"},
}
