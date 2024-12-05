---
title: Memory-Allocation
order: 13
---

# 内存分配

使用 After Effects 来处理任何相当大的内存分配。对于小规模的分配，可以使用 new 和 delete，但这是例外。在低内存条件下(如在 RAM 预览期间)，插件要优雅地处理内存不足，而不是与 After Effects 争夺操作系统内存，通过使用内存分配函数，After Effects 知晓何时释放缓存的图像，以避免内存交换。如果不使用我们的函数来分配大量的内存，会导致锁定、崩溃和技术支持电话。务必别这样做。

如果要包装现有的 C++类，请创建一个基类，为该类实现 new 和 delete，并从它派生。为了重载 STL，我们不建议重载全局 new 和 delete。相反，要提供一个分配器作为模板定义的一部分。

After Effects 传递给你的处理程序在被调用之前会被锁定，一旦返回就会被解锁。

## PF_HandleSuite1

| **函数**           | **目的**                                                                                                                                          | **替换**         |
| ---------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------- |
| `host_new_handle`      | 分配一个新句柄<br />`PF_Handle (*host_new_handle)(A_HandleSize size);`                                                                      | `PF_NEW_HANDLE`      |
| `host_lock_handle`     | 锁定一个句柄<br />`void (*host_lock_handle)(PF_Handle pf_handle);`                                                                                | `PF_LOCK_HANDLE`     |
| `host_unlock_handle`   | 解锁一个句柄<br />`void (*host_unlock_handle)( PF_Handle pf_handle);`                                                                           | `PF_UNLOCK_HANDLE`   |
| `host_dispose_handle`  | 释放一个句柄<br />`void (*host_dispose_handle)( PF_Handle pf_handle);`                                                                            | `PF_DISPOSE_HANDLE`  |
| `host_get_handle_size` | 返回其句柄被传入的可重新分配块的大小（以字节为单位）。<br />A_HandleSize (*host_get_handle_size)( PF_Handle pf_handle);` | `PF_GET_HANDLE_SIZE` |
| `host_resize_handle`   | 句柄重新分配size<br />`PF_Err (*host_resize_handle)( A_HandleSize new_sizeL, PF_Handle *handlePH);`                                                 | `PF_RESIZE_HANDLE`   |
