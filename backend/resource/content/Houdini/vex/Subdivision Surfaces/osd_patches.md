---
title: osd_patches
order: 8
---
`int [] osd_patches(<geometry>geometry, int face_id)`

Each face in a subdivision hull may create one or more patches. This function lists the patch ids for a corresponding face.

This is implemented using the following algorithm:

```vex
int []
osd_patches(const string file; const face_id)
{
    int        patches[] = {};
    int        first = osd_firstpatch(file, face_id);
    if (first >= 0)
    {
        int        npatches = osd_patchcount(file, face_id);
        for (int i = 0; i < npatches; i++)
            append(patches, first+i);
    }
    return patches;
}

```
