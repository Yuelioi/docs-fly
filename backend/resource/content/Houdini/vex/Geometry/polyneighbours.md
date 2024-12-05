---
title: polyneighbours
order: 24
---
| Since | 17.0 |
| --- | --- |

`int [] polyneighbours(<geometry>geometry, int primnum)`

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

This returns an array of primitive numbers of polygons that share an
edge with this polygon. It uses the Half-Edge data structure so will
work with geometry that supports that (ie, polygons)
Examples

## examples

This is roughly equivalent to the following code:

```vex
int[] polyneighbours(const string opname; const int primnum)
{
    int result[] = {};

    int start = primhedge(opname, primnum);

    for (int hedge = start; hedge != -1; )
    {
        for (int nh = hedge_nextequiv(opname, hedge);
             nh != hedge;
             nh = hedge_nextequiv(opname, nh))
        {
            int prim = hedge_prim(opname, nh);
            if (prim != -1 && prim != primnum)
            {
                append(result, prim);
            }
        }
        hedge = hedge_next(opname, hedge);
        if (hedge == start)
            break;
    }

    return result;
}

```
