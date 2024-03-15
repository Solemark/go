package employee

type Employee struct {
	Email string
	Name  string
	Phone string
	Rate  float32
}

/*Get Employee name*/
func (emp *Employee) GetName() string {
	return emp.Name
}

/*Set Employee email*/
func (emp *Employee) SetEmail(email string) {
	emp.Email = email
}

/*Set Employee name*/
func (emp *Employee) SetName(name string) {
	emp.Name = name
}

/*Set Employee phone*/
func (emp *Employee) SetPhone(phone string) {
	emp.Phone = phone
}

/*Set Employee rate*/
func (emp *Employee) SetRate(rate float32) {
	emp.Rate = rate
}
