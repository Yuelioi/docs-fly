---
title: usd_isindexediprimvar
order: 70
---
| Since | 19.0 |
| --- | --- |

`int  usd_isindexediprimvar(<stage>stage, string primpath, string name)`

This function checks whether the given primvar is indexed, if it’s found directly on the given primitive or is inherited from primitive’s ancestor.

Some primvars may contain a compacted array of unique values, and an additional array of indices into the value array. They are called indexed primvars. The length of the value array depends on the number of unique elements, but the length of the index array corresponds to the number of entities the primvar applies to.

Show/hide arguments

`<stage>`

When running in the context of a node (such as a wrangle LOP), this argument can be an integer representing the input number (starting at 0) to read the stage from. The integer is equivalent to the string form referencing a particular input, e.g., “opinput:0”.

You can also use this argument to refer to a USD file (e.g., “/path/to/file.usd”), or to another LOP node’s cooked stage using the `op:` as the path prefix (e.g., “op:/stage/lop_node”).

`primpath`

The path to the primitive.

`name`

Primvar name (without namespace).

Returns

`1` if the primvar exists and is indexed, or `0` otherwise.

Examples

## examples

```vex
// Check if primvar "some_primvar" is indexed on sphere or its ancestor.
int is_indexed = usd_isindexedprimvar(0, "/geometry/sphere", "some_primvar");

```
