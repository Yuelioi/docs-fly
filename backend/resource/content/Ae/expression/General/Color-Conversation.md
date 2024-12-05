---
title: Color-Conversation
order: 1
---

# Color Conversation 颜色转换

## rgbToHsl()

用法：rgbToHsl(rgbaArray)

说明：输入红色、绿色、蓝色和 Alpha 通道值数组，全部介于 0.0 到 1.0 范围内。颜色拾取器默认 Alpha 值为 1.0。生成色相、饱和度、亮度以及 Alpha 通道数组，同样介于 0.0 到 1.0 。

参数：rgba 为四维数组

类型：hsla 的四维数组

示例：可以去该网站在线转换[https://tool.lu/color/](https://tool.lu/color/)

![](https://mir.yuelili.com/user/AE/expression/exp-3-1.bmp)

```javascript
rgbToHsl(effect("Change Color")("Color To Change")); //红色从rgb转为hsl是从(1,0,0,1)变成(0,1,0.5,1)。剩下一个1代表Alpha，默认为100%。
```

## hslToRgb()

用法：hslToRgb(hslaArray)

说明：输入色相、饱和度、亮度以及 Alpha 通道数组，介于 0.0 到 1.0 。生成红色、绿色、蓝色和 Alpha 通道值数组，也是介于 0.0 到 1.0 范围内。颜色拾取器默认 Alpha 值为 1.0。

参数：hsla 的四维数组

类型：rgba 为四维数组

示例：可以去该网站在线转换[https://tool.lu/color/](https://tool.lu/color/)

```javascript
hslToRgb([0, 1, 0.5, 1]); //返回(1,0,0,1) 红色。跟rgbToHsl()为逆运算
```

## hexToRgb()

用法：hexToRgb(hexString)

说明：将十六进制颜色值（例如#FF0000）转换为 RGBA 样式颜色。

参数：hexString 为 16 进制的颜色值

类型：RGBA 四维颜色数组

示例：

```javascript
hexToRgb("#FF0000"); //返回(1,0,0,1) 红色
```
