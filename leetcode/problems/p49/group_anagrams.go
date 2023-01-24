package p49

func groupAnagrams(strs []string) [][]string {
	var m = make(map[[26]int][]string)
	var template [26]int
	for _, str := range strs {
		clear(&template)
		zip(str, &template)
		m[template] = append(m[template], str)
	}

	ans := make([][]string, 0, len(m))
	for _, v := range m {
		ans = append(ans, v)
	}
	return ans
}

func zip(str string, template *[26]int) {
	for _, ch := range str {
		(*template)[ch-'a']++
	}
	return
}

func clear(template *[26]int) {
	for i := range template {
		if template[i] > 0 {
			template[i] = 0
		}
	}
}
