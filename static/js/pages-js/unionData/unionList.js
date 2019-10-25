$(function() {

    indexModel.onlyZeuidShow();

    var unionListViewModel = function(){
        this.dataList           = ko.observableArray();
        this.isExpand           = ko.observable(false);
        this.isShow             = ko.observable(false);
        this.chartTitle         = ko.observable();
        this.unionNumber        = ko.observable();
        this.time               = ko.observable(new Date());
        this.exportExcel        = function () {
            $("#unionTable").table2excel({
                filename: "公会信息",
                exclude:".noExl"        //不导出的
            })
        };
        this.preDay             = function(){
            this.time(addDays(this.time(), -1));
            initCharts();
        };
        this.nextDay            = function() {
            this.time(addDays(this.time(), 1));
            initCharts();
        }
    }
    var newUnionListModel = new unionListViewModel();
    ko.applyBindings(newUnionListModel, $("#unionList").get(0));

    inits();
    
    function inits(){
        var jsonData = {
            zeusid: getServerId,
            startTime: getStartDate(),
            endTime: getEndDate()
        }
        
        $.ajax({
            url:"/league/leaguedata",
            data:jsonData,
            dataType:"json",
            error:function(){

            },
            success:function(data){
                console.log(jsonData);
                console.log("leaguedata: ",data);

                newUnionListModel.dataList(data.info.rows);
                
                $("#unionTable td").on("click", "button", function () {
                    newUnionListModel.isShow(true);
                    $("#contentMain").slideUp("slow");
                    newUnionListModel.isExpand(true);

                    // indexModel.systemChannelHide();
                    // indexModel.onlyTime(false);
                    
                    newUnionListModel.chartTitle($(this).parent().siblings().eq(1).text());      //chartTitle
                    newUnionListModel.unionNumber($(this).parent().siblings().eq(0).text());
                    //unionNumber
                    
                    $(".expandTitle").on("click", function () {
                        $("#contentMain").slideDown("slow");
                        // newUnionListModel.isExpand(false);
                    });

                    initCharts();
                });
            }
        })
        
    }
    
    function initCharts(){
        var jsonData = {
            zeusid: getServerId,
            startTime: newUnionListModel.time().format("yyyy-MM-dd"),
            number: newUnionListModel.unionNumber(),
        }

        $.ajax({
            url:"/league/leagueinfo",
            data:jsonData,
            dataType:"json",
            error:function(){},
            success:function(data){
                console.log(jsonData);
                console.log("leagueinfo: ", data);
              
                var timeArr = newUnionListModel.time().format("yyyy-MM-dd").split("-");

                $("#chart-content").highcharts({
                    title: {
                        text: "公会名字：" + newUnionListModel.chartTitle() + " #公会编号：" + newUnionListModel.unionNumber() + " (" + newUnionListModel.time().format("yyyy-MM-dd") + ")"
                    },
                    chart: {
                        type: 'spline'
                    },
                    xAxis: {
                        crosshair: true,
                        type: 'datetime',
                        labels: {
                            overflow: 'justify'
                        },
                        title: {
                            text: "时间（单位：小时）"
                        }
                    },
                    yAxis: {
                        title: {
                            text: "数量"
                        },
                        labels: {
                            formatter: function() {
                                if (this.value/1000 > 0.1) {
                                    return this.value / 1000 + "k";
                                } else {
                                    return this.value;
                                }
                            }
                        }
                    },
                    tooltip: {
                        shared: true,         //是否启动提示框共享
                    },
                    plotOptions: {
                        spline: {
                            lineWidth: 2,
                            states: {
                                hover: {
                                    lineWidth: 3
                                }
                            },
                            marker: {
                                enabled: false
                            },
                            pointInterval: 900000,  // 15 min(时间间隔)
                            pointStart: Date.UTC(timeArr[0], timeArr[1] - 1, timeArr[2], 0, 0, 0)
                        }
                    },
                    series: [{
                        name: "公会等级",
                        data: data.info.rows[0].level,
                    }, {
                        name: "公会资产",
                        data: data.info.rows[0].assets,
                    }, {
                        name: "公会积分",
                        data: data.info.rows[0].credits,
                    }, {
                        name: "公会健康度",
                        data: data.info.rows[0].health,
                    }, {
                        name: "公会成员",
                        data: data.info.rows[0].member,
                    }, {
                        name: "公会学徒",
                        data: data.info.rows[0].student,
                    }]
                });
            }
        })
    }
    
});