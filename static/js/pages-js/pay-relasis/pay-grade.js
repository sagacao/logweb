$(function(){

    indexModel.systemChannelHide();

    var payGradeViewModel = function () {
    }

    var newGradeModel = new payGradeViewModel();
    ko.applyBindings(newGradeModel, document.getElementById("payGrade"));

    var jsonData = { startTime: getStartDate(), endTime: getEndDate(), pageSize: 10 };
    
    var initTable = function (pageclickednumber) {
        var index = layer.load(1,{
            shade: [0.3, '#666666'],
            content: "数据获取中......"
        });

        jsonData.pageIndex = pageclickednumber;
        getServerId ? $.extend(jsonData, { zeusid: getServerId }) : $.extend(jsonData, {});

        $.ajax({
            type: "get",
            async: true,
            url: "/pay/payleveldata",
            data: jsonData,
            dataType: "json",
            error: function (xml,status) {
                layer.close(index);
                console.log("error:"+status);
            },
            success: function (data) {
                console.log(jsonData);
                console.log("pay-grade:",data);
                layer.close(index);
                if (data.code == 200) {
                    if (data.info.rows) {

                        var xA = [];
                        var arrAmount = [];
                        var arrPayerCount = [];
                        var arrPayCount = [];

                        $.each(data.info.rows, function(i, obj){
                            xA.push("Lv"+obj.level);
                            arrAmount.push(obj.amount);
                            arrPayerCount.push(obj.payerCount);
                            arrPayCount.push(obj.payCount);
                        })
                        
                        $("#chart-content").highcharts({
                            title:{
                                text:"付费用户等级分布"
                            },
                            chart: {
                                type: 'spline'
                            },
                            xAxis:{
                                crosshair: true,
                                title:{
                                    text:"等级"
                                },
                                categories:xA
                            },
                            yAxis:{
                                title:{
                                    text:"数量"
                                }
                            },
                            tooltip: {
                                shared: true,         //是否启动提示框共享
                            },
                            plotOptions: {
                                spline: {
                                    dataLabels: {
                                        enabled: true          // 开启数据标签
                                    },
                                }
                            },
                            series: [{
                                name:"付费金额",
                                data: arrAmount
                            },{
                                name:"付费角色数",
                                data: arrPayerCount
                            },{
                                name:"付费次数",
                                data: arrPayCount
                            }]
                        });
                    }else{
                        console.log("payleveldata info.rows is no data");
                    }
                   
                }
            }
        });
    }
    initTable(1);

});

