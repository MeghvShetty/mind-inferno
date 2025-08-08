# Learning to write testing code for golang. 

Run go test in your terminal. It should've passed! Just to check, try deliberately breaking the test by changing the want string.

Notice how you have not had to pick between multiple testing frameworks and then figure out how to install them. Everything you need is built into the language, and the syntax is the same as the rest of the code you will write.

### Writing tests
Writing a test is just like writing a function, with a few rules

- It needs to be in a file with a name like ```xxx_test.go```
- The test function must start with the word ```Test```
- The test function takes one argument only ```t *testing.T```
- To use the *testing.T type, you need to ```import "testing"```, like we did with ```fmt``` in the other file

```t``` of type ```*testing.T``` is your "hook" into the testing framework so you can do things like ```t.Fail()``` when you want to fail.


Let's start by capturing these requirements in a test. This is basic test-driven development and allows us to make sure our test is actually testing what we want. When you retrospectively write tests, there is the risk that your test may continue to pass even if the code doesn't work as intended.
