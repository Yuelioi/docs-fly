---
title: metaweight
order: 5
---
`float  metaweight(<geometry>geometry, vector p)`

Returns the metaweight of the geometry at position p.
This is the result of evaluating all the metaballs in the
geometry at that position. Usually this is the sum of their
values, but meta expressions can change that.
