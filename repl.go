package repl

import (
    "bufio"
    "os"
    "github.com/fatih/color"
)

type Repl interface {
    Eval(string) string
    ShouldEval(string) bool
    ShouldQuit(string) bool
}

func StartReplWithStdIO(repl Repl) {
    reader := bufio.NewReader(os.Stdin)
    writer := bufio.NewWriter(os.Stdout)
    StartRepl(repl, reader, writer)
}

func StartRepl(repl Repl, reader *bufio.Reader, writer *bufio.Writer) error {
    currInput := ""
    for {
        if currInput == "" {
            write("> ", writer, color.FgGreen)
        } else {
            write("~ ", writer, color.FgBlue)
        }

        input, err := reader.ReadString('\n')

        if err != nil {
            return err
        }

        currInput += input

        if currInput == "" {
            break
        }


        if repl.ShouldEval(currInput) {
            if repl.ShouldQuit(currInput) {
                return nil
            }
            output := repl.Eval(currInput)
            write(output, writer, color.FgWhite)
            write("\n", writer, color.FgWhite)
            currInput = ""
        }
    }
    return nil
}

func write(text string, writer *bufio.Writer, col color.Attribute) {
    color.New(col).Fprint(writer, text)
    writer.Flush()
}


