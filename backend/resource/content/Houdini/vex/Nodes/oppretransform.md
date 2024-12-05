---
title: oppretransform
order: 33
---
`matrix  oppretransform(string path)`

`matrix  oppretransform(string path, float time)`

`matrix  oppretransform(int opid)`

`matrix  oppretransform(int opid, float time)`

Returns the transform associated with an OP. If the OP specified has no
transform associated with it (for example a COP), then an identity
matrix is returned. It is possible to specify the time to evaluate the
transform at (in seconds, not frames).

Note
The op: syntax can be used to simulate this behavior using the standard transform functions.
