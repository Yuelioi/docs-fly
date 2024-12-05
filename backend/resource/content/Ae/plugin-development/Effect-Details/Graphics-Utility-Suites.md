---
title: Graphics-Utility-Suites
order: 9
---

# 图形实用套件

After Effects 通过以下函数套件公开了其内部的变换和图形实用程序。

---

## Transform Worlds

这些函数以有趣的方式组合`PF_EffectWorlds`。使用这些函数时，使用的代码与AE内部的代码相同。

### PF_WorldTransformSuite1

#### composite_rect

使用 After Effects 的一个混合模式，将矩形从`PF_EffectWorld`合成到另一个。

```cpp
PF_Err composite_rect (
 PF_ProgPtr effect_ref,
 PF_Rect *src_rect,
 A_long src_opacity,
 PF_EffectWorld *src_world,
 A_long dst_x,
 A_long dst_y,
 PF_Field field_rdr,
 PF_XferMode xfer_mode,
 PF_EffectWorld *dst);

// 示例:https://youtu.be/aqq73u091gg&t=434
PF_CompositeMode ourCompositeMode;
switch(gip.modeInt) // 切换下拉框

case 1:
    ourCompositeMode.xfer = PF_Xfer_ADD; // 将图层混合模式改成ADD
    break;
case 2:
    ourCompositeMode.xfer = PF_Xfer_SCREEN; // 屏幕
    break;
// ...
ourCompositeMode.opacity = 255;
ourCompositeMode.opacitySu = A_u_short(1.0);
 // 高bit不透明度

```

`field_rdr`可以是upper, lower 或两者都是。
`xfer_mode`是以下的一种。

- `PF_Xfer_COPY`
- `PF_Xfer_BEHIND`
- `PF_Xfer_IN_FRONT`

#### blend

混合两幅图像，α-加权。不处理不同大小的源，尽管目标可以是 PF_EffectWorld。

```cpp
PF_Err blend (
 PF_ProgPtr effect_ref,
 const PF_EffectWorld *src1,
 const PF_EffectWorld *src2,
 PF_Fixed ratio,
 PF_EffectWorld *dst);

```

#### convolve

在 a、r、g 和 b 通道上分别用一个任意大小的内核对图像进行卷积。
可以指定一个矩形来进行卷积(例如，[PF_EffectWorld 结构](../effect-basics/PF_EffectWorld.html#effect-basics-pf-effectworld-structure)中的`extent_hint`)，或者使用 0 来卷积整个图像。
如果源是目标，则不要使用。
使用[内核开关] (#effect-details-graphics-utility-suites-kernel-flags)进行卷积。

```cpp
PF_Err convolve(
 PF_EffectWorld *src,
 const PF_Rect *area,
 PF_KernelFlags flags,
 A_long kernel_size,
 void *a_kernel,
 void *r_kernel,
 void *g_kernel,
 void *b_kernel,
 PF_EffectWorld *dst);
```

```cpp
// 示例 7*7卷积 简单模糊效果
// https://en.wikipedia.org/wiki/Kernel_(image_processing)
// https://youtu.be/DYKk2Rcub6U

// SKELETON_GAIN 为插件的滑块
// Render()
A_long convKer[49] = { };
for (size_t i = 0; i < 49; i++)
{
convKer[i] = A_long(params[SKELETON_GAIN]->u.fs_d.value*2.55);
}

if (params[SKELETON_GAIN]->u.fs_d.value) {
if (in_data->appl_id != 'PrMr') {
  ERR(suites.WorldTransformSuite1()->convolve(
  in_data->effect_ref,
  &params[0]->u.ld,
  &in_data->extent_hint,
  PF_KernelFlag_2D | PF_KernelFlag_CLAMP,
  7,
  convKer, convKer, convKer, convKer,
  output
  ));
}
}
```

#### copy

将一个区域从一个 PF_EffectWorld 复制到另一个，保留 alpha(与 macOS 的 CopyBits 不同)。

```cpp
PF_Err copy (
 PF_EffectWorld *src,
 PF_EffectWorld *dst,
 PF_Rect *src_r,
 PF_Rect *dst_r);

// 示例:https://youtu.be/aqq73u091gg
ERR(
  PF_COPY(&params[0]->u.ld, 
  output, 
  NULL, 
  NULL));
```

#### copy_hq

上述的高bpc版本(参数相同)。

#### transfer_rect

使用混合模式进行混合，有一个可选的遮罩。

```cpp
PF_Err transfer_rect (
 PF_ProgPtr effect_ref,
 PF_Quality quality,
 PF_ModeFlags m_flags,
 PF_Field field,
 const PF_Rect *src_rec,
 const PF_EffectWorld *src_world,
 const PF_CompositeMode *comp_mode,
 const PF_MaskWorld *mask_world0,
 A_long dest_x,
 A_long dest_y,
 PF_EffectWorld *dst_world);

// 示例
PF_CompositeMode ourCompositeMode;
ourCompositeMode.xfer = PF_Xfer_ADD; // 设置混合模式
ourCompositeMode.opacity = 255;
ourCompositeMode.opacitySu = A_u_short(1.0);


PF_Rect rect1; // 定义一个矩形
rect1.left = 0;
rect1.top = 0;
rect1.right = params[0]->u.ld.width;
rect1.right = params[0]->u.ld.height;

ERR(suites.WorldTransformSuite1()->
transfer_rect(
  in_data->effect_ref, 
  in_data->quality, 
  NULL,
  NULL,
  &rect1,
  &params[0]->u.ld, 
  &ourCompositeMode, 
  NULL, 
  0, 
  0, 
  output)
);

```

#### transform_world

给定一个 PF_EffectWorld 和一个矩阵(或矩阵数组)，使用 After Effects 的混合模式进行转换和混合，并有一个可选的遮罩。
矩阵的指针指向一个用于运动模糊的矩阵阵列。
什么时候变换不是变换？Z-尺寸 变换不是，除非变换后的图层是其他层的父层，这些层并非都位于z=0平面中。

```cpp
PF_Err transform_world (
 PF_InData *in_data,
 PF_Quality quality,
 PF_ModeFlags m_flags,
 PF_Field field,
 const PF_EffectWorld *src_world,
 const PF_CompositeMode *comp_mode,
 const PF_MaskWorld *mask_world0,
 const PF_FloatMatrix *matrices,
 A_long num_matrices,
 Boolean src2dst_matrix,
 const PF_Rect *dest_rect,
 PF_EffectWorld *dst_world);

```

---

## Kernel Flags

诸如卷积或高斯核之类的函数与核或过滤器权重值的矩阵一起工作。这些矩阵可以是任何格式。内核标志描述了应该如何创建和使用矩阵。或者一起使用您需要的任何标志。

与给定例程相关的标志与例程原型一起记录。左侧列中的第一个条目始终是默认值，值为0。

|Kernel Flags|介绍|
|---|---|
|`PF_KernelFlag_2D` <br/> `PF_KernelFlag_1D`|指定一维或二维内核。|
|`PF_KernelFlag_UNNORMALIZED`<br/>`PF_KernelFlag_NORMALIZED`|`NORMALIZED`均衡内核；内核表面下的体积与像素覆盖区域下的体积相同。|
|`PF_KernelFlag_CLAMP`<br/>`PF_KernelFlag_NO_CLAMP`|`CLAMP`将值限制为其数据类型的有效范围。|
|`PF_KernelFlag_USE_LONG`<br/>`PF_KernelFlag_USE_CHAR`<br/>`PF_KernelFlag_USE_FIXED`<br/>`PF_KernelFlag_USE_UNDEFINED`|`USE_LONG`将内核定义为值从0到255的长数组。<br/>`USE_CHAR`将内核定义为从0到255的无符号字符数组。<br/>`USE_FIXED`将内核定义为从0到1的fixeds数组。<br/>`USE_LONG`是唯一实现的标志。|
|`PF_KernelFlag_HORIZONTAL`<br/>`PF_KernelFlag_VERTICAL`|指定卷积的方向。|
|`PF_KernelFlag_TRANSPARENT_BORDERS`<br/>`PF_KernelFlag_REPLICATE_BORDERS`|使用`REPLICATE_BORDERS`复制边框像素在边缘外采样时，使用`TRANSPARENT_BORDERS`将边缘外像素视为alpha零（黑色）。<br/>`REPLICATE_BORDERS`不执行，将被忽略。|
|`PF_KernelFlag_STRAIGHT_CONVOLVE`<br/>`PF_KernelFlag_ALPHA_WEIGHT_CONVOLVE`|使用`STRAIGHT_CONVOLVE`表示直卷积(straight convolution)，使用`ALPHA_WEIGHT_CONVOLVE`告诉卷积 alpha 加权像素对结果卷积输出的贡献。<br/>`ALPHA_WEIGHT_CONVOLVE`不执行，将被忽略。
|

## 把它们填满

FillMatteSuite 可以用来填充一个`PF_EffectWorld'，可以用一个特定的颜色或者预乘一个 alpha 值。

### PF_FillMatteSuite2

`fill`用一种颜色填充一个矩形(或者，如果颜色指针是空的，则用黑色和 alpha 值填充)。

如果矩形是空的，它将填充整个图像。

```cpp
PF_Err fill (
 PF_ProgPtr effect_ref,
 const PF_Pixel *color,
 const PF_Rect *dst_rect,
 PF_EffectWorld *world);

```

#### fill16

与 fill 相同，但需要一个指向 PF_Pixel16 颜色的指针。

#### fill_float

获取指向 PF_PixelFloat 颜色的指针。

#### premultiply

转换为(或从)r、g 和 b 的颜色值，与黑色预乘，以表示 alpha 通道。
质量独立。

- `forward`作为布尔值使用。
- `true` 意味着将非预乘转换为预乘，
- `false` 意味着非预乘

```cpp
PF_Err premultiply (
 A_long forward,
 PF_EffectWorld *dst);

```

#### premultiply_color

转换为(或从)有 r、g、b 颜色值预乘以任何颜色来表示 alpha 通道。

```cpp
PF_Err premultiply_color (
 PF_ProgPtr effect_ref,
 PF_EffectWorld *src,
 PF_Pixel *color,
 A_long forward,
 PF_EffectWorld *dst);

```

#### premultiply_color16

同上，但需要一个指向 PF_Pixel16 颜色的指针。

#### premultiply_color_float

获取指向 PF_PixelFloat 颜色的指针。

---

## 采样图像

注意被采样图像边界之外的区域被视为零alpha。为方便起见，此表中列出了PF_Sampling8Suite1、PF_Sampling16Suite1和PF_SamplingFloatSuite1的函数。

### PF_SamplingSuite Functions (Multiple Suites)

#### nn_sample

执行近邻采样。

```cpp
PF_Err nn_sample (
 PF_ProgPtr effect_ref,
 PF_Fixed x,
 PF_Fixed y,
 const PF_SampPB *params,
 PF_Pixel *dst_pixel );

```

#### nn_sample16

同上，但需要一个指向` PF_Pixel16``dst_pixel `的指针。

#### nn_sample_float

获取指向`PF_PixelFloat`dst_pixel`的指针。

#### subpixel_sample

以高质量查询源图像中非积分点处颜色的适当α加权插值。最近邻采样用于低质量。

因为如果使用采样例程，通常会被多次调用，所以将函数指针复制到回调结构并复制到寄存器或堆栈上以加快内部循环是很方便的。

有关示例，请参阅示例代码。

注意：采样假设0,0是左上角像素的中心。

```cpp
PF_Err subpixel_sample (
 PF_ProgPtr effect_ref,
 PF_Fixed x,
 PF_Fixed y,
 const PF_SampPB *params,
 PF_Pixel *dst_pixel);

```

#### subpixel_sample16

同上，但需要一个指向` PF_Pixel16*``dst_pixel `的指针。

#### subpixel_sample_float

获取指向` PF_PixelFloat*``dst_pixel `的指针。

#### area_sample

使用它来计算高质量的源图像中轴对齐的非积分矩形颜色的适当α加权平均值。

最近邻采样用于低质量。由于溢出问题，这只能平均最大256 x 256像素区域（即x和y半径<128像素）。

注意：采样半径必须至少为x和y中的一个。

```cpp
PF_Err area_sample (
 PF_ProgPtr effect_ref,
 PF_Fixed x,
 PF_Fixed y,
 const PF_SampPB *params,
 PF_Pixel *dst_pixel);

```

注意：在图层边界之外的区域被认为与零alpha相同，用于采样。

#### area_sample16

同上，但需要一个` PF_Pixel16*``dst_pixel `。

### PF_BatchSamplingSuite1 Functions

#### begin_sampling

你的效果要进行一些批量采样。

After Effects 将执行设置任务来优化你的采样。

```cpp
PF_Err (*begin_sampling)(
 PF_ProgPtr effect_ref,
 PF_Quality qual,
 PF_ModeFlags mf,
 PF_SampPB *params);

```

#### end_sampling

告诉 After Effects 已经完成了采样。

```cpp
PF_Err (*end_sampling)(
 PF_ProgPtr effect_ref,
 PF_Quality qual,
 PF_ModeFlags mf,
 PF_SampPB *params);

```

#### get_batch_func

获得一个指向 After Effects 的批量采样功能的指针(高度优化)。

```cpp
PF_Err (*get_batch_func)(
 PF_ProgPtr effect_ref,
 PF_Quality quality,
 PF_ModeFlags mode_flags,
 const PF_SampPB *params,
 PF_BatchSampleFunc *batch);

```

#### get_batch_func16

获得一个指向 After Effects 的 16-bpc 批量取样功能的指针(也是高度优化的)。

```cpp
PF_Err (*get_batch_func16)(
 PF_ProgPtr effect_ref,
 PF_Quality quality,
 PF_ModeFlags mode_flags,
 const PF_SampPB *params,
 PF_BatchSample16Func *batch);

```

## 数学函数

除了各种图形实用程序之外，我们还提供了一个ANSI标准例程块，以便插件无需包含其他库即可使用标准函数。

我们给出了大量数学函数的函数指针(三角函数、平方根、对数等)。

使用我们的套间函数提供了一些(应用层面的)错误处理，并防止了包含不同版本的多个 "标准 "库的问题。

所有函数都返回双精度。所有的角度都用弧度表示，如果需要的话，可以使用`PF_RAD_PER_DEGREE`(AE_EffectCB.h 中的一个常数)来从度转换为弧度。

### PF_ANSICallbackSuite1

| **Function**                                                            | **Purpose**                                                  | **Replaces** |
| ----------------------------------------------------------------------- | ------------------------------------------------------------ | ------------ |
| `acos`                                                                  | 返回 arc cosine of x.                                 | `PF_ACOS`    |
| `asin`                                                                  | 返回 arc sine of x.                                   | `PF_ASIN`    |
| `atan`                                                                  | 返回 arc tangent of x.                                | `PF_ATAN`    |
| `atan2`                                                                 | 返回 atan(y/x).                                           | `PF_ATAN2`   |
| `ceil`                                                                  | 返回 next integer above x.                            | `PF_CEIL`    |
| `cos`                                                                   | 返回 cosine of x.                                     | `PF_COS`     |
| `exp`                                                                   | 返回  e to the power of x.                                 | `PF_EXP`     |
| `fabs`                                                                  | 返回 absolute value of x.                             | `PF_FABS`    |
| `floor`                                                                 | 返回 closest integer below x.                         | `PF_FLOOR`   |
| `fmod`                                                                  | 返回  x modulus y.                                         | `PF_FMOD`    |
| `hypot`                                                                 | 返回 hypotenuse of x and y, which is sqrt(x*x + y*y). | `PF_HYPOT`   |
| `log`                                                                   | 返回 natural log (ln) of x.                           | `PF_LOG`     |
| `log10`                                                                 | 返回 log (base 10) of x.                              | `PF_LOG10`   |
| `pow`                                                                   | 返回  x to the power of y.                                 | `PF_POW`     |
| `sin`                                                                   | 返回 sine of x.                                       | `PF_SIN`     |
| `sqrt`                                                                  | 返回 square root of x.                                | `PF_SQRT`    |
| `tan`                                                                   | 返回 tangent of x.                                    | `PF_TAN`     |
| *(while not strictly math functions, these emulate ANSI functionality)* |
| `sprintf`                                                               | 模拟 C sprintf function.                             | `PF_SPRINTF` |
| `strcpy`                                                                | 模拟 C strcpy function.                              | `PF_STRCPY`  |
