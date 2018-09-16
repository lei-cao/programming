// Source : https://oj.leetcode.com/problems/reverse-integer/
// Author : Lei Cao
// Date   : 2018-09-01

/********************************************************************************** 
* 
* Given a 32-bit signed integer, reverse digits of an integer.
*
* Example1: x =  123, return  321
* Example2: x = -123, return -321
* Example3: x =  120, return  21
*
* 
* Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−231,  231 − 1].
* For the purpose of this problem, assume that your function returns 0 when the reversed integer overflows.
*
* Have you thought about this?
*
* Here are some good questions to ask before coding. Bonus points for you if you have already thought through this!
* 
* > If the integer's last digit is 0, what should the output be? ie, cases such as 10, 100.
* 
* > Did you notice that the reversed integer might overflow? Assume the input is a 32-bit integer, 
*   then the reverse of 1000000003 overflows. How should you handle such cases?
* 
* > Throw an exception? Good, but what if throwing an exception is not an option? 
*   You would then have to re-design the function (ie, add an extra parameter).
* 
*               
**********************************************************************************/

package main

import (
	"math"
)

func reverse(x int) int {
	max := int(math.Pow(2, 31)) - 1
	min := -int(math.Pow(2, 31))

	reversed := 0
	digits := digits(x)
	for i := 0; i < len(digits); i++ {
		reversed += digits[i] * int(math.Pow(10, float64(len(digits)-i-1)))
	}
	if reversed < min || reversed > max {
		reversed = 0
	}

	return reversed
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
