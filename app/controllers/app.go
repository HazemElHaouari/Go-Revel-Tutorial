package controllers

import (
	"github.com/revel/revel"
	"Go-Revel-Tutorial/app/models"
	"github.com/revel/revel/cache"
	"time"
)

type App struct {
	*revel.Controller
}
//Fonction Index
func (c App) Index() revel.Result {
	greeting := "Ajouter nouveau Brainee"
	return c.Render(greeting)
}
//Fonction qui ajoute brainee dans la cache et redirige l'application vers Hello.html
func (c App) Hello(brainee models.Brainee) revel.Result {
	go cache.Set("brainee_"+brainee.BraineeId, brainee, 300000*time.Minute)
    return c.Render(brainee)
}
//Fonction qui affiche le Brainee ajoutée précedement avec son id et l'affiche dans la page Brainee.html
func (c App) Brainee(braineeId string) revel.Result {
	var brainee models.Brainee 
	cache.Get("brainee_"+braineeId, &brainee)
	return c.Render(brainee)
}
