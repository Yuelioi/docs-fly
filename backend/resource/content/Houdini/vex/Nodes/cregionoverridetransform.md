---
title: cregionoverridetransform
order: 23
---
| Since | 18.0 |
| --- | --- |

`matrix  cregionoverridetransform(string path)`

`matrix  cregionoverridetransform(string path, float time)`

`matrix  cregionoverridetransform(int op_id)`

`matrix  cregionoverridetransform(int op_id, float time)`

Returns the capture or deform transform associated with a Capture Region SOP based on the global capture override flag.
The transform is built from the parameters of the SOP without cooking the SOP.
It is possible to specify the time to evaluate the transform at (in seconds, not frames).

Note
The op: syntax can be used to simulate this behavior using the standard transform functions.
