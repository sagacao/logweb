$(function(){
    indexModel.allHide();
    
    var logOnlineSummaryViewModel = function(){
        this.dataList       = ko.observableArray();
        this.rolenumSum     = ko.observable(0);
        this.exportExcel    = function () {
            $("#onlineTable").table2excel({
                filename:"当前在线汇总"
            });
        }
    }
    var newLogOnlineSummaryModel = new logOnlineSummaryViewModel();
    ko.applyBindings(newLogOnlineSummaryModel, $("#onlineSummary").get(0));
        
    inits();
    
    function inits(){

        $.ajax({
            url:"/onlinedata/curonline",
            dataType:"json",
            error:function(){},
            success:function (data) {
                console.log(data);
                if (data.info.rows) {
                    newLogOnlineSummaryModel.dataList(data.info.rows);

                    //汇总信息
                    var sum = 0;
                    $.each(data.info.rows, function(i, obj){
                        sum += obj.rolenum;
                    });
                    newLogOnlineSummaryModel.rolenumSum(sum);

                } else {
                    console.log("onlineSummary info.rows is no data.")
                }
            }
        });
    }
    
});