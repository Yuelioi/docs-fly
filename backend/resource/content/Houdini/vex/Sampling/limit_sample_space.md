---
title: limit_sample_space
order: 3
---
`float  limit_sample_space(float minu, float maxu, float u)`

`float  limit_sample_space(float maxu, float u)`

Show/hide arguments

`minu`

Minimum desired value of `u`. `minu` will be clamped to between 0 and 1.
If not specified, `minu` is 0.

`maxu`

Maximum desired value of `u`. `maxu` will be clamped to between 0 and 1.

`u`

Number between 0 and 1.

If `u` is outside of `[minu,maxu]`, `u` is wrapped in the space in such a
way that uniform random `u` in `[0,1)` will yield uniform random samples in
`[minu,maxu]` and returned. This avoids the extra samples at the bounds
of the range that clamping to the range would introduce. It also avoids the
changing of samples inside the range that fitting to the range would
introduce, i.e. if `u` is already in the range, the return value is exactly `u`.

However, this is much slower than fitting or clamping, so only use
it when both uniformity and consistency are needed. For example, it can be
useful in avoiding outliers in probability distributions without affecting
samples that are not outliers. It also introduces the issue that results
will no longer be monotone increasing with respect to `u`.
Fitting will often suffice for avoiding outliers too, at the expense
of slightly affecting samples that are not outliers.

To find `minu` and `maxu` given a `minvalue` and `maxvalue` of some probability
distribution, `minu = CDF(minvalue)` and `maxu = CDF(maxvalue)`, where `CDF`
is the cumulative distribution function (not inverse) of the probability
distribution. The versions of [sample_exponential](sample_exponential.html "Samples the exponential distribution."),
[sample_cauchy](sample_cauchy.html "Samples the Cauchy (Lorentz) distribution."), [sample_normal](sample_normal.html "Samples the normal (Gaussian) distribution."), [sample_lognormal](sample_lognormal.html "Samples the log-normal distribution based on parameters of the underlying normal distribution."), and
[sample_lognormal_by_median](sample_lognormal_by_median.html "Samples the log-normal distribution based on median and standard deviation.") that take a `minvalue` or `maxvalue` use fitting
instead of this limiting, because it maintains monotonicity, but this
function can be applied to `u` before sampling, in order to have better
consistency for samples in the range.
