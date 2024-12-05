---
title: chadd
order: 1
---
| Context(s) | [chop](../contexts/chop.html) |
| --- | --- |

`int  chadd(string channel_names)`

`int  chadd(string channel_names[])`

This function adds new channels to a CHOP node. It works only when iterating over Clip, Channel or Samples, not when iterating over ChannelSample. You can’t control the default channel values when adding channels using this function. You need to add another `Channel Wrangle` and compute the channel data inside it.

Returns 1 on success, 0 otherwise.

Show/hide arguments

`channel_names`

The names of the attributes to add. Can be an array or a space separated list of channel names.
