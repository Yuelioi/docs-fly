---
title: usd_applyapi
order: 15
---
| Since | 20.0 |
| --- | --- |

`int  usd_applyapi(int stagehandle, string primpath, string apischemaname)`

This function applies an API schema to a primitive. This function will only work for single-apply API schemas (such as GeomModelAPI), not for multiple-apply API schemas (such as CollectionAPI).

Show/hide arguments

`stagehandle`

A handle to the stage to write to. Currently the only valid value is `0`, which means the current stage in a node. (This argument may be used in the future to allow writing to other stages.)

`primpath`

The path to the primitive.

`apischemaname`

API schema name.

Returns

The value of `stagehandle` on success, or `-1` on failure.
