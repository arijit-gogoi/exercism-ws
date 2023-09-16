package airportrobot

import "fmt"

type Greeter interface {
	LanguageName() string
	Greet(s string) string
}

func SayHello(n string, g Greeter) string {
	return g.Greet(n)
}

type Italian struct {
}

func (l Italian) LanguageName() string {
	return "Italian"
}

func (l Italian) Greet(name string) string {
	return fmt.Sprintf("I can speak %s: Ciao %s!", l.LanguageName(), name)
}

type Portuguese struct {
}

func (l Portuguese) LanguageName() string {
	return "Portuguese"
}

func (l Portuguese) Greet(name string) string {
	return fmt.Sprintf("I can speak %s: Olá %s!", l.LanguageName(), name)
}
