---
title: chattrnames
order: 3
---
| Context(s) | [chop](../contexts/chop.html) |
| --- | --- |

`string [] chattrnames(int opinput, string attribclass)`

`string [] chattrnames(string attrclass)`

Returns all the CHOP attribute names of a given attribute class from a CHOP input.

CHOP attributes store metadata on clips, channels, samples, or channel/sample pairs.

Show/hide arguments

`opinput`

The input number to read from, starting from 0. For example, the first input is 0, the second input is 1, and so on.

`attribclass`

The “level” of the attribute:

`"clip"`

An attribute on a whole clip.

`"channel"`

An attribute on a whole channel.

`"sample"`

An attribute on a sample (across all channels).

`"channelsample"`

An attribute on a specific channel/sample pair.

`""`

Pass an empty string to have the function figure out the class based on the other arguments.

Returns

The attribute names as a string array.
