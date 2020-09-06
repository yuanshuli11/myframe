package utils

import "strconv"

func StringTofloat64(str string) (float64, error) {

	num, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0.00, err
	}
	return num, nil
}
func Float64Round(floatNum float64, accuracy int) float64 {
	stringNum := strconv.FormatFloat(floatNum, 'f', accuracy, 64)
	newFloatNum, err := strconv.ParseFloat(stringNum, 64)
	if err != nil {
		return floatNum
	}
	return newFloatNum
}

func Float64ToUint(floatNum float64) (uint, error) {
	if floatNum <= 0 {
		return uint(0), nil
	}
	str := strconv.FormatFloat(floatNum, 'f', 0, 64)
	returnInt, err := strconv.Atoi(str)
	if err != nil {
		return uint(0), err
	}
	return uint(returnInt), err
}
