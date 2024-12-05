---
order: 2
title: ImportOptions
---

# ImportOptions 导入设置

## 导入选项

new ImportOptions();  
new ImportOptions(file);

描述：ImportOptions 对象封装了用于使用 Project.importFile()方法导入文件的选项。构造函数采用一个可选参数，即文件的 ExtendScript
File 对象。如果未提供，则 file 在将对象与 importFile 方法一起使用之前，必须显式设置属性的值。

例如：new ImportOptions().file = new File("myfile.psd");

## 属性

### file 导入文件

importOptions.file

描述：要导入的文件。如果在构造函数中设置了文件，则可以通过此属性访问它。

类型：ExtendScript File 对象；读/写。

### forceAlphabetical 强制字母顺序

importOptions.forceAlphabetical

描述：如果为 true，等同于“文件>导入>文件”中设置“强制字母顺序排列”

类型：布尔值 读/写。

### importAs 作为某源导入

importOptions.importAs

描述：导入文件将作为源的对象的类型。设置之前，使用 canImportAs 检查是否可以将给定文件作为给定对象类型的源导入。

类型：ImportAsType，枚举值; 读/写：

- ImportAsType.COMP_CROPPED_LAYERS
- ImportAsType.FOOTAGE
- ImportAsType.COMP
- ImportAsType.PROJECT

### rangeEnd 导入序列结束剪切

importOptions.rangeEnd

::: danger
此方法/属性无正式记录，是否存在以及是否报错未知，如果有更多信息，请贡献！

:::

描述：设置要导入的序列的结束剪切范围。

如果 rangeEnd 超过要导入的序列的持续时间，则创建“丢失帧”(视频条)。

如果 sequence 设置为 false，则无效。

如果将 forceAlphabetical 设置为 true，则引发异常。

如果 rangeEnd 小于 rangeStart，则引发异常，并重置范围以包括所有文件。

类型：整数; 读/写。

### rangeStart 导入序列起始剪切

importOptions.rangeStart

::: danger
此方法/属性无正式记录，是否存在以及是否报错未知，如果有更多信息，请贡献！

:::

描述：设置要导入的序列的开始剪切范围。

如果 sequence 设置为 false，则无效。

如果将 forceAlphabetical 设置为 true，则引发异常。

如果 rangeEnd 值为 0，则引发异常。

如果 rangeStart 大于 rangeEnd 则引发异常，并将范围重置为包括所有文件。

类型：整数; 读/写。

示例

```javascript
// Import 20 frames of the sequence, starting at frame 10 and ending at frame 30

var mySequence = "~/Desktop/sequence/image_000.png";

var importOptions = new ImportOptions();
importOptions.file = new File(mySequence);
importOptions.sequence = true;
importOptions.forceAlphabetical = false;
importOptions.rangeStart = 10;
importOptions.rangeEnd = 30;

var item = app.project.importFile(importOptions);
```

### sequence 导入序列

importOptions.sequence

描述：如果为 true，则导入序列；若为 false，将导入单个文件。

类型：布尔值 读/写。

## 方法

### canImportAs() 能否作为某源导入

importOptions.canImportAs(type)

描述：报告是否可以将文件作为特定对象类型的源导入。如果此方法返回 true，则可以将给定类型设置为 importAs 属性的值。

参数：

- type，可以导入的文件类型。
- ImportAsType，枚举值：
  - ImportAsType.COMP
  - ImportAsType.FOOTAGE
  - ImportAsType.COMP_CROPPED_LAYERS
  - ImportAsType.PROJECT

返回：布尔值。

示例

```javascript
var io = new ImportOptions(new File("c:\\myFile.psd"));
if (io.canImportAs(ImportAsType.COMP)) {
  io.importAs = ImportAsType.COMP;
}
```

### isFileNameNumbered() 文件名是否含数字

importOptions.isFileNameNumbered(file)

::: danger
此方法/属性无正式记录，是否存在以及是否报错未知，如果有更多信息，请贡献！

:::

描述：报告文件对象是否编号，即文件名是否有数字。

参数：file ExtendScript 文件对象。

返回：对象，包含 2 个键：

- isNumbered：布尔值; 文件名是否包含任何数字，
- num： 整数; 在文件名中找到一个数字。如果文件名没有数字 ，则返回 0 。

示例

```JavaScript
var importOptions = new ImportOptions();
importOptions.isFileNameNumbered('image.png'); // "isNumbered": false, "num": 0
importOptions.isFileNameNumbered('003image.png'); // "isNumbered": true, "num": 3
importOptions.isFileNameNumbered('ima0102ge.png'); // "isNumbered": true, "num": 102
importOptions.isFileNameNumbered('image0120.png'); // "isNumbered": true, "num": 120
```
