---
title: opfullpath
order: 25
---
`string  opfullpath(string relative_path)`

This function returns the absolute path for a given relative path to the object being evaluated.

The function is currently only meaningful in Houdini.
Examples

## examples

- `opfullpath(".")` - The full path of the current node being evaluated
- `opfullpath("..")` - The full path of the current node’s parent
