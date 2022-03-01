/* SPDX-License-Identifier: GPL-2.0 */
#ifndef _LINUX_IPC_H
#define _LINUX_IPC_H

#include <linux/spinlock.h>
#include <linux/uidgid.h>
#include <linux/rhashtable-types.h>
#include <uapi/linux/ipc.h>
#include <linux/refcount.h>

/* used by in-kernel data structures */
struct kern_ipc_perm {
	spinlock_t	lock;
	bool		deleted;
	int		id;
	key_t		key; // 键（系统支持两种键：公有和私有）
	kuid_t		uid; // 对象拥有者对应进程的有效用户识别号和有效组织识别号
	kgid_t		gid;
	kuid_t		cuid; // 对象创建者对应进程的有效用户识别号和有效组识别号
	kgid_t		cgid;
	umode_t		mode; // 存取模式
	unsigned long	seq; // 序列号
	void		*security;

	struct rhash_head khtnode;

	struct rcu_head rcu;
	refcount_t refcount;
} ____cacheline_aligned_in_smp __randomize_layout;

#endif /* _LINUX_IPC_H */
