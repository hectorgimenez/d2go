// Code generated by cmd/txttocode. DO NOT EDIT.
// source: cmd/txttocode/txt/itemtypes.txt
package item

const (
	TypeNone = "none"
	TypeShield = "shie"
	TypeArmor = "tors"
	TypeGold = "gold"
	TypeBowQuiver = "bowq"
	TypeCrossbowQuiver = "xboq"
	TypePlayerBodyPart = "play"
	TypeHerb = "herb"
	TypePotion = "poti"
	TypeRing = "ring"
	TypeElixir = "elix"
	TypeAmulet = "amul"
	TypeCharm = "char"
	TypeBoots = "boot"
	TypeGloves = "glov"
	TypeBook = "book"
	TypeBelt = "belt"
	TypeGem = "gem"
	TypeTorch = "torc"
	TypeScroll = "scro"
	TypeScepter = "scep"
	TypeWand = "wand"
	TypeStaff = "staf"
	TypeBow = "bow"
	TypeAxe = "axe"
	TypeClub = "club"
	TypeSword = "swor"
	TypeHammer = "hamm"
	TypeKnife = "knif"
	TypeSpear = "spea"
	TypePolearm = "pole"
	TypeCrossbow = "xbow"
	TypeMace = "mace"
	TypeHelm = "helm"
	TypeMissilePotion = "tpot"
	TypeQuest = "ques"
	TypeBodyPart = "body"
	TypeKey = "key"
	TypeThrowingKnife = "tkni"
	TypeThrowingAxe = "taxe"
	TypeJavelin = "jave"
	TypeWeapon = "weap"
	TypeMeleeWeapon = "mele"
	TypeMissileWeapon = "miss"
	TypeThrownWeapon = "thro"
	TypeComboWeapon = "comb"
	TypeAnyArmor = "armo"
	TypeAnyShield = "shld"
	TypeMiscellaneous = "misc"
	TypeSocketFiller = "sock"
	TypeSecondHand = "seco"
	TypeStavesAndRods = "rod"
	TypeMissile = "misl"
	TypeBlunt = "blun"
	TypeJewel = "jewl"
	TypeClassSpecific = "clas"
	TypeAmazonItem = "amaz"
	TypeBarbarianItem = "barb"
	TypeNecromancerItem = "necr"
	TypePaladinItem = "pala"
	TypeSorceressItem = "sorc"
	TypeAssassinItem = "assn"
	TypeDruidItem = "drui"
	TypeHandtoHand = "h2h"
	TypeOrb = "orb"
	TypeVoodooHeads = "head"
	TypeAuricShields = "ashd"
	TypePrimalHelm = "phlm"
	TypePelt = "pelt"
	TypeCloak = "cloa"
	TypeRune = "rune"
	TypeCirclet = "circ"
	TypeHealingPotion = "hpot"
	TypeManaPotion = "mpot"
	TypeRejuvPotion = "rpot"
	TypeStaminaPotion = "spot"
	TypeAntidotePotion = "apot"
	TypeThawingPotion = "wpot"
	TypeSmallCharm = "scha"
	TypeMediumCharm = "mcha"
	TypeLargeCharm = "lcha"
	TypeAmazonBow = "abow"
	TypeAmazonSpear = "aspe"
	TypeAmazonJavelin = "ajav"
	TypeHandtoHand2 = "h2h2"
	TypeMagicBowQuiv = "mboq"
	TypeMagicXbowQuiv = "mxbq"
	TypeChippedGem = "gem0"
	TypeFlawedGem = "gem1"
	TypeStandardGem = "gem2"
	TypeFlawlessGem = "gem3"
	TypePerfectGem = "gem4"
	TypeAmethyst = "gema"
	TypeDiamond = "gemd"
	TypeEmerald = "geme"
	TypeRuby = "gemr"
	TypeSapphire = "gems"
	TypeTopaz = "gemt"
	TypeSkull = "gemz"
	TypeSwordsandKnives = "blde"
	TypeSpearsandPolearms = "sppl"
)

var ItemTypes = map[string]Type{
    TypeNone: {ID: 0, Name: "None", Code: "none", Throwable: false, Beltable: false},
    TypeShield: {ID: 1, Name: "Shield", Code: "shie", Throwable: false, Beltable: false},
    TypeArmor: {ID: 2, Name: "Armor", Code: "tors", Throwable: false, Beltable: false},
    TypeGold: {ID: 3, Name: "Gold", Code: "gold", Throwable: false, Beltable: false},
    TypeBowQuiver: {ID: 4, Name: "Bow Quiver", Code: "bowq", Throwable: false, Beltable: false},
    TypeCrossbowQuiver: {ID: 5, Name: "Crossbow Quiver", Code: "xboq", Throwable: false, Beltable: false},
    TypePlayerBodyPart: {ID: 6, Name: "Player Body Part", Code: "play", Throwable: false, Beltable: false},
    TypeHerb: {ID: 7, Name: "Herb", Code: "herb", Throwable: false, Beltable: false},
    TypePotion: {ID: 8, Name: "Potion", Code: "poti", Throwable: false, Beltable: true},
    TypeRing: {ID: 9, Name: "Ring", Code: "ring", Throwable: false, Beltable: false},
    TypeElixir: {ID: 10, Name: "Elixir", Code: "elix", Throwable: false, Beltable: true},
    TypeAmulet: {ID: 11, Name: "Amulet", Code: "amul", Throwable: false, Beltable: false},
    TypeCharm: {ID: 12, Name: "Charm", Code: "char", Throwable: false, Beltable: false},
    TypeBoots: {ID: 13, Name: "Boots", Code: "boot", Throwable: false, Beltable: false},
    TypeGloves: {ID: 14, Name: "Gloves", Code: "glov", Throwable: false, Beltable: false},
    TypeBook: {ID: 15, Name: "Book", Code: "book", Throwable: false, Beltable: false},
    TypeBelt: {ID: 16, Name: "Belt", Code: "belt", Throwable: false, Beltable: false},
    TypeGem: {ID: 17, Name: "Gem", Code: "gem", Throwable: false, Beltable: false},
    TypeTorch: {ID: 18, Name: "Torch", Code: "torc", Throwable: false, Beltable: false},
    TypeScroll: {ID: 19, Name: "Scroll", Code: "scro", Throwable: false, Beltable: true},
    TypeScepter: {ID: 20, Name: "Scepter", Code: "scep", Throwable: false, Beltable: false},
    TypeWand: {ID: 21, Name: "Wand", Code: "wand", Throwable: false, Beltable: false},
    TypeStaff: {ID: 22, Name: "Staff", Code: "staf", Throwable: false, Beltable: false},
    TypeBow: {ID: 23, Name: "Bow", Code: "bow", Throwable: false, Beltable: false},
    TypeAxe: {ID: 24, Name: "Axe", Code: "axe", Throwable: false, Beltable: false},
    TypeClub: {ID: 25, Name: "Club", Code: "club", Throwable: false, Beltable: false},
    TypeSword: {ID: 26, Name: "Sword", Code: "swor", Throwable: false, Beltable: false},
    TypeHammer: {ID: 27, Name: "Hammer", Code: "hamm", Throwable: false, Beltable: false},
    TypeKnife: {ID: 28, Name: "Knife", Code: "knif", Throwable: false, Beltable: false},
    TypeSpear: {ID: 29, Name: "Spear", Code: "spea", Throwable: false, Beltable: false},
    TypePolearm: {ID: 30, Name: "Polearm", Code: "pole", Throwable: false, Beltable: false},
    TypeCrossbow: {ID: 31, Name: "Crossbow", Code: "xbow", Throwable: false, Beltable: false},
    TypeMace: {ID: 32, Name: "Mace", Code: "mace", Throwable: false, Beltable: false},
    TypeHelm: {ID: 33, Name: "Helm", Code: "helm", Throwable: false, Beltable: false},
    TypeMissilePotion: {ID: 34, Name: "Missile Potion", Code: "tpot", Throwable: true, Beltable: false},
    TypeQuest: {ID: 35, Name: "Quest", Code: "ques", Throwable: false, Beltable: false},
    TypeBodyPart: {ID: 36, Name: "Body Part", Code: "body", Throwable: false, Beltable: false},
    TypeKey: {ID: 37, Name: "Key", Code: "key", Throwable: false, Beltable: false},
    TypeThrowingKnife: {ID: 38, Name: "Throwing Knife", Code: "tkni", Throwable: true, Beltable: false},
    TypeThrowingAxe: {ID: 39, Name: "Throwing Axe", Code: "taxe", Throwable: true, Beltable: false},
    TypeJavelin: {ID: 40, Name: "Javelin", Code: "jave", Throwable: true, Beltable: false},
    TypeWeapon: {ID: 41, Name: "Weapon", Code: "weap", Throwable: false, Beltable: false},
    TypeMeleeWeapon: {ID: 42, Name: "Melee Weapon", Code: "mele", Throwable: false, Beltable: false},
    TypeMissileWeapon: {ID: 43, Name: "Missile Weapon", Code: "miss", Throwable: false, Beltable: false},
    TypeThrownWeapon: {ID: 44, Name: "Thrown Weapon", Code: "thro", Throwable: true, Beltable: false},
    TypeComboWeapon: {ID: 45, Name: "Combo Weapon", Code: "comb", Throwable: true, Beltable: false},
    TypeAnyArmor: {ID: 46, Name: "Any Armor", Code: "armo", Throwable: false, Beltable: false},
    TypeAnyShield: {ID: 47, Name: "Any Shield", Code: "shld", Throwable: false, Beltable: false},
    TypeMiscellaneous: {ID: 48, Name: "Miscellaneous", Code: "misc", Throwable: false, Beltable: false},
    TypeSocketFiller: {ID: 49, Name: "Socket Filler", Code: "sock", Throwable: false, Beltable: false},
    TypeSecondHand: {ID: 50, Name: "Second Hand", Code: "seco", Throwable: false, Beltable: false},
    TypeStavesAndRods: {ID: 51, Name: "Staves And Rods", Code: "rod", Throwable: false, Beltable: false},
    TypeMissile: {ID: 52, Name: "Missile", Code: "misl", Throwable: false, Beltable: false},
    TypeBlunt: {ID: 53, Name: "Blunt", Code: "blun", Throwable: false, Beltable: false},
    TypeJewel: {ID: 54, Name: "Jewel", Code: "jewl", Throwable: false, Beltable: false},
    TypeClassSpecific: {ID: 55, Name: "Class Specific", Code: "clas", Throwable: false, Beltable: false},
    TypeAmazonItem: {ID: 56, Name: "Amazon Item", Code: "amaz", Throwable: false, Beltable: false},
    TypeBarbarianItem: {ID: 57, Name: "Barbarian Item", Code: "barb", Throwable: false, Beltable: false},
    TypeNecromancerItem: {ID: 58, Name: "Necromancer Item", Code: "necr", Throwable: false, Beltable: false},
    TypePaladinItem: {ID: 59, Name: "Paladin Item", Code: "pala", Throwable: false, Beltable: false},
    TypeSorceressItem: {ID: 60, Name: "Sorceress Item", Code: "sorc", Throwable: false, Beltable: false},
    TypeAssassinItem: {ID: 61, Name: "Assassin Item", Code: "assn", Throwable: false, Beltable: false},
    TypeDruidItem: {ID: 62, Name: "Druid Item", Code: "drui", Throwable: false, Beltable: false},
    TypeHandtoHand: {ID: 63, Name: "Hand to Hand", Code: "h2h", Throwable: false, Beltable: false},
    TypeOrb: {ID: 64, Name: "Orb", Code: "orb", Throwable: false, Beltable: false},
    TypeVoodooHeads: {ID: 65, Name: "Voodoo Heads", Code: "head", Throwable: false, Beltable: false},
    TypeAuricShields: {ID: 66, Name: "Auric Shields", Code: "ashd", Throwable: false, Beltable: false},
    TypePrimalHelm: {ID: 67, Name: "Primal Helm", Code: "phlm", Throwable: false, Beltable: false},
    TypePelt: {ID: 68, Name: "Pelt", Code: "pelt", Throwable: false, Beltable: false},
    TypeCloak: {ID: 69, Name: "Cloak", Code: "cloa", Throwable: false, Beltable: false},
    TypeRune: {ID: 70, Name: "Rune", Code: "rune", Throwable: false, Beltable: false},
    TypeCirclet: {ID: 71, Name: "Circlet", Code: "circ", Throwable: false, Beltable: false},
    TypeHealingPotion: {ID: 72, Name: "Healing Potion", Code: "hpot", Throwable: false, Beltable: true},
    TypeManaPotion: {ID: 73, Name: "Mana Potion", Code: "mpot", Throwable: false, Beltable: true},
    TypeRejuvPotion: {ID: 74, Name: "Rejuv Potion", Code: "rpot", Throwable: false, Beltable: true},
    TypeStaminaPotion: {ID: 75, Name: "Stamina Potion", Code: "spot", Throwable: false, Beltable: true},
    TypeAntidotePotion: {ID: 76, Name: "Antidote Potion", Code: "apot", Throwable: false, Beltable: true},
    TypeThawingPotion: {ID: 77, Name: "Thawing Potion", Code: "wpot", Throwable: false, Beltable: true},
    TypeSmallCharm: {ID: 78, Name: "Small Charm", Code: "scha", Throwable: false, Beltable: false},
    TypeMediumCharm: {ID: 79, Name: "Medium Charm", Code: "mcha", Throwable: false, Beltable: false},
    TypeLargeCharm: {ID: 80, Name: "Large Charm", Code: "lcha", Throwable: false, Beltable: false},
    TypeAmazonBow: {ID: 81, Name: "Amazon Bow", Code: "abow", Throwable: false, Beltable: false},
    TypeAmazonSpear: {ID: 82, Name: "Amazon Spear", Code: "aspe", Throwable: false, Beltable: false},
    TypeAmazonJavelin: {ID: 83, Name: "Amazon Javelin", Code: "ajav", Throwable: true, Beltable: false},
    TypeHandtoHand2: {ID: 84, Name: "Hand to Hand 2", Code: "h2h2", Throwable: false, Beltable: false},
    TypeMagicBowQuiv: {ID: 85, Name: "Magic Bow Quiv", Code: "mboq", Throwable: false, Beltable: false},
    TypeMagicXbowQuiv: {ID: 86, Name: "Magic Xbow Quiv", Code: "mxbq", Throwable: false, Beltable: false},
    TypeChippedGem: {ID: 87, Name: "Chipped Gem", Code: "gem0", Throwable: false, Beltable: false},
    TypeFlawedGem: {ID: 88, Name: "Flawed Gem", Code: "gem1", Throwable: false, Beltable: false},
    TypeStandardGem: {ID: 89, Name: "Standard Gem", Code: "gem2", Throwable: false, Beltable: false},
    TypeFlawlessGem: {ID: 90, Name: "Flawless Gem", Code: "gem3", Throwable: false, Beltable: false},
    TypePerfectGem: {ID: 91, Name: "Perfect Gem", Code: "gem4", Throwable: false, Beltable: false},
    TypeAmethyst: {ID: 92, Name: "Amethyst", Code: "gema", Throwable: false, Beltable: false},
    TypeDiamond: {ID: 93, Name: "Diamond", Code: "gemd", Throwable: false, Beltable: false},
    TypeEmerald: {ID: 94, Name: "Emerald", Code: "geme", Throwable: false, Beltable: false},
    TypeRuby: {ID: 95, Name: "Ruby", Code: "gemr", Throwable: false, Beltable: false},
    TypeSapphire: {ID: 96, Name: "Sapphire", Code: "gems", Throwable: false, Beltable: false},
    TypeTopaz: {ID: 97, Name: "Topaz", Code: "gemt", Throwable: false, Beltable: false},
    TypeSkull: {ID: 98, Name: "Skull", Code: "gemz", Throwable: false, Beltable: false},
    TypeSwordsandKnives: {ID: 99, Name: "Swords and Knives", Code: "blde", Throwable: false, Beltable: false},
    TypeSpearsandPolearms: {ID: 100, Name: "Spears and Polearms", Code: "sppl", Throwable: false, Beltable: false},
}