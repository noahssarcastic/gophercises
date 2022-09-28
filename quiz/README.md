# Part 1

- [x] Read in a quiz provided via a CSV file 
- [x] Keep track of how many questions answered correctly
- [x] The CSV file should default to problems.csv
- [x] Specify CSV file with a flag
- [x] Output the total number of questions correct and how many questions there were in total

## Problem Format

The first column is a question and the second column in the same row is the answer to that question. Quizzes will be relatively short (< 100 questions) and will have single word/number answers.

eg.

```csv
5+5,10
7+3,10
1+1,2
8+3,11
1+2,3
8+6,14
3+1,4
1+4,5
5+1,6
2+3,5
3+3,6
2+4,6
5+2,7
```

# Part 2

- [x] Add a timer
- [x] Time limit should default to 30 seconds
- [x] Customize time limit via a flag
- [x] Quiz should stop as soon as the time limit has exceeded
- [x] Prompt user to press enter before the timer starts

# Bonus

- [x] Add string trimming and cleanup so extra whitespace & capitalization are ignored
- [ ] Add an option to shuffle the quiz order
