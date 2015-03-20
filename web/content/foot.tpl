{{define "foot"}}
<div id="x-comand" class="x-comand form-inline">
      <div class="form-group">
        <div class="btn-group dropup">
          <button type="button" class="btn btn-default dropdown-toggle" data-toggle="dropdown" aria-expanded="false">
            Module <span class="caret"></span>
          </button>
          <ul class="dropdown-menu" role="menu">
            {{range $tab := .Config.Tabs}}
              <li><a href="/{{$tab.Handle}}">{{$tab.Title}}</a></li>
            {{end}}
          </ul>
        </div>
        <div class="input-group">
          <div class="input-group-addon">Comand:</div>
          <input type="text" placeholder="Comand" id="x-msg" class="form-control" style="width:300px;">
        </div>
      	<button id="x-exec" class="button button-caution" type="button">Execute</button>
      </div>
</div>
</body>
<script type="text/javascript" src="js/require.js" data-main="js/config"></script>
</html>
{{end}}