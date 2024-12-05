---
title: chremove
order: 18
---
| Context(s) | [chop](../contexts/chop.html) |
| --- | --- |

`int  chremove(int channel_index)`

`int  chremove(int channel_indices[])`

`int  chremove(string channel_name)`

`int  chremove(string channel_names[])`

This function removes channels from a CHOP node. It works only when iterating over Clip, Channel or Samples, not when iterating over ChannelSample.

Returns 1 if all channels were removed or 0 otherwise.

Show/hide arguments

`channel_index`

The channel index to remove.

`channel_indices`

The array of channel indices to remove.

`channel_name`

The channel name to remove.

`channel_names`

The array of channel names to remove.
