---
title: usd_istype
order: 82
---
| Since | 18.0 |
| --- | --- |

`int  usd_istype(<stage>stage, string primpath, string type)`

This function checks whether the given primitive is of a given type or derives
from a given type, e.g., an Cube, UsdGeomBoundable, etc.

Show/hide arguments

`<stage>`

When running in the context of a node (such as a wrangle LOP), this argument can be an integer representing the input number (starting at 0) to read the stage from. The integer is equivalent to the string form referencing a particular input, e.g., “opinput:0”.

You can also use this argument to refer to a USD file (e.g., “/path/to/file.usd”), or to another LOP node’s cooked stage using the `op:` as the path prefix (e.g., “op:/stage/lop_node”).

`primpath`

The path to the primitive.

`type`

The name or an alias of a type to check.

Returns

1 if the primitive is of a given type, and 0 otherwise.

Examples

## examples

```vex
// Check if the primitive is a Cube and is boundable
int is_cube_by_alias     = usd_istype(0, "/geo/cube", "Cube");
int is_cube_by_name      = usd_istype(0, "/geo/cube", "UsdGeomCube");
int is_boundable_by_name = usd_istype(0, "/geo/cube", "UsdGeomBoundable");

```
