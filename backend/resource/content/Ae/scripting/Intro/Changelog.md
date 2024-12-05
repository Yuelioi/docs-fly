---
title: Changelog
order: 2
---

# 更新日志

## 脚本更新了啥?

### After Effects 18.0(2021 年 3 月)

- 支持媒体替换的脚本编制方法和属性

  - 新增：[AVItem.isMediaReplacementCompatible](https://ae-scripting.docsforadobe.dev/items/avitem/#avitem-ismediareplacementcompatible)
  - 新增：[AVLayer.addToMotionGraphicsTemplate()](https://ae-scripting.docsforadobe.dev/layers/avlayer/#avlayer-addtomotiongraphicstemplate)
  - 新增：[AVLayer.addToMotionGraphicsTemplateAs()](https://ae-scripting.docsforadobe.dev/layers/avlayer/#avlayer-addtomotiongraphicstemplateas)
  - 新增：[AVLayer.canAddToMotionGraphicsTemplate()](https://ae-scripting.docsforadobe.dev/layers/avlayer/#avlayer-canaddtomotiongraphicstemplate)
  - 新增：[Property.alternateSource](https://ae-scripting.docsforadobe.dev/properties/property/#property-alternatesource)
  - 新增：[Property.canSetAlternateSource](https://ae-scripting.docsforadobe.dev/properties/property/#property-cansetalternatesource)
  - 添加：[Property.setAlternateSource()](https://ae-scripting.docsforadobe.dev/properties/property/#property-setalternatesource)
  - 添加了相关的[匹配名称](https://ae-scripting.docsforadobe.dev/matchnames/layer/avlayer/#matchnames-layer-avlayer)

- [为基本属性](https://ae-scripting.docsforadobe.dev/matchnames/layer/avlayer/#matchnames-layer-avlayer)属性组添加了[匹配名称](https://ae-scripting.docsforadobe.dev/matchnames/layer/avlayer/#matchnames-layer-avlayer)。

---

### After Effects 17.1.1(2020 年 5 月)

- 脚本访问形状层笔划锥度，笔划波形，偏移路径副本，偏移路径副本偏移

  - 添加了相关的[匹配名称](https://ae-scripting.docsforadobe.dev/matchnames/layer/shapelayer/#matchnames-layer-shapelayer)

- 修复了允许[CompItem.displayStartTime](https://ae-scripting.docsforadobe.dev/items/compitem/#compitem-displaystarttime)为[负值的问题](https://ae-scripting.docsforadobe.dev/items/compitem/#compitem-displaystarttime)：

  - 添加了[CompItem.displayStartFrame](https://ae-scripting.docsforadobe.dev/items/compitem/#compitem-displaystartframe)
  - 现在，与在“合成设置”对话框中设置“开始时间码”时所允许的有效范围相匹配(-3：00：00：00 到 23：59：00：00)。

---

### After Effects 17.0.1(2019 年 11 月)

- 脚本创建和修改下拉菜单控件项：

  - 新增：[Property.isDropdownEffect](https://ae-scripting.docsforadobe.dev/properties/property/#property-isdropdowneffect)
  - 添加：[Property.setPropertyParameters()](https://ae-scripting.docsforadobe.dev/properties/property/#property-setpropertyparameters)

---

### After Effects 16.1

- 脚本访问查看器指南和标尺布尔值：

  - 新增：[Viewer.guidesLocked](https://ae-scripting.docsforadobe.dev/other/viewer/#viewer-guideslocked)
  - 新增：[Viewer.guidesSnap](https://ae-scripting.docsforadobe.dev/other/viewer/#viewer-guidessnap)
  - 新增：[Viewer.guidesVisibility](https://ae-scripting.docsforadobe.dev/other/viewer/#viewer-guidesvisibility)
  - 新增：[Viewer.rulers](https://ae-scripting.docsforadobe.dev/other/viewer/#viewer-rulers)

- 脚本访问以添加，删除和设置现有指南：

  - 新增：[Item.addGuide()](https://ae-scripting.docsforadobe.dev/items/item/#item-addguide)
  - 新增：[Item.removeGuide()](https://ae-scripting.docsforadobe.dev/items/item/#item-removeguide)
  - 新增：[Item.setGuide()](https://ae-scripting.docsforadobe.dev/items/item/#item-setguide)

- 脚本访问其他 EGP 属性属性：

  - 新增：[CompItem.motionGraphicsTemplateControllerCount](https://ae-scripting.docsforadobe.dev/items/compitem/#compitem-motiongraphicstemplatecontrollercount)
  - 新增：[CompItem.getMotionGraphicsTemplateControllerName()](https://ae-scripting.docsforadobe.dev/items/compitem/#compitem-getmotiongraphicstemplatecontrollername)
  - 新增：[CompItem.setMotionGraphicsControllerName()](https://ae-scripting.docsforadobe.dev/items/compitem/#compitem-setmotiongraphicscontrollername)
  - 新增：[Property.addToMotionGraphicsTemplateAs()](https://ae-scripting.docsforadobe.dev/properties/property/#property-addtomotiongraphicstemplateas)

---

### After Effects 16.0(2018 年 10 月)

- 脚本访问标记标签和 protectedRegion 属性：

  - 新增：[MarkerValue.label](https://ae-scripting.docsforadobe.dev/other/markervalue/#markervalue-label)
  - 新增：[MarkerValue.protectedRegion](https://ae-scripting.docsforadobe.dev/other/markervalue/#markervalue-protectedregion)

- 脚本访问其他项目颜色管理设置：

  - 新增：[Project.workingSpace](https://ae-scripting.docsforadobe.dev/general/project/#project-workingspace)
  - 新增：[Project.workingGamma](https://ae-scripting.docsforadobe.dev/general/project/#project-workinggamma)
  - 新增：[Project.listColorProfiles()](https://ae-scripting.docsforadobe.dev/general/project/#project-listcolorprofiles)
  - 新增：[Project.linearizeWorkingSpace](https://ae-scripting.docsforadobe.dev/general/project/#project-linearizeworkingspace)
  - 新增：[Project.compensateForSceneReferredProfiles](https://ae-scripting.docsforadobe.dev/general/project/#project-compensateforscenereferredprofiles)

- 脚本访问表达式引擎属性：

  - 新增：[Project.expressionEngine](https://ae-scripting.docsforadobe.dev/general/project/#project-expressionengine)

- 添加了项目方法[Project.setDefaultImportFolder](https://ae-scripting.docsforadobe.dev/general/project/#project-setdefaultimportfolder)，该方法设置了将在文件导入对话框中显示的文件夹。
- 添加了应用程序属性[app.disableRendering](https://ae-scripting.docsforadobe.dev/general/application/#app-disablerendering)，该功能通过与 Caps Lock 键相同的机制来禁用呈现。

---

### After Effects 15.1(2018 年 4 月)

- [Project.autoFixExpressions()](https://ae-scripting.docsforadobe.dev/general/project/#project-autofixexpressions)现在将用单引号(例如，(“效果名称”))和双引号修复表达式名称引用。
- 修复了[CompItem.exportAsMotionGraphicsTemplate()](https://ae-scripting.docsforadobe.dev/items/compitem/#compitem-exportasmotiongraphicstemplate)不按预期返回布尔值的问题

---

### After Effects 15.0(2017 年 10 月)

- 脚本访问运动图形模板

  - 添加：[CompItem.motionGraphicsTemplateName](https://ae-scripting.docsforadobe.dev/items/compitem/#compitem-motiongraphicstemplatename)
  - 新增：[CompItem.exportAsMotionGraphicsTemplate()](https://ae-scripting.docsforadobe.dev/items/compitem/#compitem-exportasmotiongraphicstemplate)
  - 新增：[CompItem.openInEssentialGraphics()](https://ae-scripting.docsforadobe.dev/items/compitem/#compitem-openinessentialgraphics)
  - 新增：[Property.addToMotionGraphicsTemplate()](https://ae-scripting.docsforadobe.dev/properties/property/#property-addtomotiongraphicstemplate)
  - 新增：[Property.canAddToMotionGraphicsTemplate()](https://ae-scripting.docsforadobe.dev/properties/property/#property-canaddtomotiongraphicstemplate)

---

### After Effects 14.2.1(CC 2017.2)(2017 年 6 月)

- ScriptUI 面板中的按钮已还原为 After Effects 14.1 和以前版本中看到的矩形外观。
- 该[AVItem.setProxyToNone()](https://ae-scripting.docsforadobe.dev/items/avitem/#avitem-setproxytonone)脚本方法不再失败，出现错误消息“After Effects 的错误：AEGP 试图添加无效的素材”。
- 现在，[System.callSystem()](https://ae-scripting.docsforadobe.dev/general/system/#system-callsystem)脚本方法等待命令调用的所有任务完成，而不是在命令花费很长时间才能完成时失败。

---

### After Effects 14.2(CC 2017.1)(2017 年 4 月)

- 脚本访问文本引导

  - 新增：[TextDocument.lead](https://ae-scripting.docsforadobe.dev/other/textdocument/#textdocument-leading)

- 脚本访问团队项目(测试版)

  - 新增：[Project.newTeamProject()](https://ae-scripting.docsforadobe.dev/general/project/#project-newteamproject)
  - 新增：[Project.openTeamProject()](https://ae-scripting.docsforadobe.dev/general/project/#project-openteamproject)
  - 新增：[Project.shareTeamProject()](https://ae-scripting.docsforadobe.dev/general/project/#project-shareteamproject)
  - 新增：[Project.syncTeamProject()](https://ae-scripting.docsforadobe.dev/general/project/#project-syncteamproject)
  - 新增：[Project.closeTeamProject()](https://ae-scripting.docsforadobe.dev/general/project/#project-closeteamproject)
  - 新增：[Project.convertTeamProjectToProject()](https://ae-scripting.docsforadobe.dev/general/project/#project-convertteamprojecttoproject)
  - 新增：[Project.listTeamProjects()](https://ae-scripting.docsforadobe.dev/general/project/#project-listteamprojects)
  - 新增：[Project.isTeamProjectOpen()](https://ae-scripting.docsforadobe.dev/general/project/#project-isteamprojectopen)
  - 新增：[Project.isAnyTeamProjectOpen()](https://ae-scripting.docsforadobe.dev/general/project/#project-isanyteamprojectopen)
  - 新增：[Project.isTeamProjectEnabled()](https://ae-scripting.docsforadobe.dev/general/project/#project-isteamprojectenabled)
  - 新增：[Project.isLoggedInToTeamProject()](https://ae-scripting.docsforadobe.dev/general/project/#project-isloggedintoteamproject)
  - 新增：[Project.isSyncCommandEnabled()](https://ae-scripting.docsforadobe.dev/general/project/#project-issynccommandenabled)
  - 新增：[Project.isShareCommandEnabled()](https://ae-scripting.docsforadobe.dev/general/project/#project-issharecommandenabled)
  - 新增：[Project.isResolveCommandEnabled()](https://ae-scripting.docsforadobe.dev/general/project/#project-isresolvecommandenabled)
  - 新增：[Project.resolveConflict()](https://ae-scripting.docsforadobe.dev/general/project/#project-resolveconflict)

- ScriptUI 面板中的下拉菜单不再位于 Windows 的 HiDPI 显示器上。
- ScriptUI 嵌入式面板中的按钮，滑块，显示三角形(“三角箭头”)，滚动条，进度条，单选按钮和复选框的外观已更新，以匹配 After Effects 本机控件的外观。
- 当 AVLayer.compPointToSource 脚本方法与 3D 文本层一起使用时，After Effects 不再崩溃。
- 快速框模糊效果的匹配名称是“ ADBE Box Blur2”。较早的匹配名称“ ADBE Box Blur”将继续起作用：当用于添加效果时，“ ADBE Box Blur”将应用“快速框模糊”效果，但较旧的名称为“ Box Blur”；Iterations 参数将设置为新的默认值 3。

---

### After Effects 14.0(CC 2017)(2016 年 11 月)

- 脚本访问工具

  - 新增：[Project.toolType](https://ae-scripting.docsforadobe.dev/general/project/#project-tooltype)

- 脚本访问合成标记

  - 添加：[CompItem.markerProperty](https://ae-scripting.docsforadobe.dev/items/compitem/#compitem-markerproperty)

- 脚本访问 AME 中的队列

  - 新增：[RenderQueue.queueInAME()](https://ae-scripting.docsforadobe.dev/renderqueue/renderqueue/#renderqueue-queueiname)

- 脚本访问可用的 GPU 加速选项

  - 新增：[app.availableGPUAccelTypes](https://ae-scripting.docsforadobe.dev/general/application/#app-availablegpuacceltypes)

---

### After Effects 13.8(CC 2015.3)(2016 年 6 月)

- 通过脚本启用 GPU 效果渲染

  - 新增：[Project.gpuAccelType](https://ae-scripting.docsforadobe.dev/general/project/#project-gpuacceltype)

- 添加了具有匹配名称的新的高斯模糊效果 `ADBE Gaussian Blur 2`

---

### After Effects 13.6(CC 2015)(2015 年 11 月)

- 脚本访问文本基准

  - 新增：[baselineLocs](https://ae-scripting.docsforadobe.dev/other/textdocument/#textdocument-baselinelocs)

- 生成随机数的新脚本方法

  - 添加：[generateRandomNumber()](https://ae-scripting.docsforadobe.dev/general/globals/#generaterandomnumber)

- 当图层具有父级时，使用[copyToComp()](https://ae-scripting.docsforadobe.dev/layers/layer/#layer-copytocomp)脚本方法不再导致 After Effects 崩溃。
- 现在，[valueAtTime()](https://ae-scripting.docsforadobe.dev/properties/property/#property-valueattime)脚本方法将等待时间密集的表达式(如)`sampleImage`在返回结果之前完成求值。
- ScriptUI 面板现在可以在 Windows 的高 DPI 显示器上正确显示并调整其大小。
- 在带有选项卡式面板的 scriptUI 对话框中单击“确定”或“取消”按钮时，After Effects 不再崩溃。

---

### After Effects 13.2(CC 2014.2)(2014 年 12 月)

- 文本层的脚本改进(只读)

  - 返回布尔值：

    - 新增：[fauxBold](https://ae-scripting.docsforadobe.dev/other/textdocument/#textdocument-fauxbold)
    - 新增：[fauxItalic](https://ae-scripting.docsforadobe.dev/other/textdocument/#textdocument-fauxitalic)
    - 新增：[allCaps](https://ae-scripting.docsforadobe.dev/other/textdocument/#textdocument-allcaps)
    - 补充：[小写](https://ae-scripting.docsforadobe.dev/other/textdocument/#textdocument-smallcaps)
    - 补充：[上标](https://ae-scripting.docsforadobe.dev/other/textdocument/#textdocument-superscript)
    - 添加：[下标](https://ae-scripting.docsforadobe.dev/other/textdocument/#textdocument-subscript)

  - 返回浮点数：

    - 新增：[verticalScale](https://ae-scripting.docsforadobe.dev/other/textdocument/#textdocument-verticalscale)
    - 新增：[horizo​​ntalScale](https://ae-scripting.docsforadobe.dev/other/textdocument/#textdocument-horizontalscale)
    - 新增：[baselineShift](https://ae-scripting.docsforadobe.dev/other/textdocument/#textdocument-baselineshift)
    - 添加：[tsume](https://ae-scripting.docsforadobe.dev/other/textdocument/#textdocument-tsume)

  - 返回[[X，Y])个位置坐标的数组(仅适用于段落文本图层)：

    - 新增：[boxTextPos](https://ae-scripting.docsforadobe.dev/other/textdocument/#textdocument-boxtextpos)

- 图层空间/比较空间转换：

  - 新增：[sourcePointToComp()](https://ae-scripting.docsforadobe.dev/layers/avlayer/#avlayer-sourcepointtocomp)
  - 补充：compPointToSource()

---

### After Effects 13.1(CC 2014.1)(2014 年 9 月)

- 文本层的脚本改进(只读)

  - 返回字符串：

    - 新增：[fontLocation](https://ae-scripting.docsforadobe.dev/other/textdocument/#textdocument-fontlocation)
    - 新增：[fontStyle](https://ae-scripting.docsforadobe.dev/other/textdocument/#textdocument-fontstyle)
    - 新增：[fontFamily](https://ae-scripting.docsforadobe.dev/other/textdocument/#textdocument-fontfamily)

- 实施了“使用旧版用户界面”切换

---

### After Effects 13.0(CC 2014)(2014 年 6 月)

- 脚本访问渲染设置和输出模块设置

  - 新增：RenderQueueItem 对象[getSetting](https://ae-scripting.docsforadobe.dev/renderqueue/renderqueueitem/#renderqueueitem-getsetting)，[setSetting](https://ae-scripting.docsforadobe.dev/renderqueue/renderqueueitem/#renderqueueitem-setsetting)方法
  - 新增：RenderQueueItem 对象[getSettings](https://ae-scripting.docsforadobe.dev/renderqueue/renderqueueitem/#renderqueueitem-getsettings)，[setSettings](https://ae-scripting.docsforadobe.dev/renderqueue/renderqueueitem/#renderqueueitem-setsettings)方法
  - 补充：OutputModule 对象的[getSetting](https://ae-scripting.docsforadobe.dev/renderqueue/outputmodule/#outputmodule-getsetting)，[setSetting](https://ae-scripting.docsforadobe.dev/renderqueue/outputmodule/#outputmodule-setsetting)方法
  - 补充：OutputModule 对象的[getSettings](https://ae-scripting.docsforadobe.dev/renderqueue/outputmodule/#outputmodule-getsettings)，[setSettings](https://ae-scripting.docsforadobe.dev/renderqueue/outputmodule/#outputmodule-setsettings)方法

- 通过 ID 获取项目项目：[Project.itemByID()](https://ae-scripting.docsforadobe.dev/general/project/#project-itembyid)
- 实施 CEP 面板

---

### After Effects 12.0(CC)(2013 年 6 月)

- 访问效果的内部版本字符串

  - 补充：应用程序效果对象的 version 属性，请参阅[app.effects](https://ae-scripting.docsforadobe.dev/general/application/#app-effects)

- 能够获取和设置预览模式

  - 新增：[Viewer.fastPreview](https://ae-scripting.docsforadobe.dev/other/viewer/#viewer-fastpreview)

- 获取层采样方法(请参阅[amplingQuality](https://ae-scripting.docsforadobe.dev/layers/avlayer/#avlayer-samplingquality))
- 更改了首选项和设置方法(请参阅[“设置”对象](https://ae-scripting.docsforadobe.dev/other/settings/#settings))
- ScriptUI 现在基于与主应用程序相同的控件。

---

### After Effects 11.0(CS6)(2012 年 4 月)

- 补充：访问[查看器对象](https://ae-scripting.docsforadobe.dev/other/viewer/#viewer)对象和控件

  - 新增：[app.activeViewer](https://ae-scripting.docsforadobe.dev/general/application/#app-activeviewer)
  - 新增：[AVLayer.openInViewer()](https://ae-scripting.docsforadobe.dev/layers/avlayer/#avlayer-openinviewer)在图层查看器中打开图层
  - 添加：[CompItem.openInViewer()](https://ae-scripting.docsforadobe.dev/items/compitem/#compitem-openinviewer)在合成查看器中打开合成
  - 添加：[footageItem.openInViewer()](https://ae-scripting.docsforadobe.dev/items/footageitem/#footageitem-openinviewer)在素材查看器中打开素材项目

- 新增：[Property.canSetExpression](https://ae-scripting.docsforadobe.dev/properties/property/#property-cansetexpression)
- 新增：[AVLayer.environmentLayer](https://ae-scripting.docsforadobe.dev/layers/avlayer/#avlayer-environmentlayer)
- 新增：[MaskPropertyGroup.maskFeatherFalloff](https://ae-scripting.docsforadobe.dev/properties/maskpropertygroup/#maskpropertygroup-maskfeatherfalloff)
- 通过脚本访问 Shape Feather 属性

  - 新增：[Shape.featherSegLocs](https://ae-scripting.docsforadobe.dev/other/shape/#shape-featherseglocs)
  - 新增：[Shape.featherRelSegLocs](https://ae-scripting.docsforadobe.dev/other/shape/#shape-featherrelseglocs)
  - 新增：[Shape.featherRadii](https://ae-scripting.docsforadobe.dev/other/shape/#shape-featherradii)
  - 新增：[Shape.featherInterps](https://ae-scripting.docsforadobe.dev/other/shape/#shape-featherinterps)
  - 新增：[Shape.featherTensions](https://ae-scripting.docsforadobe.dev/other/shape/#shape-feathertensions)
  - 新增：[Shape.featherTypes](https://ae-scripting.docsforadobe.dev/other/shape/#shape-feathertypes)
  - 新增：[Shape.featherRelCornerAngles](https://ae-scripting.docsforadobe.dev/other/shape/#shape-featherrelcornerangles)

---

### After Effects 10.5(CS5.5)(2011 年 4 月)

- 添加到[Project 对象的](https://ae-scripting.docsforadobe.dev/general/project/#project)对象：

  - [Project.framesCountType](https://ae-scripting.docsforadobe.dev/general/project/#project-framescounttype)
  - [Project.feetFramesFilmType](https://ae-scripting.docsforadobe.dev/general/project/#project-feetframesfilmtype)
  - [Project.framesUseFeetFrames](https://ae-scripting.docsforadobe.dev/general/project/#project-framesusefeetframes)
  - [Project.footageTimecodeDisplayStartType](https://ae-scripting.docsforadobe.dev/general/project/#project-footagetimecodedisplaystarttype)
  - [Project.timeDisplayType](https://ae-scripting.docsforadobe.dev/general/project/#project-timedisplaytype)

- 从[Project 对象中](https://ae-scripting.docsforadobe.dev/general/project/#project)删除的对象：

  - `timecodeDisplayType` 属性
  - `timecodeBaseType` 属性
  - `timecodeNTSCDropFrame` 属性
  - `timecodeFilmType` 属性
  - `TimecodeDisplayType` 枚举
  - `TimecodeFilmType` 枚举
  - `TimecodeBaseType` 枚举

- 新增：[CompItem.dropFrame](https://ae-scripting.docsforadobe.dev/items/compitem/#compitem-dropframe)
- 添加了对“段落框文本”的支持：

  - 添加了[LayerCollection.addBoxText()](https://ae-scripting.docsforadobe.dev/layers/layercollection/#layercollection-addboxtext)
  - 添加了[TextDocument.boxText](https://ae-scripting.docsforadobe.dev/other/textdocument/#textdocument-boxtext)
  - 添加了[TextDocument.pointText](https://ae-scripting.docsforadobe.dev/other/textdocument/#textdocument-pointtext)
  - 添加了[TextDocument.boxTextSize](https://ae-scripting.docsforadobe.dev/other/textdocument/#textdocument-boxtextsize)

- 添加了[LightLayer.lightType](https://ae-scripting.docsforadobe.dev/layers/lightlayer/#lightlayer-lighttype)

---

### After Effects 9.0(CS4)(2008 年 9 月)

- 新增：[app.isoLanguage](https://ae-scripting.docsforadobe.dev/general/application/#app-isolanguage)
- 新增：[MarkerValue.duration](https://ae-scripting.docsforadobe.dev/other/markervalue/#markervalue-duration)
- 新增：[OutputModule.includeSourceXMP](https://ae-scripting.docsforadobe.dev/renderqueue/outputmodule/#outputmodule-includesourcexmp)
- 新增：[Project.xmpPacket](https://ae-scripting.docsforadobe.dev/general/project/#project-xmppacket)
- 添加了以下与“单独尺寸”功能相关的属性方法和属性：

  - [属性尺寸分离](https://ae-scripting.docsforadobe.dev/properties/property/#property-dimensionsseparated)
  - [Property.getSeparationFollower()](https://ae-scripting.docsforadobe.dev/properties/property/#property-getseparationfollower)
  - [Property.isSeparation 跟随](https://ae-scripting.docsforadobe.dev/properties/property/#property-isseparationfollower)
  - [Property.isSeparationLeader](https://ae-scripting.docsforadobe.dev/properties/property/#property-isseparationleader)
  - [Property.separationDimension](https://ae-scripting.docsforadobe.dev/properties/property/#property-separationdimension)
  - [Property.separationLeader](https://ae-scripting.docsforadobe.dev/properties/property/#property-separationleader)

- 添加了[TextDocument 对象](https://ae-scripting.docsforadobe.dev/other/textdocument/#textdocument)访问权限，包括：

  - 新增：[TextDocument.applyFill](https://ae-scripting.docsforadobe.dev/other/textdocument/#textdocument-applyfill)
  - 新增：[TextDocument.applyStroke](https://ae-scripting.docsforadobe.dev/other/textdocument/#textdocument-applystroke)
  - 新增：[TextDocument.fillColor](https://ae-scripting.docsforadobe.dev/other/textdocument/#textdocument-fillcolor)
  - 新增：[TextDocument.font](https://ae-scripting.docsforadobe.dev/other/textdocument/#textdocument-font)
  - 新增：[TextDocument.fontSize](https://ae-scripting.docsforadobe.dev/other/textdocument/#textdocument-fontsize)
  - 添加：[TextDocument.justification](https://ae-scripting.docsforadobe.dev/other/textdocument/#textdocument-justification)
  - 新增：[TextDocument.resetCharStyle()](https://ae-scripting.docsforadobe.dev/other/textdocument/#textdocument-resetcharstyle)
  - 新增：[TextDocument.resetParagraphStyle()](https://ae-scripting.docsforadobe.dev/other/textdocument/#textdocument-resetparagraphstyle)
  - 新增：[TextDocument.strokeColor](https://ae-scripting.docsforadobe.dev/other/textdocument/#textdocument-strokecolor)
  - 新增：[TextDocument.strokeOverFill](https://ae-scripting.docsforadobe.dev/other/textdocument/#textdocument-strokeoverfill)
  - 新增：[TextDocument.strokeWidth](https://ae-scripting.docsforadobe.dev/other/textdocument/#textdocument-strokewidth)
