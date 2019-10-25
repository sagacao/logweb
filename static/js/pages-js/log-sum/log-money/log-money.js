$(function(){

	indexModel.systemChannelHide();

	var logMoneyViewModel = function(){
		this.dataList				= ko.observableArray();
		this.roleId					= ko.observable("");
		//this.moneyType				= ko.observableArray(indexModel.moneyTypeArr());
		//this.moneyType				= ko.observableArray([{id:1,name:'Money'},{id:3,name:'LockGold'},{id:4,name:'Gold'}]);
		this.moneyType				= ko.observableArray([{id:1,name:'银币'},{id:2,name:'经验'}, {id:3,name:'金币'},{id:4,name:'勾玉'}, 
														{id:7,name:'帮贡'}, {id:9,name:'技能经验'}, {id:13,name:'阵营威望'}, 
														{id:27,name:'蘑菇币I'}, {id:28,name:'蘑菇币II'}, {id:29,name:'蘑菇币III'}]);
		this.selectedMoneyType		= ko.observable();
		this.changeMoneyValue		= function(){
			pageClick(1);
		};
		this.numType				= ko.observableArray([{id:1,name:'产出'},{id:-1,name:'消耗'}]);
		this.selectedNumType		= ko.observable();
		this.changeNumValue			= function(){
			pageClick(1);
		};
		this.getServerIdChange		= ko.computed(function(){
			if (indexModel.zeusidChange() == "" && $(".pageContainer").attr("data-url") == "/static/pages-html/log-sum/log-money/log-money.html"){
				var idCurrent = $('.clothLis li:first').attr("id");
				$('.clothLis li:first').parents("ul").siblings("input").attr("data-id",idCurrent);
				indexModel.zeusidChange($('.clothLis li:first').children("a").text());
				getServerId = idCurrent;
			}
		},this);

		this.saveRoleid 			= ko.observableArray();
	}
	
	var newLogModel = new logMoneyViewModel();
	ko.applyBindings(newLogModel,document.getElementById("logMoney"));

	
	var pageClick = function( pageclickednumber ){
		var index = layer.load(1, {
		  	shade: [0.3,'#666666'],
		  	content:"数据获取中......"
		});
		var jsonData = {zeusid:getServerId,startTime:getStartDate(),endTime:getEndDate(),pageIndex:pageclickednumber,pageSize:50};
		newLogModel.roleId() ? $.extend(jsonData,{roleId: newLogModel.roleId()}) : $.extend(jsonData,{});
		newLogModel.selectedNumType() ? $.extend(jsonData,{delta: newLogModel.selectedNumType()}) : $.extend(jsonData,{});
		newLogModel.selectedMoneyType() ? $.extend(jsonData,{moneyType: newLogModel.selectedMoneyType()}) : $.extend(jsonData,{});

		// 保存记录到sessionStorage
		if (newLogModel.roleId()) {
			if (window.sessionStorage.getItem('roleid')) {
				let sessArrRole = JSON.parse(window.sessionStorage.getItem('roleid'))
				sessArrRole.push(newLogModel.roleId())
				window.sessionStorage.setItem('roleid', JSON.stringify(sessArrRole))
			} else {
				let sessArrRole = [newLogModel.roleId()]
				window.sessionStorage.setItem('roleid', JSON.stringify(sessArrRole))
			}
			// 从sessionStorage 获取输入记录
			newLogModel.saveRoleid(JSON.parse(window.sessionStorage.getItem('roleid')))
		}
		
		$.ajax({
			type : "get",
			async:true,
			url:"/gameLog/getMoneyChangeLogPage",
			data : jsonData,
			dataType : "json",
			error:function(){
				layer.close(index); 
			},
			success : function(data){
				console.log("游戏币管理日志：",data)
				layer.close(index); 
				if( data.code == 200 ){
					var tabDom = $("#logBuild_tbody").empty();
					newLogModel.dataList.removeAll();
					if (data.info.rows) {
						for (var i = 0; i < data.info.rows.length; i++) {
							//data.info.rows[i].channelidName = indexModel.channelidMap()[data.info.rows[i].channelid];
							newLogModel.dataList.push(data.info.rows[i]);
						}
						$("#pager").pager({ pagenumber: pageclickednumber, pagecount: data.info.totalPage, buttonClickCallback: pageClick });
					}else{
						console.log("log-money data.info.rows is no data.");
					}
					
				}else{
					layer.alert('数据加载失败！', {
						skin: 'layui-layer-lan',//样式类名
						closeBtn: 0
					});
				}
			}
		});
	}
	pageClick(1);
	
	var roleForm = $("#logMoneyForm").Validform({
		btnSubmit:"#accountIdMatch",
		tiptype:function(msg,o,cssctl){
			var objtip=$("#tipShow");
			cssctl(objtip,o.type);
			objtip.text(msg);
		},
		callback: function(form){
			pageClick(1);
			return false;
		}
	});
})
