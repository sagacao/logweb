$(function () {
    indexModel.systemChannelHide();

    inits();  
   
    function inits() {
        var jsonData = {
            startTime: getStartDate(),
            endTime: getStartDate()
        }
        getServerId ? $.extend(jsonData, { zeusid: getServerId }) : $.extend(jsonData, {});
        
        $.ajax({
            url:"/statistics/daudata",
            dataType:"json",
            data:jsonData,
            error:function(){
                spinner.spin();
            },
            beforeSend: function () {
                spinner.spin($("#chart-content").get(0));
            },
            success:function(data){
                console.log(jsonData);
                console.log(data);

                var serData = [];
                if (data.info.rows) {
                    for (v of data.info.rows) {
                        if (v.dimDay == jsonData.startTime) {
                            serData.push(v.dau)
                            serData.push(v.roleDau)
                            serData.push(v.macDau)
                        }
                    }
                }else{
                    console.log("dau info.rows is no data!");
                }
                
                spinner.spin();
                $("#chart-content").highcharts({
                    chart: {
                        type: "column"
                    },
                    title: {
                        text: "日活跃" + "(" + jsonData.startTime.format("yyyy-MM-dd") + ")"
                    },
                    xAxis: {
                        categories:["活跃账号","活跃角色","活跃设备"],
                        crosshair: true
                    },
                    yAxis: {
                        title: {
                            text: "人数"
                        }
                    },
                    plotOptions: {
                        series: {
                            dataLabels: {
                                enabled: true
                            }
                        }
                    },
                    series: [{
                        name: '日活跃',
                        data: serData,
                        color:'#439fae'
                    }]
                });
            }
        });
        
    }

});