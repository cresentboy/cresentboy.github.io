# 如何成为 Apache 孵化器项目

项目发起者在申请Apache孵化器项目前一定要想清楚自己有时间可以从事开源建设工作，大家的精力总是有限的，服务于公司内部业务和开源社区之间往往有一方要牺牲一部分。

如果可以的话，建议要尽早地去构建开源社区。如果前期花费精力构建好了项目的开源社区，吸引了大量积极的贡献者，会对项目的发展和可持续性大有裨益。但另一方面，没有建立社区的项目往往在后期运营和投入上会非常痛苦。

- 当前处于孵化器状态的项目在: [All Incubator Projects](http://incubator.apache.org/projects/index.html)
- 项目孵化状态：[Incubator Clutch Status](https://incubator.apache.org/clutch/) 
- 最好的Apache Incubator入门文档在: [The Apache Incubator Cookbook](https://incubator.apache.org/cookbook/)

## 孵化器申请书

申请孵化器项目首先需要撰写申请书。

> It' s not required to have a good proposal, but having a good proposal will increase the chances of a positive outcome.

申请书通常有一套既定的模板，你需要详述项目的定位和挑战，团队是否曾经有过开源经验，是否有参与Apache项目的经历等等，这些信息能有助于让社区接纳你的项目。你可以参考其他孵化中或者已毕业项目的申请书：

- [Proposals - INCUBATOR - Apache Software Foundation](https://cwiki.apache.org/confluence/display/INCUBATOR/Proposals)

申请书写好后，我们要发给IPMC（Incubator Project Management Committee），具体方式是直接发邮件给 general@incubator.apache.org。

发邮件之前首先需要**订阅** general@incubator.apache.org 邮件列表，否则邮件会被卡住无法发出。跟订阅ASF的任何邮件列表一样，你可以手动发一则空邮件(随意标题，无内容)至 [general-subscribe@incubator.apache.org](mailto:general-subscribe@incubator.apache.org) ，该邮件组的robot会自动引导你进行订阅。

## 公开提案讨论

讨论通常有3个目的：**1. 寻找导师 2. 接受提醒 3. 收获关注**。

**1. 寻找导师**

想要加入ASF的项目(通常称作 podlings)首先需要1个champion，至少2个mentor负责作为ASF的中间人，这个人需要是 Apache Incubator PMC (简称 IPMC) 成员。

- IPMC 人员列表: [Apache Phone Book](http://people.apache.org/phonebook.html?pmc=incubator)

> To enter the Incubator, your project needs a ***champion\*** and at least two or three ***mentors\***. These people need to be part of the Incubator PMC, which ASF Members can join just by asking.

在申请书公开发布的过程中，可能会有许多慕名而来的IPMC自荐为Champion或Mentor。他们会作为你项目的长期导师，引导你走接下来的流程。理论上，如果你希望事无巨细一一询问你的导师，那么本文后续内容你全都不用再看。你可以提早获取导师的联系方式，并时常咨询。当然如果你希望预先了解流程，请往下再阅读。

**2. 接受提醒**

你的提案可能会收获他人的提醒与警示。例如项目名称是否合适，项目的定位是否与已有Apache项目重叠，项目的第三方库License是否包含GPL等与Apache不兼容的开源协议等等。

**3. 收获关注**

> A good proposal should target the wider audience and not just the IPMC. Use this time to engage and inform potential developers and users.

一个好的做法是趁热打铁，在进入孵化器前，发布申请书时做项目推广，往往能事半功倍。

**4. 结束讨论**

讨论的结束取决于你有没有找到足够多的导师。如上所述，你必须在集齐至少2名mentors，1名champion之后才能继续下一阶段。

**5. 案例**

- Dubbo的Proposal：[https://lists.apache.org/thread.html/1e4a74cc5af9cd0e298dc6085d5be57da8dfd02fef1ebd33829a6084@%3Cgeneral.incubator.apache.org%3E](https://lists.apache.org/thread.html/1e4a74cc5af9cd0e298dc6085d5be57da8dfd02fef1ebd33829a6084@)
- Doris的Proposal：(Palo时期) [https://lists.apache.org/thread.html/74a3f3f945403b50515c658047d3284955288a637207e4f97ecc15d1@%3Cgeneral.incubator.apache.org%3E](https://lists.apache.org/thread.html/74a3f3f945403b50515c658047d3284955288a637207e4f97ecc15d1@)

## 孵化器投票

为了更好地发起号召，在投票时建议通知你的 champion 和 mentors。由他们先投一票，从而吸引关注者。又或者你可以直接让 champion 来发起投票，这样更能起个好头。

## IP Clearance

投票成功后基本可以认为项目成功进入Apache孵化器。后续的工作就是 IP Clearance，即知识产权的合规检查。这是一个长期的工作，我们需要严格保证代码里声明了Apache的版权。

参考：

[Podling IP Clearanceincubator.apache.org/guides/ip_clearance.html](https://incubator.apache.org/guides/ip_clearance.html)

**1. 授权协议**

为了以合法的形式正式捐献项目到Apache Incubator，接下来有以下事务：

1 如果以公司为捐献方，我们需由公司签署 [SGA (Software Grant Agreement)](https://www.apache.org/licenses/software-grant-template.pdf)。

2 每个initial committer都需要签署 [ICLA (Individual Contributor License Agreement)](https://www.apache.org/licenses/icla.pdf)，从而获得一个Apache账号。

> Please send only one document per email.

上述文件都需发送给Apache秘书处（secretary@apache.org）。注意一封邮件只发一份文件。

参考：

- [ASF Contributor Agreements](https://www.apache.org/licenses/contributor-agreements.html#submitting)

**2. PPMC (Podling Project Management Committee)**

参考：

- [Podling Project Management Committee](https://incubator.apache.org/guides/ppmc.html#voting_in_a_new_ppmc_member)

**3. 基础服务搭建**

1 需要由 champion 创建一个JIRA主站，存放项目的Issues。

2 同时也需要搭建孵化状态网站，这个网站会显示当前 孵化进展，记录我们开源的重要时间点，例如新 committer 的加入，例如版权检查的完成时间。

3 每个Apache项目都会有多个邮件列表，这样后来者也可以检索邮件列表了解状况。你可以通过网页 [ASF Mailing List Subscription Helper](https://whimsy.apache.org/committers/subscribe) 完成对项目邮件列表的订阅。