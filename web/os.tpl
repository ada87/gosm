{{ template "head" .}}
<div class="container-fluid fullScreen">
<div class="row">
  <div class="col-xs-2 x-left">
  	<ul>
  		<li comand="net">网络</li>
  		<li comand="user">用户</li>
  		<li comand="detail">计算机信息</li>
  		<li comand="env">环境变量</li>
  	</ul>

  	<ul>
  		<li style="border:none;margin:0;"><h4>快速启动</h4></li>
  	</ul>
  	<ul class="faststart">
  		<li comand="write">写字板</li>
  		<li comand="notepad">记事本 </li>
  		<li comand="calc">计算器 </li>
  		<li comand="mspaint">画图</li>
  		<li comand="mstsc">远程桌面连接</li>

  		<li comand="devmgmt.msc">设备管理器</li>
  		<li comand="services.msc">本地服务设置</li>
  		<li comand="taskmgr">任务管理器</li>
  		<li comand="regedit">注册表</li>

  		<li comand="compmgmt.msc">计算机管理</li>
  		<li comand="fsmgmt.msc">共享文件夹管理器</li>
  		<li comand="msconfig">开机启动</li>
  		<li comand="winver">检查Windows版本</li>
	</ul>

  </div>
  <div id="ospanel" class="col-xs-10 x-right"></div>
</div>
</div>
{{template "foot" .}}