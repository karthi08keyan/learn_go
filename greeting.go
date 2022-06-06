package greetings
import "fmt"

func hello(name string) string {
var message string
  message = fmt.Sprintf("hi, %v. welcome!",name) // shortcut for declaring and intilazing {message := fmt.Sprintf("Hi, %v. Welcome!", name)}
  return message
}
.////.//././/././././././../././././.
package greetings

import "fmt"

// Hello returns a greeting for the named person.
func Hello(name string) string {
    // Return a greeting that embeds the name in a message.
    message := fmt.Sprintf("Hi, %v. Welcome!", name)
    return message
}
