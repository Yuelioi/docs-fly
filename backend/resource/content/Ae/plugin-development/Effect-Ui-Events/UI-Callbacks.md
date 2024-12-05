---
title: UI-Callbacks
order: 7
---

# UI 回调

After Effects 提供了回调功能，用于在坐标系之间进行转换，并获得操作系统特定的绘图上下文信息，无需猜测或询问操作系统。请使用这些回调! 这些回调的指针在 `PF_EventCallbacks` 中提供。使用 `AE_EffectUI.h` 和 `AE_EffectCB.h` 中的宏来访问这些程序。

有可能在不实现坐标系换位回调的情况下建立一个利用自定义 UI 的有效插件。然而当用户放大图层面板或旋转图层时，插件将表现得很糟糕。我们添加了这些宏和回调，以便自定义用户界面可以很容易地集成到 After Effects 的用户界面中，而不会给开发者带来用户界面的开销。再次强调，请使用它们!

为了简单起见，这些宏默认了 refcon 和上下文句柄。refcon 假定你有一个名为 "extra "的局部变量。默认的上下文是当前上下文。这些默认参数在 `PF_EventCallbacks` 结构中定义(在 `AE_EffectUI.h` 中)。可以通过 `PF_EventExtra` 结构访问这些回调来覆盖默认参数。我们不建议(或支持)修改头文件中的宏。请不要这样做!

### layer_to_comp

将图层坐标转换为合成面板坐标。

```cpp
PF_Err layer_to_comp (
  void           *refcon,
  PF_ContextH    context,
  A_long         curr_time,
  A_long         time_scale,
  PF_FixedPoint  *pt);
```

### comp_to_layer

将合成面板坐标转换为图层面板坐标。

```cpp
PF_Err comp_to_layer (
  void           *refcon,
  PF_ContextH    context,
  A_long         curr_time,
  A_long         time_scale,
  PF_FixedPoint  *pt);
```

### get_comp2layer_xform

返回用于从合成面板转换到图层面板的矩阵。如果`*exists`返回`FALSE`，则无法计算矩阵，因为图层缩放为零。

```cpp
PF_Err get_comp2layer_xform (
  void            *refcon,
  PF_ContextH     context,
  A_long          curr_time,
  long            time_scale,
  long            *exists,
  PF_FloatMatrix  *comp2layer);
```

### get_layer2comp_xform

返回用于从图层面板转换到合成面板的转换矩阵。这始终存在。

```cpp
PF_Err get_layer2comp_xform (
  void            *refcon,
  PF_ContextH     context,
  A_long          curr_time,
  A_long          time_scale,
  PF_FloatMatrix  *layer2comp);
```

### source_to_frame

将当前上下文中的源坐标转换为屏幕坐标。屏幕（帧）坐标受当前缩放级别的影响。

```cpp
PF_Err source_to_frame(
  void           *refcon,
  PF_ContextH    context,
  PF_FixedPoint  *pt);
```

### frame_to_source

将`*pt`标识的屏幕坐标转换为当前上下文的源坐标。

```cpp
PF_Err frame_to_source(
  void           *refcon,
  PF_ContextH    context,
  PF_FixedPoint  *pt);
```

### PF_GET_PLATFORM_DATA

检索特定于平台的数据。对于加载了本地化资源文件的插件，`PF_PlatData_RES_FILE_PATH` 将指向外部文件，而不是插件文件。如果需要插件的路径，请使用`PF_PlatData_EXE_FILE_PATH`

从CS6开始，使用 `PF_PlatData_EXE_FILE_PATH_W` and `PF_PlatData_RES_FILE_PATH_W` 代替旧的非宽调用。

```cpp
PF_Err PF_GET_PLATFORM_DATA (
  PF_PlatDataID  which,
  void           *ppData);
```

PF_PlatDataID 可以具有以下值：

> - `PF_PlatData_MAIN_WND`
> - `PF_PlatData_EXE_FILE_PATH_DEPRECATED`
> - `PF_PlatData_RES_FILE_PATH_DEPRECATED`
> - `PF_PlatData_RES_REFNUM` // macOS
> - `PF_PlatData_RES_DLLINSTANCE` // Win
> - `PF_PlatData_BUNDLE_REF`
> - `PF_PlatData_EXE_FILE_PATH_W` // new CS6
> - `PF_PlatData_RES_FILE_PATH_W` // new CS6
