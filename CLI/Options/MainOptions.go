package Options

import "github.com/c-bata/go-prompt"

var MainOptionsSuggests = []prompt.Suggest{
	{Text: "Listener", Description: "设置监听器"},
	//{Text: "Log", Description: "进入log视图"},
	{Text: "control", Description: "进入控制窗口"},
	{Text: "GetClientList", Description: "显示在线列表"},
	{Text: "help", Description: "查看帮助"},
	//{Text: "setRHOST", Description: "Listen Port"},
}
