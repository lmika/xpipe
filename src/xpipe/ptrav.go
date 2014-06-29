// Pipelines that performs traversals over the nodes.  Non-nodes will
// be filtered out.

package xpipe

import (
)

// Selects the document of each datum.  If the datum is a node datum, returns
// the parent document of the node.  If the datum is a document datum, simply
// lets it pass through.  All other datums are filtered out.
type SelectDocumentProcess struct {
}

// Configures the process using the arguments from pipeline definition
func (p *SelectDocumentProcess) Config(args []ConfigArg) error {
    return nil
}

func (p *SelectDocumentProcess) Open(ctx *ProcessContext, sink ProcessSink) error {
    return SendOpen(sink, ctx)
}

func (p *SelectDocumentProcess) Close(ctx *ProcessContext, sink ProcessSink) error {
    return SendClose(sink, ctx)
}

// Applies the process with the specific datum.
func (p *SelectDocumentProcess) Apply(ctx *ProcessContext, in Datum, sink ProcessSink) error {
    switch x := in.(type) {
    case NodeDatum:
        return SendToSink(sink, ctx, DocDatum{x.Node.MyDocument()})
    case DocDatum:
        return SendToSink(sink, ctx, in)
    default:
        return nil
    }
}
