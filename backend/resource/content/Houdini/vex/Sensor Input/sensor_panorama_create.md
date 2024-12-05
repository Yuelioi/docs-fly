---
title: sensor_panorama_create
order: 1
---
`int  sensor_panorama_create(float time, vector pos, int size, float near, float far, string candidateobj, string includeobj, string excludeobj, int uselit)`

This function will render the surrounding environment using the GL render and
provides a handle to use for querying the results.

Note
Because this needs to render the scene, it only works in interactive sessions of Houdini.

Show/hide arguments

`time`

The period in time when the render should be performed.

`pos`

The location in world space coordinates where the render should be performed.

`size`

The resolution of the performed render.

`near`

The near plane restriction.

`far`

The far plane restriction.

`candidateobj`

A bundle, group, or expression that represents what objects will be displayed if their display setting is enabled.

`includeobj`

A bundle, group, or expression that represents what objects will always be displayed.

`excludeobj`

A bundle, group, or expression that represents what objects will never be displayed.

`uselit`

Usually for AI purposes you want to not have any lighting as you are using
color as a key to differentiate actors. However, if you want to display
what a creature sees, lighting makes things more visually accurate.
