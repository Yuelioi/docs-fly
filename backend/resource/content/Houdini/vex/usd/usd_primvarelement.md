---
title: usd_primvarelement
order: 107
---
| Since | 18.0 |
| --- | --- |

`<type> usd_primvarelement(<stage>stage, string primpath, string name, int index)`

`<type> usd_primvarelement(<stage>stage, string primpath, string name, int index, float timecode)`

This function returns a value of an element in given array primvar on a given primitive.

Show/hide arguments

`<stage>`

When running in the context of a node (such as a wrangle LOP), this argument can be an integer representing the input number (starting at 0) to read the stage from. The integer is equivalent to the string form referencing a particular input, e.g., “opinput:0”.

You can also use this argument to refer to a USD file (e.g., “/path/to/file.usd”), or to another LOP node’s cooked stage using the `op:` as the path prefix (e.g., “op:/stage/lop_node”).

`primpath`

The path to the primitive.

`name`

Primvar name (without namespace).

`index`

The index into the array.

`timecode`

The USD time code at which to evaluate the attribute. A USD time code roughly corresponds to a frame in Houdini. If not given, the time code corresponding to the current frame is used.

Returns

The value of an element in an existing array primvar, or zero/empty value if the primvar does not exist. Use [usd_isprimvar](usd_isprimvar.html "Checks if the primitive has a primvar of the given name.") if you want to check whether the primvar exists.

Examples

## examples

```vex
// Get the value of some primvars on the cube primitive.
float value    = usd_primvarelement(0, "/geo/cube", "primvar_name", 3);

v@element_2_at_current_frame  = usd_primvarelement(0, "/geo/sphere", "foo", 2);
v@element_2_at_frame_8        = usd_primvarelement(0, "/geo/sphere", "foo", 2, 8.0);

```
