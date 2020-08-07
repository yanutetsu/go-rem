package rem

import "fmt"

// Print prints a message by rem.
// It works like fmt.Print.
// You can use it only by changing from `fmt` to `rem`.
func Print(a ...interface{}) {
	fmt.Printf("(///ᴗㆁ✿) ＜ %s", fmt.Sprint(a...))
}
