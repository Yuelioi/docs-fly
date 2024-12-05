---
title: opparentbonetransform
order: 27
---
`matrix  opparentbonetransform(string path)`

`matrix  opparentbonetransform(string path, float time)`

`matrix  opparentbonetransform(int opid)`

`matrix  opparentbonetransform(int opid, float time)`

Returns the parent bone transform associated with an OP. If the specified OP has no
transform associated with it (such as a COP), then an identity matrix is returned. It is possible to specify the time at which to evaluate the transform (in seconds, not frames). Returns the transform at the root of the parent bone or the parent transform otherwise.

Note
The op: syntax can be used to simulate this behavior using the standard transform functions.
