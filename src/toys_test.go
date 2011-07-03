package irc

import "testing"

import "utils"

type SecsToTimeTest struct {
	in_seconds int64
	out        string
}

var SecsToTimeTests = []SecsToTimeTest{
	SecsToTimeTest{0, "up 0 days, 0:00"},
	SecsToTimeTest{60, "up 0 days, 0:01"},
	SecsToTimeTest{61, "up 0 days, 0:01"},
	SecsToTimeTest{119, "up 0 days, 0:01"},
	SecsToTimeTest{120, "up 0 days, 0:02"},
	SecsToTimeTest{3600, "up 0 days, 1:00"},
	SecsToTimeTest{4800, "up 0 days, 1:20"},
	SecsToTimeTest{86400, "up 1 day, 0:00"},
	SecsToTimeTest{86460, "up 1 day, 0:01"},
	SecsToTimeTest{305100, "up 3 days, 12:45"},
}

func TestSecsToTime(t *testing.T) {
	for _, sttt := range SecsToTimeTests {
		result := utils.SecsToTime(sttt.in_seconds)
		if result != sttt.out {
			t.Errorf("STTT Failed, exp: %s, got: %s", sttt.out, result)
		}
	}
}

type WikifyTest struct {
	in_term string
	exp     string
}

var WikifyTests = []WikifyTest{
	WikifyTest{"!wiki apple", "http://en.wikipedia.org/wiki/apple"},
	WikifyTest{"!wiki apple pie", "http://en.wikipedia.org/wiki/apple_pie"},
	WikifyTest{"!wikiapple", "http://en.wikipedia.org/wiki/apple"},
	WikifyTest{"!wiki", "usage: !wiki <term>"},
}

func TestWikify(t *testing.T) {
	for _, wt := range WikifyTests {
		result := utils.Wikify(wt.in_term)
		if result != wt.exp {
			t.Errorf("WT Failed, exp: %s, got: %s", wt.exp, result)
		}
	}
}
