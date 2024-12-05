---
title: usd_variantselection
order: 152
---
| Since | 17.5 |
| --- | --- |

`string  usd_variantselection(<stage>stage, string primpath, string variantset)`

This function returns the current variant in a given variant set on a given primitive.

Show/hide arguments

`<stage>`

When running in the context of a node (such as a wrangle LOP), this argument can be an integer representing the input number (starting at 0) to read the stage from. The integer is equivalent to the string form referencing a particular input, e.g., “opinput:0”.

You can also use this argument to refer to a USD file (e.g., “/path/to/file.usd”), or to another LOP node’s cooked stage using the `op:` as the path prefix (e.g., “op:/stage/lop_node”).

`primpath`

The path to the primitive.

`variantset`

The variant set name.

Returns

The currently selected variant in a given variant set on a given primitive.

Examples

## examples

```vex
// Get the currently selected variant in the variant set "shapes" on a "shape_shifter" primitive.
string selected_variant = usd_variantselection(0, "/geo/shape_shifter", "shapes");

```
