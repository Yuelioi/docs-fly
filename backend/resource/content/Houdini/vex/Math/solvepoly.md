---
title: solvepoly
order: 76
---
`int  solvepoly(float coef[], float &roots[], int maxiter=0)`

Show/hide arguments

`coef`

An array of coefficients of the polynomial.

You must order the coefficients such that `coef[i]` should be `x^i`.
**This is reverse of the order you would write the polynomial out normally**.

`&roots`

The function overwrites this array with the real roots of the polynomial,
in ascending order.

Returns

The number of real roots.
