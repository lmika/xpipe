// Parser for pipe expressions
package xpipe

import (
    "io"
    "fmt"
    "text/scanner"
)


// Scanner
type Scanner struct {
    S           scanner.Scanner
    Token       rune
    TokenText   string
}

// Creates a new scanner
func NewScanner(r io.Reader, filename string) *Scanner {
    s := scanner.Scanner{}
    s.Init(r)

    tok := s.Scan()
    toktext := s.TokenText()

    return &Scanner{s, tok, toktext}
}

// Fetch the next token
func (s *Scanner) Scan() {
    s.Token = s.S.Scan()
    s.TokenText = s.S.TokenText()
}

// Returns the position
func (s *Scanner) Position() string {
    return s.S.Pos().String()
}

// --------------------------------------------------------------------------
//

// A Parser
type Parser struct {
    Scanner     *Scanner
}

// Creates a scan error
func (p *Parser) Error(msg string) error {
    return fmt.Errorf("%s: error - %s", p.Scanner.Position(), msg)
}

// Creates a new parser 
func NewParser(r io.Reader, filename string) *Parser {
    return &Parser{ NewScanner(r, filename) }
}

// Parses a script.
//  <script> = <items>*
func (p *Parser) ParseScript() (*AstScript, error) {
    items := make([]AstItem, 0)
    for p.Scanner.Token != scanner.EOF {
        item, err := p.parseItem()
        if err != nil {
            return nil, err
        }

        items = append(items, item)
    }

    return &AstScript{items}, nil
}

// Parse a single item.
//  <item> = <pipeline>
func (p *Parser) parseItem() (AstItem, error) {
    return p.parsePipeline()
}

// Parses a pipeline
//  <pipeline> = <process> ("|" <process>)
func (p *Parser) parsePipeline() (*AstPipeline, error) {
    var firstPr *AstProcess = nil
    var currPr *AstProcess = nil

    for p.Scanner.Token != scanner.EOF {
        pr, err := p.parseProcess()
        if err != nil {
            return nil, err
        }

        if firstPr == nil {
            firstPr = pr
        }
        if currPr != nil {
            currPr.Next = pr
        }
        currPr = pr

        // If the next token is not '|' then break
        if (p.Scanner.Token != '|') {
            break
        } else {
            p.Scanner.Scan()
        }
    }

    return &AstPipeline{firstPr}, nil
}

// Parses a single process
//  <process> = IDENT
func (p *Parser) parseProcess() (*AstProcess, error) {
    if p.Scanner.Token != scanner.Ident {
        return nil, p.Error("Expected identifier")
    }

    pr := &AstProcess{p.Scanner.TokenText, nil}
    p.Scanner.Scan()

    return pr, nil
}
