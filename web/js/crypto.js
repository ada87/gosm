define(['socket'],function(socket){

	var type = 'encode';
	var method = 'md5';
	var el_type=$('.rowtype div');
	var el_method=$('.rowmethod div');
	var txt_code = $('#strencode');
	var div_output = $('.output');

	el_type.click(function(){
		el_type.removeClass('selected');
		var el=$(this);
		el.addClass('selected');
		type = el.attr('etype');
	});
	el_method.click(function(){
		el_method.removeClass('selected');
		var el=$(this);
		el.addClass('selected');
		method = el.text();
	});
	txt_code.keyup(function(){
		var txt = $.trim(txt_code.val());
		if(txt!=''){
			socket.sendMessage(method+' '+type +' '+ txt)
		}
	})

	function receive(data){
		var txt = method+' ( '+$.trim(txt_code.val())+' ) = ' + data.result; 
		div_output.text(txt);
	}

	function init(){
		socket.bindReceive(receive)
	}

	return {init:init}
})
