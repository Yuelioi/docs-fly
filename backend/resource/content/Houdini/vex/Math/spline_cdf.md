---
title: spline_cdf
order: 80
---
| On this page | * [Overview](#overview) * [Usage](#usage) |
| --- | --- |
| Since | 18.5 |

Overview

## overview

Whereas the [create_cdf](create_cdf.html "Creates a cumulative distribution function (CDF) from an array of probability density function (PDF) values.") function creates a CDF from sample values, this functions creates a CDF by randomly sampling a curve defined by a list of `values` and corresponding `positions` (similar to a Ramp parameter).

See the [create_cdf](create_cdf.html "Creates a cumulative distribution function (CDF) from an array of probability density function (PDF) values.") function docs for example code using a CDF.

Usage

## usage

`float [] spline_cdf(string bases[], float values[], float positions[], ...)`

Takes an array of bases, an array of key values, a corresponding array of key
positions, and an optionally an int, res (resolution), corresponding to the number of samples
to build the CDF with. Function samples the spline according to the resolution,
then creates and returns a CDF from said samples. Note that there is only support
for single dimension splines.

Show/hide arguments

`bases`

An array of strings describing how to interpret the corresponding `values`: each string can be `"constant"`, `"linear"`, `"cubic"` (or `"catmullrom"`, `"cspline"`), `"linearsolve"` (or `"solvelinear"`), or `"monotonecubic"`. See the [spline](spline.html "Samples a value along a polyline or spline curve.") function docs for information on how these options control the interpretation of the values.

"res",

`=128`

When building a CDF, some splines may require more samples to accurately represent the spline. The `res` (resolution) controls how many samples the function takes when constructing the CDF (and consequently the size of the CDF). The default is 128.

Returns

A float array representing a CDF (like the array returned by [create_cdf](create_cdf.html "Creates a cumulative distribution function (CDF) from an array of probability density function (PDF) values.")).
