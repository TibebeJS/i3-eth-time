package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func parseTime(body string) string {
	weekDays := []string{"ሰኞ", "ማክሰኞ", "ረቡዕ", "ሐሙስ", "ዓርብ", "ቅዳሜ", "እሑድ"}
	months := []string{"መስከረም", "ጥቅምት", "ህዳር", "ታህሳስ", "ጥር", "የካቲት", "መጋቢት", "ሚያዝያ", "ግንቦት", "ሰኔ", "ሐምሌ", "ነሐሴ", "ጳጉሜ"}

	split := strings.Split(body, "-")

	year := strings.TrimSpace(split[0])
	month, _ := strconv.Atoi(split[1])
	day := split[2]
	hour := split[3]
	minute := split[4]
	dow, _ := strconv.Atoi(split[7])

	return weekDays[dow-1] + "  " + months[month-1] + " " + day + " " + year + " | " + hour + ":" + minute

}

func main() {
	for {

		resp, err := http.Get("https://time.ertale.com/gett.php")

		if err != nil {
			log.Fatalln(err)
		}

		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println(parseTime(string(body)))

		resp.Body.Close()

		time.Sleep(10 * time.Second)
	}
}
