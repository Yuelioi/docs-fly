---
title: Iteration-Suites
order: 12
---

# 迭代函数套件

插件经常对图像中的所有像素进行迭代，对每个像素进行过滤。通过利用迭代套件，可以让 After Effects 把任务分到尽可能多的处理器上，利用特定硬件加速。

After Effects 还将自动管理进度报告和用户取消。

使用这些套件! 确保传递给这些迭代器回调的像素处理函数是可重入(re-entrant)的。

::: tip
2021 年 10 月的 SDK 更新增加了并发迭代线程的数量，最多可达可用的系统 CPU 核心，而不是以前硬编码的32个核心限制。
:::

## PF_Iterate8Suite1, PF_Iterate16Suite1, PF_IterateFloatSuite1

### iterate

遍历源图像的像素，改变并填充到目标图像中。

可以指定一个长方形像素区域来进行迭代；如果不，AE会在每个重叠像素上进行迭代。给出 refcon，这个函数会被调用，并加上当前像素的 x 和 y 坐标，以及该像素在源图像和目标图像中的指针。如果传递一个NULL源图像，它将在dst图像上进行迭代。此函数与质量无关。

不要依赖于以任何特定顺序遍历的像素。图像可能是不同CPU的子集，因此在AE处理时，请考虑所有参数（dst除外）都是只读的。此回调自动包括进度和中止检查，因此不要在像素函数中这样做。

```cpp
iterate(
  PF_InData       *in_data,
  A_long          progress_base,
  A_long          progress_final,
  PF_EffectWorld  *src,
  const PF_Rect   *area,
  void            *refcon,
  PF_Err (*pix_fn)(
    void      *refcon,
    A_long    x,
    A_long    y,
    PF_Pixel  *in,
    PF_Pixel  *out),
  PF_EffectWorld  *dst);
```

### iterate_origin

允许您指定从输入到输出的偏移量。例如，如果您的输出缓冲区小于输入缓冲区，则将 (in - data>output_origin_x, in_data>output_origin_y)` 作为原点，并将NULL作为区域，此函数将根据您的像素函数适当偏移src像素指针。

```cpp
iterate_origin(
  PF_InData     *in_data,
  A_long      progress_base,
  A_long        progress_final,
  PF_EffectWorld  *src,
  const PF_Rect  *area,
  const PF_Point  *origin,
  void       *refcon,
  PF_Err (*pix_fn)(
    void     *refcon,
    A_long    x,
    A_long    y,
    PF_Pixel   *in,
    PF_Pixel   *out),
  PF_EffectWorld  *dst);
```

### iterate_lut

仅限PF_Iterate8Suite。允许传递查找表（LUT）以进行迭代；可以为每个颜色通道传递相同或不同的LUT。如果没有传递LUT，则使用identity LUT。

```cpp
iterate_lut(
  PF_InData       *in_data,
  A_long        prog_base,
  A_long        prog_final,
  PF_EffectWorld    *src,
  const PF_Rect    *area,
  A_u_char       *a_lut0,
  A_u_char       *r_lut0,
  A_u_char       *g_lut0,
  A_u_char       *b_lut0,
  PF_EffectWorld  *dst);
```

### iterate_origin_non_clip_src

允许在源层和目标层交叉点之外的像素之间进行迭代。对于这些像素，您将被传递一个

PF_Pixel值{0,0,0,0}

```cpp
iterate_origin_non_clip_src(
  PF_InData       *in_data,
  A_long          progress_base,
  A_long          progress_final,
  PF_EffectWorld  *src,
  const PF_Rect   *area,
  const PF_Point  *origin,
  void            *refcon,
  PF_Err (*pix_fn)(
    void      *refcon,
    A_long    x,
    A_long    y,
    PF_Pixel  *in,
    PF_Pixel  *out),
  PF_EffectWorld  *dst);
```

### iterate_generic

仅限 PF_Iterate8Suite。如果想在每个可用CPU上执行一次操作，请使用本函数（传递`PF_Iterations_ONCE_PER_PROCESSOR`用于iterationsL）。仅从线程索引0调用中止和进度函数。

注意：可以迭代多个像素。在内部，我们将其用于基于行的图像处理，以及复杂序列数据的每个实体一次更新。

```cpp
iterate_generic(
  A_long iterationsL,
  void   *refconPV,
  PF_Err (*fn_func)(
    void    *refconPV,
    A_long  thread_idxL,
    A_long  i,
    A_long  itrtL));
```
