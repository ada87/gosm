define(['socket'],function(socket){

	var btn_edit = $('#edit');
	var btn_back = $('#back');
	var btn_test = $('#test');

	// var btn_update = $('.btnupdate');
	// var btn_done = $('.btndone');
	var btn_new = $('.btnnew');

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

	function receive(data){
		console.log(data);
	}
	function insert(fid,fval,fdes){
		var data ={
			act:"new",
			fid:fid,
			fval:fval,
			fdes:fdes
		}
		socket.sendMessage(data);
	}
	function update(vid,fval,fdes){
		var data ={
			act:"update",
			vid:vid,
			fval:fval,
			fdes:fdes
		}
		socket.sendMessage(data);
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
		});

		$('.tab-content').on('click','.btnupdate',function(){
			var el = $(this);
			$('.row').removeClass('onedit');
			var root = el.parent().parent();
			root.addClass('onedit');
			root.children().attr('contenteditable',true);
			return false;
		});

		$('.tab-content').on('click','.btndone',function(){
			var el = $(this);
			$('.row').removeClass('onedit');
			var root = el.parent().parent();
			var attrs = root.children();
			var fval = attrs[0].textContent;
			var fdesc = attrs[1].textContent;
			var vid = root.attr('vid');
			if(vid){
				update(vid, fval, fdesc);
			}else{
				var fid = root.attr('fid');
				insert(fid, fval, fdesc);
			}
			attrs.attr('contenteditable',false);
			return false;
		});
		btn_new.click(function(){
			var el = $(this);
			var fid = el.attr('fid');
			$('.row').removeClass('onedit');
			var html = '<div fid="' + fid + '" class="row onedit">'
				html += '<div contenteditable="true" class="col-xs-4 attr"></div>';
				html += '<div contenteditable="true" class="col-xs-6 attr"></div>';
				html += '<div contenteditable="true" class="col-xs-2 attr">';
				html += '	<a fid="' + fid + '" class="btnupdate">Update</a>';
				html += '	<a fid="' + fid + '" class="btndone">Done</a>';
				html += '</div></div>';
			el.before(html);
		});
		socket.bindReceive(receive)
	}


	return {init:init}
})
