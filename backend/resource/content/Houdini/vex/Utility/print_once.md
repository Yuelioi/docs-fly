---
title: print_once
order: 11
---
`void  print_once(string msg, ...)`

Prints the string passed to the function exactly one time, even in a loop.
This is useful to print a message before the first iteration of a loop, without having to count iterations.

Show/hide arguments

`msg`

The string to print. This string does support interpolating values.
Use [sprintf](sprintf.html "Formats a string like printf but returns the result as a string
instead of printing it.") to generate the msg string if you need to include values.

"global",
`int`
`=0`

Normally, multiple instances of `print_once()` call sites
will work independently of each other. That is, if two separate
call sites to `print_once()` are passed the same string, the string will be
printed twice (once per call site). With the “global” flag turned on, strings
are checked across all instances of the `print_once()` functions.

Examples

## examples

```vex
// Only print "Hello world" one time
for (int i = 0; i < 100; ++i)
    print_once("Hello world\n");

// Print a missing texture warning, just one time across all shaders
print_once( sprintf("Missing texture map: %s\n", texture_map), "global", 1);

```
