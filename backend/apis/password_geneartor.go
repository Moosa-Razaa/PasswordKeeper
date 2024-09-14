package apis

import (
	"backend/dtos/request"
	"encoding/json"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strings"
)

func ExtractBodyAsGeneratePasswordRequest(body []byte) (request.GeneratePasswordRequest, error) {
	var req request.GeneratePasswordRequest
	err := json.Unmarshal(body, &req)
	if err != nil {
		return request.GeneratePasswordRequest{}, err
	}

	return req, nil
}

func ValidateRequestBody(req request.GeneratePasswordRequest) bool {
	log.Printf("request body: %v", req)

	if req.Length <= 0 {
		return false
	}

	if req.Type < 0 || req.Type > 3 {
		return false
	}

	if req.IncludeSpecialCharacters && len(req.SpecialCharacters) == 0 {
		return false
	}

	return true
}

func GenerateRandomInteger(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func AlphabetsOnlyRandomPassword(req request.GeneratePasswordRequest, lowerCase string, upperCase string) string {
	var password strings.Builder
	var randomIndex int

	if req.IncludeSpecialCharacters {
		specialCharacter := req.SpecialCharacters
		for i := 0; i < req.Length; i++ {
			randomIndex = GenerateRandomInteger(0, 3)
			switch randomIndex {
			case 0:
				password.WriteByte(lowerCase[GenerateRandomInteger(0, 25)])
			case 1:
				password.WriteByte(upperCase[GenerateRandomInteger(0, 25)])
			case 2:
				password.WriteByte(specialCharacter[GenerateRandomInteger(0, len(specialCharacter)-1)])
			}
		}
	}

	for i := 0; i < req.Length; i++ {
		randomIndex = GenerateRandomInteger(0, 2)
		switch randomIndex {
		case 0:
			password.WriteByte(lowerCase[GenerateRandomInteger(0, 25)])
		case 1:
			password.WriteByte(upperCase[GenerateRandomInteger(0, 25)])
		}
	}

	return password.String()
}

func AlphanumericsOnlyRandomPassword(req request.GeneratePasswordRequest, lowerCase string, upperCase string, numbers string) string {
	var password strings.Builder
	var randomIndex int

	if req.IncludeSpecialCharacters {
		specialCharacter := req.SpecialCharacters
		for i := 0; i < req.Length; i++ {
			randomIndex = GenerateRandomInteger(0, 4)
			switch randomIndex {
			case 0:
				password.WriteByte(lowerCase[GenerateRandomInteger(0, 25)])
			case 1:
				password.WriteByte(upperCase[GenerateRandomInteger(0, 25)])
			case 2:
				password.WriteByte(numbers[GenerateRandomInteger(0, 9)])
			case 3:
				password.WriteByte(specialCharacter[GenerateRandomInteger(0, len(specialCharacter)-1)])
			}
		}
	}

	for i := 0; i < req.Length; i++ {
		randomIndex = GenerateRandomInteger(0, 3)
		switch randomIndex {
		case 0:
			password.WriteByte(lowerCase[GenerateRandomInteger(0, 25)])
		case 1:
			password.WriteByte(upperCase[GenerateRandomInteger(0, 25)])
		case 2:
			password.WriteByte(numbers[GenerateRandomInteger(0, 9)])
		}
	}

	return password.String()
}

func NumericsOnlyRandomPassword(req request.GeneratePasswordRequest, numbers string) string {
	var password strings.Builder
	var randomIndex int

	if req.IncludeSpecialCharacters {
		specialCharacter := req.SpecialCharacters
		for i := 0; i < req.Length; i++ {
			randomIndex = GenerateRandomInteger(0, 3)
			switch randomIndex {
			case 0:
				password.WriteByte(numbers[GenerateRandomInteger(0, 9)])
			case 1:
				password.WriteByte(specialCharacter[GenerateRandomInteger(0, len(specialCharacter)-1)])
			}
		}
	}

	for i := 0; i < req.Length; i++ {
		randomIndex = GenerateRandomInteger(0, 1)
		switch randomIndex {
		case 0:
			password.WriteByte(numbers[GenerateRandomInteger(0, 9)])
		}
	}

	return password.String()
}

func GenerateRandomPassword(req request.GeneratePasswordRequest) string {
	lowerCase := "abcdefghijklmnopqrstuvwxyz"
	upperCase := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers := "0123456789"

	switch req.Type {
	case request.AlphabetsOnly:
		return AlphabetsOnlyRandomPassword(req, lowerCase, upperCase)

	case request.AlphanumericsOnly:
		return AlphanumericsOnlyRandomPassword(req, lowerCase, upperCase, numbers)

	case request.NumericsOnly:
		return NumericsOnlyRandomPassword(req, numbers)
	}

	return "password"
}

func GeneratePassword(r *http.Request) (string, int) {
	log.Printf("request coming from %s", r.RemoteAddr)

	body, readAllError := io.ReadAll(r.Body)
	if readAllError != nil {
		return "Unable to read request body", http.StatusBadRequest
	}

	requestBody, extractingBodyError := ExtractBodyAsGeneratePasswordRequest(body)
	if extractingBodyError != nil {
		return "Unable to extract request body", http.StatusBadRequest
	}

	if !ValidateRequestBody(requestBody) {
		return "Invalid request body", http.StatusBadRequest
	}

	randomPassword := GenerateRandomPassword(requestBody)

	return randomPassword, http.StatusOK
}
