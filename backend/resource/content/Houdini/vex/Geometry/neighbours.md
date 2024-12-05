---
title: neighbours
order: 16
---
`int [] neighbours(<geometry>geometry, int ptnum)`

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

This is a simpler, array-based replacement for the combination of
[neighbourcount](neighbourcount.html "Returns the number of points that are connected to the specified point.") and [neighbour](neighbour.html "Returns the point number of the next point connected to a given point."). The array contains the numbers
of all points that share an edge with `ptnum`. The point numbers are in no particular order.
Examples

## examples

This is roughly equivalent to the following code:

```vex
int []
neighbours(int opinput, int ptnum)
{
    int     i, n;
    int     result[];
    n = neighbourcount(input, ptnum);
    resize(result, n);
    for (i = 0; i < n; i++)
        result[i] = neighbour(input, ptnum, i);
}

```
