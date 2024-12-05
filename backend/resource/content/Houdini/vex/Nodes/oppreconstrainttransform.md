---
title: oppreconstrainttransform
order: 30
---
`matrix  oppreconstrainttransform(string path)`

`matrix  oppreconstrainttransform(string path, float time)`

`matrix  oppreconstrainttransform(int opid)`

`matrix  oppreconstrainttransform(int opid, float time)`

Returns the preconstraint transform associated with an OP. If the specified OP has no
transform associated with it (such as a COP), then an identity matrix is returned. It is possible to specify the time at which to evaluate the transform (in seconds, not frames).

Note
The op: syntax can be used to simulate this behavior using the standard transform functions.
