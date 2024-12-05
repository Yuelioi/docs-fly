---
title: 填充与描边
order: 2
---
# 填充与描边

[本节的 adobe 官方文档链接](https://helpx.adobe.com/cn/after-effects/using/shape-attributes-paint-operations-path.html)

## 中英对照

![](https://mir.yuelili.com/2021/07/ea58937a60f1e3d87e3c2200a0c11a7b.png)

| Property        | 属性     |                  |                |                              |                    |         |        |
| --------------- | -------- | ---------------- | -------------- | ---------------------------- | ------------------ | ------- | ------ |
| Fill            | 填充     |                  | 混合模式不赘述 |                              |                    |         |        |
|                 |          | Composite        | 合成           |                              |                    |         |        |
|                 |          |                  |                | Below Previous in Same Group | 在同组中前一个之下 |         |        |
|                 |          |                  |                | Above Previous in Same Group | 在同组中前一个之上 |         |        |
|                 |          | Fill Rule        | 填充规则       |                              |                    |         |        |
|                 |          |                  |                | Non-Zero Winding             | 非零环绕           |         |        |
|                 |          |                  |                | Even-Odd                     | 奇偶               |         |        |
|                 |          | Color            | 颜色           |                              |                    |         |        |
|                 |          | Opacity          | 不透明度       |                              |                    |         |        |
| Stroke          | 描边     |                  | 混合模式不赘述 |                              |                    |         |        |
|                 |          | Composite        | 合成           |                              |                    |         |        |
|                 |          |                  |                | Below Previous in Same Group | 在同组中前一个之下 |         |        |
|                 |          |                  |                | Above Previous in Same Group | 在同组中前一个之上 |         |        |
|                 |          | Color            | 颜色           |                              |                    |         |        |
|                 |          | Opacity          | 不透明度       |                              |                    |         |        |
|                 |          | Stroke Width     | 描边宽度       |                              |                    |         |        |
|                 |          | Line Cap         | 线段端点       |                              |                    |         |        |
|                 |          |                  |                | Butt Cap                     | 平头端点           |         |        |
|                 |          |                  |                | Round Cap                    | 圆头端点           |         |        |
|                 |          |                  |                | Projecting Cap               | 短形端点           |         |        |
|                 |          | Line Join        | 线段连接       |                              |                    |         |        |
|                 |          |                  |                | Miter Join                   | 斜接连接           |         |        |
|                 |          |                  |                | Round join                   | 圆角连接           |         |        |
|                 |          |                  |                | Bevel join                   | 斜面连接           |         |        |
|                 |          | Miter Limit      | 尖角限制       |                              |                    |         |        |
|                 |          | Dashes           | 虚线           |                              |                    |         |        |
|                 |          |                  |                | Dash                         | 虚线               |         |        |
|                 |          |                  |                | Gap                          | 间隙               |         |        |
|                 |          |                  |                | Dash2                        | 虚线 2             |         |        |
|                 |          |                  |                | Gap2                         | 间隙 2             |         |        |
|                 |          |                  |                | Dash3                        | 虚线 3             |         |        |
|                 |          |                  |                | Gap3                         | 间隙 3             |         |        |
|                 |          |                  |                | Offset                       | 偏移               |         |        |
|                 |          | Taper            | 锥度           |                              |                    |         |        |
|                 |          |                  |                | Length Units                 | 长度单位           | Percent | 百分比 |
|                 |          |                  |                |                              |                    | Pixels  | 像素   |
|                 |          |                  |                | Start Length                 | 起始长度           |         |        |
|                 |          |                  |                | End Length                   | 结束长度           |         |        |
|                 |          |                  |                | Start Width                  | 起始宽度           |         |        |
|                 |          |                  |                | End Width                    | 结束宽度           |         |        |
|                 |          |                  |                | Start Ease                   | 起始缓和           |         |        |
|                 |          |                  |                | End Ease                     | 结束缓和           |         |        |
|                 |          | Wave             | 波形           |                              |                    |         |        |
|                 |          |                  |                | Amount                       | 数量               |         |        |
|                 |          |                  |                | Units                        | 单位               | Percent | 百分比 |
|                 |          |                  |                |                              |                    | Pixels  | 像素   |
|                 |          |                  |                | Wavelength                   | 波长               |         |        |
|                 |          |                  |                | Cycles                       | 环形               |         |        |
|                 |          |                  |                | Phase                        | 相位               |         |        |
| Gradient Stroke | 渐变描边 |                  | 混合模式不赘述 |                              |                    |         |        |
|                 |          | Composite        | 合成           |                              |                    |         |        |
|                 |          |                  |                | Below Previous in Same Group | 在同组中前一个之下 |         |        |
|                 |          |                  |                | Above Previous in Same Group | 在同组中前一个之上 |         |        |
|                 |          | Fill Ruler       | 填充规则       |                              |                    |         |        |
|                 |          |                  |                | Non-Zero Winding             | 非零环绕           |         |        |
|                 |          |                  |                | Even-Odd                     | 奇偶               |         |        |
|                 |          | Type             | 类型           |                              |                    |         |        |
|                 |          |                  |                | Linear                       | 线性               |         |        |
|                 |          |                  |                | Radial                       | 径向               |         |        |
|                 |          | Start Point      | 起始点         |                              |                    |         |        |
|                 |          | End Point        | 结束点         |                              |                    |         |        |
|                 |          | Highlight Length |                |                              |                    |         |        |
|                 |          | Hightlight Andle |                |                              |                    |         |        |
|                 |          | Color            | 颜色           |                              |                    |         |        |
|                 |          | Opacity          | 不透明度       |                              |                    |         |        |
| Gradient Fill   | 渐变填充 |                  |                |                              |                    |         |        |
|                 |          | Composite        | 合成           |                              |                    |         |        |
|                 |          |                  |                | Below Previous in Same Group | 在同组中前一个之下 |         |        |
|                 |          |                  |                | Above Previous in Same Group | 在同组中前一个之上 |         |        |
|                 |          | Type             | 类型           |                              |                    |         |        |
|                 |          | Start Point      | 起始点         |                              |                    |         |        |
|                 |          | End Point        | 结束点         |                              |                    |         |        |
|                 |          | Color            | 颜色           |                              |                    |         |        |
|                 |          | Opacity          | 不透明度       |                              |                    |         |        |
|                 |          | Stroke Width     | 描边宽度       |                              |                    |         |        |
|                 |          | Line Cap         | 线段端点       |                              |                    |         |        |
|                 |          |                  |                | Butt Cap                     | 平头端点           |         |        |
|                 |          |                  |                | Round Cap                    | 圆头端点           |         |        |
|                 |          |                  |                | Projecting Cap               | 短形端点           |         |        |
|                 |          | Line Join        | 线段连接       |                              |                    |         |        |
|                 |          |                  |                |                              |                    |         |        |
|                 |          |                  |                | Miter Join                   | 斜接连接           |         |        |
|                 |          |                  |                | Round join                   | 圆角连接           |         |        |
|                 |          |                  |                | Bevel join                   | 斜面连接           |         |        |
|                 |          | Miter Limit      | 尖角限制       |                              |                    |         |        |
|                 |          | Dashes           | 虚线           |                              |                    |         |        |

卷首语

默认情况下，描边是基于边向形状内外扩散，如果在形状图层中改变他俩（描边&填充）的顺序，则可以更改覆盖顺序

## 填充

### 1.Composite 合成

- Below Previous in Same Group 在同组中前一个之下
- Above Previous in Same Group 在同组中前一个之上

以下为"描边-合成"的解释，原理一样（懒得截图了）

不难理解，众所周知，形状图层里可以放置多个效果，比如我添加了 3 个描边，颜色如下

![](https://mir.yuelili.com/user/AE/mg/shape-layer/shape-stroke2.png)

如果 3 个都勾选"在同组中前一个之下"，那么最终为红色。 如果 3 个都勾选"在同组中前一个之上"，那么最终为黄色。

原理：AE 是从上往下计算效果的，假设都选"之上"，（描边的其他参数同理）

运行到红色描边，结果：本层应为描边顶层，显示为红色； 再运行到绿色描边，结果：本层应为描边顶层，显示为绿色； 再运行到黄色描边，结果：本层应为描边顶层，显示为红色；

### 2.Fill Rule 填充规则

填充是指在路径的内部区域内绘制颜色。路径非常简单（例如是一个圆形）时，很容易确定将哪个区域视为路径内部。然而，如果路径与它本身相交，或者复合路径包含由其他路径圈起的路径，确定路径内部就不那么容易。

因此制定以下两个规则确定路径内部。统计从某个点穿过路径向路径环绕的区域外部绘制直线的次数。非零环绕填充规则考虑路径方向；奇偶填充规则不考虑。

![](https://mir.yuelili.com/user/AE/mg/shape-layer/rs_24.png)

非零环绕”（左图）与“奇偶填充规则”（右图）

- Non-Zero Winding 非零环绕

直线的交叉计数是直线穿过路径的自左向右部分的总次数减去直线穿过路径的自右向左部分的总次数。如果按任意方向从该点绘制的直线的交叉计数为零，则该点位于路径外部；否则，该点位于路径内部。

- Even-Odd 奇偶

如果从某个点按任意方向穿过路径绘制直线的次数为奇数次，则该点位于路径内部；否则，该点位于路径外部。

### 3.Color 颜色

填充的颜色

### 4.Opacity 不透明度

填充的不透明度

## 描边

### 1.Composite 合成

Below Previous in Same Group 在同组中前一个之下

Above Previous in Same Group 在同组中前一个之上

不难理解，众所周知，形状图层里可以放置多个效果，比如我添加了 3 个描边，颜色如下

![](https://mir.yuelili.com/user/AE/mg/shape-layer/shape-stroke2.png)

如果 3 个都勾选"在同组中前一个之下"，那么最终为红色。 如果 3 个都勾选"在同组中前一个之上"，那么最终为黄色。

原理：AE 是从上往下计算效果的，假设都选"之上"，（描边的其他参数同理）

运行到红色描边，结果：本层应为描边顶层，显示为红色； 再运行到绿色描边，结果：本层应为描边顶层，显示为绿色； 再运行到黄色描边，结果：本层应为描边顶层，显示为红色；

### 2.Color 颜色 & Opacity 不透明度 & Stroke Width 描边宽度

![](https://mir.yuelili.com/user/AE/mg/shape-layer/shape-stroke3.png)

描边的颜色、不透明度以及宽度。注意：如果在顶部菜单栏更改，则会同时改变这些参数

### 3.Line Cap 线段端点

线段的端点模式，需要增加 Dash（虚线）才可以看到。一共三种模式

- Butt Cap 平头端点

随着描边宽度的增大，会垂直于描边向垂线方向扩散

- Round Cap 圆头端点

随着描边宽度的增大，向四周扩散，只不过带圆角

- Projecting Cap 短形端点

随着描边宽度的增大，向四周扩散

下图为随着描边宽度增大，三种模式的变换形态

![](https://mir.yuelili.com/user/AE/mg/shape-layer/Line-Cap.gif)

### 4.Line Join 线段连接

拐角的连接方式，见下图（Miter Join 斜接连接、Round join 圆角连接、Bevel join 斜面连接）

![](https://mir.yuelili.com/user/AE/mg/shape-layer/Line-Join1.png)

### 5.Miter Limit 尖角限制

“尖角限制”值确定哪些情况下使用斜面连接而不是斜接连接。如果尖角限制是 4，则当点的长度达到描边粗细的四倍时，改用斜面连接。如果尖角限制为
1，则产生斜面连接。

### 6.Dashes 虚线

添加最多 3 组的虚线和间隙来创建虚线形描边

下图红色描边就是虚线，中间白色的就是间隙，结合参数应该不难理解。

offset 可以控制间隙在路径上流动

![](https://mir.yuelili.com/user/AE/mg/shape-layer/dash1.png)

![](https://mir.yuelili.com/user/AE/mg/shape-layer/dash2.png)

### 7.Taper

2020.5 月新增

Taper 可以使描边的开头和结尾变窄

- Length Units 描边长度的计算方法
  - Percent 按百分比计算
  - Pixels 按像素计算（为了准确演示，以下示例用像素计算）
  - Start Length 起始长度 ：描边开头区域需要变窄的范围（下图黄色的 200）
  - End Length 结束长度：描边结束区域需要变窄的范围（下图蓝色的 200px）
  - Start Width 起始宽度：开头要变多窄的宽度值（下图的 0%代表变为一个点）
  - End Width 结束宽度：结尾要变多窄的宽度值（下图的 50%代表窄了一半）
  - Start Ease 起始缓和：起始端头的圆度（见图三）
  - End Ease 结束缓和：结束端头的圆度（见图三）

![](https://mir.yuelili.com/user/AE/mg/shape-layer/Taper4.png)

![](https://mir.yuelili.com/user/AE/mg/shape-layer/Taper1.png)

![](https://mir.yuelili.com/user/AE/mg/shape-layer/Taper2.png)

简单示例：使用 Taper 和 Trim Path 制作的简单动画

![](https://mir.yuelili.com/user/AE/mg/shape-layer/Taper3.gif)

### 8.Wave

2020 新增

![](https://mir.yuelili.com/user/AE/mg/shape-layer/wave1.png)

- Amount 数量

与其说是数量，不如说是波浪高度。如上图左所示：

50%波浪高 = 描边一半

100%波浪高 = 描边高度

- Units 单位

Percent ：按百分比计算波浪形态

Cycles ：按圈数计算波浪形态

- Wavelength 波长

按像素计算波浪长度，如上图右所示，一波为 125px

- Cycles 圈数

按圈数计算，如上图右所示，整个路径一共有 4 圈

- Phase 相位

波浪在路径上移动，有点像虚线（dashes）里的 offset

## 渐变填充

### 1.Composite 合成

Below Previous in Same Group 在同组中前一个之下

Above Previous in Same Group 在同组中前一个之上

以下为"描边-合成"的解释，原理一样（懒得截图了）

不难理解，众所周知，形状图层里可以放置多个效果，比如我添加了 3 个描边，颜色如下

![](https://mir.yuelili.com/user/AE/mg/shape-layer/shape-stroke2.png)

如果 3 个都勾选"在同组中前一个之下"，那么最终为红色。 如果 3 个都勾选"在同组中前一个之上"，那么最终为黄色。

原理：AE 是从上往下计算效果的，假设都选"之上"，（描边的其他参数同理）

运行到红色描边，结果：本层应为描边顶层，显示为红色； 再运行到绿色描边，结果：本层应为描边顶层，显示为绿色； 再运行到黄色描边，结果：本层应为描边顶层，显示为红色；

### 2.Fill Ruler 填充规则

Non-Zero Winding 非零环绕
Even-Odd 奇偶

填充是指在路径的内部区域内绘制颜色。路径非常简单（例如是一个圆形）时，很容易确定将哪个区域视为路径内部。然而，如果路径与它本身相交，或者复合路径包含由其他路径圈起的路径，确定路径内部就不那么容易。

因此制定以下两个规则确定路径内部。统计从某个点穿过路径向路径环绕的区域外部绘制直线的次数。非零环绕填充规则考虑路径方向；奇偶填充规则不考虑。

![](https://mir.yuelili.com/user/AE/mg/shape-layer/rs_24.png)

非零环绕”（左图）与“奇偶填充规则”（右图）

### 3.Type 类型

![](https://mir.yuelili.com/user/AE/mg/shape-layer/linear-radial.png)

- Linear 线性（线性渐变，见左图）
  - Radial 径向（径向渐变，见右图）

### 4.起始结束点

控制渐变的起始/结束点，注意，形状图层的原点位置[0,0]默认在图层的正中央，如下图大红点所示。

- Start Point 起始点
  - End Point 结束点

![](https://mir.yuelili.com/user/AE/mg/shape-layer/start-end-point.png)

线性渐变（左图）白黑交界

起始点为[-50,0] 黄点
结束点为[50,0] 蓝点

径向渐变（右图）

起始点为[0,0] 红点
结束点为[50,50] 蓝点

### 5.Highlight 高光

本属性为径向渐变独有

Highlight Length 高光长度：高光的长度。见下图一 Hightlight Andle 高光角度：高光点与初始点的夹角。见下图二

| ![](https://mir.yuelili.com/user/AE/mg/shape-layer/highlight.png) | ![](https://mir.yuelili.com/user/AE/mg/shape-layer/highlight-angle.png) |
| --------------------------------------------------------------- | --------------------------------------------------------------------- |
| 图一：高光长度                                                  | 图二：高光角度                                                        |

### 6.Color 颜色

填充颜色，没啥好说的

### 7.Opacity 不透明度

填充不透明度，没啥好说的

## 渐变描边

渐变属性与渐变填充的渐变属性一致

描边属性与渐变填充的描边属性一致。

如果学习了"[描边](https://mir.yuelili.com/docs/shape-layer/fill-stoke#%e6%8f%8f%e8%be%b9)"与"[渐变填充](https://mir.yuelili.com/docs/shape-layer/fill-stoke#%e6%b8%90%e5%8f%98%e5%a1%ab%e5%85%85)"，那么本节可以快速跳过

### 1.Composite 合成

- Below Previous in Same Group 在同组中前一个之下
- Above Previous in Same Group 在同组中前一个之上

以下为"描边-合成"的解释，原理一样（懒得截图了）

判断同组多个属性，谁在顶层。图是描边的图，道理一样。

![](https://mir.yuelili.com/user/AE/mg/shape-layer/shape-stroke2.png)

如果 3 个都勾选"在同组中前一个之下"，那么最终为红色。 如果 3 个都勾选"在同组中前一个之上"，那么最终为黄色。

原理：AE 是从上往下计算效果的，假设都选"之上"，（描边的其他参数同理）

运行到红色描边，结果：本层应为描边顶层，显示为红色； 再运行到绿色描边，结果：本层应为描边顶层，显示为绿色； 再运行到黄色描边，结果：本层应为描边顶层，显示为红色；

### 2.Type 类型

![](https://mir.yuelili.com/user/AE/mg/shape-layer/gradient-stroke.png)

### 3.起始结束点

控制渐变的起始/结束点，注意，形状图层的原点位置[0,0]默认在图层的正中央，如下图大红点所示。

- Start Point 起始点
  - End Point 结束点

![](https://mir.yuelili.com/user/AE/mg/shape-layer/start-end-point.png)

线性渐变（左图）白黑交界

起始点为[-50,0] 黄点
结束点为[50,0] 蓝点

径向渐变（右图）

起始点为[0,0] 红点
结束点为[50,50] 蓝点

### 4.Highlight 高光

本属性为径向渐变独有

Highlight Length 高光长度：高光的长度。见下图一 Hightlight Andle 高光角度：高光点与初始点的夹角。见下图二

| ![](https://mir.yuelili.com/user/AE/mg/shape-layer/highlight.png) | ![](https://mir.yuelili.com/user/AE/mg/shape-layer/highlight-angle.png) |
| --------------------------------------------------------------- | --------------------------------------------------------------------- |
| 图一：高光长度                                                  | 图二：高光角度                                                        |

### 5.Color 颜色 & Opacity 不透明度 & Stroke Width 描边宽度

![](https://mir.yuelili.com/user/AE/mg/shape-layer/shape-stroke3.png)

描边的颜色、不透明度以及宽度。注意：如果在顶部菜单栏更改，则会同时改变这些参数

### 6.Line Cap 线段端点

线段的端点模式，需要增加 Dash（虚线）才可以看到。一共三种模式

- Butt Cap 平头端点

随着描边宽度的增大，会垂直于描边向垂线方向扩散

- Round Cap 圆头端点

随着描边宽度的增大，向四周扩散，只不过带圆角

- Projecting Cap 短形端点

随着描边宽度的增大，向四周扩散

下图为随着描边宽度增大，三种模式的变换形态

![](https://mir.yuelili.com/user/AE/mg/shape-layer/Line-Cap.gif)

### 7.Line Join 线段连接

拐角的连接方式，见下图（Miter Join 斜接连接、Round join 圆角连接、Bevel join 斜面连接）

![](https://mir.yuelili.com/user/AE/mg/shape-layer/Line-Join1.png)

### 8.Miter Limit 尖角限制

“尖角限制”值确定哪些情况下使用斜面连接而不是斜接连接。如果尖角限制是 4，则当点的长度达到描边粗细的四倍时，改用斜面连接。如果尖角限制为
1，则产生斜面连接。

### 9.Dashes 虚线

添加最多 3 组的虚线和间隙来创建虚线形描边

下图红色描边就是虚线，中间白色的就是间隙，结合参数应该不难理解。

offset 可以控制间隙在路径上流动

| ![](https://mir.yuelili.com/user/AE/mg/shape-layer/dash1.png) | ![](https://mir.yuelili.com/user/AE/mg/shape-layer/dash2.png) |
| ----------------------------------------------------------- | ----------------------------------------------------------- |

### 10.Taper

2020.5 月新增

Taper 可以使描边的开头和结尾变窄

- Length Units 描边长度的计算方法
- Percent 按百分比计算

  - Pixels 按像素计算（为了准确演示，以下示例用像素计算）
  - Start Length 起始长度 ：描边开头区域需要变窄的范围（下图黄色的 200）
  - End Length 结束长度：描边结束区域需要变窄的范围（下图蓝色的 200px）
  - Start Width 起始宽度：开头要变多窄的宽度值（下图的 0%代表变为一个点）
  - End Width 结束宽度：结尾要变多窄的宽度值（下图的 50%代表窄了一半）
  - Start Ease 起始缓和：起始端头的圆度（见图三）
  - End Ease 结束缓和：结束端头的圆度（见图三）

| ![](https://mir.yuelili.com/user/AE/mg/shape-layer/Taper4.png) | ![](https://mir.yuelili.com/user/AE/mg/shape-layer/Taper1.png) | ![](https://mir.yuelili.com/user/AE/mg/shape-layer/Taper2.png) |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| 图一：参数显示                                               | 图二：长度/宽度演示 形（矩形大小：600X400）                  | 图三：缓和度演示                                             |

简单示例：使用 Taper 和 Trim Path 制作的简单动画 ![](https://mir.yuelili.com/user/AE/mg/shape-layer/Taper3.gif)

#### 11.Wave

2020 新增

![](https://mir.yuelili.com/user/AE/mg/shape-layer/wave1.png)

- Amount 数量

与其说是数量，不如说是波浪高度。如上图左所示：

50%波浪高 = 描边一半

100%波浪高 = 描边高度

- Units

Percent ：按百分比计算波浪形态

Cycles ：按圈数计算波浪形态

- Wavelength 波浪长度

按像素计算波浪长度，如上图右所示，一波为 125px

- Cycles 圈数

按圈数计算，如上图右所示，整个路径一共有 4 圈，

- Phase 相位

波浪在路径上移动，有点像虚线（dashes）里的 offset

## 技巧

tip1：放置在形状图层的效果是有层级的。描边是基于边线向四周扩展，如果把描边放置在填充上，则是正常的扩散（图 1 左）；而把填充放置在描边上，则只向外扩散（图 1 右）

![](https://mir.yuelili.com/user/AE/mg/shape-layer/shape-stroke.png)

tip2：路径可以更改方向，某些规则也会因此更改，路径-“反转路径方向（开）![](https://mir.yuelili.com/user/AE/mg/shape-layer/P_ReversePathOn_Md_N.png)。
