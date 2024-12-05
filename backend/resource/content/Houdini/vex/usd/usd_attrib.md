---
title: usd_attrib
order: 16
---
| Since | 17.5 |
| --- | --- |

`<type> usd_attrib(<stage>stage, string primpath, string name)`

`<type>[] usd_attrib(<stage>stage, string primpath, string name)`

`<type> usd_attrib(<stage>stage, string primpath, string name, float timecode)`

`<type>[] usd_attrib(<stage>stage, string primpath, string name, float timecode)`

This function returns a value of a given attribute on a given primitive.

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

The value of an existing attribute, or zero/empty value if the attribute does not exist. Use [usd_isattrib](usd_isattrib.html "Checks if the primitive has an attribute by the given name.") if you want to check whether the attribute exists.

Examples

## examples

```vex
// Get the value of some attributes on the cube primitive.
float a = usd_attrib("opinput:0", "/geo/cube", "attribute_name_a");
vector b[] = usd_attrib(0, "/geo/cube", "attribute_name_b");

// Get the value of attribute "bar" at various time codes.
f[]@b_at_current_frame = usd_attrib(0, "/geo/sphere", "bar");
f[]@b_at_frame_7       = usd_attrib(0, "/geo/sphere", "bar", 7.0);

```
