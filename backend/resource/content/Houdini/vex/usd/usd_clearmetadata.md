---
title: usd_clearmetadata
order: 29
---
| Since | 18.0 |
| --- | --- |

`int  usd_clearmetadata(int stagehandle, string path, string name)`

This function clears the value of the given metadata.

Show/hide arguments

`stagehandle`

A handle to the stage to write to. Currently the only valid value is `0`, which means the current stage in a node. (This argument may be used in the future to allow writing to other stages.)

`path`

The path to the object. I.e, a primitive, an attribute, or a relationship.

`name`

The metadata name.

The name can be namespaced to acces values inside (possibly nested) VtDictionaries, such as custom data dictionary, e.g., “customData:name” or “customData:name:subname”. For non-namespaced names, the object schema needs to declare a given metadata for it to be accessible, e.g., “active” or “documentation”.

Returns

The value of `stagehandle` on success, or `-1` on failure.

Examples

## examples

```vex
// Clear the metadata value.
usd_clearmetadata(0, "/geo/sphere", "customData:some_name");

```
