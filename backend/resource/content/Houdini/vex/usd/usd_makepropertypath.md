---
title: usd_makepropertypath
order: 88
---
| Since | 18.0 |
| --- | --- |

`string  usd_makepropertypath(<stage>stage, string primpath, string name)`

This function returns the full path of a given property.

Show/hide arguments

`<stage>`

When running in the context of a node (such as a wrangle LOP), this argument can be an integer representing the input number (starting at 0) to read the stage from. The integer is equivalent to the string form referencing a particular input, e.g., “opinput:0”.

You can also use this argument to refer to a USD file (e.g., “/path/to/file.usd”), or to another LOP node’s cooked stage using the `op:` as the path prefix (e.g., “op:/stage/lop_node”).

`primpath`

The path to the primitive.

`name`

Property name.

Returns

The full path of a given property.

Examples

## examples

```vex
// Obtain the full path to the property "prop_name" on the cube primitive.
string prop_path = usd_makepropertypath(0, "/geo/cube", "prop_name");

```
