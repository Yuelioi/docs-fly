---
title: OutputModule
order: 2
---

# OutputModule 对象

## 输出模块对象

app.project.renderQueue.item(index).outputModule(index)

描述：渲染队列项目（RenderQueueItem）的输出模块对象（OutputModule）通过渲染操作生成单个文件或序列，并包含与要渲染的文件有关的属性和方法。

## 属性

### file 不懂

app.project.renderQueue.item(index).outputModule(index).file

描述：此输出模块设置为呈现的文件的 ExtendScript File 对象。

类型：ExtendScript File 对象；读/写。

### includeSourceXMP

app.project.renderQueue.item(index).outputModule(index).includeSourceXMP

描述：如果为 true，则将所有源素材 XMP 元数据写入输出文件。对应于“输出模块>输出模块设置>包含源 XMP 元数据”。

类型：布尔值 读/写。

### name

app.project.renderQueue.item(index).outputModule(index).name

描述：输出模块的名称，如用户界面中所示。

类型：字符串; 只读。

### postRenderAction

app.project.renderQueue.item(index).outputModule(index).postRenderAction

描述：渲染操作完成时要执行的动作。

类型：PostRenderAction，枚举值读/写：

- PostRenderAction.NONE:无
- PostRenderAction.IMPORT：导入
- PostRenderAction.IMPORT_AND_REPLACE_USAGE：导入并且替换使用
- PostRenderAction.SET_PROXY：设置代理

### templates

app.project.renderQueue.item(index).outputModule(index).templates

描述：After Effects 的本地安装中可用的所有输出模块模板的名称。

类型：字符串数组；只读。

## 方法

### applyTemplate()

app.project.renderQueue.item(index).outputModule(index).applyTemplate(templateName)

描述：应用指定的现有输出模板。

参数：templateName ，字符串，应用的模板名称。

返回：无。

### getSetting()

app.project.renderQueue.item(index).outputModule(index).getSetting()

::: info 提示
此功能已添加到 After Effects 13.0(CC 2014)中
:::

描述：获取给定输出模块的特定设置。

折旧资源：https://blogs.adobe.com/creativecloud/new-changed-after-effects-cc-2014/?segment=dva

存档版本：https://web.archive.org/web/20200622100656/
https://blogs.adobe.com/creativecloud/new-changed-after-effects-cc-2014/?segment=dva

示例

有关结构参考，请参见 RenderQueueItem.getSetting()中的示例。

### getSettings()

app.project.renderQueue.item(index).outputModule(index).getSettings()

::: info 提示
此功能已添加到 After Effects 13.0(CC 2014)中
:::

描述：获取给定输出模块的所有设置。

折旧资源：https://blogs.adobe.com/creativecloud/new-changed-after-effects-cc-2014/?segment=dva

存档版本：https://web.archive.org/web/20200622100656/https://blogs.adobe.com/creativecloud/new-changed-after-effects-cc-2014/?segment=dva

示例

```javascript
// Get object that contains the string version of all current output module
setting
// values of output module item 1 from render queue item 1.
// To get the values in the number format, use GetSettingsFormat.NUMBER as an
argument.

var omItem1_all_str=app.project.renderQueue.item(1).outputModule(1).getSettings(GetSettingsFormat.STRING );

// Convert to JSON format so that it is human-readable.

var omItem1_all_str_json = omItem1_all_str.toSource();

// Get object that contains string version of settable output module setting values
// of output module item 1 from render queue item 1.
// If you want to get the values in the number format, use
// GetSettingsFormat.NUMBER_SETTABLE as an argument.

var omItem1_settable_str = app.project.renderQueue
.item(1)
.outputModule(1)
.getSettings(
GetSettingsFormat.STRING_SETTABLE
);

// Currently, the format setting in the output module is not settable, but it
// is readable. The next line will tell you the current format of output
module
// item 1 from render queue item 1.

var current_format = app.project.renderQueue.item(1).outputModule(1).getSettings(GetSettingsFormat.STRING).Format;

// This line will tell you the output module file info.

var current_omFileTempalte =
app.project.renderQueue.item(1).outputModule(1).getSettings(GetSettingsFormat.STRING)["Output File Info"]["file template"];
```

### remove()

app.project.renderQueue.item(index).outputModule(index).remove()

描述：从集合中删除此 OutputModule 对象。

参数：无。

返回：无。

### saveAsTemplate()

app.project.renderQueue.item(index).outputModule(index).saveAsTemplate(name)

描述：将此输出模块另存为模板，并将其添加到 template 数组。

参数：name，新模板名称，字符串。

返回：无。

### setSetting()

app.project.renderQueue.item(index).outputModule(index).setSetting()

::: info 提示
此功能已添加到 After Effects 13.0(CC 2014)中
:::

描述：为给定的输出模块设置特定设置。

折旧资源：https://blogs.adobe.com/creativecloud/new-changed-after-effects-cc-2014/?segment=dva

存档版本：https://web.archive.org/web/20200622100656/
https://blogs.adobe.com/creativecloud/new-changed-after-effects-cc-2014/?segment=dva

示例

有关结构参考，请参见 RenderQueueItem.setSetting()中的示例。

### setSettings()

app.project.renderQueue.item(index).outputModule(index).setSettings()

::: info 提示
此功能已添加到 After Effects 13.0(CC 2014)中
:::

描述：折旧资源：https://blogs.adobe.com/creativecloud/new-changed-after-effects-cc-2014/?segment=dva

存档版本：https://web.archive.org/web/20200622100656/
https://blogs.adobe.com/creativecloud/new-changed-after-effects-cc-2014/?segment=dva

:::danger
修改输出模块设置后，有一个错误会导致 OutputModule 对象失效，因此，修改它后，需要再次检索输出模块。
:::

示例

从一项的输出模块获取设置，然后在另一项上使用它们：

```javascript
// If you want to get the values in the number format, use
// GetSettingsFormat.NUMBER_SETTABLE as an argument.

var omItem1_settable_str =
app.project.renderQueue.item(1).outputModule(1).getSettings(
GetSettingsFormat.STRING_SETTABLE );

// Set output module item 1 of render queue item 2 with values that you get
from
// output module 1 of render queue item 1

app.project.renderQueue.item(2).outputModule(1).setSettings(
omItem1_settable_str );
使用您创建的值设置渲染队列项目 3 的输出模块项目 1：

var crop_data = {
"Crop": true,
"Crop Bottom": 0,
"Crop Left": 0,
"Crop Right": 8,
"Crop Top": 10
};

app.project.renderQueue.item(1).outputModule(3).setSettings( crop_data );
将输出文件路由到用户目录：

var om1 = app.project.renderQueue.item(1).outputModule(1);
var file_name = File.decode( om1.file.name ); // Name contains special
character, space?
var new_dir = new Folder( "~/new_output" );
var new_path = new_dir.fsName;

var new_data = {
"Output File Info": {
"Base Path": new_path,
"Subfolder Path": "draft",
"File Name": file_name
}
};

om1.setSettings( new_data );
在此示例中，输出文件被路由到用户目录，但是这次使用完整路径：

var om1 = app.project.renderQueue.item(1).outputModule(1);

// Name contains special character, such as space?
var file_name = File.decode( om1.file.name );
var new_path = "/Users/myAccount/new_output";
var separator = "/";

if ($.os.indexOf("Mac") == -1) {
new_path = "C:\Users\myAccount\new_output";
separator = "\\";
}

var new_data = {
"Output File Info": {
"Full Flat Path": new_path + separator + file_name
}
};

om1.setSettings( new_data );
```
