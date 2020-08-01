package models

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

var Martabaks map[string]Martabak

func init() {
	Martabaks = make(map[string]Martabak)
	m := Martabak{ID: "ID1", Field1: "words", Field2: 100}
	Martabaks[m.ID] = m
	m = Martabak{ID: "ID2", Field1: "sentence", Field2: 200}
	Martabaks[m.ID] = m
}

type Martabak struct {
	ID     string `json:"id"`
	Field1 string `json:"field1"`
	Field2 int    `json:"field2"`
}

type MartabakRequest struct {
	Field1 string `json:"field1"`
	Field2 int    `json:"field2"`
}

func GetAllMartabak() []Martabak {
	mm := make([]Martabak, 0)
	for _, m := range Martabaks {
		mm = append(mm, m)
	}
	return mm
}

func CreateMartabak(mr MartabakRequest) Martabak {
	ID := strings.ToUpper(strconv.FormatInt(time.Now().Unix(), 16))
	mtb := Martabak{
		ID:     ID,
		Field1: mr.Field1,
		Field2: mr.Field2,
	}

	Martabaks[ID] = mtb
	return mtb
}

func FindMartabak(ID string) (m Martabak, err error) {
	if m, ok := Martabaks[ID]; ok {
		return m, nil
	}

	return m, errors.New("martabak not found")
}

func UpdateMartabak(mtb Martabak, mr MartabakRequest) Martabak {
	mtb.Field1 = mr.Field1
	mtb.Field2 = mr.Field2
	Martabaks[mtb.ID] = mtb

	return mtb
}

func DeleteMartabak(ID string) {
	delete(Martabaks, ID)
}
