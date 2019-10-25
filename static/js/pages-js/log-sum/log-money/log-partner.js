$(function () {

    indexModel.systemChannelHide();

    var logPartnerViewModel = function () {
        this.roleId         = ko.observable("");
        this.numType        = ko.observableArray([{ id: 1, name: '产出' }, { id: -1, name: '消耗' }]);
        this.selectedNumType = ko.observable();
        this.changeNumValue = function () {
            pageClick(1);
        };

        this.dataList_partner_soul = ko.observableArray();
        this.dataList_pantner = ko.observableArray();

        this.saveRoleid = ko.observableArray();
    }

    var newPartnerModel = new logPartnerViewModel();
    ko.applyBindings(newPartnerModel, document.getElementById("logPartner"));

    var pageClick = function (pageclickednumber) {
        var index = layer.load(1, {
            shade: [0.3, '#666666'],
            content: "数据获取中......"
        });
        
        var jsonData = {
            zeusid: getServerId,
            startTime: getStartDate(), 
            endTime: getEndDate(),
            pageIndex: pageclickednumber, 
            pageSize: 50
        }

        newPartnerModel.roleId() ? $.extend(jsonData, { roleId: newPartnerModel.roleId() }) : $.extend(jsonData, {});
        newPartnerModel.selectedNumType() ? $.extend(jsonData, { delta: newPartnerModel.selectedNumType() }) : $.extend(jsonData, {});

        // 保存记录到sessionStorage
        if (newPartnerModel.roleId()) {
            if (window.sessionStorage.getItem('roleid')) {
                let sessArrRole = JSON.parse(window.sessionStorage.getItem('roleid'))
                sessArrRole.push(newPartnerModel.roleId())
                window.sessionStorage.setItem('roleid', JSON.stringify(sessArrRole))
            } else {
                let sessArrRole = [newPartnerModel.roleId()]
                window.sessionStorage.setItem('roleid', JSON.stringify(sessArrRole))
            }
            // 从sessionStorage 获取输入记录
            newPartnerModel.saveRoleid(JSON.parse(window.sessionStorage.getItem('roleid')))
        }
        
        $.ajax({
            type: "get",
            async: true,
            url: "/gameLog/getVirtualPartnerSoulLogPage",
            data: jsonData,
            dataType: 'json',
            error: function () {
                layer.close(index); 
            },
            success: function (data) {
                layer.close(index); 
                console.log("PartnerSoul:", data);

                var tabDom = $("#logpartner_soul_tbody").empty();
                newPartnerModel.dataList_partner_soul.removeAll();
                if (data.info.rows) {
                    for (var i = 0; i < data.info.rows.length; i++) {
                        newPartnerModel.dataList_partner_soul.push(data.info.rows[i]);
                    }
                    $("#pager").pager({ pagenumber: pageclickednumber, pagecount: data.info.totalPage, buttonClickCallback: pageClick });
                }
            }
        })

        $.ajax({
            type: "get",
            async: true,
            url: "/gameLog/getVirtualPartnerLogPage",
            data: jsonData,
            dataType: 'json',
            error: function () {
                layer.close(index);
            },
            success: function (data) {
                layer.close(index);
                console.log("Partner:", data);

                var tabDom = $("#logpartner_tbody").empty();
                newPartnerModel.dataList_pantner.removeAll();
                if (data.info.rows) {
                    for (var i = 0; i < data.info.rows.length; i++) {
                        newPartnerModel.dataList_pantner.push(data.info.rows[i]);
                    }
                    $("#pager").pager({ pagenumber: pageclickednumber, pagecount: data.info.totalPage, buttonClickCallback: pageClick });
                }
            }
        })

        
    }
    
    var roleForm = $("#logPartnerForm").Validform({
        btnSubmit: "#accountIdMatch",
        tiptype: function (msg, o, cssctl) {
            var objtip = $("#tipShow");
            cssctl(objtip, o.type);
            objtip.text(msg);
        },
        callback: function (form) {
            pageClick(1);
            return false;
        }
    });
    
})