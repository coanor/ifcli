package ifcli

import (
	"github.com/c-bata/go-prompt"
)

var (
	additionalSugKey = map[string]bool{}

	suggestions = []prompt.Suggest{
		// A
		{Text: "ALTER", Description: "..."},

		// B

		// C
		{Text: "CREATE", Description: "..."},

		// D
		{Text: "DATABASE", Description: "..."},
		{Text: "DATABASES", Description: "..."},
		{Text: "DELETE", Description: "..."},
		{Text: "DROP", Description: "..."},
		{Text: "DURATION", Description: "..."},

		// E

		// F
		{Text: "FROM", Description: "..."},
		{Text: "FIELD", Description: "..."},

		// G
		{Text: "GROUP", Description: "..."},

		// H
		// I
		// J
		// K
		{Text: "KEYS", Description: "..."},
		{Text: "KEY", Description: "..."},

		// L
		{Text: "LIMIT", Description: "..."},

		// M
		{Text: "MEASUREMENTS", Description: "..."},

		// N
		{Text: "NAME", Description: "..."},

		// O
		{Text: "ON", Description: "..."},
		{Text: "OFFSET", Description: "..."},

		// P
		{Text: "POLICIES", Description: "..."},
		{Text: "POLICY", Description: "..."},

		// Q

		// R
		{Text: "RETENTION", Description: "..."},
		{Text: "REPLICATION", Description: "..."},

		// S
		{Text: "SHOW", Description: "..."},
		{Text: "SELECT", Description: "..."},
		{Text: "SHARD", Description: "..."},
		{Text: "SERIES", Description: "..."},

		// T
		{Text: "TAG", Description: "..."},

		// U
		// V
		{Text: "VALUES", Description: "..."},

		// W
		{Text: "WHERE", Description: "..."},
		{Text: "WITH", Description: "..."},

		// X
		// Y
		// Z

		// self key words
		{Text: "ENABLE_NIL", Description: "..."},
		{Text: "DISABLE_NIL", Description: "..."},
		{Text: "RESET_SUG", Description: "..."}, // remove suggestions
		{Text: "USE", Description: "..."},       // switch databases

		// additional suggestions
	}
)

func AddSug(key string) {
	if ok, _ := additionalSugKey[key]; !ok {
		suggestions = append(suggestions, prompt.Suggest{
			Text: key, Description: "---",
		})

		additionalSugKey[key] = true
	}
}

// remove additional suggestions
func ResetSug() {

	sug := []prompt.Suggest{}
	for _, s := range suggestions {
		if ok, _ := additionalSugKey[s.Text]; !ok {
			sug = append(sug, s)
		}

	}

	additionalSugKey = map[string]bool{}
	suggestions = suggestions[:]
	suggestions = sug
}

func SugCompleter(d prompt.Document) []prompt.Suggest {

	w := d.GetWordBeforeCursor()
	if w == "" {
		return []prompt.Suggest{}
	}

	return prompt.FilterHasPrefix(suggestions, w, true)
}
