---
title: Vex compiler (vcc)
order: 8
---
| On this page | * [Command-line options](#command-line-options)   + [Compiler options](#compiler-options)  + [Preprocessor options](#preprocessor-options)  + [Diagnostic options](#diagnostic-options)  + [VEX Context options](#vex-context-options)  + [Asset options](#otl)  + [Dialog Script options](#ds) * [Pre-processor](#pre-processor)   + [Predefined macros](#predefined-macros) |
| --- | --- |
The `vcc` compiler compiles VEX source code into a form executable by
Houdini. The VEX compiler (vcc) is capable of compiling VEX code,
generating dialog scripts for VEX functions, and also giving quick help
by listing the global variables and functions available in any given context.
Command-line options

## command-line-options

Note

Most options have a short and a long form. The long form is shown in
square brackets after the short form.

`-h [--help]`

Show help message for the compiler.

### Compiler options

compiler-options

`-o [--vex-output] [file|-]`

Specify that a VEX code output is required. If the filename given
is 'stdout', then the output will be printed to the console.
If none of `--hda-output`, `--otl-output`, or `--ds-output` are specified, this
option is implied. By default the output filename is constructed from the
context function. If no context function is defined, then the input
filename is used as a base. If multiple input files are given, then the
filename argument will be ignored, and the output filename constructed
from the aforementioned rules.

`-d [--compile-all]`

Compile all functions into the VEX code, even if they're not used by the
context function, either directly or indirectly. This option is useful for
doing syntax checking on include files.

`-z [--no-optimize]`

Generate unoptimized VEX code.

`-V [--no-version-id]`

Don’t embed a Houdini version identifier in the VEX code.

### Preprocessor options

preprocessor-options

`-E [--parse-only]`

Only parse the input files to the standard output. No compilation will take
place.

`-D [--define] name[=value]`

Define a macro for the pre-processor. If no value is given with the
name, the name is defined as 1.

`-I [--include-dir] path`

Add the path specified to the include path (the list of directories
search for files referenced by the `#include` directive to the
pre-processor). The standard Houdini include path is under
`vex/include`.

### Diagnostic options

diagnostic-options

`-w id[,id...]`

Suppress printing of certain warnings and information messages.
wlist is a comma-separated list of warning numbers to suppress.

`-F [--Werror]`

Treat all non-suppressed warnings as errors.

`-e [--Werror-output] file`

Redirect all diagnostic output to the filename given, rather than print
them to the standard error output.

`-q [--Wno-info]`

Suppress informational messages.

`-Q [--Werror-only]`

Suppress informational and warning messages. Overrides `--Wno-info`.

`--fmessage-limit <N>`

Set a maximum number of messages that will be printed before stopping.
Set to 0 (the default) for no limit.

### VEX Context options

vex-context-options

`-c [--context] name`

If no context function is defined, this can be used to specify which
VEX context to use when compiling the source(s).

`-X [--list-context] context`

Print out global variables and function signatures defined for the given
VEX context. The argument value of `contexts` can be used to list all
available VEX contexts.

### Asset options

otl

`-L [--hda-append] [file|-]`

Append a digital asset generated from the VEX source to the specified operator type library file. If the file doesn’t exist, it will be created.

`-l [--hda-output] [file|-]`

Write a digital asset definition for the context function to the specified operator type library file. If the library file already exists, it will be overwritten.

`-K [--hda-vex-section] name`

Store the generated VEX code in the given section name in the HDA, rather
than the standard section name for the given VEX context.

`-a [--hda-dialog-script] file`

Use the parameter definitions in the given file instead of the ones
automatically generated from the VEX source. The dialog script’s operator
definition will still be taken from the VEX source.

`-U [--hda-dialog-script-only]`

Only embed the dialog script into the OTL. No VEX code will be added.

`-n [--op-name] name`

Use the given name as the name for the operator. This overrides any
[`#pragma opname`](pragmas.html#opname) statement in the code.

`-S [--op-script-name] name`

Use the given name as the script name for the operator. This overrides any
[`#pragma opscript`](pragmas.html#opscript) statement in the code.

`-N [--op-label] name`

Use the given name as the UI label for the operator. This overides any
[`#pragma oplabel`](pragmas.html#opname) statement in the code.

`-C [--op-icon] name`

Use the given name as the icon to use for the operator. This overides any
[`#pragma opicon`](pragmas.html#opicon) statement in the code.

`-t [--op-min-inputs] N`

Set the minimum number of inputs for the operator. This overides any
[`#pragma opmininputs`](pragmas.html#opinputs) statement in the code. The minimum
input value will be adjusted to fit the operator type being generated.

`-T [--op-max-inputs] N`

Set the maximum number of inputs for the operator. This overides any
[`#pragma opmaxinputs`](pragmas.html#opinputs) statement in the code. The maximum
input value will be adjusted to fit the operator type being generated.

### Dialog Script options

ds

`-u [--ds-output] [file|-]`

Write a dialog script to the filename given. As with `--vex-output` and
`--hda-output`, if multiple input files are given, then the filename given
is ignored and the output filename automatically constructed from either
the context function name, or the input filename, if no context function is
defined.

Pre-processor

## pre-processor

The compiler has a pre-processor which strips comments, reads include files,
and expands macros.

The pre-processor supports many of the usual C Pre-Processor
directives:

`#define name token-string`

Replace subsequent uses of name with token-string.

`#define name(arg,...,arg) token-string`

Replace subsequent instances of name with token-string. Each
argument to name is replaced in token-string during
expansion.

`#undef name`

“Undefine” the macro so subsequent uses of name are not
expanded.

`#include "filename"`

Inserts the contents of the file at this point in the source code.
When you use quotes, the directory containing the current file is
searched for filename before the standard locations (including
the path).

`#sinclude "filename"`

Like `include` but silent (no warnings or errors when the file isn’t found).

`#includeall "filename"`

Like `include` but scans the include search path, including all files found.

`#ifdef name`

The following lines until the next `#endif` or `else` directive will
be compiled if and only if name is a defined macro.

`#ifndef name`

The lines following will be compiled if and only if name is *not* a
defined macro.

`#if constant-expr`

The following lines until the next `#endif` or `#else` directive
will be compiled if and only if constant-expr evaluates to
non-zero.

The expression can use the following operators:

- Comparisons (`==`, `!=`, `<=`, `>=`, `<`, `>`)
- Logical AND (`&&`), OR (`||`), and NOT (`!`).
- Bitwise AND (`&`), OR (`|`), exclusive OR (`^`), and NOT (`~`).
- Arithmetic (`+`, `-`, `*`, `/`, `%`).
- Parentheses.

The expression can also use the following functions:

`defined(name)`

Returns 1 if the name is a defined macro, or 0 if it is
not.

```vex
#if defined(foo) && defined(fum)

```

`environment(name)`

Returns 1 if name is a defined environment variable.

`access(filename)`

Returns 1 if filename can be read by the application, or 0
if the file cannot be read.

```vex
#if access("/etc/passwd")
#include </etc/passwd>
#endif

```

`strcmp(str1, str2)`

Works the same as the C/C++ function of the same name, if the
two strings have the same contents, the function returns 0. Each
argument should be a quoted string or a macro that expands to a
quoted string.

```vex
#define VALUE "foo"
#if strcmp(VALUE, "bar") != 0
This statement is false since "foo" != "bar"
#endif
#if !strcmp(VALUE, "foo")
This statement is TRUE since strcmp("foo", "bar") == 0
#endif

```

Expressions are evaluated from left to right (**unlike the ANSI C
standard of right to left**). As with the ANSI pre-procssor, all
numbers must be integers.

`#else`

The following lines until the next `#endif` directive will be
compiled if and only if the previous `#if` directive evaluated to
zero.

`#endif`

Marks the end of a section of conditional code. Every test directive
must have a matching `#endif`.

`#pragma ...`

Specifies extended language features. See the [list of pragmas](pragmas.html).

### Predefined macros

predefined-macros
The following macros are pre-defined:

`__vex`

This symbol is always defined. You can use this in an `if` pre-processor directive to check that the program is being compiled by vcc.

`__vex_major`

The major version number of the compiler (Houdini version number).

`__vex_minor`

The minor version number of the compiler (Houdini version number).

`__vex_build`

The build version number of the compiler (Houdini version number).

`__vex_patch`

The patch version number of the compiler (Houdini version number).

`__LINE__`

The current line number of the source file.

`__FILE__`

The file being compiled.

`__DATE__`

The current date (as a quoted string). Example: `"Dec 31 1999"`

`__TIME__`

The current time (as a quoted string). Example: `"23:59:59"`

```vex
printf("Starting shader %s at %s", __FILE__, __DATE__);

```
