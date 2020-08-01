// Controller Format: 
// parameterName      parameterSendingType     parameterDataType    required      descriptionComment
package controllers

import (
	"beego-api-server/models"
	"encoding/json"
	"github.com/astaxie/beego"
)

type MartabakController struct {
	beego.Controller
}

// @Description Create New Martabak
// @Param	content		body	models.MartabakRequest		true	"The martabak content to be created"
// @Success 200 {object} models.Martabak
// @router / [post]
func (m *MartabakController) Create() {
	var mr models.MartabakRequest
	err := json.Unmarshal(m.Ctx.Input.RequestBody, &mr)
	if err != nil {
		m.Data["json"] = "failed to create martabak: " + err.Error()
		m.ServeJSON()
		return
	}

	mtb := models.CreateMartabak(mr)

	m.Data["json"] = mtb
	m.ServeJSON()
}

// @Description Find Martabak by ID
// @Param	martabakID		path	string		true	"the martabak ID you want to get"
// @Success 200 {object} models.Martabak
// @router /:martabakID [get]
func (m *MartabakController) Find() {
	martabakID := m.Ctx.Input.Param(":martabakID")
	mtb, err := models.FindMartabak(martabakID)
	if err != nil {
		m.Data["json"] = "failed to find martabak with id: " + martabakID
		m.ServeJSON()
		return
	}

	m.Data["json"] = mtb
	m.ServeJSON()
}

// @Description Find All Martabak
// @Success 200 {object} []models.Martabak
// @router / [get]
func (m *MartabakController) FindAll() {
	mtb := models.GetAllMartabak()

	m.Data["json"] = mtb
	m.ServeJSON()
}

// @Description Update The Martabak
// @Param	martabakID		path	string						true		"the martabak ID you want to update"
// @Param	body			body 	models.MartabakRequest		true		"The content to be updated"
// @Success 200 {object} models.Martabak
// @router /:martabakID [put]
func (m *MartabakController) Update() {
	martabakID := m.Ctx.Input.Param(":martabakID")
	mtb, err := models.FindMartabak(martabakID)
	if err != nil {
		m.Data["json"] = "failed to find martabak with ID: " + martabakID
		m.ServeJSON()
		return
	}

	var mr models.MartabakRequest
	err = json.Unmarshal(m.Ctx.Input.RequestBody, &mr)
	if err != nil {
		m.Data["json"] = "failed to parse martabak update request with ID: " + martabakID
		m.ServeJSON()
		return
	}

	mtb = models.UpdateMartabak(mtb, mr)
	m.Data["json"] = mtb
	m.ServeJSON()
}

// @Description Delete The Martabak
// @Param	martabakID		path	string				true		"the martabak ID you want to delete"
// @Success 200 {string} martabak deleted successfully
// @router /:martabakID [delete]
func (m *MartabakController) Delete() {
	martabakID := m.Ctx.Input.Param(":martabakID")
	_, err := models.FindMartabak(martabakID)
	if err != nil {
		m.Data["json"] = "failed to find martabak with ID: " + martabakID
		m.ServeJSON()
		return
	}

	models.DeleteMartabak(martabakID)
	m.Data["json"] = "martabak deleted successfully"
	m.ServeJSON()
}
