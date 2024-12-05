---
title: pciterate
order: 22
---
`int  pciterate(int handle)`

This function can be used to iterate over all the points which were
found in a [pcopen](pcopen.html "Returns a handle to a point cloud file.") query. The first argument is the handle
returned by `pcopen`.
The function returns 1 while there are points left in the iteration loop,
or 0 when there are no further points. This lets you use the function as
the condition in a [while loop](../statement.html).

Warnings:

- It is not possible to nest pcunshaded or pciterate loops for the same
  handle. That is, for a single [pcopen](pcopen.html "Returns a handle to a point cloud file.") call, only one
  [pcunshaded](pcunshaded.html "Iterate over all of the points of a read-write channel which haven’t
  had any data written to the channel yet.") or `pciterate` loop may be entered.
- Computations involving derivatives inside [pcunshaded](pcunshaded.html "Iterate over all of the points of a read-write channel which haven’t
  had any data written to the channel yet.")
  loops may have slightly different results. If derivatives are required
  for variables which aren’t set by [pcimport](pcimport.html "Imports channel data from a point cloud inside a pciterate or a pcunshaded loop.") it may be
  better to pre-compute the derivatives before the
  [pcunshaded](pcunshaded.html "Iterate over all of the points of a read-write channel which haven’t
  had any data written to the channel yet.") loop is entered.
