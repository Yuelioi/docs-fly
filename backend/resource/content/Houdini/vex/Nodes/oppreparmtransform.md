---
title: oppreparmtransform
order: 31
---
| Since | 18.0 |
| --- | --- |

`matrix  oppreparmtransform(string path)`

`matrix  oppreparmtransform(string path, float time)`

`matrix  oppreparmtransform(int opid)`

`matrix  oppreparmtransform(int opid, float time)`

Returns the pre and parm transform associated with an OP. If the specified OP has no
transform associated with it (such as a COP), then an identity matrix is returned. It is possible to specify the time at which to evaluate the transform (in seconds, not frames).

Note
The op: syntax can be used to simulate this behavior using the standard transform functions.
