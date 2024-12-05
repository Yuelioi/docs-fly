---
title: FootageSource
order: 2
---

# FootageSource 素材源

## 素材源对象

app.project.item(index).mainSource  
app.project.item(index).proxySource

描述：素材源对象（FootageSource）对象保存描述某些素材来源的信息。它被用作 mainSource 一个的 FootageItem 对象，或者 proxySource 一个的 COMPITEM 对象或 FootageItem。

子级关系：素材源（footageSource）是纯色源（SolidSource）对象的基类，因此素材源的属性和方法，纯色源都可以使用。

![](https://cdn.yuelili.com/20211012170742.png)

## 属性

### alphaMode 素材 Alpha

app.project.item(index).mainSource.alphaMode  
app.project.item(index).proxySource.alphaMode

描述：定义如何解释素材中的 Alpha 信息。等同于 文件>解释素材>主要>Alpha。如果 hasAlpha 为 false，则此属性无效。

类型：Alpha 模式，枚举值；读/写：

- AlphaMode.IGNORE：忽略
- AlphaMode.STRAIGHT：直接
- AlphaMode.PREMULTIPLIED：预乘

### conformFrameRate 素材自定义帧速率

app.project.item(index).mainSource.conformFrameRate  
app.project.item(index).proxySource.conformFrameRate

描述：使用自定义帧速率而不是原始帧速率（nativeFrameRate）。等同于 文件>解释素材>主要>帧速率。

说明：如果设置为 0，则使用原始帧速率。如果素材是静止的，则无效。如果 removePulldown 不是 PulldownPhase.OFF，则将该值设置为 0 会报错。如果该值是 0，把 removePulldown 设置为除 PulldownPhase.OFF 的其他值，则自动设置为原始帧速率。

类型：浮点值，0.0~99.0；读/写。

### displayFrameRate 素材显示帧速率

app.project.item(index).mainSource.displayFrameRate  
app.project.item(index).proxySource.displayFrameRate

描述：由 After Effects 在合成中显示和渲染的有效帧速率。

说明：如果 removePulldown 为 PulldownPhase.OFF，则与 conformFrameRate(非零时)或 nativeFrameRate(如果 conformFrameRate 为 0 时)相同。如果 removePulldown 不是 PulldownPhase.OFF，则为 conformFrameRate

- 0.8，每 5 帧中删除 1 帧后的有效帧率。

类型：浮点值，0.0~99.0；只读。

### fieldSeparationType 素材场分离类型

app.project.item(index).mainSource.fieldSeparationType  
app.project.item(index).proxySource.fieldSeparationType

描述：如何在非静止画面中分离字段。

说明：如果 isStill 为 true
，则设置此属性是错误的。如果 removePulldown 不是 PulldownPhase.OFF，则将此值设置为 FieldSeparationType.OFF 报错。

类型：FieldSeparationType，枚举值; 读/写：

- FieldSeparationType.OFF
- FieldSeparationType.UPPER_FIELD_FIRST
- FieldSeparationType.LOWER_FIELD_FIRST

### hasAlpha 素材含 Alpha 信息

app.project.item(index).mainSource.hasAlpha  
app.project.item(index).proxySource.hasAlpha

描述：如果为 true，则素材包含 alpha 分量。此时，属性 alphaMode，invertAlpha 和 premulColor 具有有效值。为 false，则这些属性与素材无关。

类型：布尔值 只读。

### highQualityFieldSeparation 高质量的场分离

app.project.item(index).mainSource.highQualityFieldSeparation  
app.project.item(index).proxySource.highQualityFieldSeparation

描述：如果为 true，则 After Effects 使用特殊算法来确定如何执行高质量的场分离。

说明：如果 isStill 是 true，或者 fieldSeparationType 是 FieldSeparationType.OFF，则设置该属性报错。

类型：布尔值 读/写。

### invertAlpha 反转 Alpha 通道

app.project.item(index).mainSource.invertAlpha  
app.project.item(index).proxySource.invertAlpha

描述：如果为 true，则应反转素材片段或代理中的 Alpha 通道。仅当存在 alpha 时，此属性才有效。如果 hasAlpha 为 false，或者 alphaMode 为 is
AlphaMode.IGNORE，则忽略此属性。

类型：布尔值 读/写。

### isStill 是否静止素材

app.project.item(index).mainSource.isStill  
app.project.item(index).proxySource.isStill

描述：为 true 则镜头静止；如果为 false，则它具有基于时间的变换。

说明：静态镜头

示例：持续时间为 0 的 JPEG 文件，纯色层和占位符。非静态镜头的示例：电影文件，声音，序列和占位符。

类型：布尔值 只读。

### loop 素材循环

app.project.item(index).mainSource.loop  
app.project.item(index).proxySource.loop

描述：在合成中使用素材时，连续播放素材的次数。如果 isStill 为 true ，则设置此属性报错。

类型：整数 1~9999， 默认为 1; 读/写。

### nativeFrameRate 素材原始帧速率

app.project.item(index).mainSource.nativeFrameRate  
app.project.item(index).proxySource.nativeFrameRate

描述：素材的原始帧速率。

类型：浮点; 只读。

### premulColor 素材颜色预乘

app.project.item(index).mainSource.premulColor  
app.project.item(index).proxySource.premulColor

描述：颜色要预乘。此属性才有效，即 alphaMode 是 alphaMode.PREMULTIPLIED。

类型：三个浮点值的数组，范围 0~1；读/写。

### removePulldown 删除下拉菜单

app.project.item(index).mainSource.removePulldown  
app.project.item(index).proxySource.removePulldown

描述：使用场分离时如何删除下拉菜单。

说明：如果 isStill 为 true
，则设置此属性报错。将其设置为 PulldownPhase.OFF 以外的值在 fieldSeparationType 是 FieldSeparationType.OFF 的情况下，也会报错。

类型：PulldownPhase，枚举值; 读/写：

- PulldownPhase.RemovePulldown.OFF
- PulldownPhase.RemovePulldown.WSSWW
- PulldownPhase.RemovePulldown.SSWWW
- PulldownPhase.RemovePulldown.SWWWS
- PulldownPhase.RemovePulldown.WWWSS
- PulldownPhase.RemovePulldown.WWSSW
- PulldownPhase.RemovePulldown.WSSWW_24P_ADVANCE
- PulldownPhase.RemovePulldown.SSWWW_24P_ADVANCE
- PulldownPhase.RemovePulldown.SWWWS_24P_ADVANCE
- PulldownPhase.RemovePulldown.WWWSS_24P_ADVANCE
- PulldownPhase.RemovePulldown.WWSSW_24P_ADVANCE

## 方法

### guessAlphaMode() 获取 Alpha 模式

app.project.item(index).mainSource.guessAlphaMode()  
app.project.item(index).proxySource.guessAlphaMode()

描述：设置 alphaMode，premulColor 并 invertAlpha，为此素材源进行最佳判断。如果 hasAlpha 为 false，则不进行任何更改。

参数：无。

返回：无。

### guessPulldown() 最佳猜测值

app.project.item(index).mainSource.guessPulldown(method)  
app.project.item(index).proxySource.guessPulldown(method)

描述：设置 fieldSeparationType，并将 removePulldown 设置为此素材源的最佳猜测值。如果 isStill 为 true，则不进行任何更改。

参数：

- method，用于估计的方法。
- PulldownMethod，枚举值：
  - PulldownMethod.PULLDOWN_3_2
  - PulldownMethod.ADVANCE_24P

返回：无。
