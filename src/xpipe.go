// Main process

package main

import (
    "fmt"
    "os"
    "os/user"
    "path/filepath"
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

    rt := xpipe.NewRuntime()

    // Read the users RC file
    if user, err := user.Current() ; err == nil {
        rcFilename := filepath.Join(user.HomeDir, ".xpiperc")
        err := rt.EvalFile(rcFilename)
        if err != nil && !os.IsNotExist(err) {
            die(err)
        }
    }

    if *flagExpression == "" {
        flag.Usage()
        os.Exit(2)
    }

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
