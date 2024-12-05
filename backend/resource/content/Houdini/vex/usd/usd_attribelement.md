---
title: usd_attribelement
order: 17
---
| Since | 18.0 |
| --- | --- |

`<type> usd_attribelement(<stage>stage, string primpath, string name, int index)`

`<type> usd_attribelement(<stage>stage, string primpath, string name, int index, float timecode)`

This function returns a value of an element in given array attribute on a given primitive.

Show/hide arguments

`<stage>`

When running in the context of a node (such as a wrangle LOP), this argument can be an integer representing the input number (starting at 0) to read the stage from. The integer is equivalent to the string form referencing a particular input, e.g., “opinput:0”.

You can also use this argument to refer to a USD file (e.g., “/path/to/file.usd”), or to another LOP node’s cooked stage using the `op:` as the path prefix (e.g., “op:/stage/lop_node”).

`primpath`

The path to the primitive.

`name`

Attribute name.

`index`

The element index in the array attribute.

`timecode`

The USD time code at which to evaluate the attribute. A USD time code roughly corresponds to a frame in Houdini. If not given, the time code corresponding to the current frame is used.

Returns

The value of an element in an existing attribute, or zero/empty value if the attribute does not exist. Use [usd_isattrib](usd_isattrib.html "Checks if the primitive has an attribute by the given name.") if you want to check whether the attribute exists.

Examples

## examples

```vex
// Get the value of an element at index 3 in the array attribute.
float a = usd_attribelement("opinput:0", "/geo/cube", "array_attrib_name", 3);

// Get the value of an element at index 2 of the "bar" array attribute.
@b_element_2_at_current_frame = usd_attribelement(0, "/geo/sphere", "bar", 2);
@b_element_2_at_frame_11      = usd_attribelement(0, "/geo/sphere", "bar", 2, 11.0);

```
