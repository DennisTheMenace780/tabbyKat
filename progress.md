Aug 10, 2024
--
A small improvement was made to fix the strange rendering issue that I was
experiencing. The fix involes using the `tea.WithAltScreen()` function as a
program option. This option writes the Bubble Tea output to an alternate buffer,
which means that when the git command executes we're freeing up `stdout` for the
native git output. 

March 19, 2024
--
I got the version two of the project up and running again, so I feel comfortable
renaming everything to use tabbycat instead of tabbykat once the CI checks get
going.

For next session the following needs to be done:
  - X Clean up the poor logging throughout the codebase 
  - X Write another test case
  - X Write ReadMe for running the tests

---
The view part of `model.go` works quite well, but when I switch to various
branches it can hide the output in Jobber. Returning an empty string when the
`m.choice != ""` seems to work okay though, but I'd like to improve it just a
bit by showing the output of the command based on the user's input.
