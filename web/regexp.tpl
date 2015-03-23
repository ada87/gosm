{{ template "head" .}}
{{define "regitem"}}
<h2 id="xx">{{ .Field_desc }}</h2>
	{{range .Values}}
	<p class="check">{{.Value}}</p>
	{{end}}
{{end}}
<div class="fullScreen">
<ul class="reg9g">
	<li class="reg3h">
		<ul>
			<li class="item regb1">{{ template "regitem" .Data.regexp_form}}</li>
			<li class="item regb2 borleft borright">
				<div class="input-group" style="width:70%;margin:20px auto;">
		          <input type="text" id="txtreg" class="form-control">
		        </div>
				<button type="button" class="button button-primary" id="test">测试</button>
				<button type="button" class="button button-royal button-small" id="edit">编辑</button>
			</li>
			<li class="item regb3">{{ template "regitem" .Data.regexp_number}}</li>
		</ul>
	</li>
	
	<li class="reg3h">
		<ul>
			<li class="item regb4 bortop borbottom">{{ template "regitem" .Data.regexp_custom1}}</li>
			<li class="item regb5 borleft borright bortop borbottom">输出</li>
			<li class="item regb6 bortop borbottom">{{ template "regitem" .Data.regexp_custom2}}</li>	
		</ul>
	</li>
	<li class="reg3h">
		<ul>
			<li class="item regb7">{{ template "regitem" .Data.regexp_sl}}</li>
			<li class="item regb8 borleft borright">{{ template "regitem" .Data.regexp_xss}}</li>
			<li class="item regb9">{{ template "regitem" .Data.regexp_sp}}</li>		
		</ul>
	</li>
</ul>
<div class="editpanel container">
	<nav>
		<button class="button button-highlight" type="button" id="back">返回</button>
		<strong>正则库</strong>
	</nav>
  	
	<div role="tabpanel">

	<ul class="nav nav-tabs nav-pills nav-justified" role="tablist">
      <li class="active" role="presentation">
      	<a href="#regexp_form">{{ .Data.regexp_form.Field_desc}}</a>
      </li>
      <li role="presentation">
      	<a href="#regexp_number">{{ .Data.regexp_number.Field_desc}}</a>
      </li>
      <li role="presentation">
      	<a href="#regexp_custom1">{{ .Data.regexp_custom1.Field_desc}}</a>
      </li>
      <li role="presentation">
      	<a href="#regexp_custom2">{{ .Data.regexp_custom2.Field_desc}}</a>
      </li>
      <li role="presentation">
      	<a href="#regexp_sl">{{ .Data.regexp_sl.Field_desc}}</a>
      </li>
      <li role="presentation">
      	<a href="#regexp_xss">{{ .Data.regexp_xss.Field_desc}}</a>
      </li>
      <li role="presentation">
      	<a href="#regexp_sp">{{ .Data.regexp_sp.Field_desc}}</a>
      </li>
    </ul>

  	<div class="tab-content">
 		<div role="tabpanel" class="tab-pane fade in active" id="regexp_form">form</div>
    	<div role="tabpanel" class="tab-pane fade" id="regexp_number">regexp_number</div>
    	<div role="tabpanel" class="tab-pane fade" id="regexp_custom1">regexp_custom1</div>
    	<div role="tabpanel" class="tab-pane fade" id="regexp_custom2">regexp_custom2</div>
    	<div role="tabpanel" class="tab-pane fade" id="regexp_sl">regexp_sl</div>
    	<div role="tabpanel" class="tab-pane fade" id="regexp_xss">regexp_xss</div>
    	<div role="tabpanel" class="tab-pane fade" id="regexp_sp">regexp_sp</div>
  	</div>
	</div>
</div>
</div>
{{template "foot" .}}