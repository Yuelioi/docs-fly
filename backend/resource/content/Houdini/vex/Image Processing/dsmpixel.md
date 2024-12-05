---
title: dsmpixel
order: 9
---
`int  dsmpixel(string map, string channel, int x, int y, <type>&values[])`

Each pixel of a deep shadow map channel has multiple values. This function
will extract the list of the values associated with the channel for the given
pixel.

Returns the number of values in the deep pixel (or -1 on failure).

A DSM will always have `Pz` and `Of` channels. The `Pz` channel stores the
z-depth associated with each record. The `Of` channel stores the opacity.
