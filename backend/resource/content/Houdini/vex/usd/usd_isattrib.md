---
title: usd_isattrib
order: 67
---
| Since | 17.5 |
| --- | --- |

`int  usd_isattrib(<stage>stage, string primpath, string name)`

This function checks whether the given primitive has an attribute of a given name.

Show/hide arguments

`<stage>`

When running in the context of a node (such as a wrangle LOP), this argument can be an integer representing the input number (starting at 0) to read the stage from. The integer is equivalent to the string form referencing a particular input, e.g., “opinput:0”.

You can also use this argument to refer to a USD file (e.g., “/path/to/file.usd”), or to another LOP node’s cooked stage using the `op:` as the path prefix (e.g., “op:/stage/lop_node”).

`primpath`

The path to the primitive.

`name`

Attribute name.

Returns

`1` if the primitive has the given attribute, or `0` otherwise.

Examples

## examples

```vex
// Check if the sphere has an attribute "some_attribute".
int is_valid_attrib = usd_isattrib(0, "/geometry/sphere", "some_attribute");

```
