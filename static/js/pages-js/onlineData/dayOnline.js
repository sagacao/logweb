$(function(){
    indexModel.systemChannelHide();
    
    inits();
    
    function inits() {

        var jsonData = {
            startTime: getStartDate(),
            zeusid: getServerId
        }

        $.ajax({
            url:"/onlinedata/online",
            dataType:"json",
            data:jsonData,
            error:function(){},
            success:function(data){
                console.log(jsonData);
                console.log("dayOnline:", data);

                var seriesData = [];

                if (data.info.rows) {
                    data.info.rows.sort(function (a, b) {     //按照time进行排序
                        return a.time - b.time;
                    })

                    for (var i of data.info.rows) {
                        seriesData.push(i.rolenum);
                    }

                    var timeArr = jsonData.startTime.split("-");
                } else {
                    console.log("dayOnline info.rows is no data.");
                }
                
                $("#chart-content").highcharts({
                    chart:{
                        type:"area",
                        zoomType:"x"
                    },
                    title:{
                        text:"在线"
                    },
                    xAxis:{
                        crosshair: true,
                        type:"datetime",
                        title:{
                            text:"时间（单位：小时）"
                        }
                    },
                    yAxis:{
                        title:{
                            text:"实时在线人数（单位：人）"
                        }
                    },
                    plotOptions: {
                        area: {
                            pointStart: Date.UTC(timeArr[0], timeArr[1]-1, timeArr[2], 0, 0, 0),
                            pointInterval: 1000*60*5  //时间间隔 5 min
                        }
                    },
                    series:[{
                        name:"实时在线",
                        data: seriesData,
                        color:"#439fae"
                    }]
                });
            }
        });
        
    }

});