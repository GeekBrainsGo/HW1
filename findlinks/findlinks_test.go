package findlinks

import (
	"fmt"
	"sort"
	"testing"
)

var testCases = []struct {
	description string
	search      string
	links       []string
	expected    []string
}{
	{
		"first test case",
		"search",
		[]string{
			"https://mail.ru",
			"https://yandex.ru",
			"https://brie3.github.io/page",
		},
		[]string{
			"https://mail.ru",
			"https://yandex.ru",
		},
	},
	{
		"second test case",
		"Поиск в интернете",
		[]string{
			"https://yandex.ru",
			"https://mail.ru",
			"https://www.google.com",
			"https://brie3.github.io/page",
		},
		[]string{
			"https://yandex.ru",
			"https://mail.ru",
		},
	},
}

func equal(a []string, b []string) bool {
	if len(b) != len(a) {
		return false
	}

	sort.Strings(a)
	sort.Strings(b)
	return fmt.Sprintf("%v", a) == fmt.Sprintf("%v", b)
}

func TestFindLinks(t *testing.T) {
	msg := `
	Description: %s
	Search Query: %q
	Links: %q
	Expected: %q
	Got: %q
	`
	for _, test := range testCases {
		actual := FindLinks(test.search, test.links)
		if !equal(test.expected, actual) {
			t.Fatalf(msg, test.description, test.search, test.links, test.expected, actual)
		}
		t.Logf(msg, test.description, test.search, test.links, test.expected, actual)
	}
}
