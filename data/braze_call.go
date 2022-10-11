package data

import (
	"encoding/json"
	"io"

	"github.com/go-playground/validator/v10"
)

type BrazeCall struct {
	Braze_data    string `validate:"required"`
	Response      string `validate:"required"`
	Questionnaire string `validate:"required"`
}

type BrazeCalls []*BrazeCall

func GetBrazeCalls() BrazeCalls {
	return calls
}

var calls = BrazeCalls{
	{
		Braze_data:    "test braze_data",
		Response:      "test response",
		Questionnaire: "[test questionnaire]",
	},
	{
		Braze_data:    "{\"apiKey\":\"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb21wYW55X2lkIjoiNjJhMTM4ZGY4YjhjNzAwMDAxMTdjNGUzIiwiaWF0IjoxNjU0NzMzMDIzfQ.xJwYhQOi0_3kaVDXita5AEFZF9IjJy5vc7P3CtSuNAE\",\"surveyId\":\"62f28f6cc784b100012c3270\",\"platform\":\"ios\",\"userTraits\":{\"braze_id\":\"63377a04c1640017d45f6bd7\",\"external_id\":\"puce-colorado-9582\",\"email\":\"prakash.menon@glooko.com\",\"first_name\":\"Prakash\",\"last_name\":\"Menon\"}}",
		Response:      "{\"question1\":4,\"question2\":4,\"question3\":4,\"question4\":4,\"question5\":2,\"question6\":5,\"question7\":3,\"question8\":5,\"question9\":3,\"question10\":4,\"question11\":3,\"question12\":6,\"question13\":4,\"question14\":5,\"question15\":4,\"question16\":2,\"question17\":5}",
		Questionnaire: "[]",
	},
}

func (b BrazeCall) Validate() error {
	validtr := validator.New()
	return validtr.Struct(b)
}

func (b *BrazeCall) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(b)
}

func (b *BrazeCalls) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(b)
}

func AddBrazeCall(bc *BrazeCall) {
	calls = append(calls, bc)
}
