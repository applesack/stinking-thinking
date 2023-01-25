package p187

func findRepeatedDnaSequences(s string) []string {
	if len(s) <= 10 {
		return []string{}
	}
	mem := make(map[string]int)
	for i := 10; i <= len(s); i++ {
		mem[s[(i-10):i]]++
	}
	ans := make([]string, 0, 8)
	for k, v := range mem {
		if v > 1 {
			ans = append(ans, k)
		}
	}
	return ans
}
