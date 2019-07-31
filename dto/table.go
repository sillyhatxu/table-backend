package dto

import "fmt"

type AddDTO struct {
	Content string `json:"content"`
}

func (add AddDTO) String() string {
	return fmt.Sprintf("AddDTO{ content: %s }", add.Content)
}

type TableOne struct {
	UserName            string `json:"user_name"`
	Gender              string `json:"gender"`
	Age                 string `json:"age"`
	Education           string `json:"education"`
	Identity            string `json:"identity"`
	PhoneNumber         string `json:"phone_number"`
	Email               string `json:"email"`
	IdentityAddress     string `json:"identity_address"`
	Work                string `json:"work"`
	WorkAddress         string `json:"work_address"`
	WorkPhone           string `json:"work_phone"`
	ReasonForApplying   string `json:"reason_for_applying"`
	BankName            string `json:"bank_name"`
	ReceiveAddress      string `json:"receive_address"`
	ApplyTime           string `json:"apply_time"`
	ReferrerName        string `json:"referrer_name"`
	ReferrerPhoneNumber string `json:"referrer_phone_number"`
	Remark              string `json:"remark"`
}
