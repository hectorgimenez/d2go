package nip

// These stats are currently NOT supported, rules can not contain them otherwise it will return an error
var notSupportedStats = []string{
	"plusmindamage",
	"plusmaxdamage",
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

// These stats are partially supported, rules can contain them but the ones matching the item type in this list will be blocked
var blockedStatsForItemType = map[string][]string{
	"mindamage":      {"axe", "wand", "club", "scepter", "mace", "hammer", "sword", "knife", "thrownweapon", "throwingaxe", "javelin", "spear", "polearm", "staff", "bow", "crossbow", "assassinclaw", "orb", "amazonbow", "amazonspear", "amazonjavelin"},
	"maxdamage":      {"axe", "wand", "club", "scepter", "mace", "hammer", "sword", "knife", "thrownweapon", "throwingaxe", "javelin", "spear", "polearm", "staff", "bow", "crossbow", "assassinclaw", "orb", "amazonbow", "amazonspear", "amazonjavelin"},
	"enhanceddamage": {"axe", "wand", "club", "scepter", "mace", "hammer", "sword", "knife", "thrownweapon", "throwingaxe", "javelin", "spear", "polearm", "staff", "bow", "crossbow", "assassinclaw", "orb", "amazonbow", "amazonspear", "amazonjavelin"},
}
