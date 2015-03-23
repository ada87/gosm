/**
* WebSocket Api
*/
define([],function(){
	var receive = null,path=window.location.pathname.split('/')[1];
	var socket = new WebSocket('ws://'+window.location.host+'/_');
	socket.onmessage = function(event) {
		var data = JSON.parse(event.data);
		receive.call(null,data);
	}
	function sendMessage(cmd){
		if(receive!=null){
			var data = {
				path:path
			}
			if(typeof cmd == 'string'){
				data.cmd = cmd;
			}else{
				for( i in cmd){
					data[i] = cmd[i];
				}
			}
	    	socket.send(JSON.stringify(data));
		}else{
			alert('not vaild');
		}

	}
	$('#x-comand #x-exec').click(function(e) {
	    var message = $('#x-msg').val()
	    $('#x-msg').val('');
	    sendMessage(message);
	});
	$('#x-msg').keypress(function(e) {
		if(e.charCode == 13 || e.keyCode == 13) {
			$('#x-comand #x-exec').click();
			e.preventDefault();
		}
	});
	return {
		sendMessage:sendMessage,
		bindReceive:function(fn){
			receive = fn;
		}
	}
})
