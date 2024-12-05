---
title: chrampderiv
order: 15
---
| Since | 18.0 |
| --- | --- |

`float  chrampderiv(string channel, float ramppos)`

`float  chrampderiv(string channel, float ramppos, float time)`

`vector  chrampderiv(string channel, float ramppos)`

`vector  chrampderiv(string channel, float ramppos, float time)`

Evaluates the derivative of a parm parameter with respect to position.

The ramppos is where on the ramp to evaluate. The ramppos is clamped to the range `[0,1]`

The time parameter can be used if the ramp is animated to evaluate
at other than the current time.
Note that this does not take a derivative with respect to time.
