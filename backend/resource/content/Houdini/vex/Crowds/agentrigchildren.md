---
title: agentrigchildren
order: 35
---
`int [] agentrigchildren(<geometry>geometry, int prim, int transform)`

Returns a list of the direct children of the given transform.

Returns an empty array if `transform` is [out of range](agenttransformcount.html "Returns the number of transforms in an agent primitive’s rig."), `prim` is out of range, or `prim` is not an agent primitive.

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

`prim`

The primitive number.

`transform`

Index of a transform in the agent’s rig.

Examples

## examples

Visit all of the children of a given transform.

```vex
int[] queue = { transform };

while (len(queue) > 0) {
int i = removeindex(queue, 0);
printf("%d\n", i);

foreach (int child; agentrigchildren(0, @primnum, i))
push(queue, child);
}

```
