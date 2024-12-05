---
title: optransform
order: 35
---
`matrix  optransform(string path)`

`matrix  optransform(string path, float time)`

`matrix  optransform(int op_id)`

`matrix  optransform(int op_id, float time)`

Returns the transform associated with an OP. If the OP specified has no
transform associated with it (for example a COP), then an identity
matrix is returned. It is possible to specify the time to evaluate the
transform at (in seconds, not frames).

NOTE: The op: syntax can be used to simulate this behavior using the
standard transform functions.
