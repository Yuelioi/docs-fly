---
title: hscript_turb
order: 17
---
`float  hscript_turb(vector pos, int depth)`

Matches the output of [turb](../../expressions/turb.html "Generates spatially coherent 3D noise."). This function can be useful if you convert a workflow to VEX, or have VEX work in tandem with Hscript expressions, and need the turbulence to look the same as in an expression.

Show/hide arguments

`pos`

Position at which to sample the turbulent noise.

`depth`

Number of fractal iterations of noise.

Returns

The range is usually within -1 to 1, but can exceed it depending on the depth. The maximum range is -2 to 2 for high depths.
