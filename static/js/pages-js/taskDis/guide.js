$(function(){
    indexModel.SystemHide();

    inits();

    function inits(){
        var jsonData = {
            startTime: getStartDate(),
            endTime: getEndDate(),
            zeusid: getServerId,
            channelid: getChannelId
        }

        $.ajax({
            url:"/static/test.json",
            dataType:"json",
            data:jsonData,
            error:function() {},
            success:function(data){
                $("#chart-content").highcharts({
                    chart: {
                        type: "column",
                        zoomType: "x",
                    },
                    title: {
                        text: "指引任务"
                    },
                    xAxis: {
                        title:{
                            text:"指引编号"
                        },
                        crosshair: true
                    },
                    yAxis: {
                        title: {
                            text: "指引任务分布"
                        }
                    },
                    series: [{
                        name: "指引任务分布",
                        data: data.test02.data,
                        color:"#439fae",
                        dataLabels: {
                            enabled: true,
                            color: "#000"
                        }
                    }]
                });
            }
        });
        
    }
    
});