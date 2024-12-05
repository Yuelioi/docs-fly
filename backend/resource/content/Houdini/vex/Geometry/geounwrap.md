---
title: geounwrap
order: 6
---
`string  geounwrap(<geometry>geometry, string unwrap_attribute)`

Returns an oppath string that will cause the file or geometry to be unwrapped inplace based on a vector attribute.
This function adds an “unwrap:attrname” prefix followed by the unwrap_attribute to the path.
path can be a filename, an oppath with the “op:” prefix or an opinput.
An input index can be supplied instead of a string.

Adding an “unwrap:attrname” prefix to an oppath will create a copy of the geometry and overwrite the point positions based on the unwrap attribute. The topology can change, if the attribute is a vertex attribute.

This then enables all the vex functions working on the point positions to work in a custom space such as UV space or Color space.
