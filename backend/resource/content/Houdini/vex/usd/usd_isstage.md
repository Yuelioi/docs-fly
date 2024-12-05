---
title: usd_isstage
order: 80
---
| Since | 17.5 |
| --- | --- |

`int  usd_isstage(<stage>stage)`

This function verifies whether a given input contains a valid USD stage. If so, all other USD functions will be able to access it to perform queries and actions. Otherwise, they will fail.

Show/hide arguments

`<stage>`

When running in the context of a node (such as a wrangle LOP), this argument can be an integer representing the input number (starting at 0) to read the stage from. The integer is equivalent to the string form referencing a particular input, e.g., “opinput:0”.

You can also use this argument to refer to a USD file (e.g., “/path/to/file.usd”), or to another LOP node’s cooked stage using the `op:` as the path prefix (e.g., “op:/stage/lop_node”).

Returns

1 if the stage is valid, and 0 otherwise.

Examples

## examples

```vex
// Check if the first input has a valid stage.
int is_valid_stage_on_first_input = usd_isstage(0);

```
