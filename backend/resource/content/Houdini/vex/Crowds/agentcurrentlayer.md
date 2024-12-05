---
title: agentcurrentlayer
order: 22
---
Warning
This function has been deprecated. Use [agentcurrentlayers](agentcurrentlayers.html "Returns the names of an agent primitive’s current layers.") instead.

Returns an empty string if `prim` is out of range or is not an agent primitive.

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

`prim`

The primitive number.
