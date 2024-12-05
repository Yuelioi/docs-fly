---
title: FootageItem
order: 4
---

# FootageItem 素材

## FootageItem object

app.project.item(index)  
app.project.items[index]

描述：footageItem，素材项目对象，代表导入到项目中的素材项目，在“项目”面板中显示。可以通过项目集的位置索引进行访问。

父级关系：FootageItem 是 AVItem 对象的子类，AVItem 对象是 Item 对象的子类。除以下列出的方法外，AVItem 和 Item 的所有方法和属性，在均可在 FootageItem 使用。

使用方式：footageItem.XX → app.project.item(index).XX

## 属性篇

### file 素材源文件

app.project.item(index).file

描述：素材源文件，ExtendScript
File 对象。如果 FootageItem 的 mainSource 是文件源，则与 FootageItem.mainSource.file 相同。否则为 null。

类型：文件对象；只读。

### mainSource 素材源对象

app.project.item(index).mainSource

描述：素材源对象，包含与该素材项目相关的所有设置，包括“文件 - 解释素材”里的设置。只读，若要更改，请用 FootageItem 的“
replace”方法。请参见 footageSource 对象及其三种类型：

- SolidSource 对象
- FileSource 对象
- PlaceholderSource 对象

如果是 FileSource 对象，并且 footageMissing 值为 true，则缺少的素材文件的路径在 FileSource.missingFootagePath 属性中。

类型：footageSource 对象; 只读。

## 方法

### openInViewer() 在查看器打开素材

app.project.item(index).openInViewer()

描述：在“素材”面板中打开素材，然后将“素材”面板移到最前面并为其聚焦。

::: info 提示
丢失素材和占位符素材可以使用此脚本方法打开，但不能手动打开(通过双击)。
:::

参数：无。

返回：“素材”面板的查看器对象；如果无法打开素材，则为 null。

### replace() 更改素材源

app.project.item(index).replace(file)

描述：将 FootageItem 的源更改为指定文件。除了加载文件之外，该方法还会为文件创建一个新的 FileSource 对象，并将 mainSource 设置到该对象。在新的源对象，可以设置 name，width，height，frameDuration，和 duration 属性(详见 AVItem 对象)基于文件的相关内容。该方法保留了先前 mainSource 对象的解释参数。如果指定的文件具有未标记的 Alpha 通道，则该方法将估算 Alpha 解释。

参数：

- file：用于素材主要源的文件，ExtendScript File 对象。

### replaceWithPlaceholder() 占位符替换素材

app.project.item(index).replaceWithPlaceholder(name, width, height, frameRate,
duration)

描述：将此 FootageItem 的源更改为指定的占位符。创建一个新的占位符源（PlaceholderSource）对象，从参数中设置其值，然后设置 mainSource 到该对象。

参数：

- name ：字符串，包含占位符的名称。
- width ：占位符的宽度(像素)，整数,范围为 4~30000
- height ：占位符的高度(像素)，整数，范围为 4~30000
- frameRate ：占位符的帧速率，范围为 1.0.~99.0
- duration 占位符的持续时间(像素)，浮点数，范围为 0.0~10800.0

### replaceWithSequence() 序列替换素材

app.project.item(index).replaceWithSequence(file, forceAlphabetical)

描述：将此素材项目（FootageItem）的来源更改为指定的图像序列。除了加载文件之外，该方法还会为文件创建一个新的 FileSource 对象，并将 mainSource 设置到该对象。在新的源对象，可以设置 name，width，height，frameDuration，和 duration 属性(详见 AVItem 对象)基于文件的相关内容。该方法保留了先前 mainSource 对象的解释参数。如果指定的文件具有未标记的 Alpha 通道，则该方法将估算 Alpha 解释。

参数：

- file：序列中第一个文件的 ExtendScript File 对象，用作素材的主要来源。
- forceAlphabetical：如果为 true，则使用“强制字母顺序”选项。

### replaceWithSolid() 纯色层替换素材

app.project.item(index).replaceWithSolid(color, name, width, height,
pixelAspect)

描述：将此 FootageItem 的源更改为指定的纯色层。创建一个新的纯色源（SolidSource）对象，从参数中设置其值，然后把 mainSource 设置到该对象。

参数：

- color：纯色层颜色，三个浮点值的数组，范围为[R, G, B][0.0~1.0]。比如[0.5,0.5,0.5]代表灰色
- name：纯色层名称，字符串
- width：纯色层宽度(像素)，整数，范围为： 4~30000
- height：纯色层高度(像素)，整数，范围为： 4~30000
- pixelAspect：纯色层像素长宽比，浮点数，范围为 0.01~100.0
