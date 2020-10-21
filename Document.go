package ZYB
/**
	##项目链接地址
	#黑蝙蝠 http://xuxiao-docker.suanshubang.com/static/blackbats/index.html#/agency-manage?title=%E4%BB%A3%E7%90%86%E5%95%86%E6%9C%BA%E6%9E%84%E7%AE%A1%E7%90%86

    #金丝雀 http://xuxiao-docker.suanshubang.com/static/hy/kdrpmis/index.html#/login

	部署分支
	//金丝雀
	gg deploy afxmis dev

	//黑蝙蝠
	gg deploy actplatmis dev

	//金丝雀fe
	gg deployfe kdrpmis issue#05

	//黑蝙蝠fe
	gg deployfe blackbats issues#5
*/

/**
###办公相关地址
#Docker容器  http://docker.suanshubang.com/#/login
#git远程服务器 https://git.zuoyebang.cc/
#wiki项目需求 http://wiki.zuoyebang.cc/pages/viewpage.action?pageId=47455223
#wiki 团队简介 http://wiki.zuoyebang.cc/pages/viewpage.action?pageId=12883276
#思维导图 https://www.processon.com
#接口文档http://yapi.zuoyebang.cc/group/907
#需求空间 http://wiki.zuoyebang.cc/pages/viewpage.action?pageId=65384307
#运维op http://op.zuoyebang.cc/static/odin/index.html#/deploy/module
#relay http://ssp.zuoyebang.cc/index.php?action=resetrelay

##办公软件
# 钉钉
# Foxmail  邮件相关
*/

/**
##开发工具相关
# phpStorm
# goLand
# xShell
# gitBash
# postMan
# Navicat
# Authy  Relay账号口令生成器

###开发环境
goLang
*/
/***
  	gitBash ---相关操作
	 A、####配置git用户和邮箱
      # 设置用户姓名
  　　git config --global user.name "姓名"
  　　# 设置可以联系的邮箱地址
  　　git config --global user.email "联系邮箱"
  　　# 查看设置信息
  　　git config --global --list
  	B、####生成SSH密钥
  		//密钥命令
  		ssh-keygen -t rsa
  		//id_rsa 文件是私钥，一定保存好不能泄露；id_rsa.pub 文件是公钥，内容放在云端提供验证。
	C、####云端验证
  		登录 https://git.zuoyebang.cc/ 点击自己头像 --setings ---add SSHKEy 添加保存即可
  	D、####克隆
  		git clone -b 分支名 git服务器地址
    E grep -rn ">>>>>>" *  //冲突文件 vim 编辑修改

*/

/**
	项目 pull 和 push

	1、进入对应git分支上传代码
		git add .
		git commit -m 'xpx'
		git pull origin 分支名
		git push origin 分支名
	2、进入replay
		打开xShell 新建会话 主机名：relay.zuoyebang.cc 登录名： 邮箱前缀
		选择使用键盘输入身份验证 如果需要输入登录名  邮箱前缀
		Verification code: 手机打开 Authy 输入动态口令
		输入密码：一致密码

   3、relay 链接Docker
		ssh rd@172.29.254.7
		PassWord:MhxzKhl --M和K大写
   4、常用命令
	goto        进入各自容器
    redis       链接公共redis
    stored      链接公共stored
    memcached   链接公共memcached
    db51        链接公共数据库5.1
    db57        链接公共数据库5.7

   5、进入容器
		goto xuxiao(项目) zhibo（容器）
	    此时文件目录为pdf ODP框架
   6、上传完git 拉代码
	 gg deploy app(为app目录下的模块) 分支

	7、浏览器 访问
		由于前后端分离----访问页面为前端 ---需部署前端的页面 gg deployfe blackbats issues#3
		登录 ----18600000000 发送验证码 1111
	8、postMan访问
		常用 用于调试接口
*/

/**
NaviCat 链接mySql相关
	 /home/homework/conf/db/cluster/对应项目的集群
打开 找到ip 和 端口  username 和 password 链接即可
*/

/**
Ap框架规范 ---略 和前辈写的统一即可
*/

/**
日志相关
/home/homework/clog   /php php错误日志 || /webserver nginx 访问和错误日志

//控制器错误日志---
	/home/homework/log/项目目录/afxmis.log.wf ---错误日志
*/

/**
###重新启动php
/home/homework/php/sbin/php-fpm restart

####重启nginx
webserver loadnginx.sh start
*/

/**
	#####周报相关
	发送人sunxiancan@zuoyebang.com;yikesellrd-op@zuoyebang.com
    抄送 zhengchangshuai@zuoyebang.com

*/
/**
定时脚本
http://cron-zbcore.zuoyebang.cc
用户：rd
密码：tHxi5k0Q
*/

/**
	数据库操作申请
    http://rds.zuoyebang.cc/service/455#jira_sql_job
    自助调教
	生成执行ID

	jira权限申请
	http://jira.zuoyebang.cc/browse/JOBS-47566
	工单系统--自助修改线上数据库数据---输入执行ID
	创建成功---发送献灿审批
    审批完解决问题---回归rds 验证问题正确性
	连接备库---查看
*/

/**
	代码上线申请
    合并到dev分支
	pms：请求http://pms.zuoyebang.cc/all?searchText=%E8%AF%95%E5%90%AC%E8%AF%BE
    上线
*/

/**
tips 权限申请---
http://op.zuoyebang.cc/static/odin/#/monitor/graph
权限管理 -- 权限申请---机器临时权限--机器列表--
192.168.0.122
192.168.128.86
192.168.133.30
192.168.7.171
192.168.9.149
192.168.9.151
192.168.128.121
192.168.128.125
192.168.128.129
192.168.128.208
192.168.128.212
192.168.128.214
192.168.128.216
192.168.128.117
192.168.128.94
192.168.128.98
172.30.128.7
172.29.208.232
172.29.208.236
172.29.208.238
172.29.209.20
172.29.209.30
172.29.211.14
192.168.133.117
192.168.133.118

联系chenkun

	tips
 ssh rd@192.168.0.122
 fsh na tips
fsh all actplat(集群) ssh "tail "
fsh all actplat ssh "tail /home/homework/log/afxmis/afxmis.log.wf"

wiki
http://wiki.zuoyebang.cc/pages/viewpage.action?pageId=123334039
线上日志查看

金丝雀线上地址：
http://www.zybang.com/static/hy/kdrpmis/index.html#/dashboard
黑蝙蝠线上地址：
http://actmis.zuoyebang.cc/static/blackbats/index.html#/home

联系人：梁博
*/

/**
   okr: https://i.zuoyebang.cc/#/news
 */


//uda权限
var UDA = "http://uda.zuoyebang.cc"

