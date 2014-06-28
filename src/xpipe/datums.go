// Datums of XPipe
//

package xpipe

import (
    "github.com/moovweb/gokogiri/xml"
)


// Base datum of an XProc pipeline.
type Datum interface {
}


// A string datum
type StringDatum        string


// A boolean datum
type BoolDatum          bool


// A number datum
type NumberDatum        float64


// A node datum
type NodeDatum          struct {
    Node        *xml.XmlNode
}


// A document datum
type DocDatum           struct {
    Doc         *xml.XmlDocument
}


// --------------------------------------------------------------------------
//

// Execution context for a pipeline.
type ProcessContext    struct {
    Filename    string
}
