---
order: 1
title: AE-Object-Model
---

# AE 对象模型

## 说明

浏览本参考部分（按对象按字母顺序排列）时，可以参考下图，概述各种对象在层次结构中的位置以及它们与用户界面的对应关系。

### AE 脚本主要对象的层次结构图

![](https://mir.yuelili.com/2021/07/1ae9650c46f14bb1310fa3b81fcccf99.png)
![](https://mir.yuelili.com/2021/07/5c3d5e676644ade2a14106a01e0e2b0d.png)

请注意，File(文件)，Folder(文件夹)和 Socket(外部通讯)对象由 ExtendScript 定义，记录在《[JavaScript 工具指南》中](http://estk.aenhancers.com/)。ExtendScript 还定义了 ScriptUI 模块：窗口和用户界面控制对象，这些对象也可用于 After
Effects 脚本。这些也记录在《[JavaScript 工具指南](http://estk.aenhancers.com/)》中。脚本编制中的对象层次结构对应于用户界面（深灰色）的层次结构。

![After Effects用户界面](https://mir.yuelili.com/2021/07/3d4c924e25a07dda2e1f28769be5682c.png)

应用程序包含一个 Project(项目)面板，现在显示一个项目。项目中有合成，合成又包含图层。图层的来源可以是素材文件，占位符或自己创建，也可以在“项目”面板中列出。每个图层都包含属性，属性又可以包含标记和关键帧。渲染队列包含渲染队列项以及渲染设置和输出模块。这些均由脚本中的对象表示。

::: info 提示
在本指南中，“属性（properties）”称为“属性（attributes）”，避免与 After Effects 自己的属性（properties）混淆。
:::

### 对象汇总

下表按字母顺序列出了所有对象，并提供了每个对象的文档页面链接。

| 目的                   | 描述                                                                                        |
| ---------------------- | ------------------------------------------------------------------------------------------- |
| Global 函数            | 全局可用的函数，可以显示文本以进行脚本调试，并帮助在秒和帧之间转换时间值。                  |
| Application 对象       | 单个全局对象（按其名称（应用程序）可用）可用于访问 AE 应用程序的对象和应用程序设置。        |
| AVItem 对象            | 表示导入到 After Effects 中的音频/视频文件。                                                |
| AVLayer 对象           | 表示包含 AVItem 对象的那些图层（合成图层，素材图层，实体(solid)图层，文本图层和声音图层）。 |
| CameraLayer 对象       | 表示合成中的相机图层。                                                                      |
| collection 收集对象    | 将一组对象或值关联为一个逻辑组，并通过索引提供对它们的访问。                                |
| CompItem 对象          | 表示一个合成，允许操纵并获取相关信息。                                                      |
| FileSource 对象        | 表示来自文件的素材。                                                                        |
| FolderItem 对象        | 在“项目”面板中表示一个文件夹。                                                              |
| FootageItem 对象       | 表示导入到项目中的素材项目，该项目显示在“项目”面板中。                                      |
| FootageSource 对象     | 表示一些素材的文件来源。                                                                    |
| ImportOptions 对象     | 封装用于将文件导入 After Effects 的选项。                                                   |
| Item 对象              | 表示出现在“项目”面板中项目的子项目。                                                        |
| ItemCollection 对象    | 收集项目中的子项目。                                                                        |
| KeyframeEase 对象      | 将关键帧缓动值封装在 After Effects 属性中。                                                 |
| Layer 对象             | 图层基类。                                                                                  |
| LayerCollection 对象   | 收集项目中的图层。                                                                          |
| LightLayer 对象        | 表示合成中的等光图层。                                                                      |
| MarkerValue 对象       | 将标记值封装在 After Effects 属性中。                                                       |
| MaskPropertyGroup 对象 | 将蒙版属性封装在图层中。                                                                    |
| OMCollection 对象      | 将输出模块收集在渲染队列中。                                                                |
| OutputModule 对象      | 表示渲染队列的输出模块。                                                                    |
| PlaceholderSource 对象 | 表示素材的占位符。                                                                          |
| Project 对象           | 表示一个 After Effects 项目。                                                               |
| Property 对象          | 表示 After Effects 属性。                                                                   |
| PropertyBase 对象      | After Effects 属性和属性组的基类。                                                          |
| PropertyGroup 对象     | 表示 After Effects 属性组。                                                                 |
| RenderQueue 对象       | 表示 After Effects 渲染队列。                                                               |
| RenderQueueItem 对象   | 表示渲染队列中的可渲染项。                                                                  |
| RenderQueueItem 对象   | 收集渲染队列中的渲染队列项。                                                                |
| RQItemCollection 对象  | 用于访问应用程序设置和首选项。                                                              |
| Shape 对象             | 封装蒙版的轮廓形状信息。                                                                    |
| ShapeLayer 对象        | 表示合成中的形状图层。                                                                      |
| SolidSource 对象       | 描述纯色是某些素材的来源。                                                                  |
| System 对象            | 提供从应用程序访问操作系统的权限。                                                          |
| TextDocument 对象      | 将文本封装在文本图层中。                                                                    |
| TextLayer 对象         | 表示合成中的文本图层。                                                                      |
| Viewer 对象            | 表示“合成”，“层”或“素材”面板。                                                              |
