{{ template "head" .}}
<div class="index">
	<h1>{{.Config.Title}}</h1>
	<p class="lead">Author:xd</p>
	<ul>
	{{range $tab := .Config.Tabs}}
		<li><a href="/{{$tab.Handle}}">{{$tab.Title}}</a></li>
	{{end}}
	</ul>
</div>
{{template "foot" .}}