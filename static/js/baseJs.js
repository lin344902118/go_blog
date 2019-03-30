function ajax(type, url, datas, msg) {
    $.ajax({
        type: type,
        url: url,
        data: datas,
        async: false,
        success: function (data, status) {
            if (data.ret == 0) {
　　　　　　　　location.reload()

            }
            else {
                alert(msg+"失败!失败原因:" + data.message);
            }
        },
        error: function (data) {
            alert(msg+"错误!错误原因:" + data);
        }
    });
}