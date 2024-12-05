---
title: usd_relationshipforwardedtargets
order: 117
---
| Since | 18.0 |
| --- | --- |

`string [] usd_relationshipforwardedtargets(<stage>stage, string primpath, string name)`

This function returns the list of forwarded targets of a given relationship. This is a convenience function to expand all the nested relationships, since a target in a relationship can be another relationship.

Show/hide arguments

`<stage>`

When running in the context of a node (such as a wrangle LOP), this argument can be an integer representing the input number (starting at 0) to read the stage from. The integer is equivalent to the string form referencing a particular input, e.g., “opinput:0”.

You can also use this argument to refer to a USD file (e.g., “/path/to/file.usd”), or to another LOP node’s cooked stage using the `op:` as the path prefix (e.g., “op:/stage/lop_node”).

`primpath`

The path to the primitive.

`name`

The relationship name.

Returns

List of forwarded targets in a relationship.

Examples

## examples

```vex
// Get the list of forwarded targets in cube's "some_relationship" relationship.
string targets[] = usd_relationshipforwardedtargets(0, "/geo/cube", "some_relationship");

```
