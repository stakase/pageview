var param = "path" + location.pathname;
//window.alert(param);

var param2 = "?keyword=" + location.pathname;
//location.href="http://18.221.113.158:8080/pageview" + param2;

var url = "http://18.221.113.158:8080/pageview" + param2; // リクエスト先URL
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
