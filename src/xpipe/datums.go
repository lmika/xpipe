// Datums of XPipe
//

package xpipe

import (
    "fmt"
    "github.com/moovweb/gokogiri/xml"
)


// Base datum of an XProc pipeline.
type Datum interface {

    // Returns the standard string representation of the datum
    String() string
}


// A string datum
type StringDatum        string

func (s StringDatum) String() string {
    return string(s)
}


// A boolean datum
type BoolDatum          bool

func (b BoolDatum) String() string {
    if bool(b) {
        return "true"
    } else {
        return "false"
    }
}


// A number datum
type NumberDatum        float64

func (n NumberDatum) String() string {
    return fmt.Sprintf("%g", float64(n))
}


// A node datum
type NodeDatum          struct {
    Node        *xml.XmlNode
}

func (n NodeDatum) String() string {
    return n.Node.String()
}


// A document datum
type DocDatum           struct {
    Doc         *xml.XmlDocument
}

func (d DocDatum) String() string {
    return d.Doc.String()
}

// --------------------------------------------------------------------------
//

// Execution context for a pipeline.
type ProcessContext    struct {
    Filename    string
}
