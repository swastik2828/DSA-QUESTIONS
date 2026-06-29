func myAtoi(s string) int {
	state := "START"
	sign := 1
	result := 0

	for i := 0; i < len(s); i++ {
		char := s[i]

		switch state {
		case "START":
			if char == ' ' {
				// Stay in START
				continue
			} else if char == '+' || char == '-' {
				state = "SIGNED"
				if char == '-' {
					sign = -1
				}
			} else if char >= '0' && char <= '9' {
				state = "IN_NUMBER"
				result = int(char - '0')
			} else {
				state = "END"
			}
		case "SIGNED":
			if char >= '0' && char <= '9' {
				state = "IN_NUMBER"
				result = int(char - '0')
			} else {
				state = "END"
			}
		case "IN_NUMBER":
			if char >= '0' && char <= '9' {
				digit := int(char - '0')
				
				// Overflow check
				if result > (math.MaxInt32-digit)/10 {
					if sign == 1 {
						return math.MaxInt32
					}
					return math.MinInt32
				}
				result = result*10 + digit
			} else {
				state = "END"
			}
		}

		if state == "END" {
			break
		}
	}

	return sign * result
}