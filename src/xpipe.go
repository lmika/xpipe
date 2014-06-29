// Main process

package main

import (
    "fmt"
    "os"
    "flag"

    "./xpipe"
)

// Flag for describing the expression
var flagExpression *string = flag.String("e", "", "Process expression")


// Die with an error message
func die(err error) {
    fmt.Fprintf(os.Stderr, "%s: %s\n", os.Args[0], err.Error())
    os.Exit(1)
}

func main() {
    var err error
    flag.Parse()

    if *flagExpression == "" {
        flag.Usage()
        os.Exit(2)
    }

    rt := xpipe.NewRuntime()

    // Parse the expression
    err = rt.EvalString(*flagExpression, "-e")
    if err != nil {
        die(err)
    }

    // Execute for files
    if flag.NArg() == 0 {
        err = rt.ExecuteForFile("-")
        if err != nil {
            die(err)
        }
    } else {
        for _, filename := range flag.Args() {
            err = rt.ExecuteForFile(filename)
            if err != nil {
                die(err)
            }
        }
    }
}
