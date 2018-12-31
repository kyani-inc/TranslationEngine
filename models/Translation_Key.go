package models

import "strings"

type TranslationKeys struct {
	Language string `json:"language"`
	KeyMap map[string]map[string]map[string]string `json:"keys"`
}

func (ky *TranslationKeys) AddKey(key, keyValue string) {

	keys := parseDotSeperatedKeys(key)

	if ky.KeyMap == nil {
		ky.KeyMap = make(map[string]map[string]map[string]string)
	}
	if _, ok := ky.KeyMap[keys[0]]; !ok {

		ky.KeyMap[keys[0]] = make(map[string]map[string]string)
		ky.KeyMap[keys[0]][keys[1]] = make(map[string]string)
	}
	if _, ok := ky.KeyMap[keys[0]][keys[1]]; !ok {
		ky.KeyMap[keys[0]][keys[1]] = make(map[string]string)
	}
	if _, ok := ky.KeyMap[keys[0]][keys[1]][keys[2]]; !ok {
		ky.KeyMap[keys[0]][keys[1]][keys[2]] = keyValue
	}
}

func (ky *TranslationKeys) DeleteKey(key string) {

	keys := parseDotSeperatedKeys(key)

	delete(ky.KeyMap[keys[0]][keys[1]], keys[2])

	if len(ky.KeyMap[keys[0]][keys[1]]) < 2 {
		delete(ky.KeyMap[keys[0]], keys[1])
	}

	if len(ky.KeyMap[keys[0]]) < 2 {
		delete(ky.KeyMap, keys[0])
	}
}

func (ky *TranslationKeys) GetKey(key string) string {

	keys := parseDotSeperatedKeys(key)

	return ky.KeyMap[keys[0]][keys[1]][keys[2]]
}

func parseDotSeperatedKeys(key string) []string {
	keys := strings.Split(key, ".")
	if len(keys) == 3 {
		return keys
	}
	return []string{}
}
