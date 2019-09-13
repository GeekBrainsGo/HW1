// Package yandex implement Yandex.Disk api interactions.
package yandex

import (
	"os"
	"testing"
)

var testCases = []struct {
	description         string
	publicLink          string
	expected            bool
	deleteFileAfterTest bool
}{
	{
		"first case",
		"https://yadi.sk/i/ywY0WpUZCnysnA", // your link
		true,
		false, // delete file after test
	},
}

func TestFindLinks(t *testing.T) {
	msg := `
	Description: %s
	Public Link: %q
	Expected: %v
	Got: %v
	`
	var (
		ok  bool
		err error
	)
	for _, test := range testCases {
		if err = Disk(test.publicLink); err != nil {
			t.Fatalf(msg, test.description, test.publicLink, test.expected, err)
		}
		if _, err = os.Stat(filename); err != nil {
			t.Fatalf(msg, test.description, test.publicLink, test.expected, err)
		}
		if ok = test.expected != os.IsNotExist(err); !ok {
			t.Fatalf(msg, test.description, test.publicLink, test.expected, ok)
		}
		t.Logf(msg, test.description, test.publicLink, test.expected, ok)
		if test.deleteFileAfterTest {
			os.Remove(filename)
		}
	}
}
