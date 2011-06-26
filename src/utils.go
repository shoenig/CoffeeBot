package utils

import "bufio"
import "fmt"
import "http"
import "os"
import "rand"
import "strings"
import "time"

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
	return fmt.Sprintf(" (%v:%v)", fix(t.Hour), fix(t.Minute))
}

func fix(h int) string {
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

func GetWeather() string {
	var c http.Client
	r, _, herr := c.Get("http://www.weather.com/weather/today/USCA0987") //SF hardcoded
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
	next := false
	for _, line := range lines {
		if next {
			sky = strings.TrimSpace(line)
			next = false
			continue
		}
		if strings.Contains(line, "realTemp") {
			temp = line[strings.IndexAny(line, "\"")+1 : strings.LastIndex(line, "\"")]
		} else if strings.Contains(line, "<td class=\"twc-col-1\">") {
			next = true
		}
	}
	if temp == "" {
		fmt.Printf("Error in weather report, couldn't get temp")
	}
	if sky == "" {
		fmt.Printf("Error in weather report, did not get sky")
	}
	return "SanFran, " + temp + "° F, " + sky
}
