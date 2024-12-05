---
title: usd_iscollection
order: 68
---
| Since | 18.0 |
| --- | --- |

`int  usd_iscollection(<stage>stage, string collectionpath)`

This function checks whether the given path points to an existing collection.

Show/hide arguments

`<stage>`

When running in the context of a node (such as a wrangle LOP), this argument can be an integer representing the input number (starting at 0) to read the stage from. The integer is equivalent to the string form referencing a particular input, e.g., “opinput:0”.

You can also use this argument to refer to a USD file (e.g., “/path/to/file.usd”), or to another LOP node’s cooked stage using the `op:` as the path prefix (e.g., “op:/stage/lop_node”).

`collectionpath`

The path to the collection.

Returns

`1` if the path points to an existing collection, or `0` otherwise.

Examples

## examples

```vex
// Check if cube has a collection "some_collection".
string collection_path = usd_makecollectionpath(0, "/geo/cube", "some_collection");
int is_collection_existing = usd_iscollection(0, collection_path);

```
