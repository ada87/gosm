define(['socket'],function(socket){

	function receive(data){
		console.log(data)
	}

	function init(){
		socket.bindReceive(receive)
	}

	return {init:init}
})
