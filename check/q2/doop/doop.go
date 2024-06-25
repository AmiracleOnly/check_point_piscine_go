package main

import "os"

func main() {
	if len(os.Args) != 4 {
		return
	}

	num1, ok1 := atoi(os.Args[1])
	operator := os.Args[2]
	num2, ok2 := atoi(os.Args[3])

	if !ok1 || !ok2 {
		return
	}

	if num1+num2 > 1<<31-1 || num1+num2 < -1<<31 {
		return
	}

	switch operator {
	case "+":
		sum := (num1 + num2)
		result := itoa(sum)
		os.Stdout.WriteString(result)
		os.Stdout.WriteString("\n")
	case "-":
		sum := (num1 - num2)
		result := itoa(sum)
		os.Stdout.WriteString(result + "\n")
	case "*":
		sum := (num1 * num2)
		result := itoa(sum)
		os.Stdout.WriteString(result + "\n")
	case "/":
		if num2 == 0 {
			os.Stdout.WriteString("No Division by 0\n")
			os.Stdout.WriteString("\n")
			return
		}
		sum := (num1 / num2)
		result := itoa(sum)
		os.Stdout.WriteString(result + "\n")
	case "%":
		if num2 == 0 {
			os.Stdout.WriteString("No Division by 0\n")
			return
		}
		sum := (num1 % num2)
		result := itoa(sum)
		os.Stdout.WriteString(result + "\n")
	default:
		return
	}

}

func atoi(s string) (int64, bool) {
	var res int64
	var sign int64 = 1

	for i, char := range s {
		if i == 0 && char == '-' || char == '+' {
			if char == '-' {
				sign = -1
			}
		} else if char >= '0' && char <= '9' {
			res = res*10 + int64(char-'0')
		} else if char >= 'a' && char <= 'z' {
			return 0, false
		}
	}
	return res * sign, true
}

func itoa(num int64) string {
	sign := ""
	result := ""
	if num == 0 {
		return "0"
	}

	if num < 0 {
		sign = "-"
		num *= -1
	}

	for num > 0 {
		result = string(rune(num%10)+'0') + result
		num /= 10
	}

	return sign + result
}
