---
title: usd_relationshipnames
order: 118
---
| Since | 18.0 |
| --- | --- |

`string [] usd_relationshipnames(<stage>stage, string primpath)`

This function returns the relationship names that are available on the given primitive.

Show/hide arguments

`<stage>`

When running in the context of a node (such as a wrangle LOP), this argument can be an integer representing the input number (starting at 0) to read the stage from. The integer is equivalent to the string form referencing a particular input, e.g., “opinput:0”.

You can also use this argument to refer to a USD file (e.g., “/path/to/file.usd”), or to another LOP node’s cooked stage using the `op:` as the path prefix (e.g., “op:/stage/lop_node”).

`primpath`

The path to the primitive.

Returns

String array containing the names of the primitive’s relationships.

Examples

## examples

```vex
// Get the relationship names from the primitive.
string relationship_names[] = usd_relationshipnames(0, "/geo/cube");

```
