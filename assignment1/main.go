package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

type Participants struct {
	Participants []Participant `json:"participants"`
}

type Participant struct {
	ID        string `json:"id"`
	Code      string `json:"student_code"`
	Nama      string `json:"student_name"`
	Alamat    string `json:"student_address"`
	Pekerjaan string `json:"student_occupation"`
	Alasan    string `json:"joining_reason"`
}

func searchParticipant(kode string, arrayOfParticipants []Participant) (Participant, error) {
	kode = strings.ToLower(kode)
	var theRightPartcipant Participant
	for _, choosenParticipant := range arrayOfParticipants {
		if strings.ToLower(choosenParticipant.Code) == kode {
			theRightPartcipant = choosenParticipant
			return theRightPartcipant, nil
		}
	}
	return theRightPartcipant, errors.New("Data Not Found")

}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func main() {
	jsonByte, err := os.ReadFile("participants.json")
	checkError(err)

	var data Participants
	checkError(json.Unmarshal(jsonByte, &data))

	participant, err := searchParticipant(os.Args[1], data.Participants)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("ID\t\t: %s\n", participant.ID)
	fmt.Printf("Nama\t\t: %s\n", participant.Nama)
	fmt.Printf("Alamat\t\t: %s\n", participant.Alamat)
	fmt.Printf("Pekerjaan\t: %s\n", participant.Pekerjaan)
	fmt.Printf("Alasan\t\t: %s", participant.Alasan)
}
