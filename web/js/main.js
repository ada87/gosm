define(['os','regexp','crypto'],function(os,regexp,crypto){

	var comand = $('#x-comand');
	var page = window.location.pathname.split('/')[1];
	function resize(){
		$('.fullScreen').height(document.documentElement.clientHeight);
	}
	function hideCommand(){
		comand.addClass('x-comand-min');
	}
	$(window).resize(resize);
	resize();
	$('.dropup button').html((page?page:'月光宝盒') + '<span class="caret"></span>');
	switch (page){
		case 'os':
			os.init();
		break;
		case 'regexp':
			hideCommand();
			regexp.init();
		break;
		case 'crypto':
			crypto.init();
		break;
		default:
		break
	}
})