package entity

//TODO error in importing db package, it created import cycle
type phone string

type User struct {
	Id          string `json:"id" `
	FirstName   string `json:"firstname" binding:"required" validate:"required, min=4, max=30"`
	SecondName  string `json:"secondname" binding:"required" validate:"required, min=4, max=30"`
	LastName    string `json:"lastname" binding:"required" validate:"required, min=4, max=30"`
	Email       string `json:"email" binding:"required,email" validate:"email"`
	Password    string `json:"password" binding:"required" validate:"gte=8"`
	ImageURL    string `json:"imageurl"  `
	PhoneNumber phone  `json:"phonenumber" binding:"required" validate:"phone"`
	Status      string `json:"status"  `
}

type LoginInfo struct {
	Email    string `json:"email" binding:"email"`
	Password string `json:"password" binding:"required"`
}

type UpdateUser struct {
	Id          string `json:"id" `
	FirstName   string `json:"firstname" validate:"required, min=4, max=30"`
	SecondName  string `json:"secondname" validate:"required, min=4, max=30"`
	LastName    string `json:"lastname" validate:"required, min=4, max=30"`
	Email       string `json:"email"  validate:"email"`
	Password    string `json:"password" validate:"gte=8"`
	ImageURL    string `json:"imageurl" `
	PhoneNumber phone  `json:"phonenumber"  validate:"phone"`
	Status      string `json:"status" `
}

//VALIDATION BELOW

// const tagName = "validate"

// var mailRe = regexp.MustCompile(`\A[\w+\-.]+@[a-z\d\-]+(\.[a-z]+)*\.[a-z]+\z`)
// var phoneRe = regexp.MustCompile("^\\+[1-9]?[0-9]{7,14}$")

// type Validator interface {
// 	Validate(interface{}) (bool, error)
// }

// type DefaultValidator struct {
// }

// func (v DefaultValidator) Validate(val interface{}) (bool, error) {
// 	return true, nil
// }

// type EmailValidator struct {
// }

// func (e EmailValidator) Validate(val interface{}) (bool, error) {
// 	if !mailRe.MatchString(val.(string)) {
// 		return false, fmt.Errorf("invalid email")
// 	}
// 	if result, err := db.Search(val.(string)); err != nil || len(result) != 0 {
// 		return false, fmt.Errorf("user by this email already exist")
// 	}

// 	return true, nil
// }

// //Phone validation

// type PhoneValidator struct {
// }

// var digitRe = regexp.MustCompile(`\d{8}`)

// func (p PhoneValidator) Validate(val interface{}) (bool, error) {
// 	if string(val.(string)[:2]) == "09" {
// 		if digitRe.MatchString(val.(string)[2:]) {
// 			return true, nil
// 		}
// 	}
// 	if string(val.(string)[:1]) == "9" {
// 		if digitRe.MatchString(val.(string)[1:]) {
// 			return true, nil
// 		}
// 	}
// 	if string(val.(string)[:4]) == "+251" {
// 		if string(val.(string)[4]) == "9" {
// 			if digitRe.MatchString(val.(string)[2:]) {
// 				return true, nil
// 			}
// 		}
// 	}
// 	if string(val.(string)[:3]) == "251" {
// 		if string(val.(string)[3]) == "9" {
// 			if digitRe.MatchString(val.(string)[2:]) {
// 				return true, nil
// 			}
// 		}
// 	}
// 	return false, nil

// }

// //Adress validation
// type AddressValidator struct {
// }

// var addressRe = regexp.MustCompile(`\D*`)

// func (a AddressValidator) Validate(val interface{}) (bool, error) {

// 	if addressRe.MatchString(val.(string)) {
// 		return true, nil
// 	}
// 	return false, nil

// }

// func getValidatorFromTag(tag string) Validator {
// 	args := strings.Split(tag, ",")
// 	switch args[0] {
// 	case "email":
// 		//validator :=
// 		//fmt.Sscanf(strings.Join(args[1:], ","), "min=%d,max=%d", &validator.Min, &validator.Max)
// 		return EmailValidator{}
// 	case "phone":
// 		//validator :=
// 		//fmt.Sscanf(strings.Join(args[1:], ","), "min=%d,max=%d", &validator.Min, &validator.Max)
// 		return PhoneValidator{}
// 	case "address":
// 		return AddressValidator{}
// 	}
// 	return DefaultValidator{}
// }

// func ValidateStruct(s interface{}) []error {
// 	errs := []error{} // ValueOf returns a Value representing the run-time data
// 	v := reflect.ValueOf(s)
// 	for i := 0; i < v.NumField(); i++ {
// 		// Get the field tag value
// 		tag := v.Type().Field(i).Tag.Get(tagName) // Skip if tag is not defined or ignored
// 		if tag == "" || tag == "-" {
// 			continue
// 		} // Get a validator that corresponds to a tag
// 		validator := getValidatorFromTag(tag)                    // Perform validation
// 		valid, err := validator.Validate(v.Field(i).Interface()) // Append error to results
// 		if !valid && err != nil {
// 			errs = append(errs, fmt.Errorf("%s %s", v.Type().Field(i).Name, err.Error()))
// 		}
// 	}
// 	return errs
// }
