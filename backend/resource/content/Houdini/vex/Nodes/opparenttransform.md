---
title: opparenttransform
order: 28
---
`matrix  opparenttransform(string path)`

`matrix  opparenttransform(string path, float time)`

`matrix  opparenttransform(int opid)`

`matrix  opparenttransform(int opid, float time)`

Returns the parent transform associated with an OP. If the specified OP has no
transform associated with it (such as a COP), then an identity matrix is returned. It is possible to specify the time at which to evaluate the transform (in seconds, not frames).

Note
The op: syntax can be used to simulate this behavior using the standard transform functions.
