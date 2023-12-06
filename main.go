package main

import (
	"fmt"
	"strings"
)

func containsSubstring(placas []string, val string) bool {
	for _, item := range placas {
		if strings.EqualFold(item, val) {
			return true
		}
	}
	return false
}

func generatePlateValues(start, end string) []string {
	var possibleValues []string

	current := start

	for current != end {
		possibleValues = append(possibleValues, current)

		if current[2] != 'Z' {
			current = current[:2] + string(current[2]+1)
		} else {
			current = current[:1] + string(current[1]+1) + "A"
		}
	}

	possibleValues = append(possibleValues, end)
	return possibleValues
}

func main() {

	//Setando as placas de cada estado como uma matriz de strings
	plateListCeara := [][]string{
		{"NQL", "NRE"}, {"NUM", "NVF"}, {"HTX", "HZA"}, {"OCB", "OCU"},
		{"OHX", "OIQ"}, {"ORN", "OSV"}, {"OZA", "OZA"}, {"PMA", "POZ"},
		{"RIA", "RIN"}, {"SAN", "SBV"},
	}
	plateListMaranhao := [][]string{
		{"HOL", "HQE"}, {"NHA", "NHT"}, {"NMP", "NNI"}, {"NWS", "NXQ"},
		{"OIR", "OJQ"}, {"OXQ", "OXZ"}, {"PSA", "PTZ"}, {"ROA", "ROZ"},
	}
	plateListPiaui := [][]string{
		{"LVF", "LWQ"}, {"NHU", "NIX"}, {"ODU", "OEI"}, {"OUA", "OUE"},
		{"OVW", "OVY"}, {"PIA", "PIZ"}, {"QRN", "QRZ"}, {"RSG", "RST"},
	}

	var possibleValuesCeara []string
	for _, plates := range plateListCeara {
		possibleValuesCeara = append(possibleValuesCeara, generatePlateValues(plates[0], plates[1])...)
	}

	var possibleValuesMaranhao []string
	for _, plates := range plateListMaranhao {
		possibleValuesMaranhao = append(possibleValuesMaranhao, generatePlateValues(plates[0], plates[1])...)
	}

	var possibleValuesPiaui []string
	for _, plates := range plateListPiaui {
		possibleValuesPiaui = append(possibleValuesPiaui, generatePlateValues(plates[0], plates[1])...)
	}

	var plate string
	fmt.Print("Qual a placa? ")
	fmt.Scanln(&plate)
	plate = plate[0:3]

	switch {
	case containsSubstring(possibleValuesCeara, plate):
		fmt.Println("\nA Placa é do Ceará")
	case containsSubstring(possibleValuesMaranhao, plate):
		fmt.Println("\nA Placa é do Maranhão")
	case containsSubstring(possibleValuesPiaui, plate):
		fmt.Println("\nA Placa é do Piauí")
	default:
		fmt.Println("\nNão está no Ceará, Maranhão ou Piauí")
	}
}
