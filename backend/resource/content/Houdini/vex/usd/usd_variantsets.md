---
title: usd_variantsets
order: 153
---
| Since | 17.5 |
| --- | --- |

`string [] usd_variantsets(<stage>stage, string primpath)`

This function returns the variant sets for the given primitive.

Show/hide arguments

`<stage>`

When running in the context of a node (such as a wrangle LOP), this argument can be an integer representing the input number (starting at 0) to read the stage from. The integer is equivalent to the string form referencing a particular input, e.g., “opinput:0”.

You can also use this argument to refer to a USD file (e.g., “/path/to/file.usd”), or to another LOP node’s cooked stage using the `op:` as the path prefix (e.g., “op:/stage/lop_node”).

`primpath`

The path to the primitive.

Returns

The variant set names available on a given primitive.

Examples

## examples

```vex
// Get the variant sets available on the "shape_shifter" primitive.
string variant_sets[] = usd_variantsets(0, "/geo/shape_shifter");

```
