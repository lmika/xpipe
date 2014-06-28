// Main process

package main

import (
    "./xpipe"
)


func main() {
    rt := xpipe.NewRuntime()
    rt.EvalString("test | print", "-")

    rt.ExecuteForFile("test.txt")
}
