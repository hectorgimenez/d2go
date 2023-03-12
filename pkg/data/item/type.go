package item

import "strings"

var typeMapping = map[string][]string{
	"Axes":               {"HandAxe", "Axe", "DoubleAxe", "MilitaryPick", "WarAxe", "LargeAxe", "BroadAxe", "BattleAxe", "GreatAxe", "GiantAxe", "Hatchet", "Cleaver", "TwinAxe", "Crowbill", "Naga", "MilitaryAxe", "BeardedAxe", "Tabar", "GothicAxe", "AncientAxe", "Tomahawk", "SmallCrescent", "EttinAxe", "WarSpike", "BerserkerAxe", "FeralAxe", "SilverEdgedAxe", "Decapitator", "ChampionAxe", "GloriousAxe"},
	"Wands":              {"Wand", "YewWand", "BoneWand", "GrimWand", "BurntWand", "PetrifiedWand", "TombWand", "GraveWand", "PolishedWand", "GhostWand", "LichWand", "UnearthedWand"},
	"Clubs":              {"Club", "SpikedClub", "Cudgel", "BarbedClub", "Truncheon", "TyrantClub"},
	"Scepters":           {"Scepter", "GrandScepter", "WarScepter", "RuneScepter", "HolyWaterSprinkler", "DivineScepter", "MightyScepter", "SeraphRod", "Caduceus"},
	"Maces":              {"Mace", "MorningStar", "Flail", "FlangedMace", "JaggedStar", "Knout", "ReinforcedMace", "DevilStar", "Scourge"},
	"Hammers":            {"WarHammer", "Maul", "GreatMaul", "BattleHammer", "WarClub", "MartelDeFer", "LegendaryMallet", "OgreMaul", "ThunderMaul"},
	"Swords":             {"ShortSword", "Scimitar", "Sabre", "Falchion", "CrystalSword", "BroadSword", "LongSword", "WarSword", "TwoHandedSword", "Claymore", "GiantSword", "BastardSword", "Flamberge", "GreatSword", "Gladius", "Cutlass", "Shamshir", "Tulwar", "DimensionalBlade", "BattleSword", "RuneSword", "AncientSword", "Espandon", "DacianFalx", "TuskSword", "GothicSword", "Zweihander", "ExecutionerSword", "Falcata", "Ataghan", "ElegantBlade", "HydraEdge", "PhaseBlade", "ConquestSword", "CrypticSword", "MythicalSword", "LegendSword", "HighlandBlade", "BalrogBlade", "ChampionSword", "ColossusSword", "ColossusBlade"},
	"Daggers":            {"Dagger", "Dirk", "Kris", "Blade", "Poignard", "Rondel", "Cinquedeas", "Stiletto", "BoneKnife", "MithrilPoint", "FangedKnife", "LegendSpike"},
	"ThrowingKnifes":     {"ThrowingKnife", "BalancedKnife", "BattleDart", "WarDart", "FlyingKnife", "WingedKnife"},
	"ThrowingAxes":       {"ThrowingAxe", "BalancedAxe", "Francisca", "Hurlbat", "FlyingAxe", "WingedAxe"},
	"Javelins":           {"Javelin", "Pilum", "ShortSpear", "Glaive", "ThrowingSpear", "WarJavelin", "GreatPilum", "Simbilan", "Spiculum", "Harpoon", "HyperionJavelin", "StygianPilum", "BalrogSpear", "GhostGlaive", "WingedHarpoon"},
	"Spears":             {"Spear", "Trident", "Brandistock", "Spetum", "Pike", "WarSpear", "Fuscina", "WarFork", "Yari", "Lance", "HyperionSpear", "StygianPike", "Mancatcher", "GhostSpear", "WarPike"},
	"Polearms":           {"Bardiche", "Voulge", "Scythe", "Poleaxe", "Halberd", "WarScythe", "LochaberAxe", "Bill", "BattleScythe", "Partizan", "BecDeCorbin", "GrimScythe", "OgreAxe", "ColossusVoulge", "Thresher", "CrypticAxe", "GreatPoleaxe", "GiantThresher"},
	"Staves":             {"ShortStaff", "LongStaff", "GnarledStaff", "BattleStaff", "WarStaff", "JoStaff", "QuarterStaff", "CedarStaff", "GothicStaff", "RuneStaff", "WalkingStick", "Stalagmite", "ElderStaff", "Shillelagh", "ArchonStaff"},
	"Bows":               {"ShortBow", "HuntersBow", "LongBow", "CompositeBow", "ShortBattleBow", "LongBattleBow", "ShortWarBow", "LongWarBow", "EdgeBow", "RazorBow", "CedarBow", "DoubleBow", "ShortSiegeBow", "LargeSiegeBow", "RuneBow", "GothicBow", "SpiderBow", "BladeBow", "ShadowBow", "GreatBow", "DiamondBow", "CrusaderBow", "WardBow", "HydraBow"},
	"Crossbows":          {"LightCrossbow", "Crossbow", "HeavyCrossbow", "RepeatingCrossbow", "Arbalest", "SiegeCrossbow", "Ballista", "ChuKoNu", "PelletBow", "GorgonCrossbow", "ColossusCrossbow", "DemonCrossBow"},
	"Helms":              {"Cap", "SkullCap", "Helm", "FullHelm", "GreatHelm", "Crown", "Mask", "BoneHelm", "WarHat", "Sallet", "Casque", "Basinet", "WingedHelm", "GrandCrown", "DeathMask", "GrimHelm", "Shako", "Hydraskull", "Armet", "GiantConch", "SpiredHelm", "Corona", "DemonHead", "BoneVisage"},
	"Armors":             {"QuiltedArmor", "LeatherArmor", "HardLeatherArmor", "StuddedLeather", "RingMail", "ScaleMail", "ChainMail", "BreastPlate", "SplintMail", "PlateMail", "FieldPlate", "GothicPlate", "FullPlateMail", "AncientArmor", "LightPlate", "GhostArmor", "SerpentskinArmor", "DemonhideArmor", "TrellisedArmor", "LinkedMail", "TigulatedMail", "MeshArmor", "Cuirass", "RussetArmor", "TemplarCoat", "SharktoothArmor", "EmbossedPlate", "ChaosArmor", "OrnatePlate", "MagePlate", "DuskShroud", "Wyrmhide", "ScarabHusk", "WireFleece", "DiamondMail", "LoricatedMail", "Boneweave", "GreatHauberk", "BalrogSkin", "HellforgePlate", "KrakenShell", "LacqueredPlate", "ShadowPlate", "SacredArmor", "ArchonPlate"},
	"Shields":            {"Buckler", "SmallShield", "LargeShield", "KiteShield", "TowerShield", "GothicShield", "BoneShield", "SpikedShield", "Defender", "RoundShield", "Scutum", "DragonShield", "Pavise", "AncientShield", "GrimShield", "BarbedShield", "Heater", "Luna", "Hyperion", "Monarch", "Aegis", "Ward", "TrollNest", "BladeBarrier"},
	"Gloves":             {"LeatherGloves", "HeavyGloves", "ChainGloves", "LightGauntlets", "Gauntlets", "DemonhideGloves", "SharkskinGloves", "HeavyBracers", "BattleGauntlets", "WarGauntlets", "BrambleMitts", "VampireboneGloves", "Vambraces", "CrusaderGauntlets", "OgreGauntlets"},
	"Boots":              {"Boots", "HeavyBoots", "ChainBoots", "LightPlatedBoots", "Greaves", "DemonhideBoots", "SharkskinBoots", "MeshBoots", "BattleBoots", "WarBoots", "WyrmhideBoots", "ScarabshellBoots", "BoneweaveBoots", "MirroredBoots", "MyrmidonGreaves"},
	"Belts":              {"Sash", "LightBelt", "Belt", "HeavyBelt", "PlatedBelt", "DemonhideSash", "SharkskinBelt", "MeshBelt", "BattleBelt", "WarBelt", "SpiderwebSash", "VampirefangBelt", "MithrilCoil", "TrollBelt", "ColossusGirdle"},
	"Circlets":           {"Circlet", "Coronet", "Tiara", "Diadem"},
	"AssassinKatars":     {"Katar", "WristBlade", "HatchetHands", "Cestus", "Claws", "BladeTalons", "ScissorsKatar", "Quhab", "WristSpike", "Fascia", "HandScythe", "GreaterClaws", "GreaterTalons", "ScissorsQuhab", "Suwayyah", "WristSword", "WarFist", "BattleCestus", "FeralClaws", "RunicTalons", "ScissorsSuwayyah"},
	"SorceressOrbs":      {"EagleOrb", "SacredGlobe", "SmokedSphere", "ClaspedOrb", "JaredsStone", "GlowingOrb", "CrystallineGlobe", "CloudySphere", "SparklingBall", "SwirlingCrystal", "HeavenlyStone", "EldritchOrb", "DemonHeart", "VortexOrb", "DimensionalShard"},
	"AmazonBows":         {"StagBow", "ReflexBow", "AshwoodBow", "CeremonialBow", "MatriarchalBow", "GrandMatronBow"},
	"AmazonSpears":       {"MaidenSpear", "MaidenPike", "CeremonialSpear", "CeremonialPike", "MatriarchalSpear", "MatriarchalPike"},
	"AmazonJavelins":     {"MaidenJavelin", "CeremonialJavelin", "MatriarchalJavelin"},
	"DruidHelms":         {"WolfHead", "HawkHelm", "Antlers", "FalconMask", "SpiritMask", "AlphaHelm", "GriffonHeaddress", "HuntersGuise", "SacredFeathers", "TotemicMask", "BloodSpirit", "SunSpirit", "EarthSpirit", "SkySpirit", "DreamSpirit"},
	"BarbarianHelms":     {"JawboneCap", "FangedHelm", "HornedHelm", "AssaultHelmet", "AvengerGuard", "JawboneVisor", "LionHelm", "RageMask", "SavageHelmet", "SlayerGuard", "CarnageHelm", "FuryVisor", "DestroyerHelm", "ConquerorCrown", "GuardianCrown"},
	"PaladinShields":     {"Targe", "Rondache", "HeraldicShield", "AerinShield", "CrownShield", "AkaranTarge", "AkaranRondache", "ProtectorShield", "GildedShield", "RoyalShield", "SacredTarge", "SacredRondache", "KurastShield", "ZakarumShield", "VortexShield"},
	"NecromancerShields": {"PreservedHead", "ZombieHead", "UnravellerHead", "GargoyleHead", "DemonHeadShield", "MummifiedTrophy", "FetishTrophy", "SextonTrophy", "CantorTrophy", "HierophantTrophy", "MinionSkull", "HellspawnSkull", "OverseerSkull", "SuccubusSkull", "BloodlordSkull"},
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
