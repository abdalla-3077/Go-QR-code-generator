package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"github.com/yeqown/go-qrcode/v2"
	"github.com/yeqown/go-qrcode/writer/standard"
)

type pepper struct {
	Name string
	Code string
}

func main() {

	EmailValidate := func(input string) error {
		if !strings.Contains(input, "@") {
			return errors.New("invalid input: missing '@'")
		}
		return nil
	}

	PhoneValidate := func(input string) error {
		if len(input) < 8 {
			return errors.New("invalid input: small number")
		}
		return nil
	}

	TypePromt := promptui.Select{
		Label: "Select Type",
		Items: []string{"Email", "Phone", "URL", "Text"},
	}

	_, result, err := TypePromt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	if result == "Email" {
		Promt := promptui.Prompt{
			Label:    "Enter Your Email",
			Validate: EmailValidate,
		}

		result, err := Promt.Run()

		if err != nil {
			fmt.Printf("failed %v\n", err)
			return
		}

		QrCode("mailto:" + result)
	} else if result == "Phone" {

		peppers := []pepper{
			{Name: "Egypt", Code: "+20"},
			{Name: "Saudi Arabia", Code: "+966"},
			{Name: "United Arab Emirates", Code: "+971"},
			{Name: "Qatar", Code: "+974"},
			{Name: "Kuwait", Code: "+965"},
			{Name: "Bahrain", Code: "+973"},
			{Name: "Oman", Code: "+968"},
			{Name: "Jordan", Code: "+962"},
			{Name: "Lebanon", Code: "+961"},
			{Name: "Iraq", Code: "+964"},
			{Name: "Syria", Code: "+963"},
			{Name: "Yemen", Code: "+967"},
			{Name: "Sudan", Code: "+249"},
			{Name: "Libya", Code: "+218"},
			{Name: "Algeria", Code: "+213"},
			{Name: "Morocco", Code: "+212"},
			{Name: "Tunisia", Code: "+216"},
			{Name: "Palestine", Code: "+970"},
			{Name: "Turkey", Code: "+90"},
			{Name: "Iran", Code: "+98"},
			{Name: "Pakistan", Code: "+92"},
			{Name: "India", Code: "+91"},
			{Name: "Afghanistan", Code: "+93"},
			{Name: "China", Code: "+86"},
			{Name: "Japan", Code: "+81"},
			{Name: "South Korea", Code: "+82"},
			{Name: "Russia", Code: "+7"},
			{Name: "Germany", Code: "+49"},
			{Name: "France", Code: "+33"},
			{Name: "United Kingdom", Code: "+44"},
			{Name: "Italy", Code: "+39"},
			{Name: "Spain", Code: "+34"},
			{Name: "United States", Code: "+1"},
			{Name: "Canada", Code: "+1"},
			{Name: "Brazil", Code: "+55"},
			{Name: "Mexico", Code: "+52"},
			{Name: "South Africa", Code: "+27"},
			{Name: "Nigeria", Code: "+234"},
			{Name: "Indonesia", Code: "+62"},
			{Name: "Australia", Code: "+61"},
		}

		searcher := func(input string, index int) bool {
			pepper := peppers[index]
			name := strings.Replace(strings.ToLower(pepper.Name), " ", "", -1)
			input = strings.Replace(strings.ToLower(input), " ", "", -1)
			return strings.Contains(name, input)
		}

		TypePromt := promptui.Select{
			Label:    "Select Country",
			Items:    peppers,
			Searcher: searcher,
		}

		_, ContryResult, err := TypePromt.Run()

		if err != nil {
			fmt.Printf("failed %v\n", err)
			return
		}

		var selectedPepper pepper
		for _, pepper := range peppers {
			if strings.Contains(ContryResult, pepper.Name) {
				selectedPepper = pepper
				break
			}
		}
		Promt := promptui.Prompt{
			Label:    "Enter Your Number",
			Validate: PhoneValidate,
		}

		result, err := Promt.Run()

		if err != nil {
			fmt.Printf("failed %v\n", err)
			return
		}

		QrCode(`tel:` + selectedPepper.Code + result)
	}

	if result == "URL" {
		Promt := promptui.Prompt{
			Label: "Enter Your Url",
		}

		result, err := Promt.Run()

		if err != nil {
			fmt.Printf("failed %v\n", err)
			return
		}

		QrCode(result)
	}

	if result == "Text" {
		Promt := promptui.Prompt{
			Label: "Enter Your Text",
		}

		result, err := Promt.Run()

		if err != nil {
			fmt.Printf("failed %v\n", err)
			return
		}

		QrCode(result)
	}

}

func QrCode(inpt string) {

	prompt := promptui.Select{
		Label: "Select Qr Style",
		Items: []string{"Square", "Circle"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	qrc, err := qrcode.New(inpt)
	if err != nil {
		fmt.Printf("could not generate QRCode: %v", err)
		return
	}
	files, _ := os.ReadDir("./Qr")

	options := []standard.ImageOption{
		standard.WithBgColorRGBHex("#ffffff"),
		standard.WithFgColorRGBHex("#000000"),
	}
	if result == "Circle" {
		options = append(options, standard.WithCircleShape())
	}

	w, err := standard.New("./Qr/"+fmt.Sprint(len(files)+1)+".png", options...)
	if err != nil {
		fmt.Printf("standard.New failed: %v", err)
		return
	} else {
		color.Green("QRCode generated successfully With Name: " + fmt.Sprint(len(files)+1) + ".png")

	}

	// save file
	if err = qrc.Save(w); err != nil {
		fmt.Printf("could not save image: %v", err)
	}
}
