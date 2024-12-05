---
title: chname
order: 5
---
| Context(s) | [cop2](../contexts/cop2.html)  [chop](../contexts/chop.html) |
| --- | --- |

COPs

## cops

`string  chname(int plane_index, int chindex)`

Returns the name of the channel on the plane (for example, `"r"`, or `"x"`).

CHOPs

## chops

`string  chname(int channel_index)`

`string  chname(int opinput, int channel_index)`

Returns the name of the channel, for example `"tx"`.
To get a list of all channel names, use [chnames](chnames.html "Returns all the CHOP channel names of a given CHOP input.").

Show/hide arguments

`opinput`

The input number to read from, starting from 0. For example, the first input is 0, the second input is 1, and so on.

If you specify `-1`, the function uses the current CHOP node or input `0` if it is connected.
