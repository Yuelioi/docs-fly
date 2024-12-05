---
title: usd_setcollectionincludes
order: 127
---
| Since | 18.0 |
| --- | --- |

`int  usd_setcollectionincludes(int stagehandle, string collectionpath, string includes[])`

This function sets the includes list on the collection.

Show/hide arguments

`stagehandle`

A handle to the stage to write to. Currently the only valid value is `0`, which means the current stage in a node. (This argument may be used in the future to allow writing to other stages.)

`collectionpath`

The path to the collection.

`includes`

A list of object paths to set as an includes list on the collection.

Returns

The value of `stagehandle` on success or `-1` on failure.

Examples

## examples

```vex
// Set the exludes list on the cube's collection.
string collection_path = usd_makecollectionpath(0, "/geo/cube", "some_collection");
usd_setcollectionincludes(0, collection_path, array("/geo/sphere1", "/geo/sphere2"));

```
