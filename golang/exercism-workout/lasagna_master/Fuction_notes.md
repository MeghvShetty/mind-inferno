## Return Values
The function parameters are followed by zero or more return values which must also be explicitly typed. Single return values are left bare, multiple return values are wrapped in parenthesis. Values are returned to the calling code from functions using the return keyword. There can be multiple return statements in a function. The execution of the function ends as soon as it hits one of those return statements. If multiple values are to be returned from a function, they are comma separated.

```
func Hello(name string) string {
  return "Hello " + name
}

func HelloAndGoodbye(name string) (string, string) {
  return "Hello " + name, "Goodbye " + name
}
```