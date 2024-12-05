---
title: usd_iprimvarindices
order: 54
---
| Since | 19.0 |
| --- | --- |

`int [] usd_iprimvarindices(<stage>stage, string primpath, string name)`

`int [] usd_iprimvarindices(<stage>stage, string primpath, string name, float timecode)`

This function returns the index array of an indexed primvar found directly on the given primitive or inherited from primitive’s ancestor.

Show/hide arguments

`<stage>`

When running in the context of a node (such as a wrangle LOP), this argument can be an integer representing the input number (starting at 0) to read the stage from. The integer is equivalent to the string form referencing a particular input, e.g., “opinput:0”.

You can also use this argument to refer to a USD file (e.g., “/path/to/file.usd”), or to another LOP node’s cooked stage using the `op:` as the path prefix (e.g., “op:/stage/lop_node”).

`primpath`

The path to the primitive.

`name`

Primvar name (without namespace).

`timecode`

The USD time code at which to evaluate the attribute. A USD time code roughly corresponds to a frame in Houdini. If not given, the time code corresponding to the current frame is used.

Returns

The index array of an indexed primvar, or zero/empty value if the primvar does not exist or is not indexed. Use [usd_isiprimvar](usd_isiprimvar.html "Checks if the primitive or its ancestor has a primvar of the given name.") if you want to check whether the primvar exists and [usd_isindexediprimvar](usd_isindexediprimvar.html "Checks if there is an indexed primvar directly on the USD primitive or on USD primitive’s ancestor.") to check whether it is indexed.

Examples

## examples

```vex
// Get the index array of an indexed primvar.
int indices[] = usd_iprimvarindices(0, "/geo/cube", "indexed_primvar_name");

```
