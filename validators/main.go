package validators

import (
	"errors"
	"regexp"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	entranslations "github.com/go-playground/validator/v10/translations/en"

	"github.com/mdalbrid/models"
	"github.com/mdalbrid/utils/logger"
)

var regexpValidators = map[string]*regexp.Regexp{
	"uuid":     regexp.MustCompile(`(?i)^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}`),
	"email":    regexp.MustCompile(`^[^\\'"]{5,32}$`),
	"login":    regexp.MustCompile(`^{1,32}$`),
	"role":     regexp.MustCompile(`^[0-9a-zA-Z_-]{5,32}$`),
	"name":     regexp.MustCompile(`^[ 0-9a-zA-Z_-]{5,32}$`),
	"token":    regexp.MustCompile(`^[0-9a-zA-Z_\.-]+$`),
	"password": regexp.MustCompile("^([ ~:;!@#%&-_'\"/]|[0-9а-яА-Яa-zA-Z]|[^$*+?{}\\[\\]\\\\|()\\s]){5,64}$"),
}

var regexpErrorTexts = map[string]string{
	"uuid":     "must be valid UUID v5",
	"email":    "must be valid Email",
	"login":    "must be not less than 1 characters and no more than 32 characters",
	"role":     "must be not less than 5 characters and no more than 32 characters. Only contains alphanumeric characters, underscore and dash",
	"name":     "must be not less than 5 characters and no more than 32 characters. Only contains alphanumeric characters, underscore, space and dash",
	"token":    "must be valid Bearer Authorization token",
	"password": "must be not less than 5 characters and no more than 64 characters. Only contains alphanumeric characters and special chars: .,%!_#-",
}

var validate *validator.Validate
var uni *ut.UniversalTranslator
var trans ut.Translator

func init() {
	validate = validator.New()
	if err := validate.RegisterValidation("UUIDOrNull", validateUUIDOrNullTag, true); err != nil {
		logger.Error(err)
	}
	if err := validate.RegisterValidation("UUID", validateUUIDTag); err != nil {
		logger.Error(err)
	}
	if err := validate.RegisterValidation("regexOrNull", validateRegexOrNullTag, true); err != nil {
		logger.Error(err)
	}
	if err := validate.RegisterValidation("regex", validateRegexTag); err != nil {
		logger.Error(err)
	}

	enLang := en.New()
	uni = ut.New(enLang, enLang)
	trans, _ = uni.GetTranslator("en")

	if err := entranslations.RegisterDefaultTranslations(validate, trans); err != nil {
		logger.Error(err)
	}

	if err := validate.RegisterTranslation("regex", trans, translationsRegisterFn, translationFn); err != nil {
		logger.Error(err)
	}
}

func translationsRegisterFn(ut ut.Translator) error {
	return ut.Add("regex", "{0} {1}", true) // see universal-translator for details
}
func translationFn(ut ut.Translator, fe validator.FieldError) string {
	t, err := ut.T("regex", fe.Field(), regexpErrorTexts[fe.Param()])
	if err != nil {
		logger.Error(err)
	}
	return t
}

func validateUUIDOrNullTag(fl validator.FieldLevel) bool {
	if value, ok := fl.Field().Interface().(*model.UUID); ok && value == nil {
		return true
	}
	return validateUUIDTag(fl)
}
func validateUUIDTag(fl validator.FieldLevel) bool {
	if value, ok := fl.Field().Interface().(model.UUID); ok {
		return regexpValidators["uuid"].MatchString(value.String())
	}
	return false
}

func validateRegexOrNullTag(fl validator.FieldLevel) bool {
	if value, ok := fl.Field().Interface().(*string); ok && value == nil {
		return true
	}
	return validateRegexTag(fl)
}
func validateRegexTag(fl validator.FieldLevel) bool {
	if value, ok := fl.Field().Interface().(string); ok {
		return regexpValidators[fl.Param()].MatchString(value)
	}
	return false
}

// ValidateStruct - Do struct with data validation by struct validation tag
func ValidateStruct(structure interface{}) error {
	err := validate.Struct(structure)
	if err != nil {
		translated := make([]string, 0)
		errs := err.(validator.ValidationErrors)

		for _, e := range errs {
			translated = append(translated, e.Translate(trans))
		}
		return errors.New(strings.Join(translated, "\n"))
	}
	return nil
}
