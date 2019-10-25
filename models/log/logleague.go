package log

type LeagueData struct {
	GuildId   string `json:"guildId" bson:"roleid"`
	GuildName string `json:"guildName" bson:"leadername"`
	Level    string `json:"level" bson:"state"`
	Money     string `json:"money" bson:"assets"`
	Liveness   string `json:"liveness" bson:"health"`
}