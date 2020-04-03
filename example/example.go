package main

import (
    "github.com/thomas-armena/repl"
)

type echoRepl struct {}

func (e echoRepl) Eval(input string) string {
    return input[:len(input)-1]
}

func (e echoRepl) ShouldEval(input string) bool {
    input = input[:len(input)-1]
    return input[len(input)-1] == ';'
}

func (e echoRepl) ShouldQuit(input string) bool {
    input = input[:len(input)-1]
    return input == "quit;"
}

func main() {
    echorepl := echoRepl{}
    repl.StartReplWithStdIO(echorepl)
}
