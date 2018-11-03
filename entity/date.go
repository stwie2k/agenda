package entity

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	// "agenda/loghelper"
)

// Date : class with y, m, d, h and m
type Date struct {
	Year, Month, Day, Hour, Minute int
}

func (mDate Date) init(tYear, tMonth, tDay, tHour, tMinute int) {
	mDate.Year = tYear
	mDate.Month = tMonth
	mDate.Day = tDay
	mDate.Hour = tHour
	mDate.Minute = tMinute
}

// GetYear : getter of year
func (mDate Date) GetYear() int {
	return mDate.Year
}

// SetYear : setter of year
func (mDate Date) SetYear(tYear int) {
	mDate.Year = tYear
}

// GetMonth : getter of month
func (mDate Date) GetMonth() int {
	return mDate.Month
}

// SetMonth : setter of month
func (mDate Date) SetMonth(tMonth int) {
	mDate.Month = tMonth
}

// GetDay : getter of day
func (mDate Date) GetDay() int {
	return mDate.Day
}

// SetDay : setter of day
func (mDate Date) SetDay(tDay int) {
	mDate.Day = tDay
}

// GetHour : getter of Hour
func (mDate Date) GetHour() int {
	return mDate.Hour
}

// SetHour : setter of Hour
func (mDate Date) SetHour(tHour int) {
	mDate.Hour = tHour
}

// GetMinute : getter of minute
func (mDate Date) GetMinute() int {
	return mDate.Minute
}

// SetMinute : setter of minute
func (mDate Date) SetMinute(tMinute int) {
	mDate.Minute = tMinute
}

/**
*   @brief check whether the date is valid or not
*   @return the bool indicate valid or not
 */
func IsValid(tDate Date) bool {
	var dayOfMonths = [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	var currentYear = tDate.GetYear()
	var currentMonth = tDate.GetMonth()
	var currentDay = tDate.GetDay()
	//Check if date in normal range
	if currentYear < 1000 || currentYear > 9999 || currentMonth < 1 ||
		currentMonth > 12 || currentDay < 1 || tDate.GetHour() < 0 ||
		tDate.GetHour() >= 24 || tDate.GetMinute() < 0 ||
		tDate.GetMinute() >= 60 {
		return false
	}

	if currentMonth != 2 && currentDay > dayOfMonths[currentMonth-1] {
		return false
	} else {
		//Check if leap year.
		if (currentYear%4 == 0 && currentYear%100 != 0) ||
			(currentYear%400 == 0) {
			if currentDay > 29 {
				return false
			}
		} else {
			if currentDay > 28 {
				return false
			}
		}
	}
	defer fmt.Println("Error: Invalid Date.")
	return true
}

/**
* @brief convert string to int
 */
func String2Int(s string) int {
	result, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}
	return result
}

/**
* @brief convert a string to date, if the format is not correct return
* 0000-00-00/00:00
* @return a date
 */
func StringToDate(tDateString string) (Date, error) {
	var datePattern = `[\d]{4}-[\d]{2}-[\d]{2}\/[\d]{2}:[\d]{2}`
	var resultDate Date

	//Use regexp to check if parttern matched

	isValid, _ := regexp.Match(datePattern, []byte(tDateString))

	if isValid != true {
		return resultDate, errors.New("Invalid string format")
	}

	resultDate.Year = String2Int(tDateString[0:4])
	resultDate.Month = String2Int(tDateString[5:7])
	resultDate.Day = String2Int(tDateString[8:10])
	resultDate.Hour = String2Int(tDateString[11:13])
	resultDate.Minute = String2Int(tDateString[14:])
	return resultDate, nil
}

/**
*   @brief convert the date to string, if result length is 1, add padding 0
 */
func Int2String(a int) string {
	var resultString string
	resultString = strconv.Itoa(a)
	return resultString
}

/**
* @brief convert a date to string, if the format is not correct return
* 0000-00-00/00:00
 */
func DateToString(tDate Date) (string, error) {
	var dateString = ""
	var wString = ""
	var initTime = "0000-00-00/00:00"
	if !IsValid(tDate) {
		dateString = initTime
		return dateString, nil
	}
	// dateString = Int2String(tDate.GetYear()) + "-" + Int2String(tDate.GetMonth()) +
	// 	"-" + Int2String(tDate.GetDay()) + "/" + Int2String(tDate.GetHour()) +
	// 	":" + Int2String(tDate.GetMinute())

	dateStringWithZero := fmt.Sprintf("%04d", tDate.GetYear()) + "-" + fmt.Sprintf("%02d", tDate.GetMonth()) +
		"-" + fmt.Sprintf("%02d", tDate.GetDay()) + "/" + fmt.Sprintf("%02d", tDate.GetHour()) +
		":" + fmt.Sprintf("%02d", tDate.GetMinute())

	if dateStringWithZero != wString {
		return dateStringWithZero, nil
	}
	return dateStringWithZero, errors.New("wrong")

}

/**
*  @brief overload the assign operator
 */
func (mDate Date) CopyDate(tDate Date) Date {
	mDate.SetYear(tDate.GetYear())
	mDate.SetMonth(tDate.GetMonth())
	mDate.SetDay(tDate.GetDay())
	mDate.SetHour(tDate.GetHour())
	mDate.SetMinute(tDate.GetMinute())
	return mDate
}

/**
* @brief check whether the CurrentDate is equal to the tDate
 */
func (mDate Date) IsSameDate(tDate Date) bool {
	return (tDate.GetYear() == mDate.GetYear() &&
		tDate.GetMonth() == mDate.GetMonth() &&
		tDate.GetDay() == mDate.GetDay() &&
		tDate.GetHour() == mDate.GetHour() &&
		tDate.GetMinute() == mDate.GetMinute())
}

/**
* @brief check whether the CurrentDate is  greater than the tDate
 */
func (mDate Date) MoreThan(tDate Date) bool {
	if mDate.Year > tDate.GetYear() {
		return true
	}
	if mDate.Year < tDate.GetYear() {
		return false
	}
	if mDate.Month > tDate.GetMonth() {
		return true
	}
	if mDate.Month < tDate.GetMonth() {
		return false
	}
	if mDate.Day > tDate.GetDay() {
		return true
	}
	if mDate.Day < tDate.GetDay() {
		return false
	}
	if mDate.Hour > tDate.GetHour() {
		return true
	}
	if mDate.Hour < tDate.GetHour() {
		return false
	}
	if mDate.Minute > tDate.GetMinute() {
		return true
	}
	if mDate.Minute < tDate.GetMinute() {
		return false
	}
	return false
}

// LessThan : ommited.
func (mDate Date) LessThan(tDate Date) bool {
	if mDate.IsSameDate(tDate) == false && mDate.MoreThan(tDate) == false {
		return true
	}
	return false
}

/**
* @brief check whether the CurrentDate is  greater or equal than the
* tDate
 */
func (mDate Date) GreateOrEqual(tDate Date) bool {
	return mDate.IsSameDate(tDate) || mDate.MoreThan(tDate)
}

/**
* @brief check whether the CurrentDate is  less than or equal to the
* tDate
 */
func (mDate Date) LessOrEqual(tDate Date) bool {
	return !mDate.MoreThan(tDate)
}
