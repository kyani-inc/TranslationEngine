package translation_key

import (
	"fmt"
	"github.com/kyani-inc/TranslationEngine/helpers"
	"strings"
)

var TK *TranslationKeysList

type TranslationKeysList struct {
	List []TranslationKeys `json:"list"`
}

type TranslationKeys struct {
	Locale string                                  `json:"locale"`
	KeyMap map[string]map[string]map[string]string `json:"keys"`
}

func (tkl *TranslationKeysList) AddTranslationKeys(tks TranslationKeys) *TranslationKeysList {
	tkl.List = append(tkl.List, tks)
	return tkl
}

func (tkl *TranslationKeysList) DeleteTranslationKeys(tks TranslationKeys) *TranslationKeysList {
	for i, val := range tkl.List {
		if val.Locale == tks.Locale {
			tkl.List[i] = tkl.List[len(tkl.List)-1]
			tkl.List = tkl.List[:len(tkl.List)-1]
		}
	}
	return tkl
}

func (tkl *TranslationKeysList) Sync() *TranslationKeysList {
	TranslationKeysList := GetAll()

	isSame, sizes, err := helpers.Compare(TranslationKeysList, tkl)

	if err != nil {
		fmt.Println("Failed to Sync Because of Comparison Error", err)
	}

	if !isSame && sizes[1] > sizes[0] {
		for _, val := range tkl.List {
			Put(val)
		}
	}
	if !isSame && sizes[1] < sizes[0] {
		TK.UpdateFromDB()
	}

	return tkl
}

func (tkl *TranslationKeysList) UpdateFromDB() TranslationKeys {

	tkl_db := GetAll()

	for _, val := range tkl_db.List {
		i, current_tk := tkl.GetByLocale(val.Locale)
		isSame, sizes, err := helpers.Compare(val, current_tk)

		if err != nil {
			fmt.Println("Failed to Compare : ", val, current_tk)
		}

		if !isSame && sizes[0] > sizes[1] {
			if current_tk.Locale == "" {
				tkl.AddTranslationKeys(val)
			} else {
				tkl.List[i].KeyMap = val.KeyMap
			}
		}
	}
	return TranslationKeys{}
}

func (tkl *TranslationKeysList) GetByLocale(lang string) (int, TranslationKeys) {
	for i, val := range tkl.List {
		if val.Locale == lang {
			return i, val
		}
	}
	return 0, TranslationKeys{}
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
	if _, ok := ky.KeyMap[keys[0]][keys[1]][keys[2]]; !ok {
		ky.KeyMap[keys[0]][keys[1]][keys[2]] = keyValue
	}
}

func (ky *TranslationKeys) DeleteKey(key string) {

	keys := parseDotSeperatedKeys(key)

	delete(ky.KeyMap[keys[0]][keys[1]], keys[2])

	if len(ky.KeyMap[keys[0]][keys[1]]) < 1 {
		delete(ky.KeyMap[keys[0]], keys[1])
	}

	if len(ky.KeyMap[keys[0]]) < 1 {
		delete(ky.KeyMap, keys[0])
	}
}

func (ky *TranslationKeys) IsEmpty() bool {
	return ky.Locale == ""
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
