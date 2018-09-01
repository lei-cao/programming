// Source : https://oj.leetcode.com/problems/longest-substring-without-repeating-characters/
// Author : Lei Cao
// Date   : 2018-08-31

/********************************************************************************** 
* 
* Given a string, find the length of the longest substring without repeating characters. 
* For example, the longest substring without repeating letters for "abcabcbb" is "abc", 
* which the length is 3. For "bbbbb" the longest substring is "b", with the length of 1.
*               
**********************************************************************************/

package main

func main() {

}

// brutal force
func lengthOfLongestSubstring(s string) int {
	var longest int
	for i := 0; i < len(s); i++ {
		n := 1
		for j := i + 1; j < len(s); j++ {
			repeated := false
			for k := i; k < j; k++ {
				if string(s[j]) == string(s[k]) {
					repeated = true
					break
				}
			}
			if repeated {
				break
			}
			n++

		}
		if longest < n {
			longest = n
		}

	}
	return longest
}

// map
func lengthOfLongestSubstringMap(s string) int {
	var longest int
	start := 0
	lastCharMap := map[uint8]int{}
	currentLength := 0
	for i := 0; i < len(s); i++ {
		if lastCharMap[s[i]] > 0 && lastCharMap[s[i]] >= start+1 {
			// compare currentLength with the longest
			if currentLength > longest {
				longest = currentLength
			}

			// update start from the last occurrence
			start = lastCharMap[s[i]]
			currentLength = i - start+1
		} else {
			currentLength ++
		}
		// update the last occurrence index, avoid default 0 value
		lastCharMap[s[i]] = i + 1
	}
	if currentLength > longest {
		longest = currentLength
	}

	return longest
}


