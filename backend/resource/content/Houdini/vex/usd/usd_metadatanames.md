---
title: usd_metadatanames
order: 95
---
| Since | 18.0 |
| --- | --- |

`string [] usd_metadatanames(<stage>stage, string path)`

This function returns the metadata names that are available on the given USD object.

Show/hide arguments

`<stage>`

When running in the context of a node (such as a wrangle LOP), this argument can be an integer representing the input number (starting at 0) to read the stage from. The integer is equivalent to the string form referencing a particular input, e.g., “opinput:0”.

You can also use this argument to refer to a USD file (e.g., “/path/to/file.usd”), or to another LOP node’s cooked stage using the `op:` as the path prefix (e.g., “op:/stage/lop_node”).

`path`

The path to the object. I.e, a primitive, an attribute, or a relationship.

Returns

String array containing the names of the object’s metadata.

Examples

## examples

```vex
// Get the metadata names from the primitive.
string prim_metadata_names[] = usd_metadatanames(0, "/geo/sphere");

// Get the metadata names from the attribute.
string attrib_path = usd_makeattribpath(0, "/geo/sphere", "attrib_name");
string attrib_metadata_names[] = usd_metadatanames(0, attrib_path);

```
