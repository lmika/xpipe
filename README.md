XPIPE
=====

A small XML processing tool which doesn't require XML to work.


EXPRESSIONS
-----------

XPipe expressions are built using a pipeline language.  Pipelines, by default
start off with the entire DOM being worked on, and end with displaying the resulting
values.

Values can be either:

    - Strings
    - Numbers
    - Booleans
    - Nodes of a DOM
    - A DOM

Example: displaying results from an XPath expression:

    xpath "/something/here"

Example: setting the value of nodes that match an XPath expression:

    xpath "/something/here" | setto "A Value"

Example: adding an attribute

    xpath "/something/here" | setattr "abc" "123"

Example: declaraing a profile

    profile mvn {
        ns m="http:bladibla";
    }

    xpath "/msomething" | setto "Fla"


Language syntax:

    <script> = <statements>
    <statements> = <statement> ((";" | "\n") <statement>)*

    <statement> = <nsmapping> | <profiledecl> | <pipeline>
    <nsmapping> = "ns" <ident> "=" <string>
    <profiledecl> = "profile" <ident> "{" <statements> "}"

    <pipeline> = <processchain>
    <processchain> = <process> ("|" <process>)*
    <process> = <ident> <arg>*
    <arg> = <string>
