package log

type DeleteMail struct {
	ReceiverId 	string `json:"receiverId" bson:"roleid"`
	MailId  	string `json:"mailId" bson:"mailid"`
	LogTime  	string `json:"logTime" bson:"logtime"`
}

type SendMail struct {
	SenderId 	string `json:"senderId" bson:"roleid"`
	ReceiverId  string `json:"receiverId" bson:"receiveid"`
	MailType 	string `json:"mailType" bson:""`
	MailId 		string `json:"mailId" bson:"gameuserid"`
	Gold 		string `json:"gold" bson:"lockgold"`
	ItemId 		string `json:"itemId" bson:"attachinfo"`
	ItemCount 	string `json:"itemCount" bson:"account"`
	LogTime 	string `json:"logTime" bson:"logtime"`
}

