package main

var textFiles = []TextFileDesc{
	{
		SourceFile: "cmd/txttocode/txt/itemtypes.txt",
		DestFile:   "pkg/data/item/itemtypes.go",
		Template:   templateItemType,
	},
	{
		SourceFile: "cmd/txttocode/txt/levels.txt",
		DestFile:   "pkg/data/area/areas.go",
		Template:   templateLevels,
	},
	{
		SourceFile: "cmd/txttocode/txt/skills.txt",
		DestFile:   "pkg/data/skill/skills.go",
		Template:   templateSkills,
	},
	{
		SourceFile: "cmd/txttocode/txt/skilldesc.txt",
		DestFile:   "pkg/data/skill/skilldesc.go",
		Template:   templateSkillDesc,
	},
	{
		SourceFile: "cmd/txttocode/txt/objects.txt",
		DestFile:   "pkg/data/object/objects.go",
		Template:   templateObjects,
	},
	{
		SourceFile: "cmd/txttocode/txt/lvlwarp.txt",
		DestFile:   "pkg/data/entrance/entrances.go",
		Template:   templateEntrances,
	},
	{
		SourceFile: "cmd/txttocode/txt/rareprefix.txt",
		DestFile:   "pkg/data/item/rareprefixes.go",
		Template:   templateRarePrefixes,
	},
	{
		SourceFile: "cmd/txttocode/txt/raresuffix.txt",
		DestFile:   "pkg/data/item/raresuffixes.go",
		Template:   templateRareSuffixes,
	},
	{
		SourceFile: "cmd/txttocode/txt/magicprefix.txt",
		DestFile:   "pkg/data/item/magicprefixes.go",
		Template:   templateMagicPrefixes,
	},
	{
		SourceFile: "cmd/txttocode/txt/magicsuffix.txt",
		DestFile:   "pkg/data/item/magicsuffixes.go",
		Template:   templateMagicSuffixes,
	},
	{
		SourceFile: "cmd/txttocode/txt/uniqueitems.txt",
		DestFile:   "pkg/data/item/uniqueitems.go",
		Template:   templateUniqueItems,
	},
	{
		SourceFile: "cmd/txttocode/txt/setitems.txt",
		DestFile:   "pkg/data/item/setitems.go",
		Template:   templateSetItems,
	},
}

type TextFileDesc struct {
	SourceFile string
	DestFile   string
	Template   string
}
