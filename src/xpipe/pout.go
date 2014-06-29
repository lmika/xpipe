// Pipelines that print out things

package xpipe

import (
    "fmt"
)

//// print
//
//  Prints the results of each datum to standard out.
//
//  ```
//  xpath "/xpath/result" | print
//  ```
type PrintProcess struct {
}

// Configures the process using the arguments from pipeline definition
func (p *PrintProcess) Config(args []ConfigArg) error {
    return nil
}

func (p *PrintProcess) Open(ctx *ProcessContext, sink ProcessSink) error {
    return SendOpen(sink, ctx)
}

func (p *PrintProcess) Close(ctx *ProcessContext, sink ProcessSink) error {
    return SendClose(sink, ctx)
}

// Applies the process with the specific datum.
func (p *PrintProcess) Apply(ctx *ProcessContext, in Datum, sink ProcessSink) error {
    switch in.(type) {
    case DocDatum:
        fmt.Print(in.String())
    default:
        fmt.Println(in.String())
    }
    return SendToSink(sink, ctx, in)
}
