var indexViewModel = function(){
    var self                = this;
    this.typeOfSystem       = ko.observable(true); 
    this.filterByDate       = ko.observable(false); 
    this.filterByChannel    = ko.observable(false); 
    this.filterByZeuid      = ko.observable(false); 
    this.onlyTime           = ko.observable(true);
    
    //标签栏的各种状态
    this.allShow            = function () {
        self.typeOfSystem(true);
        self.filterByDate(true);
        self.filterByChannel(true);
        self.filterByZeuid(true);
    };
    this.allHide            = function(){
        self.typeOfSystem(false);
        self.filterByDate(false);
        self.filterByChannel(false);
        self.filterByZeuid(false);
    }
    this.channelZeuidHide   = function () {
        self.typeOfSystem(true);
        self.filterByDate(true);
        self.filterByChannel(false);
        self.filterByZeuid(false);
    };
    this.zeuidHide          = function(){
        self.typeOfSystem(true);
        self.filterByDate(true);
        self.filterByChannel(true);
        self.filterByZeuid(false);
    };
    this.channelHide = function () {
        self.typeOfSystem(true);
        self.filterByDate(true);
        self.filterByChannel(false);
        self.filterByZeuid(true);
    };
    this.dataChannelHide    = function(){
        self.typeOfSystem(true);
        self.filterByDate(false);
        self.filterByChannel(false);
        self.filterByZeuid(true);
    };
    this.systemChannelHide  = function(){
        self.typeOfSystem(false);
        self.filterByDate(true);
        self.filterByChannel(false);
        self.filterByZeuid(true);
    };
    this.onlyZeuidShow      = function(){
        self.typeOfSystem(false);
        self.filterByDate(false);
        self.filterByChannel(false);
        self.filterByZeuid(true);
    };
    this.onlySystemShow     = function(){
        self.typeOfSystem(true);
        self.filterByDate(false);
        self.filterByChannel(false);
        self.filterByZeuid(false);
    }
    this.onlyTimeShow     = function () {
        self.typeOfSystem(false);
        self.filterByDate(true);
        self.filterByChannel(false);
        self.filterByZeuid(false);
    }
    this.SystemHide         = function(){
        self.typeOfSystem(false);
        self.filterByDate(true);
        self.filterByChannel(true);
        self.filterByZeuid(true);
    }
    this.onlyTimeFn         = function () {
        this.onlyTime(false);
    }.bind(this);
    
    //configUrl.实时数据
    this.configUrl = {
        "付费等级": {
            url: "/static/pages-html/pay-relasis/pay-grade.html",
        },
        "用户贡献": {
            url: "/static/pages-html/pay-relasis/pay-user-con.html",
        },
        "付费习惯": {
            url: "/static/pages-html/pay-relasis/pay-habit.html",
            initTab: 1
        },
        "付费排行": {
            url: "/static/pages-html/pay-relasis/pay-rank.html",
            initTab: 1
        },
        "VIP分布": {
            url: "/static/pages-html/pay-relasis/pay-vip-spread.html",
        },
        "付费用户": {
            url: "/static/pages-html/pay-relasis/pay-user.html",
            initTab: 1
        },
        "付费流失": {
            url: "/static/pages-html/pay-relasis/pay-loss.html",
            initTab: 1
        },
        "账号登录日志":{
            url:"/static/pages-html/log-sum/log-account/log-accountIn.html"
        },
        "账号登出日志":{
            url:"/static/pages-html/log-sum/log-account/log-accountOut.html"
        },
        "角色登录日志":{
            url:"/static/pages-html/log-sum/log-role/log-roleIn.html"
        },
        "角色登出日志":{
            url:"/static/pages-html/log-sum/log-role/log-roleOut.html"
        },
        "角色升级日志":{
            url:"/static/pages-html/log-sum/log-role/log-roleUp.html"
        },
        "邮件发送日志":{
            url:"/static/pages-html/log-sum/log-mail/log-mailSend.html"
        },
        "邮件删除日志":{
            url:"/static/pages-html/log-sum/log-mail/log-mailDel.html"
        },
        "商城日志":{
            url:"/static/pages-html/log-sum/log-mall/log-mallSales.html"
        },
        "道具上架":{
            url: "/static/pages-html/log-sum/log-auction/log-propsOn.html"
        },
        "道具下架":{
            url:"/static/pages-html/log-sum/log-auction/log-propsOut.html"
        },
        "拍卖成交":{
            url:"/static/pages-html/log-sum/log-auction/log-auctionSuccess.html"
        },
        "任务接取日志":{
            url:"/static/pages-html/log-sum/log-task/log-acceptTask.html"
        },
        "任务完成日志":{
            url:"/static/pages-html/log-sum/log-task/log-completeTask.html"
        },
        "任务放弃日志":{
            url:"/static/pages-html/log-sum/log-task/log-giveUpTask.html"
        },
        "任务失败日志":{
            url:"/static/pages-html/log-sum/log-task/log-failTask.html"
        },
        "创建公会日志": {
            url: "/static/pages-html/log-sum/log-union/log-createUnion.html",
        },
        "解散公会日志": {
            url: "/static/pages-html/log-sum/log-union/log-dissolutionOfUnion.html",
        },
        "加入公会日志": {
            url: "/static/pages-html/log-sum/log-union/log-joinUnion.html",
        },
        "离开公会日志": {
            url: "/static/pages-html/log-sum/log-union/log-leaveUnion.html",
        },
        "转让公会日志": {
            url: "/static/pages-html/log-sum/log-union/log-transferUnion.html",
        },
        "公会信息": {
            url: "/static/pages-html/log-sum/log-union/log-messageOfUnion.html",
        },
        "公会成员": {
            url: "/static/pages-html/log-sum/log-union/log-memberMsg.html",
        },
        "物品获取": {
            url: "/static/pages-html/log-sum/log-items/log-getItems.html",
        },
        "物品丢失": {
            url: "/static/pages-html/log-sum/log-items/log-loseItems.html",
        },
        "虚拟货币": {
            url: "/static/pages-html/log-sum/log-money/log-money.html",
        },
        "侠客道具": {
            url: "/static/pages-html/log-sum/log-money/log-partner.html",
        },
        "查询原始日志": {
            url: "/static/pages-html/log-sum/log-raw/log-raw.html",
        },
        "查询分类日志": {
            url: "/static/pages-html/log-sum/log-raw/log-format.html",
        },
        "侠客背包日志": {
            url: "/static/pages-html/log-sum/log-partnerBag/log-partnerBag.html",
            initTab: 1,
        },
        "充值订单日志": {
            url: "/static/pages-html/log-sum/log-recharge/log-recharge.html",
        },
        "留存汇总":{
            url:"/static/pages-html/register/keepSummary.html",
            callback: this.onlyTimeFn
        },
        "渠道留存":{
            url:"/static/pages-html/register/keepChannel.html",
            callback: this.onlyTimeFn
        },
        "每日活跃":{
            url:"/static/pages-html/register/dau.html",
            callback: this.onlyTimeFn
        },
        "平台活跃":{
            url:"/static/pages-html/register/pau.html",
            callback: this.onlyTimeFn
        },
        "流失数据":{
            url:"/static/pages-html/register/loseData.html",
            callback: this.onlyTimeFn
        },
        "等级分布":{
            url:"/static/pages-html/level/levelDis.html"
        },
        "等级流失":{
            url:"/static/pages-html/level/levelLose.html"
        },
        "等级排行":{
            url:"/static/pages-html/level/levelRank.html"
        },
        "每日在线":{
            url:"/static/pages-html/onlineData/dayOnline.html",
            callback: this.onlyTimeFn
        },
        "当前在线汇总":{
            url:"/static/pages-html/onlineData/onlineSummary.html"
        },
        "日平均在线时长":{
            url:"/static/pages-html/onlineData/dayAverageOnline.html"
        },
        "公会统计":{
            url:"/static/pages-html/onlineData/union.html",
            callback: this.onlyTimeFn
        },
        "完成任务":{
            url:"/static/pages-html/taskDis/finishTask.html"
        },
        "新手指引":{
            url:"/static/pages-html/taskDis/guide.html"
        },
        "公会列表":{
            url:"/static/pages-html/unionData/unionList.html"
        },
        "公会活动":{
            url:"/static/pages-html/unionData/unionActive.html",
            callback: this.onlyTimeFn
        },
        "货币产出/消耗":{
            url:"/static/pages-html/coinUse/coin.html",
            initTab: 1
        },
        "商城货币":{
            url:"/static/pages-html/coinUse/mall.html",
            initTab: 1
        },
        "货币排行":{
            url:"/static/pages-html/coinUse/virCoin.html",
            initTab: 1
        },
        "充值":{
            url:"/static/pages-html/pay-relasis/recharge.html",
            initTab: 1
        },
        "每小时充值": {
            url: "/static/pages-html/pay-relasis/hour-recharge.html",
            callback: this.onlyTimeFn
        },
        "账号过滤":{
            url:"/static/pages-html/dataManage/accountFilter.html"
        }

    };
    //渠道值
    this.channelIdChange    = ko.observable("");   
    this.channelidList      = ko.observableArray();
    this.channelIdChangeFn  = ko.computed(function () {
        if (this.channelIdChange() == "") {
            getChannelId = '';                  //渠道默认值清除
            $("#areaClothing").removeAttr("data-id");
        }
    }, this);
    this.getChannelValue    = function (data, event) {     
        var idCurrent = $(event.currentTarget).attr("id");
        $(event.currentTarget).parents("ul").siblings("input").attr("data-id", idCurrent);
        indexModel.channelIdChange($(event.currentTarget).children("a").text());
        $("#channelGo").trigger("click");
    };
    this.channelGo          = function (data, event) {
        if (indexModel.channelIdChange() == "") {
            getChannelId = '';                  //渠道默认值清除
            reload();   
        } else {
            getChannelId = $(event.currentTarget).prev("input").attr("data-id");
            reload();
        };
    };
    //区服值
    this.zeusidChange       = ko.observable("");    
    this.zeuidList          = ko.observableArray();
   
    this.zeusidChangeFn     = ko.computed(function () {
        if (self.zeusidChange() == "") {
            getServerId = '';
            $("#areaClothing").removeAttr("data-id");
        }
    }, this);
    this.getServerValue     = function (data, event) {
        var idCurrent = $(event.currentTarget).attr("id");
        $(event.currentTarget).parents("ul").siblings("input").attr("data-id", idCurrent);
        self.zeusidChange($(event.currentTarget).children("a").text());
        $("#serverGo").trigger("click");
    };
    this.serverGo           = function (data, event) {
        if (self.zeusidChange() == "") {
            getServerId = '';
            reload();
        } else {
            getServerId = $(event.currentTarget).prev("input").attr("data-id");
            reload();
        };
    };
    /* self.showZeus           = ko.computed(function(){           //获取区服列表
        $.ajax("global/config").done(function(data){
            console.log("config:", data);
            
            if (data.code == 200 && data.info.zeusTypeMap) {
                indexModel.zeuidList.removeAll();
                for (i in data.info.zeusTypeMap) {
                    self.zeuidList.push({ zeusid: i, name: data.info.zeusTypeMap[i] })
                    zeusidsGlobal = data.info.zeusTypeMap;
                }
            }
        }).fail(function (xhr, status){
            console.log("showZeus error:" + status);
        }).done(function(){
            $("#cloth_wrap input").attr("data-id", $("#cloth_wrap li").eq(0).attr("id"));       //默认加载列表第一项
            self.zeusidChange($("#cloth_wrap li").eq(0).find("a").text());
            getServerId = $("#cloth_wrap li").eq(0).attr("id");
        });
    },this); */
    
    //SystemLog
    this.iosLogo            = ko.observable(true);
    this.androidLogo        = ko.observable(true);
    this.systemList         = ko.observableArray([{ id: 1, name: "安卓" }, { id: 2, name: "IOS" }, { id: 0, name: "IOS/安卓" }]);
    this.systemLogo         = function (data, event) {
        var id = $(event.currentTarget).attr("id");
        switch (id) {
            case "1":
                self.iosLogo(false);
                self.androidLogo(true);
                break;
            case "2":
                self.iosLogo(true);
                self.androidLogo(false);
                break;
            default:
                self.iosLogo(true);
                self.androidLogo(true);
                break;
        }
    };
    //ztree树配置
    this.userSettingTree    = {          
        view: {							//可视界面相关配置
            dblClickExpand: false,		//双击节点时，是否自动展开父节点的标识
            showLine: true,				 //设置是否显示节点与节点之间的连线
            selectedMulti: false		//设置是否能够同时选中多个节点
        },
        data: {							//数据相关配置
            simpleData: {
                enable: true,			//设置是否启用简单数据格式
                idKey: "id",			//设置启用简单数据格式时id对应的属性名称
                pIdKey: "pId",			//设置启用简单数据格式时parentId对应的属性名称
                rootPId: ""                    
            }
        },
        callback: {
            //beforeClick: 用于捕获单击节点之前的事件回调函数，并且根据返回值确定是否允许单击操作
            beforeClick: function (treeId, treeNode) {
                //treeId: 对应 zTree 的 treeId
                //treeNode: 被单击的节点 JSON 数据对象
                var zTree = $.fn.zTree.getZTreeObj("tree");
                //获取 id 为 tree 的 zTree 对象
                if (treeNode.isParent) {    //查看当前被选中的节点是否是父节点
                    zTree.expandNode(treeNode, null, null, null, true);
                    //expandNode: 展开 / 折叠 指定的节点
                    return false;
                } else {
                    var name = treeNode.name;
                    if (indexModel.configUrl[name]) {

                        $(".pageContainer").attr("data-url",indexModel.configUrl[name].url);
                        $(".pageContainer").load(indexModel.configUrl[name].url+"?number="+Math.random(),function(){
                        });
                        
                        if (indexModel.configUrl[name].callback) {
                            (indexModel.configUrl[name].callback)();    //self.onlyTime(false)
                        } else{
                            self.onlyTime(true);
                        }

                        if (indexModel.configUrl[name].url == "/static/pages-html/pay-relasis/recharge.html") {
                            $("#deStartTime").data("datetimepicker").setDate(self.fourteenStarTime());
                            $("#deEndTime").data("datetimepicker").setDate(self.fourteenEndTime());
                            $(".time-start").text($("#deStartTime").val());
                            $(".time-end").text($("#deEndTime").val());
                        } else if (self.onlyTime() || indexModel.configUrl[name].url == "/static/pages-html/register/keepSummary.html") {
                            $("#deStartTime").data("datetimepicker").setDate(self.twoDayStarTime());
                            $("#deEndTime").data("datetimepicker").setDate(self.twoDayEndTime());
                            $(".time-start").text($("#deStartTime").val());
                            $(".time-end").text($("#deEndTime").val());
                        } else {
                            $("#deStartTime").data("datetimepicker").setDate(self.oneDayTime());
                            $(".time-start").text($("#deStartTime").val());
                        }
                        // indexModel.zeusidChange("");
                       /*  indexModel.channelIdChange("");
                        $("#areaClothing").removeAttr("data-id");
                        $("#channel").removeAttr("data-id"); */
                    }
                    return true;
                }
            },
            
            beforeExpand: function (treeId, treeNode) {     //用于捕获父节点展开之前的事件回调函数，并且根据返回值确定是否允许展开操作

                var pNode = curExpandNode ? curExpandNode.getParentNode() : null;
                //getParentNode: 获取 treeNode 节点的父节点
                var treeNodeP = treeNode.parentTId ? treeNode.getParentNode() : null;
                
                var zTree = $.fn.zTree.getZTreeObj("tree");
                for (var i = 0, l = !treeNodeP ? 0 : treeNodeP.children.length; i < l; i++) {
                    if (treeNode !== treeNodeP.children[i]) {
                        zTree.expandNode(treeNodeP.children[i], false);
                    }
                }
                while (pNode) {
                    if (pNode === treeNode) {
                        break;
                    }
                    pNode = pNode.getParentNode();
                }
                if (!pNode) {
                    singlePath(treeNode);
                }

            },
            onExpand: function onExpand(event, treeId, treeNode) {
                curExpandNode = treeNode;
            }

        }
    };
    //logout & userName
    this.userName = ko.observable("shao nian");
    this.newPassword = ko.observable();
    this.oldPassword = ko.observable();
    this.logout             = function(){
        $.ajax({
            type:"GET",
            url:"",
            async:true,
            cache:false,
            dataType:"json",
            error:function(){
                layer.open({
                    skin: 'layui-layer-lan',
                    closeBtn:1,
                    shadeClose:true,
                    content:'请求失败,请稍后再试！'
                });
            },
            success:function(){
                if (data.code == 200) {
                    window.location.herf = 'login.html';
                }else{
                    layer.open({
                        skin: 'layui-layer-lan',
                        closeBtn: 1,
                        shadeClose:true,
                        content:'请求失败！'
                    });
                }
            }
        });
    }
    //日历时间
    this.datetimepickerObj  = {     
        language: 'zh-CN',
        weekStart: 1,
        todayBtn: 1,
        autoclose: 1,
        todayHighlight: 1,
        startView: 2,
        minView: 2,
        forceParse: 0
    };
    this.timePicker         = ko.observable();
    
    self.deStartTime        = ko.observable("");
    self.deEndTime          = ko.observable("");

    self.oneDayTime         = ko.observable(new Date());
    self.twoDayStarTime     = ko.observable(addDays(new Date(), -6));
    self.twoDayEndTime      = ko.observable(new Date());      
    
    self.fourteenStarTime   = ko.observable(addDays(new Date(), -13));
    self.fourteenEndTime    = ko.observable(new Date());

    this.timeError          = ko.computed(function () {
        if (this.deStartTime() != "" && this.deEndTime() != "" && this.onlyTime()) {
            var startTime = this.deStartTime(),
                endTime   = this.deEndTime(),
                startNum  = parseInt(startTime.replace(/-/g, ''), 10),
                endNum    = parseInt(endTime.replace(/-/g, ''), 10);
            return startNum > endNum ? this.timeErrorShow(true) : this.timeErrorShow(false);
        }
    }, this);
    this.timeErrorShow      = ko.observable(false);
    this.today              = function () {
        $("#deStartTime").data("datetimepicker").setDate(new Date());
        $("#deEndTime").data("datetimepicker").setDate(new Date());
        $('#ensure').trigger("click");
    };
    this.yesterday          = function () {
        $("#deStartTime").data("datetimepicker").setDate(addDays(new Date(), -1));
        $("#deEndTime").data("datetimepicker").setDate(addDays(new Date(), -1));
        $('#ensure').trigger("click");
    };
    this.lastSeven          = function () {
        $("#deStartTime").data("datetimepicker").setDate(addDays(new Date(), -6));
        $("#deEndTime").data("datetimepicker").setDate(new Date());
        $('#ensure').trigger("click");
    };
    this.lastThirty         = function () {
        $("#deStartTime").data("datetimepicker").setDate(addDays(new Date(), -29));
        $("#deEndTime").data("datetimepicker").setDate(new Date());
        $('#ensure').trigger("click");
    };
    this.timeAll            = function () {
        $("#deStartTime").data("datetimepicker").setDate(new Date(0));
        $("#deEndTime").data("datetimepicker").setDate(new Date());
        $('#ensure').trigger("click");
    };
    this.resetTime          = function () {
        this.deStartTime("");
        this.deEndTime("");
    };
    this.correctTime        = function () {
        if ($(".pageContainer").attr("data-url") == "/static/pages-html/pay-relasis/recharge.html") {
            self.fourteenStarTime(new Date(self.deStartTime()));
            self.fourteenEndTime(new Date(self.deEndTime()));
        } else if (self.onlyTime() || $(".pageContainer").attr("data-url") =="/static/pages-html/register/keepSummary.html"){
            self.twoDayStarTime(new Date(self.deStartTime()));
            self.twoDayEndTime(new Date(self.deEndTime()));
        } else {
            self.oneDayTime(new Date(self.deStartTime()));
        }
        
        $(".time-start").text($("#deStartTime").val());
        $(".time-end").text($("#deEndTime").val());
        reload();
        $('#dropdown_time').trigger("click");
    };
    //
    this.itemMallMap        = ko.observable();      //商城道具
    this.mailTypeMap        = ko.observable();
    this.mailTypeArr        = ko.observableArray();
    this.auctionTypeArr     = ko.observableArray();
}

var indexModel = new indexViewModel();
var curExpandNode = null;
function singlePath(newNode) {
    if (newNode === curExpandNode) return;

    var zTree = $.fn.zTree.getZTreeObj("tree"),
        rootNodes, tmpRoot, tmpTId, i, j, n;

    if (!curExpandNode) {
        tmpRoot = newNode;
        while (tmpRoot) {
            tmpTId = tmpRoot.tId;
            tmpRoot = tmpRoot.getParentNode();
        }
        rootNodes = zTree.getNodes();
        for (i = 0, j = rootNodes.length; i < j; i++) {
            n = rootNodes[i];
            if (n.tId != tmpTId) {
                zTree.expandNode(n, false);
            }
        }
    } else if (curExpandNode && curExpandNode.open) {
        if (newNode.parentTId === curExpandNode.parentTId) {
            zTree.expandNode(curExpandNode, false);
        } else {
            var newParents = [];
            while (newNode) {
                newNode = newNode.getParentNode();
                if (newNode === curExpandNode) {
                    newParents = null;
                    break;
                } else if (newNode) {
                    newParents.push(newNode);
                }
            }
            if (newParents != null) {
                var oldNode = curExpandNode;
                var oldParents = [];
                while (oldNode) {
                    oldNode = oldNode.getParentNode();
                    if (oldNode) {
                        oldParents.push(oldNode);
                    }
                }
                if (newParents.length > 0) {
                    zTree.expandNode(oldParents[Math.abs(oldParents.length - newParents.length) - 1], false);
                } else {
                    zTree.expandNode(oldParents[oldParents.length - 1], false);
                }
            }
        }
    }
    curExpandNode = newNode;
}
function reload() {
    
    // 遍历 input标签， 获取value存入 sessionStorage
    $(".form-inline").find('input').each(function (i, v) {
        window.sessionStorage.setItem(`autosaveinput${i}`, $(v).val())
    })

    // 加载页面
    $(".pageContainer").load($(".pageContainer").attr("data-url")+"?number="+Math.random());

    setTimeout(() => {
        if (window.sessionStorage.getItem("autosaveinput0")) {
            // 从 sessionStorage 获取值，填入input
            $(".form-inline").find('input').each(function (i, v) {
                $(v).val(window.sessionStorage.getItem(`autosaveinput${i}`))
            })
        }
    }, 50);
}
var getChannelId, getServerId, zeusidsGlobal, platGlobal, platGlobalKey, platGlobalValue, moneyChangeReason, spinner, rechargeGlobal;

function getStartDate(){
    return $(".time-start").text();
}
function getEndDate(){
    return $(".time-end").text();
}
function getTimeList(startTime, endTime,n) {          //返回特定时间格式的数组
    let startNum = new Date(startTime).getTime(),
        endNum = new Date(endTime).getTime();

    var timeStep = (endNum - startNum) / (1000 * 60 * 60 * 24) + 1;
    var ArrayTime = [];

    for (let i = 0; i < timeStep; i+=n) {
        ArrayTime.push(new Date(startNum).format("yyyy-MM-dd"));
        startNum += 1000 * 60 * 60 * 24 * n;
    }
    return ArrayTime;
}
function heatmapDataformat(xA, yA, yG, data) {         //热力图数据转换(x轴数据，y轴数据，y轴标识，需要转换格式数据)
    var serData = [];

    $.each(data, function (index, obj) {
        for (v in obj) {
            var t = [];
            var y, x, z;
            
            y = yA.indexOf(obj[yG])        // y轴的坐标
            if (y == -1) {
                y = yA.length;
            }
            x = xA.indexOf(v)              // x轴的坐标(根据键名获取坐标)
            z = obj[v] || 0                // 值      
            if (x != -1) {
                t.push(x)
                t.push(y)
                t.push(z)
            }
            if (t.length > 0) {
                serData.push(t)
            }
        }
    });
    return serData;
}

/* $(document).ajaxSuccess(function(event,jqxhr,settings){
    if ((jqxhr.responseText).search(/{"code":100,"message":"用户未登录"}/) != -1) {
        var jqxhrObj = eval('(' + jqxhr.responseText + ')')
        if (jqxhrObj.code == 100) {
            layer.confirm(jqxhrObj.message + ",是否重新登录？", {
                btn: ['确定']
            }, function (index) {
                layer.close(index);
                window.location.href = 'login.html';
            });
        }
    }
}); */

$(document).ajaxError(function (event, jqxhr, settings) {
    layer.alert("Unexpected error, please try again later!", {
        skin: 'layui-layer-lan',
        closeBtn: 1
    });
});

$(function(){
    
    ko.applyBindings(indexModel);       //绑定视图模型
    var minHeight = window.innerHeight - $("header").height()-14;  
    $(".pageContainer").css("min-height",minHeight);

    (function () {                          //默认加载
        indexModel.datetimepickerObj.format = 'yyyy-mm-dd';
        $('.form_datetime').datetimepicker(indexModel.datetimepickerObj);
        
        $("#deStartTime").data("datetimepicker").setDate(addDays(new Date(), -6));
        $("#deEndTime").data("datetimepicker").setDate(new Date());
        $(".time-start").text($("#deStartTime").val());
        $(".time-end").text($("#deEndTime").val());
    })();    
   
    var initsIndex = function () {          //加载配置数据
        $.ajax({
            type:"get",
            url:"global/config",
            async:true,
            cache:false,
            dataType:"json",
            error:function(xhr,status){
                console.log("initsIndex error:",status);
            },
            success:function(data){
                
                
                if (data.code == 200) {
                    console.log("config:", data);
                    
                    if (data.info.platTypeMap) {                    //渠道列表
                        var platKey = [];
                        var platValue = [];
                        
                        for (var j in data.info.platTypeMap) {
                            platKey.push(j)
                            platValue.push(data.info.platTypeMap[j])
                            platGlobal = data.info.platTypeMap;
                        }    
                        platGlobalKey = platKey;
                        platGlobalValue = platValue;    

                    } else {
                        console.log("platTypeMap is no data!")
                    }

                    if (data.info.zeusTypeMap) {                    //区服列表
                        indexModel.zeuidList.removeAll();
                        for (i in data.info.zeusTypeMap) {
                            indexModel.zeuidList.push({ zeusid: i, name: data.info.zeusTypeMap[i] })
                            zeusidsGlobal = data.info.zeusTypeMap;
                        }
                        $("#cloth_wrap input").attr("data-id", $("#cloth_wrap li").eq(0).attr("id"));       //默认加载列表第一项
                        indexModel.zeusidChange($("#cloth_wrap li").eq(0).find("a").text());
                        getServerId = $("#cloth_wrap li").eq(0).attr("id");
                    } else {
                        console.log("zeusTypeMap is no data!");
                    }

                    if (data.info.moneyChangeReasonMap) {           //...
                        moneyChangeReason = data.info.moneyChangeReasonMap;
                    } else {
                        console.log("moneyChangeReasonMap is no data!");
                    }

                    if (data.info.rechargeMap) {                //付费习惯-充值类型
                        rechargeGlobal = data.info.rechargeMap;
                    } else {
                        console.log("rechargeMap is no data.");
                    }
                   
                }
                
                /* if (data.code == 200) {
                    indexModel.channelidList.removeAll();
                    indexModel.mailTypeArr.removeAll();
                    indexModel.auctionTypeArr.removeAll();

                    for (var j in data.info.channelMap) {
                        indexModel.channelidList.push({ channelid: j, name: data.info.channelMap[j] });
                    };
                    for (var k in data.info.mailTypeMap) {
                        indexModel.mailTypeArr.push({ id: k, name: data.info.mailTypeMap[k] });
                    };
                    for (var m in data.info.auctionTypeMap) {
                        indexModel.auctionTypeArr.push({ id: m, name: data.info.auctionTypeMap[m] });
                    };
                    
                    indexModel.itemMallMap(data.info.itemMallMap);      //商城道具
                    indexModel.channelidMap(data.info.channelMap);
                    indexModel.mailTypeMap(data.info.mailTypeMap);
                    zeusidsGlobal       = data.info.zeusMap;
                    consumptionCoinType = data.info.moneyChangeReasonMap;   //货币类型
                } */
            }
        });
    }
    initsIndex();

    $.getJSON("static/index.json", function (result) {      //初始化zTree
        var zTreeObjUserIndex = $.fn.zTree.init($("#tree"), indexModel.userSettingTree, result.TreeList);
    });

    var passwordForm = $(".changePassForm").Validform({
        btnSubmit: "#storePass",
        tiptype: 3,
        showAllError: true,
        callback: function (form) {
            if (indexModel.newPassword() == 123456) {
                $("#newPassword").siblings().css("color", "red").text("新密码不得与初始密码相同！")
                return false;
            } else {
                $('#passwordModal').modal('hide');
                var index = layer.load(1, {
                    shade: [0.3, '#666666'],
                    content: "提交中,请等待......"
                });
                $.ajax({
                    type: "get",
                    url: "",
                    async: true,
                    data: { oldPwd: hex_md5(indexModel.oldPassword()), newPwd: hex_md5(indexModel.newPassword()) },
                    dataType: "json",
                    error: function () {
                        layer.close(index);
                    },
                    success: function (data) {
                        console.log(data)
                        layer.close(index);
                        if (data.code == 200) {
                            layer.msg('修改密码成功！', {
                                icon: 1,
                                shadeClose: true,
                                shade: [0.1, '#fff']
                            });
                        } else {
                            layer.msg('修改密码失败！', {
                                icon: 2,
                                shadeClose: true,
                                shade: [0.1, '#fff']
                            });
                        }
                    }
                });
            }
            return false;
        }
    });
    passwordForm.addRule([
        {
            ele: "#oldPassword",
            datatype: "*5-30",
            nullmsg: "请输入旧密码！",
            errormsg: "请输入5-15字符长度，支持汉字、字母、数字及_ !",
            sucmsg: ""
        },
        {
            ele: "#newPassword",
            datatype: "*5-30",
            nullmsg: "请输入新密码！",
            errormsg: "请输入5-15字符长度，支持汉字、字母、数字及_ !",
            sucmsg: ""
        }
    ]);

    Highcharts.setOptions({             //highcharts 全局设置
        lang:{
            printChart: "打印图表",
            downloadPNG: '下载 PNG  文件',
            downloadJPEG: '下载 JPEG 文件',
            downloadPDF: '下载 PDF   文件',
            downloadSVG: '下载 SVG  文件',
            downloadCSV: '下载 CSV  文件',
            downloadXLS: '下载 XLS   文件',
            viewData: '查看数据表格',
            resetZoom: "恢复缩放",
        },
        credits: {                      //版权信息
            enabled: true,
            text: "磐火网络",
            href: "javascript:void(0);",
            style: {
                fontSize: "12px"
            }
        }
    });

    var opts = {                        //spinner 全局配置
        lines: 7, // The number of lines to draw
        length: 0, // The length of each line
        width: 10, // The line thickness
        radius: 18, // The radius of the inner circle
        scale: 0.85, // Scales overall size of the spinner
        corners: 1, // Corner roundness (0..1)
        color: '#3e4452', // CSS color or array of colors
        fadeColor: 'transparent', // CSS color or array of colors
        opacity: 0.25, // Opacity of the lines
        rotate: 0, // The rotation offset
        direction: 1, // 1: clockwise, -1: counterclockwise
        speed: 1, // Rounds per second
        trail: 60, // Afterglow percentage
        fps: 20, // Frames per second when using setTimeout() as a fallback in IE 9
        zIndex: 2e9, // The z-index (defaults to 2000000000)
        className: 'spinner', // The CSS class to assign to the spinner
        top: '26%', // Top position relative to parent
        left: '50%', // Left position relative to parent
        position: 'absolute' // Element positioning
    };
    spinner = new Spinner(opts);
    
});