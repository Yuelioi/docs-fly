---
title: sample_cdf
order: 7
---
| On this page | * [Overview](#overview) * [Usage](#usage) |
| --- | --- |

Overview

## overview

Use this function to sample values from a CDF array created with the [create_cdf](create_cdf.html "Creates a cumulative distribution function (CDF) from an array of probability density function (PDF) values.") function. See [create_cdf](create_cdf.html "Creates a cumulative distribution function (CDF) from an array of probability density function (PDF) values.") for more information.

See the [create_cdf](create_cdf.html "Creates a cumulative distribution function (CDF) from an array of probability density function (PDF) values.") function docs for example code.

Usage

## usage

`int  sample_cdf(float cdf[], float uniform_rand)`

Returns the index of the sampled CDF.

`void  sample_cdf(float cdf[], float uniform_rand, int &index, float &x)`

Writes the index of the sampled CDF and value to output arguments.

`void  sample_cdf(float cdf[], float uniform_rand, int &index, float &x, float &pdf)`

Writes the index of the sampled CDF, the sampled value, and the corresponding PDF to output arguments.

Show/hide arguments

`cdf`

The CDF to sample from (create this using [create_cdf](create_cdf.html "Creates a cumulative distribution function (CDF) from an array of probability density function (PDF) values.")).

`uniform_rand`

A uniform random variable (must be in range 0 to 1).

`&index`

Outputs the index of the sampled CDF element.

`&x`

Outputs the value of the sampled CDF element.

`&pdf`

Outputs the PDF of the sampled CDF element.

Returns

The first form returns the index of the sampled value. The other forms write the index to an output argument instead.
