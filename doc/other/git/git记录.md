# 1. 基本命令

```
git status
git add .
git rm -r --cached + 要删除的暂存区的文件
git commit -m ''
git push origin master(branch)
```

```
git init
git add .
git commit -m ""
git branch -M main(第一次传需要创建新分支)
git remote add origin https://github.com/cresentboy/666.git          			git@github.com:cresentboy/cresentboy.github.io.git
git push -u origin main    （-u代表--set-upstream）
```

# 2. 报错non-fast-forward

error: failed to push some refs to 'https://github.com/cresentboy/resume.git'
 ! [rejected]        master -> master (non-fast-forward)

上传被拒绝，一般是因为在远程仓库别人改过了文件，自己本地没有修改，

先pull --rebase，解决冲突后再push
解决

```
$ git pull --rebase origin master/main
```

后续变成：(main|REBASE 1/1)

解决方案： (fix conflicts and then run "git rebase --continue")

 (use "git rebase --skip" to skip this patch)

(use "git rebase --abort" to check out the original branch)

```
$ git add .
$ git status
(all conflicts fixed: run "git rebase --continue")
$ git commit -m "merge path"
```

```
$ git rebase --continue
Successfully rebased and updated refs/heads/main.
```

```
$ git push -u origin master/main
```

或者强制覆盖远程代码：（不推荐）

```
$ git push -f origin master/main
```

报错：
lfs-large file storage

# 3. error: RPC failed

报错：
lfs-large file storage
error: RPC failed; curl 56 OpenSSL SSL_read: Connection was reset, errno 10054
send-pack: unexpected disconnect while reading sideband packet
fatal: the remote end hung up unexpectedly

方案：先设置大文件存储，然后：

```
git config http.postBuffer 524288000
```

# 4. GitHub存储大文件：

```
git init
git lfs install
git add .gitattributes
git commit -m ""
git branch -M main
git remote add origin [....]
git push -u origin main
```

 

```
从公用的仓库fetch代码：git fetch origin
git branch  [branchName]  新建分支
git checkout [branchName]      切换分支
git merge  [branchName]    合并分支
删除本地分支： git branch -d
强制删除，git branch -D
删除远程分支(慎用)：git push origin --delete
```

#  errno 10054

OpenSSL SSL_read: Connection was reset, errno 10054
产生原因：一般是这是因为服务器的SSL证书没有经过第三方机构的签署，所以才报错
参考网上解决办法：解除ssl验证后，再次git即可
git config --global http.sslVerify "false"

第一步：修改本地分支名称

```
git branch -m old_branch new_branch
```

第二步：删除远程分支

```
git push origin :old_branch
```

第三步：将本地新分支 push 到远程

```
git push origin new_branch
```

第四步：取消之前设置的 upstream

```
git branch --unset-upstream
```

第五步：将新分支设为 upstream

```
git push --set-upstream origin new_branch
```

备注

如果老分支配置过 upstream，前三步之后，本地仍然默认为老分支，不能用简洁的 git push 命令提交代码

进行第四步和第五步，将 upstream 设为新分支，就可以用 git push 命令来提交代码了

# fetch first

 ! [rejected]        master -> master (fetch first)
error: failed to push some refs to 'https://github.com/cresentboy/backend.git'

# unable to access

git push 时出现：
fatal: unable to access 'https://github.com/cresentboy/backend.git/': Failed to connect to github.com port 443: Timed out 

不能连接到远程仓库，排除网络问题和远程地址问题

# Connection reset by peer  

kex_exchange_identification: read: Connection reset by peer
Connection reset by 20.205.243.166 port 22
fatal: Could not read from remote repository.

解决方案：网络问题 重启电脑

# 记录

formalhaut@formalhaut MINGW32 ~/OneDrive/桌面/gitTest
$ git init
Initialized empty Git repository in C:/Users/formalhaut/OneDrive/桌面/gitTest/.git/

formalhaut@formalhaut MINGW32 ~/OneDrive/桌面/gitTest (master)
$ ll -a
total 28
drwxr-xr-x 1 formalhaut 197121 0 Jun  1 13:42 ./
drwxr-xr-x 1 formalhaut 197121 0 Jun  1 13:42 ../
drwxr-xr-x 1 formalhaut 197121 0 Jun  1 13:42 .git/

formalhaut@formalhaut MINGW32 ~/OneDrive/桌面/gitTest (master)
$ git status
On branch master

No commits yet

nothing to commit (create/copy files and use "git add" to track)

formalhaut@formalhaut MINGW32 ~/OneDrive/桌面/gitTest (master)
$ vim hello.txt

formalhaut@formalhaut MINGW32 ~/OneDrive/桌面/gitTest (master)
$ git status
On branch master

No commits yet

Untracked files:
  (use "git add <file>..." to include in what will be committed)
        hello.txt

nothing added to commit but untracked files present (use "git add" to track)

formalhaut@formalhaut MINGW32 ~/OneDrive/桌面/gitTest (master)
$ git add hello.txt
warning: LF will be replaced by CRLF in hello.txt.
The file will have its original line endings in your working directory

formalhaut@formalhaut MINGW32 ~/OneDrive/桌面/gitTest (master)
$ git status
On branch master

No commits yet

Changes to be committed:
  (use "git rm --cached <file>..." to unstage)
        new file:   hello.txt


formalhaut@formalhaut MINGW32 ~/OneDrive/桌面/gitTest (master)
$ git commit -m "test for git" hello.txt
warning: LF will be replaced by CRLF in hello.txt.
The file will have its original line endings in your working directory
[master (root-commit) f1d3a2f] test for git
 1 file changed, 1 insertion(+)
 create mode 100644 hello.txt

formalhaut@formalhaut MINGW32 ~/OneDrive/桌面/gitTest (master)
$ git status
On branch master
nothing to commit, working tree clean

formalhaut@formalhaut MINGW32 ~/OneDrive/桌面/gitTest (master)
$ vim hello.txt

formalhaut@formalhaut MINGW32 ~/OneDrive/桌面/gitTest (master)
$ git status
On branch master
Changes not staged for commit:
  (use "git add <file>..." to update what will be committed)
  (use "git restore <file>..." to discard changes in working directory)
        modified:   hello.txt

no changes added to commit (use "git add" and/or "git commit -a")

formalhaut@formalhaut MINGW32 ~/OneDrive/桌面/gitTest (master)
$ git add hello.txt
warning: LF will be replaced by CRLF in hello.txt.
The file will have its original line endings in your working directory

formalhaut@formalhaut MINGW32 ~/OneDrive/桌面/gitTest (master)
$ git status
On branch master
Changes to be committed:
  (use "git restore --staged <file>..." to unstage)
        modified:   hello.txt


formalhaut@formalhaut MINGW32 ~/OneDrive/桌面/gitTest (master)
$ git reflog
f1d3a2f (HEAD -> master) HEAD@{0}: commit (initial): test for git

formalhaut@formalhaut MINGW32 ~/OneDrive/桌面/gitTest (master)
$ git commit -m 'second commit' hello.txt
warning: LF will be replaced by CRLF in hello.txt.
The file will have its original line endings in your working directory
[master 1580915] second commit
 1 file changed, 2 insertions(+)

formalhaut@formalhaut MINGW32 ~/OneDrive/桌面/gitTest (master)
$ git status
On branch master
nothing to commit, working tree clean

formalhaut@formalhaut MINGW32 ~/OneDrive/桌面/gitTest (master)
$ git reflog
1580915 (HEAD -> master) HEAD@{0}: commit: second commit
f1d3a2f HEAD@{1}: commit (initial): test for git

formalhaut@formalhaut MINGW32 ~/OneDrive/桌面/gitTest (master)
$ git reset --hard f1d3a2f
HEAD is now at f1d3a2f test for git

formalhaut@formalhaut MINGW32 ~/OneDrive/桌面/gitTest (master)
$ git staus
git: 'staus' is not a git command. See 'git --help'.

The most similar command is
        status

formalhaut@formalhaut MINGW32 ~/OneDrive/桌面/gitTest (master)
$ git status
On branch master
nothing to commit, working tree clean

formalhaut@formalhaut MINGW32 ~/OneDrive/桌面/gitTest (master)
$ git reflog
f1d3a2f (HEAD -> master) HEAD@{0}: reset: moving to f1d3a2f
1580915 HEAD@{1}: commit: second commit
f1d3a2f (HEAD -> master) HEAD@{2}: commit (initial): test for git

formalhaut@formalhaut MINGW32 ~/OneDrive/桌面/gitTest (master)
$ cat hello.txt
test

formalhaut@formalhaut MINGW32 ~/OneDrive/桌面/gitTest (master)
$ git branch -v
* master f1d3a2f test for git

formalhaut@formalhaut MINGW32 ~/OneDrive/桌面/gitTest (master)
$ git branch hot-fix

formalhaut@formalhaut MINGW32 ~/OneDrive/桌面/gitTest (master)
$ git branch -v
  hot-fix f1d3a2f test for git
* master  f1d3a2f test for git

formalhaut@formalhaut MINGW32 ~/OneDrive/桌面/gitTest (master)
$ vim hello.txt

formalhaut@formalhaut MINGW32 ~/OneDrive/桌面/gitTest (master)
$ cat hello.txt
test
fneownfweoibgvuwa

formalhaut@formalhaut MINGW32 ~/OneDrive/桌面/gitTest (master)
$ git sta
git: 'sta' is not a git command. See 'git --help'.

The most similar commands are
        status
        stage
        stash

formalhaut@formalhaut MINGW32 ~/OneDrive/桌面/gitTest (master)
$ git status
On branch master
Changes not staged for commit:
  (use "git add <file>..." to update what will be committed)
  (use "git restore <file>..." to discard changes in working directory)
        modified:   hello.txt

no changes added to commit (use "git add" and/or "git commit -a")

formalhaut@formalhaut MINGW32 ~/OneDrive/桌面/gitTest (master)
$ git add hello.txt

formalhaut@formalhaut MINGW32 ~/OneDrive/桌面/gitTest (master)
$ git status
On branch master
Changes to be committed:
  (use "git restore --staged <file>..." to unstage)
        modified:   hello.txt


formalhaut@formalhaut MINGW32 ~/OneDrive/桌面/gitTest (master)
$ git commit -m 'my forth commit' hello.txt
[master f802532] my forth commit
 1 file changed, 1 insertion(+)

formalhaut@formalhaut MINGW32 ~/OneDrive/桌面/gitTest (master)
$ git status
On branch master
nothing to commit, working tree clean

formalhaut@formalhaut MINGW32 ~/OneDrive/桌面/gitTest (master)
$ git branch -v
  hot-fix f1d3a2f test for git
* master  f802532 my forth commit

formalhaut@formalhaut MINGW32 ~/OneDrive/桌面/gitTest (master)
$ git checkout hot-fix
Switched to branch 'hot-fix'

formalhaut@formalhaut MINGW32 ~/OneDrive/桌面/gitTest (hot-fix)
$ cat hello.txt
test

formalhaut@formalhaut MINGW32 ~/OneDrive/桌面/gitTest (hot-fix)
$ vim hello.txt

formalhaut@formalhaut MINGW32 ~/OneDrive/桌面/gitTest (hot-fix)
$ cat hello.txt
test
wnepiof

fnweo

nfeow
now

formalhaut@formalhaut MINGW32 ~/OneDrive/桌面/gitTest (hot-fix)
$ git status
On branch hot-fix
Changes not staged for commit:
  (use "git add <file>..." to update what will be committed)
  (use "git restore <file>..." to discard changes in working directory)
        modified:   hello.txt

no changes added to commit (use "git add" and/or "git commit -a")

formalhaut@formalhaut MINGW32 ~/OneDrive/桌面/gitTest (hot-fix)
$ git add hello.txt

formalhaut@formalhaut MINGW32 ~/OneDrive/桌面/gitTest (hot-fix)
$ git status
On branch hot-fix
Changes to be committed:
  (use "git restore --staged <file>..." to unstage)
        modified:   hello.txt


formalhaut@formalhaut MINGW32 ~/OneDrive/桌面/gitTest (hot-fix)
$ git commit -m 'hot-fix branch commit' hello.txt
[hot-fix 5b6ef9c] hot-fix branch commit
 1 file changed, 6 insertions(+)

formalhaut@formalhaut MINGW32 ~/OneDrive/桌面/gitTest (hot-fix)
$ git status
On branch hot-fix
nothing to commit, working tree clean

formalhaut@formalhaut MINGW32 ~/OneDrive/桌面/gitTest (hot-fix)
$ git reflog
5b6ef9c (HEAD -> hot-fix) HEAD@{0}: commit: hot-fix branch commit
f1d3a2f HEAD@{1}: checkout: moving from master to hot-fix
f802532 (master) HEAD@{2}: commit: my forth commit
f1d3a2f HEAD@{3}: reset: moving to f1d3a2f
1580915 HEAD@{4}: commit: second commit
f1d3a2f HEAD@{5}: commit (initial): test for git

formalhaut@formalhaut MINGW32 ~/OneDrive/桌面/gitTest (hot-fix)
$ git branch
* hot-fix
  master

formalhaut@formalhaut MINGW32 ~/OneDrive/桌面/gitTest (hot-fix)
$ git branch -v
* hot-fix 5b6ef9c hot-fix branch commit
  master  f802532 my forth commit

formalhaut@formalhaut MINGW32 ~/OneDrive/桌面/gitTest (hot-fix)
$ git merge master
Auto-merging hello.txt
CONFLICT (content): Merge conflict in hello.txt
Automatic merge failed; fix conflicts and then commit the result.

formalhaut@formalhaut MINGW32 ~/OneDrive/桌面/gitTest (hot-fix|MERGING)
$ git status
On branch hot-fix
You have unmerged paths.
  (fix conflicts and run "git commit")
  (use "git merge --abort" to abort the merge)

Unmerged paths:
  (use "git add <file>..." to mark resolution)
        both modified:   hello.txt

no changes added to commit (use "git add" and/or "git commit -a")

formalhaut@formalhaut MINGW32 ~/OneDrive/桌面/gitTest (hot-fix|MERGING)
$ git merge --abort

formalhaut@formalhaut MINGW32 ~/OneDrive/桌面/gitTest (hot-fix)
$ git status
On branch hot-fix
nothing to commit, working tree clean

formalhaut@formalhaut MINGW32 ~/OneDrive/桌面/gitTest (hot-fix)
$ git checkout master
Switched to branch 'master'

formalhaut@formalhaut MINGW32 ~/OneDrive/桌面/gitTest (master)
$ git branch -v
  hot-fix 5b6ef9c hot-fix branch commit
* master  f802532 my forth commit

formalhaut@formalhaut MINGW32 ~/OneDrive/桌面/gitTest (master)
$ git merge hot-fix
Auto-merging hello.txt
CONFLICT (content): Merge conflict in hello.txt
Automatic merge failed; fix conflicts and then commit the result.

formalhaut@formalhaut MINGW32 ~/OneDrive/桌面/gitTest (master|MERGING)
$ cat hello.txt
test
<<<<<<< HEAD

fneownfweoibgvuwa

wnepiof

fnweo

nfeow
now

\> \> \> \> \> \> \> hot-fix

formalhaut@formalhaut MINGW32 ~/OneDrive/桌面/gitTest (master|MERGING)
$ vim hello.txt

formalhaut@formalhaut MINGW32 ~/OneDrive/桌面/gitTest (master|MERGING)
$ git add hello.txt

formalhaut@formalhaut MINGW32 ~/OneDrive/桌面/gitTest (master|MERGING)
$ git status
On branch master
All conflicts fixed but you are still merging.
  (use "git commit" to conclude merge)

Changes to be committed:
        modified:   hello.txt


formalhaut@formalhaut MINGW32 ~/OneDrive/桌面/gitTest (master|MERGING)
$ git commit -m -m 'merge hot-fix'
fatal: cannot do a partial commit during a merge.

formalhaut@formalhaut MINGW32 ~/OneDrive/桌面/gitTest (master|MERGING)
$ git commit -m 'merge hot-fix'
[master 1516402] merge hot-fix

formalhaut@formalhaut MINGW32 ~/OneDrive/桌面/gitTest (master)
$