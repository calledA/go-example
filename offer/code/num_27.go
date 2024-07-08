package code

func Num27MethodOne(str string) []string {
	var result []string
	if str == "" {
		return result
	}
	permutation([]byte(str), 0, &result)
	return result
}

func permutation(strByte []byte, i int, result *[]string) {
	length := len(strByte)
	if i != length {
		//利用map排除重复元素
		strMap := make(map[string]int)
		//固定首字符，递归剩余字符；首字符依次与后面交换，继续递归剩余字符
		for j := i; j < length; j++ {
			_, ok := strMap[string(strByte[j])]
			if !ok {
				strMap[string(strByte[j])] = 1
				if j != i {
					strByte[i], strByte[j] = strByte[j], strByte[i]
				}
				permutation(strByte, i+1, result)
				if j != i {
					strByte[i], strByte[j] = strByte[j], strByte[i]
				}
				
			}
		}
	} else {
		*result = append(*result, string(strByte))
	}
}
