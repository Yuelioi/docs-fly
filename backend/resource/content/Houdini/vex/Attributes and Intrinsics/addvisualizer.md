---
title: addvisualizer
order: 6
---
This function creates the `visualizer` detail string array attribute if it doesn’t exist, then appends the given visualizer string to it. If the visualizer string already exists in the array, the function does not add it again.

`int  addvisualizer(int geohandle, string op_url)`

Show/hide arguments

`geohandle`

A handle to the geometry to write to. Currently the only valid value is `0` or [geoself](geoself.html "Returns a handle to the current geometry."), which means the current geometry in a node. (This argument may be used in the future to allow writing to other geometries.)

`op_url`

A string in the form`"op:/path/to/node"`. The geometry will use that node’s visualizers.

- Houdini looks up the visualizers by the `op:` reference, so changes to the visualizers on the referenced nodes will be reflected in the visualization of the geometry.
