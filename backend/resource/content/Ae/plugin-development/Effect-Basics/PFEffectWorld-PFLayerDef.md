---
title: PF_EffectWorld--PF_LayerDef
order: 5
---

# PF_EffectWorld   PF_LayerDef

After Effects 使用 PF_EffectWorlds 表示图像，也叫 PF_LayerDefs。

## PF_EffectWorld 结构体

| **Item**        | **Description**                                                                                                                                                                                                                                                                                                                                        |
| --------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `world_flags`  | 目前唯一的开关是<br />16-bpc 为 _PF_WorldFlag_DEEP_ <br /> 可以更改世界的图像数据时则为 _PF_WorldFlag_WRITEABLE_ <br />通常效果不能改变输入图像数据；只输出                                                                                       |
|`data`|指向图像数据的指针，存储为 `PF_PixelPtr`,不要直接访问；使用PF_PixelPtr 访问器宏。<br />After Effects 中的图像数据始终按从低字节到高字节的顺序进行组织，包含 Alpha、Red、Green、Blue。|
|`rowbytes`|图像像素块中每一行的长度（以字节为单位）。像素块包含高度线(height lines)，每个高度线都有宽度像素，后跟一些填充字节。宽度像素（乘以 4，因为每个像素长 4 个字节）加上可选的额外填充加起来为 rowbytes 字节。使用该值遍历图像数据。<br />特定于平台的行尾填充使得遍历整个缓冲区是不明智的。相反，使用高度和行字节查找每一行的开头。<br />注意：此值不会因场(fielt)渲染是否处于活动状态而有所不同。<br />注意：具有相同维度的输入和输出世界可以使用不同的行字节值。|
| `width`        | 像素缓冲区的宽度和高度。                                                                                                                                                                                                                                                                                                                  |
|`height`||
| `extent_hint`  | 包含图层中所有不透明（非零alpha）像素的最小矩形。这定义了需要输出的区域。如果您的插件随范围而变化（如扩散抖动），请忽略它并每次渲染完整帧。                                                                                              |
|`pix_aspect_ratio`|宽高像素比表示为PF_Rational。注意：效果可以对签出的图层使用此值，但必须对应用它们的图层使用PF_InData.pixel_aspect_ratio。|
| `platform_ref` | 不再在 CS5使用.<br />特定于平台的参考信息。在Windows上，它包含一个不透明的值。在macOS上，`PF_GET_PLATFORM_REFS` 提供来自 `PF_EffectWorld的CGrafPtr` 和`GDeviceHandle`<br />注意：您不能在_PF_Cmd_GLOBAL_SETUP_期间获取platform_ref，因为还没有任何输出上下文。耐心点，我的伙计。 |
| `dephault`     | 仅限图层参数. `PF_LayerDefault_MYSELF` 或 PF_LayerDefault_NONE`.                                                                                                                                                                                                                                                                    |

## 16.0 新内容

在 PF_Cmd_SMART_RENDER_GPU 期间，PF_LayerDef 的填充与普通的 CPU 渲染相同，但 PF_LayerDef.data 将为空；所有其他字段都是有效的。

## PF_EffectWorlds 中的 Rowbytes

不要以为你可以用`(width * sizeof(current_pixel_type))+4`的方式到达`PF_EffectWorld`的下一个扫描线。或者其他什么；用 PF_EffectWorld 的 "rowbytes "代替。

不要在 PF_EffectWorld 的指定区域外写入，这可能会破坏不属于你的缓存图像缓冲区。

为了测试你的效果是否遵守了`PF_EffectWorld>rowbytes`，在你的效果之后应用 Grow Bounds 效果。

输出的缓冲区将比输入的行字节更大(尽管它仍然有相同的逻辑大小)。

## 字节对齐 Byte Alignment

`PF_EffectWorld`中的像素不保证是 16 字节对齐的。一个效果可以得到一个更大的 PF_EffectWorld 的子区域。苹果公司的像素处理优化示例代码的用户，你已经被警告了。

除了每通道 8 位的颜色外，After Effects 还支持每通道 16 位和 32 位的浮点颜色。

插件 永远不会收到比特深度不同的输入和输出世界，也不会收到比特深度高于其声称能够处理的世界。

## 不透明（数据类型）像素的访问器宏

使用下面的宏来访问(不透明的)PF_PixelPtrs 中的数据。

强调的是，将一种类型的指针简单地投到另一种类型的指针中是不安全的！要想让它正常工作，就必须对其进行修改。要使它正常工作，就必须进行转换，而且没有任何东西可以阻止你错误地转换它。我们可能会在以后改变它的实现(到那时你会感谢我们强制执行这个抽象级别)。

## PF_PixelPtr 访问器宏

### PF_GET_PIXEL_DATA16

在指定的世界中获得一个指向 16-bpc 像素的指针。

如果世界不是 16-bpc，返回的像素指针将是 NULL。

第二个参数是可选的；如果它不是 NULL，返回的像素将是对传入的像素值的解释，就像它在指定的 PF_EffectWorld 中一样。

```cpp
PF_GET_PIXEL_DATA16(
PF_EffectWorldwP,
PF_PixelPtrpP0,
PF_Pixel16*outPP);
```

### PF_GET_PIXEL_DATA8

获得一个指向指定世界中的 8-bpc 像素的指针。

如果世界不是 8-bpc，返回的像素指针将是 NULL。

第二个参数是可选的；如果它不是 NULL，返回的像素将是对传入的像素值的解释，就像它在指定的 PF_EffectWorld 中一样。

```cpp
PF_GET_PIXEL_DATA8(
PF_EffectWorldwP,
PF_PixelPtrpP0,
PF_Pixel8*outPP);
```

---

把`PF_GET_PIXEL_DATA16`和`PF_GET_PIXEL_DATA8`看作是安全的(嗯)

需要的代码实际上非常简单，从 PF_EffectWorld 的输出中得到一个`PF_Pixel16*`。

```cpp
{
 PF_Pixel16 *deep_pixelP = NULL;
 PF_Err err = PF_Err_NONE;
 err = PF_GET_PIXEL_DATA16(output, NULL, &deep_pixelP);
}

```

如果世界上没有深层像素，这将返回 deep_pixelP 作为 NULL。

第二个参数并不常用，应该作为 NULL 传递；传递一个 PF_PixelPtr，它是包含在 PF_EffectWorld 中的，以强制到 PF_EffectWorld 的深度)。
