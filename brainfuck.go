/*
Brainfuck Interpreter
*/
package main

import (
    "fmt"
    "io/ioutil"
    "os"
)

func check_syntax(data []byte) {
    if len(data) == 0 {
        panic("Syntax check failed")
    }
    return
}

func cleanup(data []byte) []byte {
    // Todo: Implement me
    return data
}

func buildBracemap(data []byte) map[int]int {
    stack := make([]int, 0, 100)
    bracemap := make(map[int]int)
    for pos, c := range data {
        switch string(c) {
        case "[":
            stack = append(stack, pos)
        case "]":
            start := stack[len(stack)-1]
            bracemap[start] = pos
            bracemap[pos] = start
            stack = stack[:len(stack)-1]
        }
    }
    return bracemap
}

func run_brainfuck_source(data []byte) (err error) {
    code := cleanup(data)
    check_syntax(data)
    bracemap := buildBracemap(data)
    tape := make([]byte, 1, 1000)

    code_ptr := 0
    tape_ptr := 0
    instructions := 0
    for code_ptr < len(code) {
        switch string(code[code_ptr]) {
        case ">":
            tape_ptr++
            if tape_ptr == len(tape) {
                tape = append(tape, 0)
            }
        case "<":
            tape_ptr--
            if tape_ptr == -1 {
                tape = append(tape, 0)
                copy(tape[1:], tape)
                tape[0] = 0
                tape_ptr = 0
            }
        case "+":
            tape[tape_ptr]++
        case "-":
            tape[tape_ptr]--
        case ".":
            fmt.Print(string(tape[tape_ptr]))
        case ",":
            fmt.Println("Read")
        case "[":
            if tape[tape_ptr] == 0 {
                code_ptr = bracemap[code_ptr]
            }
        case "]":
            if tape[tape_ptr] != 0 {
                code_ptr = bracemap[code_ptr]
            }
        }
        code_ptr++
        instructions++
    }

    fmt.Println("")
    fmt.Println("Required tape length:", len(tape))
    fmt.Println("Instruction count:", instructions)
    return
}

func main() {
    if len(os.Args) == 2 {
        // read input file
        data, err := ioutil.ReadFile(os.Args[1])
        if err != nil {
            panic(err)
        }

        // run brainfuck code
        err = run_brainfuck_source(data)
        if err != nil {
            panic(err)
        }
    } else {
        fmt.Println("Usage:", os.Args[0], "filename")
        return
    }
}
