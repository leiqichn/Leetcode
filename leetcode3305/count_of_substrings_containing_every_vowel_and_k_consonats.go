package leetcode3305

func countOfSubstrings(word string, k int) int {
	vowels := map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}
	n := len(word)
	res := 0
	for i := 0; i < n; i++ {
		occur := map[byte]bool{}
		consonants := 0
		for j := i; j < n; j++ {
			if vowels[word[j]] {
				occur[word[j]] = true
			} else {
				consonants++
			}
			if len(occur) == 5 && consonants == k {
				res++
			}
		}
	}
	return res
}
