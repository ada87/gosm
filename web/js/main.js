define(['os'],function(os){

	 // // Create a socket
	 //  var socket = new WebSocket('ws://'+window.location.host+'/websocket/room/socket?user={{.user}}')
	 //  // Display a message
	 //  var display = function(event) {
	 //    $('#thread').append(tmpl('message_tmpl', {event: event}));
	 //    $('#thread').scrollTo('max')
	 //  }
	 //  // Message received on the socket
	 //  socket.onmessage = function(event) {
	 //    display(JSON.parse(event.data))
	 //  }
	 //  $('#send').click(function(e) {
	 //    var message = $('#message').val()
	 //    $('#message').val('')
	 //    socket.send(message)
	 //  });
	 //  $('#message').keypress(function(e) {
	 //    if(e.charCode == 13 || e.keyCode == 13) {
	 //      $('#send').click()
	 //      e.preventDefault()
	 //    }
	 //  })
	var page = window.location.pathname.split('/')
	switch (page[1]){
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