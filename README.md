# leetcode-go

This repository contains my solutions to LeetCode problems. I've been actively practicing to enhance my problem-solving skills and algorithmic knowledge.

<a href="https://leetcode.com/nongbritee">
  <img src="https://img.shields.io/badge/Leetcode-orange?style=for-the-badge&logo=leetcode&logoColor=black"/>
</a>
<a href="https://www.linkedin.com/in/napong-jantaranimi-273983171">
  <img src="https://img.shields.io/badge/LinkedIn-0077B5?style=for-the-badge&logo=linkedin&logoColor=white"/> 
 </a> 
<a href="nongbriteenapong@gmail.com">
  <img src="https://img.shields.io/badge/Gmail-D14836?style=for-the-badge&logo=gmail&logoColor=white"/>
</a>

| No.  | Title              |                                            Solution                                             | Difficulty | Note                                                                                                                                                                |
|:----:|:-------------------|:-----------------------------------------------------------------------------------------------:|:----------:|:--------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| 0026 | Remove Duplicates from Sorted Array |        [Go](https://github.com/NongBritee/leetcode-go/blob/main/leetcode/0026.Remove%20Duplicates%20from%20Sorted%20Array.go)        |   Easy     | <sub>1. loop through nums array then append to new array when it's not duplicate to the latest value in new array</sub>                                             |
| 0189 | Rotate Array       |    [Go](https://github.com/NongBritee/leetcode-go/blob/main/leetcode/0189.Rotate%20Array.go)    |    Easy    | <sub>1. Split array into 2 array by k.<br/> 2. Concat it back new order (2nd array + 1st array).<br/> 3. modify k % len(nums) to avoid bug when k > len(nums)</sub> |
| 0217 | Contains Duplicate | [Go](https://github.com/NongBritee/leetcode-go/blob/main/leetcode/0217.Contains%20Duplicate.go) |    Easy    | <sub>1. Use map to check the duplicate number</sub>                                                                                                                 |
