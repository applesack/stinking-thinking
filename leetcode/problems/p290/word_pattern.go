package p290

import "strings"

func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")
	if len(words) != len(pattern) {
		return false
	}

	word2ch := make(map[string]byte)
	ch2word := make(map[byte]string)
	for i := range words {
		ch, word := pattern[i], words[i]
		// 如果存了单词到符号的映射,那么单词到符号的映射一定等于当前符号
		// 如果存了符号到单词的映射,那么符号到单词的映射一定等于当前单词
		if w, ok := ch2word[ch]; ok {
			if w != word {
				return false
			}
		}
		if c, ok := word2ch[word]; ok {
			if c != ch {
				return false
			}
		}
		// 更新单词和符号的映射
		ch2word[ch] = word
		word2ch[word] = ch
	}
	return true
}
