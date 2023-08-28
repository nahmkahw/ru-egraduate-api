package students

import (
	"RU-Smart-Workspace/ru-smart-api/repositories/studentr"
	"context"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

type (
	studentServices struct {
		studentRepo studentr.StudentRepoInterface
		redis_cache *redis.Client
	}

	AuthenPlayload struct {
		Std_code      string `json:"std_code"`
		Refresh_token string `json:"refresh_token"`
	}

	RegisterPlayload struct {
		Std_code        string `json:"std_code"`
		Course_year     string `json:"course_year"`
		Course_semester string `json:"course_semester"`
	}

	AuthenPlayloadRedirect struct {
		Std_code     string `json:"std_code"`
		Access_token string `json:"access_token"`
	}

	AuthenTestPlayload struct {
		Std_code string `json:"std_code"`
	}

	AuthenServicePlayload struct {
		ServiveId string `json:"service_id"`
	}

	TokenResponse struct {
		AccessToken  string `json:"accessToken"`
		RefreshToken string `json:"refreshToken"`
		IsAuth       bool   `json:"isAuth"`
		Message      string `json:"message"`
		StatusCode   int    `json:"status_code"`
	}

	TokenRedirectResponse struct {
		IsAuth     bool   `json:"isAuth"`
		Message    string `json:"message"`
		StdCode    string `json:"std_code"`
		StatusCode int    `json:"status_code"`
	}

	// claims คือข้อมูลที่อยู่ในส่วน Payload ของ Token
	// -iss (issuer) : เว็บหรือบริษัทเจ้าของ token
	// -sub (subject) : subject ของ token
	// -aud (audience) : ผู้รับ token
	// -exp (expiration time) : เวลาหมดอายุของ token
	// -nbf (not before) : เป็นเวลาที่บอกว่า token จะเริ่มใช้งานได้เมื่อไหร่
	// -iat (issued at) : ใช้เก็บเวลาที่ token นี้เกิดปัญหา
	// -jti (JWT id) : เอาไว้เก็บไอดีของ JWT แต่ละตัวนะครับ
	// -name (Full name) : เอาไว้เก็บชื่อ
	ClaimsToken struct {
		Issuer              string `json:"issuer"`
		Subject             string `json:"subject"`
		Role                string `json:"role"`
		ExpiresAccessToken  string `json:"expires_access_token"`
		ExpiresRefreshToken string `json:"expiration_refresh_token"`
	}

	StudentProfileService struct {
		ENROLL_YEAR          string `json:"ENROLL_YEAR"`
		ENROLL_SEMESTER      string `json:"ENROLL_SEMESTER"`
		STD_CODE             string `json:"STD_CODE"`
		PRENAME_THAI_S       string `json:"PRENAME_THAI_S"`
		PRENAME_ENG_S        string `json:"PRENAME_ENG_S"`
		FIRST_NAME           string `json:"FIRST_NAME"`
		LAST_NAME            string `json:"LAST_NAME"`
		FIRST_NAME_ENG       string `json:"FIRST_NAME_ENG"`
		LAST_NAME_ENG        string `json:"LAST_NAME_ENG"`
		THAI_NAME            string `json:"THAI_NAME"`
		PLAN_NO              string `json:"PLAN_NO"`
		SEX                  string `json:"SEX"`
		REGINAL_NAME         string `json:"REGINAL_NAME"`
		SUBSIDY_NAME         string `json:"SUBSIDY_NAME"`
		STATUS_NAME_THAI     string `json:"STATUS_NAME_THAI"`
		BIRTH_DATE           string `json:"BIRTH_DATE"`
		STD_ADDR             string `json:"STD_ADDR"`
		ADDR_TEL             string `json:"ADDR_TEL"`
		JOB_POSITION         string `json:"JOB_POSITION"`
		STD_OFFICE           string `json:"STD_OFFICE"`
		OFFICE_TEL           string `json:"OFFICE_TEL"`
		DEGREE_NAME          string `json:"DEGREE_NAME"`
		BSC_DEGREE_NO        string `json:"BSC_DEGREE_NO"`
		BSC_DEGREE_THAI_NAME string `json:"BSC_DEGREE_THAI_NAME"`
		BSC_INSTITUTE_NO     string `json:"BSC_INSTITUTE_NO"`
		INSTITUTE_THAI_NAME  string `json:"INSTITUTE_THAI_NAME"`
		CK_CERT_NO           string `json:"CK_CERT_NO"`
		CHK_CERT_NAME_THAI   string `json:"CHK_CERT_NAME_THAI"`
		ENG_NAME             string `json:"ENG_NAME"`
	}

	RegisterResponse struct {
		STUDENT_CODE    string                   `json:"std_code"`
		COURSE_YEAR     string                   `json:"course_year"`
		COURSE_SEMESTER string                   `json:"course_semester"`
		REGISTER        []RegisterResponseFromDB `json:"register"`
	}

	StudentResponse struct {
		STUDENT_CODE string `json:"std_code"`
	}

	RegisterResponseFromDB struct {
		ID                   string `json:"id"`
		COURSE_YEAR          string `json:"course_year"`
		COURSE_SEMESTER      string `json:"course_semester"`
		COURSE_NO            string `json:"course_no"`
		COURSE_METHOD        string `json:"course_method"`
		COURSE_METHOD_NUMBER string `json:"course_method_number"`
		DAY_CODE             string `json:"day_code"`
		TIME_CODE            string `json:"time_code"`
		ROOM_GROUP           string `json:"room_group"`
		INSTR_GROUP          string `json:"instr_group"`
		COURSE_METHOD_DETAIL string `json:"course_method_detail"`
		DAY_NAME_S           string `json:"day_name_s"`
		TIME_PERIOD          string `json:"time_period"`
		COURSE_ROOM          string `json:"course_room"`
		COURSE_INSTRUCTOR    string `json:"course_instructor"`
		SHOW_RU30            string `json:"show_ru30"`
		COURSE_CREDIT        string `json:"course_credit"`
		COURSE_PR            string `json:"course_pr"`
		COURSE_COMMENT       string `json:"course_comment"`
		COURSE_EXAMDATE      string `json:"course_examdate"`
	}

	StudentServicesInterface interface {
		Authentication(stdCode string) (*TokenResponse, error)
		AuthenticationService(service_id string) (*TokenResponse, error)
		AuthenticationRedirect(stdCode, accessToken string) (*TokenRedirectResponse, error)
		RefreshAuthentication(refreshToken, stdCode string) (*TokenResponse, error)
		Unauthorization(token string) bool
		CheckExistsToken(token string) bool
		GetStudentProfile(stdCode string) (*StudentProfileService, error)
		GetRegister(studentCode, courseYear, courseSemester string) (*RegisterResponse, error)
		GetRegisterAll(studentCode, courseYear string) (*RegisterAllResponse, error)

		GetStudentAll() (*[]StudentResponse, error)
	}
)

func NewStudentServices(studentRepo studentr.StudentRepoInterface, redis_cache *redis.Client) StudentServicesInterface {
	return &studentServices{
		studentRepo: studentRepo,
		redis_cache: redis_cache,
	}
}
