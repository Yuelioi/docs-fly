---
title: chnumchan
order: 11
---
| Context(s) | [chop](../contexts/chop.html) |
| --- | --- |

`int  chnumchan()`

Uses `-1` for `opinput`

`int  chnumchan(int opinput)`

Returns the number of channels in the input specified.

Show/hide arguments

`opinput`

The input number to read from, starting from 0. For example, the first input is 0, the second input is 1, and so on.

If you specify `-1`, the function uses the current CHOP node or input `0` if it is connected.
