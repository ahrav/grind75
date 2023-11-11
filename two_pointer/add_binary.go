package two_pointer

import (
	"strconv"
	"strings"
)

func addBinary(str1, str2 string) string {
	if str1 == "" {
		return str2
	}

	if str2 == "" {
		return str1
	}

	i, j := len(str1)-1, len(str2)-1
	var result []byte
	carry := 0

	for i >= 0 || j >= 0 || carry > 0 {
		sum := carry
		if i >= 0 {
			sum += int(str1[i] - '0')
		}

		if j >= 0 {
			sum += int(str2[j] - '0')
		}

		// Append '0' or '1' as a character, not as a byte value.
		result = append(result, byte(sum%2+'0'))
		carry = sum / 2
		i--
		j--
	}

	return reverse(result)
}

func reverse(str []byte) string {
	left, right := 0, len(str)-1
	for left < right {
		str[left], str[right] = str[right], str[left]
		left++
		right--
	}

	return string(str)
}

func addBinarySlower(str1, str2 string) string {
	if str1 == "" {
		return str2
	}

	if str2 == "" {
		return str1
	}

	carry := 0
	var result strings.Builder
	for i, j := len(str1)-1, len(str2)-1; i >= 0 || j >= 0 || carry > 0; i, j = i-1, j-1 {
		sum := carry
		if i >= 0 {
			sum += int(str1[i] - '0')
		}

		if j >= 0 {
			sum += int(str2[j] - '0')
		}

		result.WriteString(strconv.Itoa(sum % 2))
		carry = sum / 2
	}

	return reverseSlower(result.String())
}

func reverseSlower(str string) string {
	var result strings.Builder
	for i := len(str) - 1; i >= 0; i-- {
		result.WriteByte(str[i])
	}

	return result.String()
}
