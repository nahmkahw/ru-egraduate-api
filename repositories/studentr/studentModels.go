package studentr

import (
	"github.com/jmoiron/sqlx"
)

type (
	studentRepoDB struct {
		oracle_db *sqlx.DB
	}

	StudentProfileRepo struct {
		ENROLL_YEAR          string `db:"ENROLL_YEAR"`
		ENROLL_SEMESTER      string `db:"ENROLL_SEMESTER"`
		STD_CODE             string `db:"STD_CODE"`
		PRENAME_THAI_S       string `db:"PRENAME_THAI_S"`
		PRENAME_ENG_S        string `db:"PRENAME_ENG_S"`
		FIRST_NAME           string `db:"FIRST_NAME"`
		LAST_NAME            string `db:"LAST_NAME"`
		FIRST_NAME_ENG       string `db:"FIRST_NAME_ENG"`
		LAST_NAME_ENG        string `db:"LAST_NAME_ENG"`
		THAI_NAME            string `db:"THAI_NAME"`
		PLAN_NO              string `db:"PLAN_NO"`
		SEX                  string `db:"SEX"`
		REGINAL_NAME         string `db:"REGINAL_NAME"`
		SUBSIDY_NAME         string `db:"SUBSIDY_NAME"`
		STATUS_NAME_THAI     string `db:"STATUS_NAME_THAI"`
		BIRTH_DATE           string `db:"BIRTH_DATE"`
		STD_ADDR             string `db:"STD_ADDR"`
		ADDR_TEL             string `db:"ADDR_TEL"`
		JOB_POSITION         string `db:"JOB_POSITION"`
		STD_OFFICE           string `db:"STD_OFFICE"`
		OFFICE_TEL           string `db:"OFFICE_TEL"`
		DEGREE_NAME          string `db:"DEGREE_NAME"`
		BSC_DEGREE_NO        string `db:"BSC_DEGREE_NO"`
		BSC_DEGREE_THAI_NAME string `db:"BSC_DEGREE_THAI_NAME"`
		BSC_INSTITUTE_NO     string `db:"BSC_INSTITUTE_NO"`
		INSTITUTE_THAI_NAME  string `db:"INSTITUTE_THAI_NAME"`
		CK_CERT_NO           string `db:"CK_CERT_NO"`
		CHK_CERT_NAME_THAI   string `db:"CHK_CERT_NAME_THAI"`
		ENG_NAME             string `db:"ENG_NAME"`
	}

	RegisterRepo struct {
		ID                   string `db:"ID"`
		COURSE_YEAR          string `db:"COURSE_YEAR"`
		COURSE_SEMESTER      string `db:"COURSE_SEMESTER"`
		COURSE_NO            string `db:"COURSE_NO"`
		COURSE_METHOD        string `db:"COURSE_METHOD"`
		COURSE_METHOD_NUMBER string `db:"COURSE_METHOD_NUMBER"`
		DAY_CODE             string `db:"DAY_CODE"`
		TIME_CODE            string `db:"TIME_CODE"`
		ROOM_GROUP           string `db:"ROOM_GROUP"`
		INSTR_GROUP          string `db:"INSTR_GROUP"`
		COURSE_METHOD_DETAIL string `db:"COURSE_METHOD_DETAIL"`
		DAY_NAME_S           string `db:"DAY_NAME_S"`
		TIME_PERIOD          string `db:"TIME_PERIOD"`
		COURSE_ROOM          string `db:"COURSE_ROOM"`
		COURSE_INSTRUCTOR    string `db:"COURSE_INSTRUCTOR"`
		SHOW_RU30            string `db:"SHOW_RU30"`
		COURSE_CREDIT        string `db:"COURSE_CREDIT"`
		COURSE_PR            string `db:"COURSE_PR"`
		COURSE_COMMENT       string `db:"COURSE_COMMENT"`
		COURSE_EXAMDATE      string `db:"COURSE_EXAMDATE"`
	}

	RegisterAllRepo struct {
		YEAR      string `db:"YEAR"`
		SEMESTER  string `db:"SEMESTER"`
		COURSE_NO string `db:"COURSE_NO"`
		STD_CODE  string `db:"STD_CODE"`
		CREDIT    string `db:"CREDIT"`
	}

	PrepareTokenRepo struct {
		STD_CODE string `db:"STD_CODE"`
		STATUS   int    `db:"STATUS"`
	}

	StudentRepo struct {
		STD_CODE string `db:"STD_CODE"`
	}

	StudentRepoInterface interface {
		GetStudentProfile(studentCode string) (*StudentProfileRepo, error)
		GetRegisterAll(studentCode, courseYear string) (*[]RegisterAllRepo, error)
		GetRegister(studentCode, courseYear, courseSemester string) (*[]RegisterRepo, error)
		Authentication(studentCode string) (*PrepareTokenRepo, error)
		GetStudentAll() (*[]StudentRepo, error)
	}
)

func NewStudentRepo(oracle_db *sqlx.DB) StudentRepoInterface {
	return &studentRepoDB{oracle_db: oracle_db}
}
