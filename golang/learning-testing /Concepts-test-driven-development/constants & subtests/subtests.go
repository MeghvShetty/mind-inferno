/*
Sometimes, it is useful to group tests around a "thing" and then have subtests describing different scenarios.
A benefit of this approach is you can set up shared code that can be used in the other tests.
*/

package main

import "fmt"

const englishHelloPrefix = "Hello, "

func Hello(name string) string {
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("Chris"))
}
