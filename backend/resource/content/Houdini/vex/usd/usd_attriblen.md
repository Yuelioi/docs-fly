---
title: usd_attriblen
order: 18
---
| Since | 17.5 |
| --- | --- |

`int  usd_attriblen(<stage>stage, string primpath, string name)`

`int  usd_attriblen(<stage>stage, string primpath, string name, float timecode)`

This function returns the length of a given attribute.

For array attributes it is the length of the array, and for non-array attributes the length is 1.

Show/hide arguments

`<stage>`

When running in the context of a node (such as a wrangle LOP), this argument can be an integer representing the input number (starting at 0) to read the stage from. The integer is equivalent to the string form referencing a particular input, e.g., “opinput:0”.

You can also use this argument to refer to a USD file (e.g., “/path/to/file.usd”), or to another LOP node’s cooked stage using the `op:` as the path prefix (e.g., “op:/stage/lop_node”).

`primpath`

The path to the primitive.

`name`

Attribute name.

`timecode`

The USD time code at which to evaluate the attribute. A USD time code roughly corresponds to a frame in Houdini. If not given, the time code corresponding to the current frame is used.

Returns

The length of the array attribute, or 1 if the attribute is not an array. Use [usd_isarray](usd_isarray.html "Checks if the attribute is an array.") if you want to check whether the attribute is an array.

Examples

## examples

```vex
// Get the array length of an attribute on the cube primitive.
int length = usd_attriblen(0, "/geo/cube", "attribute_name");

```
