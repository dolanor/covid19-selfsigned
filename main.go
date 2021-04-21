package main

import (
	"image"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/skip2/go-qrcode"
)

func main() {
	go func() {
		err := run()
		if err != nil {
			panic(err)
		}
	}()
	app.Main()
}

func run() error {
	code, err := qrcode.New(AttestationEncode(), qrcode.Highest)
	if err != nil {
		return err
	}
	img := code.Image(-5)
	_ = img
	gui(img)
	return nil
}

func gui(img image.Image) {
	w := app.NewWindow()
	if err := loop(w, "Attestation Dérogatoire COVID19", img); err != nil {
		log.Fatal(err)
	}
	os.Exit(0)
}

type C = layout.Context
type D = layout.Dimensions

func loop(w *app.Window, title string, img image.Image) error {
	th := material.NewTheme(gofont.Collection())
	var ops op.Ops

	for e := range w.Events() {
		switch e := e.(type) {
		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)
			in := layout.UniformInset(unit.Dp(8))
			layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Rigid(func(gtx C) D {
					return in.Layout(gtx, func(gtx C) D {
						t := material.Body1(th, title)
						return t.Layout(gtx)
					})
				}),
				layout.Flexed(1, func(gtx C) D {
					return in.Layout(gtx, func(gtx C) D {
						qrOp := paint.NewImageOp(img)
						imgWidget := widget.Image{Src: qrOp, Fit: widget.ScaleDown}
						return imgWidget.Layout(gtx)
					})
				}),
			)

			e.Frame(gtx.Ops)
		case system.DestroyEvent:
			return e.Err
		}
	}

	return nil
}

func AttestationEncode() string {
	return `Cree le: 22/04/2021 a 00h27;
Nom: Dupont;
Prenom: Camille;
Naissance: 01/01/1970;
Adresse: 999 avenue de France 75001 Paris;
Sortie: 22/04/2021 a 01:27;
Motifs: travail, sante, famille, convocation_demarches, animaux, travail, sante, famille, convocation_demarches, demenagement, achats_culte_culturel, sport;`
}
