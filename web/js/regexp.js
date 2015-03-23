define(['jquery'],function($){

	var btn_edit = $('#edit');
	var btn_back = $('#back');
	var btn_test = $('#test');

	var checkItems = $('li.item p.check');

	function test(){
		var regstr = $.trim($('#txtreg').val());
		if(regstr==''){
			checkItems.each(function(i,o){
				o.className='check';
			});
			return;
		}
		var reg = new RegExp(regstr, 'gm');
		checkItems.each(function(i,o){
			if(reg.test(o.textContent)){
				o.className='check checky';
			}else{
				o.className='check checkn';
			}
		});
	}


	function init(){
		btn_edit.click(function(){
			$('.reg9g').addClass('toleft');
		});
		btn_back.click(function(){
			$('.reg9g').removeClass('toleft');
		});
		btn_test.click(test);

		$('.editpanel ul li a').click(function (e) {
		  e.preventDefault();
		  $(this).tab('show');
		})
	}

	return {init:init}
})
