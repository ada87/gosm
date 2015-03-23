define(['socket'],function(socket){

	var method= 'encode';
	var type = 'md5';
	var el_type=$('.rowtype div');
	var el_method=$('.rowmethod div');
	var txt_code = $('#strencode');
	var txt_key = $('#strkey');
	var div_output = $('.output');


	var sec = ['aes','des'];

	el_type.click(function(){
		el_type.removeClass('selected');
		var el=$(this);
		el.addClass('selected');
		method = el.attr('etype');
	});
	el_method.click(function(){
		el_method.removeClass('selected');
		var el=$(this);
		el.addClass('selected');
		type = el.text();
		if(sec.indexOf(type)!=-1){
			txt_key.show();
		}else{
			txt_key.hide();
		}
	});

	var fnCoder = function(){
		var txt = $.trim(txt_code.val());
		if(txt!=''){
			var data = {
				method:method,
				waygo:type,
				code:txt
			}
			if(sec.indexOf(type)!=-1){
				var key = $.trim(txt_key.val());
				if(key ==''){
					receive({result:type+' ：必须输入密钥'});
					return;
				}
				data.key = key;
			}else{
				data.key = '';
			}
			socket.sendMessage(data)
		}
	}
	txt_code.keyup(fnCoder);
	txt_key.keyup(fnCoder);


	function receive(data){
		var txt = type+' ( '+$.trim(txt_code.val())+' ) = ' + data.result; 
		div_output.text(txt);
	}

	function init(){
		socket.bindReceive(receive)
	}

	return {init:init}
})
