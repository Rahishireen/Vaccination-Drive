package date

import (
	"fmt"
	"math"
	"time"
)
const(
	//Different layout for time 
	apiTimeFormat="2006-01-02 15:04:05"
	timeFormat="01-02-2006"
)

//Get the Current Time
func GetTime() time.Time {
	return time.Now()
}

//Get Current time in String
func GetTimeString() string{
	return GetTime().Format(apiTimeFormat)
}

//String to Time
func ConvertStringtoTime(inputDate string) time.Time{
	date,err:=time.Parse(timeFormat,inputDate)
	fmt.Println(err)
	return date
}

//Time to String
func ConvertTimetoString(inputDate time.Time)string{
	return inputDate.Format(timeFormat)
}

//Check the No of days between 2 dates
func CheckTimeGap(date1 time.Time,date2 time.Time)int{

	//Check the greater Date
	if date1.Before(date2){
		temp:=date2
		date2=date1
		date1=temp
	}
	
	//get the difference in hours
	duration:=date1.Sub(date2)

	//convert into days
	gap:=int(math.Ceil(duration.Hours()/24))

	return gap

}
