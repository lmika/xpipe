// Pipelines that print out things

package xpipe

import (
    "fmt"
)

// A pipeline that print each datum
type PrintProcess struct {
}

// Configures the process using the arguments from pipeline definition
func (p *PrintProcess) Config(args []ConfigArg) error {
    return nil
}

// Applies the process with the specific datum.
func (p *PrintProcess) Apply(ctx *ProcessContext, in Datum, sink ProcessSink) error {
    fmt.Println(in.String())
    return nil
}

// --------------------------------------------------------------------------

// A pipeline that writes a test message
type TestProcess struct {
}

// Configures the process using the arguments from pipeline definition
func (p *TestProcess) Config(args []ConfigArg) error {
    return nil
}

// Applies the process with the specific datum.
func (p *TestProcess) Apply(ctx *ProcessContext, in Datum, sink ProcessSink) error {
    return SendToSink(sink, ctx, StringDatum("Hello, world"))
}
