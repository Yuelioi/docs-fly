---
title: opparmtransform
order: 29
---
`matrix  opparmtransform(string path)`

`matrix  opparmtransform(string path, float time)`

`matrix  opparmtransform(int opid)`

`matrix  opparmtransform(int opid, float time)`

Returns the parm transform associated with an OP. If the specified OP has no
transform associated with it (such as a COP), then an identity matrix is returned. It is possible to specify the time at which to evaluate the transform (in seconds, not frames).

Note
The op: syntax can be used to simulate this behavior using the standard transform functions.
