package model

import "strconv"

func FizzBuzz(int1 int64, int2 int64, limit int64, str1 string, str2 string) (res string) {
	for i := int64(1); i <= limit; i++ {
		next := strconv.Itoa(int(i))
		if i%int1 == 0 {
			next = str1
		}
		if i%int2 == 0 {
			next = str2
		}
		res = res + next
	}
	return
}
