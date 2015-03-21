define(['os','regexp','crypto'],function(os,regexp,crypto){

	function resize(){
		$('.fullScreen').height(document.documentElement.clientHeight);
	}
	$(window).resize(resize);
	resize();
	var page = window.location.pathname.split('/')[1];
	$('.dropup button').html((page?page:'月光宝盒') + '<span class="caret"></span>');
	switch (page){
		case 'os':
		os.init();
		break;
		case 'regexp':
		regexp.init();
		break;
		case 'crypto':
		crypto.init();
		break;
		default:
		break
	}
})