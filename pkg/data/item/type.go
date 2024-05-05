package item

import "strings"

var typeMapping = map[string][]string{
	"amulet":        {"Amulet"},
	"gold":          {"Gold"},
	"ring":          {"Ring"},
	"jewel":         {"Jewel"},
	"axe":           {"HandAxe", "Axe", "DoubleAxe", "MilitaryPick", "WarAxe", "LargeAxe", "BroadAxe", "BattleAxe", "GreatAxe", "GiantAxe", "Hatchet", "Cleaver", "TwinAxe", "Crowbill", "Naga", "MilitaryAxe", "BeardedAxe", "Tabar", "GothicAxe", "AncientAxe", "Tomahawk", "SmallCrescent", "EttinAxe", "WarSpike", "BerserkerAxe", "FeralAxe", "SilverEdgedAxe", "Decapitator", "ChampionAxe", "GloriousAxe"},
	"wand":          {"Wand", "YewWand", "BoneWand", "GrimWand", "BurntWand", "PetrifiedWand", "TombWand", "GraveWand", "PolishedWand", "GhostWand", "LichWand", "UnearthedWand"},
	"club":          {"Club", "SpikedClub", "Cudgel", "BarbedClub", "Truncheon", "TyrantClub"},
	"scepter":       {"Scepter", "GrandScepter", "WarScepter", "RuneScepter", "HolyWaterSprinkler", "DivineScepter", "MightyScepter", "SeraphRod", "Caduceus"},
	"mace":          {"Mace", "MorningStar", "Flail", "FlangedMace", "JaggedStar", "Knout", "ReinforcedMace", "DevilStar", "Scourge"},
	"hammer":        {"WarHammer", "Maul", "GreatMaul", "BattleHammer", "WarClub", "MartelDeFer", "LegendaryMallet", "OgreMaul", "ThunderMaul"},
	"sword":         {"ShortSword", "Scimitar", "Sabre", "Falchion", "CrystalSword", "BroadSword", "LongSword", "WarSword", "TwoHandedSword", "Claymore", "GiantSword", "BastardSword", "Flamberge", "GreatSword", "Gladius", "Cutlass", "Shamshir", "Tulwar", "DimensionalBlade", "BattleSword", "RuneSword", "AncientSword", "Espandon", "DacianFalx", "TuskSword", "GothicSword", "Zweihander", "ExecutionerSword", "Falcata", "Ataghan", "ElegantBlade", "HydraEdge", "PhaseBlade", "ConquestSword", "CrypticSword", "MythicalSword", "LegendSword", "HighlandBlade", "BalrogBlade", "ChampionSword", "ColossusSword", "ColossalSword", "ColossusBlade"},
	"knife":         {"Dagger", "Dirk", "Kris", "Blade", "Poignard", "Rondel", "Cinquedeas", "Stiletto", "BoneKnife", "MithrilPoint", "FangedKnife", "LegendSpike"},
	"thrownweapon":  {"ThrowingKnife", "BalancedKnife", "BattleDart", "WarDart", "FlyingKnife", "WingedKnife"},
	"throwingaxe":   {"ThrowingAxe", "BalancedAxe", "Francisca", "Hurlbat", "FlyingAxe", "WingedAxe"},
	"javelin":       {"Javelin", "Pilum", "ShortSpear", "Glaive", "ThrowingSpear", "WarJavelin", "GreatPilum", "Simbilan", "Spiculum", "Harpoon", "HyperionJavelin", "StygianPilum", "BalrogSpear", "GhostGlaive", "WingedHarpoon"},
	"spear":         {"Spear", "Trident", "Brandistock", "Spetum", "Pike", "WarSpear", "Fuscina", "WarFork", "Yari", "Lance", "HyperionSpear", "StygianPike", "Mancatcher", "GhostSpear", "WarPike"},
	"polearm":       {"Bardiche", "Voulge", "Scythe", "Poleaxe", "Halberd", "WarScythe", "LochaberAxe", "Bill", "BattleScythe", "Partizan", "BecDeCorbin", "GrimScythe", "OgreAxe", "ColossusVoulge", "Thresher", "CrypticAxe", "GreatPoleaxe", "GiantThresher"},
	"staff":         {"ShortStaff", "LongStaff", "GnarledStaff", "BattleStaff", "WarStaff", "JoStaff", "QuarterStaff", "CedarStaff", "GothicStaff", "RuneStaff", "WalkingStick", "Stalagmite", "ElderStaff", "Shillelagh", "ArchonStaff"},
	"bow":           {"ShortBow", "HuntersBow", "LongBow", "CompositeBow", "ShortBattleBow", "LongBattleBow", "ShortWarBow", "LongWarBow", "EdgeBow", "RazorBow", "CedarBow", "DoubleBow", "ShortSiegeBow", "LargeSiegeBow", "RuneBow", "GothicBow", "SpiderBow", "BladeBow", "ShadowBow", "GreatBow", "DiamondBow", "CrusaderBow", "WardBow", "HydraBow"},
	"crossbow":      {"LightCrossbow", "Crossbow", "HeavyCrossbow", "RepeatingCrossbow", "Arbalest", "SiegeCrossbow", "Ballista", "ChuKoNu", "PelletBow", "GorgonCrossbow", "ColossusCrossbow", "DemonCrossBow"},
	"helm":          {"Cap", "SkullCap", "Helm", "FullHelm", "GreatHelm", "Crown", "Mask", "BoneHelm", "WarHat", "Sallet", "Casque", "Basinet", "WingedHelm", "GrandCrown", "DeathMask", "GrimHelm", "Shako", "Hydraskull", "Armet", "GiantConch", "SpiredHelm", "Corona", "DemonHead", "BoneVisage"},
	"armor":         {"QuiltedArmor", "LeatherArmor", "HardLeatherArmor", "StuddedLeather", "RingMail", "ScaleMail", "ChainMail", "BreastPlate", "SplintMail", "PlateMail", "FieldPlate", "GothicPlate", "FullPlateMail", "AncientArmor", "LightPlate", "GhostArmor", "SerpentskinArmor", "DemonhideArmor", "TrellisedArmor", "LinkedMail", "TigulatedMail", "MeshArmor", "Cuirass", "RussetArmor", "TemplarCoat", "SharktoothArmor", "EmbossedPlate", "ChaosArmor", "OrnatePlate", "MagePlate", "DuskShroud", "Wyrmhide", "ScarabHusk", "WireFleece", "DiamondMail", "LoricatedMail", "Boneweave", "GreatHauberk", "BalrogSkin", "HellforgePlate", "KrakenShell", "LacqueredPlate", "ShadowPlate", "SacredArmor", "ArchonPlate"},
	"shield":        {"Buckler", "SmallShield", "LargeShield", "KiteShield", "TowerShield", "GothicShield", "BoneShield", "SpikedShield", "Defender", "RoundShield", "Scutum", "DragonShield", "Pavise", "AncientShield", "GrimShield", "BarbedShield", "Heater", "Luna", "Hyperion", "Monarch", "Aegis", "Ward", "TrollNest", "BladeBarrier"},
	"gloves":        {"LeatherGloves", "HeavyGloves", "ChainGloves", "LightGauntlets", "Gauntlets", "DemonhideGloves", "SharkskinGloves", "HeavyBracers", "BattleGauntlets", "WarGauntlets", "BrambleMitts", "VampireboneGloves", "Vambraces", "CrusaderGauntlets", "OgreGauntlets"},
	"boots":         {"Boots", "HeavyBoots", "ChainBoots", "LightPlatedBoots", "Greaves", "DemonhideBoots", "SharkskinBoots", "MeshBoots", "BattleBoots", "WarBoots", "WyrmhideBoots", "ScarabshellBoots", "BoneweaveBoots", "MirroredBoots", "MyrmidonGreaves"},
	"belt":          {"Sash", "LightBelt", "Belt", "HeavyBelt", "PlatedBelt", "DemonhideSash", "SharkskinBelt", "MeshBelt", "BattleBelt", "WarBelt", "SpiderwebSash", "VampirefangBelt", "MithrilCoil", "TrollBelt", "ColossusGirdle"},
	"circlet":       {"Circlet", "Coronet", "Tiara", "Diadem"},
	"assassinclaw":  {"Katar", "WristBlade", "HatchetHands", "Cestus", "Claws", "BladeTalons", "ScissorsKatar", "Quhab", "WristSpike", "Fascia", "HandScythe", "GreaterClaws", "GreaterTalons", "ScissorsQuhab", "Suwayyah", "WristSword", "WarFist", "BattleCestus", "FeralClaws", "RunicTalons", "ScissorsSuwayyah"},
	"orb":           {"EagleOrb", "SacredGlobe", "SmokedSphere", "ClaspedOrb", "JaredsStone", "GlowingOrb", "CrystallineGlobe", "CloudySphere", "SparklingBall", "SwirlingCrystal", "HeavenlyStone", "EldritchOrb", "DemonHeart", "VortexOrb", "DimensionalShard"},
	"amazonbow":     {"StagBow", "ReflexBow", "AshwoodBow", "CeremonialBow", "MatriarchalBow", "GrandMatronBow"},
	"amazonspear":   {"MaidenSpear", "MaidenPike", "CeremonialSpear", "CeremonialPike", "MatriarchalSpear", "MatriarchalPike"},
	"amazonjavelin": {"MaidenJavelin", "CeremonialJavelin", "MatriarchalJavelin"},
	"pelt":          {"WolfHead", "HawkHelm", "Antlers", "FalconMask", "SpiritMask", "AlphaHelm", "GriffonHeaddress", "HuntersGuise", "SacredFeathers", "TotemicMask", "BloodSpirit", "SunSpirit", "EarthSpirit", "SkySpirit", "DreamSpirit"},
	"primalhelm":    {"JawboneCap", "FangedHelm", "HornedHelm", "AssaultHelmet", "AvengerGuard", "JawboneVisor", "LionHelm", "RageMask", "SavageHelmet", "SlayerGuard", "CarnageHelm", "FuryVisor", "DestroyerHelm", "ConquerorCrown", "GuardianCrown"},
	"auricshields":  {"Targe", "Rondache", "HeraldicShield", "AerinShield", "CrownShield", "AkaranTarge", "AkaranRondache", "ProtectorShield", "GildedShield", "RoyalShield", "SacredTarge", "SacredRondache", "KurastShield", "ZakarumShield", "VortexShield"},
	"voodooheads":   {"PreservedHead", "ZombieHead", "UnravellerHead", "GargoyleHead", "DemonHeadShield", "MummifiedTrophy", "FetishTrophy", "SextonTrophy", "CantorTrophy", "HierophantTrophy", "MinionSkull", "HellspawnSkull", "OverseerSkull", "SuccubusSkull", "BloodlordSkull"},
	"runes":         {"ElRune", "EldRune", "TirRune", "NefRune", "EthRune", "IthRune", "TalRune", "RalRune", "OrtRune", "ThulRune", "AmnRune", "SolRune", "ShaelRune", "DolRune", "HelRune", "IoRune", "LumRune", "KoRune", "FalRune", "LemRune", "PulRune", "UmRune", "MalRune", "IstRune", "GulRune", "VexRune", "OhmRune", "LoRune", "SurRune", "BerRune", "JahRune", "ChamRune", "ZodRune"},
	"ubers":         {"KeyOfTerror", "KeyOfHate", "KeyOfDestruction", "DiablosHorn", "BaalsEye", "MephistosBrain"},
	"tokens":        {"TokenofAbsolution", "TwistedEssenceOfSuffering", "ChargedEssenceOfHatred", "BurningEssenceOfTerror", "FesteringEssenceOfDestruction"},
	"chippedgems":   {"ChippedAmethyst", "ChippedDiamond", "ChippedEmerald", "ChippedRuby", "ChippedSapphire", "ChippedSkull", "ChippedTopaz"},
	"flawedgems":    {"FlawedAmethyst", "FlawedDiamond", "FlawedEmerald", "FlawedRuby", "FlawedSapphire", "FlawedSkull", "FlawedTopaz"},
	"gems":          {"Amethyst", "Diamond", "Emerald", "Ruby", "Skull", "Sapphire", "Topaz"},
	"flawlessgems":  {"FlawlessAmethyst", "FlawlessDiamond", "FlawlessEmerald", "FlawlessRuby", "FlawlessSapphire", "FlawlessSkull", "FlawlessTopaz"},
	"perfectgems":   {"PerfectAmethyst", "PerfectDiamond", "PerfectEmerald", "PerfectRuby", "PerfectSapphire", "PerfectSkull", "PerfectTopaz"},
}

var QuestItems = []string{
	"TheGidbinn",
	"WirtsLeg",
	"HoradricMalus",
	"HellforgeHammer",
	"HoradricStaff",
	"StaffOfKings",
	"KhalimsFlail",
	"AmuletOfTheViper",
	"KhalimsEye",
	"KhalimsHeart",
	"KhalimsBrain",
	"KhalimsWill",
	"ScrollOfInifuss",
	"KeyToTheCairnStones",
	"HoradricCube",
	"HoradricScroll",
	"MephistosSoulstone",
	"BookOfSkill",
}

func TypeForItemName(itemName string) (string, bool) {
	for t, itemNames := range typeMapping {
		for _, name := range itemNames {
			if strings.EqualFold(itemName, name) {
				return t, true
			}
		}
	}

	return "", false
}
