---
title: expandvertexgroup
order: 4
---
| Since | 17.0 |
| --- | --- |

`int [] expandvertexgroup(<geometry>geometry, string groupname)`

`int [] expandvertexgroup(<geometry>geometry, string groupname, string mode)`

This can use ad-hoc groups, like `0v3 1v2`.
This uses the SOP group naming convention, in particular that an empty string means *all*.

`mode` can be `ordered`, `unordered` or `split`.
`ordered` is the default mode and will return numbers in the order of appearance in the string, but only for numbers. The order won’t be kept when using expressions such as `@Cd.x>0.5`. The same number won’t appear twice in returned array.
`unordered` mode returns the resolved group following sorted point numbers order.
`split` mode starts by splitting the `groupname` string on `@` characters and then does one resolution per sub string. The order is kept between the sub strings, but will fallback to unordered when resolving a group expression. This same number can appear multiple time when resolving using this mode.
