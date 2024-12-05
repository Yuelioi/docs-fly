---
title: Sample-Projects
order: 13
---

# 示例工程

目前的 API 所支持的每一种类型的插件都有至少一个样本，还有一些工程用来说明特定的概念。

在示例工程中，我们尽可能地保持代码简单。高大尚的功能可能让代码看起来很牛 X，但不能帮你理解如何使用 API 的功能。

现在，我们来解释如何构建这些示例工程，请继续阅读下文

## 示例工程描述

| 工程                   | 描述                                                                                                                                                                                                                                                                                                                  |
| ---------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| AEGPs                  | AEGPs 绑定 AE 菜单和其他 UI 区域. 详情请看 AEGP 在 UI 的表现.                                                                                                                                                                                                                                                         |
| Artie                  | Artisan 处理给定合成的 3D 图层渲染, 这与我们的内部 3D 渲染器所使用的 API 相同; 它非常复杂，而且暴露了大量关于 After Effects 渲染管道的隐性信息。除非你有一个令人信服的理由来替换 After Effects 处理 3D 渲染的方式，否则永远不需要使用这个样本。 Artisan 出现在合成>合成设置中，在进阶选项卡中，渲染插件的下拉菜单中。 |
| Easy Cheese            | A keyframer (动画 > 关键帧助手子菜单), Easy Cheese 显示了如何操作关键帧的各种特性（其方式与我们的 shipping 插件相似，Easy Ease）。                                                                                                                                                                                    |
| FBIO                   | After Effects Input/Output (AEIO) API. 与 IO 示例类似, 但是支持 frame-based .ffk file 格式. 请注意，我们现在建议使用 [Premiere Pro Importers](../intro/other-integration-possibilities.html)                                                                                                                          |
| Grabba                 | 获取工程中任何合成的帧 (格式化为插件请求)                                                                                                                                                                                                                                                                             |
| IO                     | After Effects Input/Output (AEIO) API. 支持 fictitious.fak 文件格式, 并处理来自 After Effects 的所有请求，以从此类文件中检索数据或输出到此类文件。请注意，我们现在建议使用[Premiere Pro Importers](../intro/other-integration-possibilities.html)                                                                 |
| Mangler                | 关键帧器，演示使用 ADM 调色板，就像我们自己的一样。                                                                                                                                                                                                                                                 |
| Panelator              | 可以与其他标准面板一起停靠的面板。注意：以这种方式创建面板比使用 HTML5 面板 SDK 要多得多。我们建议从该 SDK 开始。                                                                                                                                                                                             |
| Persisto               | 从 After Effects 首选项文件读取和写入信息。                                                                                                                                                                                                                                                                   |
| ProjDumper             | 建一个文本文件，表示 After Effects 项目中的每个元素。                                                                                                                                                                                                                                                                 |
| Projector              | 导入（fictitious）. sdk 文件格式，并使用 AEGP API 调用创建项目。每当您想知道如何获取或设置项目元素的某些特征时，请先看这里。注意：Projector. h 中有一些硬编码路径。如果您不将这些设置为引用磁盘上的实际媒体，运行此插件时将会收到错误。不要责怪我们；改变它们！                                                       |
| QueueBert              | 发音为“Cue-BARE！”，QueueBert 操作渲染队列项的所有方面以及与之关联的输出模块。                                                                                                                                                                                                                                        |
| Streamie               | 操纵流，包括动态流和固定流。                                                                                                                                                                                                                                                                                          |
| Sweetie                | 使用 PICA（或“Suite Pea”）API 来提供函数套件，供其他插件使用。如果您正在编写依赖于同一图像处理库的多个插件，您可以使用这样的套件提供库功能。                                                                                                                                                                  |
| Text Twiddler          | 操作文本层及其内容。                                                                                                                                                                                                                                                                                                  |
| Effects                | 所有效果都出现在效果和预设面板以及效果菜单中。                                                                                                                                                                                                                                                                        |
| Checkout               | 在指定时间从另一个层签出（在 After Effects 的帧缓存中）输入帧。这是具有层参数的所有效果的重要概念。Premiere Pro 兼容。                                                                                                                                                                                                |
| Convolutrix            | 图像卷积回调。Premiere Pro 兼容。                                                                                                                                                                                                                                                                           |
| Gamma Table            | 管理序列数据，并使用我们的迭代回调。出于怀旧的考虑，我们将此示例保留在 C 中；它还与许多第三方插件主机兼容，因为它依赖于 3.x 版 API 功能。                                                                                                                                                                     |
| GLator                 | CC 2017 的新功能。在效果插件中演示正确的 OpenGL 上下文管理。                                                                                                                                                                                                                                                          |
| Paramarama             | 其他示例中未使用的任性参数类型。与 Premiere Pro 兼容。                                                                                                                                                                                                                                                            |
| PathMaster             | 从效果中访问路径。                                                                                                                                                                                                                                                                                            |
| Portable               | 检测和响应几个不同的插件主机。Premiere Pro 兼容。                                                                                                                                                                                                                                                               |
| Resizer                | 调整输出缓冲区的大小（令人惊讶！）。这对于发光和阴影等效果很有用，如果它们不扩展输出缓冲区，这些效果将在图层边缘被截断。Premiere Pro 兼容。                                                                                                                                                                   |
| SDK Backwards          | 反转图层的音频，并将其与关键帧的正弦波混合。                                                                                                                                                                                                                                                                          |
| SDK Noise              | Premiere Pro 兼容，演示 Premiere Pro 中的 32 位和 YUV 渲染。                                                                                                                                                                                                                                                          |
| Shifter                | 在输出缓冲区中移动图像，并探究我们的 transform_world 和子像素采样功能。                                                                                                                                                                                                                                               |
| SmartyPants            | 支持浮点像素所需的 SmartFX API。                                                                                                                                                                                                                                                                                  |
| Transformer            | 探究图像变换回调                                                                                                                                                                                                                                                                                                      |
| Effect Template        |                                                                                                                                                                                                                                                                                                                       |
| Skeleton               | Skeleton 是开发效果的起点。Premiere Pro 兼容。                                                                                                                                                                                                                                                                        |
| Effects with Custom UI |                                                                                                                                                                                                                                                                                                                       |
| CCU                    | 在合成和图层窗口中实现自定义用户交互界面，支持宽高像素比和降频采样比。Premiere Pro 兼容。                                                                                                                                                                                                                               |
| ColorGrid              | 使用任意数据类型参数。还有一个不错的自定义UI。Premiere Pro兼容                                                                                                                                                                                                                  |
| Custom ECW UI          | 演示如何使用任意数据类型参数。还有一个漂亮的自定义 UI。与 Premiere Pro 兼容。                                                                                                                                                                                                                                         |
| Histogrid              | CC 2015（13.5）的新功能。自定义 UI 如何访问异步渲染的上游帧以在 CC 2015 及更高版本中进行轻量级处理的示例。此效果从上游帧计算采样的 10x10 颜色网格，并显示该颜色格的预览。在渲染中，计算更高质量的网格并用于修改输出图像，创建颜色网格与原始图像的混合。                                                               |
| Supervisor             | 显示如何根据其他参数的值控制参数（值和 UI）。Premiere Pro 兼容。                                                                                                                                                                                                                                                      |
| BlitHook               |                                                                                                                                                                                                                                                                                                                       |
| EMP                    | 外部监视器预览。将此用作将支持从合成面板输出视频添加到视频硬件的起点。                                                                                                                                                                                                                                                |

## 编译示例项目

我们将样本项目合并为一个主项目，存放在 SDK 的 Examples 文件夹中。macOS 的 Buildall.xcodeproj； Windows 的 BuildAll.sln。

在 IDE 中，需要改变项目的输出文件夹，以构建到 After Effects 的插件文件夹。

对于开发来说，我们建议 macOS 使用以下路径：`/Library/Application Support/Adobe/Common/Plug-ins/[版本]/MediaCore/`。

所有 CC 版本的版本都锁定在 7.0，早期版本锁定在 CSx。

例如：`/Library/Application Support/Adobe/Common/Plug-ins/7.0/MediaCore/`。

或。`/Library/Application Support/Adobe/Common/Plug-ins/CS6/MediaCore/`。

和 Windows 的：`[Program Files]\Adobe\Common\Plug-ins\[version]\MediaCore\`。

例如：`C:Program Files\Adobe\Common\Plug-ins\7.0\MediaCore\`。

或者。`C:\Program Files\CommonAdobe\Plug-ins\CS6\MediaCore\`。

请注意，这个 Windows 路径只推荐用于开发目的。Windows 的安装应该遵循这里的指南。[Where Installers Should Put Plug-ins](Where-installers-should-put-plug-ins.html) 。

在 Xcode 中，你可以在 Xcode 首选项>位置>派生数据>高级 中为所有项目设置一次这个路径。在*Build Location*下选择*Custom*，并填写路径。

在 Visual Studio 中，为了方便起见，我们使用环境变量 AE *PLUGIN_BUILD_DIR* 来指定所有样本项目的输出路径。你需要把它设置为你系统的用户环境变量。在 Windows 7 上，右键单击 我的电脑 > 高级计算机属性。

在新的对话框中，点击 *环境变量* 。在用户变量区，创建一个名为 AE_PLUGIN_BUILD_DIR 的新变量，其路径如上所述。退出 Windows 并重新登录，这样该变量就会被设置。

另外，可以在 Visual Studio 中为每个项目单独设置输出路径，在解决方案资源管理器中右击项目，选择属性，然后在配置属性 > 链接器 > 常规中，设置输出文件。

在编译插件时，如果你看到一个链接错误，比如 "无法打开文件"[MediaCore 插件路径]plugin.prm"，确保以管理员模式启动 Visual Studio。在 Visual Studio 安装中，右键单击 devenv.exe，属性 > 兼容性 > 权限级别，点击 "以管理员身份运行此程序"。
