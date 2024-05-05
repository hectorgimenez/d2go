package nip

// These stats are currently NOT supported, rules can not contain them otherwise it will return an error
var notSupportedStats = []string{
	"plusmindamage",
	"mindamage",
	"plusmaxdamage",
	"maxdamage",
	"enhanceddamage",
	"itemarmorpercent",
	"itemmindamagepercent",
	"itemslashdamage",
	"itemslashdamagepercent",
	"itemcrushdamage",
	"itemcrushdamagepercent",
	"itemthrustdamage",
	"itemthrustdamagepercent",
	"secondarymindamage",
	"secondarymaxdamage",
	"damagepercent",
}
