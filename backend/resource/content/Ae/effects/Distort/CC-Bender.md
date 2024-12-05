---
title: CC扭弯
order: 4
---

# CC Bender - CC扭弯

## 简述

在指定的两点之间弯曲。可以扭曲形状并使其反弹

## 效果展示

## 教程

7:20~

<iframe src="https://player.bilibili.com/player.html?bvid=BV1e34y1X7Vj&page=56&high_quality=1" width="100%" allowfullscreen="allowfullscreen" frameborder="0"></iframe>

## 中英日对照

![](https://mir.yuelili.com/user/AE/effects/AE-Effects-Distort-CC_Bender.png)

## 参数详解

### Amount 数量

![](https://cdn.yuelili.com/20211222133720.png)

### style 风格

如何弯曲图层

Bend：弯曲

Marilyn：玛丽莲梦露（妖娆的舞姿，摇摆你的臀部）

Sharp：尖锐（中间）

Boxer：拳击手（稳住下盘，通过上半身挪动来规避攻击）

![](https://cdn.yuelili.com/20211222134822.png)

### Adjust to Distance 按距离调整

如果选中，扭曲则根据 **Top** 和 **Base** 之间的宽度而变化。

![](https://cdn.yuelili.com/20211222135103.png)

### 其他

普通图层，正常使用，但是对于矢量图层（比如形状图层），如果起始、结束位置不跟随形状本身，可以用表达式链接一下（视频：4：02）

起始、结束位置表达式

```javascript
value + transform.position;
```

### 友情效果

[CC Bend it](https://www.yuelili.com/?p=15186) ：同样是扭曲图层

[范围扩散](https://www.yuelili.com/?p=15386) ：可以打破效果作用范围（默认是图层范围）
