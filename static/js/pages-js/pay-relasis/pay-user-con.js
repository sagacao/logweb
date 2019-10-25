$(function(){

    indexModel.systemChannelHide();

    var userConViewModel = function(){              //创建视图模型
        this.dataList       = ko.observableArray();
        this.exportExcel    = function () {
            $("#userconTable").table2excel({
                filename: "用户贡献"
            });
        }
    };
    var newUserConModel  = new userConViewModel();
    ko.applyBindings( newUserConModel,document.getElementById("userCon") );

    var initData = function (){
        var jsonData = {
            startTime: getStartDate(),
            endTime: getEndDate()
        }
        getServerId ? $.extend(jsonData, { zeusid: getServerId }) : $.extend(jsonData, {});

        console.log("chart jsonData:", jsonData);
        var timeArr = jsonData.startTime.split("-");

        $.ajax({
            url:"/pay/payltv",
            dataType:"json",
            data: jsonData,
            error: function(){},
            success: function(data){
                console.log(data);
                if (data.info.rows) {

                    // chart
                    var pay3 = [];  // 3日付费总金额
                    var pay7 = [];  // 7日付费总金额
                    var pay14 = []; // 14日付费总金额
                    var pay30 = []; // 30日付费总金额
                    var dnp3 = [];  // 3日活跃人数
                    var dnp7 = [];  // 7日活跃人数
                    var dnp14 = []; // 14日活跃人数
                    var dnp30 = []; // 30日活跃人数
                    
                    for (let i = 0, len = data.info.rows.length; i < len; i++) {
                        pay3.push(data.info.rows[i].pay3);
                        pay7.push(data.info.rows[i].pay7);
                        pay14.push(data.info.rows[i].pay14);
                        pay30.push(data.info.rows[i].pay30);
                        dnp3.push(data.info.rows[i].dnp3);
                        dnp7.push(data.info.rows[i].dnp7);
                        dnp14.push(data.info.rows[i].dnp14);
                        dnp30.push(data.info.rows[i].dnp30);
                    }
                    
                    var options = {
                        chart: {
                            type: "spline",
                            zoomType: "x"
                        },
                        title: {
                            text: "LTV"
                        },
                        xAxis: {
                            title: {
                                text: "时间"
                            },
                            type: "datetime",
                            crosshair: true
                        },
                        yAxis: {
                            title: {
                                text: "付费金额"
                            }
                        },
                        tooltip: {
                            shared: true
                        },
                        plotOptions: {
                            spline: {
                                dataLabels: {
                                    enabled: true
                                },
                                pointStart: Date.UTC(timeArr[0], timeArr[1] - 1, timeArr[2], 0, 0, 0),
                                pointInterval: 1000 * 60 * 60 * 24
                            }
                        },
                        series: [{
                            name:"3日付费总金额",
                            data: pay3
                        },{
                            name: "7日付费总金额",
                            data: pay7
                        },{
                            name: "14日付费总金额",
                            data: pay14
                        },{
                            name: "30日付费总金额",
                            data: pay30
                        },{
                            name: "3日ltv",
                            data: dnp3
                        },{
                            name:"7日ltv",
                            data: dnp7
                        },{
                            name:"14日ltv",
                            data: dnp14
                        },{
                            name:"30日ltv",
                            data: dnp30
                        }]
                    }
                    Highcharts.chart("chart-spline", options);

                    // table
                    var tabDom = $("#user_con").empty();
                    newUserConModel.dataList.removeAll();

                    newUserConModel.dataList(data.info.rows);
                    
                } else {
                    console.log("pay-user-con data.info.rows is no data.");
                }
            }
        })
        
    }

    initData();


});