---
title: albedo
order: 1
---
`vector  albedo(bsdf b, ...)`

`vector  albedo(bsdf b, int mask, ...)`

`vector  albedo(bsdf b, vector viewer, ...)`

`vector  albedo(bsdf b, vector viewer, int mask, ...)`

Show/hide arguments

`viewer`

Vector toward viewer.

`mask`

A bitmask composed from values representing different shading components.

See [bouncemask](bouncemask.html) for information on component label bitmasks.
