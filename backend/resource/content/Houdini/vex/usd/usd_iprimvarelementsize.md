---
title: usd_iprimvarelementsize
order: 53
---
| Since | 19.0 |
| --- | --- |

`int  usd_iprimvarelementsize(<stage>stage, string primpath, string name)`

This function returns the element size of a primvar found directly on the given primitive or inherited from primitive’s ancestor.

The primvar element size applies to array primvars, but it does not encode the length of the array. It specifies how many consecutive array elements should be taken as an atomic element to be interpolated over a gprim. So, on a mesh, array length relates to element size like this `array_length = element_size * face_count`.

In most cases, the element size is `1`.

Note, element size is a USD concept and differs from the VEX tuple size obtained with [usd_iprimvarsize](usd_iprimvarsize.html "Returns the tuple size of the primvar directly on the USD primitive or on USD primitive’s ancestor.") or the VEX array length obtained with [usd_iprimvarlen](usd_iprimvarlen.html "Returns the length of the array primvar directly on the USD primitive or on USD primitive’s ancestor.").

Show/hide arguments

`<stage>`

When running in the context of a node (such as a wrangle LOP), this argument can be an integer representing the input number (starting at 0) to read the stage from. The integer is equivalent to the string form referencing a particular input, e.g., “opinput:0”.

You can also use this argument to refer to a USD file (e.g., “/path/to/file.usd”), or to another LOP node’s cooked stage using the `op:` as the path prefix (e.g., “op:/stage/lop_node”).

`primpath`

The path to the primitive.

`name`

Primvar name (without namespace).

Returns

The primvar’s element size.

Examples

## examples

```vex
// Get the element size of a primvar on the cube primitive or its ancestor.
int element_size = usd_iprimvarelementsize(0, "/geo/cube", "primvar_name");

```
