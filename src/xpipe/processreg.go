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
    return &ProcessRegistry{make(map[string]*ProcessRegEntry)}
}

// Creates and configures a new process
func (pr *ProcessRegistry) NewProcess(name string, args []Datum) (Process, error) {
    ent, hasEnt := pr.Entries[name]
    if !hasEnt {
        return nil, fmt.Errorf("No such process: %s", name)
    }

    p := ent.Factory()
    return p.Config(args), nil
}
