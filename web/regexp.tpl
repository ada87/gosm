{{ template "head" .}}
{{define "regitem"}}
<h2>{{ .Field_desc }}</h2>
	{{range .Values}}
	<p>{{.Value}} / {{.Desc}} </p>
	{{end}}
{{end}}
<div class="fullScreen">
<ul class="reg9g">
	<li class="reg3h">
		<ul>
			<li class="item regb1">{{ template "regitem" .Data.regexp_form}}</li>
			<li class="item regb2 borleft borright">输入<button type="button" class="btn" id="edit">编辑</button></li>
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
<div class="editpanel"></div>
</div>
{{template "foot" .}}