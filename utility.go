package util

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const ADMIN_TOKEN string = "OdHtwfBpaXM08BcbB6oCiiacD9jOgp0z4uurp3g4Zr3F7gWhjb13pdh5vGYAjSM8"
const USERNAME_REGEX string = "^[a-zA-Z0-9\u00E4\u00F6\u00FC\u00DC\u00C4\u00D6_.-]{3,16}$"
const EMAIL_REGEX string = "^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+.[a-zA-Z]{2,}$"

func Contains(s, substr string) bool {
	for i := 0; i < len(s); i++ {
		if HasPrefix(s[i:], substr) {
			return true
		}
	}
	return false
}

// Returns -1 if the string does not contain the substring
func Find(s, substr string) int {
	for i := 0; i < len(s); i++ {
		if HasPrefix(s[i:], substr) {
			return i
		}
	}
	return -1
}

func FindLast(s, substr string) int {
	for i := len(s) - 1; i >= 0; i-- {
		if HasPrefix(s[i:], substr) {
			return i
		}
	}
	return -1
}

func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

func GetHttpRequestBody(req *http.Request) ([]byte, error) {
	fmt.Println("UTIL UPDATE")
	body, err := io.ReadAll(req.Body)
	if err != nil && Contains(err.Error(), "EOF") == false {
		return nil, err
	}
	return body, nil
}

func HasInternet() bool {
	_, err := http.Get("http://google.com")
	return err == nil
}

func TimeInHttpFormat(tm time.Time) string {
	gmtLocation, _ := time.LoadLocation("GMT")
	gmtTime := tm.In(gmtLocation)
	timeInHttpFormat := gmtTime.Weekday().String()[:3] + ", " + fmt.Sprint(gmtTime.Day()) + " " + gmtTime.Month().String()[:3] + " " + fmt.Sprint(gmtTime.Year()) + " " + NumberToTwoDigitString(gmtTime.Hour()) + ":" + NumberToTwoDigitString(gmtTime.Minute()) + ":" + NumberToTwoDigitString(gmtTime.Second()) + " GMT"
	return timeInHttpFormat
}

func NumberToTwoDigitString(num int) string {
	if num < 10 {
		return fmt.Sprintf("0%d", num)
	}
	return fmt.Sprint(num)
}

func HashString(str string) string {
	concatHashed := sha256.Sum256([]byte(str))
	concatHashedHex := fmt.Sprintf("%x", concatHashed)
	return concatHashedHex
}

func UnixMilliToGmtString(unixMs int64) string {
	goTime := time.UnixMilli(unixMs)
	return TimeInHttpFormat(goTime)
}

func Uint64StringToUint64(str string) uint64 {
	ui, _ := strconv.ParseUint(str, 10, 64)
	return ui
}

func SetupTestRequest(body interface{}, sessionToken string) *http.Request {
	bodyJson, _ := json.Marshal(body)
	stringReader := strings.NewReader(string(bodyJson))
	request, _ := http.NewRequest("", "", stringReader)
	if sessionToken != "" {
		cookie := http.Cookie{Name: "sessionToken", Value: sessionToken}
		request.AddCookie(&cookie)
	}
	return request
}

func TimeDiff(earlyTime time.Time, lateTime time.Time) time.Duration {
	return lateTime.Sub(earlyTime)
}

func WasXHoursAgoUTC(t time.Time, hours int) bool {
	now := time.Now().UTC()
	duration, _ := time.ParseDuration(strconv.Itoa(hours) + "h")
	return t.Add(duration).Before(now)
}
