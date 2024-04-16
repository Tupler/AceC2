package Utils

import "github.com/c-bata/go-prompt"

func StringsToSuggests(stringList []string) []prompt.Suggest {
	s := []prompt.Suggest{}
	for _, str := range stringList {
		s = append(s, prompt.Suggest{Text: str})
	}
	return s
}
