# Processes

### first

```
first
```

Sends through the first datum encountered.  Filters all other datums out.

### print

```
print
```

Prints the results of each datum to standard out.

```
xpath "/xpath/result" | print
```

### settext

```
settext <newvalue>
```

Changes the contents of nodes and document datums to *newvalue*.  All other datums
are sent to the sink unchanged.

### thisdoc

```
thisdoc
```

Selects the document of nodes and sends them to the sink.  Any documents encountered
will pass through unmodified.  All other datums are filtered out.

