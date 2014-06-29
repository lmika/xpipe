// Processors which select XPath expressions.

package xpipe

import (
    "fmt"

    "github.com/moovweb/gokogiri/xpath"
    "github.com/moovweb/gokogiri/xml"
)

// Pipeline which select nodes using an XPath expression and sends them to the
// sink as Node datums.

type XPathProcess struct {
    Expr        *xpath.Expression
}

func (xp *XPathProcess) Config(args []ConfigArg) error {
    if (len(args) != 1) {
        return fmt.Errorf("Expected a least 1 argument")
    }

    xpathStr := args[0].String()
    if err := xpath.Check(xpathStr) ; err != nil {
        return err
    }

    xp.Expr = xpath.Compile(xpathStr)
    return nil
}

func (xp *XPathProcess) Open(ctx *ProcessContext, sink ProcessSink) error {
    return SendOpen(sink, ctx)
}

func (xp *XPathProcess) Close(ctx *ProcessContext, sink ProcessSink) error {
    return SendClose(sink, ctx)
}

// Apply the XPath to documents
func (xp *XPathProcess) Apply(ctx *ProcessContext, in Datum, sink ProcessSink) error {
    switch x := in.(type) {
    case DocDatum:
        return xp.applyTo(ctx, x.Doc, x.Doc.Root(), sink)
    case NodeDatum:
        return xp.applyTo(ctx, x.Node.MyDocument(), x.Node, sink)
    default:
        // Do not forward anything else
        return nil
    }
}

func (xp *XPathProcess) applyTo(ctx *ProcessContext, doc xml.Document, node xml.Node, sink ProcessSink) error {
    xpathCtx := xpath.NewXPath(doc.DocPtr())
    defer xpathCtx.Free()

    // TODO: Register namespaces

    if err := xpathCtx.Evaluate(node.NodePtr(), xp.Expr) ; err != nil {
        return err
    }

    return xp.sendResultsToSink(xpathCtx, doc, ctx, sink)
}

func (xp *XPathProcess) sendResultsToSink(xpathCtx *xpath.XPath, doc xml.Document, ctx *ProcessContext, sink ProcessSink) error {
    switch xpathCtx.ReturnType() {
    case xpath.XPATH_NODESET:
        ns, err := xpathCtx.ResultAsNodeset()
        if err != nil {
            return err
        }

        for _, np := range ns {
            node := xml.NewNode(np, doc)
            err = SendToSink(sink, ctx, NodeDatum{node})
            if err != nil {
                return err
            }
        }
        return nil
    case xpath.XPATH_BOOLEAN:
        if b, err := xpathCtx.ResultAsBoolean() ; err != nil {
            return err
        } else {
            return SendToSink(sink, ctx, BoolDatum(b))
        }
    case xpath.XPATH_NUMBER:
        if n, err := xpathCtx.ResultAsNumber() ; err != nil {
            return err
        } else {
            return SendToSink(sink, ctx, NumberDatum(n))
        }
    case xpath.XPATH_STRING:
        if s, err := xpathCtx.ResultAsString() ; err != nil {
            return err
        } else {
            return SendToSink(sink, ctx, StringDatum(s))
        }
    default:
        return fmt.Errorf("Unreognised return type from XPath")
    }
}
