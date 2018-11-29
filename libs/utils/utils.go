package utils

import "strconv"

func ParseInt(int_string string) int{
	if int_string == ""{
		return 1
	}
	int_data, err := strconv.Atoi(int_string)
	if err != nil{
		panic("parse Error")
	}
	return int_data
}