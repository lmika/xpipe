// Base types for a pipeline.

package xpipe

import (
)


// A process factory
type ProcessFactory func() Process

type ProcessSink interface {
    // Accepts data from a process
    Accept(ctx *ProcessContext, d Datum) error
}



// A single process.
type Process interface {

    // Configures the process using the arguments from pipeline definition
    Config(args []Datum) error

    // Applies the process with the specific datum.
    Apply(ctx *ProcessContext, in Datum, sink ProcessSink) error
}


// ----------------------------------------------------------------------

// A pipline chain
type PipelineChain struct {
    Process         Process
    Next            *PipelineChain
}

// Implementation of the ProcessSink interface.  This will forward the data to
// the next process in the chain if one is defined.
func (pc *PipelineChain) Accept(ctx *ProcessContext, out Datum) error {
    if pc.Next != nil {
        return pc.Next.Process.Apply(ctx, out, pc.Next)
    } else {
        return nil
    }
}

// ----------------------------------------------------------------------

// A pipeline
type Pipeline struct {
    Start           *PipelineChain
    End             *PipelineChain
}

// Creates a new pipeline
func NewPipeline() *Pipeline {
    return &Pipeline{nil, nil}
}

// Appends a process to the end of the pipeline.
func (p *Pipeline) Append(proc Process, configArgs []Datum) {
    pc := &PipelineChain{proc, configArgs, nil}

    if (p.Start == nil) && (p.End == nil) {
        // Pipeline is empty
        p.Start = pc
        p.End = pc
    } else if (p.Start != nil) && (p.End != nil) {
        p.End.Next = pc
        p.End = pc
    } else {
        panic("Invariant violated: either start or end is not nil")
    }
}

// Prepends a process to the start of the pipeline.
func (p *Pipeline) Prepend(proc Process, configArgs []Datum) {
    pc := &PipelineChain{proc, configArgs, nil}

    if (p.Start == nil) && (p.End == nil) {
        // Pipeline is empty
        p.Start = pc
        p.End = pc
    } else if (p.Start != nil) && (p.End != nil) {
        pc.Next = p.Start
        p.Start = pc
    } else {
        panic("Invariant violated: either start or end is not nil")
    }
}

// Executes the process, starting a single datum
func (p *Pipeline) Accept(ctx *ProcessContext, d Datum) error {
    if p.Start != nil {
        return p.Start.Accept(ctx, d)
    } else {
        return nil
    }
}
