---
title: AEGP-Suites
order: 2
---

# AEGP 套件

如前所述，AEGPs 通过套件来完成一切。以下套装被所有类型的 AEGP 使用，并且可以从任何钩子函数中调用(除了 RegisterSuite，它必须从 AEGP 的入口函数中使用)。下面是每个套件中每个函数的描述，以及在适当的地方使用这些函数的细节。

| **Suite**| **Description** |
| --- | --- |
| Memory Suite| 管理内存资源。使用此套件！每当遇到与内存相关的错误时，AE都可以为您报告错误。 |
| Command Suite | 管理您的AEGP菜单项。与注册套件结合使用。 |
| Register Suite  | 与Command Suite连用,用于菜单命令添加功能。AEIO和工匠必须使用此套件的功能向AE指示他们想要接收适当的消息流。您可以使用此套件替换一些AE命令。 |
| Project Suite | 读取和修改项目数据。 |
| Item Suite | 管理项目或合成中的项目。文件夹、合成、纯色层和素材都是项目。  |
| Collection Suite | 查询当前选择了哪些项目，并创建您自己的选择集。选择您的AEGP修改的所有项目通常是一个很好的UI举措，只是为了让用户知道您做了什么。|
| Composition Suite  |项目中管理（和创建）合成，以及合成特定的项目，如纯色层。  |
| Footage Suite | 管理素材|
| Layer Suite  | 提供有关合成中图层的信息，以及源和图层时间之间的关系。纯色层、文本、绘制、相机、灯光、图像和图像序列都可以成为图层。|
| Effect Suite| 提供对应用于图层的效果的访问。使用流套件获取效果关键帧信息。使用`AEGP_EffectCallGeneric()`(in AEGP_EffectSuite4) 与您提前设置的效果进行通信，以响应您的AEGP。 |  | Stream Suite | 用于访问图层的关键帧属性的值。  |  | Dynamic Stream Suite | 用于访问与图层相关联的动态流的特性。 |  | Keyframe Suite | 用于访问和操作所有关键帧数据。 |  | Marker Suite | 用于操纵标记。使用 `AEGP_GetNewCompMarkerStream()` (in AEGP_CompSuite11) 来获取合成的标记流|
| Mask Suite | 提供检索有关图层蒙版的信息的访问。 |
| Mask Outline Suite| 与Stream Suite结合使用时，该套件提供了有关为制作图层蒙版而渲染的路径的详细信息。|
| Text Document Suite |用于访问文本图层上的实际文本。 |
| Text Layer Suite | 用于访问构成文本图层轮廓的路径。|
| Utility Suite | 提供错误消息处理、AEGP版本检查和对AE撤消堆栈的访问。 |
| Persistent Data Suite| 查询和管理所有持久(persistent)数据（即首选项文件）。AEGP还可以将自己的数据添加到首选项中。  |
| Color Settings Suite  | 获取有关AE当前颜色管理设置的信息。  |
| Render Suite| 从AEGP中获取渲染的帧（和音频样本）。  |
| World Suite  | 分配、处置和查询AEGP_Worlds。还提供了一种将`PF_EffectWorld`转换为`AEGP_World`的方法，用于使用效果插件。   |
| Composite Suite| 公开AE的合成功能，包括混合模式、轨道遮罩和良好的老式位复制(fashioned bit copying)。  |
| Sound Data Suite |用于管理和访问声音数据的功能。 |
| Render Queue Suite| 从渲染队列中添加和删除项目。|
| Render Queue Item Suite  | 查询和修改渲染队列中的项目。|
| Render Options Suite  |  查询和管理渲染队列项的选项对话框中公开的所有项目。|
| Output Module Suite | 查询和修改附加到渲染队列中项目的输出模块。  |
| PF Interface Suite| 该套件中的功能虽然在技术上是AEGP API的一部分，但供效果使用。   |
| AEGP Iterate Suite| 为AEGP提供了一种在任何或所有可用处理器上运行函数（具有所需签名）的方法。 |
| File Import Manager Suite | 注册AEGP文件和项目导入器作为AE文件处理的一部分。|

## 示例代码

在skeleton.cpp文件里 有几个入口函数,比如`About`,`GlobalSetup`,如果没有其他说明,可以把我写的示例放在这些入口函数里进行测试

比如以下代码代表,我在`Render`入口函数里进行测试,`...` 代表省略了函数本身的代码

```cpp
static PF_Err 
Render (

// ...

AEGP_Command cmd;
ERR(suites.CommandSuite1()->AEGP_InsertMenuCommand(
  suites.CommandSuite1()->AEGP_GetUniqueCommand(&cmd),
  "Test",
  AEGP_Menu_WINDOW,
  AEGP_MENU_INSERT_SORTED));

// ...

```

## 优雅地出错

如果一个套件不存在，要尽力做到优雅地失败。向用户显示一条信息，说明问题的性质。尝试使用同一套件的早期版本。

由于 AEGPs 与 After Effects 结合得如此之深，要确保用户知道是谁或什么东西遇到了某个问题。

表明自己的身份! 尽可能的向用户提供支持和/或帮助信息。

## 处理 Handles

使用 AEGP 内存套件来管理 AEGP 使用的内存。每当遇到与内存有关的错误，After Effects 可以报告错误，让你及早发现。

`AEGP_MemHandle`是一个结构，它包含的不仅仅是被引用的内存。所以它不应该被直接解引用。使用`AEGP_LockMemHandle`来获得`AEGP_MemHandle`所引用的内存的指针。

当然，当你完成后要解锁它。

### AEGP_MemorySuite1

#### AEGP_NewMemHandle

创建一个新的内存句柄。`plugin_id`是通过[入口函数](implement.html#aegps-implementation-entry-point)传入的 ID，或者是你从`AEGP_RegisterWithAEGP()`获得的。
(来自[AEGP_UtilitySuite6])。
使用`whatZ`来识别你所要求的内存。After Effects 使用这个字符串来显示任何相关的错误信息。

```cpp
AEGP_NewMemHandle(
 AEGP_PluginID *plugin_id,
 const A_char *whatZ,
 AEGP_MemSize size,
 AEGP_MemFlag flags,
 AEGP_MemHandle *memPH);
```

#### AEGP_FreeMemHandle

释放你用 AEGP_NewMemHandle()分配的一个句柄。

```cpp
AEGP_FreeMemHandle(
 AEGP_MemHandle memH);
```

#### AEGP_LockMemHandle

在使用`AEGP_NewMemHandle`分配的内存之前，使用这个函数来锁定手柄(不能被操作系统移动)。可以被嵌套。

```cpp
AEGP_LockMemHandle(
 AEGP_MemHandle memH,
 void **ptr_to_ptr);
```

#### AEGP_UnlockMemHandle

允许操作系统移动引用的内存。始终平衡锁调用和解锁。

```cpp
AEGP_UnlockMemHandle(
 AEGP_MemHandle memH);
```

#### AEGP_GetMemHandleSize

返回句柄的分配大小。

```cpp
AEGP_GetMemHandleSize AEGP_MemHandle memH,
 AEGP_MemSize *sizeP);
```

#### AEGP_ResizeMemHandle

改变句柄的分配大小。

```cpp
AEGP_ResizeMemHandle(
 const char *whatZ,
 AEGP_MemSize new_size,
 AEGP_MemHandle memH);
```

#### AEGP_SetMemReportingOn

如果 After Effects 在内存处理方面遇到问题，应该向用户报告错误。在开发过程中要利用这一点!
只有使用该套件分配并泄漏的内存才会被报告，所以例如使用[PF_HandleSuite1](../effect-details/memory-allocation.html#effect-details-memory-allocation-pf-handlesuite)分配的内存不会被报告。
将不会被报告。

```cpp
AEGP_SetMemReportingOn(
 A_Boolean turn_OnB);
```

#### AEGP_GetMemStats

获取关于当前分配的句柄数量及其总大小的信息。
只有使用这个套件分配的内存才会被跟踪，并使用这个调用进行报告，所以例如使用[PF_HandleSuite1](../effect-details/memory-allocation.html#effect-details-memory-allocation-pf-handlesuite)分配的内存不会在这里报告。

```cpp
AEGP_GetMemStats(
 AEGP_MemID mem_id,
 A_long *countPL,
 A_long *sizePL);
```

## Managing Menu Items

命令套件允许你创建和处理任何菜单事件。

要添加你自己的菜单命令，你还必须使用[注册套件]来为菜单事件分配处理程序。

### AEGP_CommandSuite1

#### AEGP_GetUniqueCommand

获得一个唯一的命令标识符。使用*注册套件*为该命令注册一个处理程序。

```cpp
AEGP_GetUniqueCommand(
 AEGP_Command *unique_commandP);
```

注意：After Effects 有时会发送命令 0(零)。
所以不要把它作为你的命令处理逻辑的一部分。

#### AEGP_InsertMenuCommand

添加一个新的菜单命令。使用 nameZ = "-" 将插入一个分隔符。 menu_ID 可以是:

- `AEGP_Menu_NONE`
- `AEGP_Menu_APPLE`
- `AEGP_Menu_FILE`
- `AEGP_Menu_EDIT`
- `AEGP_Menu_COMPOSITION`
- `AEGP_Menu_LAYER`
- `AEGP_Menu_EFFECT`
- `AEGP_Menu_WINDOW`
- `AEGP_Menu_FLOATERS`
- `AEGP_Menu_KF_ASSIST`
- `AEGP_Menu_IMPORT`
- `AEGP_Menu_SAVE_FRAME_AS`
- `AEGP_Menu_PREFS`
- `AEGP_Menu_EXPORT`
- `AEGP_Menu_ANIMATION`
- `AEGP_Menu_PURGE`
- `AEGP_Menu_NEW` (CC 之后才支持)

位置可以设置为菜单中的一个特定位置，也可以是 After Effects 指定的一个位置。

- `AEGP_MENU_INSERT_SORTED`
- `AEGP_MENU_INSERT_AT_BOTTOM`
- `AEGP_MENU_INSERT_AT_TOP`

对于`AEGP_Menu_WINDOW`，从 CS4 开始就不支持 BOTTOM 和 TOP 选项，会返回一个错误。
我们推荐`SORTED`。

```cpp
AEGP_InsertMenuCommand(
 AEGP_Command command,
 const A_char *nameZ,
 AEGP_MenuID menu_id,
 A_long after_itemL);
```

示例:注册一个名为"test"的菜单(在"窗口"菜单栏)

:::warn
菜单名称不支持中文
:::

```cpp

// GlobalSetup ...
PF_Err    err = PF_Err_NONE;
AEGP_SuiteHandler suites(in_data->pica_basicP);
AEGP_Command cmd; // 定义一个命令

ERR(suites.CommandSuite1()->AEGP_InsertMenuCommand(
  suites.CommandSuite1()->AEGP_GetUniqueCommand(&cmd), // 获取唯一ID
  "Test", // 非中文名称
  AEGP_Menu_WINDOW, // 窗口菜单
  AEGP_MENU_INSERT_SORTED));
```

#### AEGP_RemoveMenuCommand

删除一个菜单命令。如果你有这么大的动力，你可以删除所有 After Effects 的菜单项。

```cpp
AEGP_RemoveMenuCommand(
 AEGP_Command command);
```

#### AEGP_SetMenuCommandName

设置一个命令的菜单名称。

```cpp
AEGP_SetMenuCommandName(
 AEGP_Command command,
 const A_char *nameZ);
```

#### AEGP_EnableCommand

启用一个菜单命令。

```cpp
AEGP_EnableCommand(
 AEGP_Command command);
```

#### AEGP_DisableCommand

禁用一个菜单命令。

```cpp
AEGP_DisableCommand(
 AEGP_Command command);
```

#### AEGP_CheckMarkMenuCommand

After Effects 会在菜单命令的旁边画一个复选标记。

```cpp
AEGP_CheckMarkMenuCommand(
 AEGP_Command command,
 A_Boolean checkB);
```

#### AEGP_DoCommand

调用指定菜单命令的处理程序。每个 After Effects 菜单项都有一个相关的命令。
注意，我们不保证各版本的命令 ID 是一致的。

```cpp
AEGP_DoCommand(
 AEGP_Command command);
```

在给出了上述免责声明后，这里有几个已经提供给其他开发者的命令号，可能会有兴趣。

- 3061 - 打开选择，忽略任何修改键。
- 10314 - 播放/停止(在 13.5 及以后版本中有效)
- 2285 - RAM 预览(在 13.5 之前有效)
- 2415 - 播放(空格键)(在 13.5 之前有效)。
- 2997 - 裁剪合成到目标区域。
- 2372 - 编辑 > 清除 > 图像缓存

如果你的 AEGP 需要调用一些其他的 After Effects 菜单项，有一个相当简单的方法来找出你想要的大多数命令，使用脚本。

```cpp
cmd = app.findMenuCommandId(text); // e.g. text = "Open Project…"
alert(cmd);
```

在 AE 运行的情况下，只需打开 Adobe ExtendScript Toolkit CC，将上述脚本复制进去，并在应用程序下拉菜单中选择你正在运行的 After Effects 版本，然后点击播放按钮在 AE 中运行脚本。
。

## 在 After Effects 中注册

注册函数供 After Effects 使用。

### AEGP_RegisterSuite5

#### AEGP_RegisterCommandHook

在 After Effects 中注册一个钩子(命令处理)函数。如果要替换 After Effects 也要处理的函数，`AEGP_HookPriority`决定是插件先处理还是AE先处理。

- `AEGP_HP_BeforeAE`
- `AEGP_HP_AfterAE`

对于添加的每个菜单项，需要使用`AEGP_GetUniqueCommand()`预先注册单个`command_hook_func`获得你自己的`AEGP_Command`。判断哪个命令是在这个钩子函数中发送的，并采取相应的行动。

目前，`AEGP_HookPriority`被忽略了。

```cpp
AEGP_RegisterCommandHook(
 AEGP_PluginID aegp_plugin_id,
 AEGP_HookPriority hook_priority,
 AEGP_Command command,
 AEGP_CommandHook command_hook_func
 void *refconPV);
```

#### AEGP_RegisterUpdateMenuHook

注册你的菜单更新函数(决定项目是否处于活动状态)，在每次绘制任何菜单时调用。
这个钩子函数处理所有菜单的更新。

```cpp
AEGP_RegisterUpdateMenuHook(
 AEGP_PluginID aegp_plugin_id,
 AEGP_UpdateMenuHook update_menu_hook_func,
 void *refconPV);
```

#### AEGP_RegisterDeathHook

注册你的终止函数。当应用程序退出时被调用。

```cpp
AEGP_RegisterDeathHook(
 AEGP_PluginID aegp_plugin_id,
 AEGP_DeathHook death_hook_func,
 void *refconPV);
```

#### AEGP_RegisterVersionHook

目前不被调用。

#### AEGP_RegisterAboutStringHook

目前未被调用。

#### AEGP_RegisterAboutHook

目前未被调用。

#### AEGP_RegisterArtisan

注册你的 Artisan。参见 [Artisans](../artisans/artisans.html#artisans-artisans) 以了解更多细节。

```cpp
AEGP_RegisterArtisan(
 A_Version api_version,
 A_Version Artisan_version,
 long aegp_plugin_id,
 void *aegp_refconPV,
 const A_char *match_nameZ,
 const A_char *Artisan_nameZ,
 PR_ArtisanEntryPoints *entry_funcsP);
```

#### AEGP_RegisterIO

注册你的 AEIO 插件。参见[AEIOs](../aeios/aeios.html#aeios-aeios)以了解更多细节。

```cpp
AEGP_RegisterIO (
 AEGP_PluginID aegp_plugin_id,
 AEGP_IORefcon aegp_refconP,
 const AEIO_ModuleInfo *io_infoP,
 const AEIO_FunctionBlock4 *aeio_fcn_blockP);
```

#### AEGP_RegisterIdleHook

注册你的 IdleHook 函数。After Effects 会在用户做出困难的艺术决定时(或在他们得到更多的咖啡时)，零星地调用这个函数。

```cpp
AEGP_RegisterIdleHook(
 AEGP_PluginID aegp_plugin_id,
 AEGP_IdleHook idle_hook_func,
 AEGP_IdleRefcon refconP);
```

#### AEGP_RegisterInteractiveArtisan

将你的 AEGP 注册为一个交互式工匠(Artisans)，用于预览和渲染特定合成中的所有层。

```cpp
AEGP_RegisterInteractiveArtisan (
 A_Version api_version,
 A_Version artisan_version,
 AEGP_PluginID aegp_plugin_id,
 void *aegp_refconPV,
 const A_char *match_nameZ,
 const A_char *artisan_nameZ,
 PR_ArtisanEntryPoints *entry_funcsP);
```

#### AEGP_RegisterPresetLocalizationString

调用它来注册尽可能多的字符串，以便在加载预置时进行名称替换。
任何时候发现一个属性名称，或者在表达式中提到它，并且它以 ASCII tab 字符('t')开头，后面是一个英文名称，它将被替换成本地化的名称。
(在英语中，制表符将被简单地删除)。

```cpp
AEGP_RegisterPresetLocalizationString(
 const A_char *english_nameZ,
 const A_char *localized_nameZ);
```

## Manage Projects

这些功能访问和修改项目数据。包括对多个项目的支持，以便为将来的扩展做准备。
目前 After Effects 坚持采用单一项目模式。

要在 After Effects 的首选项中保存特定项目的数据(因此，在项目本身之外)，请使用"持久性数据套件"

请注意：用于打开和创建项目的函数在被调用时不会保存对当前打开的项目的修改!

### AEGP_ProjSuite6

#### AEGP_NumProjects

目前不会返回超过 1，After Effects 一次只能打开一个项目。

```cpp
AEGP_GetNumProjects(
 A_long *num_projPL)
```

#### AEGP_GetIndProject

通过索引检索一个特定的项目。

```cpp
AEGP_GetProjectProjectByIndex(
 A_long proj_indexL,
 AEGP_ProjectH *projPH);
```

#### AEGP_GetProjectName

获取项目名称(最多到`AEGP_MAX_PROJ_NAME_LEN + 1`)的长度。

```cpp
AEGP_GetProjectName(
 AEGP_ProjectH projH,
 A_char *nameZ);
```

#### AEGP_GetProjectPath

获取项目的路径(空字符串表示项目还没有被保存)。
路径是一个以 NULL 结尾的 A_UTF16Char 字符串的句柄，必须用`AEGP_FreeMemHandle`处置。

```cpp
AEGP_GetProjectPath(
 AEGP_ProjectH projH,
 AEGP_MemHandle *unicode_pathPH)
```

#### AEGP_GetProjectRootFolder

获取项目的根，After Effects 也将其视为一个文件夹。

```cpp
AEGP_GetProjectRootFolder(
 AEGP_ProjectH projH,
 AEGP_ItemH *root_folderPH)
```

#### AEGP_SaveProjectToPath

将整个项目保存到指定的完整路径。
文件路径是一个以 NULL 结尾的 UTF-16 字符串，带有平台分隔符。

```cpp
AEGP_SaveProjectToPath(
 AEGP_ProjectH projH,
 const A_UTF16Char *pathZ);
```

#### AEGP_GetProjectTimeDisplay

检索当前的时间显示设置。

```cpp
AEGP_GetProjectTimeDisplay(
 AEGP_ProjectH projH,
 AEGP_TimeDisplay3 *time_displayP);

typedef struct {
 AEGP_TimeDisplayMode display_mode;
 AEGP_SourceTimecodeDisplayMode footage_display_mode;
 A_Boolean display_dropframeB;
 A_Boolean use_feet_framesB;
 A_char timebaseC;
 A_char frames_per_footC;
 AEGP_FramesDisplayMode frames_display_mode;
} AEGP_TimeDisplay3;

enum {
 AEGP_TimeDisplay_TIMECODE = 0,
 AEGP_TimeDisplay_FRAMES
};

typedef char AEGP_TimeDisplayMode;

enum {
 AEGP_SourceTimecode_ZERO= 0,
 AEGP_SourceTimecode_SOURCE_TIMECODE
};

typedef char AEGP_SourceTimecodeDisplayMode;

enum {
 AEGP_Frames_ZERO_BASED= 0,
 AEGP_Frames_ONE_BASED,
 AEGP_Frames_TIMECODE_CONVERSION
};

typedef char AEGP_FramesDisplayMode;
```

#### AEGP_SetProjectTimeDisplay

指定用于显示时间的设置。

```cpp
AEGP_SetProjectTimeDisplay(
 AEGP_ProjectH projH,
 const AEGP_TimeDisplay3 *time_displayP);
```

#### AEGP_ProjectIsDirty

如果项目在打开后被修改过，返回 TRUE。

```cpp
AEGP_ProjectIsDirty(
 AEGP_ProjectH projH,
 A_Boolean *is_dirtyPB);
```

#### AEGP_SaveProjectAs

将项目保存到指定路径。
文件路径是一个以 NULL 结尾的 UTF-16 字符串，带有平台分隔符。
注意：这将覆盖一个现有的文件。

```cpp
AEGP_SaveProjectAs(
 AEGP_ProjectH projH,
 const A_UTF16Char *pathZ);
```

#### AEGP_NewProject

创建一个新的项目。注意：将关闭当前的项目，而不先保存它。

```cpp
AEGP_NewProject(
 AEGP_ProjectH *new_projectPH);
```

#### AEGP_OpenProjectFromPath

从提供的路径打开一个项目，并返回其`AEGP_ProjectH`。
文件路径是一个以 NULL 结尾的 UTF-16 字符串，带有平台分隔符。
注意：将关闭当前项目而不先保存它

```cpp
AEGP_OpenProjectFromPath(
 const A_UTF16Char *pathZ,
 AEGP_ProjectH *projectPH);
```

#### AEGP_GetProjectBitDepth

检索项目的比特深度。

```cpp
AEGP_GetProjectBitDepth(
 AEGP_Projec tH projectH,
 AEGP_ProjBitDepth *bit_depthP);
```

AEGP_ProjBitDepth 将是以下之一。

- `AEGP_ProjBitDepth_8`
- `AEGP_ProjBitDepth_16`
- `AEGP_ProjBitDepth_32`

#### AEGP_SetProjectBitDepth

设置项目的比特深度。可撤销。

```cpp
AEGP_SetProjectBitDepth(
 AEGP_ProjectH projectH,
 AEGP_ProjBitDepth bit_depth);
```

### AEGP_TimeDisplay2

::: tip

当 After Effects 使用不同的显示类型时，未使用的字段中的值会持续存在。
:::

#### AEGP_TimeDisplayType type

以下之一。

- `AEGP_TimeDisplayType_TIMECODE`
- `AEGP_TimeDisplayType_FRAMES`
- `AEGP_TimeDisplayType_FEET_AND_FRAMES`

#### A_char timebaseC

0 - 100. 只用于`AEGP_TimeDisplayType_TIMECODE`。

#### A_Boolean non_drop_30B

当时间基数为 30，项目的帧率为 29.97 时，决定是否显示为非降帧。

#### A_char frames_per_footC

只用于`AEGP_TimeDisplayType_FEET_AND_FRAMES`。

#### A_long starting_frameL

通常为 0 或 1。当类型通常为 0 或 1 时不使用，不用于`AEGP_TimeDisplayType_TIMECODE`。

#### A_Boolean auto_timecode_baseB

如果`TRUE`，项目的时间码显示设置为自动。

## Control Items Within Projects

访问和修改项目或合成中的项目。

项目库中的任何东西都是`AEGP_Item`。注意，摄像机没有源，因此没有`AEGP_ItemH`。

除非你使用的函数需要更多的细节，否则尽可能保持抽象；AEGP_Comps 作为 AEGP_Items 被传入大多数函数并从其返回。

### AEGP_ItemSuite9

#### AEGP_GetFirstProjItem

检索给定项目中的第一个项目。

```cpp
AEGP_GetFirstProjItem(
 AEGP_ProjectH projectH,
 AEGP_ItemH *itemPH);
```

#### AEGP_GetNextProjItem

检索下一个项目；`*next_itemPH`在最后一个项目之后将是`NULL`。

```cpp
AEGP_GetNextProjItem(
 AEGP_ProjectH projectH,
 AEGP_ItemH itemH,
 AEGP_ItemH *next_itemPH);
```

#### AEGP_GetActiveItem

如果项目窗口处于活动状态，活动项目是选定的项目(如果只有一个项目被选中)。如果合成、时间轴或镜头窗口处于活动状态，返回与窗口中最前面的标签相关的层的父级。
如果没有项目处于活动状态，则返回 NULL。

```cpp
AEGP_GetActiveItem(
 AEGP_ItemH *itemPH,
```

#### AEGP_IsItemSelected

如果项目窗口处于活动状态并且项目被选中，则返回 true。

```cpp
AEGP_IsItemSelected(
 AEGP_ItemH itemH,
 A_Boolean *selectedPB)
```

#### AEGP_SelectItem

切换项目的选择状态，(取决于`deselect_othersB`)可以取消对其他项目的选择。这个调用可以选择项目面板中的项目。
要在合成面板中进行选择，请使用[AEGP_CompSuite11]的`AEGP_SetSelection`。

```cpp
AEGP_SelectItem(
 AEGP_ItemH itemH,
 A_Boolean selectB,
 A_Boolean deselect_othersB);
```

#### AEGP_GetItemType

获取项目的类型。注意：实体不会出现在项目中，但可以成为一个图层的来源。

```cpp
AEGP_GetItemType(
 AEGP_ItemH itemH,
 AEGP_ItemType *item_typeP);
```

项目是以下类型中的一种。

- `AEGP_ItemType_NONE`
- `AEGP_ItemType_FOLDER`
- `AEGP_ItemType_COMP`
- `AEGP_ItemType_SOLID`
- `AEGP_ItemType_FOOTAGE`

#### AEGP_GetTypeName

获取类型的名称。(名称长度到`AEGP_MAX_TYPE_NAME_LEN + 1`)。

```cpp
AEGP_GetTypeName(
 AEGP_ItemType item_type,
 A_char *nameZ);
```

#### AEGP_GetItemName

获取项目名称。(名称长度没有限制)。`unicode_namePH`指向`A_UTF16Char`(包含空尾的 UTF16 字符串)。
它必须用`AEGP_FreeMemHandle`处置。

```cpp
AEGP_GetItemName(
 AEGP_PluginID pluginID,
 AEGP_ItemH itemH,
 AEGP_MemHandle *unicode_namePH);
```

#### AEGP_SetItemName

指定 AEGP_ItemH 的名称。(名称长度没有限制)。可撤销。

```cpp
AEGP_SetItemName(
 AEGP_ItemH itemH,
 const A_UTF16Char *nameZ);
```

#### AEGP_GetItemID

返回项目的唯一 ID，该 ID 在项目的保存和加载过程中持续存在。

```cpp
AEGP_GetItemID(
 AEGP_ItemH itemH,
 A_long *item_idPL);
```

#### AEGP_GetItemFlags

获取一个项目的属性。

```cpp
AEGP_GetItemFlags(
 AEGP_ItemH itemH,
 AEGP_ItemFlags *item_flagsP);
```

标志值(可以是 OR'd together)。

- `AEGP_ItemFlag_MISSING`
- `AEGP_ItemFlag_HAS_PROXY`
- `AEGP_ItemFlag_USING_PROXY`
- `AEGP_ItemFlag_MISSING_PROXY`
- `AEGP_ItemFlag_HAS_VIDEO`
- `AEGP_ItemFlag_HAS_AUDIO`
- `AEGP_ItemFlag_STILL`
- `AEGP_ItemFlag_HAS_ACTIVE_AUDIO`

与 "HAS_AUDIO "标志不同，这个位标志只有在编译器至少有一个层实际有音频时才会设置。

#### AEGP_SetItemUseProxy

切换项目的代理使用。可撤销。

```cpp
AEGP_SetItemUseProxy(
 AEGP_ItemH itemH,
 A_Boolean use_proxyB);
```

#### AEGP_GetItemParentFolder

获取包含项目的文件夹。

```cpp
AEGP_GetItemParentFolder(
 AEGP_ItemH itemH,
 AEGP_ItemH *parent_itemPH);
```

#### AEGP_SetItemParentFolder

设置一个项目的父文件夹。可撤销。

```cpp
AEGP_SetItemParentFolder(
 AEGP_ItemH itemH,
 AEGP_ItemH parent_folderH);
```

#### AEGP_GetItemDuration

获取项目的持续时间，以秒为单位。

```cpp
AEGP_GetItemDuration(
 AEGP_ItemH itemH,
 A_Time *durationPT);
```

#### AEGP_GetItemCurrentTime

获取项目中的当前时间。渲染时不更新。

```cpp
AEGP_GetItemCurrentTime(
 AEGP_ItemH itemH,
 A_long *curr_timePT);
```

#### AEGP_GetItemDimensions

获取项目的宽度和高度。

```cpp
AEGP_GetItemDimensions(
 AEGP_ItemH itemH,
 A_long *widthPL)
 A_long *heightPL);
```

#### AEGP_GetItemPixelAspectRatio

获取一个像素的宽度，假设其高度为 1.0，作为分子大于分母。

```cpp
AEGP_GetItemPixelAspectRatio(
 AEGP_ItemH itemH,
 A_Ratio *ratioPRt);
```

#### AEGP_DeleteItem

从所有合成中移除项目。可以撤销。
调用此函数后不要使用`AEGP_ItemH`。

```cpp
AEGP_DeleteItem(
 AEGP_ItemH itemH);
```

#### AEGP_GetItemSolidColor

在`AEGP_ItemSuite4`中被删除。请参见[AEGP_FootageSuite5]中的`AEGP_GetSolidFootageColor`。
给定一个实体项目，返回其颜色。

```cpp
AEGP_GetItemSolidColor(
 AEGP_ItemH itemH,
 PF_Pixel *PF_Pixel);
```

#### AEGP_SetSolidColor

在`AEGP_ItemSuite4`中被删除。参见[AEGP_FootageSuite5]的`AEGP_SetSolidFootageColor`。
设置现有实体的颜色(如果`itemH`不是实体，则错误)。

```cpp
AEGP_SetSolidColor(
 AEGP_ItemH itemH,
 AEGP_ColorVal color);
```

#### AEGP_SetSolidDimensions

在`AEGP_ItemSuite4`中删除。参见[AEGP_FootageSuite5]的`AEGP_SetSolidFootageDimensions`。
设置现有实体的尺寸(如果 "itemH "不是一个实体，则错误)。

```cpp
AEGP_SetSolidDimensions(
 AEGP_ItemH itemH,
 A_short widthS,
 A_short heightS);
```

#### AEGP_CreateNewFolder

在项目中创建一个新的文件夹。新创建的文件夹由 After Effects 分配和拥有。
Passing `NULL` for `parent_folderH0` creates the folder at the project’s root.

```cpp
AEGP_CreateNewFolder(
 const A_UTF16Char *nameZ,
 AEGP_ProjectH projH),
 AEGP_ItemH parentH0),
 AEGP_ItemH *new_folderPH);
```

#### AEGP_SetItemCurrentTime

Sets the current time within a given `itemH`.

```cpp
AEGP_SetItemCurrentTime(
 AEGP_ItemH itemH,
 const A_Time *new_timePT);
```

#### AEGP_GetItemCommentLength

Removed in `ItemSuite9`. Retrieves the length (in characters) of the `itemH's` comment.

```cpp
AEGP_GetItemCommentLength(
 AEGP_ItemH itemH,
 A_u_long *buf_sizePLu);
```

#### AEGP_GetItemComment

Updated to support Unicode in `ItemSuite9`, available in 14.1. Retrieves the `itemH's` comment.

```cpp
AEGP_GetItemComment(
 AEGP_ItemH itemH,
 AEGP_MemHandle *unicode_namePH);
```

#### AEGP_SetItemComment

Updated to support Unicode in `ItemSuite9`, available in 14.1. Sets the `itemH's` comment.

```cpp
AEGP_SetItemComment(
 AEGP_ItemH itemH,
 const A_UTF16Char *commentZ);
```

#### AEGP_GetItemLabel

Retrieves an item’s label.

```cpp
AEGP_GetItemLabel(
 AEGP_ItemH itemH,
 AEGP_LabelID *labelP);
```

#### AEGP_SetItemLabel

Sets an item’s label.

```cpp
AEGP_SetItemLabel(
 AEGP_ItemH itemH,
 AEGP_LabelID label);
```

#### AEGP_GetItemMRUView

Gets an item’s most recently used view.The view can be used with two calls in the `AEGP_ColorSettingsSuite`,to perform a color transform on a pixel buffer from working to view color space.

```cpp
AEGP_GetItemMRUView(
 AEGP_ItemH itemH,
 AEGP_ItemViewP *mru_viewP);
```

::: tip

`AEGP_RenderNewItemSoundData()` used to be here, but is now part of [AEGP_RenderSuite4].

:::

## Managing Selections

This suite manages selection states, mirroring the functionality supplied by vectors in the C++ Standard Template Library.

Many types of items may be simultaneously selected in After Effects; `AEGP_CollectionItems` are unions of layer, mask, effect, stream, mask vertex, and keyframe items.

First acquire the current collection, then iterate across its members to ensure that whatever your AEGP does is applicable to each.

We’ve added `AEGP_Collection2H` and `AEGP_CollectionItemV2` so that selected dynamic streams can be handled with the `AEGP_CollectionSuite`.

### AEGP_CollectionSuite2

#### AEGP_NewCollection

Creates and returns a new, empty collection.To obtain the current composition’s selection as a collection, use `AEGP_GetNewCollectionFromCompSelection`.

```cpp
AEGP_NewCollection(
 AEGP_PluginID plugin_id,
 AEGP_Collection2H *collectionPH);
```

#### AEGP_DisposeCollection

Disposes of a collection.

```cpp
AEGP_DisposeCollection(
 AEGP_Collection2H collectionH);
```

#### AEGP_GetCollectionNumItems

Returns the number of items contained in the given collection.

```cpp
AEGP_GetCollectionNumItems(
 AEGP_Collection2H collectionH,
 A_u_long *num_itemsPL);
```

#### AEGP_GetCollectionItemByIndex

Retrieves (creates and populates) the index’d collection item.

```cpp
AEGP_GetCollectionItemByIndex(
 AEGP_Collection2H collectionH,
 A_u_long indexL,
 AEGP_CollectionItemV2 *itemP);
```

#### AEGP_CollectionPushBack

Adds an item to the given collection.

```cpp
AEGP_CollectionPushBack(
 AEGP_Collection2H collectionH,
 const AEGP_CollectionItemV2 *itemP);
```

#### AEGP_CollectionErase

Removes an index’d item (or items) from a given collection. NOTE: this range is exclusive,like STL iterators. To erase the first item, you would pass 0 and 1, respectively.

```cpp
AEGP_CollectionErase(
 AEGP_Collection2H collectionH,
 A_u_long index_firstL,
 A_u_long index_lastL);
```

### Ownership Of Collection Items

When `AEGP_StreamRefHs` are inserted into a collection, they are adopted by the collection; do not free them.

`AEGP_EffectRefHs`, on the other hand, are not adopted, and must be freed by the calling AEGP.

## Manipulate Compositions

Provide information about the compositions in a project, and create cameras, lights, and solids.

### AEGP_CompSuite11

#### AEGP_GetCompFromItem

Retrieves the handle to the composition, given an item handle.
Returns `NULL` if `itemH` is not an `AEGP_CompH`.

```cpp
AEGP_GetCompFromItem(
 AEGP_ItemH itemH,
 AEGP_CompH *compPH);
```

#### AEGP_GetItemFromComp

Used to get the item handle, given a composition handle.

```cpp
AEGP_GetItemFromComp(
 AEGP_CompH compH,
 AEGP_ItemH *itemPH);
```

#### AEGP_GetCompDownsampleFactor

Returns current downsample factor. Measured in pixels X by Y.Users can choose a custom downsample factor with independent X and Y.

```cpp
AEGP_GetCompDownsampleFactor(
 AEGP_CompH compH,
 AEGP_DownsampleFactor *dsfP);
```

#### AEGP_SetCompDownsampleFactor

Sets the composition’s downsample factor.

```cpp
AEGP_SetCompDownsampleFactor(
 AEGP_CompH compH,
 AEGP_DownsampleFactor *dsfP);
```

#### AEGP_GetCompBGColor

Returns the composition background color.

```cpp
AEGP_GetCompBGColor(
 AEGP_CompH compH,
 AEGP_ColorVal *bg_colorP);
```

#### AEGP_SetCompBGColor

Sets a composition’s background color.

```cpp
AEGP_SetCompBGColor(
 AEGP_CompH compH,
 const AEGP_ColorVal *bg_colorP);
```

#### AEGP_GetCompFlags

Returns composition flags, or’d together.

```cpp
AEGP_GetCompFlags(
 AEGP_CompH compH,
 AEGP_CompFlags *AEGP_CompFlags);
```

- `AEGP_CompFlag_SHOW_ALL_SHY`
- `AEGP_CompFlag_ENABLE_MOTION_BLUR`
- `AEGP_CompFlag_ENABLE_TIME_FILTER`
- `AEGP_CompFlag_GRID_TO_FRAME`
- `AEGP_CompFlag_GRID_TO_FIELDS`
- `AEGP_CompFlag_USE_LOCAL_DSF`
- `AEGP_CompFlag_DRAFT_3D`
- `AEGP_CompFlag_SHOW_GRAPH`

#### AEGP_GetShowLayerNameOrSourceName

New in CC. Passes back true if the Comp’s timeline shows layer names, false if source names.This will open the comp as a side effect.

```cpp
AEGP_GetShowLayerNameOrSourceName(
 AEGP_CompH compH,
 A_Boolean *layer_names_shownPB);
```

#### AEGP_SetShowLayerNameOrSourceName

New in CC. Pass in true to have the Comp’s timeline show layer names, false for source names.This will open the comp as a side effect.

```cpp
AEGP_SetShowLayerNameOrSourceName(
 AEGP_CompH compH,
 A_Boolean *layer_names_shownPB);
```

#### AEGP_GetShowBlendModes

New in CC. Passes back true if the Comp’s timeline shows blend modes column, false if hidden.
This will open the comp as a side effect.

```cpp
AEGP_GetShowBlendModes(
 AEGP_CompH compH,
 A_Boolean *blend_modes_shownPB);
```

#### AEGP_SetShowBlendModes

New in CC. Pass in true to have the Comp’s timeline show the blend modes column, false to hide it.
This will open the comp as a side effect.

```cpp
AEGP_GetCompFlags(
 AEGP_CompH compH,
 A_Boolean show_blend_modesB);
```

#### AEGP_GetCompFramerate

Returns the composition’s frames per second.

```cpp
AEGP_GetCompFramerate(
 AEGP_CompH compH,
 A_FpLong *fpsPF);
```

#### AEGP_SetCompFramerate

Sets the composition’s frames per second.

```cpp
AEGP_SetCompFramerate(
 AEGP_CompH compH,
 A_FpLong *fpsPF);
```

#### AEGP_GetCompShutterAnglePhase

The composition shutter angle and phase.

```cpp
AEGP_GetCompShutterAnglePhase(
 AEGP_CompH compH,
 A_Ratio *angle,
 A_Ratio *phase);
```

#### AEGP_GetCompShutterFrameRange

The duration of the shutter frame, in seconds.

```cpp
AEGP_GetCompShutterFrameRange(
 AEGP_CompH compH,
 const A_Time *comp_timeP);
```

#### AEGP_GetCompSuggestedMotionBlurSamples

Retrieves the number of motion blur samples After Effects will perform in the given composition.

```cpp
AEGP_GetCompSuggestedMotionBlurSamples(
 AEGP_CompH compH,
 A_long *samplesPL)
```

#### AEGP_SetCompSuggestedMotionBlurSamples

Specifies the number of motion blur samples After Effects will perform in the given composition. Undoable.

```cpp
AEGP_SetCompSuggestedMotionBlurSamples(
 AEGP_CompH compH,
 A_long samplesL);
```

#### AEGP_GetCompMotionBlurAdaptiveSampleLimit

New in CC. Retrieves the motion blur adaptive sample limit for the given composition. As of CC, a new comp defaults to 128.

```cpp
AEGP_GetCompMotionBlurAdaptiveSampleLimit(
 AEGP_CompH compH,
 A_long *samplesPL)
```

#### AEGP_SetCompMotionBlurAdaptiveSampleLimit

New in CC. Specifies the motion blur adaptive sample limit for the given composition.As of CC, both the limit and the suggested values are clamped to [2,256] range and the limit value will not be allowed less than the suggested value.
Undoable.

```cpp
AEGP_SetCompMotionBlurAdaptiveSampleLimit(
 AEGP_CompH compH,
 A_long samplesL);
```

#### AEGP_GetCompWorkAreaStart

Get the time where the current work area starts.

```cpp
AEGP_GetCompWorkAreaStart(
 AEGP_CompH compH,
 A_Time *startPT);
```

#### AEGP_GetCompWorkAreaDuration

Get the duration of a composition’s current work area, in seconds.

```cpp
AEGP_GetCompWorkAreaDuration(
 AEGP_CompH compH,
 A_Time *durationPT);
```

#### AEGP_SetCompWorkAreaStartAndDuration

设置工作区的开始和持续时间，以秒为单位。可以不做。对这个函数的一次调用就足以设置图层的起始点和持续时间；没有必要对它调用两次，对每个时间空间调用一次。

```cpp
AEGP_SetCompWorkAreaStartAndDuration(
 AEGP_CompH compH,
 const A_Time *startPT)
 const A_Time *durationPT);
```

#### AEGP_CreateSolidInComp

在合成中创建一个具有指定宽度、高度、颜色和持续时间的新实体。可以撤销。如果你为持续时间传入`NULL`，After Effects 会使用它对新静止的持续时间的偏好。
如果传递 NULL 或无效的时间尺度，持续时间将被设置为合成的长度。

```cpp
AEGP_CreateSolidInComp(
 const A_UTF16Char *utf_nameZ,
 A_Long widthL,
 A_Long heightL,
 const PF_Pixel *color,
 AEGP_CompH parent_compH,
 const A_Time *durationPT0,
 AEGP_LayerH *new_solidPH);
```

#### AEGP_CreateCameraInComp

创建一个摄像机并将其添加到指定的合成中。创建后，你可以使用 AEGP[流套件]来操作摄像机的参数流。
要指定一个双节点摄像机，请使用`AEGP_SetLayerFlag`。
来设置`AEGP_LayerFlag_LOOK_AT_POI`。

```cpp
AEGP_CreateCameraInComp(
 const A_UTF16Char *utf_nameZ,
 A_FloatPoint center_point,
 AEGP_CompH parent_compH,
 AEGP_LayerH *new_cameraPH);
```

#### AEGP_CreateLightInComp

创建并添加一个灯光到指定的合成。创建后，你可以使用 AEGP[流套件]操作灯光的参数流。

```cpp
AEGP_CreateLightInComp(
 const A_UTF16Char *utf_nameZ,
 A_FloatPoint center_point,
 AEGP_CompH parent_compH,
 AEGP_LayerH *new_lightPH);
```

#### AEGP_CreateComp

为该项目创建一个新的合成。
如果你没有提供一个父文件夹，该合成将在项目的根层。可撤销。

```cpp
AEGP_CreateComp(
 AEGP_ItemH parent_folderHO,
 const A_UTF16Char *utf_nameZ,
 A_Long widthL,
 A_Long heightL,
 const A_Ratio *pixel_aspect_ratioPRt,
 const A_Time *durationPT,
 const A_Ratio *frameratePRt,
 AEGP_CompH *new_compPH);
```

#### AEGP_GetNewCollectionFromCompSelection

从给定的合成中选择的项目创建一个新的 AEGP_Collection2H。
插件负责处理`AEGP_Collection2H`。

```cpp
AEGP_GetNewCollectionFromCompSelection(
 AEGP_PluginID plugin_id,
 AEGP_CompH compH,
 AEGP_Collection2H *collectionPH);
```

#### AEGP_SetSelection

将在给定的合成中的选择设置为给定的`AEGP_Collection2H`。 如果 "AEGP_Collection2H "的成员不可用，将返回一个错误。 不要假设一个合成在操作之间没有变化；总是使用一个新的`AEGP_Collection2H`。

```cpp
AEGP_SetSelection(
 AEGP_CompH compH,
 AEGP_Collection2H collectionH);
```

#### AEGP_GetCompDisplayStartTime

获取一个合成的显示的开始时间。

```cpp
AEGP_GetCompDisplayStartTime(
 AEGP_CompH compH,
 const A_Time *start_timePT);
```

#### AEGP_SetCompDisplayStartTime

不能撤销。设置一个合成的开始时间(对合成的持续时间没有影响)。

```cpp
AEGP_SetCompDisplayStartTime(
 AEGP_CompH compH,
 const A_Time *start_timePT);
```

#### AEGP_SetCompDuration

可撤消。设置给定合成的持续时间。

```cpp
AEGP_SetCompDuration(
 AEGP_CompH compH,
 const A_Time *durationPT);
```

#### AEGP_CreateNullInComp

在合成中创建一个 "空对象"(对于从 3D 应用程序翻译到 After Effects 的项目很有用)。如果你把 "NULL "作为持续时间，After Effects 会使用它对新静止的持续时间的偏好。
如果传入 0，或者传入一个无效的时间尺度，持续时间就会被设置为合成的长度。

```cpp
AEGP_CreateNullInComp(
 const A_UTF16Char *utf_nameZ,
 AEGP_CompH parent_compH,
 const A_Time *durationPT0,
 AEGP_LayerH *new_null_solidPH);
```

#### AEGP_SetCompPixelAspectRatio

设置合成物的像素长宽比。

```cpp
AEGP_SetCompPixelAspectRatio(
 AEGP_CompH compH,
 const A_Ratio *parPRt);
```

#### AEGP_CreateTextLayerInComp

在 CS6 中更新。在合成中创建一个文本层，并返回其`AEGP_LayerH`。

```cpp
AEGP_CreateTextLayerInComp(
 AEGP_CompH parent_compH,
 A_Boolean select_new_layerB,
 AEGP_LayerH *new_text_lyrPH);
```

#### AEGP_CreateBoxTextLayerInComp

在 CS6 中更新。创建一个新的文本图层，并返回其`AEGP_LayerH`。

```cpp
AEGP_CreateBoxTextLayerInComp(
 AEGP_CompH parent_compH,
 A_Boolean select_new_layerB,
 A_FloatPoint box_dimensions,
 AEGP_LayerH *new_text_layerPH);
```

#### AEGP_SetCompDimensions

设置合成的尺寸。可撤销。

```cpp
AEGP_SetCompDimensions(
 AEGP_CompH compH,
 A_long widthL,
 A_long heightL);
```

#### AEGP_DuplicateComp

复制合成。可撤销。

```cpp
AEGP_DuplicateComp(
 AEGP_CompH compH,
 AEGP_CompH *new_compPH);
```

#### AEGP_GetCompFrameDuration

检索合成中某一帧的持续时间。

```cpp
AEGP_GetCompFrameDuration(
 AEGP_CompH compH,
 A_Time *timeP);
```

#### AEGP_GetMostRecentlyUsedComp

返回最近一次使用的合成。

```cpp
AEGP_GetMostRecentlyUsedComp(
 AEGP_CompH *compPH);
```

#### AEGP_CreateVectorLayerInComp

创建并返回一个新矢量图层的句柄。

```cpp
AEGP_CreateVectorLayerInComp(
 AEGP_CompH parent_compH,
 AEGP_LayerH *new_vec_layerPH);
```

#### AEGP_GetNewCompMarkerStream

返回一个 AEGP_StreamRefH 到合成的标记流。必须由调用者处理。

```cpp
AEGP_GetNewCompMarkerStream(
 AEGP_PluginID aegp_plugin_id,
 AEGP_CompH parent_compH,
 AEGP_StreamRefH *streamPH);
```

#### AEGP_GetCompDisplayDropFrame

传回一个布尔值，表示指定的合成是否使用降帧时间码。

```cpp
AEGP_GetCompDisplayDropFrame(
 AEGP_CompH compH,
 A_Boolean *dropFramePB);
```

#### AEGP_SetCompDisplayDropFrame

在指定的合成中设置时间码的落格性。

```cpp
AEGP_SetCompDisplayDropFrame(
 AEGP_CompH compH,
 A_Boolean dropFrameB);
```

#### AEGP_ReorderCompSelection

将选区移动到某一图层索引。与`AEGP_SetSelection()`一起使用。

```cpp
AEGP_SetCompDisplayDropFrame(
 AEGP_CompH compH,
 A_long index);
```

## Work with Footage

提供有关镜头或项目或合成的信息。当获取和设置片段的解释时，有可能指定不兼容的选项。

如果你在开发过程中遇到警告和错误，一定要以原子方式进行所有相关的修改，并重新评估你所执行的操作的逻辑。

例如，改变镜头的下拉解释是行不通的，除非它的原始帧率和符合的帧率之间存在差异。

根据你要完成的任务，在这一点上中止你的所有操作可能是有意义的，通知用户遇到的问题。

### AEGP_FootageSuite5

#### AEGP_GetMainFootageFromItem

如果 item 不是一个 footage item，则返回错误。用于将一个项目句柄转换为一个镜头句柄。

```cpp
AEGP_GetMainFootageFromItem(
 AEGP_ItemH itemH,
 AEGP_FootageH *footagePH);
```

#### AEGP_GetProxyFootageFromItem

如果项目没有代理，则返回错误。返回代理镜头句柄。注意：一个合成可以有一个代理。

```cpp
AEGP_GetProxyFootageFromItem(
 AEGP_ItemH itemH,
 AEGP_FootageH *proxy_ftgPH);
```

#### AEGP_GetFootageNumFiles

返回数据(RGBA 或音频)文件的数量，以及每一帧的文件数量(如果镜头有辅助通道，可能大于 1)。

```cpp
AEGP_GetFootageNumFiles(
 AEGP_FootageH footageH,
 A_long *num_filesPL0,
 A_long *files_per_frmPL0);
```

#### AEGP_GetFootagePath

获取完全实现的镜头源文件的路径。
检索一段录像(或一个录像序列中的指定帧)的镜头路径。
`frame_numL`的范围是`0到num_main_files`，使用[AEGP_FootageSuite5]中的`AEGP_GetFootageNumFiles`获得。
`AEGP_FOOTAGE_MAIN_FILE_INDEX`是主文件。路径是一个以 NULL 结尾的`A_UTF16Char`字符串的句柄，必须用 AEGP_FreeMemHandle 处置。

```cpp
AEGP_GetFootagePath(
 AEGP_FootageH footageH,
 A_long frame_numL,
 A_long file_indexL,
 AEGP_MemHandle *unicode_pathPH);
```

#### AEGP_GetFootageSignature

检索指定镜头的签名。

```cpp
AEGP_GetFootageSignature(
 AEGP_FootageH footageH,
 AEGP_FootageSignature *sigP);
```

签名将是以下内容之一。

- `AEGP_FootageSignature_NONE`
- `AEGP_FootageSignature_MISSING`
- `AEGP_FootageSignature_SOLID`

#### AEGP_NewFootage

创建一个新的镜头项目。文件路径是一个以 NULL 结尾的 UTF-16 字符串，带有平台分隔符。
注意带冒号的镜头文件名是不允许的，因为冒号在 HFS+文件系统中被用作路径分隔符。

```cpp
AEGP_NewFootage(
 AEGP_PluginID aegp_plugin_id,
 const A_UTF16Char *pathZ,
 const AEGP_FootageLayerKey *layer_infoP0,
 const AEGP_FileSequenceImportOptions *sequence_optionsP0,
 AEGP_InterpretationStyle interp_style,
 void *reserved,
 AEGP_FootageH *footagePH);
```

注意可选的参数。如果`allow_interpretation_dialogB`为 FALSE，After Effects 将猜测 alpha 解释。

```cpp
typedef struct {
 A_long layer_idL;
 A_long layer_indexL
 char *nameAC;
 AEGP_LayerDrawStyle draw_style;
} AEGP_FootageLayerKey;
```

`AEGP_LayerDrawStyle`可以是。

- `AEGP_LayerDrawStyle_LAYER_BOUNDS`
- `AEGP_LayerDrawStyle_DOCUMENT_BOUNDS`

`AEGP_InterpretationStyle`可以是。

- `AEGP_InterpretationStyle_NO_DIALOG_GUESS` Will guess alpha interpretation even if file contains unknown alpha interpretation and user pref says to ask user.
- `AEGP_InterpretationStyle_DIALOG_OK` Optionally can show a dialog.
- `AEGP_InterpretationStyle_NO_DIALOG_NO_GUESS` Used for replace footage implementation.

#### AEGP_AddFootageToProject

在项目中添加一个镜头项目。镜头将被项目采用，并且只能添加一次。这是可以撤消的；如果撤消了，不要处理返回的添加项目。

```cpp
AEGP_AddFootageToProject(
 AEGP_FootageH footageH,
 AEGP_ItemH folderH,
 AEGP_ItemH *add_itemPH0);
```

#### AEGP_SetItemProxyFootage

设置镜头作为项目的代理。将被项目采用。这是可以撤消的；如果撤消了，不要处理返回的新增项目。

```cpp
AEGP_SetItemProxyFootage(
 AEGP_FootageH footageH,
 AEGP_ItemH itemH);
```

#### AEGP_ReplaceItemMainFootage

替换一个项目的镜头。该项目将取代该项目的主要尺码。这是可以撤销的；如果撤销了，不要处理返回的新增项目。

```cpp
AEGP_ReplaceItemMainFootage(
 AEGP_FootageH footageH,
 AEGP_ItemH itemH);
```

#### AEGP_DisposeFootage

删除一个镜头项目。不要处理你没有创建的镜头，或已经添加到项目中的镜头。

```cpp
AEGP_DisposeFootage(
 AEGP_FootageH footageH);
```

#### AEGP_GetFootageInterpretation

弹出一个 AEGP_FootageInterp，描述`AEGP_FootageH`的设置。除了使用这个函数外，没有办法创建一个有效的`AEGP_FootageInterp`。

```cpp
AEGP_GetFootageInterpretation(
 const AEGP_ItemH itemH,
 A_Boolean proxyB,
 AEGP_FootageInterp *interpP);
```

如果 proxyB 是`TRUE`
，代理脚步的设置就会被检索出来。

#### AEGP_SetFootageInterpretation

将`AEGP_FootageInterp`中的设置应用于`AEGP_FootageH`。可撤销。

```cpp
AEGP_SetFootageInterpreta tion(
 const AEGP_ItemH itemH,
 A_Boolean proxyB,
 const AEGP_FootageInterp *interpP);
```

如果`proxyB`为`TRUE`，代理镜头的设置将被修改。

#### AEGP_GetFootageLayerKey

填充一个描述镜头的`AEGP_FootageLayerKey`。

```cpp
AEGP_GetFootageLayerKey(
 AEGP_FootageH footageH,
 AEGP_FootageLayerKey* layerKeyP);
```

#### AEGP_NewPlaceholderFootage

弃用。在项目中添加一个新的占位符镜头项目。对丢失的镜头使用这个函数将导致用户搜索每一个丢失的文件，不管它们是否都在同一个目录中。
可撤销。

```cpp
AEGP_NewPlaceholderFootage(
 AEGP_PluginID plugin_id,
 const A_char *nameZ,
 A_long width,
 A_long height,
 const A_Time *durationPT,
 AEGP_FootageH *footagePH);
```

#### AEGP_NewPlaceholderFootageWithPath

这是添加对此刻无法找到的片段的引用的时髦新方法。文件路径是一个以 NULL 结尾的 UTF-16 字符串，带有平台分隔符。
在 CS6 和更早的版本中，file_type 被忽略了，我们之前建议将其设置为`AEIO_FileType_NONE`。从 CC 开始，`AEIO_FileType_NONE`现在是一个警告条件。如果你通过`AEIO_FileType_ANY`，那么路径必须存在。
如果路径可能不存在，传递`AEIO_FileType_DIR`代表文件夹，或`AEIO_FileType_GENERIC`代表文件。

```cpp
AEGP_NewPlaceholderFootageWithPath(
 AEGP_PluginID plugin_id,
 const A_UTF16Char *pathZ,
 AEGP_Platform path_platform,
 AEIO_FileType file_type,
 A_long widthL,
 A_long heightL,
 const A_Time *durationPT,
 AEGP_FootageH *footagePH);
```

#### AEGP_NewSolidFootage

这就是添加实体的方法。在镜头被添加到项目中之前，调用者拥有`AEGP_FootageH`。
(并且必须在它没有被添加到项目中的情况下处理掉它)。

```cpp
AEGP_NewSolidFootage(
 const A_char *nameZ,
 A_long width,
 A_long height,
 const AEGP_ColorVal *colorP,
 AEGP_FootageH *footagePH);
```

#### AEGP_GetSolidFootageColor

返回给定实体的颜色。如果`AEGP_ItemH`不是一个实体，则返回错误。

```cpp
AEGP_GetSolidFootageColor(
 AEGP_ItemH itemH,
 A_Boolean proxyB,
 AEGP_ColorVal *colorP);
```

如果`proxyB`是`TRUE`，将检索代理实体的颜色。

#### AEGP_SetSolidFootageColor

设置实体的颜色。可以撤销。

```cpp
AEGP_SetSolidFootageColor(
 AEGP_ItemH itemH,
 A_Boolean proxyB,
 AEGP_ColorVal *colorP);
```

如果`proxyB`是`TRUE`，代理实体的颜色被设置。

#### AEGP_SetSolidFootageDimensions

设置实体的尺寸。可以撤消。

```cpp
AEGP_SetSolidFootageDimensions(
 AEGP_ItemH itemH,
 A_Boolean proxyB,
 A_long widthL,
 A_long heightL);
```

如果`proxyB`是`TRUE`，代理实体的尺寸会被修改。如果该项目不是实体，则返回错误。

#### AEGP_GetFootageSoundDataFormat

检索片段中的音频数据信息(通过填充你传入的`AEGP_SoundDataFormat`)。

```cpp
AEGP_GetFootageSoundDataFormat(
 AEGP_FootageH footageH,
 AEGP_SoundDataFormat *formatP);
```

#### AEGP_GetFootageSequenceImportOptions

填充并返回一个`AEGP_FileSequenceImportOptions`，描述给定的`AEGP_FootageH`。

```cpp
AEGP_GetFootageSequenceImportOptions(
 AEGP_FootageH footageH,
 AEGP_FileSequenceImportOptions *optionsP);
```

### AEGP_FootageInterp Structure

#### AEGP_InterlaceLabel il

镜头项目的隔行扫描设置。

```cpp
A_u_long signature; // 'FIEL'
A_short version;
FIEL_Type type;
FIEL_Order order;
A_u_long reserved;
```

FIEL_Type 是以下之一。

- `FIEL_Type_FRAME_RENDERED`
- `FIEL_Type_INTERLACED`
- `FIEL_Type_HALF_HEIGHT`
- `FIEL_Type_FIELD_DOUBLED`

`FIEL_Type_FIELD_DOUBLED`表示每秒 60 个全尺寸场的双倍帧。
`FIEL_Order`是`FIEL_Order_UPPER_FIRST`或`FIEL_Order_LOWER_FIRST`。

#### AEGP_AlphaLabel al

```cpp
AEGP_AlphaFlag flags;
A_u_char redCu;
A_u_char greenCu;
A_u_char blueCu;
```

`AEGP_AlphaFlags`是以下的一个或多个，或一起。

- `AEGP_AlphaPremul`
- `AEGP_AlphaInverted`
- `AEGP_AlphaIgnore`

如果`AEGP_AlphaPremul`没有被设置，则假定为直线α。`AEGP_AlphaInverted`表示高值透明，而不是低值。

#### AEGP_PulldownPhase pd

表示在 3:2 pulldown 中使用的相位。以下之一。

- `AEGP_PulldownPhase_NO_PULLDOWN`,
- `AEGP_PulldownPhase_WSSWW`,
- `AEGP_PulldownPhase_SSWWW`,
- `AEGP_PulldownPhase_SWWWS`,
- `AEGP_PulldownPhase_WWWSS`,
- `AEGP_PulldownPhase_WWSSW`,
- `AEGP_PulldownPhase_WWWSW`,
- `AEGP_PulldownPhase_WWSWW`,
- `AEGP_PulldownPhase_WSWWW`,
- `AEGP_PulldownPhase_SWWWW`,
- `AEGP_PulldownPhase_WWWWS`

#### AEGP_LoopBehavior loop

表示镜头应该循环的次数。

```cpp
A_long loops;
A_long reserved;
```

#### A_Ratio pix_aspect_ratio

表示镜头的像素长宽比(x 大于 y)。

#### A_FpLong native_fpsF

镜头项目的原始帧率(以每秒帧数计)。

#### A_FpLong conform_fpsF

镜头项目所使用的帧速率。

#### A_long depthL

镜头的像素深度。以下内容之一。

- `AEGP_Footage_Depth_1`
- `AEGP_Footage_Depth_2`
- `AEGP_Footage_Depth_4`
- `AEGP_Footage_Depth_8`
- `AEGP_Footage_Depth_16`
- `AEGP_Footage_Depth_24`
- `AEGP_Footage_Depth_30`
- `AEGP_Footage_Depth_32`
- `AEGP_Footage_Depth_GRAY_2`
- `AEGP_Footage_Depth_GRAY_4`
- `AEGP_Footage_Depth_GRAY_8`
- `AEGP_Footage_Depth_48`
- `AEGP_Footage_Depth_64`
- `AEGP_Footage_Depth_GRAY_16`

#### A_Boolean motion_dB

表示是否正在对该片段项目应用运动去隔行扫描。

## Manage Layers

AEGP_LayerSuite "提供合成中的层的信息，以及源和层时间之间的关系。

由于大多数 After Effects 的使用都归结为层的操作，这是我们 API 中最大的功能套件之一。

### AEGP_LayerSuite8

#### AEGP_GetCompNumLayers

获取一个合成中的层数。

```cpp
AEGP_GetCompNumLayers(
 AEGP_CompH compH,
 A_long *num_layersPL);
```

#### AEGP_GetCompLayerByIndex

从一个合成物中获得一个`AEGP_LayerH`。零是最前面的层。

```cpp
AEGP_GetCompLayerByIndex(
 AEGP_CompH compH,
 A_long layer_indexL,
 AEGP_LayerH *layerPH);
```

#### AEGP_GetActiveLayer

获取活动层。如果图层或效果控制调色板是活动的，活动图层是与窗口中最前面的标签相关的。如果一个合成或时间轴窗口处于活动状态，活动层是选定的层(如果只有一个被选定；否则返回`NULL`)。

```cpp
AEGP_GetActiveLayer(
 AEGP_LayerH *layerPH);
```

#### AEGP_GetLayerIndex

获取图层的索引(0 是合成中最上面的图层)。

```cpp
AEGP_GetLayerIndex(
 AEGP_LayerH layerH,
 A_long *layer_indexPL);
```

#### AEGP_GetLayerSourceItem

获取该层的源项的 AEGP_ItemH。

```cpp
AEGP_GetLayerSourceItem(
 AEGP_LayerH layerH,
 AEGP_ItemH *source_itemPH);
```

#### AEGP_GetLayerSourceItemID

检索给定的`AEGP_LayerH`的 ID。当在`AEGP_StreamVal`中寻找特定层的 ID 时，这很有用。

```cpp
AEGP_GetLayerSourceItemID(
 AEGP_LayerH layerH,
 A_long *source_idPL);
```

#### AEGP_GetLayerParentComp

获取包含该层的合成的 AEGP_CompH。

```cpp
AEGP_GetLayerParentComp(
 AEGP_LayerH layerH,
 AEGP_CompH *compPH);
```

#### AEGP_GetLayerName

获取一个图层的名称。`utf_layer_namePH0`和`utf_source_namePH0`都指向空尾的 UTF-16 字符串。它们必须用`AEGP_FreeMemHandle`来处理。

```cpp
AEGP_GetLayerName(
 AEGP_PluginID pluginID,
 AEGP_LayerH layerH,
 AEGP_MemHandle *utf_layer_namePH0,
 AEGP_MemHandle *utf_source_namePH0);
```

#### AEGP_GetLayerQuality

获取一个图层的质量。

```cpp
AEGP_GetLayerQuality(
 AEGP_LayerH layerH,
 AEGP_LayerQuality *qualityP);
```

图层质量是以下标志之一。

- `AEGP_LayerQual_NONE`
- `AEGP_LayerQual_WIREFRAME`
- `AEGP_LayerQual_DRAFT`
- `AEGP_LayerQual_BEST`

#### AEGP_SetLayerQuality

设置一个图层的质量(见上面的标志值)。可撤销。

```cpp
AEGP_SetLayerQuality(
 AEGP_LayerH layerH,
 AEGP_LayerQuality quality);
```

#### AEGP_GetLayerFlags

获取一个图层的标志。

```cpp
AEGP_GetLayerFlags(
 AEGP_LayerH layerH,
 AEGP_LayerFlags *layer_flagsP);
```

- `AEGP_LayerFlag_NONE`
- `AEGP_LayerFlag_VIDEO_ACTIVE`,
- `AEGP_LayerFlag_AUDIO_ACTIVE`,
- `AEGP_LayerFlag_EFFECTS_ACTIVE`,
- `AEGP_LayerFlag_MOTION_BLUR`,
- `AEGP_LayerFlag_FRAME_BLENDING`,
- `AEGP_LayerFlag_LOCKED`,
- `AEGP_LayerFlag_SHY`,
- `AEGP_LayerFlag_COLLAPSE`,
- `AEGP_LayerFlag_AUTO_ORIENT_ROTATION`,
- `AEGP_LayerFlag_ADJUSTMENT_LAYER`,
- `AEGP_LayerFlag_TIME_REMAPPING`,
- `AEGP_LayerFlag_LAYER_IS_3D`,
- `AEGP_LayerFlag_LOOK_AT_CAMERA`,
- `AEGP_LayerFlag_LOOK_AT_POI`,
- `AEGP_LayerFlag_SOLO`,
- `AEGP_LayerFlag_MARKERS_LOCKED`,
- `AEGP_LayerFlag_NULL_LAYER`,
- `AEGP_LayerFlag_HIDE_LOCKED_MASKS`,
- `AEGP_LayerFlag_GUIDE_LAYER`,
- `AEGP_LayerFlag_ENVIRONMENT_LAYER`,
- `AEGP_LayerFlag_ADVANCED_FRAME_BLENDING`, `True` only if pixel motion frame blending is on for the layer.
- `AEGP_LayerFlag_SUBLAYERS_RENDER_SEPARATELY`, Used to get/set the state of per-character 3D enablement on a text layer.
- `AEGP_LayerFlag_ENVIRONMENT_LAYER`, New in CS6.

#### AEGP_SetLayerFlag

一次设置一个图层标志。可撤消。

```cpp
AEGP_SetLayerFlag(
 AEGP_LayerH layerH,
 AEGP_LayerFlags single_flag,
 A_Boolean valueB);
```

#### AEGP_IsLayerVideoReallyOn

确定该层的视频是否可见。这是必要的，以便考虑到合成中其他层的 "独奏 "状态；非独奏的层仍然是打开的。

```cpp
AEGP_IsLayerVideoReallyOn(
 AEGP_LayerH layerH,
 A_Boolean *onPB);
```

#### AEGP_IsLayerAudioReallyOn

解释合成中其他层的独奏状态。

```cpp
AEGP_IsLayerAudioReallyOn(
 AEGP_LayerH layerH,
 A_Boolean *onPB);
```

#### AEGP_GetLayerCurrentTime

获取当前时间，在图层或合成的时间空间。这个值在渲染过程中不被更新。
注意：如果一个图层在时间 0 以外的地方开始，或在时间上被拉伸到 100%以外，图层时间和合成时间是不同的。

```cpp
AEGP_GetLayerCurrentTime(
 AEGP_LayerH layerH,
 AEGP_LTimeMode time_mode,
 A_Time *curr_timePT);
```

#### AEGP_GetLayerInPoint

在合成时间或图层时间中获取第一个可见帧的时间。
在图层时间中，"in_pointPT "始终为 0。

```cpp
AEGP_GetLayerInPoint(
 AEGP_LayerH layerH,
 AEGP_LTimeMode time_mode,
 A_Time *in_pointPT);
```

#### AEGP_GetLayerDuration

在合成或图层时间中，获取图层的持续时间，单位是秒。

```cpp
AEGP_GetLayerDuration(
 AEGP_LayerH layerH,
 AEGP_LTimeMode time_mode,
 A_Time *durationPT);
```

#### AEGP_SetLayerInPointAndDuration

设置合成或图层时间中图层的持续时间和在点。可撤销。

```cpp
AEGP_SetLayerInPointAndDuration(
 AEGP_LayerH layerH,
 AEGP_LTimeMode time_mode,
 const A_Time *in_pointPT,
 const A_Time *durationPT);
```

#### AEGP_GetLayerOffset

获取从合成开始到图层时间 0 的偏移，以合成时间为单位。

```cpp
AEGP_GetLayerOffset(
 AEGP_LayerH layerH,
 A_Time *offsetPT);
```

#### AEGP_SetLayerOffset

设置从合成开始到图层第一帧的偏移，以合成时间为单位。可撤销。

```cpp
AEGP_SetLayerOffset(
 AEGP_LayerH layerH,
 A_Time *offsetPT);
```

#### AEGP_GetLayerStretch

获取图层的拉伸系数。

```cpp
AEGP_GetLayerStretch(
 AEGP_LayerH layerH,
 A_Ratio *stretchPRt);
```

#### AEGP_SetLayerStretch

设置图层的拉伸系数。

```cpp
AEGP_SetLayerStretch(
 AEGP_LayerH layerH,
 A_Ratio *stretchPRt);
```

#### AEGP_GetLayerTransferMode

获取图层的混合模式。

```cpp
AEGP_GetLayerTransferMode(
 AEGP_LayerH layerH,
 AEGP_LayerTransferMode *modeP);
```

#### AEGP_SetLayerTransferMode

设置图层的混合模式。可撤消。

```cpp
AEGPSetLayerTransferMode(
 AEGP_LayerH layerH,
 AEGP_LayerTransferMode *modeP);
```

从 6.5 版本开始，当你把一个图层变成轨道哑光时，它前面的图层将被禁用。
就像你通过界面做这件事一样。

#### AEGP_IsAddLayerValid

测试当前是否可以将一个给定的项目添加到一个合成中。一个合成不能被添加到它自己或它所包含的任何合成中；其他条件也可以排除成功添加。不先使用这个函数而添加一个层，将产生未定义的结果。

```cpp
AEGP_IsAddLayerValid(
 AEGP_ItemH item_to_addH,
 AEGP_CompH into_compH,
 A_Boolean *validPB);
```

#### AEGP_AddLayer

添加一个项目到合成中，高于所有其他层。可以撤销。首先使用`AEGP_IsAddLayerValid()`，以确认它是可行的。

```cpp
AEGP_AddLayer(
 AEGP_ItemH item_to_addH,
 AEGP_CompH into_compH,
 A_Boolean *added_layerPH0);
```

#### AEGP_ReorderLayer

改变层的顺序。可撤销。

```cpp
AEGP_ReorderLayer(
 AEGP_LayerH layerH,
 A_long layer_indexL);
```

要将一个图层添加到合成的最后，要使用`layer_indexL = AEGP_REORDER_LAYER_TO_END`。

#### AEGP_GetLayerMaskedBounds

给定一个图层的句柄和时间，返回应用了遮罩的可见区域的边界。

```cpp
AEGP_GetLayerMaskedBounds(
 AEGP_LayerH layerH,
 const A_Time *comp_timePT,
 A_FloatRect *boundsPR);
```

#### AEGP_GetLayerObjectType

返回一个图层的对象类型。

```cpp
AEGP_GetLayerObjectType(
 AEGP_LayerH layerH,
 AEGP_ObjectType *object_type);
```

- `AEGP_ObjectType_AV`
- `AEGP_ObjectType_LIGHT`,
- `AEGP_ObjectType_CAMERA`,
- `AEGP_ObjectType_TEXT`

#### AEGP_IsLayer3D

录像项目是否为 3D 层。所有 AV 层都是 2D 或 3D。

```cpp
AEGP_IsLayer3D(
 AEGP_LayerH layerH,
 A_Boolean *is_3DPB);
```

#### AEGP_IsLayer2D

录像项目是否为 2D 层。所有的 AV 层都是 2D 或 3D 的。

```cpp
AEGP_IsLayer2D(
 AEGP_LayerH layerH,
 A_Boolean *is_2DPB);
```

#### AEGP_IsVideoActive

给出组成时间和一个图层，看该图层是否会渲染。时间模式是`AEGP_LTimeMode_LayerTime`或`AEGP_LTimeMode_CompTime`。

```cpp
AEGP_IsVideoActive(
 AEGP_LayerH layerH,
 AEGP_LTimeMode time_mode,
 A_Time *comp_timePT,
 A_Boolean *is_activePB);
```

#### AEGP_IsLayerUsedAsTrackMatte

该层是否作为轨道消光使用？

```cpp
AEGP_IsLayerUsedAsTrackMatte(
 AEGP_LayerH layerH,
 A_Boolean fill_must_be_activeB,
 A_Boolean *is_track_mattePB);
```

#### AEGP_DoesLayerHaveTrackMatte

这个图层是否有轨迹哑光？

```cpp
AEGP_DoesLayerHaveTrackMatte(
 AEGP_LayerH layerH,
 A_Boolean *has_track_mattePB);
```

#### AEGP_ConvertCompToLayerTime

给出合成空间中的时间，返回相对于图层源片段的时间。

```cpp
AEGP_ConvertCompToLayerTime(
 AEGP_LayerH layerH,
 const A_Time *comp_timeP,
 A_Time *layer_timeP);
```

#### AEGP_ConvertLayerToCompTime

给出一个图层空间的时间，找到组成空间的相应时间。

```cpp
AEGP_ConvertLayerToCompTime(
 AEGP_LayerH layerH,
 const A_Time *layer_timePT,
 A_Time *comp_timePT);
```

#### AEGP_GetLayerDancingRandValue

由舞蹈溶解转移函数使用。

```cpp
AEGP_GetLayerDancingRandValue(
 AEGP_LayerH layerH,
 const A_Time *comp_timePT,
 A_long *rand_valuePL);
```

#### AEGP_GetLayerID

提供该层的唯一 ID。这个 ID 在项目的生命周期内不会改变。

```cpp
AEGP_GetLayerID(
 AEGP_LayerH layerH,
 AEGP_LayerIDVal *id_valP);
```

#### AEGP_GetLayerToWorldXform

给定一个图层句柄和时间，返回图层到世界的转换矩阵。

```cpp
AEGP_GetLayerToWorldXform(
 AEGP_LayerH aegp_layerH,
 const A_Time *comp_timeP,
 A_Matrix4 *transform);
```

#### AEGP_GetLayerToWorldXformFromView

给出一个图层句柄、当前(组成)时间和请求的视图时间。
返回用户的视图和图层之间的转换，并根据合成的当前长宽比进行校正。

```cpp
AEGP_GetLayerToWorldXformFromView(
 AEGP_LayerH aegp_layerH,
 const A_Time *view_timeP,
 const A_Time *comp_timeP,
 A_Matrix4 *transform);
```

#### AEGP_SetLayerName

设置一个图层的名称。new_nameZ 指向一个空尾的 UTF-16 字符串。

```cpp
AEGP_SetLayerName(
 AEGP_LayerH aegp_layerH,
 const A_UTF16Char *new_nameZ);
```

#### AEGP_GetLayerParent

检索图层的父层句柄(如果没有父层，则没有)。

```cpp
AEGP_GetLayerParent(
 AEGP_LayerH layerH,
 AEGP_LayerH *parent_layerPH);
```

#### AEGP_SetLayerParent

设置一个图层的父图层。

```cpp
AEGP_SetLayerParent(
 AEGP_LayerH layerH,
 const AEGP_LayerH parent_layerH);
```

#### AEGP_DeleteLayer

删除一个图层。你能相信我们花了三个套件的时间来添加一个删除函数吗？我们也不能。

```cpp
AEGP_DeleteLayer(
 AEGP_LayerH layerH);
```

#### AEGP_DuplicateLayer

复制图层。可以撤销。

```cpp
AEGP_DuplicateLayer(
 AEGP_LayerH orig_layerH,
 AEGP_LayerH *dupe_layerPH);
```

#### AEGP_GetLayerFromLayerID

检索与给定的`AEGP_LayerIDVal`相关的`AEGP_LayerH`(这是你在访问效果的层参数流时得到的)。

```cpp
AEGP_GetLayerFromLayerID(
 AEGP_CompH parent_compH,
 AEGP_LayerIDVal id,
 AEGP_LayerH *layerPH);
```

#### AEGP_GetLayerLabel

获取一个图层的`AEGP_LabelID`。

```cpp
AEGP_GetLayerLabel(
 AEGP_LayerH layerH,
 AEGP_LabelID *labelP);
```

#### AEGP_SetLayerLabel

设置一个图层的`AEGP_LabelID`。可撤销。

```cpp
AEGP_SetLayerLabel(
 AEGP_LayerH layerH,
 AEGP_LabelID label);
```

#### AEGP_GetLayerSamplingQuality

CC 中的新功能。获取一个图层的采样质量。

```cpp
AEGP_GetLayerSamplingQuality(
 AEGP_LayerH layerH,
 AEGP_LayerSamplingQuality *label);
```

图层采样质量是以下标志之一。

- `AEGP_LayerSamplingQual_BILINEAR`
- `AEGP_LayerSamplingQual_BICUBIC`

#### AEGP_SetLayerSamplingQuality

在 CC 中新出现。设置一个图层的采样质量(见上面的标志值)。
该选项在图层上被明确设置，与图层质量无关。
如果你想强制开启它，你必须用`AEGP_SetLayerQuality`将图层质量设置为`AEGP_LayerQual_BEST`。
否则，只要图层质量被设置为`AEGP_LayerQual_BEST`，它就只会使用指定的图层采样质量。
可撤销。

```cpp
AEGP_SetLayerSamplingQuality(
 AEGP_LayerH layerH,
 AEGP_LayerSamplingQuality label);
```

## Layer Creation Notes

所有使用 AEGP 调用创建的图层将从合成时间 0 开始，并具有合成的持续时间。

使用[AEGP_LayerSuite8]的`AEGP_SetLayerOffset()`和`AEGP_SetLayerInPointAndDuration()`来正确设置层的时间信息。

当图层拉伸系数(使用[AEGP_LayerSuite8]中的`AEGP_GetLayerStretch`获得，自然)不是 100%时，将需要进行以下计算以产生正确的图层偏移。

```cpp
offset = compIn - stretch * layerIn;
```

## Communication With A Layer’s Effects

访问应用于一个层的效果。这个套件提供对所有参数数据流的访问。

使用[流套件]来处理这些流。

一个`AEGP_Effect_RefH`是对一个应用效果的引用。一个`AEGP_InstalledEffectKey`是对一个已安装的效果的引用，它可能是也可能不是当前应用于任何东西。

如果 Foobarocity 在一个图层上应用了两次，会有两个不同的`AEGP_Effect_RefH`，但它们都会返回同一个`AEGP_InstalledEffectKey`。

### AEGP_EffectSuite4

#### AEGP_GetLayerNumEffects

获取应用于一个图层的效果数量。

```cpp
AEGP_GetLayerNumEffects(
 AEGP_LayerH layerH,
 A_long *num_effectsPL);
```

#### AEGP_GetLayerEffectByIndex

检索(通过索引)应用到该层的效果的引用。

```cpp
AEGP_GetLayerEffectByIndex(
 AEGP_PluginID aegp_plugin_id,
 AEGP_LayerH layerH,
 AEGP_EffectIndex effect_indexL,
 AEGP_EffectRefH *effectPH);
```

#### AEGP_GetInstalledKeyFromLayerEffect

给定一个`AEGP_EffectRefH`，检索其相关的`AEGP_InstalledEffectKey`。

```cpp
AEGP_GetInstalledKeyFromLayerEffect(
 AEGP_EffectRefH effect_refH,
 AEGP_InstalledEffectKey *installed_keyP);
```

#### AEGP_GetEffectParamUnionByIndex

返回效果参数的描述。不要使用这个函数返回的 ParamDef 中的值(使用`AEGP_GetNewStreamValue()`代替)。
提供这个函数是为了让 AEGP 能够访问参数默认值、复选框名称和弹出式字符串。
使用[AEGP_StreamSuite5]的`AEGP_GetEffectNumParamStreams()`来获取流的数量。
这对确定最大的 "param_index "很有用。最后一个参数是可选的。

```cpp
AEGP_GetEffectParamUnionByIndex(
 AEGP_PluginID aegp_plugin_id,
 AEGP_EffectRefH effectH,
 PF_ParamIndex param_index,
 PF_ParamType *param_typeP
 PF_ParamDefUnion *uP0);
```

#### AEGP_GetEffectFlags

获取给定的`AEGP_EffectRefH`的标志。

```cpp
AEGP_GetEffectFlags(
 AEGP_EffectRefH effect_refH,
 AEGP_EffectFlags *effect_flagsP);
```

标志将是以下内容的合成。

- `AEGP_EffectFlags_NONE`
- `AEGP_EffectFlags_ACTIVE`
- `AEGP_EffectFlags_AUDIO_ONLY`
- `AEGP_EffectFlags_AUDIO_TOO`
- `AEGP_EffectFlags_MISSING`

#### AEGP_SetEffectFlags

为给定的`AEGP_EffectRefH`设置标志(上面列举的)，用一组不同的效果标志来掩盖。

```cpp
AEGP_SetEffectFlags(
 AEGP_EffectRefH effect_refH,
 AEGP_EffectFlags mask,
 AEGP_EffectFlags effect_flags);
```

#### AEGP_ReorderEffect

改变应用效果的顺序(传递要求的索引)。

```cpp
AEGP_ReorderEffect(
 AEGP_EffectRefH effect_refH,
 A_long effect_indexL);
```

#### AEGP_EffectCallGeneric

调用一个效果插件，并传递给它一个指向任何你喜欢的数据的指针；效果可以修改它。这就是 AEGP 与效果的通信方式。
传递`PF_Cmd_COMPLETELY_GENERAL`给`effect_cmd`。

```cpp
AEGP_EffectCallGeneric(
 AEGP_PluginID aegp_plugin_id,
 AEGP_EffectRefH effectH,
 const A_Time *timePT,
 PF_Cmd effect_cmd,
 void *extraPV);
```

#### AEGP_DisposeEffect

处置一个 AEGP_EffectRefH。使用它来处理由 After Effects 返回的任何 AEGP_EffectRefH。

```cpp
AEGP_DisposeEffect(
 AEGP_EffectRefH effectH);
```

#### AEGP_ApplyEffect

在给定的图层上应用一个效果。返回新创建的`AEGP_EffectRefH`。

```cpp
AEGP_ApplyEffect(
 AEGP_PluginID aegp_plugin_id,
 AEGP_LayerH layerH,
 AEGP_InstalledEffectKey installed_key,
 AEGP_EffectRefH *effect_refPH);
```

#### AEGP_DeleteLayerEffect

删除一个已应用的效果。

```cpp
AEGP_DeleteLayerEffect(
 AEGP_EffectRefH effect_refH);
```

#### AEGP_GetNumInstalledEffects

返回 After Effects 中安装的效果的数量。

```cpp
AEGP_GetNumInstalledEffects(
 A_long *num_installed_effectsPL);
```

#### AEGP_GetNextInstalledEffect

返回下一个安装的效果的`AEGP_InstalledEffectKey`。将`AEGP_InstalledEffectKey_NONE`作为第一个参数，获得第一个`AEGP_InstalledEffectKey`。

```cpp
AEGP_GetNextInstalledEffect(
 AEGP_InstalledEffectKey key,
 AEGP_InstalledEffectKey *next_keyPH);
```

#### AEGP_GetEffectName

获取效果的名称。`nameZ`的长度可以达到`AEGP_MAX_EFFECT_NAME_SIZE + 1`。

```cpp
AEGP_GetEffectName(
 AEGP_InstalledEffectKey installed_key,
 A_char *nameZ);
```

注意：使用`AEGP_SetStreamName`来改变一个效果的显示名称。

#### AEGP_GetEffectMatchName

获取效果的匹配名称(在 PiPL 中定义)。`match_nameZ`最长为`AEGP_MAX_EFFECT_MATCH_NAME_SIZE + 1`长。

```cpp
AEGP_GetEffectMatchName(
 AEGP_InstalledEffectKey installed_key,
 A_char *match_nameZ);
```

匹配名称为 7 位 ASCII。UI 名称是当前应用程序运行时的编码；例如，对于 Windows 上的大多数语言，ISO 8859-1。

#### AEGP_GetEffectCategory

效果的菜单类别。`categoryZ`的长度可以达到`AEGP_MAX_EFFECT_CATEGORY_NAME_SIZE + 1`。

```cpp
AEGP_GetEffectCategory(
 AEGP_InstalledEffectKey installed_key,
 A_char *categoryZ);
```

#### AEGP_DuplicateEffect

复制一个给定的`AEGP_EffectRefH`。调用者必须在完成后处理掉重复的内容。

```cpp
AEGP_DuplicateEffect(
 AEGP_EffectRefH orig_effect_refH,
 AEGP_EffectRefH *dupe_refPH);
```

#### AEGP_NumEffectMask

CC2014 中的新内容。这个效果上有多少个遮罩？

```cpp
AEGP_NumEffectMask(
 AEGP_EffectRefH effect_refH,
 A_u_long *num_masksPL);
```

#### AEGP_GetEffectMaskID

CC2014 中的新功能。对于一个给定的遮罩indexL，返回相应的`AEGP_MaskIDVal`，用于唯一识别遮罩。

```cpp
AEGP_GetEffectMaskID(
 AEGP_EffectRefH effect_refH,
 A_u_long mask_indexL,
 AEGP_MaskIDVal *id_valP);
```

#### AEGP_AddEffectMask

CC2014 中的新功能。添加一个效果掩码，可以使用[掩码管理]创建。
返回效果反射的本地流--如果你想添加关键帧，则很有用。调用者必须在完成后处理`AEGP_StreamRefH`。
可撤销。

```cpp
AEGP_AddEffectMask(
 AEGP_EffectRefH effect_refH,
 AEGP_MaskIDVal id_val,
 AEGP_StreamRefH streamPH0);
```

#### AEGP_RemoveEffectMask

CC 2014 的新功能。移除一个效果遮罩。
可撤销。

```cpp
AEGP_RemoveEffectMask(
 AEGP_EffectRefH effect_refH,
 AEGP_MaskIDVal id_val);
```

#### AEGP_SetEffectMask

在 CC 2014 中新增。在一个现有的索引上设置一个效果遮罩。
返回效果参考的本地流 - 如果你想添加关键帧，很有用。
完成后调用者必须处理`AEGP_StreamRefH`。
可撤销。

```cpp
AEGP_SetEffectMask(
 AEGP_EffectRefH effect_refH,
 A_u_long mask_indexL,
 AEGP_MaskIDVal id_val,
 AEGP_StreamRefH *streamPH0);
```

## Exploiting Effect UI Behavior To Look Cool

即使你操作了一个图层的效果，它的效果控制也不一定会变得可见。

然而，如果你应用然后立即删除一个效果，该层的效果控制将变得可见。

很狡猾，是吗？

## StreamRefs And EffectRefs

你如何获得一个效果的 AEGP_StreamRef？首先，通过调用`AEGP_GetNewEffectForEffect()`获得效果的`AEGP_EffectRef`。

然后调用`AEGP_GetNewEffectStreamByIndex()`，比如说针对参数索引 1，这时会传回一个参数流。

然后调用`AEGP_GetNewParentStreamRef()`--瞧，你的`AEGP_StreamRef`先生!

如果你获得了对效果流的引用，在你完成对流的处理之前，不要处置`AEGP_EffectRefH`，否则你会使 After Effects 的检出机制不平衡。还要记住，AEGP_StreamRefHs 是不透明的；而`AEGP_StreamValue2`s 则不是(完全)。

要获得一个效果的实例名称(由用户重新命名)，需要获得该效果本身的 AEGP_StreamRef，然后调用`AEGP_GetStreamName`。

## Diving Into Streams

在 After Effects 中，几乎所有的东西都是一个流。效果参数、图层、遮罩和形状都在内部由流表示。AEGP API 几乎可以访问每个流的各个方面。

After Effects 的时间轴可以包含许多对象类型；每个对象都支持一组称为流的参数。所有的流，无论连接到哪种类型的对象，在概念上都是相似的(After Effects 的处理方式也是相似的。但由于其包含的内容不同，你访问每种类型的流的方式也不同。

一个流，一旦获得，代表一个可能随时间变化的值。不是所有的流都能随时间变化，一个特定的流在访问时可能不是随时间变化的。

有两种方法来访问一个流的值。如果流有关键帧，你可以使用[Working With Keyframes](#aegps-aegp-suites-keyframe-suite)。提供的值不会反映表达式的影响。注意：在任何表达式中，当前关键帧的值总是可以作为变量值。

你也可以使用来自[AEGP_StreamSuite5]的`AEGP_GetNewStreamValue`，它在特定时间对流的值进行采样。对于没有表达式或关键帧的流，时间参数是没有意义的，该函数返回的基本上是流的常量值。使用`AEGP_SetStreamValue`(它不以时间为参数)来设置这些流。

许多 StreamSuite 函数会填充一个 StreamH，完成后你的 AEGP 必须处理掉它。After Effects 分配并传递给你一个值的副本，而不是一个直接指向原始值的句柄。`AEGP_GetNewLayerStream()`仅限于那些不需要分配内存就可以访问其值的流。

## Okay, What Did I Just Get?

一个流值是一个大的联合体，其中只有一个结构(取决于流的类型)被填充。注意与[PF_ParamDef](../effect-basics/PF_ParamDef.html#effect-basics-pf-paramdef)的相似性。

```cpp
typedef union {
 AEGP_FourDVal four_d;
 AEGP_ThreeDVal three_d;
 AEGP_TwoDVal two_d;
 AEGP_OneDVal one_d;
 AEGP_ColorVal color;
 AEGP_ArbBlockVal arbH;
 AEGP_MarkerValP markerP;
 AEGP_LayerIDVal layer_id;
 AEGP_MaskIDVal mask_id;
 AEGP_MaskOutlineValH mask;
 AEGP_TextDocumentH text_documentH;
} AEGP_StreamVal2;
```

## Layers

`AEGP_GetLayerStreamValue`用于访问参数，如锚点和位置，几乎是 AE 中所有层的原生参数。

使用`IsStreamLegal`可以让你确定该层是否提供该流类型。

## Masks

由于一个图层可以有多个遮罩，使用`AEGP_GetLayerMaskByIndex`从[AEGP_MaskSuite6](#aegps-aegp-suites-aegp-masksuite)访问遮罩。

面具不像层那样有流；它们有自己的枚举。使用[AEGP_StreamSuite5]中的`AEGP_GetNewMaskStream`访问它们的流。

## Effects

它们可以有数量不等的流/参数，而且在编写 AEGP 时，它们的顺序和定义并不清楚。

因此，我们不能提供一个枚举来选择它们，而是必须通过索引来获得它们，因此，从[AEGP_StreamSuite5]中的`GetNewEffectStreamByIndex`。

## Stream Suite

访问和操作一个图层的流的值。对于绘画和文本流，使用[动态流]代替。

### AEGP_StreamSuite5

#### AEGP_IsStreamLegal

确定给定的流是否适合给定的层。

```cpp
AEGP_IsStreamLegal(
 AEGP_LayerH layerH,
 AEGP_LayerStream which_stream,
 A_Boolean* is_legalP);
```

#### AEGP_CanVaryOverTime

给定一个流，返回一个流是否是时间变化的(并且可以是关键帧)。

```cpp
AEGP_CanVaryOverTime(
 AEGP_StreamRefH streamH,
 A_Boolean *can_varyPB);
```

#### AEGP_GetValidInterpolations

检索一个`AEGP_KeyInterpolationMask`，表明哪些插值类型对`AEGP_StreamRefH`有效。

```cpp
AEGP_GetValidInterpolations(
 AEGP_StreamRefH streamH,
 AEGP_KeyInterpolationMask *valid_interpP);
```

`AEGP_KeyInterpolationMask`将是以下的合成。

- `AEGP_KeyInterpMask_NONE`
- `AEGP_KeyInterpMask_LINEAR`
- `AEGP_KeyInterpMask_BEZIER`
- `AEGP_KeyInterpMask_HOLD`
- `AEGP_KeyInterpMask_CUSTOM`
- `AEGP_KeyInterpMask_ANY`

#### AEGP_GetNewLayerStream

获取一个图层的数据流。插件必须处理掉`streamPH`。注意，这将不提供关键帧的访问。
使用[AEGP_KeyframeSuite]代替。

```cpp
AEGP_GetNewLayerStream(
 AEGP_PluginID id,
 AEGP_LayerH layerH,
 AEGP_LayerStream which_stream,
 AEGP_StreamRefH *streamPH);
```

- `AEGP_LayerStream_ANCHORPOINT`,
- `AEGP_LayerStream_POSITION`,
- `AEGP_LayerStream_SCALE`,
- `AEGP_LayerStream_ROTATION`,
- `AEGP_LayerStream_ROTATE_Z`,
- `AEGP_LayerStream_OPACITY`,
- `AEGP_LayerStream_AUDIO`,
- `AEGP_LayerStream_MARKER`,
- `AEGP_LayerStream_TIME_REMAP`,
- `AEGP_LayerStream_ROTATE_X`,
- `AEGP_LayerStream_ROTATE_Y`,
- `AEGP_LayerStream_ORIENTATION`

只对`AEGP_ObjectType_CAMERA`有效。

- `AEGP_ObjectType_CAMERA`
- `AEGP_LayerStream_ZOOM`,
- `AEGP_LayerStream_DEPTH_OF_FIELD`,
- `AEGP_LayerStream_FOCUS_DISTANCE`,
- `AEGP_LayerStream_APERTURE`,
- `AEGP_LayerStream_BLUR_LEVEL`,
- `AEGP_LayerStream_IRIS_SHAPE`,
- `AEGP_LayerStream_IRIS_ROTATION`,
- `AEGP_LayerStream_IRIS_ROUNDNESS`,
- `AEGP_LayerStream_IRIS_ASPECT_RATIO`,
- `AEGP_LayerStream_IRIS_DIFFRACTION_FRINGE`,
- `AEGP_LayerStream_IRIS_HIGHLIGHT_GAIN`,
- `AEGP_LayerStream_IRIS_HIGHLIGHT_THRESHOLD`,
- `AEGP_LayerStream_IRIS_HIGHLIGHT_SATURATION`,

只对 "AEGP_ObjectType_LIGHT "有效。

- `AEGP_LayerStream_INTENSITY`,
- `AEGP_LayerStream_COLOR`,
- `AEGP_LayerStream_CONE_ANGLE`,
- `AEGP_LayerStream_CONE_FEATHER`,
- `AEGP_LayerStream_SHADOW_DARKNESS`,
- `AEGP_LayerStream_SHADOW_DIFFUSION`,
- `AEGP_LayerStream_LIGHT_FALLOFF_TYPE`,
- `AEGP_LayerStream_LIGHT_FALLOFF_START`,
- `AEGP_LayerStream_LIGHT_FALLOFF_DISTANCE`,

只对 "AEGP_ObjectType_AV "有效。

- `AEGP_LayerStream_ACCEPTS_SHADOWS`,
- `AEGP_LayerStream_ACCEPTS_LIGHTS`,
- `AEGP_LayerStream_AMBIENT_COEFF`,
- `AEGP_LayerStream_DIFFUSE_COEFF`,
- `AEGP_LayerStream_SPECULAR_INTENSITY`,
- `AEGP_LayerStream_SPECULAR_SHININESS`,
- `AEGP_LayerStream_METAL`,
- `AEGP_LayerStream_LIGHT_TRANSMISSION`,

只对`AEGP_ObjectType_AV`有效，CS6 新增：\* `AEGP_LayerStream_REFLECTION_INTENSITY`。

- `AEGP_LayerStream_REFLECTION_SHARPNESS`,
- `AEGP_LayerStream_REFLECTION_ROLLOFF`,
- `AEGP_LayerStream_TRANSPARENCY_COEFF`,
- `AEGP_LayerStream_TRANSPARENCY_ROLLOFF`,
- `AEGP_LayerStream_INDEX_OF_REFRACTION`,
- `AEGP_LayerStream_EXTRUSION_BEVEL_STYLE`,
- `AEGP_LayerStream_EXTRUSION_BEVEL_DIRECTION`,
- `AEGP_LayerStream_EXTRUSION_BEVEL_DEPTH`,
- `AEGP_LayerStream_EXTRUSION_HOLE_BEVEL_DEPTH`,
- `AEGP_LayerStream_EXTRUSION_DEPTH`,
- `AEGP_LayerStream_PLANE_CURVATURE`,
- `AEGP_LayerStream_PLANE_SUBDIVISION`,

只对`LIGHT`和`AV`有效。

- `AEGP_LayerStream_CASTS_SHADOWS`,
- `AEGP_LayerStream_SOURCE_TEXT`
- `AEGP_LayerStream_BEGIN` = `AEGP_LayerStream_ANCHORPOINT`,
- `AEGP_LayerStream_END` = `AEGP_LayerStream_LIGHT_FALLOFF_DISTANCE + 1`

```cpp
enum {
 AEGP_LightFalloff_NONE = 0,
 AEGP_LightFalloff_SMOOTH,
 AEGP_LightFalloff_INVERSE_SQUARE_CLAMPED
};

typedef A_u_long AEGP_LightFalloffType;
```

#### AEGP_GetEffectNumParamStreams

获取与一个效果相关的参数流的数量。

```cpp
AEGP_GetEffectNumParamStreams(
 AEGP_EffectRefH effect_refH,
 A_long *num_parmsPL);
```

#### AEGP_GetNewEffectStreamByIndex

获取一个效果的参数流。插件必须处理`streamPH`。

```cpp
AEGP_GetNewEffectStreamByIndex(
 AEGP_PluginID id,
 AEGP_EffectRefH effect_refH,
 PF_ParamIndex param_index,
 AEGP_StreamRefH *streamPH);
```

#### AEGP_GetNewMaskStream

获取一个面具的流。该流必须被处理掉。
也可以参考[AEGP_MaskSuite]和[AEGP_MaskOutlineSuite]中的其他 Mask 函数\* `AEGP_MaskStream_OUTLINE`。

- `AEGP_MaskStream_OPACITY`,
- `AEGP_MaskStream_FEATHER`,
- `AEGP_MaskStream_EXPANSION`,

对迭代有用。

- `AEGP_MaskStream_BEGIN` = `AEGP_MaskStream_OUTLINE`,
- `AEGP_MaskStream_END` = `AEGP_MaskStream_EXPANSION + 1`

```cpp
AEGP_GetNewMaskStream(
 AEGP_PluginID aegp_plugin_id,
 AEGP_MaskRefH mask_refH,
 AEGP_MaskStream which_stream,
 AEGP_StreamRefH *mask_strmPH);
```

#### AEGP_DisposeStream

处理一个流(对这些函数传递给插件的所有流进行处理)。

```cpp
AEGP_DisposeStream(
 AEGP_StreamRefH streamH);
```

#### AEGP_GetNewMaskOpacity

获取面具的不透明度流。该流必须被处理掉。

```cpp
AEGP_GetNewMaskOpacity(
 AEGP_PluginID aegp_plugin_id,
 AEGP_MaskH maskH,
 PF_ParamIndex param_index,
 AEGP_StreamRefH *mask_opacity_streamPH);
```

#### AEGP_GetStreamName

获取流的名称(本地化或强制英文)。 是`A_UTF16Char`的句柄(包含空尾的 UTF16 字符串)。
必须用`AEGP_FreeMemHandle`处置。

```cpp
AEGP_GetStreamName(
 AEGP_PluginID pluginID,
 AEGP_StreamRefH streamH,
 A_Boolean force_englishB,
 AEGP_MemHandle *utf_stream_namePH);
```

注意：如果`force_englishB`为 true，默认的名字将覆盖任何已经完成的流重命名(无论是程序上的，还是用户的)。

#### AEGP_GetStreamUnitsText

获取流的单位，格式为文本(本地化或强制英语)；`unitsZ`最长为`AEGP_MAX_STREAM_NAME_LEN + 1`。

```cpp
AEGP_GetStreamUnitsText(
 AEGP_StreamRefH streamH,
 A_Boolean force_englishB,
 A_char *unitsZ);
```

#### AEGP_GetStreamProperties

获取流的标志，以及最小和最大值(浮动)，如果流*有*分钟和最大值。
StreamFlags 值。

- `AEGP_StreamFlag_NONE`
- `AEGP_StreamFlag_HAS_MIN`
- `AEGP_StreamFlag_HAS_MAX`

```cpp
AEGP_GetStreamProperties(
 AEGP_StreamRefH streamH,
 AEGP_StreamFlags *flagsP,
 A_FpLong *minP0,
 A_FpLong *maxP0);
```

#### AEGP_IsStreamTimevarying

返回流是否受到表达式的影响。

```cpp
AEGP_IsStreamTimevarying(
 AEGP_StreamRefH streamH,
 A_Boolean *is_timevaryPB);
```

#### AEGP_GetStreamType

获取流的类型(尺寸)。

```cpp
AEGP_GetStreamType(
 AEGP_StreamRefH streamH,
 AEGP_StreamType *stream_typeP);
```

- `AEGP_StreamType_NO_DATA`,
- `AEGP_StreamType_TwoD_SPATIAL`,
- `AEGP_StreamType_TwoD`,
- `AEGP_StreamType_ThreeD`,
- `AEGP_StreamType_ThreeD_SPATIAL`,
- `AEGP_StreamType_OneD`,
- `AEGP_StreamType_COLOR`,
- `AEGP_StreamType_ARB`,
- `AEGP_StreamType_MARKER`,
- `AEGP_StreamType_LAYER_ID`,
- `AEGP_StreamType_MASK_ID`,
- `AEGP_StreamType_MASK`,
- `AEGP_StreamType_TEXT_DOCUMENT`

注意：总是返回`ThreeD_Spatial`的位置，不管该层是否是 3D 的。

#### AEGP_GetNewStreamValue

在你指定的时间，获取流的值。`valueP`必须由插件处理。
`AEGP_LTimeMode`表示时间是合成时间还是层时间。

```cpp
AEGP_GetNewStreamValue(
 AEGP_PluginID aegp_plugin_id,
 AEGP_StreamRefH streamH,
 AEGP_LTimeMode time_mode,
 const A_Time *timePT,
 A_Boolean pre_exprB,
 AEGP_StreamValue2 *valueP);
```

#### AEGP_DisposeStreamValue

处置流值。始终将传递给插件的值去掉。

```cpp
AEGP_DisposeStreamValue(
 AEGP_StreamValue2 *valueP);
```

#### AEGP_SetStreamValue

只有当流不是时间变化的时候才合法。

```cpp
AEGP_SetStreamValue(
 AEGP_PluginID aegp_plugin_id,
 AEGP_StreamRefH streamH,
 AEGP_StreamValue2 *valueP);
```

#### AEGP_GetLayerStreamValue

注意：这个便利函数只对具有原始数据类型的流有效，而对`AEGP_ArbBlockVal`、`AEGP_MarkerValH`或`AEGP_MaskOutlineValH`无效。
对于这些和其他复杂的类型，使用上面描述的`AEGP_GetNewStreamValue`。

```cpp
AEGP_GetLayerStreamValue(
 AEGP_LayerH layerH,
 AEGP_LayerStream which_stream,
 AEGP_LTimeMode time_mode,
 const A_Time *timePT,
 A_Boolean pre_expB,
 AEGP_StreamVal *stream_valP,
 AEGP_StreamType *strm_typeP0);
```

#### AEGP_GetExpressionState

确定是否在给定的`AEGP_StreamRefH`上启用表达式。

```cpp
AEGP_GetExpressionState(
 AEGP_PluginID aegp_plugin_id,
 AEGP_StreamRefH streamH,
 A_Boolean *enabledPB);
```

#### AEGP_SetExpressionState

在给定的`AEGP_StreamRefH`上启用和禁用表达式。

```cpp
AEGP_SetExpressionState(
 AEGP_PluginID aegp_plugin_id,
 AEGP_StreamRefH streamH,
 A_Boolean enabledB);
```

#### AEGP_GetExpression

获取表达式的文本。从 suite 版本 5 开始(在 15.0 及更高版本中可用)，它现在支持 Unicode。

```cpp
AEGP_GetExpression(
 AEGP_PluginID aegp_plugin_id,
 AEGP_StreamRefH streamH,
 AEGP_MemHandle *unicodeHZ);
```

#### AEGP_SetExpression

设置表达式的文本。从第 5 版开始(在 15.0 及以后版本中可用)，现在支持 Unicode。

```cpp
AEGP_SetExpression(
 AEGP_PluginID aegp_plugin_id,
 AEGP_StreamRefH streamH,
 const A_UTF16Char* expressionP);
```

#### AEGP_DuplicateStreamRef

复制一个给定的`AEGP_StreamRefH`。处置重复的内容。

```cpp
AEGP_DuplicateStreamRef(
 AEGP_PluginID aegp_plugin_id,
 AEGP_StreamRefH streamH,
 AEGP_StreamRefH *dup_streamPH);
```

## Dynamic Streams

`AEGP_DynamicStreamSuite`访问和操作绘画和文本流。

使用`AEGP_GetStreamGroupingType`和`AEGP_GetDynamicStreamFlags`来识别流，然后再尝试使用只对某些流类型工作的函数。

还要注意，通常情况下，你可以简单地使用[Stream Suite]调用来处理动态流。另一方面，只有那些专门针对动态流的函数才在这个套件中。

### AEGP_DynamicStreamSuite4

#### AEGP_GetNewStreamRefForLayer

检索与该层对应的 AEGP_StreamRefH。这个函数被用来启动该层流的递归行走。

```cpp
AEGP_GetNewStreamRefForLayer(
 AEGP_PluginID aegp_plugin_id,
 AEGP_LayerH layerH,
 AEGP_StreamRefH *streamPH);
```

#### AEGP_GetNewStreamRefForMask

检索与掩码对应的 AEGP_StreamRefH。

```cpp
AEGP_GetNewStreamRefForMask(
 AEGP_PluginID aegp_plugin_id,
 AEGP_MaskRefH maskH,
 AEGP_StreamRefH *streamPH);
```

#### AEGP_GetStreamDepth

检索与给定的`AEGP_StreamRefH`相关的子流的数量。
初始层的深度为 0。

```cpp
AEGP_GetStreamDepth(
 AEGP_StreamRefH streamH,
 A_long *depthPL);
```

#### AEGP_GetStreamGroupingType

检索给定的`AEGP_StreamRefH`的分组类型。

```cpp
AEGP_GetStreamGroupingType(
 AEGP_StreamRefH streamH,
 AEGP_StreamGroupingType *group_typeP);
```

AEGP_StreamGroupingType 将是以下之一。

- `AEGP_StreamGroupingType_NONE`
- `AEGP_StreamGroupingType_LEAF`
- `AEGP_StreamGroupingType_NAMED_GROUP`
- `AEGP_StreamGroupingType_INDEXED_GROUP`

#### AEGP_GetNumStreamsInGroup

检索与给定的`AEGP_StreamRefH`相关的流的数量。 如果用类型为`AEGP_StreamGroupingType_LEAF`的`AEGP_StreamRefH`调用此函数将返回一个错误。

```cpp
AEGP_GetNumStreamsInGroup(
 AEGP_StreamRefH streamH,
 A_long *num_streamsPL);
```

#### AEGP_GetDynamicStreamFlags

检索一个给定的 AEGP_StreamRefH 的标志。

```cpp
AEGP_GetDynamicStreamFlags(
 AEGP_StreamRefH streamH,
 AEGP_DynStreamFlags *flagsP);
```

`AEGP_DynStreamFlags`将是以下之一：\* `AEGP_DynStreamFlag_ACTIVE_EYEBALL`意味着流可用于读写。

- `AEGP_DynStreamFlag_HIDDEN` means that, while the stream is still readable/writable, it may not currently be visible in the UI.
- `AEGP_DynStreamFlag_DISABLED` A read-only flag. Indicates whether the `AEGP_StreamRefH` is grayed out in the UI.
  注意，从 CS5 开始，如果一个参数被禁用，这个标志将不会被返回。
  相反，请检查[Parameter UI Flags](../effect-basics/PF_ParamDef.html#effect-basics-pf-paramdef-parameter-ui-flags)中的`PF_PUI_DISABLED`。
- `AEGP_DynStreamFlag_ELIDED` A read-only flag. Indicates that the `AEGP_StreamRefH` is read-only, the user never sees it.
  但是，在时间轴面板中仍然可以看到子节点，而且不缩进。
- `AEGP_DynStreamFlag_SHOWN_WHEN_EMPTY` New in CS6. A read-only flag. Indicates that this stream group should be shown when empty.
- `AEGP_DynStreamFlag_SKIP_REVEAL_WHEN_UNHIDDEN` New in CS6. A read-only flag. Indicates that this stream property will not be automatically revealed when un-hidden.

#### AEGP_SetDynamicStreamFlag

为`AEGP_StreamRefH`设置指定的标志。
注意：标志必须被单独设置。如果`undoableB`为`TRUE`，则可撤销。

```cpp
AEGP_SetDynamicStreamFlag(
 AEGP_StreamRefH streamH,
 AEGP_DynStreamFlags one_flag,
 A_Boolean undoableB,
 A_Boolean setB);
```

这个调用可以通过设置和清除`AEGP_DynStreamFlag_HIDDEN`来动态显示或隐藏参数。
然而，`AEGP_DynStreamFlag_DISABLED`不能被设置。

#### AEGP_GetNewStreamRefByIndex

从给定的`AEGP_StreamRefH`中按索引检索一个子流。不能用于类型为`AEGP_StreamGroupingType_LEAF`的流。

```cpp
AEGP_GetNewStreamRefByIndex(
 AEGP_PluginID aegp_plugin_id,
 AEGP_StreamRefH parent_groupH,
 A_long indexL,
 AEGP_StreamRefH *streamPH);
```

#### AEGP_GetNewStreamRefByMatchname

从给定的`AEGP_StreamRefH`中按匹配名称检索一个子流。只对`AEGP_StreamGroupingType_NAMED_GROUP`合法。

```cpp
AEGP_GetNewStreamRefByMatchname(
 AEGP_PluginID aegp_plugin_id,
 AEGP_StreamRefH parent_groupH,
 const A_char *match_nameZ,
 AEGP_StreamRefH *streamPH);
```

这里有一些方便的流的名称，可以检索到它们的引用。

- `AEGP_StreamGroupName_MASK_PARADE`
- `AEGP_StreamGroupName_MASK_ATOM`
- `AEGP_StreamName_MASK_FEATHER`
- `AEGP_StreamName_MASK_OPACITY`
- `AEGP_StreamName_MASK_OFFSET`
- `AEGP_StreamGroupName_EFFECT_PARADE`
- `AEGP_StreamGroupName_LAYER`
- `AEGP_StreamGroupName_AV_LAYER`
- `AEGP_StreamGroupName_TEXT_LAYER`
- `AEGP_StreamGroupName_CAMERA_LAYER`
- `AEGP_StreamGroupName_LIGHT_LAYER`
- `AEGP_StreamGroupName_AUDIO`
- `AEGP_StreamGroupName_MATERIAL_OPTIONS`
- `AEGP_StreamGroupName_TRANSFORM`
- `AEGP_StreamGroupName_LIGHT_OPTIONS`
- `AEGP_StreamGroupName_CAMERA_OPTIONS`

#### AEGP_DeleteStream

从一个流组中删除指定的流。
注意，调用者仍然必须处理它已经通过 API 获得(分配)的任何`AEGP_StreamRefH`。可撤销。
只对`AEGP_StreamGroupingType_INDEXED_GROUP`类型的子集有效。

```cpp
AEGP_DeleteStream(
 AEGP_StreamRefH streamH);
```

注意：从 6.5 开始，如果一个流被删除，而它或任何子流被选中，当前的合成选择将变成`NULL`。

#### AEGP_ReorderStream

设置指定的`AEGP_StreamRefH`的新索引。可撤销。
只对`AEGP_StreamGroupingType_INDEXED_GROUP`的子流有效。
`AEGP_StreamRefH`被更新以引用新排序的流。

```cpp
AEGP_ReorderStream(
 AEGP_StreamRefH streamH
 A_long new_indexL);
```

#### AEGP_DuplicateStream

复制指定的流并将其附加到流组中。
可撤销。
只对`AEGP_StreamGroupingType_INDEXED_GROUP`的子集有效。

```cpp
AEGP_DuplicateStream(
 AEGP_PluginID aegp_plugin_id,
 AEGP_StreamRefH streamH,
 A_long *new_indexPL0);
```

#### AEGP_SetStreamName

设置给定的`AEGP_StreamRefH`的名称。nameZ 指向一个空尾的 UTF-16 字符串，可撤销。
只对`AEGP_StreamGroupingType_INDEXED_GROUP`的子集有效。 注意：如果你在 force_englishB 设置为`TRUE`的情况下检索名称，你将得到流的典型的、未改变的名称。

```cpp
AEGP_SetStreamName(
 AEGP_StreamRefH streamH,
 const A_UTF16Char *nameZ);
```

注意：在一个效果流的组上使用这个功能来改变效果的显示名称。

#### AEGP_CanAddStream

返回当前是否可以通过 API 添加一个流。

```cpp
AEGP_CanAddStream(
 AEGP_StreamRefH group_streamH,
 const A_char *match_nameZ,
 A_Boolean *can_addPB);
```

#### AEGP_AddStream

将一个流添加到指定的流组中。可撤销。只对`AEGP_StreamGroupingType_INDEXED_GROUP`有效。

```cpp
AEGP_AddStream(
 AEGP_PluginID aegp_plugin_id,
 AEGP_StreamRefH indxd_grp_streamH,
 const A_char *match_nameZ,
 AEGP_StreamRefH *streamPH0);
```

#### AEGP_GetMatchName

检索指定的`AEGP_StreamRefH`的匹配名称。
注意，这可能与显示名称不同，显示名称可以用`AEGP_GetStreamName`检索，在[AEGP_StreamSuite5]。
`nameZ`的长度可以达到`AEGP_MAX_STREAM_MATCH_NAME_SIZE`。

```cpp
AEGP_GetMatchName(
 AEGP_StreamRefH streamH,
 A_char *nameZ);
```

#### AEGP_GetNewParentStreamRef

检索指定的`AEGP_StreamRefH`的父类的`AEGP_StreamRefH`。

```cpp
AEGP_GetNewParentStreamRef(
 AEGP_PluginID plugin_id,
 AEGP_StreamRefH streamH,
 AEGP_StreamRefH *parentPH);
```

#### AEGP_GetStreamIsModified

返回指定的`AEGP_StreamRefH`是否被修改。
注意：同样的结果可以通过 After Effect 用户界面，在选择合成时键入 "UU"。

```cpp
AEGP_GetStreamIsModified(
 AEGP_StreamRefH streamH,
 A_Boolean *modifiedPB);
```

#### AEGP_GetStreamIndexInParent

检索一个给定流的索引，相对于它的父流。
只对`AEGP_StreamGroupingType_INDEXED_GROUP`的子集有效。

```cpp
AEGP_GetStreamIndexInParent(
 AEGP_StreamRefH streamH,
 A_long *indexPL);
```

注意：正如在其他地方提到的，`AEGP_StreamRefHs`不会在不同的函数调用中持久存在。 如果流被重新排序，添加或删除，所有以前检索到的`AEGP_StreamRefHs`可能会被废止。

#### AEGP_IsSeparationLeader

只对叶子流有效。如果这个流是一个多维流，可以分离它的维度，返回 true，尽管它们目前可能没有被分离。
术语。Leader 是可以被分离的流，Follower 是 N 个自动流中的一个，对应于 Leader 的 N 个维度。
Leader 并不总是分离的，调用`AEGP_AreDimensionsSeparated`来确定它是否被分离。
从 CS4 开始，唯一被分离的流是图层的位置属性。
请不要编写假设的代码，我们预计将来会允许分离更多的数据流。

```cpp
AEGP_IsSeparationLeader(
 AEGP_StreamRefH streamH,
 A_Boolean *leaderPB);
```

#### AEGP_AreDimensionsSeparated

诸如`AEGP_GetNewKeyframeValue`这样对关键帧指数起作用的方法肯定不会对 Leader 属性起作用，你将需要明确地检索和操作 Followers。

```cpp
AEGP_AreDimensionsSeparated(
 AEGP_StreamRefH streamH,
 A_Boolean *separatedPB);
```

#### AEGP_SetDimensionsSeparated

只有当`AEGP_IsSeparationLeader()`为`true`时才有效。

```cpp
AEGP_AreDimensionsSeparated(
 AEGP_StreamRefH streamH,
 A_Boolean *separatedPB);
```

#### AEGP_GetSeparationFollower

检索与 Leader 流的特定维度相对应的 Follower 流。`dimS`的范围从`0`到`AEGP_GetStreamValueDimensionality(lea der_streamH) - 1`。

```cpp
AEGP_GetSeparationFollower(
 AEGP_StreamRefH leader_streamH
 A_short dimS,
 AEGP_StreamRefH *follower_streamPH);
```

#### AEGP_IsSeparationFollower

只对叶子流有效。
如果这个流是一个一维的属性，代表 Leader 的一个维度，则返回`true`。 你可以使用`AEGP_GetSeparationFollower()`从 Leader 那里检索流。

```cpp
AEGP_IsSeparationFollower(
 AEGP_StreamRefH streamH
 A_Boolean *followerPB);
```

#### AEGP_GetSeparationLeader

仅对分离追随者有效，返回它所属的 Leader。

```cpp
AEGP_GetSeparationLeader(
 AEGP_StreamRefH follower_streamH,
 AEGP_StreamRefH *leader_streamPH);
```

#### AEGP_GetSeparationDimension

只对分离跟随者有效，返回它所对应的 Leader 的维度。

```cpp
AEGP_GetSeparationDimension(
 AEGP_StreamRefH follower_streamH,
 A_short *dimPS);
```

## Working With Keyframes

关键帧使 After Effects 成为它的一部分。AEGPs(以及......嘘，别告诉别人......特效)可以使用这套软件来添加、操作和删除任何可使用关键帧的流的关键帧。

### AEGP_KeyframeSuite3

#### AEGP_GetStreamNumKFs

检索给定流上的关键帧的数量。
如果流不支持关键帧，返回`AEGP_NumKF_NO_DATA`。
另外，注意没有关键帧的流不一定是恒定的；它可以被表达式所改变。

```cpp
AEGP_GetStreamNumKFs(
 AEGP_StreamRefH streamH,
 A_long *num_kfsPL);
```

#### AEGP_GetKeyframeTime

检索指定关键帧的时间。

```cpp
AEGP_GetKeyframeTime(
 AEGP_StreamRefH streamH,
 AEGP_KeyframeIndex index,
 AEGP_LTimeMode time_mode,
 A_Time *timePT);
```

#### AEGP_InsertKeyframe

给指定的流添加一个关键帧(在指定的合成或层时间)。
返回新关键帧的索引。
所有大于新索引的索引现在都是无效的(但你知道)。如果当时已经有一个关键帧，数值将被更新。

```cpp
AEGP_InsertKeyframe(
 AEGP_StreamRefH streamH,
 AEGP_LTimeMode time_mode,
 const A_Time *timePT,
 AEGP_KeyframeIndex *key_indexP);
```

#### AEGP_DeleteKeyframe

删除指定的关键帧。

```cpp
AEGP_DeleteKeyframe(
 AEGP_StreamRefH streamH,
 AEGP_KeyframeIndex key_index);
```

#### AEGP_GetNewKeyframeValue

为关键帧时的流值创建并填充一个`AEGP_StreamValue2`。
返回的`AEGP_StreamValue2`必须使用`AEGP_DisposeStreamValue`来处理。

```cpp
AEGP_GetNewKeyframeValue(
 AEGP_PluginID plugin_id,
 AEGP_StreamRefH streamH,
 AEGP_KeyframeIndex key_index,
 AEGP_StreamValue2 *valueP);
```

#### AEGP_SetKeyframeValue

设置关键帧时的流的值。

```cpp
AEGP_SetKeyframeValue(
 AEGP_StreamRefH streamH,
 AEGP_KeyframeIndex index,
 const AEGP_StreamValue2 *valP);
```

#### AEGP_GetStreamValueDimensionality

检索流值的维度。

```cpp
AEGP_GetStreamValueDimensionality(
 AEGP_StreamRefH streamH,
 A_short *value_dimPS);
```

#### AEGP_GetStreamTemporalDimensionality

检索流的时间维度。

```cpp
AEGP_GetStreamTemporalDimensionality(
 AEGP_StreamRefH streamH,
 A_short *t_dimPS);
```

#### AEGP_GetNewKeyframeSpatialTangents

返回代表关键帧时流的切向值的`AEGP_StreamValue2s`。 返回的`AEGP_StreamValue2s`必须使用`AEGP_DisposeStreamValue`进行处理。

```cpp
AEGP_GetNewKeyframeSpatialTangents(
 AEGP_PluginID plugin_id,
 AEGP_StreamRefH streamH,
 AEGP_KeyframeIndex key_index,
 AEGP_StreamValue2 *in_tanP0,
 AEGP_StreamValue2 *out_tanP0);
```

#### AEGP_SetKeyframeSpatialTangents

指定切向的`AEGP_StreamValue2s`在关键帧时用于流的值。
用于切入和切出的`AEGP_StreamValue2s`不被After Effects采用，必须用`AEGP_DisposeStreamValue`来处理。

```cpp
AEGP_SetKeyframeSpatialTangents(
 AEGP_StreamRefH streamH,
 AEGP_KeyframeIndex key_index,
 const AEGP_StreamValue2 *in_tP0,
 const AEGP_StreamValue2 *out_tP0);
```

注意：在`AEGP_KeyframeSuite2`和以前的版本中，当调用效果点控制流或锚点时，从这个函数返回的值是错误的。
它们没有乘以图层大小。现在，它们是。

#### AEGP_GetKeyframeTemporalEase

检索与关键帧时流值的指定维度相关的`AEGP_KeyframeEases`。
`dimensionL`范围从`0`到`(temporal_dimensionality -1)`。

```cpp
AEGP_GetKeyframeTemporalEase(
 AEGP_StreamRefH streamH,
 AEGP_KeyframeIndex key_index,
 A_long dimensionL,
 AEGP_KeyframeEase *in_easeP0,
 AEGP_KeyframeEase *out_easeP0);
```

注意：返回的缓和值必须乘以图层高度，以匹配 After Effects UI 中显示的值。

#### AEGP_SetKeyframeTemporalEase

指定在关键帧时用于流的值的`AEGP_KeyframeEases`。`dimensionL`范围从`0`到`(temporal_dimensionality -1)`。
所传递的`AEGP_KeyframeEases`不被 After Effects 采用。

```cpp
AEGP_SetKeyframeTemporalEase(
 AEGP_StreamRefH streamH,
 AEGP_KeyframeIndex key_index,
 A_long dimL,
 const AEGP_KeyframeEase *in_P0,
 const AEGP_KeyframeEase *outP0);
```

#### AEGP_GetKeyframeFlags

检索当前为关键帧设置的标志。

```cpp
AEGP_GetKeyframeFlags(
 AEGP_StreamRefH streamH,
 AEGP_KeyframeIndex key_index,
 AEGP_KeyframeFlags *flagsP);
```

`*flagsP`将是以下的合成。

- `AEGP_KeyframeFlag_NONE`
- `AEGP_KeyframeFlag_TEMPORAL_CONTINUOUS`
- `AEGP_KeyframeFlag_TEMPORAL_AUTOBEZIER`
- `AEGP_KeyframeFlag_SPATIAL_CONTINUOUS`
- `AEGP_KeyframeFlag_SPATIAL_AUTOBEZIER`
- `AEGP_KeyframeFlag_ROVING`

#### AEGP_SetKeyframeFlag

为关键帧设置指定的标志。标志必须单独设置。

```cpp
AEGP_SetKeyframeFlag(
 AEGP_StreamRefH streamH,
 AEGP_KeyframeIndex key_index,
 AEGP_KeyframeFlags flag,
 A_Boolean valueB);
```

#### AEGP_GetKeyframeInterpolation

检索指定关键帧的输入和输出`AEGP_KeyframeInterpolationTypes`。

```cpp
AEGP_GetKeyframeInterpolation(
 AEGP_StreamRefH streamH,
 AEGP_KeyframeIndex key_index,
 AEGP_KeyframeInterpolationType *inP0,
 AEGP_KeyframeInterpolationType *outP0);
```

`AEGP_KeyframeInterpolationType`是以下之一。

- `AEGP_KeyInterp_NONE`
- `AEGP_KeyInterp_LINEAR`
- `AEGP_KeyInterp_BEZIER`
- `AEGP_KeyInterp_HOLD`

#### AEGP_SetKeyframeInterpolation

指定用于指定关键帧的 AEGP_KeyframeInterpolationTypes 的输入和输出。

```cpp
AEGP_SetKeyframeInterpolation(
 AEGP_StreamRefH streamH,
 AEGP_KeyframeIndex key_index,
 AEGP_KeyframeInterpolationType in_interp,
 AEGP_KeyframeInterpolationType out_interp);
```

#### AEGP_StartAddKeyframes

通知 After Effects，你将在指定的流中加入几个关键帧。
After Effects 将返回一个分配的不透明的`AEGP_AddKeyframesInfoH`，供下面的调用使用。

```cpp
AEGP_StartAddKeyframes(
 AEGP_StreamRefH streamH,
 AEGP_AddKeyframesInfoH *akPH);
```

#### AEGP_AddKeyframes

在指定的(层或合成)时间向指定的流添加一个关键帧。
注意：这实际上并不对流的值做任何事情。

```cpp
AEGP_AddKeyframes(
 AEGP_AddKeyframesInfoH akH,
 AEGP_LTimeMode time_mode,
 const A_Time *timePT,
 A_long *indexPL);
```

#### AEGP_SetAddKeyframe

设置指定关键帧的值。

```cpp
AEGP_SetAddKeyframe(
 AEGP_AddKeyframesInfoH akH,
 A_long indexL,
 const AEGP_StreamValue2 *valueP);
```

#### AEGP_EndAddKeyframes

告诉 After Effects 你已经完成了关键帧的添加。

```cpp
AEGP_EndAddKeyframes(
 A_Boolean addB,
 AEGP_AddKeyframesInfoH akH);
```

## Adding Multiple Keyframes

每次你调用`AEGP_InsertKeyframe()`，整个流都会被添加到撤销堆栈中。

如果你正在添加一两个关键帧，这不是一个问题。然而，如果你在写一个关键帧，你会想用正确的方法来做事情。

在你开始添加关键帧之前，调用(非常合适的名称)`AEGP_StartAddKeyframes`，并传递给它一个不透明的`AEGP_AddKeyframesInfoH`。

对于每个要添加的关键帧，调用`AEGP_AddKeyframes`来设置要使用的时间(并获得新添加的关键帧的索引)，然后调用`AEGP_SetAddKeyframe`来指定要使用的值。

完成后，调用`AEGP_EndAddKeyframes`，让 After Effects 知道是时候将改变的参数流添加到撤销堆栈中。

## Marker Streams

`AEGP_MarkerSuite`允许直接操作标记数据。

### AEGP_MarkerSuite2

#### AEGP_NewMarker

创建一个新的标记。

```cpp
AEGP_NewMarker(
 AEGP_MarkerValP *markerPP);
```

#### AEGP_DisposeMarker

处置一个标记。

```cpp
AEGP_DisposeMarker(
 AEGP_MarkerValP markerP);
```

#### AEGP_DuplicateMarker

复制一个标记(没料到会这样，嗯？)

```cpp
AEGP_DuplicateMarker(
 AEGP_MarkerValP markerP,
 AEGP_MarkerValP *new_markerP);
```

#### AEGP_SetMarkerFlag

设置一个标记标志的值。

```cpp
AEGP_SetMarkerFlag(
 AEGP_MarkerValP markerP,
 AEGP_MarkerFlagType flagType,
 A_Boolean valueB);
```

目前，AEGP_MarkerFlagType 是以下的一种。

- `AEGP_MarkerFlag_NONE`
- `AEGP_MarkerFlag_NAVIGATION`

#### AEGP_GetMarkerFlag

获取给定的`AEGP_MarkerFlagType`的值(见上文)。

```cpp
AEGP_GetMarkerFlag(
 AEGP_ConstMarkerValP markerP,
 AEGP_MarkerFlagType flagType
 A_Boolean *valueBP);
```

#### AEGP_GetMarkerString

检索位于指定标记字段中的 UTF-16、以 NULL 为结尾的字符串。
必须由调用者使用`AEGP_FreeMemHandle`处理。

```cpp
AEGP_GetMarkerString(
 AEGP_PluginID id,
 AEGP_ConstMarkerValP markerP,
 AEGP_MarkerStringType strType,
 AEGP_MemHandle *unicodePH);
```

#### AEGP_SetMarkerString

将标记的指定字段设置为提供的文本。

```cpp
AEGP_SetMarkerString(
 AEGP_MarkerValP markerP,
 AEGP_MarkerStringType strType,
 const A_u_short *unicodeP,
 A_long lengthL);
```

#### AEGP_CountCuePointParams

返回提示点参数的数量。

```cpp
AEGP_CountCuePointParams(
 AEGP_ConstMarkerValP markerP,
 A_long *paramsLP);
```

#### AEGP_GetIndCuePointParam

返回指定索引的提示点参数(必须在`0`和`(提示点参数-1)`之间。
返回的句柄是 UTF-16，以 NULL 结尾的字符串，必须由调用者使用`AEGP_FreeMemHandle`来处理。

```cpp
AEGP_GetIndCuePointParam(
 AEGP_PluginID id,
 AEGP_ConstMarkerValP markerP,
 A_long param_indexL,
 AEGP_MemHandle *unicodeKeyPH,
 AEGP_MemHandle *uni_ValuePH);
```

#### AEGP_SetIndCuePointParam

将一个有索引的提示点参数的值设置为指定值。
`key_lengthL`是 "unicode 字符数"，`value_lenL`是提供值的长度。
`unicode_KeyP`和`unicode_ValueP`指向 UTF-16 数据。

```cpp
AEGP_SetIndCuePointParam(
 AEGP_MarkerValP markerP,
 A_long param_idxL,
 const A_u_short *unicode_KeyP,
 A_long key_lengthL,
 const A_u_short *unicode_ValueP,
 A_long value_lengthL);
```

#### AEGP_InsertCuePointParam

插入一个提示点参数。
这个调用由`AEGP_SetIndCuePointParam`跟随，以实际设置数据。

```cpp
AEGP_InsertCuePointParam(
 AEGP_MarkerValP markerP,
 A_long param_idxL);
```

| `AEGP_DeleteInd` CuePointParam | Deletes the cue point param at the specified index.

```cpp
AEGP_DeleteIndCuePointParam(
 AEGP_MarkerValP markerP,
 A_long param_idxL);
```

#### AEGP_SetMarkerDuration

```cpp
AEGP_SetMarkerDuration(
 AEGP_MarkerValP markerP,
 const A_Time *durationPT);
```

#### AEGP_GetMarkerDuration

```cpp
AEGP_GetMarkerDuration(
 AEGP_ConstMarkerValP markerP,
 A_Time *durationPT);
```

## Mask Management

访问、操作和删除一个图层的遮罩。

### AEGP_MaskSuite6

#### AEGP_GetLayerNumMasks

计算应用于一个图层的蒙版。

```cpp
AEGP_GetLayerNumMasks(
 AEGP_LayerH aegp_layerH,
 A_long *num_masksPL);
```

#### AEGP_GetLayerMaskByIndex

给定一个图层句柄和蒙版索引，返回一个指向蒙版句柄的指针。
你必须通过使用`AEGP_DisposeMask()`来销毁蒙版句柄。

```cpp
AEGP_GetLayerMaskByIndex(
 AEGP_LayerH aegp_layerH,
 A_long mask_indexL,
 AEGP_MaskRefH *maskPH);
```

#### AEGP_DisposeMask

处置一个掩码句柄。

```cpp
AEGP_DisposeMask(
 AEGP_MaskRefH maskH);
```

#### AEGP_GetMaskInvert

给定一个掩码句柄，确定该掩码是否被反转。

```cpp
AEGP_GetMaskInvert(
 AEGP_MaskRefH maskH,
 A_Boolean *invertPB);
```

#### AEGP_SetMaskInvert

设置一个掩码的反转状态。

```cpp
AEGP_SetMaskInvert(
 AEGP_MaskRefH mask_refH,
 A_Boolean invertB);
```

#### AEGP_GetMaskMode

给定一个掩码句柄，返回该掩码的当前模式。
`PF_MaskMode_NONE`不做任何事情，`PF_MaskMode_ADD`是默认行为。

- `PF_MaskMode_ADD`,
- `PF_MaskMode_SUBTRACT`,
- `PF_MaskMode_INTERSECT`,
- `PF_MaskMode_LIGHTEN`,
- `PF_MaskMode_DARKEN`,
- `PF_MaskMode_DIFFERENCE`,

```cpp
AEGP_GetMaskMode(
 AEGP_MaskRefH maskH,
 PF_MaskMode *modeP);
```

#### AEGP_SetMaskMode

设置给定遮罩的模式。

```cpp
AEGP_SetMaskMode(
 AEGP_MaskRefH maskH,
 PF_MaskMode mode);
```

#### AEGP_GetMaskMotionBlurState

检索给定遮罩的运动模糊设置。

```cpp
AEGP_GetMaskMotionBlurState(
 AEGP_MaskRefH mask_refH,
 AEGP_MaskMBlur *blur_stateP);
```

`AEGP_MaskMBlur`将是以下的一种。

- `AEGP_MaskMBlur_SAME_AS_LAYER`
- `AEGP_MaskMBlur_OFF`
- `AEGP_MaskMBlur_ON`

#### AEGP_SetMaskMotionBlurState

CS6 中的新设置。设置给定遮罩的运动模糊设置。

```cpp
AEGP_SetMaskMotionBlurState(
 AEGP_MaskRefH mask_refH,
 AEGP_MaskMBlur blur_state);
```

#### AEGP_GetMaskFeatherFalloff

CS6 中新增。获取给定蒙版的羽化淡出类型，可以是
`AEGP_MaskFeatherFalloff_SMOOTH`或`AEGP_MaskFeatherFalloff_LINEAR`。

```cpp
AEGP_SetMaskMotionBlurState(
 AEGP_MaskRefH mask_refH,
 AEGP_MaskFeatherFalloff *feather_falloffP);
```

#### AEGP_SetMaskFeatherFalloff

为给定的蒙版设置羽化淡出的类型。

```cpp
AEGP_SetMaskMotionBlurState(
 AEGP_MaskRefH mask_refH,
 AEGP_MaskFeatherFalloff feather_falloff);
```

#### AEGP_GetMaskName

在 CS4 中被删除。使用`AEGP_GetNewStreamRefForMask`和动态流套件中的名称函数代替。

#### AEGP_SetMaskName

#### AEGP_GetMaskID

为给定的`AEGP_MaskRefH`检索`AEGP_MaskIDVal`，用于唯一识别掩码。

```cpp
AEGP_GetMaskID(
 AEGP_MaskRefH mask_refH,
 AEGP_MaskIDVal *id_valP);
```

#### AEGP_CreateNewMask

在引用的`AEGP_LayerH`上创建一个新的掩码，节点数为 0。新掩码的索引被返回。

```cpp
AEGP_CreateNewMask(
 AEGP_LayerH layerH,
 AEGP_MaskRefH *mask_refPH,
 A_long *mask_indexPL0);
```

#### AEGP_DeleteMaskFromLayer

```cpp
AEGP_DeleteMaskFromLayer(
 AEGP_MaskRefH mask_refH);
```

注意：从 6.5 版本开始，如果你删除了一个遮罩，并且它或一个子流被选中，在合成中的当前选择将变成 NULL。

#### AEGP_GetMaskColor

检索指定掩码的颜色。

```cpp
AEGP_GetMaskColor(
 AEGP_MaskRefH mask_refH,
 AEGP_ColorVal *colorP);
```

#### AEGP_SetMaskColor

设置指定掩码的颜色。

```cpp
AEGP_SetMaskColor(
 AEGP_MaskRefH mask_refH,
 const AEGP_ColorVal *colorP);
```

#### AEGP_GetMaskLockState

检索指定掩码的锁定状态。

```cpp
AEGP_GetMaskLockState(
 AEGP_MaskRefH mask_refH,
 A_Boolean *is_lockedPB);
```

#### AEGP_SetMaskLockState

设置指定掩码的锁定状态。

```cpp
AEGP_SetMaskLockState(
 AEGP_MaskRefH mask_refH,
 A_Boolean lockB);
```

#### AEGP_GetMaskIsRotoBezier

返回给定的遮罩是否被用作滚轴。

```cpp
AEGP_GetMaskIsRotoBezier(
 AEGP_MaskRefH mask_refH,
 A_Boolean *is_roto_bezierPB);
```

#### AEGP_SetMaskIsRotoBezier

设置给定的蒙板是否用作滚轴。

```cpp
AEGP_SetMaskIsRotoBezier(
 AEGP_MaskRefH mask_refH,
 A_Boolean *is_roto_bezierPB);
```

#### AEGP_DuplicateMask

复制一个给定的`AEGP_MaskRefH`。调用者必须处理掉重复的部分。

```cpp
AEGP_DuplicateMask(
 AEGP_MaskRefH orig_mask_refH,
 AEGP_MaskRefH *dupe_mask_refPH);
```

## Mask Outlines

上面的 Mask Suite 告诉插件一个图层上的遮罩，但不告诉这些遮罩的细节。

这是因为 After Effects 需要处理这些信息；这些信息并不是随便乱放的。

插件使用这个蒙版轮廓套件来访问这些信息。

### AEGP_MaskOutlineSuite3

#### AEGP_IsMaskOutlineOpen

给定一个遮罩轮廓指针(可通过[Stream Suite]获得)，确定遮罩路径是打开还是关闭。

```cpp
AEGP_IsMaskOutlineOpen(
 AEGP_MaskOutlineVal *mask_outlineP,
 A_Boolean *openPB);
```

#### AEGP_SetMaskOutlineOpen

设置给定掩码轮廓的打开状态。

```cpp
AEGP_SetMaskOutlineOpen(
 AEGP_MaskOutlineValH mask_outlineH,
 A_Boolean openB);
```

#### AEGP_GetMaskOutlineNumSegments

给定一个掩码轮廓指针，返回路径中的段数。
`num_segmentsPL`是段的总数`[0...N-1]`。

```cpp
AEGP_GetMaskOutlineNumSegments(
 AEGP_MaskOutlineVal *mask_outlineP,
 A_long *num_segmentsPL);
```

#### AEGP_GetMaskOutlineVertexInfo

给出一个掩码轮廓指针和 0 与总段数之间的一个点。
对于封闭的遮罩路径，`vertex[0]`与`vertex[num_segments]`相同。

```cpp
AEGP_GetMaskOutlineVertexInfo(
 AEGP_MaskOutlineVal *mask_outlineP,
 A_long which_pointL,
 AEGP_MaskVertex *vertexP);
```

#### AEGP_SetMaskOutlineVertexInfo

设置给定索引的顶点信息。
设置顶点 0 是特别的；它的切线实际上是设置轮廓中最后一个顶点的切线。
当然，`which_pointL`必须对遮罩的轮廓有效，否则函数将返回错误。

```cpp
AEGP_SetMaskOutlineVertexInfo(
 AEGP_MaskOutlineValH mask_outlineH,
 AEGP_VertexIndex which_pointL,
 AEGP_MaskVertex *vertexP);
```

#### AEGP_CreateVertex

在索引位置创建一个顶点。
所有以前有`AEGP_VertexIndex`的位置或更高的顶点将被增加一个索引。

```cpp
AEGP_CreateVertex(
 AEGP_MaskOutlineValH mask_outlineH,
 AEGP_VertexIndex position);.
```

注意：所有蒙版必须至少有一个顶点。

#### AEGP_DeleteVertex

从遮罩中删除一个顶点。

```cpp
AEGP_DeleteVertex(
 AEGP_MaskOutlineValH mask_outlineH,
 AEGP_VertexIndex index);
```

#### AEGP_GetMaskOutlineNumFeathers

CS6 中的新功能。

```cpp
AEGP_DeleteVertex(
 AEGP_MaskOutlineValH mask_outlineH,
 A_long *num_feathersPL);
```

#### AEGP_GetMaskOutlineFeatherInfo

CS6 中的新功能。

```cpp
AEGP_GetMaskOutlineFeatherInfo(
 AEGP_MaskOutlineValH mask_outlineH,
 AEGP_FeatherIndex which_featherL,
 AEGP_MaskFeather *featherP);
```

#### AEGP_SetMaskOutlineFeatherInfo

CS6 中的新功能。
羽毛必须已经存在；如果需要，首先使用`AEGP_CreateMaskOutlineFeather`。

```cpp
AEGP_SetMaskOutlineFeatherInfo(
 AEGP_MaskOutlineValH mask_outlineH,
 AEGP_VertexIndex which_featherL,
 const AEGP_MaskFeather *featherP);
```

#### AEGP_CreateMaskOutlineFeather

CS6 中的新功能。新羽毛的索引在`insert_positionP`中传回。

```cpp
AEGP_CreateMaskOutlineFeather(
 AEGP_MaskOutlineValH mask_outlineH,
 const AEGP_MaskFeather *featherP0,
 AEGP_FeatherIndex *insert_positionP);
```

#### AEGP_DeleteMaskOutlineFeather

CS6 中的新功能。

```cpp
AEGP_DeleteMaskOutlineFeather(
 AEGP_MaskOutlineValH mask_outlineH,
 AEGP_FeatherIndex index);
```

## Mask Feathering

CS6 的新功能，遮罩可以被羽化。

`AEGP_MaskFeather`定义如下。

```cpp
typedef struct {
 A_long segment; // mask segment where feather is
 PF_FpLong segment_sF; // 0-1: feather location on segment
 PF_FpLong radiusF; // negative value allowed if type == AEGP_MaskFeatherType_INNER
 PF_FpShort ui_corner_angleF; // 0-1: angle of UI handle on corners
 PF_FpShort tensionF; // 0-1: tension of boundary at feather pt
 AEGP_MaskFeatherInterp interp;
 AEGP_MaskFeatherType type;
} AEGP_MaskFeather;
```

`AEGP_MaskFeatherInterp`是`AEGP_MaskFeatherInterp_NORMAL`或`AEGP_MaskFeatherInterp_HOLD_CW`。

`AEGP_MaskFeatherType`是`AEGP_MaskFeatherType_OUTER`或`AEGP_MaskFeatherType_INNER`。

这个套件使 AEGP 能够获取和设置与文本层相关的文本。

注意：要开始，通过调用上面的 "AEGP_GetLayerStreamValue"，并将 "AEGP_StreamType_TEXT_DOCUMENT "作为 "AEGP_StreamType "来检索一个 "AEGP_TextDocumentH"。

## Working With Text Layers

该套件使 AEGP 能够获得和设置与文本层相关的文本。

### AEGP_TextDocumentSuite1

#### AEGP_GetNewText

检索在`AEGP_TextDocumentH`中使用的 UTF-16、以 NULL 为结尾的字符串。
注意：After Effects 将分配`AEGP_MemHandle`。
你的插件必须在完成后使用`AEGP_FreeMemHandle`来处理它。

```cpp
AEGP_GetNewText(
 AEGP_PluginID id,
 AEGP_TextDocumentH text_docH,
 AEGP_MemHandle *unicodePH);
```

#### AEGP_SetText

指定 "AEGP_TextDocumentH "使用的文本。

```cpp
AEGP_SetText(
 AEGP_TextDocumentH text_docH,
 const A_u_short *unicodePS,
 long lengthL);
```

## Working With Text Outlines

`AEGP_TextLayerSuite`提供对文本层使用的实际文本轮廓的访问。

一旦有了路径，就可以用[PF_PathQuerySuite1](../effect-details/working-with-paths.html#effect-details-working-with-paths-pf-pathquerysuite)和[PF_PathDataSuite](../effect-details/working-with-paths.html#effect-details-working-with-paths-pf-pathdatasuite)来操作它。

### AEGP_TextLayerSuite1

#### AEGP_GetNewTextOutlines

分配并返回与指定图层相关的`AEGP_TextOutlinesHs`的句柄。
如果没有与`layerH`相关的`AEGP_TextOutlinesHs`(换句话说，如果它不是一个文本层)，`outlinesPH`将是 NULL。

```cpp
AEGP_GetNewTextOutlines(
 AEGP_LayerH layerH,
 const A_Time *layer_timePT,
 AEGP_TextOutlinesH *outlinesPH);
```

#### AEGP_DisposeTextOutlines

处理掉我们为你分配的那些轮廓线!

```cpp
AEGP_DisposeTextOutlines(
 AEGP_TextOutlinesH outlinesH);
```

#### AEGP_GetNumTextOutlines

检索该层的文本轮廓的数量。

```cpp
AEGP_GetNumTextOutlines(
 AEGP_TextOutlinesH outlinesH,
 A_long *num_otlnsPL);
```

#### AEGP_GetIndexedTextOutline

返回一个指定文本轮廓的 PF_PathOutlinePtr。

```cpp
AEGP_GetIndexedTextOutline(
 AEGP_TextOutlinesH outlinesH,
 A_long path_indexL,
 PF_PathOutlinePtr *pathPP);
```

## Utility Functions

Utility 套件提供了错误信息处理、AEGP 版本检查和对撤销堆栈的访问。

你需要的一切，以保持 After Effects 和你的插件整洁。

### AEGP_UtilitySuite6

#### AEGP_ReportInfo

显示对话框，显示 AEGP 的名称和所传递的字符串。

```cpp
AEGP_ReportInfo(
 AEGP_PluginID aegp_plugin_id,
 const A_char *info_stringZ);
```

#### AEGP_ReportInfoUnicode

CC 中的新内容。显示对话框，显示 AEGP 的名称和传递的 unicode 字符串。

```cpp
AEGP_ReportInfoUnicode(
 AEGP_PluginID aegp_plugin_id,
 const A_UTF16Char *info_stringP);
```

#### AEGP_GetDriverSpecVersion

返回`AEGPDriver`插件的版本(用于确定支持的功能)。

```cpp
AEGP_GetDriverSpecVersion(
 A_short *major_versionPS,
 A_short *minor_versionPS);
```

#### AEGP_StartQuietErrors

沉默错误。必须与`AEGP_EndQuietErrors`平衡。
`AEGP_ErrReportState`是一个不透明的结构，为 After Effects 私有。

```cpp
AEGP_StartQuietErrors(
 AEGP_ErrReportState *err_stateP);
```

#### AEGP_EndQuietErrors

重新启用错误。

```cpp
AEGP_EndQuietErrors(
 AEGP_ErrReportState *err_stateP)
```

#### AEGP_StartUndoGroup

将动作添加到撤销队列中。用户可以撤销在此和`AEGP_EndUndoGroup()`之间的任何动作。`undo_nameZ`将出现在编辑菜单中。

```cpp
AEGP_StartUndoGroup(
 const A_char *undo_nameZ);
```

#### AEGP_EndUndoGroup

结束撤消列表。

```cpp
AEGP_EndUndoGroup();
```

#### AEGP_RegisterWithAEGP

返回一个 AEGP_PluginID，然后效果插件可以在调用整个 AEGP API 的许多函数时使用。
在`PF_Cmd_GLOBAL_SETUP`期间，效果器应该只调用这个函数一次，并保存`AEGP_PluginID`供以后使用。
第一个参数可以是任何数值，第二个参数应该是插件的匹配名称。

```cpp
AEGP_RegisterWithAEGP(
 AEGP_GlobalRefcon global_refcon,
 const A_char *plugin_nameZ,
 AEGP_PluginID *plugin_id);
```

#### AEGP_GetMainHWND

检索 After Effects 的 HWND；在 Windows 上显示你自己的对话框时很有用。
如果你不使用 After Effects 的 HWND，你的模态对话框将无法阻止与后面的窗口互动，痛苦将随之而来。

```cpp
AEGP_GetMainHWND(
 void *main_hwnd);
```

#### AEGP_ShowHideAllFloaters

切换是否显示浮动调色板。
使用时要小心；当你意外地改变他们的用户界面时，用户会感到不安。

```cpp
AEGP_ShowHideAllFloaters(
 A_Boolean include_tool_palB);
```

#### AEGP_PaintPalGetForeColor

从调色板中检索前景颜色。

```cpp
AEGP_PaintPalGetForeColor(
 AEGP_ColorVal *fore_colorP);
```

#### AEGP_PaintPalGetBackColor

从调色板中获取背景色。

```cpp
AEGP_PaintPalGetBackColor(
 AEGP_ColorVal *back_colorP);
```

#### AEGP_PaintPalSetForeColor

在调色板中设置前景色。

```cpp
AEGP_PaintPalSetForeColor(
 const AEGP_ColorVal *fore_colorP);
```

#### AEGP_PaintPalSetBackColor

在调色板中设置背景颜色。

```cpp
AEGP_PaintPalSetBackColor(
 const AEGP_ColorVal *back_colorP);
```

#### AEGP_CharPalGetFillColor

从字符调色板中检索填充颜色。

```cpp
AEGP_CharPalGetFillColor(
 A_Boolean *is_fcolor_definedPB,
 AEGP_ColorVal *fill_colorP);
```

#### AEGP_CharPalGetStrokeColor

从字符调色板中检索笔触颜色。

```cpp
AEGP_CharPalGetStrokeColor(
 A_Boolean *is_scolor_definedPB,
 AEGP_ColorVal *stroke_colorP);
```

#### AEGP_CharPalSetFillColor

设置字符调色板中的填充色。

```cpp
AEGP_CharPalSetFillColor(
 const AEGP_ColorVal *fill_colorP);
```

#### AEGP_CharPalSetStrokeColor

在字符调色板中设置笔画颜色。

```cpp
AEGP_CharPalSetStrokeColor(
 const AEGP_ColorVal *stroke_colorP);
```

#### AEGP_CharPalIsFillColorUIFrontmost

返回填充颜色是否是最前面的。如果不是，笔画颜色就是最前面的。

```cpp
AEGP_CharPalIsFillColorUIFrontmost(
 A_Boolean *is_fcolor_selectedPB);
```

#### AEGP_ConvertFpLongToHSFRatio

返回给定的`A_FpLong`的`A_Ratio`解释。对水平比例因子的解释很有用。

```cpp
AEGP_ConvertFpLongToHSFRatio(
 A_FpLong numberF,
 A_Ratio *ratioPR);
```

#### AEGP_ConvertHSFRatioToFpLong

返回给定 "A_FpLong "的 "A_Ratio "解释。对水平比例因子的解释很有用。

```cpp
AEGP_ConvertHSFRatioToFpLong(
 A_Ratio ratioR,
 A_FpLong *numberPF);
```

#### AEGP_CauseIdleRoutinesToBeCalled

本例程可以从主线程以外的线程安全调用。
它是异步的，会在空闲处理程序被调用之前返回。
获取该函数指针的套件函数不是线程安全的；将其保存在主线程中供子线程使用。

```cpp
AEGP_CauseIdleRoutinesToBeCalled(void);
```

#### AEGP_GetSuppressInteractiveUI

返回 After Effects 是否在没有用户界面的情况下运行。

```cpp
AEGP_GetSuppressInteractiveUI(
 A_Boolean *ui_is_suppressedPB);
```

#### AEGP_WriteToOSConsole

向操作系统控制台发送一个字符串。

```cpp
AEGP_WriteToOSConsole(
 const A_char *textZ);
```

#### AEGP_WriteToDebugLog

写一条信息到调试日志，如果 After Effects 是用"-debug "选项启动的，则写到操作系统命令行。

```cpp
AEGP_WriteToDebugLog(
 const A_char *subsystemZ,
 const A_char *event_typeZ,
 const A_char *infoZ);
```

#### AEGP_GetLastErrorMessage

检索显示给用户的最后一条错误信息，以及相关的错误编号。
传入要返回的字符缓冲区的大小。

```cpp
AEGP_GetLastErrorMessage(
 A_long buffer_size,
 A_char *error_string,
 A_Err *error_num);
```

#### AEGP_IsScriptingAvailable

如果脚本对插件可用，返回`TRUE`。一般与`AEGP_ExecuteScript`配合使用(见下)

```cpp
AEGP_IsScriptingAvailable(
 A_Boolean *outAvailablePB);
```

#### AEGP_ExecuteScript

执行一个脚本。传入的脚本可以是 UTF-8 或当前应用程序的编码( platform_encodingB为 TRUE)。两个输出参数是可选的。脚本最后一行的值就是传回 outResultPH0 的值。

```cpp
AEGP_ExecuteScript(
 AEGP_PluginID inPlugin_id,   // 插件ID
 const A_char *inScriptZ,     // 脚本内容
 const A_Boolean platform_encodingB,  // 是否支持当前平台编码
 AEGP_MemHandle *outResultPH0,    // 返回结果
 AEGP_MemHandle *outErrStringPH0  // 报错信息
);
```

示例: 一个AE弹窗(随便丢在Render函数里了,在自己函数里调用suites即可)

```cpp
static PF_Err 
Render (
  PF_InData  *in_data,
  PF_OutData  *out_data,
  PF_ParamDef  *params[],
  PF_LayerDef  *output )
{
  PF_Err    err  = PF_Err_NONE;
  AEGP_SuiteHandler suites(in_data->pica_basicP); // 调用suites
  //   ... 以上为 Skeleton.cpp Render 函数自带


  A_Boolean outAvailablePB;
  AEGP_MemHandle outResultPH;
  AEGP_MemHandle outErrorStringPH;

  // 判断插件能否使用脚本
  ERR(suites.UtilitySuite4()->AEGP_IsScriptingAvailable(&outAvailablePB));
  ERR(suites.UtilitySuite4()->AEGP_ExecuteScript(NULL, "alert('我是一个弹窗');", true, &outResultPH, &outErrorStringPH));

  // 使用MemHandle后 需要释放掉
  ERR(suites.MemorySuite1()->AEGP_FreeMemHandle(outResultPH));
  ERR(suites.MemorySuite1()->AEGP_FreeMemHandle(outErrorStringPH));
}
```

#### AEGP_HostIsActivated

如果用户成功激活了 After Effects，返回 TRUE。

```cpp
AEGP_HostIsActivated(
 A_Boolean *is_activatedPB);
```

#### AEGP_GetPluginPlatformRef

在 macOS 上，返回一个`CFBundleRef`给你的 Mach-O 插件，如果是 CFM 插件则返回 NULL。
在 Windows 上总是返回`NULL`(你可以使用操作系统特定的入口函数来捕获你的 DLLInstance)。

```cpp
AEGP_GetPluginPlatformRef(
 AEGP_PluginID plug_id,
 void **plat_refPPV);
```

#### AEGP_UpdateFontList

重新扫描系统字体列表。

```cpp
AEGP_UpdateFontList();
```

#### AEGP_GetPluginPaths

在 CC 中是新的。返回一个与插件相关的特定路径。

- `AEGP_GetPathTypes_PLUGIN` - (Not Implemented) The path to the location of the plug-in itself.
- `AEGP_GetPathTypes_USER_PLUGIN` -The suite specific location of user specific plug-ins.
- `AEGP_GetPathTypes_ALLUSER_PLUGIN` - The suite specific location of plug-ins shared by all users.
- `AEGP_GetPathTypes_APP` - The After Effects .exe or .app location. Not plug-in specific.

```cpp
AEGP_GetPluginPaths(
 AEGP_PluginID aegp_plugin_id,
 AEGP_GetPathTypes path_type
 AEGP_MemHandle *unicode_pathPH);
```

## Persistent Data Suite

插件可以读写 After Effects 偏好中的持久性数据。AEGPs 可以使用以下套件添加和管理自己的持久性数据。数据条目由(section key, value key)对访问。建议插件使用它们的匹配名称作为它们的部分键，如果使用多个部分键，则作为前缀。

可用的数据类型有`A_long`，`A_FpLong`，字符串和`void*`。`A_FpLong`以小数点后 6 位的精度存储。没有规定可以指定不同的精度。字符串数据支持完整的 8 位空间。只有 0x00 被保留给字符串结束。这使得它们成为存储 UTF-8 编码的字符串、ISO 8859-1 和普通 ASCII 的理想选择。节段键和值键都是这种类型的。对于没有被提供的简单数据类型所代表的数据类型，使用包含你的自定义数据的数据句柄。 void\*非结构化数据允许你存储任何类型的数据。你必须在传递数据的同时传递一个以字节为单位的大小。

当调用任何一个函数来检索一个键的值时，如果没有找到给定的键，默认值都会被写入 blob 并作为值返回；如果没有提供默认值，将写入并返回一个空白值。

请注意，这些数据是存储在应用程序的首选项中，而不是在项目中。从 6.5 开始，没有办法在 After Effects 项目中存储不透明的 AEGP 生成的数据。

After Effects 可以处理在应用过程中改变首选项的插件；它在对可首选项的设置采取行动之前会检查首选项的内存拷贝，而不是依赖保存的首选项。就像我计划好的一样，或者其他什么东西

### AEGP_PersistentDateSuite4

#### AEGP_GetApplicationBlob

获取所有持久化应用程序数据的句柄。修改它将修改应用程序。
AEGP_PersistentType "参数在 CC 中是新的，应该被设置为以下内容之一。

- `AEGP_PersistentType_MACHINE_SPECIFIC`,
- `AEGP_PersistentType_MACHINE_INDEPENDENT`,
- `AEGP_PersistentType_MACHINE_INDEPENDENT_RENDER`,
- `AEGP_PersistentType_MACHINE_INDEPENDENT_OUTPUT`,
- `AEGP_PersistentType_MACHINE_INDEPENDENT_COMPOSITION`
- `AEGP_PersistentType_MACHINE_SPECIFIC_TEXT`,
- `AEGP_PersistentType_MACHINE_SPECIFIC_PAINT`

```cpp
AEGP_GetApplicationBlob(
 AEGP_PersistentType blob_type,
 AEGP_PersistentBlobH *blobPH);
```

#### AEGP_GetNumSections

获取应用程序 blob 中的部分数量。

```cpp
AEGP_GetNumSections(
 AEGP_PersistentBlobH blobH,
 A_long *num_sectionPL);
```

#### AEGP_GetSectionKeyByIndex

获取给定索引处的键。

```cpp
AEGP_GetSectionKeyByIndex(
 AEGP_PersistentBlobH blobH,
 A_long section_index,
 A_long max_section_size,
 A_char *section_keyZ);
```

#### AEGP_DoesKeyExist

返回 blob 中是否存在给定的键/值对。

```cpp
AEGP_DoesKeyExist(
 AEGP_PersistentBlobH blobH,
 const A_char *section_keyZ,
 const A_char *value_keyZ,
 A_Boolean *existsPB);
```

#### AEGP_GetNumKeys

检索该部分中的值键的数量。

```cpp
AEGP_GetNumKeys(
 AEGP_PersistentBlobH blobH,
 const A_char *section_keyZ,
 A_long *num_keysPL);
```

#### AEGP_GetValueKeyByIndex

检索索引的键的值。

```cpp
AEGP_GetValueKeyByIndex(
 AEGP_PersistentBlobH blobH,
 const A_char *section_keyZ,
 A_long key_index,
 A_long max_key_size,
 A_char *value_keyZ);
```

::: tip

对于下面的函数，如果没有找到给定的键，默认值会同时写入 blob 中并作为值返回。
如果没有提供默认值，将写入并返回一个空白值。
:::

#### AEGP_GetDataHandle

获取与给定章节的键相关的值。如果使用内存中的数据结构，注意 endian 问题。

```cpp
AEGP_GetDataHandle(
 AEGP_PluginID plugin_id,
 AEGP_PersistentBlobH blobH,
 const A_char *section_keyZ,
 const A_char *value_keyZ,
 AEGP_MemHandle defaultH0,
 AEGP_MemHandle *valuePH);
```

#### AEGP_GetData

获取位于给定部分的值的数据。

```cpp
AEGP_GetData(
 AEGP_PersistentBlobH blobH,
 const A_char *section_keyZ,
 const A_char *value_keyZ,
 A_u_long data_sizeLu,
 const void *defaultPV0,
 void *bufPV);
```

#### AEGP_GetString

获取给定部分键值的字符串(并以`actual_szLu0`表示其长度)。

```cpp
AEGP_GetString(
 AEGP_PersistentBlobH blobH,
 const A_char *section_keyZ,
 const A_char *value_keyZ,
 const A_char *defaultZ0,
 A_u_long buf_sizeLu,
 char *bufZ,
 A_u_long *actual_szLu0);
```

#### AEGP_GetLong

获取与给定章节键值相关的`A_long`。

```cpp
AEGP_GetLong(
 AEGP_PersistentBlobH blobH,
 const A_char *section_keyZ,
 const A_char *value_keyZ,
 A_long defaultL,
 A_long *valuePL);
```

#### AEGP_GetFpLong

获取与给定章节键值相关的 "A_FpLong"。

```cpp
AEGP_GetFpLong(
 AEGP_PersistentBlobH blobH,
 const A_char *section_keyZ,
 const A_char *value_keyZ,
 A_FpLong defaultF,
 A_FpLong *valuePF);
```

#### AEGP_GetTime

在 CC 中新增。获取与给定部分键值相关的`A_Time`。

```cpp
AEGP_GetTime(
 AEGP_PersistentBlobH blobH,
 const A_char *section_keyZ,
 const A_char *value_keyZ,
 const A_Time *defaultPT0,
 A_Time *valuePT);
```

#### AEGP_GetARGB

在 CC 中新增。获取与给定章节键值相关的`PF_PixelFloat`。

```cpp
AEGP_GetARGB(
 AEGP_PersistentBlobH blobH,
 const A_char *section_keyZ,
 const A_char *value_keyZ,
 const PF_PixelFloat *defaultP0,
 PF_PixelFloat *valueP);
```

#### AEGP_SetDataHandle

将给定的部分键值设置为传入的句柄。

```cpp
AEGP_SetDataHandle(
 AEGP_PersistentBlobH blobH,
 const A_char *section_keyZ,
 const A_char *value_keyZ,
 const AEGP_MemHandle valueH);
```

#### AEGP_SetData

将给定的部分键值设置为`dataPV`中的数据。

```cpp
AEGP_SetData(
 AEGP_PersistentBlobH blobH,
 const A_char *section_keyZ,
 const A_char *value_keyZ,
 A_u_long data_sizeLu,
 const void *dataPV);
```

#### AEGP_SetString

将给定章节键的字符串设置为`strZ`。

```cpp
AEGP_SetString(
 AEGP_PersistentBlobH blobH,
 const A_char *section_keyZ,
 const A_char *value_keyZ,
 const A_char *strZ);
```

#### AEGP_SetLong

将给定章节的键值设置为`valueL`。

```cpp
AEGP_SetLong(
 AEGP_PersistentBlobH blobH,
 const A_char *section_keyZ,
 const A_char *value_keyZ,
 A_long valueL);
```

#### AEGP_SetFpLong

将给定章节的键值设置为`valueF`。

```cpp
AEGP_SetFpLong(
 AEGP_PersistentBlobH blobH,
 const A_char *section_keyZ,
 const A_char *value_keyZ,
 A_FpLong valueF);
```

#### AEGP_SetTime

CC 中的新内容。将给定章节的键值设置为`valuePT`。

```cpp
AEGP_SetTime(
 AEGP_PersistentBlobH blobH,
 const A_char *section_keyZ,
 const A_char *value_keyZ,
 A_Time *valuePT);
```

#### AEGP_SetARGB

在 CC 中新增。将给定章节的键值设置为 "valueP"。

```cpp
AEGP_SetARGB(
 AEGP_PersistentBlobH blobH,
 const A_char *section_keyZ,
 const A_char *value_keyZ,
 PF_PixelFloat *valueP);
```

#### AEGP_DeleteEntry

从 blob 中删除给定部分的值。

```cpp
AEGP_DeleteEntry(
 AEGP_PersistentBlobH blobH,
 const A_char *section_keyZ,
 const A_char *value_keyZ);
```

#### AEGP_GetPrefsDirectory

获取包含 After Effects 偏好文件的文件夹的路径。
路径是一个以 NULL 结尾的 A_UTF16Char 字符串的句柄，必须用`AEGP_FreeMemHandle`处理。

```cpp
AEGP_GetPrefsDirectory(
 AEGP_MemHandle *unicode_pathPH);
```

## Color Management

我们提供了一个函数，以便 AEGP 可以获得 After Effects 当前颜色管理设置的信息。

### AEGP_ColorSettingsSuite2

#### AEGP_GetBlendingTables

检索当前不透明的`PF_EffectBlendingTables`，供`AEGP_TransferRect`使用。

```cpp
AEGP_GetBlendingTables(
 PR_RenderContextH render_contextH,
 PF_EffectBlendingTables *blending_tables);
```

#### AEGP_DoesViewHaveColorSpaceXform

返回是否有颜色空间转换应用到当前项目视图。

```cpp
AEGP_DoesViewHaveColorSpaceXform(
 AEGP_ItemViewP viewP,
 A_Boolean *has_xformPB);
```

#### AEGP_XformWorkingToViewColorSpace

将源的视图色彩空间改变为目标的工作色彩空间。
源点和终点可以是相同的。

```cpp
AEGP_XformWorkingToViewColorSpace(
 AEGP_ItemViewP viewP,
 AEGP_WorldH srcH,
 AEGP_WorldH dstH);
```

#### AEGP_GetNewWorkingSpaceColorProfile

检索不透明的当前工作空间 ICC 配置文件。必须被处理掉。
名称中的 "New "并不表示你在编造一个新的配置文件；相反，它是我们函数命名标准的一部分。
名称中带有 "New "的东西分配了一些东西，调用者必须处理掉。

```cpp
AEGP_GetNewWorkingSpaceColorProfile(
 AEGP_PluginID aegp_plugin_id,
 AEGP_MemHandle *icc_profPH);
```

#### AEGP_GetNewColorProfileFromICCProfile

从 After Effects 中获取一个新的`AEGP_ColorProfileP`，代表指定的 ICC 配置文件。
调用者必须使用`AEGP_DisposeColorProfile()`来处理返回的`AEGP_ColorProfileP`。

```cpp
AEGP_GetNewColorProfile FromICCProfile(
 AEGP_PluginID aegp_plugin_id,
 A_long icc_sizeL,
 const void *icc_dataPV,
 AEGP_ColorProfileP *profilePP);
```

#### AEGP_GetNewICCProfileFromColorProfile

检索一个新的 ICC 配置文件(存储在`AEGP_MemHandle`中)，代表指定的颜色配置文件。
返回的`AEGP_MemHandle`必须由调用者处理。

```cpp
AEGP_GetNewICCProfile FromColorProfile(
 AEGP_PluginID plugin_id,
 AEGP_ConstColorProfileP profileP,
 AEGP_MemHandle *profilePH);
```

#### AEGP_GetNewColorProfileDescription

返回指定颜色配置文件的文本描述。
文本将是一个空尾的 UTF16 字符串，必须由调用者处理。

```cpp
AEGP_GetNewColorProfileDescription(
 AEGP_PluginID aegp_plugin_id,
 AEGP_ConstColorProfileP profileP,
 AEGP_MemHandle *unicode_descPH);
```

#### AEGP_DisposeColorProfile

处置一个使用本套件中其他函数获得的颜色配置文件。

```cpp
AEGP_DisposeColorProfile(
 AEGP_ColorProfileP profileP);
```

#### AEGP_GetColorProfileApproximateGamma

返回一个近似于指定颜色配置文件所使用的伽马设置的浮点数。

```cpp
AEGP_GetColorProfileApproximateGamma(
 AEGP_ConstColorProfileP profileP,
 A_FpShort *approx_gammaP);
```

#### AEGP_IsRGBColorProfile

返回指定的颜色配置文件是否为 RGB。

```cpp
AEGP_IsRGBColorProfile(
 AEGP_ConstColorProfileP profileP,
 A_Boolean *is_rgbPB);
```

## Render Suites

自从我们引入 AEGP API 以来，我们一直被要求提供检索渲染帧的函数。

这些函数套件使你能够做到这一点。

首先，在[AEGP_RenderOptionsSuite4]或[AEGP_LayerRenderOptionsSuite1]中指定你想渲染的内容。

然后用[AEGP_RenderSuite4]来进行渲染。

### AEGP_RenderOptionsSuite4

#### AEGP_NewFromItem

返回与给定的`AEGP_ItemH`相关的`AEGP_RenderOptionsH`。
如果没有指定选项，After Effects 会传回一个`AEGP_RenderOptionsH`，渲染时间设置为 0，时间步长设置为当前帧长，场渲染设置为`PF_Field_FRAME`，深度设置为项目中指定的最高分辨率。

```cpp
AEGP_NewFromItem(
 AEGP_PluginID plugin_id,
 AEGP_ItemH itemH,
 AEGP_RenderOptionsH *optionsPH);
```

#### AEGP_Duplicate

复制一个`AEGP_RenderOptionsH`到`copyPH`。

```cpp
AEGP_Duplicate(
 AEGP_PluginID plugin_id,
 AEGP_RenderOptionsH optionsH,
 AEGP_RenderOptionsH *copyPH);
```

#### AEGP_Dispose

删除一个`AEGP_RenderOptionsH`。

```cpp
AEGP_Dispose(
 AEGP_RenderOptionsH optionsH);
```

#### AEGP_SetTime

设置一个`AEGP_RenderOptionsH`的渲染时间。

```cpp
AEGP_SetTime(
 AEGP_RenderOptionsH optionsH,
 A_Time time);
```

#### AEGP_GetTime

读取给定的`AEGP_RenderOptionsH`的渲染时间。

```cpp
AEGP_GetTime(
 AEGP_RenderOptionsH optionsH,
 A_Time *timeP);
```

#### AEGP_SetTimeStep

指定引用的`AEGP_RenderOptionsH`的时间步骤(一帧的持续时间)。

```cpp
AEGP_SetTimeStep(
 AEGP_RenderOptionsH optionsH,
 A_Time time_step);
```

#### AEGP_GetTimeStep

检索给定的`AEGP_RenderOptionsH`的时间步数(一帧的持续时间)。

```cpp
AEGP_GetTimeStep(
 AEGP_RenderOptionsH optionsH,
 A_Time *timePT);
```

#### AEGP_SetFieldRender

指定给定的`AEGP_RenderOptionsH`的字段设置。

```cpp
AEGP_SetFieldRender(
 AEGP_RenderOptionsH optionsH,
 PF_Field field_render);
```

#### AEGP_GetFieldRender

检索给定的 "AEGP_RenderOptionsH "的字段设置。

```cpp
AEGP_GetFieldRender(
 AEGP_RenderOptionsH optionsH,
 PF_Field *field_renderP);
```

#### AEGP_SetWorldType

指定给定的`AEGP_RenderOptionsH`的输出的 AEGP_WorldType。

```cpp
AEGP_SetWorldType(
 AEGP_RenderOptionsH optionsH,
 AEGP_WorldType type);
```

`AEGP_WorldType`将是`AEGP_WorldType_8`或`AEGP_WorldType_16`。

#### AEGP_GetWorldType

检索给定的`AEGP_RenderOptionsH`的`AEGP_WorldType`。

```cpp
AEGP_GetWorldType(
 AEGP_RenderOptionsH optionsH,
 AEGP_WorldType *typeP);
```

#### AEGP_SetDownsampleFactor

指定给定的`AEGP_RenderOptionsH`的降采样系数(有独立的水平和垂直设置)。

```cpp
AEGP_SetDownsampleFactor(
 AEGP_RenderOptionsH optionsH,
 A_short x,
 A_short y);
```

#### AEGP_GetDownsampleFactor

检索给定的 "AEGP_RenderOptionsH "的降频采样系数。

```cpp
AEGP_GetDownsampleFactor(
 AEGP_RenderOptionsH optionsH,
 A_short *xP,
 A_short *yP);
```

#### AEGP_SetRegionOfInterest

为给定的 "AEGP_RenderOptionsH "指定兴趣区域的子矩形。

```cpp
AEGP_SetRegionOfInterest(
 AEGP_RenderOptionsH optionsH,
 const A_LRect *roiP)
```

#### AEGP_GetRegionOfInterest

检索给定的 "AEGP_RenderOptionsH "的兴趣区域子矩形。

```cpp
AEGP_GetRegionOfInterest(
 AEGP_RenderOptionsH optionsH,
 A_LRect *roiP);
```

#### AEGP_SetMatteMode

为给定的`AEGP_RenderOptionsH`指定`AEGP_MatteMode`。

```cpp
AEGP_SetMatteMode(
 AEGP_RenderOptionsH optionsH,
 AEGP_MatteMode mode);
```

`AEGP_MatteMode`将是下列之一。

- `AEGP_MatteMode_STRAIGHT`
- `AEGP_MatteMode_PREMUL_BLACK`
- `AEGP_MatteMode_PREMUL_BG_COLOR`

#### AEGP_GetMatteMode

为给定的`AEGP_RenderOptionsH`检索`AEGP_MatteMode`。

```cpp
AEGP_GetMatteMode(
 AEGP_RenderOptionsH optionsH,
 AEGP_MatteMode *modeP);
```

#### AEGP_GetChannelOrder

为给定的`AEGP_RenderOptionsH`获取`AEGP_ChannelOrder`。`AEGP_ChannelOrder`将是`AEGP_ChannelOrder_ARGB`或`AEGP_ChannelOrder_BGRA`。

```cpp
AEGP_GetChannelOrder(
 AEGP_RenderOptionsH optionsH,
 AEGP_ChannelOrder *orderP);
```

事实是：这是为方便与 Premiere Pro 进行实时连接而添加的。

#### AEGP_SetChannelOrder

设置`AEGP_ChannelOrder`的`AEGP_RenderOptionsH`。

```cpp
AEGP_SetChannelOrder(
 AEGP_RenderOptionsH optionsH,
 AEGP_ChannelOrder order);
```

#### AEGP_GetRenderGuideLayers

传回一个布尔值，如果渲染引导层的设置是开启的，则为真。

```cpp
AEGP_GetRenderGuideLayers)(
 AEGP_RenderOptionsH optionsH,
 A_Boolean *will_renderPB);
```

#### AEGP_SetRenderGuideLayers

指定是否要渲染引导层。

```cpp
AEGP_SetRenderGuideLayers)(
 AEGP_RenderOptionsH optionsH,
 A_Boolean render_themB);
```

#### AEGP_GetRenderQuality

获取渲染队列项目的渲染质量。
质量可以是`AEGP_ItemQuality_DRAFT`或`AEGP_ItemQuality_BEST`。

```cpp
AEGP_GetRenderQuality)(
 AEGP_RenderOptionsH optionsH,
 AEGP_ItemQuality *qualityP);
```

#### AEGP_SetRenderQuality

设置渲染队列项目的渲染质量。

```cpp
AEGP_GetRenderQuality)(
 AEGP_RenderOptionsH optionsH,
 AEGP_ItemQuality quality);
```

### AEGP_LayerRenderOptionsSuite1

::: tip

13.0 中的新内容
:::

#### AEGP_NewFromLayer

返回与给定的`AEGP_LayerH`相关的`AEGP_LayerRenderOptionsH`。 渲染时间设置为该层的当前时间，时间步长设置为该层的帧持续时间。 ROI设置为该层的名义边界，EffectsToRender设置为 "所有"。`optionsPH`必须由调用代码处理。

```cpp
AEGP_NewFromLayer(
 AEGP_PluginID plugin_id,
 AEGP_LayerH layerH,
 AEGP_LayerRenderOptionsH *optionsPH);
```

#### AEGP_NewFromUpstreamOfEffect

返回与给定的`AEGP_EffectRefH`相关的层的`AEGP_LayerRenderOptionsH`。
渲染时间被设置为该层的当前时间，时间步长被设置为该层的帧长。
ROI 为图层的标称边界，EffectToRender 为`effectH`的索引。
`optionsPH`必须由调用代码处理。

```cpp
AEGP_NewFromUpstreamOfEffect(
 AEGP_PluginID plugin_id,
 AEGP_EffectRefH effectH,
 AEGP_LayerRenderOptionsH *optionsPH);
```

#### AEGP_Duplicate

复制一个`AEGP_LayerRenderOptionsH`到`copyPH`。

```cpp
AEGP_Duplicate(
 AEGP_PluginID plugin_id,
 AEGP_LayerRenderOptionsH optionsH,
 AEGP_LayerRenderOptionsH *copyPH);
```

#### AEGP_Dispose

删除一个`AEGP_LayerRenderOptionsH`。

```cpp
AEGP_Dispose(
 AEGP_LayerRenderOptionsH optionsH);
```

#### AEGP_SetTime

设置`AEGP_LayerRenderOptionsH`的渲染时间。

```cpp
AEGP_SetTime(
 AEGP_LayerRenderOptionsH optionsH,
 A_Time time);
```

#### AEGP_GetTime

读取给定的`AEGP_LayerRenderOptionsH`的渲染时间。

```cpp
AEGP_GetTime(
 AEGP_LayerRenderOptionsH optionsH,
 A_Time *timeP);
```

#### AEGP_SetTimeStep

指定引用的`AEGP_LayerRenderOptionsH`的时间步长(一帧的持续时间)。

```cpp
AEGP_SetTimeStep(
 AEGP_LayerRenderOptionsH optionsH,
 A_Time time_step);
```

#### AEGP_GetTimeStep

检索给定的`AEGP_LayerRenderOptionsH`的时间步骤(一帧的持续时间)。

```cpp
AEGP_GetTimeStep(
 AEGP_LayerRenderOptionsH optionsH,
 A_Time *timePT);
```

#### AEGP_SetWorldType

指定给定的`AEGP_LayerRenderOptionsH`的输出的 AEGP_WorldType。

```cpp
AEGP_SetWorldType(
 AEGP_LayerRenderOptionsH optionsH,
 AEGP_WorldType type);
```

`AEGP_WorldType`将是`AEGP_WorldType_8`或`AEGP_WorldType_16`。

#### AEGP_GetWorldType

检索给定的`AEGP_LayerRenderOptionsH`的 AEGP_WorldType。

```cpp
AEGP_GetWorldType(
 AEGP_LayerRenderOptionsH optionsH,
 AEGP_WorldType *typeP);
```

#### AEGP_SetDownsampleFactor

指定给定的`AEGP_LayerRenderOptionsH`的降采样系数(有独立的水平和垂直设置)。

```cpp
AEGP_SetDownsampleFactor(
 AEGP_LayerRenderOptionsH optionsH,
 A_short x,
 A_short y);
```

#### AEGP_GetDownsampleFactor

读取给定的 "AEGP_LayerRenderOptionsH "的降频采样系数。

```cpp
AEGP_GetDownsampleFactor(
 AEGP_LayerRenderOptionsH optionsH,
 A_short *xP,
 A_short *yP);
```

#### AEGP_SetMatteMode

为给定的`AEGP_LayerRenderOptionsH`指定 AEGP_MatteMode。

```cpp
AEGP_SetMatteMode(
 AEGP_LayerRenderOptionsH optionsH,
 AEGP_MatteMode mode);
```

AEGP_MatteMode 将是下列之一。

- `AEGP_MatteMode_STRAIGHT`
- `AEGP_MatteMode_PREMUL_BLACK`
- `AEGP_MatteMode_PREMUL_BG_COLOR`

#### AEGP_GetMatteMode

为给定的`AEGP_LayerRenderOptionsH`检索 AEGP_MatteMode。

```cpp
AEGP_GetMatteMode(
 AEGP_LayerRenderOptionsH optionsH,
 AEGP_MatteMode *modeP);
```

### AEGP_RenderSuite4

#### AEGP_RenderAndCheckoutFrame

检索请求的帧的`AEGP_FrameReceiptH`(不是实际的像素)。
使用`AEGP_CheckinFrame`签入此收据以释放内存。
使用[AEGP_RenderOptionsSuite4] 创建`AEGP_RenderOptionsH`。
可选的是，AEGP 可以传递一个函数，在用户取消当前渲染时由 After Effects 调用，以及一个 refcon(对不透明数据的常量引用)，供该函数使用。

```cpp
AEGP_RenderAndCheckoutFrame(
 AEGP_RenderOptionsH optionsH,
 AEGP_RenderSuiteCheckForCancel cancel_functionP0,
 AEGP_CancelRefcon cancel_function_refconP0,
 AEGP_FrameReceiptH *receiptPH);
```

#### AEGP_RenderAndCheckoutLayerFrame

CC 2014 的新功能。这允许在非渲染时间对应用了效果的图层进行帧检。

这对于需要帧的操作很有用，例如，当一个按钮被点击时，在渲染时等待一会儿是可以接受的。

注意：由于它不是异步的，所以它不能解决自定义 UI 需要根据框架来绘制的一般问题。

为请求的图层框架检索一个`AEGP_FrameReceiptH`(而不是实际的像素)。使用`AEGP_CheckinFrame`签入此收据以释放内存。

使用`AEGP_NewFromUpstreamOfEffect()`创建`AEGP_LayerRenderOptionsH`，在AEGP_LayerRenderOptionsSuite1。

实际上，你可以使用`AEGP_NewFromLayer()`来获得其他图层参数的图层，并应用其效果。
然而，要小心。如果你在你的效果 A 中这样做，而在另一个图层上有一个效果 B，在渲染时做同样的事情，你就会产生一个无限循环。
如果你不是为了渲染而这样做，那么它可能是好的。
如果用户取消了当前的渲染，AEGP 可以传递给 After Effects 一个函数，供其调用，同时也可以传递一个 refcon(对不透明数据的常数引用)，供该函数使用。

```cpp
AEGP_RenderAndCheckoutLayerFrame(
 AEGP_LayerRenderOptionsH optionsH,
 A_Boolean render_plain_layer_frameB,
 AEGP_RenderSuiteCheckForCancel cancel_functionP0,
 AEGP_CancelRefcon cancel_function_refconP0,
 AEGP_FrameReceiptH *receiptPH);
```

#### AEGP_CheckinFrame

一旦你的 AEGP 完成了对帧的访问，就立即调用这个函数。
After Effects 会根据哪些帧被检出而做出缓存决定，所以不要占用它们!

```cpp
AEGP_CheckinFrame(
 AEGP_FrameReceiptH receiptH);
```

#### AEGP_GetReceiptWorld

检索与引用的`AEGP_FrameReceiptH`相关的像素(`AEGP_WorldH`)。

```cpp
AEGP_GetReceiptWorld(
 AEGP_FrameReceiptH receiptH,
 AEGP_WorldH *worldPH);
```

#### AEGP_GetRenderedRegion

检索一个`A_LRect`，包含`AEGP_FrameReceiptH`的`AEGP_WorldH`中已经被渲染的区域。
记住，只有图像中被改变的部分才有可能被渲染，所以能够检查是否包括你需要的部分是很重要的。

```cpp
AEGP_GetRenderedRegion(
 AEGP_FrameReceiptH receiptH,
 A_LRect *regionP);
```

#### AEGP_IsRenderedFrameSufficient

给定两组`AEGP_RenderOptionsH`，如果已经渲染的像素对提议的`AEGP_RenderOptionsH`仍然有效，After Effects 将返回`TRUE`。

```cpp
AEGP_IsRenderedFrameSufficient(
 AEGP_RenderOptionsH rendered_optionsH,
 AEGP_RenderOptionsH proposed_optionsH,
 A_Boolean *is_sufficientPB);
```

#### AEGP_RenderNewItemSoundData

获得一个`AEGP_ItemH`的音频，在给定的时间，给定的持续时间，以给定的格式。 插件必须处理返回的`AEGP_SoundDataH`(如果没有音频，它可能是 NULL)。

```cpp
AEGP_RenderNewItemSoundData(
 AEGP_ItemH itemH,
 const A_Time *start_timePT,
 const A_Time *durationPT,
 const AEGP_SoundDataFormat *formatP,
 AEGP_SoundDataH *new_dataPH);
```

注意：这个函数，如果作为`AEGP_ItemSuite2`的一部分被调用，提供了一个可通过鼠标点击中断的渲染，与这里发布的`AEGP_RenderSuite`版本不同。

#### AEGP_GetCurrentTimestamp

检索项目的当前`AEGP_TimeStamp`。
AEGP_TimeStamp "在项目被触及并影响渲染时被更新。

```cpp
AEGP_GetCurrentTimestamp(
 AEGP_TimeStamp *time_stampP);
```

#### AEGP_HasItemChangedSinceTimestamp

返回 AEGP_ItemH 的视频在给定的`AEGP_TimeStamp`后是否有变化。
注意：这并不跟踪音频的变化。

```cpp
AEGP_HasItemChangedSinceTimestamp(
 AEGP_ItemH itemH,
 const A_Time *start_timeP,
 const A_Time *durationP,
 const AEGP_TimeStamp *time_stampP,
 A_Boolean *changedPB);
```

#### AEGP_IsItemWorthwhileToRender

返回该帧是否值得在外部渲染并检查到缓存中。
一个投机性的渲染器应该检查两次：在把帧送出去渲染之前，以及在完成之后，在调用`AEGP_NewPlatformWorld()`和签入之前。
这个函数要和`AEGP_HasItemChangedSinceTimestamp()`一起使用，不能单独使用。

```cpp
AEGP_IsItemWorthwhileToRender(
 AEGP_RenderOptionsH roH,
 const AEGP_TimeStamp *time_stampP,
 A_Boolean *worthwhilePB);
```

#### AEGP_CheckinRenderedFrame

向 After Effects 提供一个渲染的框架(`AEGP_PlatformWorldH`)，After Effects 会采用它。
`ticksL`是渲染框架所需的大致时间。

```cpp
AEGP_CheckinRenderedFrame(
 AEGP_RenderOptionsH roH,
 const AEGP_TimeStamp* time_stampP,
 A_u_long ticksL,
 AEGP_PlatformWorldH imageH);
```

#### AEGP_GetReceiptGuid

CS6 的新功能。检索一个已渲染的帧的 GUID。传回的内存句柄必须被处理掉。

```cpp
AEGP_GetReceiptGuid(
 AEGP_FrameReceiptH receiptH,
 AEGP_MemHandle *guidMH)
```

## The AEGP_World As We Know It

`AEGP_Worlds`是整个 AEGP APIs 用来描述像素帧的通用格式。

### AEGP_WorldSuite3

#### AEGP_New

返回一个分配的、初始化的`AEGP_WorldH`。

```cpp
AEGP_New(
 AEGP_PluginID plugin_id,
 AEGP_WorldType type,
 A_long widthL,
 A_long heightL,
 AEGP_WorldH *worldPH);
```

#### AEGP_Dispose

弃置一个`AEGP_WorldH`。在你分配的每个世界上使用这个。

```cpp
AEGP_Dispose(
 AEGP_WorldH worldH);
```

#### AEGP_GetType

返回一个给定的`AEGP_WorldH`的类型。

```cpp
AEGP_GetType(
 AEGP_WorldH worldH,
 AEGP_WorldType **typeP);
```

AEGP_WorldType 将是以下之一。

- `AEGP_WorldType_8`,
- `AEGP_WorldType_16`,
- `AEGP_WorldType_32`

#### AEGP_GetSize

返回给定的`AEGP_WorldH`的宽度和高度。

```cpp
AEGP_GetSize(
 AEGP_WorldH worldH,
 A_long *widthPL,
 A_long *heightPL);
```

#### AEGP_GetRowBytes

Returns the rowbytes for the given `AEGP_WorldH`. 返回给定`AEGP_WorldH`的行数。

```cpp
AEGP_GetRowBytes(
 AEGP_WorldH worldH,
 A_u_long *row_bytesPL);
```

#### AEGP_GetBaseAddr8

返回`AEGP_WorldH`的基本地址，以便在像素迭代函数中使用。
如果在非 8bpc 世界中使用，将返回错误。

```cpp
AEGP_GetBaseAddr8(
 AEGP_WorldH worldH,
 PF_Pixel8 **base_addrP);
```

#### AEGP_GetBaseAddr16

返回`AEGP_WorldH`的基本地址，用于像素迭代函数。
如果在非 16bpc 的世界中使用，将返回错误。

```cpp
AEGP_GetBaseAddr16(
 AEGP_WorldH worldH,
 PF_Pixel16 **base_addrP);
```

#### AEGP_GetBaseAddr32

返回`AEGP_WorldH`的基本地址，用于像素迭代函数。
如果在非 32bpc 的世界中使用，将返回错误。

```cpp
AEGP_GetBaseAddr32(
 AEGP_WorldH worldH,
 PF_PixelFloat **base_addrP);
```

#### AEGP_FillOutPFEffectWorld

填充并返回一个代表给定的`AEGP_WorldH`的 PF_EffectWorld，用于许多像素处理回调。
注意：这并没有给你的插件提供所引用的世界的所有权；只有在你分配了源`AEGP_WorldH`的情况下才会销毁它。
它只是填补了所提供的`PF_EffectWorld`，使其指向同一个像素缓冲区。

```cpp
AEGP_FillOutPFEffectWorld(
 AEGP_WorldH worldH,
 PF_EffectWorld *pf_worldP);
```

#### AEGP_FastBlur

对给定的`AEGP_WorldH`进行快速模糊处理。

```cpp
AEGP_FastBlur(
 A_FpLong radiusF,
 PF_ModeFlags mode,
 PF_Quality quality,
 AEGP_WorldH worldH);
```

#### AEGP_NewPlatformWorld

创建一个新的`AEGP_PlatformWorldH`(一个执行平台上的像素世界)。

```cpp
AEGP_NewPlatformWorld(
 AEGP_PluginID plugin_id,
 AEGP_WorldType type,
 A_long widthL,
 A_long heightL,
 AEGP_PlatformWorldH *worldPH);
```

#### AEGP_DisposePlatformWorld

处置一个`AEGP_PlatformWorldH`。

```cpp
AEGP_DisposePlatformWorld(
 AEGP_PlatformWorldH worldH);
```

#### AEGP_NewReferenceFromPlatformWorld

检索指向给定的`AEGP_PlatformWorldH`的 AEGP_WorldH。
注意：这并不分配一个新的世界，它只是提供一个对现有世界的引用。

```cpp
AEGP_NewReferenceFromPlatformWorld(
 AEGP_PluginID plugin_id,
 AEGP_PlatformWorldH plat_worldH,
 AEGP_WorldH *worldPH);
```

## Track Mattes and Transform Functions

使用`AEGP_CompositeSuite`来复制像素世界，操作轨道哑光，并应用转移函数。

### AEGP_CompositeSuite2

#### AEGP_ClearAlphaExceptRect

对于给定的`PF_EffectWorld`，除了指定的矩形外，将 alpha 设置为完全透明。

```cpp
AEGP_ClearAlphaExceptRect(
 A_Rect *clipped_dst_rectPR,
 PF_EffectWorld *dstP);
```

#### AEGP_PrepTrackMatte

用 src_masks 中描述的`PF_Pixel`对`PF_EffectWorld`中的像素进行消光处理，将输出结果放入一个像素数组 dst_mask 中。
注意：与 After Effects 提供的其他大多数像素处理函数不同，这个函数不接受`PF_EffectWorld`参数；相反，你可以简单地从`PF_EffectWorld`中传递数据指针。
这可能会引起混淆，但作为奖励，该函数对输出进行了适当的填充，使`num_pix`像素总是被输出。

```cpp
AEGP_PrepTrackMatte(
 A_long num_pix,
 A_Boolean deepB,
 const PF_Pixel *src_mask,
 PF_MaskFlags mask_flags,
 PF_Pixel *dst_mask);
```

#### AEGP_TransferRect

使用混合模式混合两个`PF_EffectWorlds`，有一个可选的掩码。 在`blend_tablesP0`参数中传递 NULL，以在当前的工作色彩空间中执行混合。

```cpp
AEGP_TransferRect(
 PF_Quality quality,
 PF_ModeFlags m_flags,
 PF_Field field,
 const A_Rect *src_rec,
 const PF_EffectWorld *src_world,
 const PF_CompositeMode *comp_mode,
 PF_EffectBlendingTables blend_tablesP0,
 const PF_MaskWorld *mask_world0,
 A_long dest_x,
 A_long dest_y,
 PF_EffectWorld *dst_world);
```

#### AEGP_CopyBits_LQ

从一个`PF_EffectWorld`复制一个像素矩形(传递一个`NULL`矩形来获得所有像素)到另一个`PF_EffectWorld`，质量不高。

```cpp
AEGP_CopyBits_LQ(
 PF_EffectWorld *src_worldP,
 A_Rect *src_r,
 A_Rect *dst_r,
 PF_EffectWorld *dst_worldP);
```

#### AEGP_CopyBits_HQ_Straight

从一个`PF_EffectWorld`复制一个像素矩形(通过一个`NULL`矩形来获得所有像素)到另一个`PF_EffectWorld`，质量很高，有一个直的 alpha 通道。

```cpp
AEGP_CopyBits_HQ_Straight(
 PF_EffectWorld *src,
 A_Rect *src_r,
 A_Rect *dst_r,
 PF_EffectWorld *dst);
```

#### AEGP_CopyBits_HQ_Premul

从一个`PF_EffectWorld`复制一个像素矩形(传递一个`NULL`矩形来获得所有像素)到另一个`PF_EffectWorld`，质量很高，预先乘以 alpha 通道。

```cpp
AEGP_CopyBits_HQ_Premul(
 PF_EffectWorld *src,
 A_Rect *src_r,
 A_Rect *dst_r,
 PF_EffectWorld *dst);
```

## Work With Audio

`AEGP_SoundDataSuite`允许 AEGPs 获得并操作与合成和镜头项目相关的音频。

使用`AEGP_RenderNewItemSoundData()`可以将纯音频的项目添加到渲染队列中。

### AEGP_SoundDateSuite1

#### AEGP_NewSoundData

创建一个新的`AEGP_SoundDataH`，插件必须处理它。

```cpp
AEGP_NewSoundData(
 const AEGP_SoundDataFormat *formatP,
 AEGP_SoundDataH *new_dataPH);
```

#### AEGP_DisposeSoundData

释放一个`AEGP_SoundDataH`。

```cpp
AEGP_DisposeSoundData(
 AEGP_SoundDataH sound_dataH);
```

#### AEGP_GetSoundDataFormat

获取关于给定的`AEGP_SoundDataH`的格式信息。

```cpp
AEGP_GetSoundDataFormat(
 AEGP_SoundDataH soundH,
 AEGP_SoundDataFormat *formatP);
```

#### AEGP_LockSoundDataSamples

锁定内存中的`AEGP_SoundDataH`。

```cpp
AEGP_LockSoundDataSamples(
 AEGP_SoundDataH soundH,
 void **samples);
```

#### AEGP_UnlockSoundDataSamples

解锁一个`AEGP_SoundDataH`。

```cpp
AEGP_UnlockSoundDataSamples(
 AEGP_SoundDataH soundH);
```

#### AEGP_GetNumSamples

获取给定的`AEGP_SoundDataH`中的样本数。

```cpp
AEGP_GetNumSamples(
 AEGP_SoundDataH soundH,
 A_long *numsamplesPL);
```

## Audio Settings

音频渲染设置用`AEGP_SoundDataFormat`表示。

```cpp
struct AEGP_SoundDataFormat {
 A_FpLong sample_rateF;
 AEGP_SoundEncoding encoding;
 A_long bytes_per_sampleL;
 A_long num_channelsL; // 1 for mono, 2 for stereo
} AEGP_SoundDataFormat;
```

`bytes_per_sampleL`总是`1`,`2`, 或 `4`, 如果指定浮动编码则忽略。

`AEGP_SoundEncoding`是以下之一。

> - `AEGP_SoundEncoding_UNSIGNED_PCM`。
> - `AEGP_SoundEncoding_SIGNED_PCM`。
> - `AEGP_SoundEncoding_FLOAT`。

## Render Queue Suite

该套件允许 AEGPs 向渲染队列中添加项目(使用默认选项)，并控制渲染队列的基本状态。

### AEGP_RenderQueueSuite1

#### AEGP_AddCompToRenderQueue

使用默认选项，向渲染队列中添加一个合成。

```cpp
AEGP_AddCompToRenderQueue(
 AEGP_CompH compH,
 const A_char* pathZ);
```

#### AEGP_SetRenderQueueState

将渲染队列设置为三种有效状态之一。不可能从停止状态转为暂停状态。

```cpp
AEGP_SetRenderQueueState(
 AEGP_RenderQueueState state);
```

- `AEGP_RenderQueueState_STOPPED`
- `AEGP_RenderQueueState_PAUSED`
- `AEGP_RenderQueueState_RENDERING`

#### AEGP_GetRenderQueueState

获取当前的渲染队列状态。

```cpp
AEGP_GetRenderQueueState(
 AEGP_RenderQueueState *stateP);
```

## Render Queue Item Suite

使用这个套件操纵渲染队列项目的所有方面。

### AEGP_RQItemSuite4

#### AEGP_GetNumRQItems

返回当前在渲染队列中的项目数量。

```cpp
AEGP_GetNumRQItems(
 A_long *num_itemsPL);
```

#### AEGP_GetRQItemByIndex

返回一个`AEGP_RQItemRefH`，引用索引的项目。

```cpp
AEGP_GetRQItemByIndex(
 A_long rq_item_index,
 AEGP_RQItemRefH *rq_item_refPH);
```

#### AEGP_GetNextRQItem

返回下一个`AEGP_RQItemRefH`，用于迭代目的。 要获得第一个`AEGP_RQItemRefH`，请为`current_rq_itemH`传递`RQ_ITEM_INDEX_NONE`。

```cpp
AEGP_GetNextRQItem(
 AEGP_RQItemRefH current_rq_itemH,
 AEGP_RQItemRefH *next_rq_itemPH);
```

#### AEGP_GetNumOutputModulesForRQItem

返回应用于给定`AEGP_RQItemRefH`的输出模块的数量。

```cpp
AEGP_GetNumOutputModulesForRQItem(
 AEGP_RQItemRefH rq_itemH,
 A_long *num_outmodsPL);
```

#### AEGP_GetRenderState

如果`AEGP_RQItemRefH`被设置为渲染(一旦用户点击了渲染按钮)，返回 TRUE。

```cpp
AEGP_GetRenderState(
 AEGP_RQItemRefH rq_itemH,
 A_Boolean *will_renderPB);
```

#### AEGP_SetRenderState

控制`AEGP_RQItemRefH`是否会在用户下次点击渲染按钮时渲染。
如果在渲染过程中被调用，将返回一个错误。
如果在`AEGP_RenderQueueState`不是`AEGP_RenderQueueState_STOPPED`的情况下尝试调用，该函数将返回：\*`Err_PARAMETER`。

- `Err_RANGE` if you pass a status that is illegal in any case, and
- `Err_PARAMETER` if you try to pass a status that doesn’t make sense (like trying to queue something for which there’s no output path)

```cpp
AEGP_SetRenderState(
 AEGP_RQItemRefH rq_itemH,
 A_Boolean renderB);
```

#### AEGP_GetStartedTime

返回渲染开始的时间(以秒为单位，从 1904 年开始)。

```cpp
AEGP_GetStartedTime(
 AEGP_RQItemRefH rq_itemH,
 A_Time *started_timePT);
```

#### AEGP_GetElapsedTime

返回渲染开始后所经过的时间。

```cpp
AEGP_GetElapsedTime(
 AEGP_RQItemRefH rq_itemH,
 A_Time *render_timePT);
```

#### AEGP_GetLogType

返回引用的`AEGP_RQItemRefH`的日志类型。

```cpp
AEGP_GetLogType(
 AEGP_RQItemRefH rq_itemH,
 AEGP_LogType *logtypeP);
```

`AEGP_LogtType`将有以下值之一。

- `AEGP_LogType_NONE`
- `AEGP_LogType_ERRORS_ONLY`
- `AEGP_LogType_PLUS_SETTINGS`
- `AEGP_LogType_PER_FRAME_INFO`

#### AEGP_SetLogType

指定与引用的`AEGP_RQItemRefH`一起使用的日志类型。

```cpp
AEGP_SetLogType(
 AEGP_RQItemRefH rq_itemH,
 AEGP_LogType logtype);
```

#### AEGP_RemoveOutputModule

从引用的`AEGP_RQItemRefH`中删除指定的输出模块。

```cpp
AEGP_RemoveOutputModule(
 AEGP_RQItemRefH rq_itemH,
 AEGP_OutputModuleRefH outmodH);
```

#### AEGP_GetComment

更新以支持`RQItemSuite4`中的 Unicode，在 14.1 中可用。
检索与引用的`AEGP_RQItemRefH`相关的注释。

```cpp
AEGP_GetComment(
 AEGP_RQItemRefH rq_itemH,
 AEGP_MemHandle *unicodeH);
```

#### AEGP_SetComment

更新为在`RQItemSuite4`中支持 Unicode，在 14.1 中可用。
指定与引用的 "AEGP_RQItemRefH "相关的注释。

```cpp
AEGP_SetComment(
 AEGP_RQItemRefH rq_itemH,
 const A_UTF16Char *commentZ);
```

#### AEGP_GetCompFromRQItem

检索与`AEGP_RQItemRefH`相关的`AEGP_CompH`。

```cpp
AEGP_GetCompFromRQItem(
 AEGP_RQItemRefH rq_itemH,
 AEGP_CompH *compPH);
```

#### AEGP_DeleteRQItem

删除渲染队列项目。可撤销。

```cpp
AEGP_DeleteRQItem(
 AEGP_RQItemRefH rq_itemH);
```

## Render Queue Monitor Suite

CS6 的新功能。这个套件提供了渲染队列管理器需要的所有信息，以弄清渲染中任何一点的情况。

### AEGP_RenderQueueMonitorSuite1

#### AEGP_RegisterListener

注册一组插件定义的函数，由渲染队列调用。
使用 refcon 来传递数据，当你的`AEGP_RQM_FunctionBlock1`中的插件定义的函数以后被调用时，你要使用这些数据。
如果你不需要它，可以把它设置为 NULL。

```cpp
AEGP_RegisterListener(
 AEGP_PluginID aegp_plugin_id,
 AEGP_RQM_Refcon aegp_refconP,
 const AEGP_RQM_FunctionBlock1 *fcn_blockP);
```

`AEGP_RQM_FunctionBlock1`定义如下。

```cpp
struct _AEGP_RQM_FunctionBlock1 {
 A_Err (*AEGP_RQM_RenderJobStarted)(
 AEGP_RQM_BasicData *basic_dataP,
 AEGP_RQM_SessionId jobid);

 A_Err (*AEGP_RQM_RenderJobEnded)(
 AEGP_RQM_BasicData *basic_dataP,
 AEGP_RQM_SessionId jobid);

 A_Err (*AEGP_RQM_RenderJobItemStarted)(
 AEGP_RQM_BasicData *basic_dataP,
 AEGP_RQM_SessionId jobid,
 AEGP_RQM_ItemId itemid);

 A_Err (*AEGP_RQM_RenderJobItemUpdated)(
 AEGP_RQM_BasicData *basic_dataP,
 AEGP_RQM_SessionId jobid,
 AEGP_RQM_ItemId itemid,
 AEGP_RQM_FrameId frameid);

 A_Err (*AEGP_RQM_RenderJobItemEnded)(
 AEGP_RQM_BasicData *basic_dataP,
 AEGP_RQM_SessionId jobid,
 AEGP_RQM_ItemId itemid,
 AEGP_RQM_FinishedStatus fstatus);

 A_Err (*AEGP_RQM_RenderJobItemReportLog)(
 AEGP_RQM_BasicData *basic_dataP,
 AEGP_RQM_SessionId jobid,
 AEGP_RQM_ItemId itemid,
 A_Boolean isError,
 AEGP_MemHandle logbuf);
} AEGP_RQM_FunctionBlock1;
```

`AEGP_RQM_FinishedStatus`将是以下的一个。

- `AEGP_RQM_FinishedStatus_UNKNOWN`,
- `AEGP_RQM_FinishedStatus_SUCCEEDED`,
- `AEGP_RQM_FinishedStatus_ABORTED`,
- `AEGP_RQM_FinishedStatus_ERRED`

AEGP_RQM_BasicData 定义如下。

```cpp
struct _AEGP_RQM_BasicData {
 const struct SPBasicSuite *pica_basicP;
 A_long aegp_plug_id;
 AEGP_RQM_Refcon aegp_refconPV;
} AEGP_RQM_BasicData;
```

#### AEGP_DeregisterListener

从渲染队列中取消注册。

```cpp
AEGP_DeregisterListener(
 AEGP_PluginID aegp_plugin_id,
 AEGP_RQM_Refcon aegp_refconP);
```

#### AEGP_GetProjectName

获取当前的项目名称。
项目名称是一个以 NULL 结尾的 A_UTF16Char 字符串的句柄，必须用`AEGP_FreeMemHandle`处置。

```cpp
AEGP_GetProjectName(
 AEGP_RQM_SessionId sessid,
 AEGP_MemHandle *utf_project_namePH0);
```

#### AEGP_GetAppVersion

获取应用程序的版本。
应用程序的版本是一个以 NULL 结尾的 A_UTF16Char 字符串的句柄，必须用`AEGP_FreeMemHandle`来处理。

```cpp
AEGP_GetAppVersion(
 AEGP_RQM_SessionId sessid,
 AEGP_MemHandle *utf_app_versionPH0);
```

#### AEGP_GetNumJobItems

获取工作项目的数量。

```cpp
AEGP_GetNumJobItems(
 AEGP_RQM_SessionId sessid,
 A_long *num_jobitemsPL);
```

#### AEGP_GetJobItemID

获取具有指定索引的作业。

```cpp
AEGP_GetJobItemID(
 AEGP_RQM_SessionId sessid,
 A_long jobItemIndex,
 AEGP_RQM_ItemId *jobItemID);
```

#### AEGP_GetNumJobItemRenderSettings

获取具有指定索引的作业的渲染设置的数量。

```cpp
AEGP_GetNumJobItemRenderSettings(
 AEGP_RQM_SessionId sessid,
 AEGP_RQM_ItemId itemid,
 A_long *num_settingsPL);
```

#### AEGP_GetJobItemRenderSetting

获取特定作业的特定渲染设置。
设置名称和值是以 NULL 结尾的 A_UTF16Char 字符串的句柄，必须用`AEGP_FreeMemHandle`处理。

```cpp
AEGP_GetJobItemRenderSetting(
 AEGP_RQM_SessionId sessid,
 AEGP_RQM_ItemId itemid,
 A_long settingIndex,
 AEGP_MemHandle *utf_setting_namePH0,
 AEGP_MemHandle *utf_setting_valuePH0);
```

#### AEGP_GetNumJobItemOutputModules

获取具有指定索引的工作的输出模块的数量。

```cpp
AEGP_GetNumJobItemOutputModules(
 AEGP_RQM_SessionId sessid,
 AEGP_RQM_ItemId itemid,
 A_long *num_outputmodulesPL);
```

#### AEGP_GetNumJobItemOutputModuleSettings

获取具有指定索引的输出模块的设置数量。

```cpp
AEGP_GetNumJobItemOutputModuleSettings(
 AEGP_RQM_SessionId sessid,
 AEGP_RQM_ItemId itemid,
 A_long outputModuleIndex,
 A_long *num_settingsPL);
```

#### AEGP_GetJobItemOutputModuleSetting

获取一个工作项目输出模块的特定设置。
设置名称和值是以 NULL 结尾的 A_UTF16Char 字符串的句柄，必须用`AEGP_FreeMemHandle`处理。

```cpp
AEGP_GetJobItemOutputModuleSetting(
 AEGP_RQM_SessionId sessid,
 AEGP_RQM_ItemId itemid,
 A_long outputModuleIndex,
 A_long settingIndex,
 AEGP_MemHandle *utf_setting_namePH0,
 AEGP_MemHandle *utf_setting_valuePH0);
```

#### AEGP_GetNumJobItemOutputModuleWarnings

获取一个工作项目的输出模块警告的数量。

```cpp
AEGP_GetNumJobItemOutputModuleWarnings(
 AEGP_RQM_SessionId sessid,
 AEGP_RQM_ItemId itemid,
 A_long outputModuleIndex,
 A_long *num_warningsPL);
```

#### AEGP_GetJobItemOutputModuleWarning

获取一个特定工作项目的特定输出模块的警告。
警告值是一个以 NULL 结尾的 A_UTF16Char 字符串的句柄，必须用`AEGP_FreeMemHandle`处置。

```cpp
AEGP_GetJobItemOutputModuleWarning(
 AEGP_RQM_SessionId sessid,
 AEGP_RQM_ItemId itemid,
 A_long outputModuleIndex,
 A_long warningIndex,
 AEGP_MemHandle *utf_warning_valuePH0);
```

#### AEGP_GetNumJobItemFrameProperties

获取一个工作项目框架的属性数量。

```cpp
AEGP_GetNumJobItemFrameProperties(
 AEGP_RQM_SessionId sessid,
 AEGP_RQM_ItemId itemid,
 AEGP_RQM_FrameId frameid,
 A_long *num_propertiesPL);
```

#### AEGP_GetJobItemFrameProperty

获取一个工作项目框架的特定属性。
属性名称和值是以 NULL 结尾的 A_UTF16Char 字符串，必须用`AEGP_FreeMemHandle`处理。

```cpp
AEGP_GetJobItemFrameProperty(
 AEGP_RQM_SessionId sessid,
 AEGP_RQM_ItemId itemid,
 AEGP_RQM_FrameId frameid,
 A_long propertyIndex,
 AEGP_MemHandle *utf_property_namePH0,
 AEGP_MemHandle *utf_property_valuePH0);
```

#### AEGP_GetNumJobItemOutputModuleProperties

获取一个工作项目输出模块的属性数量。

```cpp
AEGP_GetNumJobItemOutputModuleProperties(
 AEGP_RQM_SessionId sessid,
 AEGP_RQM_ItemId itemid,
 A_long outputModuleIndex,
 A_long *num_propertiesPL);
```

#### AEGP_GetJobItemOutputModuleProperty

从一个工作项目输出模块上获取一个特定的属性。
属性名称和值是以 NULL 结尾的 A_UTF16Char 字符串，必须用`AEGP_FreeMemHandle`处理。

```cpp
AEGP_GetJobItemOutputModuleProperty(
 AEGP_RQM_SessionId sessid,
 AEGP_RQM_ItemId itemid,
 A_long outputModuleIndex,
 A_long propertyIndex,
 AEGP_MemHandle *utf_property_namePH0,
 AEGP_MemHandle *utf_property_valuePH0);
```

#### AEGP_GetJobItemFrameThumbnail

获得一个缓冲区，里面有工作项目框架的 JPEG 编码的缩略图。
输入最大宽度和高度，实际尺寸将被传回。

```cpp
AEGP_GetJobItemFrameThumbnail(
 AEGP_RQM_SessionId sessid,
 AEGP_RQM_ItemId itemid,
 AEGP_RQM_FrameId frameid,
 A_long *widthPL,
 A_long *heightPL,
 AEGP_MemHandle *thumbnailPH0);
```

## Output Module Suite

渲染队列中的每个项目都至少有一个指定的输出模块。

使用这个套件可以查询和控制附属于一个给定渲染项目的输出模块的所有方面。

你也可以添加和删除输出模块。

事实是这样的。对于一个给定的渲染项目的每一帧渲染，输出模块的列表都会被遍历。所以，对于第 0 帧，会调用输出模块 0，然后是 1，然后是 2。

### AEGP_OutputModuleSuite4

#### AEGP_GetOutputModuleByIndex

检索索引的输出模块。
注意：`AEGP_OutputModuleRefH`是一个不透明的数据类型，不能直接操作；你必须使用我们的访问函数来修改它。

```cpp
AEGP_GetOutputModuleByIndex(
 AEGP_RQItemRefH rq_itemH,
 A_long outmod_indexL,
 AEGP_OutputModuleRefH *outmodPH);
```

#### AEGP_GetEmbedOptions

检索为引用的`AEGP_OutputModuleRefH`指定的嵌入设置。

```cpp
AEGP_GetEmbedOptions(
 AEGP_RQItemRefH rq_itemH,
 AEGP_OutputModuleRefH outmodH,
 AEGP_EmbeddingType *embed_optionsP);
```

`AEGP_EmbeddingType`将是以下之一。

- `AEGP_Embedding_NOTHING`
- `AEGP_Embedding_LINK`
- `AEGP_Embedding_LINK_AND_COPY`

#### AEGP_SetEmbedOptions

为引用的`AEGP_OutputModuleRefH`指定嵌入设置。

```cpp
AEGP_SetEmbedOptions(
 AEGP_RQItemRefH rq_itemH,
 AEGP_OutputModuleRefH outmodH,
 AEGP_EmbeddingType embed_options);
```

#### AEGP_GetPostRenderAction

检索引用的`AEGP_OutputModuleRefH`的后期渲染动作设置。

```cpp
AEGP_GetPostRenderAction(
 AEGP_RQItemRefH rq_itemH,
 AEGP_OutputModuleRefH outmodH,
 AEGP_PostRenderAction *actionP);
```

`AEGP_PostRenderAction`将是以下之一。

- `AEGP_PostRenderOptions_IMPORT`
- `AEGP_PostRenderOptions_IMPORT_AND_REPLACE_USAGE`
- `AEGP_PostRenderOptions_SET_PROXY`

#### AEGP_SetPostRenderAction

指定引用的`AEGP_OutputModuleRefH`的后期渲染动作设置。

```cpp
AEGP_SetPostRenderAction(
 AEGP_RQItemRefH rq_itemH,
 AEGP_OutputModuleRefH outmodH,
 AEGP_PostRenderAction action);
```

#### AEGP_GetEnabledOutputs

检索引用的`AEGP_OutputModuleRefH`启用了哪些输出类型。

```cpp
AEGP_GetEnabledOutputs(
 AEGP_RQItemRefH rq_itemH,
 AEGP_OutputModuleRefH outmodH,
 AEGP_OutputTypes *typesP);
```

`AEGP_OutputTypes`将包含以下一个或两个值。

- `AEGP_OutputType_VIDEO`
- `AEGP_OutputType_AUDIO`

注意: 这些是标志，不是枚举。

#### AEGP_SetEnabledOutputs

指定引用的`AEGP_OutputModuleRefH`启用哪些输出类型。

```cpp
AEGP_SetEnabledOutputs(
 AEGP_RQItemRefH rq_itemH,
 AEGP_OutputModuleRefH outmodH,
 AEGP_OutputTypes enabled_types);
```

#### AEGP_GetOutputChannels

检索引用的 AEGP_OutputModuleRefH 中哪些视频通道被启用输出。

```cpp
AEGP_GetOutputChannels(
 AEGP_RQItemRefH rq_itemH,
 AEGP_OutputModuleRefH outmodH,
 AEGP_VideoChannels *outchannelsP);
```

`AEGP_VideoChannels`将是以下之一。

- `AEGP_VideoChannels_RGB`
- `AEGP_VideoChannels_RGBA`
- `AEGP_VideoChannels_ALPHA`

#### AEGP_SetOutputChannels

指定在引用的`AEGP_OutputModuleRefH`中哪些视频通道被启用进行输出。

```cpp
AEGP_SetOutputChannels(
 AEGP_RQItemRefH rq_itemH,
 AEGP_OutputModuleRefH outmodH,
 AEGP_VideoChannels outchannels);
```

#### AEGP_GetStretchInfo

检索引用的`AEGP_OutputModuleRefH`启用的拉伸信息；是否启用了拉伸功能，是否将帧的宽高比锁定在合成上，以及指定了什么质量设置。

```cpp
AEGP_GetStretchInfo(
 AEGP_RQItemRefH rq_itemH,
 AEGP_OutputModuleRefH outmodH,
 A_Boolean *enabledPB,
 AEGP_StretchQuality *qualP,
 A_Boolean *lockedPB);
```

`AEGP_StretchQuality`将是以下之一。

- `AEGP_StretchQual_LOW`
- `AEGP_StretchQual_HIGH`

#### AEGP_SetStretchInfo

检索引用的`AEGP_OutputModuleRefH`所启用的拉伸信息。

```cpp
AEGP_SetStretchInfo(
 AEGP_RQItemRefH rq_itemH,
 AEGP_OutputModuleRefH outmodH,
 A_Boolean is_enabledB,
 AEGP_StretchQuality quality);
```

#### AEGP_GetCropInfo

检索引用的`AEGP_OutputModuleRefH`是否启用裁剪，以及要使用的矩形。

```cpp
AEGP_GetCropInfo(
 AEGP_RQItemRefH rq_itemH,
 AEGP_OutputModuleRefH outmodH,
 A_Boolean *is_enabledBP,
 A_Rect *crop_rectP);
```

#### AEGP_SetCropInfo

指定参考的`AEGP_OutputModuleRefH`是否启用裁剪，以及要使用的矩形。

```cpp
AEGP_SetCropInfo(
 AEGP_RQItemRefH rq_itemH,
 AEGP_OutputModuleRefH outmodH,
 A_Boolean enableB,
 A_Rect crop_rect);
```

#### AEGP_GetSoundFormatInfo

读取引用的`AEGP_OutputModuleRef`是否启用音频输出，以及要使用的设置。

```cpp
AEGP_GetSoundFormatInfo(
 AEGP_RQItemRefH rq_itemH,
 AEGP_OutputModuleRefH outmodH,
 AEGP_SoundDataFormat *formatP,
 A_Boolean *enabledPB);
```

#### AEGP_SetSoundFormatInfo

指定引用的`AEGP_OutputModuleRefH`是否启用音频输出，以及要使用的设置。

```cpp
AEGP_SetSoundFormatInfo(
 AEGP_RQItemRefH rq_itemH,
 AEGP_OutputModuleRefH outmodH,
 AEGP_SoundDataFormat format_info,
 A_Boolean enabledB);
```

#### AEGP_GetOutputFilePath

检索 "AEGP_OutputModuleRefH "的输出文件将被写入的路径。
该路径是一个以 NULL 结尾的 A_UTF16Char 字符串的句柄，必须用`AEGP_FreeMemHandle`来处理。

```cpp
AEGP_GetOutputFilePath(
 AEGP_RQItemRefH rq_itemH,
 AEGP_OutputModuleRefH outmodH,
 AEGP_MemHandle *unicode_pathPH);
```

#### AEGP_SetOutputFilePath

指定`AEGP_OutputModuleRefH`的输出文件将被写入的路径。
文件路径是一个以 NULL 结尾的 UTF-16 字符串，带有平台分隔符。

```cpp
AEGP_SetOutputFilePath(
 AEGP_RQItemRefH rq_itemH,
 AEGP_OutputModuleRefH outmodH,
 const A_UTF16Char *pathZ);
```

#### AEGP_AddDefaultOutputModule

为指定的`AEGP_RQItemRefH`添加默认的输出模块，并返回添加的输出模块的`AEGP_OutputModuleRefH`(如果你不打算乱用它，就不会添加它，对吗？)

```cpp
AEGP_AddDefaultOutputModule(
 AEGP_RQItemRefH rq_itemH,
 AEGP_OutputModuleRefH *outmodPH);
```

#### AEGP_GetExtraOutputModuleInfo

检索关于输出模块的信息。
`format_uniPH`和`info_uniPH`提供输出模块的文字描述和信息，格式为用户看到的样子。
`format_uniPH`和`info_uniPH`将包含以 NULL 结尾的 UTF16 字符串，调用者必须处理掉。

```cpp
AEGP_GetExtraOutputModuleInfo(
 AEGP_RQItemRefH rq_itemH,
 AEGP_OutputModuleRefH outmodH,
 AEGP_MemHandle *format_uniPH,
 AEGP_MemHandle *info_uniPH,
 A_Boolean *is_sequenceBP,
 A_Boolean *multi_frameBP);
```

## Working With Effects

这些函数为效果(和 AEGPs)提供了一种方法，以获得关于应用效果的上下文的信息。

::: tip

任何时候，当你修改或依赖来自正常渲染管道之外的数据时，你都有可能出现依赖性问题。
:::
After Effects 没有办法知道你依赖这些外部信息；因此，如果这些信息从你下面改变了，你将不会得到通知。

### AEGP_PFInterfaceSuite1

#### AEGP_GetEffectLayer

获取应用该效果的图层的句柄。

```cpp
AEGP_GetEffectLayer(
 PF_ProgPtr effect_ref,
 AEGP_LayerH *layerPH);
```

#### AEGP_GetNewEffectForEffect

获得与效果对应的`AEGP_EffectRefH`。

```cpp
AEGP_GetNewEffectForEffect(
 AEGP_PluginID aegp_plugin_id,
 PF_ProgPtr effect_ref,
 AEGP_EffectRefH *effectPH);
```

#### AEGP_ConvertEffectToCompTime

获取与效果图层时间相对应的合成时间。

```cpp
AEGP_ConvertEffectToCompTime(
 PF_ProgPtr effect_ref,
 long what_timeL,
 unsigned long time_scaleLu,
 A_Time *comp_timePT);
```

#### AEGP_GetEffectCamera

获取 After Effects 用于查看效果层的摄像机(如果有的话)。

```cpp
AEGP_GetEffectCamera(
 PF_ProgPtr effect_ref,
 const A_Time *comp_timePT,
 AEGP_LayerH camera_layerPH);
```

#### AEGP_GetEffectCameraMatrix

获取用于在该层的坐标空间和包含的合成空间之间移动的变换。

```cpp
AEGP_GetEffectCameraMatrix(
 PF_ProgPtr effect_ref,
 const A_Time *comp_timePT,
 A_Matrix4 *camera_matrixP,
 A_FpLong *dst_to_planePF,
 A_short *plane_widthPL,
 A_short *plane_heightPL);
```

注意：如果效果的输入层是方形像素，但在一个非方形像素的合成中。
你必须通过将矩阵预先乘以`(1/parF, 1, 1)`来校正像素长宽比。

### AEGP_GetEffectCameraMatrix Notes

摄像机矩阵的模型视图是由`AEGP_GetEffectCameraMatrix()`得到的矩阵的逆向。

还要注意，我们的矩阵是基于行的；OpenGL 的是基于列的。

## Do This Many Times

利用多个处理器(如果有的话)进行计算。

### AEGP_IterateSuite1

#### AEGP_GetNumThreads

询问 After Effects 目前有多少个线程可用。

```cpp
AEGP_GetNumThreads(
 A_long *num_threadsPL);
```

#### AEGP_IterateGeneric

指定一个函数，让 After Effects 在多个处理器上管理。
可以是由`fn_func`指定的任何函数指针，采取下面列出的参数。
参见[Private Data](implementation.html#aegps-implementation-private-data)以了解 refconPV 的使用方法。

```cpp
AEGP_IterateGeneric(
 A_long iterationsL,
 void *refconPV,
 A_Err (*fn_func)
 (void *refconPV,
 A_long thread_indexL,
 A_long i,
 A_long iterationsL));
```

## File Import Manager Suite

FIMSuite 允许由 AEGPs 处理的文件类型作为 After Effects 导入对话框的一部分出现，以及拖放信息。

这些不是给 AEIO 们使用的! 相反，它们是用来导入那些最适合作为 After Effects 合成的项目的。
这些不是给 AEIO 使用的!

### AEGP_FIMSuite3

#### AEGP_RegisterImportFlavor

注册插件所支持的文件类型的名称。
返回时，`imp_refP`将是一个有效的不透明引用，或`AE_FIM_ImportFlavorRef_NONE`。

```cpp
AEGP_RegisterImportFlavor(
 const char *nameZ,
 AE_FIM_ImportFlavorRef *imp_refP);
```

#### AEGP_RegisterImportFlavorFileTypes

注册 AEGP 支持的文件类型和文件扩展名的数组(两个数组的长度不需要相等)。

```cpp
AEGP_RegisterImportFlavorFileTypes(
 AE_FIM_ImportFlavorRef imp_ref,
 long num_filekindsL,
 const AEIO_FileKind *kindsAP,
 long num_fileextsL,
 const AEIO_FileKind *extsAP);
```

#### AEGP_RegisterImportFlavorImportCallbacks

注册 AEGP 函数，这些函数将响应不同文件类型的导入。

```cpp
AEGP_RegisterImportFlavorImportCallbacks(
 AE_FIM_ImportFlavorRef ref,
 AE_FIM_ImportFlags single_flag,
 const AE_FIM_ImportCallbacks *imp_cbsP);
```

#### AEGP_SetImportedItem

指定一个项目被导入(可能是替换现有项目)，并设置相关的导入选项。

```cpp
AEGP_SetImportedItem(
 AE_FIM_ImportOptions imp_options,
 AEGP_ItemH imported_itemH);
```
