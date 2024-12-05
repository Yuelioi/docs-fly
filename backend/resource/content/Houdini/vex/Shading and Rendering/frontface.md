---
title: frontface
order: 8
---
`vector  frontface(vector N, vector I)`

This form (which doesn’t take a reference vector) is only available in
the shading contexts, where the Ng variable is used.

`vector  frontface(vector N, vector I, vector Nref)`

If [dot](dot.html "Returns the dot product between the arguments.")(I, Nref) is less than zero, N will be negated.
