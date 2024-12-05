---
order: 9
title: Viewer
---

# Viewer 查看器

## Viewer 查看器对象

app.activeViewer

描述：查看器对象表示“合成”，“层”或“素材”面板。

示例：最大化活动的查看器面板，并在包含合成的情况下显示其类型。

```javascript
var activeViewer = app.activeViewer;
activeViewer.maximized = true;
if (activeViewer.type === ViewerType.VIEWER_COMPOSITION) {
  alert("Composition panel is active.");
}
```

## 属性

### active

viewer.active

描述：如果为 true，则指示查看器面板是否聚焦，从而位于最前面。

类型：布尔值 只读。

### fastPreview

viewer.fastPreview

描述：快速预览菜单的状态这是使用枚举值的读/写属性：

:::danger
如果尝试在“图层”或“素材”面板中获取或设置属性的值，则会收到错误消息。
:::

:::tip
草稿预览模式仅在光线跟踪的 3D 合成中可用。如果尝试在 Classic 3D 合成中使用它，则会出现错误：“无法在 Classic
3D 合成中设置草稿快速预览模式。”
:::

类型：FastPreviewType，枚举值; 读/写：

- FastPreviewType.FP_OFF：关(最终质量)
- FastPreviewType.FP_ADAPTIVE_RESOLUTION：自适应分辨率
- FastPreviewType.FP_DRAFT： 草稿
- FastPreviewType.FP_FAST_DRAFT：快速草稿
- FastPreviewType.FP_WIREFRAME：线框

示例

```javascript
app.activeViewer.views[0].options.fastPreview === FastPreviewType.FP_ADAPTIVE_RESOLUTION;
app.activeViewer.views[0].options.fastPreview === FastPreviewType.FP_DRAFT;
app.activeViewer.views[0].options.fastPreview === FastPreviewType.FP_FAST_DRAFT;
app.activeViewer.views[0].options.fastPreview === FastPreviewType.FP_OFF;
app.activeViewer.views[0].options.fastPreview === FastPreviewType.FP_WIREFRAME;
```

### guidesLocked

viewer.guidesLocked

::: info 提示
此功能已添加到 After Effects 16.1(CC 2019)中
:::

描述：如果为 true，则表示参考线已锁定在查看器中。

类型：布尔值 读/写。

示例

app.activeViewer.views[0].options.guidesLocked;

### guidesSnap

viewer.guidesSnap

::: info 提示
此功能已添加到 After Effects 16.1(CC 2019)中
:::

描述：如果为 true，则表示在查看器中拖动时图层会紧贴参考线。

类型：布尔值 读/写。

示例

app.activeViewer.views[0].options.guidesSnap;

### guidesVisibility

viewer.guidesVisibility

::: info 提示
此功能已添加到 After Effects 16.1(CC 2019)中
:::

描述：如果为 true，则表示参考线在查看器中可见。

类型：布尔值 读/写。

示例

app.activeViewer.views[0].options.guidesVisibility;

### maximized

viewer.maximized

描述：如果为 true，则指示查看器面板是否处于其最大尺寸。

类型：布尔值 读/写。

### rulers

::: info 提示
此功能已添加到 After Effects 16.1(CC 2019)中
:::

描述：如果为 true，则表示在查看器中显示了标尺。 类型：布尔值 读/写。

示例 app.activeViewer.views[0].options.rulers;

描述：如果为 true，则表示在查看器中显示了标尺。

类型布尔值 读/写。

例子

app.activeViewer.views[0].options.rulers;

### type

viewer.type

描述：查看器面板中的内容。

类型：ViewerType，枚举值; 只读：

- ViewerType.VIEWER_COMPOSITION
- ViewerType.VIEWER_LAYER
- ViewerType.VIEWER_FOOTAGE

## 方法

### setActive()

viewer.setActive()

描述：将查看器面板移到前面并对其进行聚焦，使其处于活动状态。调用此方法会将查看者的 active 属性设置为 true。

参数：无。

返回：布尔值，指示查看器面板是否处于活动状态。
