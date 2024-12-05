---
title: usd_addcollectionexclude
order: 2
---
| Since | 18.0 |
| --- | --- |

`int  usd_addcollectionexclude(int stagehandle, string collectionpath, string path)`

This function excludes the object from the collection. This is usually achieved by adding an explicit path to the collection’s exclude list, but it may just remove a path from the collection’s include list, if it’s sufficient.

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
// Exclude sphere3 from cube's collection.
string collection_path = usd_makecollectionpath(0, "/geo/cube", "some_collection");
usd_addcollectionexclude(0, collection_path, "/geo/sphere3");

```
