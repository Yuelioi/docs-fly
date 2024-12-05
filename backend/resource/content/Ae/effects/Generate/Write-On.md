---
title: 写入
order: 26
---

# Write On - 写入

    # 简述 #

可在图层上为描边设置动画。例如，您可以模拟草书文本或签名的笔迹动作。

## 效果展示

![](https://mir.yuelili.com/2021/07/481a60db53bbaab8501f9fedaded72c8.png)

## 教程

<iframe src="https://player.bilibili.com/player.html?bvid=BV1f44y167NT&page=1&high_quality=1" width="100%" allowfullscreen="allowfullscreen" frameborder="0"></iframe>

## 中英日对照

![](https://mir.yuelili.com/user/AE/effects/AE-Effects-Generate-Write_On.png)![](https://mir.yuelili.com/user/AE/effects/AE-Effects-Generate-Write_On_cn.png)

| Write On              | 写入           | ブラシアニメーション       |                       |              |                    |
| --------------------- | -------------- | -------------------------- | --------------------- | ------------ | ------------------ |
| Brush Position        | 画笔位置       | ブラシの位置               |                       |              |                    |
| Color                 | 颜色           | カラー                     |                       |              |                    |
| Brush Size            | 画笔大小       | ブラシサイズ               |                       |              |                    |
| Brush Hardness        | 画笔硬度       | ブラシの硬さ               |                       |              |                    |
| Brush Opacity         | 画笔不透明度   | ブラシの不透明度           |                       |              |                    |
| Stroke Length (secs)  | 描边长度（秒） | ストロークの長さ(秒)       |                       |              |                    |
| Brush Spacing(secs)   | 画笔间距（秒   |                            |                       |              |                    |
| Paint Time Properties | 绘画时间属性   | ペイントのタイムプロパティ | None                  | 无           | なし               |
|                       |                |                            | Opacity               | 不透明度     | 不透明度           |
|                       |                |                            | Color                 | 颜色         | カラー             |
| Brush Time Properties | 画笔时间属性   | ブラシのタイムプロパティ   | None                  | 无           | なし               |
|                       |                |                            | Size                  | 大小         | サイズ             |
|                       |                |                            | Hardness              | 硬度         | 硬さ               |
|                       |                |                            | Size & Hardness       | 大小和硬度   |                    |
| Paint Style           | 绘画样式       | ペイントスタイル           | On Original Image     | 在原始图像上 | 元のイメージ       |
|                       |                |                            | On Transparent        | 在透明背景上 | 透明               |
|                       |                |                            | Reveal Original Image | 显示原始图像 | 元のイメージを表示 |

## 参数详解

### 笔刷位置

笔刷的位置。为此属性设置动画可创建描边。

### 描边长度（秒）

每个笔刷标记的持续时间，以秒为单位。如果此值为 0，则笔刷标记的持续时间无限。使用单个非零的常量值可创建蛇形移动的描边。为此值设置动画可使描边扩展和收缩。

### 笔刷间距（秒）

笔刷标记之间的时间间隔，以秒为单位。值越小，绘画描边越平滑，但渲染时间越长。

### 绘画时间属性和笔刷时间属性

指定是将绘画属性和笔刷属性应用到每个笔刷标记还是整个描边。选择“无”可将每个时间的值应用到描边中的所有笔刷标记。选择每个笔刷标记的属性名称可保留绘制笔刷标记时该属性的值。例如，如果选择“颜色”，则每个笔刷标记均保留绘制标记时“颜色”值指定的颜色。

### 绘画样式

绘画描边与原始图像相互作用的方式：

### 在原始图像上

绘画描边显示在原始图像上。

### 在透明背景上

绘画描边显示在透明背景上；不显示原始图像。

### 显示原始图像

通过绘画描边显示原始图像。

## 案例

![](https://mir.yuelili.com/2021/07/481a60db53bbaab8501f9fedaded72c8.png)
