package Media

import (
	"strconv"
	"strings"
	"time"
)


// get current time

func LeftPad2Len(s string, padStr string, overallLen int) string {
	var padCountInt = 1 + ((overallLen - len(padStr)) / len(padStr))
	var retStr = strings.Repeat(padStr, padCountInt) + s
	return retStr[(len(retStr) - overallLen):]
}

// get current time

func GetYear() string{
	currentTime := time.Now()
	result:=LeftPad2Len(strconv.Itoa(int(currentTime.Year())), "0", 4)
	return result
}

func GetMonth() string{
	currentTime := time.Now()
	result:=LeftPad2Len(strconv.Itoa(int(currentTime.Month())), "0", 2)
	return result
}

func GetDay() string{
	currentTime := time.Now()
	result:=LeftPad2Len(strconv.Itoa(int(currentTime.Day())), "0", 2)
	return result
}

func GetHour() string{
	currentTime := time.Now()
	result := strconv.Itoa(currentTime.Hour())
	return result
}

func GetMinute() string{
	currentTime := time.Now()
	result := strconv.Itoa(currentTime.Minute())
	return result
}

func GetSecond() string{
	currentTime := time.Now()
	result := strconv.Itoa(currentTime.Second())
	return result
}
