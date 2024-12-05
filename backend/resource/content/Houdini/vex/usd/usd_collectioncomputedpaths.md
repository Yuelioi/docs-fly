---
title: usd_collectioncomputedpaths
order: 31
---
| Since | 18.0 |
| --- | --- |

`string [] usd_collectioncomputedpaths(<stage>stage, string collectionpath)`

This function returns the list of all objects that belong to the given collection.

Show/hide arguments

`<stage>`

When running in the context of a node (such as a wrangle LOP), this argument can be an integer representing the input number (starting at 0) to read the stage from. The integer is equivalent to the string form referencing a particular input, e.g., “opinput:0”.

You can also use this argument to refer to a USD file (e.g., “/path/to/file.usd”), or to another LOP node’s cooked stage using the `op:` as the path prefix (e.g., “op:/stage/lop_node”).

`collectionpath`

The path to the collection.

Returns

The list of all objects that belong to the given collection.

Examples

## examples

```vex
// Get all objects in cube's collection.
string collection_path = usd_makecollectionpath(0, "/geo/cube", "some_collection");
string members[] = usd_collectioncomputedpaths(0, collection_path);

```
