---
title: usd_primvartypename
order: 115
---
| Since | 18.0 |
| --- | --- |

`string  usd_primvartypename(<stage>stage, string primpath, string name)`

This function returns the type name of a primvar found directly on the given primitive. The type name is taken from the USD value type registry, e.g., “float”, “vector3d”, “double3”, etc.

Show/hide arguments

`<stage>`

When running in the context of a node (such as a wrangle LOP), this argument can be an integer representing the input number (starting at 0) to read the stage from. The integer is equivalent to the string form referencing a particular input, e.g., “opinput:0”.

You can also use this argument to refer to a USD file (e.g., “/path/to/file.usd”), or to another LOP node’s cooked stage using the `op:` as the path prefix (e.g., “op:/stage/lop_node”).

`primpath`

The path to the primitive.

`name`

Primvar name (without namespace).

Returns

The primvar type name used in the value type registry. E.g., “float”, “vector3d”, “double3”, etc.

Examples

## examples

```vex
// Get the type name of the primvar on cube.
string type_name = usd_primvartypename(0, "/geo/cube", "primvar_name");

```
