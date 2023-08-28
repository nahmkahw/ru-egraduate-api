package students

import (
	"encoding/json"
	"time"
)

func (s *studentServices) GetStudentProfile(studentCode string) (studentProfileResponse *StudentProfileService, err error) {

	student := StudentProfileService{}

	STUDENT_CODE := studentCode

	key := STUDENT_CODE + "::profile"

	studentCache, err := s.redis_cache.Get(ctx, key).Result()
	if err == nil {

		_ = json.Unmarshal([]byte(studentCache), &student)
		return &student, nil
	}

	sp, err := s.studentRepo.GetStudentProfile(STUDENT_CODE)
	if err != nil {

		return studentProfileResponse, err
	}

	student = StudentProfileService{
		ENROLL_YEAR:          sp.ENROLL_YEAR,
		ENROLL_SEMESTER:      sp.ENROLL_SEMESTER,
		STD_CODE:             sp.STD_CODE,
		PRENAME_THAI_S:       sp.PRENAME_THAI_S,
		PRENAME_ENG_S:        sp.PRENAME_ENG_S,
		FIRST_NAME:           sp.FIRST_NAME,
		LAST_NAME:            sp.LAST_NAME,
		FIRST_NAME_ENG:       sp.FIRST_NAME_ENG,
		LAST_NAME_ENG:        sp.LAST_NAME_ENG,
		THAI_NAME:            sp.THAI_NAME,
		PLAN_NO:              sp.PLAN_NO,
		SEX:                  sp.SEX,
		REGINAL_NAME:         sp.REGINAL_NAME,
		SUBSIDY_NAME:         sp.SUBSIDY_NAME,
		STATUS_NAME_THAI:     sp.STATUS_NAME_THAI,
		BIRTH_DATE:           sp.BIRTH_DATE,
		STD_ADDR:             sp.STD_ADDR,
		ADDR_TEL:             sp.ADDR_TEL,
		JOB_POSITION:         sp.JOB_POSITION,
		STD_OFFICE:           sp.STD_OFFICE,
		OFFICE_TEL:           sp.OFFICE_TEL,
		DEGREE_NAME:          sp.DEGREE_NAME,
		BSC_DEGREE_NO:        sp.BSC_DEGREE_NO,
		BSC_DEGREE_THAI_NAME: sp.BSC_DEGREE_THAI_NAME,
		BSC_INSTITUTE_NO:     sp.BSC_INSTITUTE_NO,
		INSTITUTE_THAI_NAME:  sp.INSTITUTE_THAI_NAME,
		CK_CERT_NO:           sp.CK_CERT_NO,
		CHK_CERT_NAME_THAI:   sp.CHK_CERT_NAME_THAI,
		ENG_NAME:             sp.ENG_NAME,
	}

	studentProfileResponse = &student

	studentProfileJSON, _ := json.Marshal(studentProfileResponse)
	timeNow := time.Now()
	redisCacheStudentProfile := time.Unix(timeNow.Add(time.Minute*20).Unix(), 0)
	_ = s.redis_cache.Set(ctx, key, studentProfileJSON, redisCacheStudentProfile.Sub(timeNow)).Err()

	return studentProfileResponse, nil
}
