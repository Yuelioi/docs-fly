---
title: usd_collectioncontains
order: 32
---
| Since | 18.0 |
| --- | --- |

`int  usd_collectioncontains(<stage>stage, string collectionpath, string path)`

This function returns `1` if the given objects belongs to the collection, otherwise returns `0`.

Show/hide arguments

`<stage>`

When running in the context of a node (such as a wrangle LOP), this argument can be an integer representing the input number (starting at 0) to read the stage from. The integer is equivalent to the string form referencing a particular input, e.g., “opinput:0”.

You can also use this argument to refer to a USD file (e.g., “/path/to/file.usd”), or to another LOP node’s cooked stage using the `op:` as the path prefix (e.g., “op:/stage/lop_node”).

`collectionpath`

The path to the collection.

`path`

The path to the object. I.e, a primitive, an attribute, or a relationship.

Returns

`1` if the given objects belongs to the collection, otherwise returns `0`.

Examples

## examples

```vex
// Check if sphere3 is in cube's collection.
string collection_path = usd_makecollectionpath(0, "/geo/cube", "some_collection");
int contains_sphere3 = usd_collectioncontains(0, collection_path, "/geo/sphere3");

```
