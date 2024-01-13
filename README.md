<style>
red { color: red }
orange { color: orange }
green { color: green }
</style>

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

### <green>Easy</green> 8/761
### <orange>Medium</orange> 2/1580
### <red>Hard</red> 0/663

|  No.  | Title                                               | Solution                                                                                                                                      |             Difficulty              | Note                                                                                                                                                                                                                                    |
|:-----:|:----------------------------------------------------|:----------------------------------------------------------------------------------------------------------------------------------------------|:-----------------------------------:|:----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| 0026  | Remove Duplicates from Sorted Array                 | [Go](https://github.com/NongBritee/leetcode-go/blob/main/leetcode/0026.Remove%20Duplicates%20from%20Sorted%20Array.go) [Array]                |         <green>Easy</green>         | <sub>1. loop through nums array then append to new array when it's not duplicate to the latest value in new array</sub>                                                                                                                 |
| 0027  | Remove Element                                      | [Go](https://github.com/NongBritee/leetcode-go/blob/main/leetcode/0027.Remove%20Element.go) [Array]                                           |         <green>Easy</green>         | <sub>1. loop through array and replace wanted number to 999<br/> 2. Sort array</sub>                                                                                                                                                    |
| 0189  | Rotate Array                                        | [Go](https://github.com/NongBritee/leetcode-go/blob/main/leetcode/0189.Rotate%20Array.go) [Array, Modular]                                    |               <orange>Medium</orange>                | <sub>1. Split array into 2 array by k.<br/> 2. Concat it back new order (2nd array + 1st array).<br/> 3. modify k % len(nums) to avoid bug when k > len(nums)</sub>                                                                     |
| 0217  | Contains Duplicate                                  | [Go](https://github.com/NongBritee/leetcode-go/blob/main/leetcode/0217.Contains%20Duplicate.go) [Array, Map]                                  |         <green>Easy</green>         | <sub>1. Use map to check the duplicate number</sub>                                                                                                                                                                                     |
| 0383  | Ransoms Note                                        | [Go](https://github.com/NongBritee/leetcode-go/blob/main/leetcode/0383.Ransom%20Note.go)  [Array, Map]                                        |         <green>Easy</green>         | <sub>1. Map magazine [char, count] </br>2. Loop through ransoms note and check char from Magazine map is enough for ransoms note by decrease -1 every time that found char in map<br/></sub>                                            |
| 0412  | Fizz Buzz                                           | [Go](https://github.com/NongBritee/leetcode-go/blob/main/leetcode/0412.Fizz%20Buzz.go) [Array, Modular]                                       |         <green>Easy</green>         | <sub>1. Loop through array and push to new array when mod by 3, 5</sub>                                                                                                                                                                 |
| 1342  | Number of Steps to Reduce a Number to Zero          | [Go](https://github.com/NongBritee/leetcode-go/blob/main/leetcode/1342.Number%20of%20Steps%20to%20Reduce%20a%20Number%20to%20Zero.go) [Math]  |         <green>Easy</green>         | <sub>1. Loop until result == 0<br/> 2. divided by 2 and -1 when result is odd number<br/> 3. Count every step of devide and minus operation</sub>                                                                                       |
| 1347  | Minimum Number of Steps to Make Two Strings Anagram | [Go]() [Map]                                                                                                                                  |               <orange>Medium</orange>                | <sub>1. create sMap of s string [char, count] by count the no. of char<br/> 2. loop through t string and find char in sMap<br/> 3. -1 from count when found (stop minus when 0)<br/> 4. return sum of remaining count in the sMap</sub> |
| 1480  | Running Sum of 1d Array                             | [Go](https://github.com/NongBritee/leetcode-go/blob/main/leetcode/1480.Running%20Sum%20of%201d%20Array.go) [Array]                            |         <green>Easy</green>         | <sub>1. Loop through array start from 2nd index and add previous value to current value</sub>                                                                                                                                           |
| 1672  | Richest Customer Wealth                             | [Go](https://github.com/NongBritee/leetcode-go/blob/main/leetcode/1672.Richest%20Customer%20Wealth.go) [Array]                                |         <green>Easy</green>         | <sub>1. Loop through array and sum each array then compare to max value</sub>                                                                                                                                                           |
|       |                                                     |                                                                                                                                               |                                     |                                                                                                                                                                                                                                         |