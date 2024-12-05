---
title: pcimport
order: 14
---
This function is only valid while looping with `pciterate` or `pcunshaded`.

`int  pcimport(int handle, string channel_name, <type>&value)`

Imports data from the point cloud file into the given variable.

Show/hide arguments

`channel_name`

There are two special channel names you can import:

`point.number`

The number of the point being processed.

`point.distance`

The distance of the point being processed from the query point.
This is only available when iterating over unshaded points.

`value`

If the import succeeds the function overwrites this variable with the channel value.

Returns

`1` if the import succeeded or `0` if the import failed (usually due to the given channel name not existing).
