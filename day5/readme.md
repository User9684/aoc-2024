# Day 5 Analysis

<details>
  <summary>[[ SPOILER WARNING ]]</summary>
  
### Input
We are given two sections, one listing "rules" for placement, and another listing "updates"

Example:  
<code>47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47</code>

### Part 1 Task

Iterate over all of the rule lines, creating a map of numbers to a list of numbers for it to be behind.
Then, we iterate over the "updates" section, checking if each number defined in the "rules" is behind the numbers they are supposed to be.
Finally, get the very middle number of the "update" and add it together if the "update" line is in the correct order according to the rules. 

The total sum of all the middle numbers is our answer to part 1.

### Part 2 Task

Part 2 is very similar to part 1. Major difference is that now you have to go in and re-order the "updates" line according to the defined rules, and then add up all the middle numbers of the reordered lines.

The total sum of all the middle numbers in the reordered "updates" is our answer to part 2.

</details>
