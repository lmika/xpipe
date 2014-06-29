// Process registry.
//

package xpipe

import (
    "fmt"
)

// A process registry entry
type ProcessRegEntry struct {
    Factory     ProcessFactory
}

// --------------------------------------------------------------------------
//

// A process registry
type ProcessRegistry struct {
    Entries     map[string]*ProcessRegEntry
}

// Creates a new process registry
func NewProcessRegistry() *ProcessRegistry {
    pr := &ProcessRegistry{make(map[string]*ProcessRegEntry)}
    pr.registerStandardProcessors()
    return pr
}

// Registers the standard processes
func (pr *ProcessRegistry) registerStandardProcessors() {
    pr.Entries["print"] = &ProcessRegEntry{func() Process { return &PrintProcess{} }}
    pr.Entries["test"] = &ProcessRegEntry{func() Process { return &TestProcess{} }}
    pr.Entries["xpath"] = &ProcessRegEntry{func() Process { return &XPathProcess{} }}
}

// Creates and configures a new process
func (pr *ProcessRegistry) NewProcess(name string, args []ConfigArg) (Process, error) {
    ent, hasEnt := pr.Entries[name]
    if !hasEnt {
        return nil, fmt.Errorf("No such process: %s", name)
    }

    p := ent.Factory()
    return p, p.Config(args)
}
