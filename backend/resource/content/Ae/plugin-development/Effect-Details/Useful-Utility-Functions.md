---
title: Useful-Utility-Functions
order: 20
---

# 实用函数

## PF_EffectUISuite

虽然不是严格意义上的参数，这个套件可以改变选项按钮的名称。

### PF_SetOptionsButtonName

改变效果控制面板中选项按钮上的文字。

注意：必须在[PF_Cmd_PARAM_SETUP](../effect-basics/command-selectors.html)中调用。

```cpp
PF_SetOptionsButtonName(
PF_ProgPtreffect_ref,
constA_char*nameZ);
```

`nameZ`的长度可以达到`A_char[31]`。

## PF_AppSuite

大约437年前，当我们发布AE 5.0时，我们在PF_AppSuite发布了一些有用的实用程序回调。它们今天和那时一样有用。AE具有用户可控制的UI亮度。

除了[PF_EffectCustomUIOverlayThemeSuite](../effect-ui-events/custom-ui-and-drawbot.html) 用于插件中的自定义 UI，使用这些调用可以无缝集成到 After Effects 的 UI 中。

还有什么比把别人的个人信息放到水印里，更能增加购买欲呢，嗯？或者将光标设置为添加遮罩顶点，只是为了迷惑人们？嘿嘿嘿。但那可不好。

### PF_AppGetBgColor

检索当前的背景颜色。

```cpp
PF_AppGetBgColor(
PF_App_Colorbg_colorP);
```

### PF_AppGetColor

检索指定的 UI 元素的颜色。参见 AE_EffectSuites.h 中的 `PF_App_Color`值的完整面板枚举；基本上可以检索AE UI中的任何颜色。

CC 为新元素增加了几个新的 `PF_App_ColorType`枚举值，可以查询。

请注意，在 CS6 中，颜色定义`从FILL_LIGHT`向下。

请使用以下仅适用于 CS6 的 psuedocode。

```cpp
GetColor(enume)
{
ifhost_is_CS6ande>=FILL_LIGHT
e+=3
callrealGetColor
}

PF_AppGetColor(
PF_App_ColorTypecolor_type,
PF_App_Color*app_colorP);
```

### PF_AppGetLanguage

CC 中的新功能。检索 AE 用户界面的活动显示语言，以便插件可以匹配。以下是 CC 中可能的语言代码。

- 中文 - `zh_CN`。
- 英语 - `en_US`
- 法语 - `fr_FR`
- 德文 - `de_DE`
- 意大利语 - `it_IT`
- 日语 - `ja_JP`
- 韩国语 - `ko_KR`
- 西班牙语 - `es_ES`

```cpp
PF_AppGetLanguage(
A_charlang_tagZ);
```

### PF_GetPersonalInfo

检索用户的注册信息。

```cpp
PF_GetPersonalInfo(
PF_AppPersonalTextInfo*ptiP);

typedefstructPF_AppPersonalTextInfo{
A_charname[PF_APP_MAX_PERS_LEN+1];
A_charorg[PF_APP_MAX_PERS_LEN+1];
A_charserial_str[PF_APP_MAX_PERS_LEN+1];
}PF_AppPersonalTextInfo;
```

### PF_GetFontStyleSheet

检索 After Effects 用户界面中使用的字体样式表信息。

琐事：从 15.0 开始，After Effects 的用户界面中使用的字体是 Adobe Clean。在此之前，Windows 上是 Tahoma，macOS X 上是 Lucida Grande。

```cpp
PF_GetFontStyleSheet(
PF_FontStyleSheetsheet,
PF_FontName*font_nameP0,
A_short*font_numPS0,
A_short*sizePS0,
A_short*stylePS0);
```

### PF_SetCursor

将光标设置为 After Effects 中的任何光标。详见见 AE_EffectUI.h 的枚举值。

设置为:

- `PF_Cursor_NONE`，允许 After Effects 设置光标。
- `PF_Cursor_CUSTOM`，如果你使用操作系统特定的调用来改变光标 (After Effects 将尊重你的改变)。

```cpp
PF_SetCursor(
PF_CursorTypecursor);
```

### PF_IsRenderEngine

如果 After Effects 在 "监视文件夹" 模式下运行，或者是仅使用渲染引擎模式下，返回 TRUE。

```cpp
PF_IsRenderEngine(
PF_Boolean*render_enginePB);
```

从 AE6.5 开始，如果安装的是渲染引擎，或者 After Effects 在运行时没有用户界面，或者 After Effects 处于 "监视文件夹" 模式，这个函数返回 TRUE`。

### PF_AppColorPickerDialog

显示 After Effects 颜色选择器对话框(可能是系统颜色选择器，取决于用户的偏好)。

如果用户取消对话框，将返回 `PF_Interrupt_CANCEL`。返回的颜色是在项目的工作色彩空间。

```cpp
PF_AppColorPickerDialog(
constA_char*dialog_titleZ0,
constPF_PixelFloat*sample_colorP,
PF_PixelFloat*result_colorP);
```

### PF_GetMouse

返回鼠标在自定义 UI 坐标空间中的位置。

```cpp
PF_GetMouse(
PF_Point*pointP);
```

### PF_InvalidateRect

排队等待自定义用户界面的特定区域的[重绘](../effect-ui-events/custom-ui-and-drawbot.html)，以获得效果。

只在处理效果中的非绘制事件时有效。

将 `rectP0`指定为`NULL`可以使整个窗口失效。重绘将在事件返回后的下一个可用的空闲时间发生。

设置 `PF_EO_UPDATE_NOW`事件输出标志，在事件返回后立即更新窗口。

```cpp
PF_InvalidateRect(
constPF_ContextHcontextH,
constPF_Rect*rectP0);
```

### PF_ConvertLocalToGlobal

从自定义 UI 坐标系统转换为全局屏幕坐标。仅在自定义 UI 事件处理中使用。

```cpp
PF_ConvertLocalToGlobal(
constPF_Point*localP,
PF_Point*globalP);
```

## 高级应用套件：你可以做到这一点

`PF_AdvAppSuite`最初是为了一些非常邪恶的目的而设计的；一个外部应用程序假装是一个AE插件，并且需要一些方法来通知AE它对项目所做的更改。我们的API杂质是你的收获。

## PF_AdvAppSuite2

### PF_SetProjectDirty

告诉 After Effects 自从上次保存后，项目已经被改变。

```cpp
PF_SetProjectDirty(void);
```

### PF_SaveProject

将项目保存到当前路径。要把项目保存在其他地方，请使用[AEGP_SaveProjectToPath()](../aegps/aegp-suites.html)。

```cpp
PF_SaveProject(void);
```

### PF_SaveBackgroundState

存储背景状态(AE在打开的应用程序和窗口堆叠顺序中的位置)。

```cpp
PF_SaveBackgroundState(void);
```

### PF_ForceForeground

将 After Effects 置顶。

```cpp
PF_ForceForeground(void);
```

### PF_RestoreBackgroundState

将 After Effects 放回原来的堆叠位置，相对于其他的应用程序和窗口而言。

```cpp
PF_RestoreBackgroundState(void);
```

### PF_RefreshAllWindows

强制更新所有 After Effects 窗口。注意，虽然合成面板会被刷新，但这并不保证会有新的帧被发送到外部监视器预览插件。

```cpp
PF_RefreshAllWindows(void);
```

### PF_InfoDrawText

将文本写入 After Effects 信息面板。

```cpp
PF_InfoDrawText(
constA_char*line1Z0,
constA_char*line2Z0);
```

### PF_InfoDrawColor

在 After Effects 信息面板中绘制指定的颜色(Alpha 被忽略)。

```cpp
PF_InfoDrawColor(
PF_Pixelcolor);
```

### PF_InfoDrawText3

在 After Effects 信息面板中写入三行文字。

```cpp
PF_InfoDrawText3(
constA_char*line1Z0,
constA_char*line2Z0,
constA_char*line3Z0);
```

### PF_InfoDrawText3Plus

在 After Effects 信息面板上写三行文字，第二行和第三行的部分内容是左右对齐的。

```cpp
PF_InfoDrawText3Plus(
constA_char*line1Z0,
constA_char*line2_jrZ0,
constA_char*line2_jlZ0,
constA_char*line3_jrZ0,
constA_char*line3_jlZ0);
```

### PF_AppendInfoText

在当前显示的信息文本中添加字符。

```cpp
PF_AppendInfoText(
constA_char*appendZ0);
```

## Formatting Time

`PF_AdvTimeSuite`提供了几个函数来配合 After Effects 显示时间的方式。事实上，这些都是我们内部使用的函数。

### PF_AdvTimeSuite4

#### PF_FormatTimeActiveItem

给定一个时间值和刻度，返回一个代表该时间的格式化字符串。如果 `durationB` 为 `TRUE`，将附加适当的单位。

```cpp
PF_FormatTimeActiveItem(
A_longtime_valueUL,
A_u_longtime_scaleL,
PF_BooleandurationB,
A_char*time_buf);
```

#### PF_FormatTime

为给定的 PF_InData 和 PF_EffectWorld(即图层时间)提供格式化的时间字符串的上下文。

```cpp
PF_FormatTime(
PF_InData*in_data,
PF_EffectWorld*world,
A_longtime_valueUL,
A_u_longtime_scaleL,
PF_BooleandurationB,
A_char*time_buf);
```

#### PF_FormatTimePlus

允许你选择合成或图层时间。

```cpp
PF_FormatTimePlus(
PF_InData*in_data,
PF_EffectWorld*world,
A_longtime_valueUL,
A_u_longtime_scaleL,
PF_Booleancomp_timeB,
PF_BooleandurationB,
A_char*time_buf);
```

#### PF_GetTimeDisplayPref

返回起始帧数(由用户在合成设置中指定)，以及合成的时间显示偏好。在 14.2 中更新，以支持更高的帧率。

```cpp
PF_GetTimeDisplayPref(
PF_TimeDisplayPref2*tdp,
A_long*starting_num);
typedefstruct{
A_chardisplay_mode;
A_longframemax;
A_longframes_per_foot;
A_charframes_start;
A_Booleannondrop30B;
A_Booleanhonor_source_timecodeB;
A_Booleanuse_feet_framesB;
}PF_TimeDisplayPrefVersion3;
```

#### PF_TimeCountFrames

15.0 版中的新内容。返回当前合成帧的索引。

```cpp
PF_TimeCountFrames(
constA_Time*start_timeTP,
constA_Time*time_stepTP,
A_Booleaninclude_partial_frameB,
A_long*frame_countL);
```

## 影响时间轴

很久以前，我们帮助开发人员将他们的独立跟踪器与AE集成，方法是公开一组功能，为他们提供某种方式来通知我们和告知它们通知时间轴的更改。。

由于有大量可用的AEGP API调用，这些调用并没有被大量使用，但它们仍然可用。

不要把这个套件和[AEGP_ItemSuite](../aegps/aegp-suites.html)混淆。

### PF_AdvItemSuite1

#### PF_MoveTimeStep

将当前时间 num_stepsL 向指定方向移动。

```cpp
PF_MoveTimeStep(
PF_InData*in_data,
PF_EffectWorld*world,
PF_Steptime_dir,
A_longnum_stepsL);
```

#### PF_MoveTimeStepActiveItem

在指定的方向上移动活动项目的 num_stepsL。

```cpp
PF_MoveTimeStepActiveItem(
PF_Steptime_dir,
A_longnum_stepsL);
```

#### PF_TouchActiveItem

告诉 After Effects 必须更新活动项目。

```cpp
PF_TouchActiveItem(void);
```

#### PF_ForceRerender

强制 After Effects 重新渲染当前帧。

```cpp
PF_ForceRerender(
PF_InData*in_data,
PF_EffectWorld*world);
```

#### PF_EffectIsActiveOrEnabled

返回拥有`PF_ContextH`的效果当前是否激活或启用（如果不是，AE将不会监听来自它的函数调用）。

```cpp
PF_EffectIsActiveOrEnabled(
PF_ContextHcontextH,
PF_Boolean*enabledPB);
```

## 访问辅助通道数据

有些文件类型不仅仅包含像素数据；使用`PF_ChannelSuite`来确定是否存在这些信息，并使用 `AE_ChannelSuites.h` 中的宏来检索你需要的格式。

### PF_ChannelSuite1

#### PF_GetLayerChannelCount

检索与索引层相关的辅助通道的数量。

```cpp
PF_GetLayerChannelCount(
PF_ProgPtreffect_ref,
PF_ParamIndexparam_index,
A_long*num_channelsPL);
```

#### PF_GetLayerChannelIndexedRefAndDesc

检索(通过索引)对指定通道的引用和描述。

```cpp
PF_GetLayerChannelIndexedRefAndDesc(
PF_ProgPtreffect_ref,
PF_ParamIndexparam_index,
PF_ChannelIndexchannel_index,
PF_Boolean*foundPB,
PF_ChannelRef*channel_refP,
PF_ChannelDesc*channel_descP);
```

#### PF_GetLayerChannelTypedRefAndDesc

按类型检索一个辅助通道。只有当 `foundPB`返回TRUE`时，返回的信息才有效。

```cpp
PF_GetLayerChannelTypedRefAndDesc(
PF_ProgPtreffect_ref,
PF_ParamIndexparam_index,
PF_ChannelTypechannel_type,
PF_Boolean*foundPB,
PF_ChannelRef*channel_refP,
PF_ChannelDesc*channel_descP);
```

`PF_DataType`, 以下之一。

> - `PF_DataType_FLOAT` - 34 字节
> - `PF_DataType_DOUBLE` - 38 字节
> - `PF_DataType_LONG` - 34 字节
> - `PF_DataType_SHORT` - 32 字节
> - `PF_DataType_FIXED_16_16` - 34 字节
> - `PF_DataType_CHAR` - 31 字节
> - `PF_DataType_U_BYTE` - 31 字节
> - `PF_DataType_U_SHORT` - 32 字节
> - `PF_DataType_U_FIXED_16_16` - 34 字节
> - `PF_DataType_RGB` - 3 字节

PF_ChannelType 将是以下之一。

> - `PF_ChannelType_DEPTH`
> - `PF_ChannelType_NORMALS`
> - `PF_ChannelType_OBJECTID`
> - `PF_ChannelType_MOTIONVECTOR`
> - `PF_ChannelType_NORMALS`
> - `PF_ChannelType_BK_COLOR`
> - `PF_ChannelType_TEXTURE`
> - `PF_ChannelType_COVERAGE`
> - `PF_ChannelType_NODE`
> - `PF_ChannelType_MATERIAL`
> - `PF_ChannelType_UNCLAMPED`
> - `PF_ChannelType_UNKNOWN`

#### PF_CheckoutLayerChannel

检索包含与给定 `PF_ChannelRefPtr`相关的数据的`PF_ChannelChunk`。

```cpp
PF_CheckoutLayerChannel(
PF_ProgPtreffect_ref,
PF_ChannelRefPtrchannel_refP,
longwhat_time,
longduration,
unsignedlongtime_scale,
PF_DataTypedata_type,
PF_ChannelChunk*channel_chunkP);
```

#### PF_CheckinLayerChannel

在 `PF_ChannelChunk`中进行检查。总总总会检查数据回来。

```cpp
PF_CheckinLayerChannel(
PF_ProgPtreffect_ref,
PF_ChannelRefPtrchannel_refP,
PF_ChannelChunk*channel_chunkP);
```
