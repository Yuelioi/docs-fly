---
title: pcunshaded
order: 32
---
`int  pcunshaded(int handle, string channel_name)`

Like [pciterate](pciterate.html "This function can be used to iterate over all the points which were
found in the pcopen query."), this function can be used to iterate over
points which were found in a [pcopen](pcopen.html "Returns a handle to a point cloud file.") query. The first argument is
the handle returned by `pcopen`.

However, where `pciterate` iterates over all the points, this function
only iterates over points where the channel in channel_name has
not yet been written to.

The function returns 1 while there are points left in the iteration loop,
or 0 when there are no further points. This lets you use the function as
the condition in a [while loop](../statement.html).

Warnings:

- This function will not work correctly when used in multi-threaded OPs.
  It is not possible to nest `pcunshaded` or [pciterate](pciterate.html "This function can be used to iterate over all the points which were
  found in the pcopen query.")
  loops for the same handle. That is, for a single [pcopen](pcopen.html "Returns a handle to a point cloud file.")
  call, only one `pcunshaded` or [pciterate](pciterate.html "This function can be used to iterate over all the points which were
  found in the pcopen query.") loop may be
  entered.
- Computations involving derivatives inside `pcunshaded` loops may have
  slightly different results. If derivatives are required for variables
  which aren’t set by [pcimport](pcimport.html "Imports channel data from a point cloud inside a pciterate or a pcunshaded loop.") it may be better to
  pre-compute the derivatives before the `pcunshaded` loop is entered.
