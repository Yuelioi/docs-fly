---
title: Image-Buffer-Management-Functions
order: 10
---

# 图像缓冲区管理函数

使用这些函数创建和销毁[PF_EffectWorld / PF_LayerDef](../effect-basics/PF_EffectWorld.html)，并找出它们的位深度。

## PF_WorldSuite2

### PF_NewWorld

创建一个新的 `PF_EffectWorld`。

```cpp
PF_ErrPF_NewWorld(
PF_ProgPtreffect_ref,
A_longwidthL,
A_longheightL,
PF_Booleanclear_pixB,
PF_PixelFormatpixel_format,
PF_EffectWorld*worldP);
```

### PF_DisposeWorld

处置一个 `PF_EffectWorld`。

```cpp
PF_ErrPF_DisposeWorld(
PF_ProgPtreffect_ref,
PF_EffectWorld*worldP);
```

### PF_GetPixelFormat

获取一个给定的 `PF_EffectWorld`的像素格式。

```cpp
PF_ErrPF_GetPixelFormat(
constPF_EffectWorld*worldP,
PF_PixelFormat*pixel_formatP);
```

`pixel_formatP`可以是。

- PF_PixelFormat_ARGB32` - 标准 8 位 RGB
- PF_PixelFormat_ARGB64` - 16-bit RGB
- PF_PixelFormat_ARGB128` - 32 位浮点 RGB
