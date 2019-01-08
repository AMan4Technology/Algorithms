package IntegerToEnglishWords

import (
	"strings"

	. "Algorithms/LeetCode/internal/box/math"
)

var wordWithInt = map[int]string{
	1: "One", 2: "Two", 3: "Three", 4: "Four", 5: "Five", 6: "Six", 7: "Seven", 8: "Eight", 9: "Nine", 10: "Ten",
	11: "Eleven", 12: "Twelve", 13: "Thirteen", 14: "Fourteen", 15: "Fifteen", 16: "Sixteen", 17: "Seventeen", 18: "Eighteen", 19: "Nineteen", 20: "Twenty",
	30: "Thirty", 40: "Forty", 50: "Fifty", 60: "Sixty", 70: "Seventy", 80: "Eighty", 90: "Ninety", 100: "Hundred", 1000: "Thousand", 1000000: "Million", 1000000000: "Billion",
}

func numberToWords(num int) (result string) {
	if num == 0 {
		return "Zero"
	}
	for i := 0; num != 0; num /= 1000 {
		if num%1000 != 0 {
			result = parseThousand(num%1000) + " " + wordWithInt[i] + " " + result
		}
		i = MaxOfTwo(i, 1) * 1000
	}
	return strings.TrimSpace(result)
}

func parseThousand(num int) (result string) {
	var (
		hundred = num / 100
		other   = num % 100
	)
	if hundred != 0 {
		result = wordWithInt[hundred] + " " + wordWithInt[100] + " "
	}
	if other <= 20 {
		result += wordWithInt[other]
	} else {
		result += wordWithInt[other/10*10] + " " + wordWithInt[other%10]
	}
	return strings.TrimSpace(result)
}
