package main

func intToRoman(num int) string {
	ones := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	tens := []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	hundreds := []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	thous := []string{"", "m", "MM", "MMM"}
	ans := thous[num/1000] + hundreds[(num%100)/100] + tens[(num%100)/10] + ones[num%10]
	return ans
}

func main() {}
