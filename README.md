xpipe
=====

A small XML processing tool which doesn't require XML to work.

Usage
-----

```
$ xpipe [SWITCHES] [FILES]
```

Valid switches are:

    - `-e`: Pipeline expression to execute.
    - `-x`: Select nodes matching an XPath expression instead of the entire DOM.
    - `-l`: Display filenames which contain results.
    - `-L`: Display filenames which contain no results.

Either `-x` or `-e` need to be defined at a minimum.

Expressions
-----------

Expressions consist of pipelines with a bunch of special statements.  Pipelines
take values, called datums, and produce other datums.  A datum can be:

    - Strings
    - Numbers
    - Booleans
    - XML Nodes
    - XML Documents

Pipelines consist of processors separated by the pipe character ("|").  Each
process transform datums in some way.

For example, a pipeline which changes the groupId of a Maven Pom file can be
written as so:

    xpath "/project/groupId" | settext "another.group"

By default, each individual pipeline starts with the entire XML document and passes
the final results to a process which prints it to standard out.  Using the `-x`
switch, an implicit `xpath` process will be place in the front of each pipeline.

Multiple pipelines can be defined, each one separated by the semicolon (";").  Each
individual pipeline is executed in the order they appear in the expression and make
use of the same DOM.  Note that some statements, line `settext`, may modify the
underlying XML DOM, which could affect pipelines running afterwards.

### Namespace Mappings

The XPath processor is namespace aware and will require explicit mapping of namespaces
that cannot be inferred.  To create a mapping from a URL to a prefix, add the `ns`
statement:

    ns m = "http://maven.apache.org/POM/4.0.0" ; xpath "/m:project/m:groupId"

Syntax
======

The full language syntax is given below:

```
<expression>        =   <statement> (";" <statement>)*
<statement>         =   <namespaceMapping> | <pipeline>

<namespaceMapping>  =   "ns" <prefix:IDENT> "=" <url:STRING>

<pipeline>          =   <process> ("|" <process>)*
<process>           =   <processName:IDENT> <processArgs:STRING>*
```
