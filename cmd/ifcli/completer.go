package main

import (
	"github.com/c-bata/go-prompt"
)

func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
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
		{Text: "USE", Description: "..."}, // switch databases
	}

	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}
