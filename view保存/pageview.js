var param = "?keyword=" + location.pathname;

var url = "http://XXXXXXXXX/pageview" + param;

var request = new XMLHttpRequest();
request.open('GET', url);
request.onreadystatechange = function () {
    if (request.readyState != 4) {
        // リクエスト中
    } else if (request.status != 200) {
        // 失敗
    } else {
        // 取得成功
        // var result = request.responseText;
    }
};
request.send(null);
