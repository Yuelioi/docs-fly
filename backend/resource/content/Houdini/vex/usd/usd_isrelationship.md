---
title: usd_isrelationship
order: 79
---
| Since | 18.0 |
| --- | --- |

`int  usd_isrelationship(<stage>stage, string primpath, string name)`

This function checks whether the primitive has a relationship of a given name.

Show/hide arguments

`<stage>`

When running in the context of a node (such as a wrangle LOP), this argument can be an integer representing the input number (starting at 0) to read the stage from. The integer is equivalent to the string form referencing a particular input, e.g., “opinput:0”.

You can also use this argument to refer to a USD file (e.g., “/path/to/file.usd”), or to another LOP node’s cooked stage using the `op:` as the path prefix (e.g., “op:/stage/lop_node”).

`primpath`

The path to the primitive.

`name`

The relationship name.

Returns

`1` if the primitive has such a relationship, or `0` otherwise.

Examples

## examples

```vex
// Check if the cube has a relationship "some_relationship".
int is_valid_relationship = usd_isrelationship(0, "/geo/cube", "some_relationship");

```
