---
title: usd_setcollectionexpansionrule
order: 126
---
| Since | 18.0 |
| --- | --- |

`int  usd_setcollectionexpansionrule(int stagehandle, string collectionpath, string rule)`

This function sets the expansion rule on the collection.

Show/hide arguments

`stagehandle`

A handle to the stage to write to. Currently the only valid value is `0`, which means the current stage in a node. (This argument may be used in the future to allow writing to other stages.)

`collectionpath`

The path to the collection.

`rule`

The expansion rule to set on the collection.

USD supports a few standard expansion rules

- `explicitOnly` - only paths in the include list and not in the exclude list belong to the collection
- `expandPrims` - all the primitives at or below the includes (but not excludes) belong to the collection
- `expanPrimsAndProperties` - like `expandPrims` but also includes properties of matched primitives

Returns

The value of `stagehandle` on success or `-1` on failure.

Examples

## examples

```vex
// Set the expansion rule on the cube's collection.
string collection_path = usd_makecollectionpath(0, "/geo/cube", "some_collection");
usd_setcollectionexpansionrule(0, collection_foo, "explicitOnly");

```
