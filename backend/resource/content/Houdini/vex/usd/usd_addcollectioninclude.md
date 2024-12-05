---
title: usd_addcollectioninclude
order: 3
---
| Since | 18.0 |
| --- | --- |

`int  usd_addcollectioninclude(int stagehandle, string collectionpath, string path)`

This function includes the object in the collection. This is usually achieved by adding an explicit path to the collection’s include list, but it may just remove a path from the collection’s exclude list, if it’s sufficient.

Show/hide arguments

`stagehandle`

A handle to the stage to write to. Currently the only valid value is `0`, which means the current stage in a node. (This argument may be used in the future to allow writing to other stages.)

`collectionpath`

The path to the collection.

`path`

The path to the object. I.e, a primitive, an attribute, or a relationship.

Returns

The value of `stagehandle` on success or `-1` on failure.

Examples

## examples

```vex
// Include sphere4 in cube's collection.
string collection_path = usd_makecollectionpath(0, "/geo/cube", "some_collection");
usd_addcollectioninclude(0, collection_path, "/geo/sphere4");

```
