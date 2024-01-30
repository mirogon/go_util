package util_test

import (
	"regexp"
	"testing"
	"time"

	util "github.com/mirogon/go_util"
)

func TestInternetDateTimeTest(t *testing.T) {
	//now := time.Now().UTC()
	//internetTime := now.Format(time.RFC3339)
	//t.Error(internetTime)
}
func TestContains(t *testing.T) {
	if !util.Contains("Hello", "ello") {
		t.Error()
	}
	if util.Contains("Hello", "x") {
		t.Error()
	}
	if util.Contains("Hello", "elloo") {
		t.Error()
	}
}

func TestFind(t *testing.T) {
	index := util.Find("Hello", "o")
	if index != 4 {
		t.Error()
	}

	index = util.Find("Hello", "H")
	if index != 0 {
		t.Error()
	}
}

func TestFindLast(t *testing.T) {
	index := util.FindLast("Character", "a")
	if index != 4 {
		t.Error()
	}
}

func TestHasPrefix(t *testing.T) {
	str := "HelloThere"
	if !util.HasPrefix(str, "Hello") {
		t.Error()
	}
	if util.HasPrefix(str, "HelloA") {
		t.Error()
	}
}

func TestCurrentTimeInHttpFormat(t *testing.T) {
	loc, _ := time.LoadLocation("GMT")
	someTime := time.Date(2023, 1, 1, 0, 7, 9, 0, loc)
	var timeInHttpFormat string = util.TimeInHttpFormat(someTime)
	if timeInHttpFormat != "Sun, 1 Jan 2023 00:07:09 GMT" {
		t.Error(timeInHttpFormat)
	}
}

func TestNumberToTwoDigitString(t *testing.T) {
	result := util.NumberToTwoDigitString(8)
	if result != "08" {
		t.Error()
	}

	result = util.NumberToTwoDigitString(10)
	if result != "10" {
		t.Error()
	}

	result = util.NumberToTwoDigitString(84)
	if result != "84" {
		t.Error()
	}

	result = util.NumberToTwoDigitString(0)
	if result != "00" {
		t.Error()
	}
}

func TestHashString(t *testing.T) {
	if util.HashString("Bongo") != "d60701e61d47abf91cad843dfc3cf97a78fcabd3bd3f1299ff29e4eab7bc1b26" {
		t.Error()
	}
}

func TestUsernameRegex(t *testing.T) {
	match, _ := regexp.MatchString(util.USERNAME_REGEX, "你好好好")
	if match {
		t.Error()
	}
}

func TestUsernameRegex2(t *testing.T) {
	match, _ := regexp.MatchString(util.USERNAME_REGEX, "11")
	if match {
		t.Error()
	}
}

func TestUsernameRegex3(t *testing.T) {
	match, _ := regexp.MatchString(util.USERNAME_REGEX, "11111111111111111")
	if match {
		t.Error()
	}
}

func TestEmailRegex(t *testing.T) {
	match, _ := regexp.MatchString(util.EMAIL_REGEX, "m1smr@hotmail.com")
	if !match {
		t.Error()
	}
}

func TestWasXHoursAgo(t *testing.T) {
	now := time.Now()
	nowMinus11h := time.UnixMilli(now.UnixMilli() - 1000*60*60*11)

	result := util.WasXHoursAgoUTC(nowMinus11h, 10)
	if result == false {
		t.Error()
	}
}

func TestWasXHoursAgo2(t *testing.T) {
	now := time.Now()
	nowMinus9h := time.UnixMilli(now.UnixMilli() - 1000*60*60*9)

	result := util.WasXHoursAgoUTC(nowMinus9h, 10)
	if result == true {
		t.Error()
	}
}
