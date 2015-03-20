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
  </div>
  <div id="ospanel" class="col-xs-10 x-right"></div>
</div>
</div>
{{template "foot" .}}