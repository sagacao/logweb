$(function(){

    indexModel.onlyZeuidShow();
    //初始化fileinput
    $("#xlsfile").fileinput({
        language:"zh",
        browseClass: "btn btn-info",
    });
    
    $(".fileinput-upload-button").on("click", function () {
        upload()
        event.preventDefault() //阻止form表单默认提交
    })
    
});

function upload() {
    var formData = new FormData($("#uploadxls")[0])
    formData.set("zeusid", getServerId)
    
    $.ajax({
        type: 'POST',
        url: '/importfilter/import',
        contentType: false,    //不设置content-Type请求头
        processData: false,    //不处理发送的数据
        data: formData
    }).success(function (message) {
        console.log('message:', message)
    })
}
