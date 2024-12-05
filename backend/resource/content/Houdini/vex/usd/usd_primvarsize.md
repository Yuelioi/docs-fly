---
title: usd_primvarsize
order: 113
---
| Since | 18.0 |
| --- | --- |

`int  usd_primvarsize(<stage>stage, string primpath, string name)`

This function returns the tuple size of a primvar found directly on the given primitive. If the primvar is an array, it returns the tuple size of the array element. E.g., for vector types, this is the number of components.

Show/hide arguments

`<stage>`

When running in the context of a node (such as a wrangle LOP), this argument can be an integer representing the input number (starting at 0) to read the stage from. The integer is equivalent to the string form referencing a particular input, e.g., “opinput:0”.

You can also use this argument to refer to a USD file (e.g., “/path/to/file.usd”), or to another LOP node’s cooked stage using the `op:` as the path prefix (e.g., “op:/stage/lop_node”).

`primpath`

The path to the primitive.

`name`

Primvar name (without namespace).

Returns

The tuple size of the primvar.

- For a vector type, this is the number of components.
- For an integer, float, or string, this returns `1`.
- For an array primvar, this returns the tuple size of the elements.

If the primvar does not exist, returns `0`.

Use [usd_primvarlen](usd_primvarlen.html "Returns the length of the array primvar directly on the USD primitive.") if you want to obtain the array primvar length.

Examples

## examples

```vex
// Get the tuple size of a primvar on the cube primitive.
int tuple_size = usd_primvarsize(0, "/geo/cube", "primvar_name");

```
