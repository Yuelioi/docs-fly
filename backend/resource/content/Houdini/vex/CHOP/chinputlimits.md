---
title: chinputlimits
order: 9
---
| Context(s) | [chop](../contexts/chop.html) |
| --- | --- |

`int  chinputlimits(int opinput, int channel, float &channel_min, float &channel_max)`

Show/hide arguments

`opinput`

CHOP Input index or -1 if omitted.

`channel`

When reading a `channel` or `channelsample` attribute, this is the index of the channel.
If you are reading a `clip` or `sample` attribute, use `-1` here.

`channel_min`

Computed minimum channel value;

`channel_max`

Computed maximum channel value;

Returns

Returns 1 on success or 0 on failure.
