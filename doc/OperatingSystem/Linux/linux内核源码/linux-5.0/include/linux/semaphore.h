/*
 * Copyright (c) 2008 Intel Corporation
 * Author: Matthew Wilcox <willy@linux.intel.com>
 *
 * Distributed under the terms of the GNU GPL, version 2
 *
 * Please see kernel/locking/semaphore.c for documentation of these functions
 */
#ifndef __LINUX_SEMAPHORE_H
#define __LINUX_SEMAPHORE_H

#include <linux/list.h>
#include <linux/spinlock.h>

/* 信号量数据结构 */
struct semaphore {
	// lock是自旋锁变量，用于保护semaphore数据结构里面的count/wait_list成员
	raw_spinlock_t		lock; 

	// 用来表示允许进入临界区的内核执行路径个数
	unsigned int		count; 

	// 使用链表来管理所有在该信号量上睡眠的进程，没有成功获取锁的进程会在这个链表上睡眠
	struct list_head	wait_list;
};

#define __SEMAPHORE_INITIALIZER(name, n)				\
{									\
	.lock		= __RAW_SPIN_LOCK_UNLOCKED((name).lock),	\
	.count		= n,						\
	.wait_list	= LIST_HEAD_INIT((name).wait_list),		\
}

#define DEFINE_SEMAPHORE(name)	\
	struct semaphore name = __SEMAPHORE_INITIALIZER(name, 1)

// 通过此函数进行信号时的初始化操作，其中__SEMAPHORE_INITIALIZER()宏会完成对semaphore数据结构的填充，
// val值通常设为1
static inline void sema_init(struct semaphore *sem, int val)
{
	static struct lock_class_key __key;
	*sem = (struct semaphore) __SEMAPHORE_INITIALIZER(*sem, val);
	lockdep_init_map(&sem->lock.dep_map, "semaphore->lock", &__key, 0);
}

extern void down(struct semaphore *sem);

// 此函数在争用信号量失败时进入可中断的睡眠状态，而down()函数进入不可中断的睡眠状态
extern int __must_check down_interruptible(struct semaphore *sem);
extern int __must_check down_killable(struct semaphore *sem);

// 此函数返回0表示成功获取锁，返回1表示获取锁失败
extern int __must_check down_trylock(struct semaphore *sem);
extern int __must_check down_timeout(struct semaphore *sem, long jiffies);
extern void up(struct semaphore *sem);

#endif /* __LINUX_SEMAPHORE_H */
