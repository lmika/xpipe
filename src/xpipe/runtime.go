// The main runtime

package xpipe

import (
    "strings"
)

// A runtime instance
type Runtime struct {
    Registry        *ProcessRegistry
    Pipelines       []*Pipeline
}


// Creates a new runtime
func NewRuntime() *Runtime {
    return &Runtime {
        Registry:       NewProcessRegistry(),
        Pipelines:      make([]*Pipeline, 0),
    }
}

// Adds a new pipeline
func (rt *Runtime) AddPipeline(p *Pipeline) {
    rt.Pipelines = append(rt.Pipelines, p)
}

// Evalutate a script from a string
func (rt *Runtime) EvalString(str string, fileName string) error {
    pr := NewParser(strings.NewReader(str), fileName)
    ast, err := pr.ParseScript()
    if err != nil {
        return err
    }

    return ast.Configure(rt)
}

// Execute for a file
func (rt *Runtime) ExecuteForFile(filename string) {
    rt.runPipelines(nil)
}

// Run the pipelines
func (rt *Runtime) runPipelines(ctx *ProcessContext) {
    for _, pl := range rt.Pipelines {
        pl.Accept(nil, nil)
    }
}
