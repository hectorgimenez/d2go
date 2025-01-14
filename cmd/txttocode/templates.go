package main

const templateItemType = `// Code generated by cmd/txttocode. DO NOT EDIT.
// source: cmd/txttocode/txt/itemtypes.txt
package item

const (
{{- range $key, $value := . }}
	Type{{ replace $value.ItemType " " "" }} = "{{ $value.Code }}"
{{- end }}
)

var ItemTypes = map[string]Type{
{{- range $key, $value := . }}
    Type{{ replace $value.ItemType " " "" }}: {ID: {{ $key }}, Name: "{{ $value.ItemType }}", Code: "{{ $value.Code }}", Throwable: {{ if eq $value.Throwable "1" }}true{{ else }}false{{ end }}, Beltable: {{ if eq $value.Beltable "1" }}true{{ else }}false{{ end }}},
{{- end }}
}`

const templateLevels = `// Code generated by cmd/txttocode. DO NOT EDIT.
// source: cmd/txttocode/txt/levels.txt
package area

var Areas = map[ID]Area{
{{- range $key, $value := . }}
    {{ $key }}: {Name: "{{ $value.LevelName }}", ID: {{ $key }}},
{{- end }}
}`

const templateSkillDesc = `// Code generated by cmd/txttocode. DO NOT EDIT.
// source: cmd/txttocode/txt/skilldesc.txt
package skill

var Desc = map[ID]Description{
{{- range $key, $value := . }}
    {{ $key }}: {Page: {{ $value.SkillPage }}, Row: {{ $value.SkillRow }}, Column: {{ $value.SkillColumn }}, ListRow: {{ $value.ListRow }}, IconCel: {{ $value.IconCel }}},
{{- end }}
}`

const templateSkills = `// Code generated by cmd/txttocode. DO NOT EDIT.
// source: cmd/txttocode/txt/skills.txt
package skill

var Skills = map[ID]Skill{
{{- range $key, $value := . }}
    {{ index $value "*Id" }}: {Name: "{{ $value.skill }}", ID: {{ $key }}, LeftSkill: {{ if eq $value.leftskill "1" }}true{{ else }}false{{ end }}, RightSkill: {{ if eq $value.rightskill "1" }}true{{ else }}false{{ end }}},
{{- end }}
}`

const templateWeapons = `// Code generated by cmd/txttocode. DO NOT EDIT.
// source: cmd/txttocode/txt/weapons.txt, cmd/txttocode/txt/armor.txt, cmd/txttocode/txt/misc.txt
package item

var Desc = map[int]Description{
{{- range $key, $value := . }}
    {{ $value.ID }}: {Name: "{{ $value.name }}", ID: {{ $value.ID }}, Code: "{{ $value.code }}", NormalCode: "{{ $value.normcode }}", UberCode: "{{ $value.ubercode }}", UltraCode: "{{ $value.ultracode }}", InventoryWidth: {{ $value.invwidth }}, InventoryHeight: {{ $value.invheight }}, MinDefense: {{ $value.minac }}, MaxDefense: {{ $value.maxac }}, MinDamage: {{ $value.mindam }}, MaxDamage: {{ $value.maxdam }}, TwoHandMinDamage: {{ index $value "2handmindam" }}, TwoHandMaxDamage: {{ index $value "2handmaxdam" }}, MinMissileDamage: {{ $value.minmisdam }}, MaxMissileDamage: {{ $value.maxmisdam }}, Speed: {{ $value.speed }}, StrengthBonus: {{ $value.StrBonus }}, DexterityBonus: {{ $value.DexBonus }}, RequiredStrength: {{ $value.reqstr }}, RequiredDexterity: {{ $value.reqdex }}, Durability: {{ $value.durability }}, RequiredLevel: {{ $value.levelreq }}, MaxSockets: {{ $value.gemsockets }}, Type: "{{ $value.type }}"},
{{- end }}`

const templateArmorAndMisc = `
{{- range $key, $value := . }}
    {{ $value.ID }}: {Name: "{{ $value.name }}", ID: {{ $value.ID }}, Code: "{{ $value.code }}", NormalCode: "{{ $value.normcode }}", UberCode: "{{ $value.ubercode }}", UltraCode: "{{ $value.ultracode }}", InventoryWidth: {{ $value.invwidth }}, InventoryHeight: {{ $value.invheight }}, MinDefense: {{ $value.minac }}, MaxDefense: {{ $value.maxac }}, MinDamage: {{ $value.mindam }}, MaxDamage: {{ $value.maxdam }}, TwoHandMinDamage: {{ index $value "2handmindam" }}, TwoHandMaxDamage: {{ index $value "2handmaxdam" }}, MinMissileDamage: {{ $value.minmisdam }}, MaxMissileDamage: {{ $value.maxmisdam }}, Speed: {{ $value.speed }}, StrengthBonus: {{ $value.StrBonus }}, DexterityBonus: {{ $value.DexBonus }}, RequiredStrength: {{ $value.reqstr }}, RequiredDexterity: {{ $value.reqdex }}, Durability: {{ $value.durability }}, RequiredLevel: {{ $value.levelreq }}, MaxSockets: {{ $value.gemsockets }}, Type: "{{ $value.type }}"},
{{- end }}
`

const templateObjects = `// Code generated by cmd/txttocode. DO NOT EDIT.
// source: cmd/txttocode/txt/objects.txt
package object

var Desc = map[int]Description{
{{- range $key, $value := . }}
	{{ index $value "*ID" }}: {Name: "{{ $value.Name }}", ID: {{ index $value "*ID" }}, SizeX: {{ $value.SizeX }}, SizeY: {{ $value.SizeY }}, Left: {{ $value.Left }}, Top: {{ $value.Top }}, Width: {{ $value.Width }}, Height: {{ $value.Height }}, Yoffset: {{ $value.Yoffset }}, Xoffset: {{ $value.Xoffset }}, HasCollision: {{ if eq $value.HasCollision0 "1" }}true{{ else }}false{{ end }}},
{{- end }}
}`

const templateEntrances = `// Code generated by cmd/txttocode. DO NOT EDIT.
// source: cmd/txttocode/txt/lvlwarp.txt
package entrance

var Desc = map[int]Description{
{{- range $key, $value := . }}
    {{ $value.UniqueId }}: {ID: {{ $value.UniqueId }}, Name: "{{ $value.Name }}", SelectX: {{ $value.SelectX }}, SelectY: {{ $value.SelectY }}, SelectDX: {{ $value.SelectDX }}, SelectDY: {{ $value.SelectDY }}, OffsetX: {{ $value.OffsetX }}, OffsetY: {{ $value.OffsetY }}, Direction: "{{ $value.Direction }}"},
{{- end }}
}`
