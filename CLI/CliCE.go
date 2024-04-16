package CLI

import (
	"ReAce/CLI/Docs"
	"ReAce/CLI/Options"
	"ReAce/Logger"
	"ReAce/Server"
	"ReAce/Utils"
	"github.com/c-bata/go-prompt"
	"strings"
)

func exitChecker(in string, breakline bool) bool {
	if in == "exit" {
		return true
	}
	return false
}

func mainExec(input string) {
	args := strings.Split(input, " ")
	if len(args) == 2 {
		switch args[0] {
		case "control":
			{
				uuid := args[1]
				if Server.Server.CheckUUID(uuid) {
					MainUUID = uuid
					options := []prompt.Option{
						prompt.OptionTitle("control"),
						prompt.OptionPrefix("[Ace]Control(" + uuid + ") > "),
						prompt.OptionPrefixTextColor(prompt.Yellow),
						prompt.OptionSetExitCheckerOnInput(exitChecker),
					}
					p := prompt.New(controlExec, controlcompleter, options...)
					p.Run()
					return
				} else {
					Logger.ALogger.Error("不正确的uuid")
					return
				}

			}

		}
	}
	if input == "Listener" {
		options := []prompt.Option{
			prompt.OptionTitle("Main"),
			prompt.OptionPrefix("[Ace]Listeners > "),
			prompt.OptionPrefixTextColor(prompt.Yellow),
			prompt.OptionSetExitCheckerOnInput(exitChecker),
		}
		p := prompt.New(listenerExec, listenercompleter, options...)
		p.Run()
		return
	} else if input == "GetClientList" {
		Server.Server.ShowSessions()
		return
	} else if input == "help" || input == "?" {
		Logger.ALogger.Info("[小提示]使用ctrl+l 可以进行清屏哦")
		Docs.MainDocs.PrintTable()
		return
	} else if input == "" {
		return
	}

	Logger.ALogger.Error("错误的指令！输入help查看帮助")
}

func maincompleter(d prompt.Document) []prompt.Suggest {

	//if d. {
	//	return MainOptionsSuggests
	//}
	arg1 := strings.ToLower(strings.Split(d.TextBeforeCursor(), " ")[0])
	if arg1 == "control" {
		return prompt.FilterHasPrefix(Utils.StringsToSuggests(Server.Server.GetSessionsUUIDS()), d.GetWordBeforeCursor(), true)
	}
	return prompt.FilterHasPrefix(Options.MainOptionsSuggests, d.TextBeforeCursor(), true)
}

func Run() {
	options := []prompt.Option{
		prompt.OptionTitle("Main"),
		prompt.OptionPrefix("[Ace] > "),
		prompt.OptionPrefixTextColor(prompt.Yellow),
		prompt.OptionSetExitCheckerOnInput(exitChecker),
	}
	p := prompt.New(mainExec, maincompleter, options...)
	p.Run()
}
