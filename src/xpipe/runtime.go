// The main runtime

package xpipe

import (
    "strings"
    "bytes"
    "os"

    "github.com/moovweb/gokogiri"
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
func (rt *Runtime) ExecuteForFile(filename string) error {
    var file *os.File
    var err error

    if (filename == "-") {
        file = os.Stdin
    } else {
        file, err = os.Open(filename)
        if err != nil {
            return err
        }
        defer file.Close()
    }

    buffer := bytes.Buffer{}
    buffer.ReadFrom(file)

    doc, err := gokogiri.ParseXml(buffer.Bytes())
    if err != nil {
        return err
    }

    return rt.runPipelines(&ProcessContext{filename}, DocDatum{doc})
}

// Run the pipelines
func (rt *Runtime) runPipelines(ctx *ProcessContext, in Datum) error {
    for _, pl := range rt.Pipelines {
        if err := rt.runPipeline(pl, ctx, in) ; err != nil {
            return err
        }
    }
    return nil
}

// Runs a pipeline
func (rt *Runtime) runPipeline(p *Pipeline, ctx *ProcessContext, in Datum) error {
    return p.Accept(ctx, in)
}
