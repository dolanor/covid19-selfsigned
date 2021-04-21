package main

import (
	"bytes"
	"image/png"
	"os"

	"github.com/skip2/go-qrcode"
)

func main() {
	err := run()
	if err != nil {
		panic(err)
	}
}
func run() error {
	code, err := qrcode.New(AttestationEncode(), qrcode.Highest)
	if err != nil {
		return err
	}
	img := code.Image(-1)

	f, err := os.Create("lol.png")
	if err != nil {
		return err
	}
	defer f.Close()

	var buf bytes.Buffer
	_ = buf
	err = png.Encode(f, img)
	if err != nil {
		return err
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
