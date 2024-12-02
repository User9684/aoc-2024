# Day 2 Analysis

<details>
  <summary>[[ SPOILER WARNING ]]</summary>
  
  ### Input
  We are given a multi-dimensional list of numbers (undefined dimension depth)  
  Example:  
  - 7 6 4 2 1
  - 1 2 7 8 9
  - 9 7 6 2 1
  - 1 3 2 4 5
  - 8 6 4 4 1
  - 1 3 6 7 9

### Part 1 Task

Iterate over each layer in the list, comparing the difference between each number.
 If the difference is lesser than 1 or greater than 3, it is unsafe.
 The differences must also be all increasing or all decreasing, it cannot have both.

The total number of "safe" layers is our answer for part 1.

### Part 2 Task

Same exact rules as in part 1, only difference being that if 1 number can be removed to make it safe, it is considered safe.

</details>
