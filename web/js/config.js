require.config({
	baseUrl: 'js/',
    paths: {
        jquery: 'jquery-2.1.3.min',
        bootstrap:'bootstrap.min'
    },
    shim: {
    	'bootstrap':['jquery']
    }
});
require(['main','jquery','bootstrap'],function(){
});
