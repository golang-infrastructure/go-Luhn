package luhn

// Check 检查银行卡号是否合法
func Check(cardNumber string) bool {
	sum := 0
	for index := len(cardNumber) - 1; index >= 0; index-- {
		if !(cardNumber[index] >= '0' && cardNumber[index] <= '9') {
			return false
		}
		t := int(cardNumber[index] - '0')
		if (len(cardNumber)-index)%2 == 0 {
			// 偶数位
			n := t * 2
			for n > 0 {
				sum += n % 10
				n /= 10
			}
		} else {
			// 奇数位
			sum += t
		}
	}
	return sum%10 == 0
}
