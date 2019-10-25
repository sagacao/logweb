package router

import (
	"logweb/handler"
	"logweb/middleware"

	"github.com/henrylee2cn/faygo"
)

// Route register router in a tree style.
func Route(frame *faygo.Framework) {
	frame.Route(
		frame.NewNamedAPI("Index", "GET", "/", handler.Index),
		frame.NewNamedAPI("test struct handler", "POST", "/test", &handler.Test{}).Use(middleware.Token),

		frame.NewGroup("global",
			frame.NewNamedAPI("Config", "GET", "config", handler.Config),
			//frame.NewNamedAPI("Index", "GET", "dynamicData", handler.DynamicData),
		),

		frame.NewGroup("summary",
			frame.NewNamedAPI("渠道汇总", "GET", "summarydata", handler.SummaryData),
			frame.NewNamedAPI("Index", "GET", "modeldata", handler.ModelData),
		),
		frame.NewGroup("statistics",
			frame.NewNamedAPI("留存汇总", "GET", "remaindata", handler.RemainData),
			frame.NewNamedAPI("渠道留存", "GET", "remainplatdata", handler.RemainPlatData),
			// frame.NewNamedAPI("机型数据", "GET", "remainmodeldata", handler.RemainModelData),
			// frame.NewNamedAPI("注册数据", "GET", "registerdata", handler.RegisterData),
			frame.NewNamedAPI("日活跃数据(dau)", "GET", "daudata", handler.DauData),
			frame.NewNamedAPI("平台活跃数据(pau)", "GET", "dauplatdata", handler.DauPlatData),
			frame.NewNamedAPI("流失数据", "GET", "runoffdata", handler.RunoffData),
		),
		frame.NewGroup("onlinedata",
			frame.NewNamedAPI("每日在线", "GET", "online", handler.Online),
			frame.NewNamedAPI("当前在线汇总", "GET", "curonline", handler.CurOnline),
			frame.NewNamedAPI("日平均在线时长", "GET", "avgonline", handler.AvgOnline),
			frame.NewNamedAPI("公会统计", "GET", "leagueline", handler.LeagueLine),
		),
		frame.NewGroup("level",
			frame.NewNamedAPI("等级分布", "GET", "leveldisdata", handler.LevelDisData),
			frame.NewNamedAPI("等级流失", "GET", "levellosedata", handler.LevelLoseData),
			// frame.NewNamedAPI("等级排行", "GET", "levelrankdata", handler.LevelRankData)
		),
		frame.NewGroup("league",
			frame.NewNamedAPI("公会列表", "GET", "leaguedata", handler.LeagueData),
			frame.NewNamedAPI("公会信息", "GET", "leagueinfo", handler.LeagueInfo),
		),
		frame.NewGroup("pay",
			frame.NewNamedAPI("付费等级", "GET", "payleveldata", handler.PayLevelData),
			frame.NewNamedAPI("用户贡献", "GET", "payltv", handler.PayAmountData),
			frame.NewNamedAPI("付费习惯", "GET", "payhabitdata", handler.PayHabitData),
			frame.NewNamedAPI("付费时间分布", "GET", "payhourdata", handler.PayHourData),
			frame.NewNamedAPI("付费排行", "GET", "payrankdata", handler.PayRankData),
			frame.NewNamedAPI("每日付费排行", "GET", "payrankdaydata", handler.PayRankDayData),
			frame.NewNamedAPI("Index", "GET", "payvipdata", handler.VipData),
			frame.NewNamedAPI("用户每日付费信息", "GET", "paydailydata", handler.PayDailyData),
			frame.NewNamedAPI("渠道用户付费信息", "GET", "paychanneldata", handler.PayChannelData),
			frame.NewNamedAPI("付费流失", "GET", "paylossdata", handler.PayLossData),
		),

		frame.NewGroup("importfilter",
			frame.NewNamedAPI("账号过滤", "POST", "import", handler.ImportFilterAccount),
		),

		frame.NewGroup("costAnalysis",

			frame.NewNamedAPI("货币消耗信息", "GET", "moneychange", handler.MoneyChangePage),
			frame.NewNamedAPI("货币消耗列表", "GET", "moneychangelist", handler.MoneyChangeList),

			frame.NewNamedAPI("商城消耗信息", "GET", "mallcost", handler.MallCostPage),
			frame.NewNamedAPI("商城消耗列表", "GET", "mallcostlist", handler.MallCostList),

			frame.NewNamedAPI("充值", "GET", "recharge", handler.Recharge),
			frame.NewNamedAPI("新增充值", "GET", "newlyrecharge", handler.NewlyRecharge),
			frame.NewNamedAPI("每小时充值", "GET", "hourrecharge", handler.HourRecharge),
		),

		frame.NewGroup("gameLog",

			frame.NewNamedAPI("账号登录日志", "GET", "getAccountLoginLogPage", handler.AccountLoginLogPage),
			frame.NewNamedAPI("账号登出日志", "GET", "getAccountLogoutLogPage", handler.AccountLogoutLogPage),

			frame.NewNamedAPI("角色登录日志", "GET", "getPlayerLoginLogPage", handler.PlayerLoginLogPage),
			frame.NewNamedAPI("角色登出日志", "GET", "getPlayerLogoutLogPage", handler.PlayerLogoutLogPage),
			frame.NewNamedAPI("角色升级日志", "GET", "getPlayerLevelUpLogPage", handler.PlayerLevelUpLogPage),

			frame.NewNamedAPI("邮件发送日志", "GET", "getMailSendLogPage", handler.MailSendLogPage),
			frame.NewNamedAPI("邮件删除日志", "GET", "getMailDeleteLogPage", handler.MailDeleteLogPage),

			frame.NewNamedAPI("商城日志", "GET", "getPlayerDoMallLogPage", handler.PlayerDoMallLogPage),

			frame.NewNamedAPI("订单查询", "GET", "getAddGoldLogPage", handler.GetAddGoldLogPage),

			frame.NewNamedAPI("道具上架", "GET", "getAddAuctionLogPage", handler.AddAuctionLogPage),
			frame.NewNamedAPI("道具下架", "GET", "getCancelAuctionLogPage", handler.CancelAuctionLogPage),
			frame.NewNamedAPI("道具成交", "GET", "getDoneAuctionLogPage", handler.DoneAuctionLogPage),

			frame.NewNamedAPI("物品获取", "GET", "getItemObtainLogPage", handler.ItemObtainLogPage),
			frame.NewNamedAPI("物品丢失", "GET", "getItemDisappearLogPage", handler.ItemDisappearLogPage),

			frame.NewNamedAPI("任务接取", "GET", "getTaskAcceptLogPage", handler.TaskAcceptLogPage),
			frame.NewNamedAPI("任务完成", "GET", "getTaskSuccessLogPage", handler.TaskSuccessLogPage),

			frame.NewNamedAPI("公会信息", "GET", "getLeagueDataLogPage", handler.LeagueDataLogPage),

			frame.NewNamedAPI("货币产出消耗", "GET", "getMoneyChangeLogPage", handler.MoneyChangeLogPage),
			frame.NewNamedAPI("侠客魂石日志", "GET", "getVirtualPartnerSoulLogPage", handler.VirtualPartnerSoulLogPage),
			frame.NewNamedAPI("侠客日志", "GET", "getVirtualPartnerLogPage", handler.VirtualPartnerLogPage),

			frame.NewNamedAPI("原始日志", "GET", "getRawLogPage", handler.RawLogPage),
			frame.NewNamedAPI("分类日志", "GET", "getRawFormatLogPage", handler.RawFormatLogPage),

			frame.NewNamedAPI("虚拟货币日志", "GET", "getVirtualMoneyLogPage", handler.VirtualMoneyLogPage),
		),
	)
}
