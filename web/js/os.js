define(['socket'],function(socket){

	var ospanel = $('#ospanel');

	var cmd = $('.x-left ul li');

	cmd.click(function(){
		socket.sendMessage($(this).attr('comand'));
	})

	function receive(data){
		ospanel.empty();
		for(i in data){
			var html = '<p>'+i+':' +data[i] +  '</p>';
			ospanel.append(html);
		}

	}

	function init(){
		socket.bindReceive(receive)
	}

	return {init:init}
})
