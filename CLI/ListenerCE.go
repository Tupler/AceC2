package CLI

import (
	"ReAce/CLI/Docs"
	"ReAce/CLI/Options"
	"ReAce/Logger"
	"ReAce/Server"
	"ReAce/Utils"
	"github.com/c-bata/go-prompt"
	"regexp"
	"strings"
)

func listenerExec(input string) {

	if input == "" {
		return
	}
	if strings.ToUpper(input) == "SHOW" {
		Server.Server.GetListeners().Show()
		return
	}
	if strings.ToUpper(input) == "HELP" || input == "?" {
		Docs.ListenerDocs.PrintTable()
	}
	args := strings.Split(input, " ")
	if len(args) > 1 {
		switch strings.ToLower(args[0]) {
		case "remove":
			{
				err := Server.Server.GetListeners().Remove(args[1])
				if err != nil {
					Logger.ALogger.Error(err.Error())
					return
				}
				Logger.ALogger.Success("监听器已删除！")
				return
			}
		case "add": //add xxx xxx xxx tcp/http/https
			{
				if len(args) < 5 {
					Logger.ALogger.Error("错误的参数数量")
					return
				}
				//获取参数
				listenerName := args[1]
				ip := args[2]
				port := args[3]
				listenerType := 0

				if strings.Contains(listenerName, " ") {
					Logger.ALogger.Error("名称中请勿添加空格")
					return
				}
				//ip匹配
				re := regexp.MustCompile(`\b(?:\d{1,3}\.){3}\d{1,3}\b`)
				matchIp := re.MatchString(ip)
				if !matchIp {
					Logger.ALogger.Error("不正确的ip格式")
					return
				}
				//端口匹配
				re = regexp.MustCompile(`\b(?:[1-9]\d{0,4}|0)\b`)
				matchPort := re.MatchString(port)
				if !matchPort {
					Logger.ALogger.Error("不正确的端口范围")
					return
				}

				//type匹配
				strings.ToUpper(args[4])
				switch strings.ToUpper(args[4]) {
				case "TCP":
					listenerType = Server.LISTEN_TCP
					break
				case "HTTP":
					listenerType = Server.LISTEN_HTTP
					break
				case "HTTPS":
					listenerType = Server.LISTEN_HTTPS
					break
				case "WEBSOCKET":
					listenerType = Server.LISTEN_WEBSOCKET
					break
				default:
					Logger.ALogger.Error("监听器类型只能为TCP/HTTP/HTTPS/WEBSOCKET")
					return

				}

				err := Server.Server.GetListeners().Add(listenerName, listenerType, ip, port)
				if err != nil {
					Logger.ALogger.Error(err.Error())
					return
				}
				Logger.ALogger.Success("成功添加监听器")
				return
			}
		case "run":
			{
				//err :=
				go func() {
					err := Server.Server.GetListeners().Run(args[1])
					if err != nil {
						Logger.ALogger.Error(err.Error())
						return
					}
				}()
				//if err != nil {
				//	Logger.ALogger.Error(err.Error())
				//	return
				//}
				Logger.ALogger.Success("成功启动监听器:", args[1])
				return
			}
		case "stop":
			{
				err := Server.Server.GetListeners().Stop(args[1])
				if err != nil {
					Logger.ALogger.Error(err.Error())
					return
				}
				Logger.ALogger.Success("成功停止监听器:", args[1])
				return
			}

		}

	}
	Logger.ALogger.Error("错误的指令")
}

func listenercompleter(d prompt.Document) []prompt.Suggest {

	//if d. {
	//	return MainOptionsSuggests
	//}
	//strings.Contains(d.GetWordBeforeCursor(), "remove ")
	arg1 := strings.ToLower(strings.Split(d.TextBeforeCursor(), " ")[0])
	if arg1 == "remove" || arg1 == "run" || arg1 == "stop" {
		return prompt.FilterHasPrefix(Utils.StringsToSuggests(Server.Server.GetListeners().GetListenersNames()), d.GetWordBeforeCursor(), true)
	}

	return prompt.FilterHasPrefix(Options.ListenerOptionsSuggests, d.TextBeforeCursor(), true)
}
