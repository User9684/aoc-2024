# Day 3 Analysis

<details>
  <summary>[[ SPOILER WARNING ]]</summary>
  
### Input
We are given a string containing a mix of random characters along with:
  - `mul(n,n)`
  - `don't()`
  - `do()`

Example: `xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`

### Part 1 Task
Search for instances of mul(\<number\>, \<number\>) and multiply the two numbers together.

The total of all the multiplications added together is our answer for part 1.

### Part 2 Task
Just like part 1, we instead search for don't() and do().
The only difference is that now we only add the multiplication if it is "enabled" as per the most recently matched "call"

The total of all the multiplications with the logic included added is our answer for part 2.

</details>
