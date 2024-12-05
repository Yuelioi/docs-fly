---
title: usd_collectionexpansionrule
order: 34
---
| Since | 18.0 |
| --- | --- |

`string  usd_collectionexpansionrule(<stage>stage, string collectionpath)`

This function returns the collection’s expansion rule.

USD supports a few standard expansion rules

- `explicitOnly` - only paths in the include list and not in the exclude list belong to the collection
- `expandPrims` - all the primitives at or below the includes (but not excludes) belong to the collection
- `expanPrimsAndProperties` - like `expandPrims` but also includes properties of matched primitives

Show/hide arguments

`<stage>`

When running in the context of a node (such as a wrangle LOP), this argument can be an integer representing the input number (starting at 0) to read the stage from. The integer is equivalent to the string form referencing a particular input, e.g., “opinput:0”.

You can also use this argument to refer to a USD file (e.g., “/path/to/file.usd”), or to another LOP node’s cooked stage using the `op:` as the path prefix (e.g., “op:/stage/lop_node”).

`collectionpath`

The path to the collection.

Returns

The collection’s expansion rule.

Examples

## examples

```vex
// Get collection's expansion rule.
string collection_path = usd_makecollectionpath(0, "/geo/cube", "some_collection");
string expansion_rule  = usd_collectionexpansionrule(0, collection_path);

```
