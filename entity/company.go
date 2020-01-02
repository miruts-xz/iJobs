package entity

type Company struct {
	ID                                                                 int
	CompanyName, Password, Email, Phone, Logo, Short_desc, Detail_info string
	Address                                                            Address
}
