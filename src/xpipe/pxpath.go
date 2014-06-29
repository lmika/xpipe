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

// Apply the XPath to documents
func (xp *XPathProcess) Apply(ctx *ProcessContext, in Datum, sink ProcessSink) error {
    switch x := in.(type) {
    case DocDatum:
        return xp.applyToDocument(ctx, x.Doc, sink)
    default:
        // Do not forward anything else
        return nil
    }
}

func (xp *XPathProcess) applyToDocument(ctx *ProcessContext, doc *xml.XmlDocument, sink ProcessSink) error {
    xpathCtx := doc.DocXPathCtx()
    defer xpathCtx.Free()

    // TODO: Register namespaces

    if err := xpathCtx.Evaluate(doc.DocPtr(), xp.Expr) ; err != nil {
        return err
    }

    return xp.sendResultsToSink(xpathCtx, doc, ctx, sink)
}

func (xp *XPathProcess) sendResultsToSink(xpathCtx *xpath.XPath, doc *xml.XmlDocument, ctx *ProcessContext, sink ProcessSink) error {
    switch xpathCtx.ReturnType() {
    case xpath.XPATH_NODESET:
        ns, err := xpathCtx.ResultAsNodeset()
        if err != nil {
            return err
        }

        for _, np := range ns {
            node := xml.NewNode(np, doc)
            err = SendToSink(sink, ctx, &NodeDatum{node})
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
