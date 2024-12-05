---
title: osd_limitsurface
order: 3
---
`osd_limitsurface` evaluates the point attribute in the geometry specified as a
subdivision surface.

For vertex attributes, use [osd_limitsurfacevertex](osd_limitsurfacevertex.html "Evaluates a vertex attribute at the subdivision limit surface using Open Subdiv.").

`int  osd_limitsurface(<geometry>geometry, string attrib_name, int patch_id, float u, float v, <type>&result)`

`int  osd_limitsurface(<geometry>geometry, string attrib_name, int patch_id, float u, float v, float &result[])`

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

`&result`

The computed attribute value is stored in the variable you pass to this argument.
The type of the variable should match the type of the attribute you are reading.

Returns

`1` if computing the attribute was successful, `0` if it failed.

Possible reasons for failure are:

- The geometry contains no polygons or the topology can’t be converted using Open Subdiv
- The attribute doesn’t exist on the input geometry.
- The attribute size/type doesn’t match the VEX type of the `result` argument.

Examples

## examples

Generate a point cloud on the limit surface of a subdivision mesh.

```vex
int npatches = osd_patchcount(file);
for (int patch = 0; patch < npatches; patch++)
{
    for (int v = 0; v < 100; v++)
    {
        vector P;
        if (osd_limitsurface(file, "P", patch, nrandom(), nrandom(), P))
        {
            int ptid = addpoint(geohandle, P);
        }
    }
}

```
