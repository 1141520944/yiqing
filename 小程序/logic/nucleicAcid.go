package logic

import (
	"time"
	"webGin/dao/mysql"
	"webGin/models"
	"webGin/pkg/snowflake"
)

// AddInformationNucleicAcid 添加核酸
func AddInformationNucleicAcid(p *models.ParamNucleicAcidAdd) error {
	uid := snowflake.GenID()
	nucleicAcid := &models.Table_nucleicAcid{
		UserID:        p.UserID,
		Time:          time.Now().Unix(),
		Address:       p.Address,
		SuperID:       p.SuperID,
		NucleicAcidID: uid,
		// State:   1,
	}
	if err := mysql.InsertOneNucleicAcidInfo(nucleicAcid); err != nil {
		return err
	}
	//修改对应student的最近状态
	if err := mysql.UpdateNucleAcideState(nucleicAcid); err != nil {
		return err
	}
	return nil
}

// GetClassStudentNucleicAcid 查看核酸记录数-学生信息
func GetClassStudentNucleicAcid(p *models.ParamGetClassNucleicAcidNumber) (*models.NucleicAcidNumberInfo, error) {
	re := new(models.NucleicAcidNumberInfo)

	re.StudentName = p.StudentName
	re.StudentCollege = p.College
	re.StudentIDNumber = p.StudentNumber
	result, er := mysql.SelectoneByClassID(p.ClassID)
	if er != nil {
		return nil, er
	}
	re.ClassName = result.Name
	re.StudentGard = result.Grade
	C_number, C_Rnumber, err := mysql.SelectByClassID(p.ClassID)
	if err != nil {
		return nil, err
	}
	re.ClassNumber = C_number
	re.ClassRNumber = C_Rnumber
	I_number, I_Rnumber, err2 := mysql.SelectByInstructorId(p.InstructorID)
	if err2 != nil {
		return nil, err2
	}
	re.InstructorNumber = I_number
	re.InstructorRNumber = I_Rnumber
	S_number, S_Rnumber, err3 := mysql.SelectByUserID(p.UserID)
	if err3 != nil {
		return nil, err2
	}
	re.StudentNumber = S_number
	re.StudentRNumber = S_Rnumber
	return re, nil
}
