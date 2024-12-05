---
title: opprerawparmtransform
order: 32
---
| Since | 18.0 |
| --- | --- |

`matrix  opprerawparmtransform(string path)`

`matrix  opprerawparmtransform(string path, float time)`

`matrix  opprerawparmtransform(int opid)`

`matrix  opprerawparmtransform(int opid, float time)`

`matrix  opprerawparmtransform(int opid, int trsorder, int xyzorder, int mask)`

Returns the pre and raw parm transform associated with an OP. If the specified OP has no
transform associated with it (such as a COP), then an identity matrix is returned. It is possible to specify the time at which to evaluate the transform (in seconds, not frames).

The raw parameter transform is built from the transform parameters and does not include the effect of the CHOP IK solver.

Note
The op: syntax can be used to simulate this behavior using the standard transform functions.
