---
title: chrename
order: 20
---
| Context(s) | [chop](../contexts/chop.html) |
| --- | --- |

`int  chrename(int channel_index, string new_name)`

`int  chrename(string channel_name, string new_name)`

This function renames a CHOP channel . It works only when iterating over Clip, Channel or Samples, not when iterating over ChannelSample.

Returns 1 if channel was renamed or 0 otherwise.

Show/hide arguments

`channel_index`

The channel index to rename.

`channel_name`

The channel to rename.

`new_name`

The new name.
