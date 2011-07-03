package utils

import "bufio"
import "fmt"
import "http"
import "os"
import "rand"
import "strings"
import "time"

const VERSION = "0.2"
const AUTHOR = "Seth Hoenig"

func SecsToNSecs(seconds int64) int64 {
	return seconds * 1000000000
}

// rand rand int [low, high)
func RandInt(low, high int) int {
	dv := high - low
	i := rand.Int() % dv
	return i + low
}

// returns HH:MM
func SimpleTime() string {
	t := time.LocalTime()
	return fmt.Sprintf(" (%v:%v)", fix(int64(t.Hour)), fix(int64(t.Minute)))
}

func fix(h int64) string {
	s := fmt.Sprintf("%v", h)
	if h < 10 {
		s = "0" + s
	}
	return s
}

func asLines(r *bufio.Reader) []string {
	var content []string
	for {
		line, _, err := r.ReadLine()
		if err == os.EOF {
			break
		} else if err != nil {
			panic("Error in asLines")
		} else {
			content = append(content, string(line))
		}
	}
	return content
}

func GetWeather(zipcode string) string {
	var c http.Client
	r, _, herr := c.Get("http://www.weather.com/weather/today/" + zipcode)
	if herr != nil {
		fmt.Printf("Error getting weather data: %v\n", herr)
		return ""
	}
	if r.StatusCode != 200 {
		fmt.Printf("Error with weather report, status code: %v\n", r.Status)
		return ""
	}

	lines := asLines(bufio.NewReader(r.Body))
	temp := ""
	sky := ""
	city := ""
	next := false
	for _, line := range lines {
		if next {
			sky = strings.TrimSpace(line)
			next = false
			continue
		}
		if strings.Contains(line, "realTemp:") {
			temp = line[strings.IndexAny(line, "\"")+1 : strings.LastIndex(line, "\"")]
		} else if strings.Contains(line, "<td class=\"twc-col-1\">") {
			next = true
		} else if strings.Contains(line, "locName:") {
			fmt.Printf("HERE, %s\n", line)
			city = line[strings.IndexAny(line, "\"")+1 : strings.LastIndex(line, "\"")]
		}
	}
	if temp == "" {
		fmt.Printf("Error in weather report, couldn't get temp")
	}
	if sky == "" {
		fmt.Printf("Error in weather report, did not get sky")
	}
	if city == "" {
		fmt.Printf("Error in weather report, did not get city")
	}
	return city + ", " + temp + "° F, " + sky
}

func Wikify(term string) string {
	term = strings.Replace(term, "!wiki", "", -1)
	term = strings.TrimSpace(term)
	splitted := strings.Fields(term)
	if len(splitted) == 0 {
		return "usage: !wiki <term>"
	}
	term = strings.Replace(term, " ", "_", -1)
	return fmt.Sprintf("http://en.wikipedia.org/wiki/%s", term)
}

func SecsToTime(seconds int64) string {
	secsPday := int64(60 * 60 * 24)
	days := seconds / secsPday

	seconds %= secsPday
	secsPhour := int64(60 * 60)
	hours := seconds / secsPhour
	
	seconds %= secsPhour
	secsPminute := int64(60)
	minutes := seconds / secsPminute

	str := ""
	if days == 1 {
		str = fmt.Sprintf("up %d day, %d:%s", days, hours, fix(minutes))
	} else {
		str = fmt.Sprintf("up %d days, %d:%s", days, hours, fix(minutes))
	}
	return str
}

// shortener for strings.Contains
func Scon(a, b string) bool {
	return strings.Contains(a, b)
}