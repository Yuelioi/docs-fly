---
title: agentmetadata
order: 32
---
| Since | 18.5 |
| --- | --- |

`dict  agentmetadata(<geometry>geometry, int prim)`

Returns the shared [metadata dictionary](../../crowds/agents.html#metadata) from the agent’s definition.
Returns an empty dictionary if `prim` is out of range or is not an agent primitive.

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

`prim`

The primitive number.
