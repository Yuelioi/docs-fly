---
title: chstart
order: 26
---
| Context(s) | [chop](../contexts/chop.html) |
| --- | --- |

`int  chstart()`

Uses `-1` as the `opinput`.

`int  chstart(int opinput)`

Returns the index of the first sample in the channel data in the given CHOP input.

Show/hide arguments

`opinput`

The input number to read from, starting from 0. For example, the first input is 0, the second input is 1, and so on.

If you specify `-1`, the function uses the current CHOP node or input `0` if it is connected.

To get the start frame, use [chstartf](chstartf.html "Returns the frame corresponding to the first sample of the input
specified."). To get the start time in seconds, use [chstartt](chstartt.html "Returns the time corresponding to the first sample of the input
specified.").
