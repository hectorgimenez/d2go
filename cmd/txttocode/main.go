package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
	"text/template"
)

func main() {
	for _, file := range textFiles {
		err := generateFile(file.SourceFile, file.DestFile, file.Template)
		if err != nil {
			panic(err)
		}
	}
	err := generateItems()
	if err != nil {
		panic(err)
	}
}

func generateFile(sourcePath, destinationPath, tpl string) error {
	fileToRead, err := os.Open(sourcePath)
	if err != nil {
		return err
	}
	defer fileToRead.Close()

	fileScanner := bufio.NewScanner(fileToRead)
	fileScanner.Split(bufio.ScanLines)

	headers := make([]string, 0)
	fileContent := make([]map[string]string, 0)
	for fileScanner.Scan() {
		fields := strings.Split(fileScanner.Text(), "\t")
		if len(headers) == 0 {
			headers = fields
			continue
		}

		// Ignore if Code is empty
		if fields[1] == "" {
			continue
		}
		lineMap := make(map[string]string)
		for i, header := range headers {
			lineMap[header] = fields[i]
		}
		fileContent = append(fileContent, lineMap)
	}

	if len(fileContent) == 0 {
		return fmt.Errorf("error: no content found for file %s", sourcePath)
	}

	t := template.Must(template.New("tpl").Funcs(template.FuncMap{
		"replace": strings.ReplaceAll,
	}).Parse(tpl))

	file, err := os.Create(destinationPath)
	if err != nil {
		return err
	}
	defer file.Close()

	return t.Execute(file, fileContent)
}

func generateItems() error {
	weapons, err := os.Open("cmd/txttocode/txt/weapons.txt")
	if err != nil {
		return err
	}
	defer weapons.Close()
	armor, err := os.Open("cmd/txttocode/txt/armor.txt")
	if err != nil {
		return err
	}
	defer armor.Close()
	misc, err := os.Open("cmd/txttocode/txt/misc.txt")
	if err != nil {
		return err
	}
	defer misc.Close()

	files := []*os.File{weapons, armor, misc}

	itemsFileContent := ""

	itemID := 0
	for k, file := range files {
		var b bytes.Buffer
		scanner := bufio.NewScanner(file)

		headers := make([]string, 0)
		fileContent := make([]map[string]string, 0)
		for scanner.Scan() {
			fields := strings.Split(scanner.Text(), "\t")
			if len(headers) == 0 {
				headers = fields
				continue
			}

			// Ignore if Code is empty
			if fields[1] == "" {
				continue
			}

			lineMap := make(map[string]string)
			for i, header := range headers {
				lineMap[header] = fields[i]
			}

			lineMap["ID"] = fmt.Sprintf("%d", itemID)
			lineMap["name"] = strings.ReplaceAll(lineMap["name"], "Heirophant", "Hierophant")
			lineMap["name"] = strings.ReplaceAll(lineMap["name"], "Colossal", "Colossus")
			lineMap["name"] = strings.ReplaceAll(lineMap["name"], "Ancient Shield", "Kurast Shield")
			lineMap["name"] = strings.ReplaceAll(lineMap["name"], "Ornate Armor", "Ornate Plate")
			lineMap["name"] = strings.ReplaceAll(lineMap["name"], "Essense", "Essence")

			fieldsToCheck := []string{
				"minac", "maxac", "mindam", "maxdam", "2handmindam", "2handmaxdam",
				"minmisdam", "maxmisdam", "speed", "StrBonus", "DexBonus",
				"reqstr", "reqdex", "durability", "level", "gemsockets",
			}

			for _, field := range fieldsToCheck {
				if value, found := lineMap[field]; !found || value == "" {
					lineMap[field] = "0"
				}
			}

			if _, found := lineMap["normcode"]; !found {
				lineMap["normcode"] = ""
				lineMap["ubercode"] = ""
				lineMap["ultracode"] = ""
			}

			fileContent = append(fileContent, lineMap)
			itemID++
		}

		tpl := templateArmorAndMisc
		if k == 0 {
			tpl = templateWeapons
		}

		t := template.Must(template.New("tpl").Funcs(template.FuncMap{
			"replace": strings.ReplaceAll,
		}).Parse(tpl))

		err = t.Execute(&b, fileContent)
		if err != nil {
			return err
		}
		itemsFileContent += b.String()
	}

	itemsFileContent += "}"
	err = os.WriteFile("pkg/data/item/items.go", []byte(itemsFileContent), 0644)
	if err != nil {
		return err
	}
	return nil
}
