---
title: osd_lookupface
order: 5
---
`int  osd_lookupface(<geometry>geometry, int patch_id, float patch_u, float patch_v, int &face_id, float &face_u, float &face_v)`

If you don’t specify a texture attribute, the function uses intrinsic polygon interpolants.

`int  osd_lookupface(<geometry>geometry, int patch_id, float patch_u, float patch_v, int &face_id, float &face_u, float &face_v, string attribute)`

If you specify a texture attribute, the function uses the UVs in that attribute to translate the patch coordinates onto the Houdini geometry.

Each polygon in a piece of geometry will generate one or more Catmull-Clark subdivision patches. A quadrilateral will generate a single patch, while a pentagon will generate five patches. This function helps map between the subdivision patch id and the Houdini polygon (face). The reverse function to map from face to patch is [osd_lookuppatch](osd_lookuppatch.html "Outputs the OSD patch and UV coordinates corresponding to the given coordinates on a Houdini polygon face.").

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

`patch_id`

The ID number of the OSD patch.

`patch_u`, `patch_v`

The coordinates in the subdivision patch to map onto a Houdini primitive. These values should be in the range 0 to 1.

`&face_id`

The function overwrites this variable with the Houdini primitive number of the corresponding face.

`&face_u`, `&face_v`

The function overwrites these variables with the corresponding U/V coordinates on the Houdini face. The output coordinates will have values in the range 0 to 1.

Returns

`1` on success or `0` on an error.

Examples

## examples

```vex
void
scatterOnLimitSurface(string file, texmap; int geo_handle; int npts)
{
    int npatches = osd_patchcount(file);
    for (int i = 0; i < npts; ++i)
    {
        int patch_id = nrandom() * npatches;
        float patch_s = nrandom();
        float patch_t = nrandom();
        int face_id;
        float face_u, face_v;
        if (osd_lookupface(file, patch_id, patch_s, patch_t, face_id, face_u, face_v, "uv"))
        {
            vector clr = texture(texmap, face_u, face_v);
            vector P;
            osd_limitsurface(file, "P", patch_id, patch_s, patch_t, P);
            int ptnum = addpoint(geo_handle, P);        // add a scattered point.
            if (ptnum >= 0)
            {
                addpointattrib(geo_handle, "Cd", clr);
                addpointattrib(geo_handle, "face_id", face_id);
            }
        }
    }
}

```
