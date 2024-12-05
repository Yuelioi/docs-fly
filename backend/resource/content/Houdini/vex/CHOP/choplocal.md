---
title: choplocal
order: 13
---
| Since | 17.5 |
| --- | --- |

`matrix  choplocal(string filename, int|stringchannel, float|intsample)`

Read a sample from the local transform channel at the given index.

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
