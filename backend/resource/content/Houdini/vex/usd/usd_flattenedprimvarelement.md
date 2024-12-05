---
title: usd_flattenedprimvarelement
order: 41
---
| Since | 18.0 |
| --- | --- |

`<type> usd_flattenedprimvarelement(<stage>stage, string primpath, string name, int index)`

`<type> usd_flattenedprimvarelement(<stage>stage, string primpath, string name, int index, float timecode)`

This function returns a value of an element of a flattened array primvar on a given primitive.

Some primvars can be indexed, where the primvar is a compacted array of unique values, and there is an index array to map an entity to the value element. This function expands the compacted array by using the index array, and returns element value from the expanded array at a given index.

Show/hide arguments

`<stage>`

When running in the context of a node (such as a wrangle LOP), this argument can be an integer representing the input number (starting at 0) to read the stage from. The integer is equivalent to the string form referencing a particular input, e.g., “opinput:0”.

You can also use this argument to refer to a USD file (e.g., “/path/to/file.usd”), or to another LOP node’s cooked stage using the `op:` as the path prefix (e.g., “op:/stage/lop_node”).

`primpath`

The path to the primitive.

`name`

Primvar name (without namespace).

`index`

The index into the expanded array.

`timecode`

The USD time code at which to evaluate the attribute. A USD time code roughly corresponds to a frame in Houdini. If not given, the time code corresponding to the current frame is used.

Returns

The element of a flattened value array of an existing primvar, or zero/empty value if the primvar does not exist. Use [usd_isprimvar](usd_isprimvar.html "Checks if the primitive has a primvar of the given name.") if you want to check whether the primvar exists.

Examples

## examples

```vex
// Get the value of a flattened primvar on the cube primitive.
float flat_value    = usd_flattenedprimvarelement(0, "/geo/cube", "primvar_name", 3);

f@flat_primvar_element_10_at_current_frame = usd_flattenedprimvarelement(0, "/geo/sphere", "bar", 10);
f@flat_primvar_element_10_at_frame_7       = usd_flattenedprimvarelement(0, "/geo/sphere", "bar", 10, 7.0);

```
