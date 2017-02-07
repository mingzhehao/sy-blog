$(function(){
    $('.content').fadeIn('fast');
    if($('.board').length>3){
        $('.content').masonry({
            itemSelector : '.board'
            ,isFitWidth:true
        });
    }else{
        $('.board').css({float:'left'});
        $('.board').parents('.content').addClass('clearfix').css({width:'850px'})
    }

    var rgba = [
        'rgba(251,34,240,0.25)'
        ,'rgba(214,17,21,0.25)'
        ,'rgba(14,251,252,0.25)'
        ,'rgba(158,134,255,0.25)'
        ,'rgba(60,255,20,0.25)'
        ,'rgba(44,158,52,0.25)'
        ,'rgba(225,211,20,0.25)'
        ,'rgba(100,117,121,0.25)'
    ];

    $.each($('.board'),function(index,item){
        var charCode = location.pathname.substr(1).charCodeAt()
        var i = (index+charCode) % rgba.length;
        $(item).css('background',rgba[i]);
    });

    $('.board').delegate('a','click',function(e){
        e.preventDefault()
        var target = e.target;
        if(!$(this).parent().hasClass('inactive')){
            window.open($(target).attr('href'));
            mixpanel.track("Cheat Link",{
                'pagename':location.pathname
                ,'href':$(target).attr('href')
            });
        }
    });

    var shareInputFocus = false;
    $('body').delegate('#at16filt','focus',function(){
        shareInputFocus = true;
    }).delegate('#at16filt','blur',function(){
        shareInputFocus = false;
    });

    //$('body').bind('keydown',function (e) {
        //if (!shareInputFocus && !e.metaKey && !e.shiftKey && !e.ctrlKey && !e.altKey && e.keyCode != 27 && e.keyCode!= 32 && e.keyCode!=33 && e.keyCode!=34 && !$('#searchApi').is(':focus')) {
            //$('#searchApi').focus();
        //}
    //});

});
