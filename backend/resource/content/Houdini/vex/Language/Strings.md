---
title: Strings
order: 5
---
| On this page | * [Overview](#overview) * [String Literals](#string-literals) * [Escape sequences in wrangles and VEXpressions](#escape-sequences-in-wrangles-and-vexpressions) * [Declaring string types](#declaring-string-types) * [Accessing and setting string values](#accessing-and-setting-string-values) * [Looping over a string](#looping-over-a-string) * [Working with strings](#working-with-strings) |
| --- | --- |

Overview

## overview

VEX includes a string datatype. This is useful in several places:

- Manipulating text
- Referencing filenames and op node names
- Manipulating binary data

String Literals

## string-literals

String literals can be enclosed in either single quotes (') or double quotes ("). Strings may also be specified using the Python or C++ raw-string format.

```vex
string s = 'foo';
string t = "bar";
string py = r"Hello world\n";        // Python style, equivalent to "Hello world\\n"
string cpp = R"(Hello world\n)";   // C++ style, equivalent to "Hello world\\n"

```

Escaped strings (non-raw strings) automatically convert known escape sequences
to their representative byte sequences. For example “\\n” will convert to the
ASCII byte to emit a newline.

Raw strings ignore escape sequences. For a raw string, the “\\n” will be
interpreted literally as a backslash and the lower case `n`.

The syntax for strings can be summarized

- Escaped strings `"text"` or `'text'`
- Python raw-strings `r"raw text"`
- C++ raw-string `R"delimiter(raw text)delimiter"`
  Where the `delimiter` is an optional string of 0 to 16 characters. Unlike
  Python raw-strings, C++ style raw strings can contain multi-line text and
  even binary data.

```vex
string escaped = 'Line 1\nLine 2';
string raw = r"Line 1\nLine 1 continues";        // "Line 1\\nLine 1 continues"
string cppraw = R"(Line 1\nLine 1 continues)";        // "Line 1\\nLine 1 continues"
string cppmultiline = R"multi(This is a long
    string which has multiple lines.  The string
    also contains an embedded raw string R"(raw string)"
    But since the delimiter doesn't match, the string isn't
    actually ended until here.)multi";

```

Escape sequences in wrangles and VEXpressions

## escape-sequences-in-wrangles-and-vexpressions

It is important to note that inside Houdini one rarely writes VEX code
directly. Instead, the snippet written will be channel-referenced into
a VOP and then used to generate code. The sequence of transformations
from an Attribute Wrangle into the final VEX will result in a lot of
substitutions.

In particular, '$' will start an environment variable expansion and ''
will trigger escape sequences **before** VEX receives the code. Use the
View Code option in VEX to see the complete result of the transformation.

You can generate a backslash with `chr(92)` and a dollar sign with
`chr(36)`.

Declaring string types

## declaring-string-types

To declare a string variable, the general form is
`string var_name`:

```vex
// My string is a normal string
string   mystring;

```

To declare a function that returns a string:

```vex
// A function which returns a string
string rgb_name()
{
...
};    

```

To specify a literal array, use double quotes or single quotes.

```vex
string a_string = "hello world!";
string another_string = 'good-bye!'

```

Accessing and setting string values

## accessing-and-setting-string-values

Use `string[index]` to look up a character by its position in the
array.

The index is a byte offset into the string, not a character offset. This is an
important distinction when you deal with Unicode strings. VEX assumes a UTF-8
encoding for all strings. If the given offset isn’t a valid UTF-8 character,
an empty string is returned. Otherwise, the full UTF-8 character is returned -
this may be a string of length greater than one!

String bounds are checked at run time. Reading out of bounds will result in an
empty string. This may generate a warning or optional run-time error in the
future.

Python-style indexing is used. This means negative indices refer to positions
from the end of the array.

The slice notation can be used with the square brackets to extract ranges of a
string. This will operate on byte sequences, side stepping the UTF-8
requirements of the normal square bracket operation. Thus if you want the
third byte, regardless of whether it is a valid UTF-8 or not, use:

```vex
string a_string = "hello world!";
string thirdbyte = a_string[2:3];

```

You cannot assign values to an array using square brackets.

(The [getcomp](functions/getcomp.html "Extracts a single component of a vector type, matrix type, or array.") function is the equivalent for using
the square brackets notation.)

Looping over a string

## looping-over-a-string

See [foreach](functions/foreach.html "Loops over the items in an array, with optional enumeration.").

Note you will get empty strings for the offsets which do not correspond to valid unicode characters.

Working with strings

## working-with-strings

The following functions let you query and manipulate arrays.

[len](functions/len.html "Returns the length of an array.")

Returns the length of a string.

[append](functions/append.html "Adds an item to an array or string.")

Adds another array to the end of this one.

[ord](functions/ord.html "Converts an UTF8 string into a codepoint.")

Converts a UTF-8 string to a codepoint.

[chr](functions/chr.html "Converts an unicode codepoint to a UTF8 string.")

Converts a codepoint to a UTF-8 string.
