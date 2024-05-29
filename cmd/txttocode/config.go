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
}

type TextFileDesc struct {
	SourceFile string
	DestFile   string
	Template   string
}
