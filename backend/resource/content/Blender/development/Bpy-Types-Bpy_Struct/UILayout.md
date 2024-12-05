---
title: UILayout
order: 12
---

# UILayout

## Description

[class
bpy.types.UILayout(bpy_struct)](https://docs.blender.org/api/master/bpy.types.UILayout.html)
基类：`bpy_struct`

### activate_init 不懂

说明：为 True 时，弹出窗口定义的按钮将在第一次显示时被激活（使用时，可以在一个字段中输入，而不必先点击它）。
类型：布尔值，默认为 False

### active 不懂

说明：
类型：布尔值，默认为 False

### active_default 不懂

说明：为 True 时，在此之后定义的操作项，将在按下回车键后激活(与弹出对话框一同使用)
类型：布尔值，默认为 False

### alert 警告

![](https://cdn.yuelili.com/20220113220929.png)
说明：警告，变红了。
类型：布尔值，默认为 False

### alignment 对齐方式

说明：对齐方式
类型：默认为 EXPAND

- EXPAND：填充
- LEFT：左对齐
- CENTER：中心对齐
- RIGHT：右对齐

### direction 不懂

说明：方向
类型：HORIZONTAL（默认） 或 VERTICAL。字符串

### emboss 按钮浮雕样式

说明：按钮的浮雕样式
类型列表：

- NORMAL：常规，绘制标准的按钮浮雕样式。
- NONE：无，只绘制文本和图标。
- PULLDOWN_MENU：下拉式菜单。
- RADIAL_MENU：径向菜单。
- NONE_OR_STATUS：无或状态，绘制无浮雕样式，除非按钮有着色状态，如动画状态。
  类型： 默认为 NORMAL，字符串

### enabled 布局激活

说明：为 False 时，此（子）布局未激活，显示为灰色
类型：布尔值，默认为 False

### operator_context 操作项上下文

说明：操作项上下文
类型列表：

- `INVOKE_DEFAULT` ：只执行 invoke()
- `INVOKE_REGION_WIN`：
- `INVOKE_REGION_CHANNELS`：
- `INVOKE_REGION_PREVIEW`：
- `INVOKE_AREA`：
- `INVOKE_SCREEN`：
- `EXEC_DEFAULT`：只执行 execute()
- `EXEC_REGION_WIN`：
- `EXEC_REGION_CHANNELS`：
- `EXEC_REGION_PREVIEW`：
- `EXEC_AREA`：
- `EXEC_SCREEN`：
  默认为 INVOKE_DEFAULT，字符串

### scale_x 布局水平缩放因子

说明：此（子）布局中项目的水平缩放因子
类型：浮点数，[0, inf]，默认为 0.0

### scale_y 布局垂直缩放因子

说明：此（子）布局中项目的垂直缩放因子
类型：浮点数，[0, inf]，默认为 0.0

### ui_units_x 布局水平固定尺寸

说明：此（子）布局中项目 X 方向的固定尺寸
类型：浮点数，[0, inf]，默认为 0.0

### ui_units_y 布局垂直固定尺寸

说明：此（子）布局中项目 Y 方向的固定尺寸
类型：浮点数，[0, inf]，默认为 0.0

### use_property_decorate 属性装饰

![](https://cdn.yuelili.com/20220113220806.png)
说明：那个点点
类型：布尔值，默认为 False

### use_property_split 不懂

类型：布尔值，默认为 False

### row() 行子布局

全名：row(align=False, heading='', heading_ctxt='', translate=True)
说明：子布局（Sub-layout），此子布局中的项目会在一排中并排排列
参数：

- align：布尔值, 可选。将按钮相互对齐。
- heading：字符串，可选，不为 None。将标题，标签插入到这个子布局的布局中。
- heading_ctxt：字符串可选，不为 None。覆盖指定标题的自动翻译环境。
- translate：布尔值, 可选。当 UI 翻译被启用时，翻译指定标题
  返回：放入项目的子布局
  返回类型：UILayout
  示例：
  ![](https://cdn.yuelili.com/20220117231107.png)

```python
import bpy
from bpy.types import Menu

class PT_ROW_TEST(bpy.types.Panel):
    bl_label = "ROW 示例"
    bl_idname = "PT_ROW_TEST"
    bl_space_type = 'PROPERTIES'
    bl_region_type = 'WINDOW'
    bl_context = "object"
    def draw(self, context):
        layout = self.layout
        row = layout.row()
        row.operator('mesh.primitive_cube_add', text="一行", icon='BLENDER')
        row.separator()
        row.operator('mesh.primitive_cube_add', text="一行")
        row.operator('mesh.primitive_cube_add', text="又一行")
# 注册
bpy.utils.register_class(PT_ROW_TEST)
```

### column() 列子布局

全名：column(align=False, heading='', heading_ctxt='', translate=True)
说明：子布局（Sub-layout），此子布局中的项目会在 **一列** 中并排排列
参数：

- align：布尔值, 可选。将按钮相互对齐。
- heading：字符串，可选，不为 None。将标题，标签插入到这个子布局的布局中。
- heading_ctxt：字符串可选，不为 None。覆盖指定标题的自动翻译环境。
- translate：布尔值, 可选。当 UI 翻译被启用时，翻译指定标题
  返回：放入项目的子布局
  返回类型：UILayout
  示例：
  ![](https://cdn.yuelili.com/20220117231353.png)

```python
import bpy
from bpy.types import Menu

class PT_LAYOUT_TEST(bpy.types.Panel):
    bl_label = "LAYOUT 示例"
    bl_idname = "PT_ROW_TEST"
    bl_space_type = 'PROPERTIES'
    bl_region_type = 'WINDOW'
    bl_context = "object"
    def draw(self, context):
        layout = self.layout
        lay = layout.column()
        lay.operator('mesh.primitive_cube_add', text="一行", icon='BLENDER')
        lay.separator()
        lay.operator('mesh.primitive_cube_add', text="一行")
        lay.operator('mesh.primitive_cube_add', text="又一行")
# 注册
bpy.utils.register_class(PT_LAYOUT_TEST)

```

### column_flow() 列浮动布局

全名：column_flow(columns=0, align=False)
说明：column_flow，列浮动排列？
参数：

- columns：整数， [0, inf]，列数，默认为 0
- align：对齐，布尔值，可选。按钮间相互对齐
  返回：放入项目的子布局
  返回类型：UILayout
  示例
  ![](https://cdn.yuelili.com/20220117231708.png)

```python
import bpy
from bpy.types import Menu

class PT_LAYOUT_TEST(bpy.types.Panel):
    bl_label = "LAYOUT 示例"
    bl_idname = "PT_ROW_TEST"
    bl_space_type = 'PROPERTIES'
    bl_region_type = 'WINDOW'
    bl_context = "object"
    def draw(self, context):
        layout = self.layout
        lay = layout.column_flow(columns=3)
        lay.operator('mesh.primitive_cube_add', text="第1个", icon='BLENDER')
        lay.separator()
        lay.operator('mesh.primitive_cube_add', text="第3个")
        lay.operator('mesh.primitive_cube_add', text="第4个")
        lay.operator('mesh.primitive_cube_add', text="第5个")
        lay.operator('mesh.primitive_cube_add', text="第6个")
# 注册
bpy.utils.register_class(PT_LAYOUT_TEST)
```

## grid_flow() 网格浮动布局

全名：grid_flow(row_major=False, columns=0, even_columns=False, even_rows=False, align=False)
说明：grid_flow
参数：

- row_major：布尔值，可选。Fill row by row, instead of column by column
- columns ：整数， [0, inf]，可选。列数，为正则是固定数字，自动是 0。为负责沿主轴自动倍数 (例如-2 的行布局只会产生 2,4,6 列，列布局会 2,4,6 行，等等)
- even_columns：布尔值，可选。所有列宽度一致
- even_rows：布尔值，可选。所有行高度一致
- align：布尔值，可选。按钮间相互对齐
  返回：放入项目的子布局
  返回类型：UILayout
  示例：
  ![](https://cdn.yuelili.com/20220117231757.png)

```python
import bpy
from bpy.types import Menu

class PT_LAYOUT_TEST(bpy.types.Panel):
    bl_label = "LAYOUT 示例"
    bl_idname = "PT_ROW_TEST"
    bl_space_type = 'PROPERTIES'
    bl_region_type = 'WINDOW'
    bl_context = "object"

    def draw(self, context):
        layout = self.layout
        lay = layout.grid_flow(columns=3)

        lay.operator('mesh.primitive_cube_add', text="第1个", icon='BLENDER')
        lay.separator()
        lay.operator('mesh.primitive_cube_add', text="第3个")
        lay.operator('mesh.primitive_cube_add', text="第4个")
        lay.operator('mesh.primitive_cube_add', text="第5个")
        lay.operator('mesh.primitive_cube_add', text="第6个")

# 注册
bpy.utils.register_class(PT_LAYOUT_TEST)
```

## box() 盒型布局

全名：box()
说明 ：子布局，(子布局中的项目彼此相邻，在盒型之中)
返回：放入项目的子布局
返回类型：UILayout
![](https://cdn.yuelili.com/20220117231914.png)

```python
import bpy
from bpy.types import Menu

class PT_LAYOUT_TEST(bpy.types.Panel):
    bl_label = "LAYOUT 示例"
    bl_idname = "PT_ROW_TEST"
    bl_space_type = 'PROPERTIES'
    bl_region_type = 'WINDOW'
    bl_context = "object"

    def draw(self, context):
        layout = self.layout
        lay = layout.box()
        lay.operator('mesh.primitive_cube_add', text="第1个", icon='BLENDER')
        lay.separator()
        lay.operator('mesh.primitive_cube_add', text="第3个")
        lay.operator('mesh.primitive_cube_add', text="第4个")
        lay.operator('mesh.primitive_cube_add', text="第5个")
        lay.operator('mesh.primitive_cube_add', text="第6个")
# 注册
bpy.utils.register_class(PT_LAYOUT_TEST)
```

## split() 分割布局

全名：split(factor=0.0, align=False)
说明：分割布局
参数：

- factor，浮点数，[0, 1]，可选。分割间距百分比 (不设置则自动计算)
- align：布尔值，可选。按钮间彼此对齐
  返回：放入项目的子布局
  返回类型：UILayout
  示例：
  ![](https://cdn.yuelili.com/20220117232054.png)

```python
import bpy
from bpy.types import Menu

class PT_LAYOUT_TEST(bpy.types.Panel):
    bl_label = "LAYOUT 示例"
    bl_idname = "PT_ROW_TEST"
    bl_space_type = 'PROPERTIES'
    bl_region_type = 'WINDOW'
    bl_context = "object"
    def draw(self, context):
        layout = self.layout
        lay = layout.split()
        lay.operator('mesh.primitive_cube_add', text="第1个", icon='BLENDER')
        lay.separator()
        lay.operator('mesh.primitive_cube_add', text="第3个")
        lay = layout.split()
        lay.operator('mesh.primitive_cube_add', text="第4个")
        lay.operator('mesh.primitive_cube_add', text="第5个")
        lay = layout.split()
        lay.operator('mesh.primitive_cube_add', text="第6个")
# 注册
bpy.utils.register_class(PT_LAYOUT_TEST)
```

## menu_pie() 饼菜单

全名：menu_pie()
说明：子布局，其中的项目会被饼状包裹（饼菜单）
返回：Sub-layout to put items in
返回类型：UILayout
示例：
![](https://cdn.yuelili.com/20220117225822.png)

```python
import bpy
from bpy.types import Menu
class MT_pie_test(Menu):
    bl_label = 'PIE TEST'
    bl_idname = 'MT_pie_test'
    def draw(self, context):
        pie = self.layout.menu_pie()
        # 顺序为：西、东、南、北、西北、东北、西南、东南
        # 使用separator跳过顺序
        pie.separator()  # 跳过西
        pie.operator('mesh.primitive_cube_add', text="东", icon='EVENT_A')
        pie.operator('mesh.primitive_cube_add', text="南", icon='EVENT_A')
        pie.operator('mesh.primitive_cube_add', text="北")
# 注册与使用
bpy.utils.register_class(MT_pie_test)
bpy.ops.wm.call_menu_pie(name='MT_pie_test')
```

## classmethod icon() 不懂

全名：classmethodicon(data)
说明：返回该数据的自定义图标, 并使用。比如获取材质或纹理的图标
参数：data，任何类型, 不会为 None。从要使用的图标获取的数据。
返回：Icon identifier
返回类型：整数，[0, inf]

## classmethod enum_item_name() 不懂

全名：classmethodenum_item_name(data, property, identifier)
说明：返回该枚举项目的用户界面名称
参数：

- data：任何类型, 不会为 None。从中获取属性的数据
- property：字符串， 不会为 None。数据中属性的标识符
- identifier：字符串， 不会为 None。枚举项目的标识符
  返回：枚举项目的名称
  返回类型：字符串， 不会为 None

## classmethod enum_item_description() 不懂

全名：classmethodenum_item_description(data, property, identifier)
说明：返回该枚举项目的 UI 介绍
参数：

- data：任何类型, 不会为 None。从中获取属性的数据
- property：字符串， 不会为 None。数据中属性的标识符
- identifier：字符串， 不会为 None。枚举项目的标识符
  返回：枚举项目的 UI 描述
  返回类型：字符串， 不会为 None

## classmethod enum_item_icon() 不懂

全名：classmethodenum_item_icon(data, property, identifier)
Return the icon for this enum item
参数：

- data：任何类型, 不会为 None。从中获取属性的数据
- property：字符串， 不会为 None。数据中属性的标识符
- identifier：字符串， 不会为 None。枚举项目的标识符
  返回：Icon identifier
  返回类型：int in [0, inf]

## prop()

全名：prop(data, property, text='', text_ctxt='', translate=True, icon='NONE', expand=False, slider=False, toggle=- 1, icon_only=False, event=False, full_event=False, emboss=True, index=- 1, icon_value=0, invert_checkbox=False)
说明：Item. Exposes，RNA 项目，并置入布局中
参数：

- data：任何类型, 不会为 None。从中获取属性的数据
- property：字符串， 不会为 None。属性标识符
- text：string ，可选。覆盖项目的自动文本。
- text_ctxt：string ，可选。覆盖给定文本的自动翻译环境。
- translate，字符串，可选。当启用 UI 翻译时，翻译给定文本。
- icon，图标字符串
- expand：布尔值，可选。扩展按钮以显示更多细节。
- slider ：布尔值，可选。数字类型的值使用滑块小工具
- toggle：切换-1,1,0。对布尔值使用切换部件，或在禁用时使用复选框（默认为-1，仅在显示图标时使用切换功能）
- icon_only：布尔值，可选。仅绘制图标，不用文字
- event：布尔值，可选。使用按钮来输入 key 事件
- full_event：布尔值，可选。使用按钮输入完整的事件，包括修改器
- emboss：布尔值，可选。绘制按钮自身，而不仅仅是 图标/文字。为 False 则对应"NONE_OR_STATUS"浮雕样式
- index：整数， [-2, inf]，可选。该按钮的索引，当设置为-1 时，可以访问数组中的单个成员，当设置为-1 时，则使用所有数组成员
- icon_value：整数，[0, inf] ，可选。图标值， 覆盖项目的自动图标
- invert_checkbox：布尔值，可选。绘制反转的复选框
  示例：
  ![](https://cdn.yuelili.com/20220119230542.png)

```python
import bpy
class PanelOne(bpy.types.Panel):
    bl_idname = "VIEW3D_PT_test_1"
    bl_label = "面板 One"
    bl_space_type = 'VIEW_3D'
    bl_region_type = 'UI'
    bl_category = "我是面板"

    def draw(self, context):
        layout = self.layout
        col = layout.column()
        # 一般情况下 RNA 属性放在场景(scene)中
        scene = context.scene
        col.prop(scene,"myStr")
bpy.utils.register_class(PanelOne)
# 需要对 RNA 属性注册，一般情况下 RNA 属性放在场景(scene)中
bpy.types.Scene.myStr = bpy.props.StringProperty(name = '输入：')
```

## props_enum() 不懂

全名：props_enum(data, property)
说明：props_enum
参数：

- data：任何类型, 不会为 None。从中获取属性的数据
- property：字符串， 不会为 None。数据中属性的标识符

## prop_menu_enum() 不懂

全名：prop_menu_enum(data, property, text='', text_ctxt='', translate=True, icon='NONE')
说明：prop_menu_enum
参数：

- data：任何类型, 不会为 None。从中获取属性的数据
- property：字符串， 不会为 None。数据中属性的标识符
- text：string ，可选。覆盖项目的自动文本。
- text_ctxt：string ，可选。覆盖给定文本的自动翻译环境。
- translate，字符串，可选。当启用 UI 翻译时，翻译给定文本。
- icon，图标字符串

## prop_with_popover() 不懂

全名：prop_with_popover(data, property, text="", text_ctxt="", translate=True, icon='NONE', icon_only=False, panel)
说明：prop_with_popover
参数：

- data：任何类型, 不会为 None。从中获取属性的数据
- property：字符串， 不会为 None。数据中属性的标识符
- text：string ，可选。覆盖项目的自动文本。
- text_ctxt：string ，可选。覆盖给定文本的自动翻译环境。
- translate，字符串，可选。当启用 UI 翻译时，翻译给定文本。
- icon，图标字符串

## prop_with_menu() 不懂

全名：prop_with_menu(data, property, text="", text_ctxt="", translate=True, icon='NONE', icon_only=False, menu)
说明：prop_with_menu
参数：

- data：任何类型, 不会为 None。从中获取属性的数据
- property：字符串， 不会为 None。数据中属性的标识符
- text：string ，可选。覆盖项目的自动文本。
- text_ctxt：string ，可选。覆盖给定文本的自动翻译环境。
- translate，字符串，可选。当启用 UI 翻译时，翻译给定文本。
- icon，图标字符串

## prop_tabs_enum() 不懂

全名：prop_tabs_enum(data, property, data_highlight=None, property_highlight='', icon_only=False)
说明：prop_tabs_enum
参数：

- data：任何类型, 不会为 None。从中获取属性的数据
- property：字符串， 不会为 None。数据中属性的标识符
- data_highlight：任何类型，可选, 不为 None。高亮属性的数据
- property_highlight：字符串，可选, 不为 None。数据中高亮属性的 ID
- icon_only：布尔值，可选。仅绘制图标，不用文字

## prop_enum() 不懂

全名：prop_enum(data, property, value, text='', text_ctxt='', translate=True, icon='NONE')
说明：prop_enum
参数：

- data：任何类型, 不会为 None。从中获取属性的数据
- property：字符串， 不会为 None。数据中属性的标识符
- value：字符串，不会为 None。枚举属性值
- text：string ，可选。覆盖项目的自动文本。
- text_ctxt：string ，可选。覆盖给定文本的自动翻译环境。
- translate，字符串，可选。当启用 UI 翻译时，翻译给定文本。
- icon，图标字符串

## prop_search() 不懂

全名：prop_search(data, property, search_data, search_property, text='', text_ctxt='', translate=True, icon='NONE')
说明：prop_search
参数：

- data：任何类型, 不会为 None。从中获取属性的数据
- property：字符串， 不会为 None。数据中属性的标识符
- search_data：任何类型，不会为 None。用于搜索的集合数据
- search_property：字符串，不会为 None。搜索的集合属性的 ID
- text：string ，可选。覆盖项目的自动文本。
- text_ctxt：string ，可选。覆盖给定文本的自动翻译环境。
- translate，字符串，可选。当启用 UI 翻译时，翻译给定文本。
- icon，图标字符串

## prop_decorator() 不懂

全名：prop_decorator(data, property, index=- 1)
说明：prop_decorator
参数：

- data：任何类型, 不会为 None。从中获取属性的数据
- property：字符串， 不会为 None。数据中属性的标识符
- index：整数，[-2, inf] ，可选。该按钮的索引，设置为-1 时，可以访问数组中的单个成员，为-1 时，则是全体数组成员

## operator()

全名：operator(operator, text='', text_ctxt='', translate=True, icon='NONE', emboss=True, depress=False, icon_value=0)
说明：项目。在布局中放置一个按钮，用于调用操作项。
参数：

- operator：字符串，不会为 None。操作项 ID
- text：string ，可选。覆盖项目的自动文本。
- text_ctxt：string ，可选。覆盖给定文本的自动翻译环境。
- translate，字符串，可选。当启用 UI 翻译时，翻译给定文本。
- icon，图标字符串
- emboss：布尔值，可选。绘制按钮自身，而不仅仅是 图标/文字
- depress：布尔值，可选。绘制“按下”样式
- icon_value：整数，[0, inf] ，可选。图标值， 覆盖项目的自动图标
  返回：Operator properties to fill in
  返回类型：OperatorProperties
  示例：物体属性栏，插入一个“添加立方体”
  ![](https://cdn.yuelili.com/20220118023306.png)

```python
import bpy
from bpy.types import Panel
class DATA_PT_test(Panel):
    bl_label = "我是测试1"
    bl_space_type = 'PROPERTIES'
    bl_region_type = 'WINDOW'
    bl_context = "object"

    def draw(self, context):
        layout = self.layout
        layout.operator("mesh.primitive_cube_add")

bpy.utils.register_class(DATA_PT_test)
```

## operator_menu_hold() 不懂

全名：operator_menu_hold(operator, text="", text_ctxt="", translate=True, icon='NONE', emboss=True, depress=False, icon_value=0, menu)
说明：项目。在布局中放置一个按钮，用于调用操作项。
参数：

- operator：字符串，不会为 None。操作项 ID
- text：string ，可选。覆盖项目的自动文本。
- text_ctxt：string ，可选。覆盖给定文本的自动翻译环境。
- translate，字符串，可选。当启用 UI 翻译时，翻译给定文本。
- icon，图标字符串
- emboss：布尔值，可选。绘制按钮自身，而不仅仅是 图标/文字
- depress：布尔值，可选。绘制“按下”样式
- icon_value：整数，[0, inf] ，可选。图标值， 覆盖项目的自动图标
- menu：字符串，不会为 None。菜单的 ID
  返回：填入的操作项属性
  返回类型：OperatorProperties
  示例：物体属性栏，插入一个“添加立方体”。
  注意：我不知道这个函数跟 operator()的区别。加了个 menu 参数能跑了反正
  ![](https://cdn.yuelili.com/20220118023306.png)

```python
import bpy
from bpy.types import Panel
class DATA_PT_test(Panel):
    bl_label = "我是测试1"
    bl_space_type = 'PROPERTIES'
    bl_region_type = 'WINDOW'
    bl_context = "object"

    def draw(self, context):
        layout = self.layout
        layout.operator_menu_hold("mesh.primitive_cube_add",menu="object.metaball_add")

bpy.utils.register_class(DATA_PT_test)

```

## operator_enum() 操作项枚举

![](https://cdn.yuelili.com/20220113224447.png)
全名：operator_enum(operator, property, icon_only=False)
说明：operator_enum
参数：

- operator：字符串，不会为 None。操作项 ID
- property：字符串， 不会为 None。数据中属性的标识符
- icon_only：布尔值，可选。仅绘制图标，不用文字
  示例：代码见下个函数 operator_menu_enum

## operator_menu_enum() 操作项菜单枚举

全名：operator_menu_enum(operator, property, text='', text_ctxt='', translate=True, icon='NONE')
说明：operator_menu_enum
参数：

- operator：字符串，不会为 None。操作项 ID
- property：字符串， 不会为 None。数据中属性的标识符
- text：string ，可选。覆盖项目的自动文本。
- text_ctxt：string ，可选。覆盖给定文本的自动翻译环境。
- translate，字符串，可选。当启用 UI 翻译时，翻译给定文本。
- icon，图标字符串
  返回：填入的操作项属性
  返回类型：OperatorProperties
  示例：把添加融球操作项菜单，放在对象属性里
  ![](https://cdn.yuelili.com/20220118021752.png)

```python
import bpy
from bpy.types import Panel

class DATA_PT_test(Panel):
    bl_label = "我是测试"
    bl_space_type = 'PROPERTIES'
    bl_region_type = 'WINDOW'
    bl_context = "object"
    # bl_options = {'HIDE_HEADER'}
    def draw(self, context):
        layout = self.layout
        layout.operator_menu_enum("object.metaball_add", "type")

bpy.utils.register_class(DATA_PT_test)
```

## label()

全名：label(text='', text_ctxt='', translate=True, icon='NONE', icon_value=0)
说明：项目。布局中显示文字或图标
参数：

- text：string ，可选。覆盖项目的自动文本。
- text_ctxt：string ，可选。覆盖给定文本的自动翻译环境。
- translate，字符串，可选。当启用 UI 翻译时，翻译给定文本。
- icon：[系统图标](https://www.yuelili.com/blender-development-use-blenders-own-icon/)列表里任选一个
- icon_value：整数，[0, inf] ，可选。图标值， 覆盖项目的自动图标
  示例：放 3 个标签
  ![](https://cdn.yuelili.com/20220118024744.png)

```python
import bpy
from bpy.types import Menu


class PT_LAYOUT_TEST(bpy.types.Panel):
    bl_label = "LAYOUT 示例"
    bl_idname = "PT_ROW_TEST"
    bl_space_type = 'PROPERTIES'
    bl_region_type = 'WINDOW'
    bl_context = "object"
    direction = "VERTICAL"

    def draw(self, context):
        layout = self.layout
        lay = layout.column()
        lay.label(text="A")
        lay.label(text="B")
        lay.label(text="C")

# 注册
bpy.utils.register_class(PT_LAYOUT_TEST)


```

## menu() 插入整个菜单

全名：menu(menu, text='', text_ctxt='', translate=True, icon='NONE', icon_value=0)
说明：将目标整个 menu 作为子面板加入
参数：

- menu：字符串，不会为 None。菜单的 ID
- text：string ，可选。覆盖项目的自动文本。
- text_ctxt：string ，可选。覆盖给定文本的自动翻译环境。
- translate，字符串，可选。当启用 UI 翻译时，翻译给定文本。
- icon，图标字符串
- icon_value：整数，[0, inf] ，可选。图标值， 覆盖项目的自动图标
  示例：加入整个 Mesh 菜单放到物体属性栏
  ![](https://cdn.yuelili.com/20220118022617.png)

```python
import bpy
from bpy.types import Panel

class DATA_PT_test(Panel):
    bl_label = "我是测试1"
    bl_space_type = 'PROPERTIES'
    bl_region_type = 'WINDOW'
    bl_context = "object"

    def draw(self, context):
        layout = self.layout
        layout.menu("VIEW3D_MT_mesh_add")

bpy.utils.register_class(DATA_PT_test)
```

## menu_contents() 插入全部菜单

全名：menu_contents(menu)
说明：把菜单所有子项全部列成一列
参数：menu ，菜单 ID，字符串， 不会为 None。
示例：把生成网格图形菜单，打包放物体属性栏
![](https://cdn.yuelili.com/20220118022543.png)

```python
import bpy
from bpy.types import Panel

class DATA_PT_test(Panel):
    bl_label = "我是测试1"
    bl_space_type = 'PROPERTIES'
    bl_region_type = 'WINDOW'
    bl_context = "object"
    # bl_options = {'HIDE_HEADER'}

    def draw(self, context):
        layout = self.layout
        layout.menu_contents("VIEW3D_MT_mesh_add")

bpy.utils.register_class(DATA_PT_test)

```

## popover() 不懂

全名：popover(panel, text='', text_ctxt='', translate=True, icon='NONE', icon_value=0)
说明：popover
参数：

- panel：字符串，不会为 None。面板的 ID
- text：string ，可选。覆盖项目的自动文本。
- text_ctxt：string ，可选。覆盖给定文本的自动翻译环境。
- translate，字符串，可选。当启用 UI 翻译时，翻译给定文本。
- icon，图标字符串
- icon_value：整数，[0, inf] ，可选。图标值， 覆盖项目的自动图标

## popover_group() 不懂

全名：popover_group(space_type, region_type, context, category)
popover_group
参数：

- space_type：VIEW_3D/IMAGE_EDITOR 等空间类型之一
- region_type：WINDOW、HEADER 等区块类型之一
- context：上下文，字符串， 不会为 None。面板的上下文
- category：分类，字符串，不会为 None：面板分类

## separator() 分割线

![](https://cdn.yuelili.com/20220117014033.png)
全名：separator(factor=1.0)
说明：项目，在项目间插入分割线
参数：factor (float in [0, inf]，可选) 。Percentage, Percentage of width to space (leave unset for default space)

## separator_spacer() 不懂

全名：separator_spacer()
说明：Item. Inserts horizontal spacing empty space into the layout between items

## context_pointer_set() 不懂

全名：context_pointer_set(name, data)
说明：context_pointer_set
参数：

- name：字符串，不会为 None。上下文中 entry 的名称
- data：任何类型, 不会为 None。从中获取属性的数据

## template_header() 不懂

全名：template_header()
Inserts common Space header UI (editor type selector)

## template_ID() 不懂

全名：template_ID(data, property, new='', open='', unlink='', filter='ALL', live_icon=False, text='', text_ctxt='', translate=True)
说明：template_ID
参数：

- data：任何类型, 不会为 None。从中获取属性的数据
- property：字符串， 不会为 None。数据中属性的标识符
- new：字符串，可选, 不为 None。Operator identifier to create a new ID block
- open：字符串，可选, 不为 None。Operator identifier to open a file for creating a new ID block
- unlink：字符串，可选, 不为 None。Operator identifier to unlink the ID block
- filter ：字符串，ALL 或 AVAILABLE，可选。限制项目是否能被选择的选项
- live_icon：布尔值，可选。Show preview instead of fixed icon
- text：string ，可选。覆盖项目的自动文本。
- text_ctxt：string ，可选。覆盖给定文本的自动翻译环境。
- translate，字符串，可选。当启用 UI 翻译时，翻译给定文本。

## template_ID_preview() 不懂

全名：template_ID_preview(data, property, new='', open='', unlink='', rows=0, cols=0, filter='ALL', hide_buttons=False)
说明：template_ID_preview
参数：

- data：任何类型, 不会为 None。从中获取属性的数据
- property：字符串， 不会为 None。数据中属性的标识符
- new：字符串，可选, 不为 None。Operator identifier to create a new ID block
- open：字符串，可选, 不为 None。Operator identifier to open a file for creating a new ID block
- unlink：字符串，可选, 不为 None。Operator identifier to unlink the ID block
- rows：整数，[0, inf] ，可选。Number of thumbnail preview rows to display
- cols：整数，[0, inf] ，可选。Number of thumbnail preview columns to display
- filter ：字符串，ALL 或 AVAILABLE，可选。限制项目是否能被选择的选项
- hide_buttons：布尔，可选。仅显示列表，而不显示按钮

## template_any_ID() 不懂

全名：template_any_ID(data, property, type_property, text='', text_ctxt='', translate=True)
说明：template_any_ID
参数：

- data：任何类型, 不会为 None。从中获取属性的数据
- property：字符串， 不会为 None。数据中属性的标识符
- type_property：字符串， 不会为 None。Identifier of property in data giving the type of the ID-blocks to use
- text：string ，可选。覆盖项目的自动文本。
- text_ctxt：string ，可选。覆盖给定文本的自动翻译环境。
- translate，字符串，可选。当启用 UI 翻译时，翻译给定文本。

## template_ID_tabs() 不懂

全名：template_ID_tabs(data, property, new='', menu='', filter='ALL')
说明：template_ID_tabs
参数：

- data：任何类型, 不会为 None。从中获取属性的数据
- property：字符串， 不会为 None。数据中属性的标识符
- new：字符串，可选, 不为 None。Operator identifier to create a new ID block
- menu：字符串，不会为 None。菜单的 ID
- filter ：字符串，ALL 或 AVAILABLE，可选。限制项目是否能被选择的选项

## template_search() 不懂

全名：template_search(data, property, search_data, search_property, new='', unlink='')
说明：template_search
参数：

- data：任何类型, 不会为 None。从中获取属性的数据
- property：字符串， 不会为 None。数据中属性的标识符
- search_data：任何类型，不会为 None。用于搜索的集合数据
- search_property：字符串，可选, 不为 None。搜索集合属性的 ID
- new：字符串，可选, 不为 None。Operator identifier to create a new ID block
- unlink：字符串，可选, 不为 None。Operator identifier to unlink the ID block

## template_search_preview() 不懂

全名：template_search_preview(data, property, search_data, search_property, new='', unlink='', rows=0, cols=0)
说明：template_search_preview
参数：

- data：任何类型, 不会为 None。从中获取属性的数据
- property：字符串， 不会为 None。数据中属性的标识符
- search_data：任何类型，不会为 None。用于搜索的集合数据
- search_property：字符串，可选, 不为 None。搜索集合属性的 ID
- new：字符串，可选, 不为 None。Operator identifier to create a new ID block
- unlink：字符串，可选, 不为 None。Operator identifier to unlink the ID block
- rows：整数，[0, inf] ，可选。Number of thumbnail preview rows to display
- cols：整数，[0, inf] ，可选。Number of thumbnail preview columns to display

## template_path_builder() 不懂

全名：template_path_builder(data, property, root, text='', text_ctxt='', translate=True)
说明：template_path_builder
参数：

- data：任何类型, 不会为 None。从中获取属性的数据
- property：字符串， 不会为 None。数据中属性的标识符
- root (ID) ：ID-block from which path is evaluated from
- text：string ，可选。覆盖项目的自动文本。
- text_ctxt：string ，可选。覆盖给定文本的自动翻译环境。
- translate，字符串，可选。当启用 UI 翻译时，翻译给定文本。

## template_modifiers()

全名：template_modifiers()
说明：为修改器栈生成 UI 布局
示例：用 operator_menu_enum 生成枚举菜单，再用本函数生成布局
![](https://cdn.yuelili.com/20220118021321.png)

```python
import bpy
from bpy.types import Panel
class DATA_PT_modifiers(Panel):
    bl_label = "谁TM才是修改器"
    bl_space_type = 'PROPERTIES'
    bl_region_type = 'WINDOW'
    bl_context = "modifier"
    # bl_options = {'HIDE_HEADER'} # 如果不想要Header，就把注释取消

    def draw(self, context):
        layout = self.layout
        layout.operator_menu_enum("object.modifier_add", "type")
        layout.template_modifiers()
bpy.utils.register_class(DATA_PT_modifiers)
```

## template_constraints() 不懂

全名：template_constraints(use_bone_constraints=True)
说明：Generates the panels for the constraint stack
参数：use_bone_constraints 布尔值，可选) 。为骨骼约束添加面板替代对象约束

## template_grease_pencil_modifiers() 不懂

全名：template_grease_pencil_modifiers()
说明：Generates the panels for the grease pencil modifier stack

## template_shaderfx() 不懂

全名：template_shaderfx()
说明：Generates the panels for the shader effect stack

## template_greasepencil_color() 不懂

全名：template_greasepencil_color(data, property, rows=0, cols=0, scale=1.0, filter='ALL')
说明：template_greasepencil_color
参数：

- data：任何类型, 不会为 None。从中获取属性的数据
- property：字符串， 不会为 None。数据中属性的标识符
- rows：整数，[0, inf] ，可选。Number of thumbnail preview rows to display
- cols：整数，[0, inf] ，可选。Number of thumbnail preview columns to display
- scale：浮点值，[0.1, 1.5]，可选。Scale of the image thumbnails
- filter ：字符串，ALL 或 AVAILABLE，可选。限制项目是否能被选择的选项

## template_constraint_header() 不懂

全名：template_constraint_header(data)
说明：生成约束面板的 header
参数：
data：任何类型, 不会为 None。从中获取属性的数据

## template_preview() 不懂

全名：template_preview(id, show_buttons=True, parent=None, slot=None, preview_id='')
说明：项目。材料、纹理、灯光或世界的预览窗口。
参数：

- id (ID) 。ID data-block
- show_buttons 布尔值，可选) 。Show preview buttons?
- parent (ID，可选) 。ID data-block
- slot (TextureSlot，可选) 。Texture slot
- preview_id ：字符串， (，可选, 不为 None)) 。Identifier of this preview widget, if not set the ID type will be used (i.e. all previews of materials without explicit ID will have the same size…)

## template_curve_mapping() 不懂

全名：template_curve_mapping(data, property, type='NONE', levels=False, brush=False, use_negative_slope=False, show_tone=False)
说明：项目。一个曲线映射部件，例如用于灯光的衰减曲线。
参数：

- data：任何类型, 不会为 None。从中获取属性的数据
- property：字符串， 不会为 None。数据中属性的标识符
- type ：要显示的曲线类型，可选。NONE、VECTOR、COLOR、HUE
- levels 布尔值，可选) 。Show black/white levels
- brush 布尔值，可选) 。Show brush options
- use_negative_slope 布尔值，可选) 。Use a negative slope by default
- show_tone 布尔值，可选) 。Show tone options

## template_curveprofile() 不懂

全名：template_curveprofile(data, property)
说明：一个用于自定义配置文件的路径编辑器
参数：

- data：任何类型, 不会为 None。从中获取属性的数据
- property：字符串， 不会为 None。数据中属性的标识符

## template_color_ramp() 不懂

全名：template_color_ramp(data, property, expand=False)
说明：项目。一个颜色渐变小部件
参数：

- data：任何类型, 不会为 None。从中获取属性的数据
- property：字符串， 不会为 None。数据中属性的标识符
- expand：布尔值，可选。扩展按钮以显示更多细节。

## template_icon() 不懂

全名：template_icon(icon_value, scale=1.0)
说明：显示一个大图标
参数：

- icon_value：整数，[0, inf] ，可选。图标值， 覆盖项目的自动图标
- scale：浮点值，[0.1, 1.5]，可选。图像缩略图的比例

## template_icon_view() 不懂

全名：template_icon_view(data, property, show_labels=False, scale=6.0, scale_popup=5.0)
说明：枚举。显示图标预览的大组件
参数：

- data：任何类型, 不会为 None。从中获取属性的数据
- property：字符串， 不会为 None。数据中属性的标识符
- show_labels 布尔值，可选) 。Show enum label in preview buttons
- scale (float in [1, 100]，可选) 。UI Units, Scale the button icon size (by the button size)
- scale_popup (float in [1, 100]，可选) 。Scale, Scale the popup icon size (by the button size)

## template_histogram() 不懂

全名：template_histogram(data, property)
说明：Item. A histogramm widget to analyze imaga data
参数：
data ：AnyType， 不会为 None。Data from which to take property
property ：字符串， 不会为 None。Identifier of property in data

## template_waveform() 不懂

全名：template_waveform(data, property)
说明：Item. A waveform widget to analyze imaga data
参数：

- data：任何类型, 不会为 None。从中获取属性的数据
- property：字符串， 不会为 None。数据中属性的标识符

## template_vectorscope() 不懂

全名：template_vectorscope(data, property)
说明：Item. A vectorscope widget to analyze imaga data
参数：

- data：任何类型, 不会为 None。从中获取属性的数据
- property：字符串， 不会为 None。数据中属性的标识符

## template_layers() 不懂

全名：template_layers(data, property, used_layers_data, used_layers_property, active_layer)
说明：template_layers
参数：

- data：任何类型, 不会为 None。从中获取属性的数据
- property：字符串， 不会为 None。数据中属性的标识符
- used_layers_data (AnyType) 。Data from which to take property
- used_layers_property ：字符串， 不会为 None。Identifier of property in data
- active_layer 整数， [0, inf]) 。Active Layer

## template_color_picker() 不懂

全名：template_color_picker(data, property, value_slider=False, lock=False, lock_luminosity=False, cubic=False)
说明：项目。一个颜色选择色轮小部件
参数：

- data：任何类型, 不会为 None。从中获取属性的数据
- property：字符串， 不会为 None。数据中属性的标识符
- value_slider 布尔值，可选) 。Display the value slider to the right of the color wheel
- lock 布尔值，可选) 。Lock the color wheel display to value 1.0 regardless of actual color
- lock_luminosity 布尔值，可选) 。Keep the color at its original vector length
- cubic 布尔值，可选) 。Cubic saturation for picking values close to white

## template_palette() 不懂

全名：template_palette(data, property, color=False)
说明：Item. A palette used to pick colors
参数：

- data：任何类型, 不会为 None。从中获取属性的数据
- property：字符串， 不会为 None。数据中属性的标识符
- color 布尔值，可选) 。Display the colors as colors or values

## template_image_layers() 不懂

全名：template_image_layers(image, image_user)
说明：template_image_layers

## template_image() 不懂

全名：template_image(data, property, image_user, compact=False, multiview=False)
说明：Item(s). User interface for selecting images and their source paths
参数：

- data：任何类型, 不会为 None。从中获取属性的数据
- property：字符串， 不会为 None。数据中属性的标识符
- compact 布尔值，可选) 。Use more compact layout
- multiview 布尔值，可选) 。Expose Multi-View options

## template_image_settings() 不懂

全名：template_image_settings(image_settings, color_management=False)
说明：用于设置图像格式选项的用户界面
参数：color_management，布尔值，可选 。显示色彩管理设置

## template_image_stereo_3d() 不懂

全名：template_image_stereo_3d(stereo_3d_format)
说明：用于设置图像立体三维选项的用户界面

## template_image_views() 不懂

全名：template_image_views(image_settings)
说明：用于设置图像视图输出选项的用户界面

## template_movieclip() 不懂

全名：template_movieclip(data, property, compact=False)
说明：项目(s). 选择电影片段及其来源路径的用户界面
参数：

- data：任何类型, 不会为 None。从中获取属性的数据
- property：字符串， 不会为 None。数据中属性的标识符
- compact 布尔值，可选) 。Use more compact layout

## template_track() 不懂

全名：template_track(data, property)
说明：项目。一个电影轨道的小部件，用于预览跟踪图像。
参数：

- data：任何类型, 不会为 None。从中获取属性的数据
- property：字符串， 不会为 None。数据中属性的标识符

## template_marker() 不懂

全名：template_marker(data, property, clip_user, track, compact=False)
说明：Item. A widget to control single marker settings.
参数：

- data：任何类型, 不会为 None。从中获取属性的数据
- property：字符串， 不会为 None。数据中属性的标识符
- compact 布尔值，可选) 。Use more compact layout

## template_movieclip_information() 不懂

全名：template_movieclip_information(data, property, clip_user)
说明：项目。电影剪辑信息数据。
参数：

- data：任何类型, 不会为 None。从中获取属性的数据
- property：字符串， 不会为 None。数据中属性的标识符

## template_list() 不懂

全名：template_list(listtype_name, list_id, dataptr, propname, active_dataptr, active_propname, item_dyntip_propname='', rows=5, maxrows=5, type='DEFAULT', columns=9, sort_reverse=False, sort_lock=False)
说明：项目。一个显示数据的列表部件，例如顶点组。
参数：

- listtype_name ：字符串， 不会为 None。要使用的列表类型的标识符
- list_id ：字符串， 不会为 None。列表部件的标识符。如果不设置则使用类名
- dataptr (AnyType) 。从集合属性中获取的数据。
- propname ：字符串， 不会为 None。数据中集合属性的标识符
- active_dataptr ：AnyType, 不会为 None。从该数据中获取整数属性，活动项目的索引
- active_propname ：字符串， 不会为 None。active_data 中整数属性的标识符，活动项目的索引。
- item_dyntip_propname ：字符串，可选, 不为 None。项目中字符串属性的标识，用作工具提示内容
- rows ：整数， [0, inf]，可选。默认和最小显示行数。
- maxrows ：整数， [0, inf]，可选。默认显示的最大行数。
- type：布局方式 \* DEFAULT：默认布局。使用默认的多行布局。

      * COMPACT：紧凑型布局。使用紧凑型的单行布局。

      * GRID：网格布局。使用基于网格的布局。

- columns ：整数， [0, inf]，可选。每行显示的项目数，用于 GRID 布局。
- sort_reverse ：boolean，可选。默认反向顺序显示项目。
- sort_lock ：boolean，可选 。将显示顺序锁定为默认值。

## template_running_jobs() 不懂

全名：template_running_jobs()
说明：template_running_jobs

## template_operator_search() 不懂

全名：template_operator_search()
说明：template_operator_search

## template_menu_search() 不懂

全名：template_menu_search()
说明：template_menu_search

## template_header_3D_mode() 不懂

全名：template_header_3D_mode()

## template_edit_mode_selection() 不懂

全名：template_edit_mode_selection()
说明：Inserts common 3DView Edit modes header UI (selector for selection mode)

## template_reports_banner() 不懂

全名：template_reports_banner()
说明：template_reports_banner

## template_input_status() 不懂

全名：template_input_status()
说明：template_input_status

## template_node_link() 不懂

全名：template_node_link(ntree, node, socket)
说明：template_node_link

## template_node_view() 不懂

全名：template_node_view(ntree, node, socket)
说明：template_node_view

## template_texture_user() 不懂

全名：template_texture_user()
说明：template_texture_user

## template_keymap_item_properties() 不懂

全名：template_keymap_item_properties(item)
说明：template_keymap_item_properties

## template_component_menu() 不懂

全名：template_component_menu(data, property, name='')
说明：项目。在一个弹出式菜单中显示扩展的属性
参数：

- data：任何类型, 不会为 None。从中获取属性的数据
- property：字符串， 不会为 None。数据中属性的标识符

## template_colorspace_settings() 不懂

全名：template_colorspace_settings(data, property)
说明：项目。一个控制输入颜色空间设置的小组件。
参数：

- data：任何类型, 不会为 None。从中获取属性的数据
- property：字符串， 不会为 None。数据中属性的标识符

## template_colormanaged_view_settings() 不懂

全名：template_colormanaged_view_settings(data, property)
说明：项目。一个小部件，控制颜色管理的视图设置设置。
参数：

- data：任何类型, 不会为 None。从中获取属性的数据
- property：字符串， 不会为 None。数据中属性的标识符

## template_node_socket() 不懂

全名：template_node_socket(color=(0.0, 0.0, 0.0, 1.0))
说明：Node Socket Icon
参数：color (float array of 4 items in [0, 1]，可选) 。Color

## template_cache_file() 不懂

全名：template_cache_file(data, property)
说明：Item(s)。选择缓存文件和其源路径的用户界面
参数：

- data：任何类型, 不会为 None。从中获取属性的数据
- property：字符串， 不会为 None。数据中属性的标识符

## template_cache_file_velocity() 不懂

全名：template_cache_file_velocity(data, property)
说明：显示缓存文件的速度属性
参数：

- data：任何类型, 不会为 None。从中获取属性的数据
- property：字符串， 不会为 None。数据中属性的标识符

## template_cache_file_procedural() 不懂

全名：template_cache_file_procedural(data, property)
说明：显示缓存文件渲染程序属性
参数：

- data：任何类型, 不会为 None。从中获取属性的数据
- property：字符串， 不会为 None。数据中属性的标识符

## template_cache_file_time_settings() 不懂

全名：template_cache_file_time_settings(data, property)
说明：显示缓存文件时间设置
参数：

- data：任何类型, 不会为 None。从中获取属性的数据
- property：字符串， 不会为 None。数据中属性的标识符

## template_recent_files() 不懂

全名：template_recent_files(rows=5)
说明：显示最近保存的.blend 文件的列表
参数：rows，整数，[1, inf] ，可选。要显示的最大项目数
返回：Number of items drawn
返回类型：int in [0, inf]

## template_file_select_path() 不懂

全名：template_file_select_path(params)
说明：项目。一个文本按钮，用于设置活动的文件浏览器路径。

## template_event_from_keymap_item() 不懂

全名：template_event_from_keymap_item(item, text='', text_ctxt='', translate=True)
说明：Display keymap item as icons/text
参数：

- item (KeyMapItem, 不会为 None。Item
- text：string ，可选。覆盖项目的自动文本。
- text_ctxt：string ，可选。覆盖给定文本的自动翻译环境。
- translate，字符串，可选。当启用 UI 翻译时，翻译给定文本。

## template_asset_view() 不懂

全名：template_asset_view(list_id, asset_library_dataptr, asset_library_propname, assets_dataptr, assets_propname, active_dataptr, active_propname, filter_id_types={}, display_options={}, activate_operator='', drag_operator='')
说明：项目。网格视图中的一个可滚动的资产列表
参数：

- list_id ：字符串， 不会为 None。该资产视图的标识符。用于区分不同的资产视图和识别从.blend 读取的资产视图
- asset_library_dataptr ：AnyType， 不会为 None。从活动资产库属性中获取的数据
- asset_library_propname ：字符串， 不会为 None。资产库属性的标识符
- assets_dataptr ：AnyType， 不会为 None。从资产清单属性中获取的数据
- assets_propname ：字符串， 不会为 None。资产清单属性的标识符
- active_dataptr ：AnyType， 不会为 None。从整数属性中获取的数据，活动项目的索引
- active_propname ：字符串， 不会为 None。active_data 中的整数属性的标识符，活动项目的索引。
- filter_id_types (enum set in {}，可选) 。filter_id_types
- display_options：显示资产视图的选项，可选 \* NO_NAMES：不在预览图像下显示每个资产的名称。

      * NO_FILTER：不显示用于过滤可用资产的按钮。

      * NO_LIBRARY：不显示选择或刷新资产库的按钮。

- activate_operator ：字符串，可选, 不为 None 。激活一个项目时要调用的自定义操作符的名称
- drag_operator ：字符串，可选, 不为 None。当开始拖动一个项目时要调用的自定义操作符的名称。从不与 active_operator（如果设置）一起调用，它要么是拖动，要么是激活。
  返回：(activate_operator_properties, drag_operator_properties)
  activate_operator_properties, Operator properties to fill in for the custom activate operator passed to the template, OperatorProperties
  drag_operator_properties, Operator properties to fill in for the custom drag operator passed to the template, OperatorProperties

## classmethodbl_rna_get_subclass() 不懂

全名：classmethodbl_rna_get_subclass(id, default=None)
参数：id (字符串)。The RNA type identifier.
返回：RNA 类型，未找到则为默认。
返回类型：bpy.types.Struct subclass

## classmethodbl_rna_get_subclass_py() 不懂

全名：classmethodbl_rna_get_subclass_py(id, default=None)
参数：id(字符串) 。The RNA type identifier.
返回：类，未找到则为默认。
返回类型：type

```

```
