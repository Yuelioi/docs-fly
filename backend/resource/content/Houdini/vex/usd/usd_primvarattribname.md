---
title: usd_primvarattribname
order: 106
---
| Since | 18.0 |
| --- | --- |

`string  usd_primvarattribname(<stage>stage, string name)`

This function returns the namespaced attribute name corresponding to the given primvar name.

Show/hide arguments

`<stage>`

When running in the context of a node (such as a wrangle LOP), this argument can be an integer representing the input number (starting at 0) to read the stage from. The integer is equivalent to the string form referencing a particular input, e.g., “opinput:0”.

You can also use this argument to refer to a USD file (e.g., “/path/to/file.usd”), or to another LOP node’s cooked stage using the `op:` as the path prefix (e.g., “op:/stage/lop_node”).

`name`

Primvar name (without namespace).

Returns

The namespaced name of an attribute corresponding to the given primvar name.

Examples

## examples

```vex
// Get the attribute name for the given primvar.
string attrib_name  = usd_primvarattribname(0, "some_primvar");
int is_attrib = usd_isattrib(0, "/geo/sphere", attrib_name);

```
