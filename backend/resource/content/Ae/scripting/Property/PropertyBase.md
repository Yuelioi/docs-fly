---
title: PropertyBase
order: 3
---

# PropertyBase 属性基类

## PropertyBase object 属性基类

app.project.item(index).layer(index).propertySpec

描述：根据图层，按照属性名称访问属性。例如，以下是“效果”组中访问属性的方式

```javascript
var effect1 = app.project.item(1).layer(1).effect("AddGrain")("Viewing Mode");
var effect2 = app.project.item(1).layer(1).effect.addGrain.viewingMode;
var effect3 = app.project.item(1).layer(1)("Effects").addGrain.viewingMode;
var effect4 = app.project.item(1).layer(1)("Effects")("Add Grain")("Viewing Mode");
```

另请参见 PropertyGroup.property()。

PropertyBase 是 Property 和 PropertyGroup 的基类，因此在使用属性和属性组时，可以使用 PropertyBase 属性和方法。

### 引用无效

如果某对象无法引用，则引用改对象的脚本可能报错。例如先删除一个对象，再访问该对象会报错：

```javascript
var selLayer = app.project.activeItem.selectedLayers[0]; //当前选择的图层
selLayer.remove(); //移除该图层
alert(selLayer.name); //调用该图层名称，报错
```

同样，如果引用已删除图层的属性，也会报错

```javascript
var selLayer = app.project.activeItem.selectedLayers[0]; //当前选择的图层
selLayer.remove(); //移除该图层
alert(selLayer.transform.position); //调用该图层位置，报错
```

从属性组中删除属性。由于索引位置更改，因此“对象无效”。例如：

```javascript
var effect1 = app.project.item(1).layer(1).effect(1);
var effect2 = app.project.item(1).layer(1).effect(2);
var effect2param = app.project.item(1).layer(1).effect(2).blendWithOriginal;
effect1.remove(); //删掉第1个效果
alert(effect2.name); // 调用第2个效果名称报错，因为现在第2个效果跑到第1了
```

## 属性

### active 活动

app.project.item(index).layer(index).active  
app.project.item(index).layer(index).propertySpec.active

描述：判断是否处于活动状态。

活动状态：

- 图层：小眼睛图标打开，且时间轴处于图层入点和出点内
- 效果和属性：fx 图标的开关。

类型：布尔值，只读。

示例：图层的写入效果未打开，故返回 false（由于截图原因，其实 Null 需要选中的）

![](https://cdn.yuelili.com/20210914160859.png)

```javascript
var seL = app.project.activeItem.selectedLayers; //当前选择的图层
alert(seL[0].effect(1).active); // 判断它的第1个效果是否激活
```

### canSetEnabled 能否可视

app.project.item(index).layer(index).propertySpec.canSetEnabled

描述：如果为 true，则可设置 enabled 属性值。也就是有小眼睛图标，否则为 false。

类型：布尔值，只读。

示例：“写入效果”的“颜色” 参数是没有可视控制的，因此返回 false

![](https://cdn.yuelili.com/20210914161628.png)

```javascript
var a = app.project.activeItem.selectedLayers;
alert(a[0].effect(1)("Color").canSetEnabled); //返回false
```

### elided 隐藏的父组

app.project.item(index).layer(index).propertySpec.elided

描述：如果为 true，则此属性是其他属性的父组。而且不会显示在用户界面中，其子属性也未在面板上缩进。

示例：对于具有两个动画器且不展开属性的文本层，可能会看到：

![](https://mir.yuelili.com/2021/07/ae38358173a69384cba8cef13d4a1c56.png)

但实际的路径为：

- Text
  - Source Text （源文本）
  - Path Options （路径选项）
  - MoreOptions （更多选项）
  - **Animators** （ **在 AE 界面不显示** ）
    - Animator 1 （动画制作工具 1）
    - Animator 2 （动画制作工具 2）

“ Animator 1”和“ Animator 2”的父组为：“Animators”。但 Animators 不会显示在 AE 界面中。

类型：布尔值，只读。

示例：找到文字图层隐藏的属性组，并判断

```javascript
var txtLayer = app.project.activeItem.selectedLayers[0]
alert(txtLayer.text(4).name) // 返回 Animators alert(txtLayer.text(4).elided) //
返回true，说明它是隐藏的组
```

因此，想访问 Animator 1

需要：txtLayer.text(4)(1)

### enabled 启用

app.project.item(index).layer(index).enabled  
app.project.item(index).layer(index).propertySpec.enabled

![](https://cdn.yuelili.com/20210914160859.png)

说明：为 true 时，将启用图层或属性；为 false 则不启用。

启用状态：

- 图层：小眼睛图标打开
- 效果和属性：fx 图标的开关。

类型：布尔值，如果 canSetEnabled 为 true，为读/写；否则只读。

### isEffect 是否为效果属性组

app.project.item(index).layer(index).propertySpec.isEffect

描述：为 true 时，该属性为效果属性组（PropertyGroup）。

类型：布尔值，只读。

示例：其他的以后研究

```javascript
var txtLayer = app.project.activeItem.selectedLayers[0];
alert(txtLayer.effect(1).isEffect); // 写入效果 是 属于“属性组”的
```

### isMask 是否为蒙版组

app.project.item(index).layer(index).propertySpec.isMask

描述：为 true 时，该属性为蒙版属性组（PropertyGroup）。

类型：布尔值，只读。

示例：Mask 1 就是蒙版组。Masks 不是蒙版组

![](https://cdn.yuelili.com/20210914163832.png)

```javascript
var txtLayer = app.project.activeItem.selectedLayers[0];
alert(txtLayer(3)(1).isMask);
```

### isModified 是否更改

app.project.item(index).layer(index).propertySpec.isModified

描述：为 true 时，该属性自创建以来已更改。

类型：布尔值，只读。

### matchName 匹配名称

app.project.item(index).layer(index).propertySpec.matchName

描述：每个属性都有一个唯一的匹配名称（matchName）标识符。更新版本，更换语言，均不会变换。一般用于制作模板或修正表达式错误

类型：字符串; 只读。

示例：读取当前图层第一个效果的匹配名称

![](https://cdn.yuelili.com/20210914164043.png)

```javascript
var txtLayer = app.project.activeItem.selectedLayers[0];
alert(txtLayer.effect(1).matchName); // 返回 ADBE_FOG_3D
```

### name 名称

app.project.item(index).layer(index).name  
app.project.item(index).layer(index).propertySpec.name

描述：

- 图层：图层名称。默认情况下，与源名称相同，除非自己改名（也就是 Layer.isNameSet 返回 false）。
- 效果和属性：效果或属性的名称。如果不是索引组的子级，则无法设置名称值。索引组请参见 PropertyBase.propertyType。

类型：字符串; 是索引组的子级则可以读/写；否则为只读。

### parentProperty 父级属性组

app.project.item(index).layer(index).propertySpec.parentProperty

描述：该属性父级的属性组；如果该属性是图层，则返回 null。

类型：属性组（PropertyGroup）对象或 null。只读。

示例：查找父级属性名称

![](https://cdn.yuelili.com/20210914163832.png)

```javascript
var txtLayer = app.project.activeItem.selectedLayers[0];
alert(txtLayer("Masks")("Mask 1").parentProperty.name); //Mask 1的父级为： Masks
```

### propertyDepth 父组深度

app.project.item(index).layer(index).propertySpec.propertyDepth

描述：属性与图层之间的级别数。图层的级别为 0。

类型：整数; 只读。

示例：图层的第 1 个效果与图层本身之间，父组级别为 2

```javascript
var selLayer = app.project.activeItem.selectedLayers[0]; //当前选择的图层
alert(selLayer.effect(1).propertyDepth);
```

### propertyIndex 属性索引

app.project.item(index).layer(index).propertySpec.propertyIndex

描述：该属性在其父组中的位置索引，前提：它是索引组的子级(索引组：请参阅 PropertyBase.propertyType)。

类型：整数; 只读。

### propertyType 属性类型

app.project.item(index).layer(index).propertySpec.propertyType

描述：该属性的类型。

类型：PropertyType，枚举值; 读/写：

- PropertyType.PROPERTY：属性，单个属性，例如位置或缩放。
- PropertyType.INDEXED_GROUP：索引组，属性组，其成员的名称和索引可编辑。效果和蒙版就是索引组。
- PropertyType.NAMED_GROUP：命名组，成员名称不可编辑的属性组。图层就是命名组。

关于索引组：

效果组的成员 - 也就是某一个效果，该效果名称可编辑

蒙版组的成员 - 也就是某一个蒙版，该蒙版名称可编辑

![](https://cdn.yuelili.com/20210914164802.png)![](https://cdn.yuelili.com/20210914164836.png)

### selected 属性选择

app.project.item(index).layer(index).propertySpec.selected

描述：如果为 true，则选择该属性，为 false 则取消选择。大量属性使用此法会降低系统性能。要读取合成或图层的所选属性的完整集合，请使用 CompItem.selectedProperties 或 Layer.selectedProperties。

类型：布尔值，读/写。

## 方法

### duplicate() 属性复制

app.project.item(index).layer(index).propertySpec.duplicate()

描述：如果此属性是索引组的子级，创建并返回与相同的属性值的新 PropertyBase 对象。如果不是索引组的子级，则报错。

索引组请参见 PropertyBase.propertyType。

参数：无。

返回：PropertyBase 对象。

### moveTo() 属性移动

app.project.item(index).layer(index).propertySpec.moveTo(newIndex)

描述：将此属性移动到其父属性组中的新位置。此方法仅对索引组的子级有效。如果不是索引组的子级，则报错。

索引组请参见 PropertyBase.propertyType。

::: danger
:::

使用此方法会改变其他子级的索引。例如图层有三个效果，移动其中一个效果会使其他效果索引发生变化。

参数：

- newIndex：放置在其父组的新索引位置，整数。

返回：无。

### propertyGroup() 查找属性父组

app.project.item(index).layer(index).propertySpec.propertyGroup([countUp])

描述：获取属性父组对象（PropertyGroup）。

参数：

- countUp：向上查找的辈数。选填，整数，范围为 1~propertyDepth。，默认为 1，获取父级（“爷级”就是 2）。

返回：PropertyGroup 对象；如果越界，则为 null。

### remove() 属性删除

app.project.item(index).layer(index).propertySpec.remove()

描述：从其父组中删除该属性。如果这是一个属性组，还也会删除子属性。此方法仅对索引组的子级有效。如果不是索引组的子级，则报错。可以在文本动画属性(文字动画制作工具)上调用此方法。

索引组请参见 PropertyBase.propertyType。

参数：无。

返回：无。
