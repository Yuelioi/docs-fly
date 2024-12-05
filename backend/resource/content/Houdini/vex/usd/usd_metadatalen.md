---
title: usd_metadatalen
order: 94
---
| Since | 18.0 |
| --- | --- |

`int  usd_metadatalen(<stage>stage, string path, string name)`

This function returns the length of a given metadata.

For array metadata it is the length of the array, and for non-array metadata the length is 1.

Show/hide arguments

`<stage>`

When running in the context of a node (such as a wrangle LOP), this argument can be an integer representing the input number (starting at 0) to read the stage from. The integer is equivalent to the string form referencing a particular input, e.g., “opinput:0”.

You can also use this argument to refer to a USD file (e.g., “/path/to/file.usd”), or to another LOP node’s cooked stage using the `op:` as the path prefix (e.g., “op:/stage/lop_node”).

`path`

The path to the object. I.e, a primitive, an attribute, or a relationship.

`name`

The metadata name.

The name can be namespaced to acces values inside (possibly nested) VtDictionaries, such as custom data dictionary, e.g., “customData:name” or “customData:name:subname”. For non-namespaced names, the object schema needs to declare a given metadata for it to be accessible, e.g., “active” or “documentation”.

Returns

The length of the array metadata, or 1 if the metadata is not an array. Use [usd_isarraymetadata](usd_isarraymetadata.html "Checks if the given metadata is an array.") if you want to check whether the metadata is an array.

Examples

## examples

```vex
// Get the array length of metadata on the cube primitive.
int length = usd_metadatalen(0, "/geo/cube", "customData:name");

```
