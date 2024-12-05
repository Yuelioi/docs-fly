---
title: chremoveattr
order: 19
---
| Context(s) | [chop](../contexts/chop.html) |
| --- | --- |

`int  chremoveattr(string attrclass, string attrname)`

`int  chremoveattr(string attrclass, string attrnames[])`

`int  chremoveattr(string attrname)`

`int  chremoveattr(string attrnames[])`

CHOP attributes store metadata on clips, channels, samples, or channel/sample pairs.

This function removes CHOP attribute.

Show/hide arguments

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

Signatures that don’t have this parameter act the same as if you passed the empty string.

`attrname`

The name of the attribute to remove.

`attrnames`

The names of the attributes to remove.

Returns

Returns `1` if the operation succeeded, or `0` otherwise.
