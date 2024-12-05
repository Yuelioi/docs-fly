---
title: Group
order: 5
---

# Group 组

## 说明

时间轴的所有属性都在"Group"中组织,用于共享诸如 `name` and `propertyIndex`.
Groups 可以拥有固定数量属性(比如单独一个效果,其属性不会变换)或者可变数量的属性(比如效果组本身,效果组可以包含任意数量的效果)

- **图层中的顶级组:**

  - Motion Trackers 动态追踪
  - Text 文本
  - Contents 内容
  - Masks 蒙版
  - Effects 效果
  - Transform 变换
  - Layer Styles 图层样式
  - Geometry Options 物理选项
  - Material Options 材质选项
  - Audio 音频
  - Data 数据(驱动?)
  - Essential Properties 动态图形属性

- **嵌套分组**

  - Individual effects 独立效果
  - Individual masks 独立蒙版
  - Shape groups 形状组
  - Text Animators 文字动画

---

## numProperties

用法：组.numProperties

说明：组内属性个数(不包含嵌套组内的属性)

参数：无

类型：数值

示例：查看文字组有多少个属性

```javascript
thisLayer.text.numProperties; // 4
Object.getOwnPropertyNames(thisLayer.text); // 文字图层中text组的子属性:sourceText, pathOption, moreOption,Animator
```

:::tip
效果一般使用匹配名称`thisLayer("ADBE Effect Parade").numProperties`来查看效果的数量
:::

## propertyGroup

用法：propertyGroup(`countUp=1`)

类型：方法

说明：返回父级属性组,参见 [propertyGroup(countUp=1)](property.html#property-propertygroup)

参数：countUp,父级层级,默认为 1(2 就是爷级)

类型：Group

示例：查看文字组父级

```javascript
thisLayer.text.propertyGroup(1); // 返回当前文字图层
```

## propertyIndex

用法：XX.propertyIndex

说明：返回当前属性在其父组的索引

类型：数值

示例：查看文字组在图层里的索引

```javascript
Object.getOwnPropertyNames(thisLayer); // 文字图层的属性:text,transform
text.propertyIndex; // 2
transform.propertyIndex; // 5
```

:::tip
由此可见,属性索引并非视觉上的 1 2,总体的索引大概也包含了子属性,还有些隐藏属性
:::

## name

说明:组的名称

返回:字符串

```javascript
thisLayer.text.name; // Text
```
