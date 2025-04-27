package kata

func ReversesWord(str string) string {
	var res string
	var word string

	for _, i := range str {
		if i == ' ' {
			res = res + word + " "
			word = ""
		} else {
			word = string(i) + word
		}
	}

	return res + word
}

// golang 資料型別

// int  			32bit, 64bit
// int8 			8bits
// int16 			16bits
// int32			32bits
// int64 			64bits
// uint, uint8 ... 	無符號整數
// float32          6位有效
// float64        	15位有效
// string			[]byte
// rune             int32

// for _, i := range string
// i 為 rune 型別
// for i := 0; i < len(str) ; i++
// str[i] 為 byte 型別
