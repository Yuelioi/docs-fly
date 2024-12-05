---
title: chop
order: 12
---
| Since | 17.0 |
| --- | --- |

`float  chop(string filename, int|stringchannel, float|intsample)`

`vector2  chop(string filename, int|stringchannel, float|intsample)`

`vector  chop(string filename, int|stringchannel, float|intsample)`

`vector4  chop(string filename, int|stringchannel, float|intsample)`

`matrix  chop(string filename, int|stringchannel, float|intsample)`

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

Returns

The value of a channel at the specified sample in a CHOP node.
