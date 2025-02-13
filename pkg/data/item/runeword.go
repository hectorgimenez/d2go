package item

type RunewordName string

const (
	RunewordNone             RunewordName = ""
	RunewordAncientsPledge   RunewordName = "Ancients' Pledge"
	RunewordBeast            RunewordName = "Beast"
	RunewordBlack            RunewordName = "Black"
	RunewordBone             RunewordName = "Bone"
	RunewordBramble          RunewordName = "Bramble"
	RunewordBrand            RunewordName = "Brand"
	RunewordBreathOfTheDying RunewordName = "Breath of the Dying"
	RunewordCallToArms       RunewordName = "Call to Arms"
	RunewordChainsOfHonor    RunewordName = "Chains of Honor"
	RunewordChaos            RunewordName = "Chaos"
	RunewordCrescentMoon     RunewordName = "Crescent Moon"
	RunewordDeath            RunewordName = "Death"
	RunewordDelerium         RunewordName = "Delerium"
	RunewordDestruction      RunewordName = "Destruction"
	RunewordDoom             RunewordName = "Doom"
	RunewordDragon           RunewordName = "Dragon"
	RunewordDream            RunewordName = "Dream"
	RunewordDuress           RunewordName = "Duress"
	RunewordEdge             RunewordName = "Edge"
	RunewordEnigma           RunewordName = "Enigma"
	RunewordEnlightenment    RunewordName = "Enlightenment"
	RunewordEternity         RunewordName = "Eternity"
	RunewordExile            RunewordName = "Exile"
	RunewordFaith            RunewordName = "Faith"
	RunewordFamine           RunewordName = "Famine"
	RunewordFlickeringFlame  RunewordName = "Flickering Flame"
	RunewordFortitude        RunewordName = "Fortitude"
	RunewordFury             RunewordName = "Fury"
	RunewordGloom            RunewordName = "Gloom"
	RunewordGrief            RunewordName = "Grief"
	RunewordHandOfJustice    RunewordName = "Hand of Justice"
	RunewordHarmony          RunewordName = "Harmony"
	RunewordHeartOfTheOak    RunewordName = "Heart of the Oak"
	RunewordHolyThunder      RunewordName = "Holy Thunder"
	RunewordHonor            RunewordName = "Honor"
	RunewordIce              RunewordName = "Ice"
	RunewordInfinity         RunewordName = "Infinity"
	RunewordInsight          RunewordName = "Insight"
	RunewordKingsGrace       RunewordName = "King's Grace"
	RunewordKingslayer       RunewordName = "Kingslayer"
	RunewordLastWish         RunewordName = "Last Wish"
	RunewordLawbringer       RunewordName = "Lawbringer"
	RunewordLeaf             RunewordName = "Leaf"
	RunewordLionheart        RunewordName = "Lionheart"
	RunewordLore             RunewordName = "Lore"
	RunewordMalice           RunewordName = "Malice"
	RunewordMelody           RunewordName = "Melody"
	RunewordMemory           RunewordName = "Memory"
	RunewordMist             RunewordName = "Mist"
	RunewordMyth             RunewordName = "Myth"
	RunewordNadir            RunewordName = "Nadir"
	RunewordOath             RunewordName = "Oath"
	RunewordObedience        RunewordName = "Obedience"
	RunewordObsession        RunewordName = "Obsession"
	RunewordPassion          RunewordName = "Passion"
	RunewordPattern          RunewordName = "Pattern"
	RunewordPeace            RunewordName = "Peace"
	RunewordPhoenix          RunewordName = "Phoenix"
	RunewordPlague           RunewordName = "Plague"
	RunewordPride            RunewordName = "Pride"
	RunewordPrinciple        RunewordName = "Principle"
	RunewordPrudence         RunewordName = "Prudence"
	RunewordRadiance         RunewordName = "Radiance"
	RunewordRain             RunewordName = "Rain"
	RunewordRhyme            RunewordName = "Rhyme"
	RunewordRift             RunewordName = "Rift"
	RunewordSanctuary        RunewordName = "Sanctuary"
	RunewordSilence          RunewordName = "Silence"
	RunewordSmoke            RunewordName = "Smoke"
	RunewordSpirit           RunewordName = "Spirit"
	RunewordSplendor         RunewordName = "Splendor"
	RunewordStealth          RunewordName = "Stealth"
	RunewordSteel            RunewordName = "Steel"
	RunewordStone            RunewordName = "Stone"
	RunewordStrength         RunewordName = "Strength"
	RunewordTreachery        RunewordName = "Treachery"
	RunewordUnbendingWill    RunewordName = "Unbending Will"
	RunewordVenom            RunewordName = "Venom"
	RunewordVoiceOfReason    RunewordName = "Voice of Reason"
	RunewordWealth           RunewordName = "Wealth"
	RunewordWhite            RunewordName = "White"
	RunewordWind             RunewordName = "Wind"
	RunewordWisdom           RunewordName = "Wisdom"
	RunewordWrath            RunewordName = "Wrath"
	RunewordZephyr           RunewordName = "Zephyr"
	RunewordHustle           RunewordName = "Hustle"
	RunewordMosaic           RunewordName = "Mosaic"
	RunewordMetamorphosis    RunewordName = "Metamorphosis"
	RunewordGround           RunewordName = "Ground"
	RunewordTemper           RunewordName = "Temper"
	RunewordHearth           RunewordName = "Hearth"
	RunewordCure             RunewordName = "Cure"
	RunewordBulwark          RunewordName = "Bulwark"
)

var RunewordIDMap = map[int16]RunewordName{
	20507: RunewordAncientsPledge,
	20510: RunewordBeast,
	20512: RunewordBlack,
	20514: RunewordBone,
	20515: RunewordBramble,
	20516: RunewordBrand,
	20517: RunewordBreathOfTheDying,
	20519: RunewordCallToArms,
	20520: RunewordChainsOfHonor,
	20522: RunewordChaos,
	20523: RunewordCrescentMoon,
	20526: RunewordDeath,
	20528: RunewordDelerium,
	20531: RunewordDestruction,
	20532: RunewordDoom,
	20533: RunewordDragon,
	20535: RunewordDream,
	20536: RunewordDuress,
	20537: RunewordEdge,
	20539: RunewordEnigma,
	20540: RunewordEnlightenment,
	20542: RunewordEternity,
	20543: RunewordExile,
	20544: RunewordFaith,
	20545: RunewordFamine,
	20546: RunewordFlickeringFlame,
	20547: RunewordFortitude,
	20550: RunewordFury,
	20551: RunewordGloom,
	20553: RunewordGrief,
	20554: RunewordHandOfJustice,
	20555: RunewordHarmony,
	20557: RunewordHeartOfTheOak,
	20560: RunewordHolyThunder,
	20561: RunewordHonor,
	20565: RunewordIce,
	20566: RunewordInfinity,
	20568: RunewordInsight,
	20571: RunewordKingsGrace,
	20572: RunewordKingslayer,
	20575: RunewordLastWish,
	20577: RunewordLawbringer,
	20578: RunewordLeaf,
	20580: RunewordLionheart,
	20581: RunewordLore,
	20586: RunewordMalice,
	20587: RunewordMelody,
	20588: RunewordMemory,
	20589: RunewordMist,
	20592: RunewordMyth,
	20593: RunewordNadir,
	20596: RunewordOath,
	20597: RunewordObedience,
	20599: RunewordObsession,
	20600: RunewordPassion,
	20602: RunewordPattern,
	20603: RunewordPeace,
	20608: RunewordPhoenix,
	20611: RunewordPlague,
	20614: RunewordPride,
	20615: RunewordPrinciple,
	20617: RunewordPrudence,
	20621: RunewordRadiance,
	20622: RunewordRain,
	20625: RunewordRhyme,
	20626: RunewordRift,
	20627: RunewordSanctuary,
	20631: RunewordSilence,
	20633: RunewordSmoke,
	20635: RunewordSpirit,
	20636: RunewordSplendor,
	20638: RunewordStealth,
	20639: RunewordSteel,
	20642: RunewordStone,
	20644: RunewordStrength,
	20653: RunewordTreachery,
	20656: RunewordUnbendingWill,
	20659: RunewordVenom,
	20661: RunewordVoiceOfReason,
	20665: RunewordWealth,
	20667: RunewordWhite,
	20668: RunewordWind,
	20670: RunewordWisdom,
	20673: RunewordWrath,
	20675: RunewordZephyr,
	27360: RunewordHustle,
	27362: RunewordMosaic,
	27363: RunewordMetamorphosis,
	27364: RunewordGround,
	27365: RunewordTemper,
	27366: RunewordHearth,
	27367: RunewordCure,
	27368: RunewordBulwark,
}
