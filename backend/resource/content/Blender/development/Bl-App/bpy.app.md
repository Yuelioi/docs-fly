---
title: 应用程序数据 (bpy.app)
order: 5
---

# 应用程序数据 (bpy.app)

## 应用程序数据 (bpy.app)

包含在运行时保持不变的应用程序值。大部分都是开发函数，可以康康后面的子模块

### 子模块

- 应用程序处理程序 (bpy.app.handlers)
- 应用程序翻译 (bpy.app.translations)
- 应用程序图标 (bpy.app.icons)
- 应用程序计时器 (bpy.app.timers)

### autoexec_fail

全名：bpy.app.autoexec_fail

说明：暂无说明

### autoexec_fail_message

全名：bpy.app.autoexec_fail_message

说明：暂无说明

### autoexec_fail_quiet

全名：bpy.app.autoexec_fail_quiet

说明：暂无说明

### debug

全名：bpy.app.debug

说明：Boolean, for debug info (started with –debug / –debug\_\* matching this
attribute name)

### debug_depsgraph

全名：bpy.app.debug_depsgraph

说明：Boolean, for debug info (started with –debug / –debug\_\* matching this
attribute name)

### debug_depsgraph_build

全名：bpy.app.debug_depsgraph_build

说明：Boolean, for debug info (started with –debug / –debug\_\* matching this
attribute name)

### debug_depsgraph_eval

全名：bpy.app.debug_depsgraph_eval

说明：Boolean, for debug info (started with –debug / –debug\_\* matching this
attribute name)

### debug_depsgraph_pretty

全名：bpy.app.debug_depsgraph_pretty

说明：Boolean, for debug info (started with –debug / –debug\_\* matching this
attribute name)

### debug_depsgraph_tag

全名：bpy.app.debug_depsgraph_tag

说明：Boolean, for debug info (started with –debug / –debug\_\* matching this
attribute name)

### debug_depsgraph_time

全名：bpy.app.debug_depsgraph_time

说明：Boolean, for debug info (started with –debug / –debug\_\* matching this
attribute name)

### debug_events

全名：bpy.app.debug_events

说明：Boolean, for debug info (started with –debug / –debug\_\* matching this
attribute name)

### debug_ffmpeg

全名：bpy.app.debug_ffmpeg

说明：Boolean, for debug info (started with –debug / –debug\_\* matching this
attribute name)

### debug_freestyle

全名：bpy.app.debug_freestyle

说明：Boolean, for debug info (started with –debug / –debug\_\* matching this
attribute name)

### debug_handlers

全名：bpy.app.debug_handlers

说明：Boolean, for debug info (started with –debug / –debug\_\* matching this
attribute name)

### debug_io

全名：bpy.app.debug_io

说明：Boolean, for debug info (started with –debug / –debug\_\* matching this
attribute name)

### debug_python

全名：bpy.app.debug_python

说明：Boolean, for debug info (started with –debug / –debug\_\* matching this
attribute name)

### debug_simdata

全名：bpy.app.debug_simdata

说明：Boolean, for debug info (started with –debug / –debug\_\* matching this
attribute name)

### debug_value

全名：bpy.app.debug_value

说明：Short, number which can be set to non-zero values for testing purposes

### debug_wm

全名：bpy.app.debug_wm

说明：Boolean, for debug info (started with –debug / –debug\_\* matching this
attribute name)

### driver_namespace

全名：bpy.app.driver_namespace

说明：Dictionary for drivers namespace, editable in-place, reset on file load
(read-only)

### render_icon_size

全名：bpy.app.render_icon_size

说明：Reference size for icon/preview renders (read-only)

### render_preview_size

全名：bpy.app.render_preview_size

说明：Reference size for icon/preview renders (read-only)

### tempdir

全名：bpy.app.tempdir

说明：String, the temp directory used by blender (read-only)

### use_event_simulate

全名：bpy.app.use_event_simulate

说明：Boolean, for application behavior (started with –enable-\* matching this
attribute name)

### use_userpref_skip_save_on_exit

全名：bpy.app.use_userpref_skip_save_on_exit

说明：Boolean, for application behavior (started with –enable-\* matching this
attribute name)

### background

全名：bpy.app.background

说明：Boolean, True when blender is running without a user interface (started
with -b)

### factory_startup

全名：bpy.app.factory_startup

说明：Boolean, True when blender is running with –factory-startup)

### build_branch

全名：bpy.app.build_branch

说明：当前 blender 来源是由哪个分支

### build_cflags

全名：bpy.app.build_cflags

说明：C 语言编译器的标志

### build_commit_date

全名：bpy.app.build_commit_date

说明：The date of commit this blender instance was built

### build_commit_time

全名：bpy.app.build_commit_time

说明：当前 blender 实例的提交日期

### build_cxxflags

全名：bpy.app.build_cxxflags

说明：C++编译器标志

### build_date

全名：bpy.app.build_date

说明：The date this blender instance was built

### build_hash

全名：bpy.app.build_hash

说明：当前 blender 实例的建立日期

### build_linkflags

全名：bpy.app.build_linkflags

说明：二进制连接标志

### build_platform

全名：bpy.app.build_platform

说明：当前 Blender 实例是为哪个平台建立的

### build_system

全名：bpy.app.build_system

说明：使用的操作系统

### build_time

全名：bpy.app.build_time

说明：当前 Blender 实例的建立时间

### build_type

全名：bpy.app.build_type

说明：构建的类型（发布、调试）。

### build_commit_timestamp

全名：bpy.app.build_commit_timestamp

说明：当前 Blender 实例建立时的 unix 时间戳。

### binary_path

全名：bpy.app.binary_path

说明：Blender 的可执行文件的位置，对打开新实例的实用程序很有用。

### version_char

全名：bpy.app.version_char

说明：已弃用，空字符串

### version_cycle

全名：bpy.app.version_cycle

说明：该版本的发布状态为 alpha/beta/rc/release

### version_string

全名：bpy.app.version_string

说明：格式化为字符串的 Blender 版本

### version

全名：bpy.app.version

说明：Blender 版本由 3 个数字组成的元组，例如：(2, 83, 1)。

### version_file

全名：bpy.app.version_file

说明：Blender 版本，作为一个元组，最后用于保存一个.blend 文件，与 bpy.data.version 兼容。当前值应该被用来处理 Blender 版本之间的兼容性变化。

### alembic

全名：bpy.app.alembic

说明：常量

### alembic()

全名：bpy.app.alembic(supported=True, version=(1, 7, 16), version_string=’ 1, 7,
16’)

说明：

### build_options

全名：bpy.app.build_options

说明：常量

### build_options()

全名：bpy.app.build_options(bullet=True, codec_avi=True, codec_ffmpeg=True,codec_sndfile=True, compositor=True, cycles=True, cycles_osl=True,freestyle=True, image_cineon=True, image_dds=True, image_hdr=True,image_openexr=True, image_openjpeg=True, image_tiff=True, input_ndof=True,audaspace=True, international=True, openal=True, opensubdiv=True, sdl=True,sdl_dynload=True, coreaudio=False, jack=False, pulseaudio=True, wasapi=False,libmv=True, mod_oceansim=True, mod_remesh=True, collada=True,opencolorio=True, openmp=True, openvdb=True, alembic=True, usd=True,fluid=True, xr_openxr=True, potrace=True, pugixml=True, haru=True)

说明：

### ffmpeg

全名：bpy.app.ffmpeg

说明：常量

### ffmpeg()

全名：bpy.app.ffmpeg(supported=True, avcodec_version=(58, 134, 100),avcodec_version_string=’58, 134, 100’, avdevice_version=(58, 13, 100),avdevice_version_string=’58, 13, 100’, avformat_version=(58, 76, 100),avformat_version_string=’58, 76, 100’, avutil_version=(56, 70, 100),avutil_version_string=’56, 70, 100’, swscale_version=(5, 9, 100),swscale_version_string=’ 5, 9, 100’)

说明：

### ocio

全名：bpy.app.ocio

说明：常量

### ocio()

全名：bpy.app.ocio(supported=True, version=(2, 0, 0), version_string=’ 2, 0, 0’)

说明：

### oiio

全名：bpy.app.oiio

说明：常量

### oiio()

全名：bpy.app.oiio(supported=True, version=(2, 2, 15), version_string=’ 2, 2,
15’)

说明：

### opensubdiv

全名：bpy.app.opensubdiv

说明：Constant value

### opensubdiv()

全名：bpy.app.opensubdiv(supported=True, version=(0, 0, 0), version_string=’ 0,
0, 0’)

说明：

### openvdb

全名：bpy.app.openvdb

说明：常量

### openvdb()

全名：bpy.app.openvdb(supported=True, version=(8, 0, 0), version_string=’ 8, 0,
0’)

说明：

### sdl

全名：bpy.app.sdl

说明：常量

### sdl()

全名：bpy.app.sdl(supported=True, version=(0, 0, 0), version_string=’Unknown’,
available=False)

说明：

### usd

全名：bpy.app.usd

说明：常量

### usd()

全名：bpy.app.usd(supported=True, version=(0, 21, 2), version_string=’ 0, 21, 2’)

说明：

## 处理程序 (bpy.app.handlers)

本模块包含回调列表

### 基本处理程序实例

此脚本展示添加处理程序的最简单示例。

```python
import bpy


def my_handler(scene):
    print("Frame Change", scene.frame_current)


bpy.app.handlers.frame_change_pre.append(my_handler)

```

### 持久处理程序实例

默认情况下，处理程序在加载新文件时被释放，在某些情况下，可能希望处理程序在多个文件中保持运行（例如，当处理程序是插件的一部分时）。

为此，需要使用 bpy.app.handlers.persistent 装饰器。

```python
import bpy
from bpy.app.handlers import persistent


@persistent
def load_handler(dummy):
    print("Load Handler:", bpy.data.filepath)


bpy.app.handlers.load_post.append(load_handler)

```

### 关于修改数据的说明

改变处理程序数据应谨慎进行。在渲染过程中，frame_change_pre 和 frame_change_post 处理程序是由一个线程调用的，而视窗更新是由另一个线程进行的。

如果处理程序改变了被视口访问的数据，可能导致 Blender 崩溃。此时，请在开始渲染前锁定接口（Render → Lock Interface 或 bpy.types.RenderSettings.use_lock_interface）。

下面是一个被处理程序改变的网格的例子。

```python
def frame_change_pre(scene):
    # 一个在Z方向移动的三角形
    zshift = scene.frame_current * 0.1
    vertices = [(-1, -1, zshift), (1, -1, zshift), (0, 1, zshift)]
    triangles = [(0, 1, 2)]

    object = bpy.data.objects["The Object"]
    object.data.clear_geometry()
    object.data.from_pydata(vertices, [], triangles)


```

### handlers.depsgraph_update_post

全名：bpy.app.handlers.depsgraph_update_post

描述：关于 depsgraph 的更新（post）

### handlers.depsgraph_update_pre

全名：bpy.app.handlers.depsgraph_update_pre

描述：在 depsgraph 更新（pre）

### handlers.frame_change_post

全名：bpy.app.handlers.frame_change_post

描述：在回放和渲染的帧变化后，在对新帧进行数据评估后调用。

### handlers.frame_change_pre

全名：bpy.app.handlers.frame_change_pre

描述：在回放和渲染的帧变化之后，在为新帧评估任何数据之前调用。

这样可以改变新帧的数据和关系（例如，将一个对象换成一个网格）。

请注意，这个处理程序不能作为 "帧变化之前 "的事件使用。在这个处理程序中依赖关系图不可用，因为数据和关系可能已经改变，依赖关系图还没有被更新。

### handlers.load_factory_preferences_post

全名：bpy.app.handlers.load_factory_preferences_post

描述：在加载工厂偏好时（after）

### handlers.load_factory_startup_post

全名：bpy.app.handlers.load_factory_startup_post

描述：在加载工厂启动时 (after)

### handlers.load_post

全名：bpy.app.handlers.load_post

描述：加载新 Blender 文件(after)

### handlers.load_pre

全名：bpy.app.handlers.load_pre

描述：加载新 Blender (before)

### handlers.redo_post

全名：bpy.app.handlers.redo_post

描述：加载重做步骤(after)

### handlers.redo_pre

全名：bpy.app.handlers.redo_pre

描述：加载重做步骤 (before)

### handlers.render_cancel

全名：bpy.app.handlers.render_cancel

描述：在取消渲染工作时

### handlers.render_complete

全名：bpy.app.handlers.render_complete

描述：在完成渲染工作后

### handlers.render_init

全名：bpy.app.handlers.render_init

描述：初始化渲染工作时

### handlers.render_post

全名：bpy.app.handlers.render_post

描述：正在渲染 (after)

### handlers.render_pre

全名：bpy.app.handlers.render_pre

描述：正在渲染 (before)

### handlers.render_stats

全名：bpy.app.handlers.render_stats

描述：正在打印渲染统计数据

### handlers.render_write

全名：bpy.app.handlers.render_write

描述：在写入渲染帧时（直接在写完帧之后）。

### handlers.save_post

全名：bpy.app.handlers.save_post

描述：在保存 BL 文件 (after)

### handlers.save_pre

全名：bpy.app.handlers.save_pre

描述：在保存 BL 文件 (before)

### handlers.undo_post

全名：bpy.app.handlers.undo_post

描述：在加载撤销步骤 (after)

### handlers.undo_pre

全名：bpy.app.handlers.undo_pre

描述：在加载撤销步骤(before)

### handlers.version_update

全名：bpy.app.handlers.version_update

描述：结束版本控制代码

### handlers.xr_session_start_pre

全名：bpy.app.handlers.xr_session_start_pre

描述：启动 XR 会话 (before)

### handlers.persistent

全名：bpy.app.handlers.persistent

描述：加载新文件时不删除回调函数的函数装饰器

## 应用程序翻译 (bpy.app.translations)

关于 Blender 国际化的数据/方法，并允许 py 脚本为自己的用户界面信息提供翻译。

### 简介

这个对象的大部分内容只有在你从 Python 中实际操作 i18n 的东西时才会有用。

如果是普通插件，应该只关心 contexts 成员，以及 register()/unregister()函数! pgettext()系列函数只在特殊情况下使用
(例如，复杂的 "合成 "UI 字符串...)。

要在你的 Python 脚本中添加翻译，你必须定义一个字典，其格式如下。{locale: {msg_key: msg_translation, ...},
...} 。

- locale 是 lang iso 代码（如 fr），lang+country 代码（如 pt_BR），lang+variant 代码（如 sr@latin），或者完整代码（如 uz_UZ@cyrilic）。
- msg_key：元组(context, org message) ，尽可能地使用预定义的 context。
- msg_translation：是给定语言的翻译信息

然后，在你的 register() 函数中调用 bpy.app.translations.register(**name**, your_dict)，并在你的 unregister() 函数中调用 bpy.app.translations.unregister(**name**) 。

管理用户界面翻译插件有几个功能，可以帮助你收集要翻译的字符串，并生成所需的 python 代码（翻译字典），以及可选中介的 po 文件（如果你想的话）......更多信息请参见如何翻译 Blender 和在 Blender 代码中使用 i18n。

### 模块参考

```python
import bpy

# 这块可以由UI翻译插件自动生成，它也可以处理与PO格式的转换。
# 另请看https://wiki.blender.org/wiki/Process/Translate_Blender#Translating_non-official_addons
# 它也可以（应该）放在一个不同的、特定的py文件中。

# ##### 开始自动生成的i18n部分 #####
# 注意：你可以安全地在这个自动生成的区块周围移动（有开始/结束标记！）。
# 并用手编辑翻译。
# 只要小心尊重元组的格式即可!

# 元组的元组 ((msgctxt, msgid), (sources, gen_comments), (lang, translation, (is_fuzzy, comments)), ...)

# Tuple of tuples ((msgctxt, msgid), (sources, gen_comments), (lang, translation, (is_fuzzy, comments)), ...)
translations_tuple = (
    (("*", ""),
     ((), ()),
     ("fr_FR", "Project-Id-Version: Copy Settings 0.1.5 (r0)\nReport-Msgid-Bugs-To: \nPOT-Creation-Date: 2013-04-18 15:27:45.563524\nPO-Revision-Date: 2013-04-18 15:38+0100\nLast-Translator: Bastien Montagne <montagne29@wanadoo.fr>\nLanguage-Team: LANGUAGE <LL@li.org>\nLanguage: __POT__\nMIME-Version: 1.0\nContent-Type: text/plain; charset=UTF-8\nContent-Transfer-Encoding: 8bit\n",
               (False,
                ("Blender's translation file (po format).",
                 "Copyright (C) 2013 The Blender Foundation.",
                 "This file is distributed under the same license as the Blender package.",
                 "FIRST AUTHOR <EMAIL@ADDRESS>, YEAR."))),
     ),
    (("Operator", "Render: Copy Settings"),
     (("bpy.types.SCENE_OT_render_copy_settings",),
      ()),
     ("fr_FR", "Rendu : copier réglages",
               (False, ())),
     ),
    (("*", "Copy render settings from current scene to others"),
     (("bpy.types.SCENE_OT_render_copy_settings",),
      ()),
     ("fr_FR", "Copier les réglages de rendu depuis la scène courante vers d’autres",
               (False, ())),
     ),
    # ... etc, all messages from your addon.
)

translations_dict = {}
for msg in translations_tuple:
    key = msg[0]
    for lang, trans, (is_fuzzy, comments) in msg[2:]:
        if trans and not is_fuzzy:
            translations_dict.setdefault(lang, {})[key] = trans

# ##### END AUTOGENERATED I18N SECTION #####

# Define remaining addon (operators, UI...) here.


def register():
   # Usual operator/UI/etc. registration...

    bpy.app.translations.register(__name__, translations_dict)


def unregister():
    bpy.app.translations.unregister(__name__)

   # Usual operator/UI/etc. unregistration...

```

### translations.locale

全名：bpy.app.translations.locale

说明：当前正在使用的实际语言（当 Blender 没有国际化支持时，将返回无效字符串）。

### translations.locales

全名：bpy.app.translations.locales

说明：目前 Blender 已知的所有地区语言（即可作为翻译的语言）。

### translations.contexts_C_to_py

全名：bpy.app.translations.contexts_C_to_py

说明：一个只读的字典，将上下文的 C-identifiers 映射到它们的 Py-identifiers。

### translations.contexts

全名：bpy.app.translations.contexts

说明：常量， bpy.app.translations.contexts(default_real=None, default=’\*’,operator_default=’Operator’, ui_events_keymaps=’UI_Events_KeyMaps’,plural=’Plural’, id_action=’Action’, id_armature=’Armature’, id_brush=’Brush’,id_camera=’Camera’, id_cachefile=’CacheFile’, id_collection=’Collection’,id_curve=’Curve’, id_fs_linestyle=’FreestyleLineStyle’, id_gpencil=’GPencil’,id_hair=’Hair’, id_id=’ID’, id_image=’Image’, id_shapekey=’Key’,id_light=’Light’, id_library=’Library’, id_lattice=’Lattice’, id_mask=’Mask’,id_material=’Material’, id_metaball=’Metaball’, id_mesh=’Mesh’,id_movieclip=’MovieClip’, id_nodetree=’NodeTree’, id_object=’Object’,id_paintcurve=’PaintCurve’, id_palette=’Palette’,id_particlesettings=’ParticleSettings’, id_pointcloud=’PointCloud’,id_lightprobe=’LightProbe’, id_scene=’Scene’, id_screen=’Screen’,id_sequence=’Sequence’, id_simulation=’Simulation’, id_speaker=’Speaker’,id_sound=’Sound’, id_texture=’Texture’, id_text=’Text’, id_vfont=’VFont’,id_volume=’Volume’, id_world=’World’, id_workspace=’WorkSpace’,id_windowmanager=’WindowManager’)

### translations.locale_explode()

全名：bpy.app.translations.locale_explode(locale)

说明：返回所有组件和在给定的 ISO 地区字符串中的组合。

```python
>>>bpy.app.translations.locale_explode("sr_RS@latin")
("sr", "RS", "latin", "sr_RS", "sr@latin")

```

对于不完整的语言区域，缺少的元素将是 None。

参数：locale - 要翻译的 ISO locale 字符串。

返回：元组，`(language, country, variant, language_country, language@variant)`

### translations.pgettext()

全名：bpy.app.translations.pgettext( _msgid_ , _msgctxt =None_)

说明：尝试翻译给定的 msgid（可选 msgctxt）。

::: tip
常规插件少用这个函数，因为所有的翻译都应该由 Blender 内部代码处理。除非是格式化字符串（如"文件：%r"），但也应该使用 pgettext_iface()/pgettext_tip()!
:::

参数

- msgid (string)：要翻译的字符串。
- msgctxt (string or None) ：翻译背景（默认为 BLT_I18NCONTEXT_DEFAULT）。

返回：翻译后的字符串（如果没找到翻译，则为 msgid）。

### translations.pgettext_data()

全名：bpy.app.translations.pgettext_data(msgid, msgctxt=None)

说明：如果新数据名称的翻译被启用，尝试翻译给定的 msgid（msgctxt 可选）。

其他：跟 pgettext()一样，别用

### translations.pgettext_iface()

全名：bpy.app.translations.pgettext_iface(msgid, msgctxt=None)

说明：如果标签的翻译功能被启用，尝试翻译给定的 msgid（可选 msgctxt）。

其他：跟 pgettext()一样，别用

### translations.pgettext_tip()

全名：bpy.app.translations.pgettext_tip(msgid, msgctxt=None)

说明：如果工具提示的翻译功能被启用，尝试翻译给定的 msgid（可选 msgctxt）。

其他：跟 pgettext()一样，别用

### translations.register()

全名：bpy.app.translations.register(module_name, translations_dict)

说明：注册插件的 UI 翻译。

其他：在 Blender 没有国际化支持的情况下，没有用。

参数

- module_name (string) ：识别插件的名称。
- translations_dict (dict) ： 字典，如{locale: {msg_key: msg_translation, ...}, ...}。

### translations.unregister()

全名：bpy.app.translations.unregister(module_name)

说明：注销插件的 UI 翻译。

其他：在 Blender 没有国际化支持的情况下，没有用。

参数

- module_name (string) ：插件名称。

## 应用程序图标 (bpy.app.icons)

### icons.new_triangles()

全名：bpy.app.icons.new_triangles(range, coords, colors)

说明：从三角形的几何形状创建一个新的图标。

参数

- range (元组)：一对整数。
- coords (byte sequence.) ： X, Y 坐标的字节序列（一个三角形有 6 个浮点数）。
- colors (byte sequence.) ：RGBA 的字节序列（一个三角形为 12 个浮点）。

返回：唯一的图标值（传递给接口 icon_value 参数）。

返回类型：int

### icons.new_triangles_from_file()

全名：bpy.app.icons.new_triangles(filename)

说明：从三角形的几何形状创建一个新的图标。

参数：

- filename ：文件路径，字符串。

返回：唯一的图标值（传递给接口 icon_value 参数）。

返回类型：int

### icons.new_triangles()

全名：bpy.app.icons.new_triangles(range, coords, colors)

说明：释放图标。

## 应用程序计时器 (bpy.app.timers)

### 在 x 秒内运行一个函数

```python
import bpy

def in_5_seconds():
    print("Hello World")

bpy.app.timers.register(in_5_seconds, first_interval=5)

```

### 每 x 秒运行一个函数

```python
import bpy

def every_2_seconds():
    print("Hello World")
    return 2.0

bpy.app.timers.register(every_2_seconds)
```

### 每 x 秒运行 n 次函数

```python
import bpy

counter = 0


def run_10_times():
    global counter
    counter += 1
    print(counter)
    if counter == 10:
        return None
    return 0.1


bpy.app.timers.register(run_10_times)
```

### 将参数分配给函数

```python
import bpy
import functools


def print_message(message):
    print("Message:", message)


bpy.app.timers.register(functools.partial(print_message, "Hello"), first_interval=2.0)
bpy.app.timers.register(functools.partial(print_message, "World"), first_interval=3.0)

```

### 使用 Timer 对其他线程事件做出反应

```python
import bpy
import queue

execution_queue = queue.Queue()

# 在另一个线程中，该函数可以安全被调用。
# 在定时器下次运行时，该函数将被执行。
def run_in_main_thread(function):
    execution_queue.put(function)


def execute_queued_functions():
    while not execution_queue.empty():
        function = execution_queue.get()
        function()
    return 1.0


bpy.app.timers.register(execute_queued_functions)

```

### timers.is_registered()

全名：bpy.app.timers.is_registered(function)

说明：检查该函数是否被注册为一个定时器。

参数：function (int) ，要检查的函数。

返回：被注册时为真，否则为假。

返回类型：bool

### timers.register()

全名：bpy.app.timers.is_registered(function)

说明：添加一个新函数，该函数将在指定秒数后被调用。

该函数没有得到任何参数，预计将返回 None 或一浮点数。

如果返回 None，定时器将被取消注册。返回的数字指定了再次调用该函数之前的延迟时间。 functools.partial 可以用来分配一些参数。

参数

- function (Callable[[], Union[float, None]) ：应该调用的函数。
- first_interval (float) ：距离回调第一次被调用的秒数。
- persistent (bool) ：当新文件加载时，不要删除定时器。

### timers.unregister()

全名：bpy.app.timers.is_registered(function)

说明：注销定时器。

参数：function (函数) ，要取消注册的函数。
