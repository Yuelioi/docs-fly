---
title: cregiondeformtransform
order: 22
---
| Since | 18.0 |
| --- | --- |

`matrix  cregiondeformtransform(string path)`

`matrix  cregiondeformtransform(string path, float time)`

`matrix  cregiondeformtransform(int op_id)`

`matrix  cregiondeformtransform(int op_id, float time)`

Returns the deform transform associated with a Capture Region SOP.
The transform is built from the parameters of the SOP without cooking the SOP.
It is possible to specify the time to evaluate the transform at (in seconds, not frames).

Note
The op: syntax can be used to simulate this behavior using the standard transform functions.
