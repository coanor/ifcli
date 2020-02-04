package ifcli

import (
	"github.com/c-bata/go-prompt"
)

var (
	additionalSugKey = map[string]bool{}

	/*
		ALL           ALTER         ANALYZE       ANY           AS            ASC
		BEGIN         BY            CREATE        CONTINUOUS    DATABASE      DATABASES
		DEFAULT       DELETE        DESC          DESTINATIONS  DIAGNOSTICS   DISTINCT
		DROP          DURATION      END           EVERY         EXPLAIN       FIELD
		FOR           FROM          GRANT         GRANTS        GROUP         GROUPS
		IN            INF           INSERT        INTO          KEY           KEYS
		KILL          LIMIT         SHOW          MEASUREMENT   MEASUREMENTS  NAME
		OFFSET        ON            ORDER         PASSWORD      POLICY        POLICIES
		PRIVILEGES    QUERIES       QUERY         READ          REPLICATION   RESAMPLE
		RETENTION     REVOKE        SELECT        SERIES        SET           SHARD
		SHARDS        SLIMIT        SOFFSET       STATS         SUBSCRIPTION  SUBSCRIPTIONS
		TAG           TO            USER          USERS         VALUES        WHERE
		WITH          WRITE
	*/

	suggestions = []prompt.Suggest{
		// A
		{Text: "ALTER", Description: "..."},
		{Text: "ALL", Description: "..."},
		{Text: "ANALYZE", Description: "..."},
		{Text: "ANY", Description: "..."},
		{Text: "AS", Description: "..."},
		{Text: "ASC", Description: "..."},

		// B
		{Text: "BEGIN", Description: "..."},
		{Text: "BY", Description: "..."},

		// C
		{Text: "CREATE", Description: "..."},
		{Text: "CONTINUOUS", Description: "..."},

		// D
		{Text: "DATABASE", Description: "..."},
		{Text: "DATABASES", Description: "..."},
		{Text: "DEFAULT", Description: "..."},
		{Text: "DELETE", Description: "..."},
		{Text: "DESC", Description: "..."},
		{Text: "DESTINATIONS", Description: "..."},
		{Text: "DIAGNOSTICS", Description: "..."},
		{Text: "DISTINCT", Description: "..."},
		{Text: "DROP", Description: "..."},
		{Text: "DURATION", Description: "..."},

		// E
		{Text: "END", Description: "..."},
		{Text: "EVERY", Description: "..."},
		{Text: "EXPLAIN", Description: "..."},

		// F
		{Text: "FIELD", Description: "..."},
		{Text: "FOR", Description: "..."},
		{Text: "FROM", Description: "..."},

		// G
		{Text: "GRANT", Description: "..."},
		{Text: "GRANTS", Description: "..."},
		{Text: "GROUP", Description: "..."},
		{Text: "GROUPS", Description: "..."},

		// H

		// I
		{Text: "IN", Description: "..."},
		{Text: "INF", Description: "..."},
		{Text: "INSERT", Description: "..."},
		{Text: "INTO", Description: "..."},

		// J

		// K
		{Text: "KEY", Description: "..."},
		{Text: "KEYS", Description: "..."},
		{Text: "KILL", Description: "..."},

		// L
		{Text: "LIMIT", Description: "..."},

		// M
		{Text: "MEASUREMENTS", Description: "..."},
		{Text: "MEASUREMENT", Description: "..."},

		// N
		{Text: "NAME", Description: "..."},

		// O
		{Text: "ON", Description: "..."},
		{Text: "OFFSET", Description: "..."},
		{Text: "ORDER", Description: "..."},

		// P
		{Text: "PASSWORD", Description: "..."},
		{Text: "POLICIES", Description: "..."},
		{Text: "POLICY", Description: "..."},
		{Text: "PRIVILEGES", Description: "..."},

		// Q
		{Text: "QUERIES", Description: "..."},
		{Text: "QUERY", Description: "..."},

		// R
		{Text: "READ", Description: "..."},
		{Text: "RESAMPLE", Description: "..."},
		{Text: "RETENTION", Description: "..."},
		{Text: "REPLICATION", Description: "..."},
		{Text: "REVOKE", Description: "..."},

		// S
		{Text: "SHOW", Description: "..."},
		{Text: "SELECT", Description: "..."},
		{Text: "SHARD", Description: "..."},
		{Text: "SERIES", Description: "..."},
		{Text: "SET", Description: "..."},
		{Text: "SHARD", Description: "..."},
		{Text: "SHARDS", Description: "..."},
		{Text: "SLIMIT", Description: "..."},
		{Text: "SOFFSET", Description: "..."},
		{Text: "STATS", Description: "..."},
		{Text: "SUBSCRIPTION", Description: "..."},
		{Text: "SUBSCRIPTIONS", Description: "..."},

		// T
		{Text: "TAG", Description: "..."},
		{Text: "TO", Description: "..."},

		// U
		{Text: "USER", Description: "..."},
		{Text: "USERS", Description: "..."},

		// V
		{Text: "VALUES", Description: "..."},

		// W
		{Text: "WHERE", Description: "..."},
		{Text: "WITH", Description: "..."},
		{Text: "WRITE", Description: "..."},

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
