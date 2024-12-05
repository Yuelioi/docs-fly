---
title: usd_metadataelement
order: 93
---
| Since | 18.0 |
| --- | --- |

`<type> usd_metadataelement(<stage>stage, string path, string name, int index)`

This function returns an element value of of a given array metadata at a given index.

Show/hide arguments

`<stage>`

When running in the context of a node (such as a wrangle LOP), this argument can be an integer representing the input number (starting at 0) to read the stage from. The integer is equivalent to the string form referencing a particular input, e.g., “opinput:0”.

You can also use this argument to refer to a USD file (e.g., “/path/to/file.usd”), or to another LOP node’s cooked stage using the `op:` as the path prefix (e.g., “op:/stage/lop_node”).

`path`

The path to the object. I.e, a primitive, an attribute, or a relationship.

`name`

The metadata name.

The name can be namespaced to acces values inside (possibly nested) VtDictionaries, such as custom data dictionary, e.g., “customData:name” or “customData:name:subname”. For non-namespaced names, the object schema needs to declare a given metadata for it to be accessible, e.g., “active” or “documentation”.

`index`

The index into the array.

Returns

The value of an element in an existing array metadata, or zero/empty value if the metadata does not exist. Use [usd_ismetadata](usd_ismetadata.html "Checks if the primitive has metadata by the given name.") if you want to check whether the metadata exists.

Examples

## examples

```vex
// Get the value of an element at index 3 in the "foo:bar" array custom data.
string docs = usd_metadataelement(0, "/geo/cube", "customData:foo:bar", 3);

```
