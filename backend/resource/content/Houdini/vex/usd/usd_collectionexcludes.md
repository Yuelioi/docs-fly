---
title: usd_collectionexcludes
order: 33
---
| Since | 18.0 |
| --- | --- |

`string [] usd_collectionexcludes(<stage>stage, string collectionpath)`

This function returns the collection’s exclude list.

Show/hide arguments

`<stage>`

When running in the context of a node (such as a wrangle LOP), this argument can be an integer representing the input number (starting at 0) to read the stage from. The integer is equivalent to the string form referencing a particular input, e.g., “opinput:0”.

You can also use this argument to refer to a USD file (e.g., “/path/to/file.usd”), or to another LOP node’s cooked stage using the `op:` as the path prefix (e.g., “op:/stage/lop_node”).

`collectionpath`

The path to the collection.

Returns

The collection’s exclude list.

Examples

## examples

```vex
// Get collection's exclude list.
string collection_path = usd_makecollectionpath(0, "/geo/cube", "some_collection");
string exclude_list[]  = usd_collectionexcludes(0, collection_path);

```
