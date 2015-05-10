/*
Brainfuck Interpreter
*/
package main

import (
    "errors"
    "fmt"
    "io/ioutil"
    "os"
)

func cleanup(data []byte) []byte {
    result := make([]byte, 0, len(data))
    for _, c := range data {
        switch c {
        case '>', '<', '+', '-', '.', ',', '[', ']':
            result = append(result, c)
        }
    }
    return result
}

func buildBracemap(data []byte) (bracemap map[int]int, err error) {
    stack := make([]int, 0, 100)
    bracemap = make(map[int]int)
    for pos, c := range data {
        switch c {
        case '[':
            stack = append(stack, pos)
        case ']':
            if len(stack) == 0 {
                err = errors.New("Syntax error: Unmatched closing brace")
                return bracemap, err
            }
            start := stack[len(stack)-1]
            bracemap[start] = pos
            bracemap[pos] = start
            stack = stack[:len(stack)-1]
        }
    }
    if len(stack) != 0 {
        err = errors.New("Syntax error: Not enough closing braces")
        return bracemap, err
    }
    return bracemap, nil
}

func runBrainfuck(data []byte) (tapeLength, instructionCount int, err error) {
    code := cleanup(data)
    bracemap, err := buildBracemap(code)
    if err != nil {
        return 0, 0, err
    }

    tape := make([]byte, 1, 1000)
    codePtr, tapePtr, instructionCount := 0, 0, 0
    for codePtr < len(code) {
        switch code[codePtr] {
        case '>':
            tapePtr++
            if tapePtr == len(tape) {
                tape = append(tape, 0)
            }
        case '<':
            tapePtr--
            if tapePtr == -1 {
                tape = append(tape, 0)
                copy(tape[1:], tape)
                tape[0] = 0
                tapePtr = 0
            }
        case '+':
            tape[tapePtr]++
        case '-':
            tape[tapePtr]--
        case '.':
            fmt.Print(string(tape[tapePtr]))
        case ',':
            fmt.Scanf("%c")
        case '[':
            if tape[tapePtr] == 0 {
                codePtr = bracemap[codePtr]
            }
        case ']':
            if tape[tapePtr] != 0 {
                codePtr = bracemap[codePtr]
            }
        }
        codePtr++
        instructionCount++
    }
    tapeLength = len(tape)
    return tapeLength, instructionCount, nil
}

func main() {
    if len(os.Args) != 2 {
        fmt.Println("Usage:", os.Args[0], "filename")
    } else {
        // read bf file
        data, err := ioutil.ReadFile(os.Args[1])
        if err != nil {
            fmt.Println(err)
            return
        }
        // run code
        tapeLength, instructionCount, err := runBrainfuck(data)
        if err != nil {
            fmt.Println(err)
        } else {
            fmt.Println("")
            fmt.Println("Tape length:", tapeLength)
            fmt.Println("Instruction count:", instructionCount)
        }
    }
}
