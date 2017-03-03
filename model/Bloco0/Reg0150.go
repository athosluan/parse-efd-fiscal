package Bloco0

import "github.com/jinzhu/gorm"

type Reg0150 struct {
	gorm.Model
	Reg string		`gorm:"type:varchar(4)"`
	CodPart string		`gorm:"type:varchar(60);unique_index"`
	Nome string		`gorm:"type:varchar(100)"`
	CodPais string		`gorm:"type:varchar(5)"`
	Cnpj string		`gorm:"type:varchar(15)"`
	Cpf string		`gorm:"type:varchar(11)"`
	Ie string		`gorm:"type:varchar(14)"`
	CodMun string		`gorm:"type:varchar(7)"`
	Suframa string		`gorm:"type:varchar(9)"`
	Endereco string		`gorm:"type:varchar(60)"`
	Num string		`gorm:"type:varchar(10)"`
	Compl string		`gorm:"type:varchar(60)"`
	Bairro string		`gorm:"type:varchar(60)"`

}

func (Reg0150) TableName() string {
	return "reg_0150"
}