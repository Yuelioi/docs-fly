---
title: Color-Space-Conversion
order: 5
---

# 颜色空间转换

不同的像素格式对不同的操作有用。AE通过`PF_ColorCallbacksSuite`暴露其内部功能。以下是支持的格式。

## 不同颜色空间的像素类型

像素类型和与数据结构

### 8 bpc ARGB

```cpp
typedef struct {
  A_u_char alpha, red, green, blue;
} PF_Pixel8;
```

### 16 bpc ARGB

```cpp
typedef struct {
  A_u_short alpha, red, green, blue;
} PF_Pixel16;
```

### 32 bpc ARGB

```cpp
typedef struct {
  PF_FpShort alpha, red, green, blue;
} PF_PixelFloat, PF_Pixel32;
```

### HLS (色相、亮度、饱和度)

```cpp
typedef PF_Fixed PF_HLS_PIXEL[3]
```

### YIQ (亮度、同相色度、正交色度)

```cpp
typedef PF_Fixed PF_YIQ_PIXEL[3]
```

插件可以使用以下回调函数利用为几乎任何颜色空间编写的图像处理算法。

## 色彩空间转换回调

| 函数  | 目的 | 替换 |
| ------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------ | --------------- |
| RGBtoHLS| 给定一个RGB像素，返回一个HLS（色调、亮度、饱和度）像素。HLS值从0到1缩放。| `PF_RGB_TO_HLS` |
| HLStoRGB| 给定一个HLS像素，返回一个RGB像素。  | `PF_HLS_TO_RGB` |
| RGBtoYIQ| 给定一个RGB像素，返回一个YIQ（亮度、同相色度、正交色度）像素。 Y为0到1，I在为-0.5959到0.5959，Q为-0.5227到0.5227。 | `PF_RGB_TO_YIQ` |
| YIQtoRGB| 给定一个YIQ像素，返回一个RGB像素。| `PF_YIQ_TO_RGB` |
| Luminance  | 给定一个RGB像素，返回其亮度值的100倍（0到25500）。  | `PF_LUMINANCE`  |
| Hue  | 给定一个RGB像素，它的色调角从0映射到255，其中0是0度，255是360度。| `PF_HUE`  |
| Lightness  | 给定一个RGB像素，返回它的亮度值（0到255）。| `PF_LIGHTNESS`  |
| Saturation | 给定一个RGB像素，返回其饱和度值（0到255）。 | `PF_SATURATION` |
