# leetcode-go

This repository contains my solutions to LeetCode problems. I've been actively practicing to enhance my problem-solving skills and algorithmic knowledge.

<a href="https://leetcode.com/nongbritee">
  <img alt="leetcode" src="https://img.shields.io/badge/Leetcode-orange?style=for-the-badge&logo=leetcode&logoColor=black"/>
</a>
<a href="https://www.linkedin.com/in/napong-jantaranimi-273983171">
  <img alt="linkedin" src="https://img.shields.io/badge/LinkedIn-0077B5?style=for-the-badge&logo=linkedin&logoColor=white"/> 
 </a> 
<a href="mailto:nongbriteenapong@gmail.com">
  <img alt="gmail" src="https://img.shields.io/badge/Gmail-D14836?style=for-the-badge&logo=gmail&logoColor=white"/>
</a>
<br/>
<br/>
<a href="https://leetcode.com/nongbritee">
  <img alt="leetcode_done_over_total" src="https://img.shields.io/badge/dynamic/json?style=for-the-badge&labelColor=black&color=%23ffa116&label=Solved&query=solved&url=https%3A%2F%2Fleetcode-badge.vercel.app%2Fapi%2Fusers%2Fnongbritee&logo=leetcode&logoColor=yellow"/>
</a>
<br/>

![](https://badges.peiyuan.ch/leetcode/nongbritee/ranking)
![](https://badges.peiyuan.ch/leetcode/nongbritee/solved?difficulty=easy)
![](https://badges.peiyuan.ch/leetcode/nongbritee/solved?difficulty=medium)
![](https://badges.peiyuan.ch/leetcode/nongbritee/solved?difficulty=hard)


| No.  | Title                                               | Solution                                                                                                                                             |     Difficulty      | Note                                                                                                                                                                                                                                                                |
|:----:|:----------------------------------------------------|:-----------------------------------------------------------------------------------------------------------------------------------------------------|:-------------------:|:--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| 0001 | Two Sum                                             | [Go](https://github.com/NongBritee/leetcode-go/blob/main/leetcode/0001.Two%20Sum.go) [Array]                                                         |       Easy         | <sub>1. Loop through array as 2d and find the sum of target<br/>2. Return the founded index pair</sub>                                                                                                                                                              |
| 0026 | Remove Duplicates from Sorted Array                 | [Go](https://github.com/NongBritee/leetcode-go/blob/main/leetcode/0026.Remove%20Duplicates%20from%20Sorted%20Array.go) [Array]                       |        Easy	        | <sub>1. loop through nums array then append to new array when it's not duplicate to the latest value in new array</sub>                                                                                                                                             |
| 0027 | Remove Element                                      | [Go](https://github.com/NongBritee/leetcode-go/blob/main/leetcode/0027.Remove%20Element.go) [Array]                                                  |        Easy         | <sub>1. loop through array and replace wanted number to 999<br/> 2. Sort array</sub>                                                                                                                                                                                |
| 0048 | Rotate Image                                        | [Go](https://github.com/NongBritee/leetcode-go/blob/main/leetcode/0048.Rotate%20Image.go) [Array, Transpose, Swap]                                   | Medium (open cheat) | <sub>1. Transpose matrix<br/> 2. Swap left and right column</sub>                                                                                                                                                                                                   |
| 0066 | Plus One | [Go](https://github.com/NongBritee/leetcode-go/blob/main/leetcode/0066.Plus%20One.go) [Array]                                                        | Easy | <sub>1. Loop through array from last index to first index<br/> 2. Add 1 to the last index<br/> 3. If the value is 10, set the value to 0 and add 1 to the previous index<br/> 4. If the first index is 10, set the value to 0 and add 1 to the new index<br/></sub> |
| 0136 | Single Number                                       | [Go](https://github.com/NongBritee/leetcode-go/blob/main/leetcode/0136.Single%20Number.go) [Map]                                                     |        Easy         | <sub>1. Use map to count the number of each element<br/> 2. Loop through map and return the key that has value = 1</sub>                                                                                                                                            |
| 0137 | Single Number II                                    | [Go](https://github.com/NongBritee/leetcode-go/blob/main/leetcode/0137.Single%20Number%20II.go) [Map]                                                |       Medium        | <sub>1. Use map to count the number of each element<br/> 2. Loop through map and return the key that has value = 1</sub>                                                                                                                                            |
| 0189 | Rotate Array                                        | [Go](https://github.com/NongBritee/leetcode-go/blob/main/leetcode/0189.Rotate%20Array.go) [Array, Modular]                                           |       Medium        | <sub>1. Split array into 2 array by k.<br/> 2. Concat it back new order (2nd array + 1st array).<br/> 3. modify k % len(nums) to avoid bug when k > len(nums)</sub>                                                                                                 |
| 0217 | Contains Duplicate                                  | [Go](https://github.com/NongBritee/leetcode-go/blob/main/leetcode/0217.Contains%20Duplicate.go) [Array, Map]                                         |        Easy         | <sub>1. Use map to check the duplicate number</sub>                                                                                                                                                                                                                 |
| 0260 | Single Number III                                   | [Go](https://github.com/NongBritee/leetcode-go/blob/main/leetcode/0137.Single%20Number%20III.go) [Map]                                               |       Medium        | <sub>1. Use map to count the number of each element<br/> 2. Loop through map, push key to new array if val == 1</sub>                                                                                                                                               |
| 0283 | Move Zeroes                                         | [Go]() [Array, 2 Pointer]                                                                                                                            |       Easy         | <sub>1. create to 2 index pointer<br/> 2. 1st pointer running for find zero index<br/>3. 2nd pointer running to find non-zero index<br/>4. swap data when 2nd pointer is greater than 1st pointer<br/>5. 2nd pointer is always +1 at the end of loop</sub>          |
| 0383 | Ransoms Note                                        | [Go](https://github.com/NongBritee/leetcode-go/blob/main/leetcode/0383.Ransom%20Note.go)  [Array, Map]                                               |        Easy         | <sub>1. Map magazine [char, count] </br>2. Loop through ransoms note and check char from Magazine map is enough for ransoms note by decrease -1 every time that found char in map<br/></sub>                                                                        |
| 0412 | Fizz Buzz                                           | [Go](https://github.com/NongBritee/leetcode-go/blob/main/leetcode/0412.Fizz%20Buzz.go) [Array, Modular]                                              |        Easy         | <sub>1. Loop through array and push to new array when mod by 3, 5</sub>                                                                                                                                                                                             |
| 1342 | Number of Steps to Reduce a Number to Zero          | [Go](https://github.com/NongBritee/leetcode-go/blob/main/leetcode/1342.Number%20of%20Steps%20to%20Reduce%20a%20Number%20to%20Zero.go) [Math]         |        Easy         | <sub>1. Loop until result == 0<br/> 2. divided by 2 and -1 when result is odd number<br/> 3. Count every step of devide and minus operation</sub>                                                                                                                   |
| 1347 | Minimum Number of Steps to Make Two Strings Anagram | [Go](https://github.com/NongBritee/leetcode-go/blob/main/leetcode/1347.Minimum%20Number%20of%20Steps%20to%20Make%20Two%20Strings%20Anagram.go) [Map] |       Medium        | <sub>1. create sMap of s string [char, count] by count the no. of char<br/> 2. loop through t string and find char in sMap<br/> 3. -1 from count when found (stop minus when 0)<br/> 4. return sum of remaining count in the sMap</sub>                             |
| 1480 | Running Sum of 1d Array                             | [Go](https://github.com/NongBritee/leetcode-go/blob/main/leetcode/1480.Running%20Sum%20of%201d%20Array.go) [Array]                                   |        Easy         | <sub>1. Loop through array start from 2nd index and add previous value to current value</sub>                                                                                                                                                                       |
| 1672 | Richest Customer Wealth                             | [Go](https://github.com/NongBritee/leetcode-go/blob/main/leetcode/1672.Richest%20Customer%20Wealth.go) [Array]                                       |        Easy         | <sub>1. Loop through array and sum each array then compare to max value</sub>                                                                                                                                                                                       |
|      |                                                     |                                                                                                                                                      |                     |                                                                                                                                                                                                                                                                     |
