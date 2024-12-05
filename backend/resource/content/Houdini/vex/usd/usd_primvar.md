---
title: usd_primvar
order: 105
---
| Since | 18.0 |
| --- | --- |

`<type> usd_primvar(<stage>stage, string primpath, string name)`

`<type>[] usd_primvar(<stage>stage, string primpath, string name)`

`<type> usd_primvar(<stage>stage, string primpath, string name, float timecode)`

`<type>[] usd_primvar(<stage>stage, string primpath, string name, float timecode)`

This function returns a value of a primvar on a given primitive.

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

The value of an existing primvar, or zero/empty value if the primvar does not exist. Use [usd_isprimvar](usd_isprimvar.html "Checks if the primitive has a primvar of the given name.") if you want to check whether the primvar exists.

Examples

## examples

```vex
// Get the value of some primvars on the cube primitive.
vector vec_value = usd_primvar(0, "/geo/cube", "vec_primvar_name"); 
float values[] = usd_primvar(0, "/geo/cube", "primvar_name");
float value    = usd_primvar(0, "/geo/cube", "primvar_name", 3);

v[]@foo_at_current_frame = usd_primvar(0, "/geo/sphere", "foo");
v[]@foo_at_frame_8       = usd_primvar(0, "/geo/sphere", "foo", 8.0);

```
