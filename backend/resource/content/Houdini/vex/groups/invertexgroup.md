---
title: invertexgroup
order: 7
---
`int  invertexgroup(string filename, string groupname, int vertexnum)`

`int  invertexgroup(int opinput, string groupname, int vertexnum)`

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

`vertexnum`

The linear vertex number of the vertex to test.

To convert a primitive number and vertex number within that primitive to a linear vertex number for the `vertexnum` parameter, use the [vertexindex](vertexindex.html "Converts a primitive/vertex pair into a linear vertex.") function.

Returns

`1` if the group exists and the vertex is in the group, or `0` otherwise.

This can use ad-hoc groups, like `42p0-2`. It matches the SOP group naming
convention, in particular that an empty string means all vertices.
