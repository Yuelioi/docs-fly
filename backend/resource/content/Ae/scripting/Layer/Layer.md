---
title: Layer
order: 3
---

# Layer 图层

## Layer object 图层对象

app.project.item(index).layer(index)

![](https://mir.yuelili.com/2021/07/bda0d1d04efa8415ec933f6035f5c1d5.png)

描述：图层对象用于访问合成中的图层。可以通过索引号或名称从项目的图层集进行访问。

父级关系：图层（Layer）是属性组（PropertyGroup）的子类，属性组（PropertyGroup）是属性基类（PropertyBase）的子类。属性组的方法和属性，Layer 都能用，除非 propertyIndex 设置为 undefined。

子级关系：图层（Layer）是摄像机对象（CameraLayer），灯光图层对象（LightLayer）和音视频图层对象（AVLayer）的基类，因此这几种图层类都可以使用 Layer 的属性和方法。图层除了包含 JavaScript 属性和方法外，还包含 AE 属性。有关如何访问图层中的属性的示例，请参见属性基类对象（PropertyBase
object）。

示例：如果项目中的第一项是合成（CompItem），禁用该合成中的第一个图层，并将其重命名。

```javascript
var firstLayer = app.project.item(1).layer(1);
firstLayer.enabled = false;
firstLayer.name = "DisabledLayer";
```

## 属性

### comment 注释

app.project.item(index).layer(index).comment

描述：图层的注释。注释再 “列数-注释” 显示

类型：字符串; 读/写。

示例：弹出活动项目第一个图层的注释

![](https://mir.yuelili.com/2021/07/f915b259b4febd7123f733882332a04e.png)

```javascript
alert(app.project.activeItem.layer(1).comment);
```

### containingComp 查询图层父级合成

app.project.item(index).layer(index).containingComp

描述：包含此图层的合成。

类型：CompItem 对象；只读。

示例：返回包含图层 1 的合成的名称

```javascript
alert(app.project.activeItem.layer(1).containingComp.name);
```

### hasVideo 视频判断

app.project.item(index).layer(index).hasVideo

描述：如果该图层在“时间轴”面板中，视频开关(小眼睛图标)是打开的，则为 true；否则为假。

类型：布尔值；只读。

### index 图层索引

app.project.item(index).layer(index).index

描述：图层的索引位置。

类型：整数; 只读。

### inPoint 图层入点

app.project.item(index).layer(index).inPoint

![](https://mir.yuelili.com/2021/07/6cfb17826ff5b2e39fcad3529b0c6a3f.png)

描述：图层的“入点”，以合成时间(秒)表示。

类型：浮点值，-10800.0~10800.0(负~正三个小时)；读/写。

### isNameSet 自动设置名称

app.project.item(index).layer(index).isNameSet

描述：如果名称是从源自动设置的，则为 false，否则为 true。（源自动设置是指：新建一个灰色纯色图层，系统会自动命名为 灰色纯色 1）

类型：布尔值 只读。

### **id 图层 ID**

app.project.item(index).layer(index).id

描述：图层 ID。当项目保存到文件并重新加载时，ID 保持不变。但将此项目导入另一个项目时， 所有图层会获得新 ID。ID 不会显示在普通用户界面。

::: info 提示
从 AE 2022 Beta 版本 22x25 开始加入
:::

返回：唯一且持久的标识号，用于在内部识别图层。

类型：整数，只读。

示例：

```javascript
alert(app.project.activeItem.layer(1).id); // 弹出一个整数
```

### label 图层标签

app.project.item(index).layer(index).label

![](https://mir.yuelili.com/2021/07/659cb125f55e2b22dd80a6d72f47283a.png)

描述：图层的标签颜色。由其数字表示(“ 0”代表“无”，“ 1 至 16”代表“标签”首选项中预设颜色编号)。

::: info 提示
自定义标签颜色无法通过编程进行设置。
:::

类型：整数(0 到 16)，读/写。

### locked 图层锁定

app.project.item(index).layer(index).locked

描述：图层是否被锁定。对应于“图层”面板中的锁定开关。

类型：布尔值，读/写。

### nullLayer 空对象图层

app.project.item(index).layer(index).nullLayer

描述：如果该图层为空对象则为 true。否则为 false。

类型：布尔值，只读。

### outPoint 图层出点

app.project.item(index).layer(index).outPoint

![](https://mir.yuelili.com/2021/07/098d23faf5186af17c348a858ffb5e0b.png)

描述：图层的“出点”，以合成时间(秒)表示。

类型：浮点值，范围为-10800.0~10800.0(负~正三个小时)；读/写。

### parent 图层设置父级

app.project.item(index).layer(index).parent

![](https://mir.yuelili.com/2021/07/4e03af348f0132bc4c68e978a2d7562b.png)

描述：该图层的父级；可以不设置（返回 null）。等同于直接拖小鞭子到父对象。要设置父级而不更改子层的变换值，请使用 setParentWithJump 方法（等同于按住 shift
再用小鞭子拖）。

类型：图层对象或 null ，读/写。

### selectedProperties 图层选择属性

app.project.item(index).layer(index).selectedProperties

描述：包含图层中所有当前选定的 Property 和 PropertyGroup 对象的数组。

类型：PropertyBase 对象数组；只读。

### shy 图层消隐

app.project.item(index).layer(index).shy

![](https://mir.yuelili.com/2021/07/fe00f4e4178833850bd83e2a159282f0.png)

描述：如果该图层为“消隐”，则为 true，否则为 false。

类型：布尔值，读/写。

### solo 图层独奏

app.project.item(index).layer(index).solo

描述：如果单独显示该层则为 true，否则为 false。

类型：布尔值，读/写。

### startTime 图层开始时间

app.project.item(index).layer(index).startTime

描述：图层的开始时间，以合成时间(秒)表示。

类型：浮点值，范围为-10800.0..10800.0(负~正三个小时)；读/写。

### stretch 图层拉伸

app.project.item(index).layer(index).stretch

描述：图层的时间范围，以百分比表示。100 表示 ​​ 无拉伸。0 和 1 之间的值设置为 1，-1 和 0 之间的值(不包括 0)设置为-1。

类型：浮点值，范围为-9900.0~9900.0；读/写。

### time 图层当前时间

app.project.item(index).layer(index).time

描述：图层的当前时间，以合成时间(秒)表示。

类型：浮点值；只读。

## 方法

### activeAtTime() 活动状态

app.project.item(index).layer(index).activeAtTime(time)

描述：如果此层在指定时间处于活动状态，则返回 true。必须启用该层或和其他图层一起独奏，并且时间必须在该层的 inPoint 和 outPoint 值之间。

参数：

- time：时间(秒)，浮点值。

返回：布尔值。

### applyPreset() 应用预设

app.project.item(index).layer(index).applyPreset(presetName);

描述：将动画预设应用于图层。动画预设文件夹在“preset”中，用户可以通过用户界面创建新的动画预设。

::: danger
好像是把预设应用到选择的图层！不是你自己 var 的图层（未测试）
:::

参数：

- presetName：动画预设文件，ExtendScript File 对象，需要使用绝对路径。

返回：无。

### copyToComp() 图层复制到合成

app.project.item(index).layer(index).copyToComp(intoComp)

描述：将图层复制到指定的合成中。原始层保持不变。这与用户界面复制并粘贴图层相同。

::: info 提示
从 After Effects 13.6 开始，当图层具有父级时，此方法不再导致 After Effects 崩溃。
:::

::: danger
从 After Effects 13.7(未测试 13.6)开始，如果复制的图层对其具有影响，并且用户撤消该操作，则 After Effects 将崩溃。
:::

技巧：它会复制到第一个选定图层的上方，未选择图层时，会复制合成顶部的图层。要检索复制过的图层，可以检索当前选择图层（CompItem.selectedLayers）中索引最大，然后用索引减 1 即可。并使用它来检索。comp.layer(
topmost_index_of_selected_layers - 1 )

参数：

- intoComp 目标合成

返回：无。

案例：复制合成中第 2 个图层，并获取副本图层

```javascript
var comp = app.project.activeItem;
comp.layer(2).copyToComp(comp); // 复制第2个图层
copylayer = comp.layer(2 - 1); // 因为副本总在最上面，所以新图层索引是2-1 = 1
```

## duplicate() 图层复制

```javascript
app.project.item(index).layer(index).duplicate();
```

描述：复制图层。等同于选择图层“编辑>重复(Ctrl D)”。不同之处在于，使用此方法时，复制以后不会选择刚刚复制的图层（还是选择原来的图层）。

参数：无。

返回：图层对象。

## layerByID() 通过 ID 获取图层

app.project.layerByID(id)

描述：将图层移动到指定图层之后(下方)。

::: info 提示
从 AE 2022 Beta 版本 22x25 开始加入
:::

参数：id ，图层的唯一识别号。整数。

返回：图层对象（如果有的话）；否则为 _null_ 。无效 ID 将引发异常。

## moveAfter() 移动图层后

app.project.item(index).layer(index).moveAfter(layer)

描述：将图层移动到指定图层之后(下方)。

参数：layer ，指定图层。（不能跨合成）

返回：无。

## moveBefore() 移动图层前

app.project.item(index).layer(index).moveBefore(layer)

描述：将图层移动到紧接在指定层之前(上方)的位置。

参数：layer ，指定图层。（不能跨合成）

返回：无。

## moveToBeginning() 移动图层顶

app.project.item(index).layer(index).moveToBeginning()

描述：将图层移动到最顶层(第一层)。

参数：无。

返回：无。

## moveToEnd() 移动图层底

app.project.item(index).layer(index).moveToEnd()

描述：将图层移动到最底层(最后一层)。

参数：无。

返回：无。

## remove() 删除图层

app.project.item(index).layer(index).remove()

描述：从合成中删除指定的图层。

参数：无。

返回：无。

## setParentWithJump() 图层设置父级（跳跃）

app.project.item(index).layer(index).setParentWithJump([newParent])

描述：将此层的父级设置为指定图层，而不更改子层的变换值。等同于按住 shift 再拖拽小鞭子。

参数：

- newParent：可选，同合成的图层对象。如果未指定，则将父级设置为“无”。

返回：无。
