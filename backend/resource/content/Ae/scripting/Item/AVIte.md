---
title: AVItem
order: 1
---

# AVItem 对象

## AVItem object

app.project.item(index)

描述：AVItem 对象，用于访问 AE 中的带有音频/视频文件。A：Audio（音频） V：Video（视频）

父级关系：AVItem 是 Item 的子类。使用 AVItem 时，Item 的方法和属性均可用。

子级关系：AVItem 是合成（CompItem）和素材（footageItem）的基类，因此在使用合成和素材对象时，AVItem 属性和方法均可用。

::: danger
合成和素材虽然是 AVItem 的逻辑后代，但实际上不是 AVItem 的子类，因为 Extendscript 中不存在 AVItem。
尝试检查 AVItem 的项目 instanceof 是否会失败，因为 AVItem 未定义。 对于项目本身也是如此。
:::

## 属性

### duration 持续时间

app.project.item(index).duration

描述：返回项目持续时间(秒)。静止镜头项目的持续时间为 0。

- 合成（CompItem）：合成持续时间，读/写。
- 素材（FootageItem）：素材持续时间，只读。

类型：浮点值，0.0~10800.0

示例：返回第 4 个项目的持续时间

```javascript
alert(app.project.item(4).duration);
```

![](https://mir.yuelili.com/2021/07/ac0a347903749db48d2103de51ede74a.png)

### footageMissing 素材缺失判断

app.project.item(index).footageMissing

描述：设置为 true 时，AVItem 是占位符，或显示为素材丢失。此时，missingFootagePath 为素材丢失时的路径

请参阅 footageItem.mainSource 和 FileSource.missingFootagePath。

类型：布尔值；只读。

```javascript
app.project.item(4).footageMissing;
```

示例：如果第 4 个素材丢失了，则返回 true

### frameDuration 单帧长度

app.project.item(index).frameDuration

描述：返回 AVItem 的单帧长度（秒）。

类型：浮点值，1/99~1.0，读/写；素材文件则只读。

示例：将第一个项目的单帧设置为 0.04，也就是 25 帧/s

```javascript
app.project.item(1).frameDuration = 0.04;
```

### frameRate 帧速率

app.project.item(index).frameRate

描述：AVItem 的帧速率。

- 在合成中，读/写。
- 在素材中，只读。要改变它，请设置 mainSource 对象的 conformFrameRate。

类型：浮点值，1.0~99.0

### hasAudio 音频判断

app.project.item(index).hasAudio

描述：如果 AVItem 有音频，则为 true。

- 在合成中，该 ​​ 值链接到合成。
- 在素材中，该 ​​ 值链接到 mainSource 对象。

类型：布尔值 只读。

### hasVideo 视频判断

app.project.item(index).hasVideo

描述：如果 AVItem 有视频，则为 true。

- 在合成中，该 ​​ 值链接到合成。
- 在素材中，该 ​​ 值链接到 mainSource 对象。

类型：布尔值 只读。

### height 高度

app.project.item(index).height

描述：项目的高度(以像素为单位)。

- 在合成中，该 ​​ 值链接到合成，读/写。
- 在素材中，该 ​​ 值链接到 mainSource 对象，当 mainSource 是 SolidSource（纯色图层）时才可读/写。否则只读。

类型：整数，（1~30000）

### isMediaReplacementCompatible !不懂

app.project.item(index).isMediaReplacementCompatible

::: info 提示
此功能已添加到 After Effects 18.0(2021)中
:::

描述：测试在调用 Property.setAlternateSource()时，是否可以将 AVItem 用作备用源。如果可以使用则返回 true；否则返回 false。

可以将合成（CompItem）或素材（footageItem）用作图层的替代源，但有一些限制：

- 如果 AVItem 是素材对象，则 FootageItem.FootageSource 不应是 SolidSource 对象。
- 如果 AVItem 是素材对象，而 FootageItem.FootageSource 是 FileSource 对象，则该 FileSource 不应指向非媒体文件，例如.JSX 脚本文件。
- 设置 AVItem 不能在项目内创建循环引用。

类型：布尔值；只读。

### name 名称

app.project.item(index).name

描述：项目名称，如“项目”面板中所示。

在素材中，该 ​​ 值链接到 mainSource 对象。如果 mainSource 对象是 FileSource，则此值控制“项目”面板中的显示名称，但不影响文件名。

类型：字符串; 读/写。

示例：设置项目面板第 1 个项目的名称，并弹窗提醒

```javascript
app.project.item(1).name = "合成111";
alert(app.project.item(1).name);
```

### pixelAspect 像素长宽比

app.project.item(index).pixelAspect

描述：项目的像素长宽比(PAR)。

- 在合成中，该 ​​ 值链接到合成。
- 在素材中，该 ​​ 值链接到 mainSource 对象。

设置后检索的值可能与您提供的值略有不同。下表将 UI 中显示的值与此属性返回的更准确的值进行比较。

| PAR in AE | pixelAspect 属性返回的 PAR |
| --------- | -------------------------- |
| 0.91      | 0.90909090909091           |
| 1         | 1                          |
| 1.5       | 1.5                        |
| 1.09      | 1.09401709401709           |
| 1.21      | 1.21212121212121           |
| 1.33      | 1.33333333333333           |
| 1.46      | 1.45868945868946           |
| 2         | 2                          |

类型：浮点值，范围 0.01~100.0；读/写。

### proxySource 代理源

app.project.item(index).proxySource

描述：被用作代理的素材源对（footageSource）。只读；

改变 AVItem 代理源的方法：setProxy(), setProxyWithSequence(), setProxyWithSolid(), 或者
setProxyWithPlaceholder().

类型：素材源对象（FootageSource）；只读。

### time 项目预览时间

app.project.item(index).time

描述：直接从“项目面板”预览项目的当前时间（秒）。使用全局方法 timeToCurrentFormat()可以将其转换为以帧。如果源是静止的(item.mainSource.isStilltrue)，则无法设置时间（否则报错）。

类型：浮点值；读/写。

![](https://cdn.yuelili.com/20210912193526.png)

示例：返回项目面板第 1 个项目的当前时间轴预览时间

```javascript
alert(app.project.item(1).time);
```

### usedIn 该合成的父级

app.project.item(index).usedIn

描述：使用该 AVItem 的所有合成。请注意，在检索时，是直接复制数组，因此不会自动更新。如果获取此值，然后将此项目添加到另一个合成中，则必须重新获取，更新一下。

类型：CompItem 对象的数组；只读。

示例：合成 1 在合成 2 与合成 3 内部，下面会返回包含合成 1 的所有合成集合，然后找到其中第 1 个合成，也就是合成 2。

![](https://mir.yuelili.com/2021/07/f40c16919f33ef474567440ec0e38a27.png)

```javascript
alert(app.project.item(1).usedIn[0].name);
```

### useProxy 是否使用代理

app.project.item(index).useProxy

描述：如果该项目使用代理用，则为 true。使用 SetProxy 可以设置代理，使用 SetProxyToNone()可以取消代理。

类型：布尔值，读/写。

### width 项目宽度

app.project.item(index).width

描述：项目的宽度（像素）。

- 在 CompItem 中，该 ​​ 值链接到合成，读/写。
- 在 FootageItem 中，该 ​​ 值链接到该 mainSource 对象，并且仅当该 mainSource 对象是 SolidSource 时才可读/写。否则只读。

类型：整数，范围为 1~30000；读/写，除非另有说明。

## 方法

### setProxy() 设置代理

app.project.item(index).setProxy(file)

描述：设置文件作为代理。将指定的文件加载到新的 FileSource 对象中，将其设置为代理源（proxySource）的值，使用代理（useProxy）为 true。

不保留解释参数，而是使用用户首选项。如果文件具有未标记的 Alpha 通道，并且用户首选项询问用户如何设置，则该方法将估算 Alpha 解析，而不会询问用户。这与设置素材项目（footageItem）的 mainSource 有所不同，但是两个操作均在用户界面中执行。

参数：

- file：用作代理的文件，ExtendScript File 对象。

返回：无。

### setProxyToNone() 删除代理

app.project.item(index).setProxyToNone()

描述：从 AVItem 中删除代理，将代理源（proxySource）设置为 null，并将使用代理（useProxy）设置为 false。

参数：无。

返回：无。

### setProxyWithPlaceholder() 使用占位符设置代理

app.project.item(index).setProxyWithPlaceholder(name, width, height
,frameRate, duration)

描述：创建具有指定值的占位符对象（PlaceholderSource），将它设置为 proxySource 对象的值，然后将 useProxy 设置为 true。不保留解析参数，而是使用用户首选项。

::: info 提示
在用户界面中无法直接将占位符设置为代理。当已设置代理，然后将其移动或删除时，才会这样。
:::

参数：

- name：包含新对象名称的字符串。
- width， height：占位符的像素尺寸，4~30000 范围的整数。
- frameRate：每秒帧数，1~99 的整数
- duration：总长度(秒)，最多 3 小时。范围 0.0~10800.0 的整数。

返回：无。

### setProxyWithSequence() 使用序列设置代理

app.project.item(index).setProxyWithSequence(file,forceAlphabetical)

描述：设置文件序列作为此 AVItem 的代理，并可以选择强制按字母顺序排列。将指定的文件序列加载到新的 FileSource 对象中，将其设置为 proxySource 属性的值，然后设置 useProxy 为 true。

不保留解释参数，而是使用用户首选项。如果文件具有未标记的 Alpha 通道，并且用户首选项询问用户如何设置，则该方法将估算 Alpha 解析，而不会询问用户。

参数：

- file：序列中第一个文件，ExtendScript File 对象。
- forceAlphabetical：如果为 true，则使用“强制字母顺序”选项。

返回：无。

### setProxyWithSolid() 使用纯色图层设置代理

app.project.item(index).setProxyWithSolid(color, name, width, height,
pixelAspect)

描述：创建具有指定值的 SolidSource 对象，将此对象设置为 proxySource 属性的值，然后设置 useProxy 为 true。不保留解析参数，而是使用用户首选项。

::: info 提示
无法在用户界面将纯色设置为代理。此功能仅可 **通过脚本** 实现。
:::

参数：

- color：颜色，[R，G，B]数组，范围为[0.0~1.0]。
- name：名称，字符串。
- width， height：占位符的像素尺寸，整数，范围为[1…30000]。
- pixelAspect：像素宽高比，浮点值，范围为[0.01…100.0]范围。

返回：无。
