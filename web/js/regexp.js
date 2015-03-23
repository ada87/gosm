define(['jquery'],function($){

	var btn_edit = $('#edit');
	var btn_done = $('#done');

	function gotoEdit(e){
		$('.reg9g').addClass('toleft');
	}

	function init(){
		btn_edit.click(gotoEdit)
	}

	return {init:init}
})
