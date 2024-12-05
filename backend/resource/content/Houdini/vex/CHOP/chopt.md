---
title: chopt
order: 15
---
| Since | 17.0 |
| --- | --- |

`float  chopt(string filename, int|stringchannel, float|intsample, float time)`

`vector2  chopt(string filename, int|stringchannel, float|intsample, float time)`

`vector  chopt(string filename, int|stringchannel, float|intsample, float time)`

`vector4  chopt(string filename, int|stringchannel, float|intsample, float time)`

`matrix  chopt(string filename, int|stringchannel, float|intsample, float time)`

Read a sample from the channel at the given index.

Show/hide arguments

`filename`

The CHOP node path to query using the op: syntax.
Doesn’t support reading directly from CHOP files yet.

`channel`

The channel index or channel name to query.

`sample`

If this is fractional, the value is linearly interpolated from the
two nearest points.

`time`

Time in seconds when the CHOP node needs to be evaluated.

Returns

The value of a channel at the specified sample in a CHOP node.
