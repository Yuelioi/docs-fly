---
title: ramp_lookup
order: 13
---
| Since | 18.5 |
| --- | --- |

`float  ramp_lookup(float pos, string ramp)`

`vector  ramp_lookup(float pos, string ramp)`

`float  ramp_lookup(float pos, string basis[], float key[], float val[])`

`vector  ramp_lookup(float pos, string basis[], float key[], vector val[])`

Evaluates a provided ramp at the given position. The ramp can be an encoded string or a triple of basis, key, and value arrays.
