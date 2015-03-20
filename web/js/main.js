define(['os'],function(os){

	function resize(){
		$('.fullScreen').height(document.documentElement.clientHeight-45);
	}
	$(window).resize(resize);
	resize();
	var page = window.location.pathname.split('/')[1];
	$('.dropup button').html(page?page:'月光宝盒'+'<span class="caret"></span>');
	switch (page){
		case 'os':
		os.init();
		break;
		case 'regexp':
		regexp.init();
		break;
		default:
		break
	}
})