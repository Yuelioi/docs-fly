---
title: oprawparmtransform
order: 34
---
| Since | 18.0 |
| --- | --- |

`matrix  oprawparmtransform(string path)`

`matrix  oprawparmtransform(string path, float time)`

`matrix  oprawparmtransform(int opid)`

`matrix  oprawparmtransform(int opid, float time)`

`matrix  oprawparmtransform(int opid, int trsorder, int xyzorder, int mask)`

Returns the raw parm transform associated with an OP. If the specified OP has no
transform associated with it (such as a COP), then an identity matrix is returned. It is possible to specify the time at which to evaluate the transform (in seconds, not frames).

The raw parameter transform is built from the transform parameters and does not include the effect of the CHOP IK solver.

Note
The op: syntax can be used to simulate this behavior using the standard transform functions.
