# Day 4 Analysis

<details>
  <summary>[[ SPOILER WARNING ]]</summary>
  
### Input
We are given a "word search" string containing the word "XMAS" along with "MAS" in an "X" formation.

Example:  
<code>MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX</code>

### Part 1 Task

Search for all occurances of "XMAS" going left or right, down or up, or in diagonal directions.

The total occurances is our answer for part 1.

### Part 2 Task

Iterate over each character, checking if it is an "A", and if so then check if there is two "MAS" attached going diagonally, creating an X.
For example, if we have an "A" on X5,Y6, we would check for an "M" to each corner on that A, and then check for an S on the opposite corner.

The total amount of "X" MAS formations is our answer for part 2.

</details>
