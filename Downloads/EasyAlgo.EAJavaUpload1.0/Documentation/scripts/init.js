var windowProxy;
    $(document).ready(function () {
 				
		$("a").each(function () {
            try {
                if ($(this).attr("href").indexOf("#") > 0) {
                    $(this).click(function (e) {
                        e.preventDefault();
						windowProxy.postMessage($(this).attr("href"));                        
                    });
                }
            }
            catch (ex) { }
        });
		
		windowProxy = new Porthole.WindowProxy('../proxy.html');
		$(window)._scrollable();
		if(window.location.hash != "")
		{
			$.scrollTo('a[name="' + window.location.hash.replace("#", "") + '"]');			
		}
    });