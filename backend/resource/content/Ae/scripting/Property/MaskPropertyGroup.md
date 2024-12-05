---
title: MaskPropertyGroup
order: 1
---

# MaskPropertyGroup 蒙版属性组

## 蒙版属性组对象

app.project.item(index).layer(index).mask

描述：MaskPropertyGroup 对象，将蒙版属性封装在一个图层中。

父级关系：蒙版属性组（MaskPropertyGroup）是属性组（PropertyGroup）的子类。PropertyBase 对象和 PropertyGroup 的所有方法和属性均可使用

## 属性

### 添加蒙版示例

```javascript
var selLayer = app.project.activeItem.selectedLayers[0];
// 创建蒙版对象
newMask = selLayer.Masks.addProperty("Mask");
newMask.color = [1, 1, 1]; //设置蒙版颜色
newMask.inverted = true; //设置蒙版反转
newMask.locked = true; //设置蒙版锁定

// 创建蒙版形状
myMaskShape = newMask.property("maskShape");
myShape = myMaskShape.value;
myShape.vertices = [
  [930, 450],
  [930, 570],
  [1050, 570],
  [1050, 450],
];
myShape.closed = true;
myMaskShape.setValue(myShape);
```

### color 蒙版颜色

app.project.item(index).layer(index).mask(index).color

描述：在用户界面(“合成”面板，“图层”面板和“时间轴”面板)中蒙版轮廓颜色。

类型：三个浮点值（0.0~1.0）组成的数组；读/写。

![](https://cdn.yuelili.com/20210930234003.png)

示例：把蒙版前面颜色改为白色

```javascript
var comp = app.project.activeItem comp.layer(1).mask(1).color = [1,1,1]
```

### inverted 蒙版反转

app.project.item(index).layer(index).mask(index).inverted

描述：设置为 true 时，蒙版反转。否则为 false。

类型：布尔值，读/写。

![](https://cdn.yuelili.com/20210930235106.png)

### locked 蒙版锁定

app.project.item(index).layer(index).mask(index).locked

描述：如果为 true，则遮罩将被锁定，无法在用户界面进行编辑；否则为 false。

类型：布尔值，读/写。

![](https://cdn.yuelili.com/20210930235135.png)

### maskFeatherFalloff 羽化衰减

app.project.item(index).layer(index).mask(index).maskFeatherFalloff

描述：蒙版的羽化衰减模式。等同于 图层>蒙版>羽化衰减

类型：MaskFeatherFalloff，枚举值; 读/写：

- MaskFeatherFalloff.FFO_LINEAR：线性
- MaskFeatherFalloff.FFO_SMOOTH：平滑

![](https://cdn.yuelili.com/20210930235255.png)

### maskMode 蒙版叠加模式

app.project.item(index).layer(index).mask(index).maskMode

描述：蒙版的叠加模式。

类型：MaskMode，枚举值; 读/写：

- MaskMode.NONE：无
- MaskMode.ADD：相加
- MaskMode.SUBTRACT：相减
- MaskMode.INTERSECT：交集
- MaskMode.LIGHTEN：变亮
- MaskMode.DARKEN：变暗
- MaskMode.DIFFERENCE：差值

![](https://cdn.yuelili.com/20210930235401.png)

### maskMotionBlur 蒙版运动模糊

app.project.item(index).layer(index).mask(index).maskMotionBlur

描述：蒙版的运动模糊。等同于 图层>蒙版>运动模糊

类型：MakMotionBlur，枚举值; 读/写：

- MaskMotionBlur.SAME_AS_LAYER：与图层相同
- MaskMotionBlur.ON：开
- MaskMotionBlur.OFF：关

![](https://cdn.yuelili.com/20210930235457.png)

### rotoBezier 蒙版贝塞尔

app.project.item(index).layer(index).mask(index).rotoBezier

描述：如果为 true，则蒙版为 RotoBezier 形状；否则为假。

类型：布尔值，读/写。

![](https://cdn.yuelili.com/20210930235746.png)
