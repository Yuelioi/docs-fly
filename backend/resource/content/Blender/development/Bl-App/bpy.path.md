---
order: 10
title: 路径工具(bpy.path)
---

# 路径工具(bpy.path)

## Description

本模块与 os.path 类似，在 Blender 中处理各种路径。

### abspath()

全名：bpy.path.abspath(path, \*, start=None, library=None)

说明：返回相对于当前 Blender 文件的绝对路径，使用"//"前缀。

参数：

- start (string or bytes) ： 相对于此路径，不设置则使用当前文件名。
- library (bpy.types.Library) ： 这个路径来自的库。当没有库的时候，它的路径会取代 start。

返回：绝对路径。

返回类型：字符串

### basename()

全名：bpy.path.basename(path)

说明：相当于 os.path.basename，但跳过"//"前缀。为了与 Windows 兼容而使用。给定路径的基本名称。

返回类型： string

### clean_name()

全名：bpy.path.clean*name(name, \*, replace='*')

说明：返回一个被替换了字符的名称，这些字符在各种情况下可能会导致问题。

例如写到文件中。除了 A：Z/a：z，0：9 以外的所有字符都被替换为"\_"或替换参数：（如果定义了）。

参数：

- name: 路径名。string or bytes
- replace: 字符串

返回： 替换无效的字符。清理后的名称。

返回类型： 字符串

### display_name()

全名：bpy.path.display_name(name, \*, has_ext=True, title_case=True)

说明：从名称中创建一个显示字符串，用于菜单和用户界面。适用于文件名和模块名。

参数：

- name (string) ： 用来显示用户界面的名称。
- has_ext (boolean) ： 从名称中移除文件扩展名。
- title_case (boolean) ： 将小写的名字转换成标题大小写。

返回：显示的字符串。

返回类型：字符串

### display_name_to_filepath()

全名：bpy.path.display_name_to_filepath(name)

说明：使用文件路径中不支持的字符的文字版本，执行 display_name 的反向操作。

参数：

- name: 要转换的显示名称，string

返回： 文件路径。

返回类型： 字符串

### display_name_from_filepath()

全名：bpy.path.display_name_from_filepath(name)

说明：返回去除目录和扩展名的路径，确保与 utf8 兼容。

参数： name，要转换的文件路径。显示名称。

返回类型： string

### ensure_ext()

全名：bpy.path.ensure_ext(filepath, ext, \*, case_sensitive=False)

说明：如果还没有设置扩展名，则返回已添加扩展名的路径。

参数：

- filepath (string) ： 文件路径。
- ext (string) ： 要检查的扩展名，可以是一个复合扩展名。应该以点开始，如'.blend'或'.tar.gz'。
- case_sensitive (boolean) ： 在比较扩展名时检查是否匹配大小写。

返回：具有给定扩展名的文件路径。

返回类型：字符串

### is_subdir()

全名：bpy.path.is_subdir(path, directory)

说明：如果路径在目录的子目录中，返回 true。两个路径必须是绝对的。

参数：path (string or bytes) ， 一个绝对路径。

返回：路径是否是一个子目录。

返回类型：布尔型

### module_names()

全名：bpy.path.module_names(path, \*, recursive=False)

说明：返回一个可以从路径导入的模块列表。

参数：

- path (string) ： 一个要扫描的目录。
- recursive (bool) ： 同时返回软件包的子模块名称。

返回：字符串列表（module_name, module_file）。

返回类型：列表

### native_pathsep()

全名：bpy.path.native_pathsep(path)

说明：用系统本地的 os.sep 来替换路径分隔符。

参数： path，要替换的路径，string

返回： 使用系统本地分隔符的路径。

返回类型： string

### reduce_dirs()

全名：bpy.path.reduce_dirs(dirs)

说明：给定一个目录序列，删除重复的目录和嵌套在其他路径中的任何目录。(对递归路径搜索很有用）。

参数：dirs (字符串的序列) ， 目录路径的序列。

返回：一个唯一的路径列表。

返回类型：字符串列表

### relpath()

全名：bpy.path.relpath(path, \*, start=None)

说明：返回相对于当前混合文件的路径，使用"//"前缀。

参数：

- path (string or bytes) ： 绝对路径。
- start（字符串或字节） ： 相对于这个路径，没有设置则使用当前文件名。

返回：相对路径。

返回类型：字符串

### resolve_ncase()

全名：bpy.path.resolve_ncase(path)

说明：在一个区分大小写的系统上解析一个不区分大小写的路径，如果找到的话，返回一个带有该路径的字符串，否则返回原始路径。

参数： path。要解析的路径名称，string

返回： 解析后的路径。

返回类型： string
