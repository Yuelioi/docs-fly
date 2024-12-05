---
order: 4
title: PropertyGroup
---

# PropertyGroup 属性组

## PropertyGroup object

app.project.item(index).layer(index).propertyGroupSpec

描述：PropertyGroup（属性组）对象代表一组属性。它可以包含属性和其他属性组。可以嵌套属性组，其中 Layer 对象位于最顶部(root)，再到单个属性对象，例如蒙版 - 蒙版羽化。

说明：要遍历组层次结构，使用 PropertyBase 方法和属性，参见 PropertyBase.propertyGroup()。访问属性和属性组，请参见 PropertyBase
object。

父级关系：属性组是属性基类（PropertyBase）的子类。PropertyBase 的所有方法和属性均可用。

子级关系：属性组是 Layer 和蒙版属性组（MaskPropertyGroup）的基类。使用图层或蒙版组时，属性组的属性和方法均可用。

## 属性

### numProperties 属性数量

app.project.item(index).layer(index).propertyGroupSpec.numProperties

描述：组中索引属性的数量。

比如图层的属性数量为 3，对应蒙版，效果和运动跟踪器组，它们是图层的索引组。

但是，图层还具有许多其他只能按名称使用的属性（命名组）。请参阅 PropertyGroup.property()。

类型：整数; 只读。

示例：查看选择图层的属性数量（返回 14）

![](https://cdn.yuelili.com/20210916221245.png)

```javascript
var selLayer = app.project.activeItem.selectedLayers[0];
alert(selLayer.numProperties);
```

## 方法

### addProperty() 添加属性

app.project.item(index).layer(index).propertyGroupSpec.addProperty(name)

描述：创建属性基类（PropertyBase）对象，并将其添加到该组中。

通常，只能将属性添加到索引组。唯一例外是文本动画设计器属性，可以将其添加到命名组中(索引组和命名组请参阅 PropertyBase.propertyType)。

如果无法创建具有指定名称的属性，则会生成异常。

是否可以向该组添加属性，请在调用前先使用 canAddProperty。请参阅 PropertyGroup.canAddProperty()

::: danger
向索引组添加新属性后，将重新创建索引组，从而组中原来属性的引用会失效。
:::

解决方法之一：将添加的属性的索引存储在 property.propertyIndex 中。

示例：一旦添加颜色控制属性，滑块就失效：

![](https://cdn.yuelili.com/20210916221840.png)

```javascript
var selLayer = app.project.activeItem.selectedLayers[0]; //当前选择的图层
var effectsProperty = selLayer.property("ADBE Effect Parade");
var slider = effectsProperty.addProperty("ADBE Slider Control"); // 添加滑块
var color = effectsProperty.addProperty("ADBE Color Control"); //添加颜色
var sliderProperty = slider.property("ADBE Slider Control-0001"); // 引用滑块失效
```

修改后：

```javascript
var selLayer = app.project.activeItem.selectedLayers[0]; //当前选择的图层
var effectsProperty = selLayer.property("ADBE Effect Parade");
var slider = effectsProperty.addProperty("ADBE Slider Control"); // 添加滑块
var sliderIndex = slider.propertyIndex; // 储存滑块索引
var color = effectsProperty.addProperty("ADBE Color Control"); //添加颜色
var sliderProperty = effectsProperty.property(sliderIndex).property("ADBE Slider Control-0001");
//正常引用
```

参数：name：要添加属性的显示名称或匹配名称。(请参阅 PropertyBase.matchName)。支持以下名称：

- 属性的匹配名称（match name）。例如：ADBE Mask Atom、ADBE Paint Atom、ADBE Text Position、ADBE Text Anchor Point
- ADBE Mask Parade（蒙版），例如：ADBE Mask Atom、Mask
- ADBE Parade（效果），具有匹配名称的效果，例如：ADBE Bulge、ADBE Glo2、APC Vegas”
- 效果显示名称，如：Bulge、Glow、Vegas
- 文本动画制作工具，如：ADBE Text Animator
- 选择器，范围选择器的名称为 ADBE Text Selector，摆动选择器的名称为 ADBE Text Wiggly Selector，表达式选择器的名称为 ADBE Text Expressible Selector

返回：PropertyBase 对象。

### canAddProperty() 可添加属性

app.project.item(index).layer(index).propertyGroupSpec.canAddProperty(name)

描述：如果可以将属性添加到此属性组，则返回 true。

例如只能将蒙版添加到蒙版组中。唯一合法的输入参数是“ mask”或“ ADBE Mask Atom”。

假设定义了蒙版组为 maskGroup

- maskGroup.canAddProperty("mask"); // true
- maskGroup.canAddProperty("ADBE Mask Atom"); // true
- maskGroup.canAddProperty("blend"); // false

参数：

- name：属性的显示名称或匹配名称。请参阅 PropertyGroup.addProperty()

返回：布尔值。

### property() 查找属性

app.project.item(index).layer(index).propertyGroupSpec.property(index)  
app.project.item(index).layer(index).propertyGroupSpec.property(name)

描述：查找并返回该组的子属性，由其索引或名称指定。引用方式与表达式语法相同。以下内容均正确且等效：

假设定义(var)了一个图层为 myLayer

- mylayer.position;
- mylayer("position");
- mylayer.property("position");
- mylayer(1); //假设 position 的索引为 1
- mylayer.property(1);

图层的某些属性(例如位置和缩放)只能通过名称访问。当使用名称向下查找多个级别的属性时，必须多次调用此方法。

例如，以下调用向下查找 2 个级别，返回蒙版组的第一个蒙版： myLayer.property("ADBE Masks").property(1)

参数：

index ：子属性在索引组的索引。整数，范围为 1~属性数量。  
name：子属性的名称。支持：

- 匹配名称。如 ADBE Mask Parade
- 表达式“括号样式”，表示显示名称或紧凑的英文名称，如：Anchor Point
- 表达式“ intercap style”语法中的任何名称。如： depthOfField

::: info 提示
Intercap style ：所有词连接在一起，中间没有空格，并且每个词的首字母都大写

有关受支持的属性名称，请参见下表。
:::

返回：PropertyBase 对象；如果未找到，则为 null。

可通过名称访问的属性

| 来源                                                                 | 匹配名与显示名称                             | 中文介绍   |
| -------------------------------------------------------------------- | -------------------------------------------- | ---------- |
| 图层                                                                 | ADBE Mask Parade、Masks                      | 蒙版       |
| ADBE Effect Parade、Effects                                          | 效果                                         |
| ADBE MTrackers、Motion Trackers                                      | 运动追踪                                     |
| AVLayer                                                              | Anchor Point、anchorPoint                    | 锚点       |
| Position、position                                                   | 位置                                         |
| Scale、scale                                                         | 缩放                                         |
| Rotation、rotation                                                   | 旋转                                         |
| Z Rotation、zRotation、Rotation Z、rotationZ                         | Z 旋转 旋转 Z                                |
| Opacity、opacity                                                     | 不透明度                                     |
| Marker、marker                                                       | 标记/关键帧                                  |
| 非静态源的 AVLayer                                                   | Time Remap、timeRemapEnabled                 | 时间重映射 |
| 音频组件的 AVLayer                                                   | Audio Levels、audioLevels                    | 音频电平   |
| 相机图层                                                             | Zoom、zoom                                   | 缩放       |
| Depth of Field、depthOfField                                         | 景深                                         |
| Focus Distance、focusDistance                                        | 焦距                                         |
| Aperture、aperture                                                   | 光圈                                         |
| Blur Level、blurLevel                                                | 模糊等级                                     |
| 灯光图层                                                             | Intensity、intensity                         | 强度       |
| Color、color                                                         | 颜色                                         |
| Cone Angle、coneAngle                                                | 锥形角度                                     |
| Cone Feather、coneFeather                                            | 锥形羽化                                     |
| Shadow Darkness、shadowDarkness                                      | 阴影深度                                     |
| Shadow Diffusion、shadowDiffusion                                    | 阴影扩散                                     |
| Casts Shadows、castsShadows                                          | 投影                                         |
| 3D 图层                                                              | Accepts Shadows、acceptsShadows              | 接受阴影   |
| Accepts Lights、acceptsLights                                        | 接受灯光                                     |
| Ambient、ambient                                                     | 环境                                         |
| Diffuse、diffuse                                                     | 漫射                                         |
| Specular、specular (these are for the Specular Intensity property)   | 镜面强度                                     |
| Shininess、shininess (these are for the Specular Shininess property) | 镜面反光度                                   |
| Casts Shadows、castsShadows                                          | 投影                                         |
| Light Transmission、lightTransmission                                | 透光率                                       |
| Metal、metal                                                         | 金属质感                                     |
| 相机、灯光、3D 图层                                                  | X Rotation、xRotation、Rotation X、rotationX | 旋转       |
| Y Rotation、yRotation、Rotation Y、rotationY                         | 旋转                                         |
| Orientation、orientation                                             | 方向                                         |
| 文字图层                                                             | Source Text、source Text、Text、text         | 源文本     |
| ADBE Mask Parade 属性组                                              | ADBE Mask Atom                               | 蒙版       |
| ADBE Mask Atom 属性组                                                | ADBE Mask Shape、maskShape、maskPath         | 蒙版路径   |
| ADBE Mask Feather、maskFeather                                       | 蒙版羽化                                     |
| ADBE Mask Opacity、maskOpacity                                       | 蒙版不透明度                                 |
| ADBE Mask Offset、maskOffset                                         | 蒙版偏移/扩展                                |

示例：

图层具有“快速方框模糊”效果，则可以通过以下任意一种方式来检索该效果：

```javascript
var myLayer = app.project.activeItem.selectedLayers[0]; //当前选择的图层
alert(myLayer.property("Effects").property("Fast Box Blur").name); //显示名称

alert(myLayer.property("Effects").property("fastBoxBlur").name); // intercap 风格名称

alert(myLayer.property("Effects").property("ADBE Box Blur2").name); //匹配名
```

图层具有名为“ Mask 1”的蒙版，则可以按以下方式进行检索：

```javascript
var myLayer = app.project.activeItem.selectedLayers[0]; //当前选择的图层
myLayer.property("Masks").property("Mask 1");
```

从凸出（Bulge）效果获取“凸出中心”参数，可以使用以下任一方法：

```javascript
var myLayer = app.project.activeItem.selectedLayers[0]; //当前选择的图层
myLayer.property("Effects").property("Bulge").property("Bulge Center");

myLayer.property("Effects").property("Bulge").property("bulgeCenter");
```
