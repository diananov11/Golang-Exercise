package airportrobot

import "fmt"

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.

type Greeter interface {
    LanguageName() string
    Greet(name string) string
}


type germanGreeter struct {
}

func (g germanGreeter) LanguageName() string {
    return "German"
}

func (g germanGreeter) Greet(n string) string {
    return fmt.Sprintf("Hallo %v!",n)
}
func SayHello(name string, g Greeter) string {
    return fmt.Sprintf("I can speak %v: %v",g.LanguageName(),g.Greet(name) )
}


type Italian struct {}

func (i Italian) LanguageName() string {
    return "Italian"
}

func (i Italian) Greet(n string) string {
    return fmt.Sprintf("Ciao %v!",n)
}

type Portuguese struct {}

func (p Portuguese) LanguageName() string {
    return "Portuguese"
}

func (p Portuguese) Greet(n string) string {
    return fmt.Sprintf("Olá %v!", n)
}
