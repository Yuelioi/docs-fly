---
title: cregioncapturetransform
order: 21
---
| Since | 18.0 |
| --- | --- |

`matrix  cregioncapturetransform(string path)`

`matrix  cregioncapturetransform(string path, float time)`

`matrix  cregioncapturetransform(int op_id)`

`matrix  cregioncapturetransform(int op_id, float time)`

Returns the capture transform associated with a Capture Region SOP.
The transform is built from the parameters of the SOP without cooking the SOP.
It is possible to specify the time to evaluate the transform at (in seconds, not frames).

Note
The op: syntax can be used to simulate this behavior using the standard transform functions.
