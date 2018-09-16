// Source : https://oj.leetcode.com/problems/longest-common-prefix/
// Author : Lei Cao
// Date   : 2018-09-01

/**********************************************************************************
*
* Write a function to find the longest common prefix string amongst an array of strings.
*
* If there is no common prefix, return an empty string "".
*
* Example 1:
*
* Input: ["flower","flow","flight"]
* Output: "fl"
* Example 2:
*
* Input: ["dog","racecar","car"]
* Output: ""
* Explanation: There is no common prefix among the input strings.
* Note:
*
* All given inputs are in lowercase letters a-z.
*
**********************************************************************************/

package main

func longestCommonPrefix(strs []string) string {
	common := ""
	if len(strs) == 0 {
		return common
	}
	for i := 0; i < len(strs[0]); i++ {
		currentChar := ""
		for _, v := range strs {
			if len(v) <= i {
				return common
			}
			if currentChar == "" {
				currentChar = string(v[i])
			}
			if currentChar != string(v[i]) {
				return common
			}
		}
		common += currentChar
	}

	return common
}
