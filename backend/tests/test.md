初始化 创建meta √

再次初始化  √ 会修改metaMaps.Local.Remain   由于第一次会在父目录创建meta.json 会导致父目录变化(无法避免, 因为无法确定是创建meta产生的变化还是创建其他文件产生的变化)

读取README.md √
读取meta.json √

修改README.md内容
 父级文件夹变动

随便添加非md文件 √
 父文件夹变动=>重新读取README=>dbDatas.Updates=>collections.Updates
随便删除非md文件 √
 父文件夹变动=>重新读取README=>dbDatas.Updates=>collections.Updates

修改文件内容
 只会影响文件本身=>dbDatas.Updates=>collections.Updates

添加md文件(ORDER有问题)
 数据库增加文件 =>collections.Creates
 文件变动导致父级文件夹变化 =>dbDatas.Updates
 重新规划meta.json=>metaMaps.Local.Remain

添加空文件夹
 数据库增加文件夹 =>collections.Creates
 文件变动导致父级文件夹变化 =>dbDatas.Updates
 重新规划meta.json=>metaMaps.Local.Remain
