---
title: osd_limitsurfacevertex
order: 4
---
This is like [osd_limitsurface](osd_limitsurface.html "Evaluates a point attribute at the subdivision limit surface using Open Subdiv.") but for vertex attributes instead of point attributes.
See [osd_limitsurface](osd_limitsurface.html "Evaluates a point attribute at the subdivision limit surface using Open Subdiv.") for more information.

`int  osd_limitsurfacevertex(<geometry>geometry, string attrib_name, int face_id, float u, float v, <type>&result)`

`int  osd_limitsurfacevertex(<geometry>geometry, string attrib_name, int face_id, float u, float v, float &result[])`

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.
