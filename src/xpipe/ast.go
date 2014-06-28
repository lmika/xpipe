// Pipeline definition file parser

package xpipe

// --------------------------------------------------------------------------
// Ast


// Top level script
type AstScript struct {
    Items       []AstItem
}

// Creates a new runtime from the script
func (a *AstScript) Configure(rt *Runtime) error {
    for _, i := range a.Items {
        if err := i.Config(rt) ; err != nil {
            return err
        }
    }

    return nil
}

// ---------------------------------------------------------------------------
//

type AstItem    interface {
    // Configure part of the runtime
    Config(rt *Runtime) error
}

// A pipeline definition
type AstPipeline struct {
    Processes   *AstProcess
}

func (a *AstPipeline) Config(rt *Runtime) error {
    pl := NewPipeline()
    err := a.Processes.AddToPipeline(rt, pl)
    if err != nil {
        return err
    }

    rt.AddPipeline(pl)
    return nil
}

// ---------------------------------------------------------------------------
//

// A process invocation
type AstProcess struct {
    Name        string
    Next        *AstProcess
}

func (a *AstProcess) AddToPipeline(rt *Runtime, pl *Pipeline) error {
    p, err := rt.Registry.NewProcess(a.Name, []Datum {})
    if err != nil {
        return err
    }

    pl.Append(p)
    if a.Next != nil {
        return a.Next.AddToPipeline(rt, pl)
    } else {
        return nil
    }
}
