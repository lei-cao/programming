// Source : https://oj.leetcode.com/problems/palindrome-number/
// Author : Lei Cao
// Date   : 2018-09-01

/**********************************************************************************
*
* Determine whether an integer is a palindrome. An integer is a palindrome when it reads the same backward as forward.
*
* Example 1:
*	Input: 121
*	Output: true
*
* Example 2:
*	Input: -121
*	Output: false
* 	Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
*
* Example 3:
*	Input: 10
*	Output: false
*	Explanation: Reads 01 from right to left. Therefore it is not a palindrome.
*
* Follow up:
* 	Coud you solve it without converting the integer to a string?
*
* Some hints:
*
* Could negative integers be palindromes? (ie, -1)
*
* If you are thinking of converting the integer to string, note the restriction of using extra space.
*
* You could also try reversing an integer. However, if you have solved the problem "Reverse Integer",
* you know that the reversed integer might overflow. How would you handle such case?
*
* There is a more generic way of solving this problem.
* 
*               
**********************************************************************************/

package main

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	digits := digits(x)
	for i := 0; i < len(digits); i++ {
		if digits[i] != digits[len(digits)-1-i ] {
			return false
		}
	}

	return true
}

func digits(x int) []int {
	digits := []int{}
	var divide = x
	var remains = 0
	for i := 1; i > -1; i++ {
		remains = divide % 10
		divide = divide / 10
		digits = append(digits, remains)
		if divide == 0 {
			break
		}
	}
	return digits
}
