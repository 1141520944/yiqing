package logic

import (
	"strconv"
	"time"
	"webGin/dao/mysql"
	"webGin/models"
	"webGin/pkg/jwt"
	"webGin/pkg/snowflake"
)

// AddInformation 添加用户记录
func AddInformation(p *models.ParamStudentInfoSignUp) error {
	if err := mysql.CheckStudentExistState(p.StudentNumber); err != nil {
		return err
	}
	uid := snowflake.GenID()
	student := &models.Table_user{
		Phone:         p.Phone,
		StudentName:   p.StudentName,
		College:       p.College,
		StudentNumber: p.StudentNumber,
		UserID:        uid,
		InstructorID:  p.InstructorID,
		ClassID:       p.ClassID,
		TodayTime:     time.Now().Unix(),
	}
	if err := mysql.InsertOneStudentInfo(student); err != nil {
		return err
	}
	return nil
}

// UpdateInformation 更新用户记录
func UpdateInformation(p *models.ParamStudentInfoUpdate) error {
	if err := mysql.CheckStudentExistState(p.StudentNumber); err != nil {
		return err
	}
	student := &models.Table_user{
		Phone:         p.Phone,
		StudentName:   p.StudentName,
		College:       p.College,
		StudentNumber: p.StudentNumber,
		UserID:        p.UserID,
		InstructorID:  p.InstructorID,
		ClassID:       p.ClassID,
	}
	if err := mysql.UpdateOneStudentInfo(student); err != nil {
		return err
	}
	return nil
}

// DeleteInformation 软删除用户记录
func DeleteInformation(uid int64) error {
	if err := mysql.SelectByUid(uid); err != nil {
		return err
	}
	student := &models.Table_user{
		UserID: uid,
	}
	if err := mysql.DeleteByUid(student); err != nil {
		return err
	}
	return nil
}

// DeleteCompletelyInformation  完全删除用户记录
func DeleteCompletelyInformation(uid int64) error {
	// if err := mysql.SelectByUid(uid); err != nil {
	// 	return err
	// }
	student := &models.Table_user{
		UserID: uid,
	}
	if err := mysql.DeleteCompleteByUid(student); err != nil {
		return err
	}
	return nil
}

// RegistInformationStudent 注册学生用户
func RegistInformationStudent(p *models.ParamStudentInfoRegist) (*models.Table_user, error) {
	if err := mysql.CheckStudentExistState(p.StudentNumber); err != nil {
		return nil, err
	} else if err := mysql.SelectByOpenIDStudent(p.OpenID); err != nil {
		//更新字段--依据open_id
		student := &models.Table_user{
			Phone:         p.Phone,
			StudentName:   p.StudentName,
			College:       p.College,
			StudentNumber: p.StudentNumber,
			InstructorID:  p.InstructorID,
			ClassID:       p.ClassID,
			OpenID:        p.OpenID,
		}
		re, err2 := mysql.UpdateByOpenIDStudent(student)
		if err2 != nil {
			return nil, err2
		} else {
			return re, nil
		}
	} else {
		uid := snowflake.GenID()
		student := &models.Table_user{
			Phone:         p.Phone,
			StudentName:   p.StudentName,
			College:       p.College,
			StudentNumber: p.StudentNumber,
			UserID:        uid,
			InstructorID:  p.InstructorID,
			ClassID:       p.ClassID,
			OpenID:        p.OpenID,
		}
		if err := mysql.RegistOneStudentInfo(student); err != nil {
			return nil, err
		} else {
			return student, nil
		}
	}
}

// IsOrNoRegistStudent 是否注册过
func IsOrNoRegistStudent(p string) (string, *models.Student, error) {
	user := &models.Table_user{
		OpenID: p,
	}
	st, err := mysql.SelectByStudentOpenid(user)
	if err != nil {
		return "", nil, err
	} else {
		token, err := jwt.GenToken(st.UserID, st.StudentName)
		if err != nil {
			return "", nil, err
		} else {
			student := &models.Student{
				CreatTime:     st.Model.CreatedAt,
				UpdateTime:    st.Model.UpdatedAt,
				StudentName:   st.StudentName,
				College:       st.College,
				Phone:         st.Phone,
				StudentNumber: st.StudentNumber,
				OpenID:        st.OpenID,
				UserID:        strconv.FormatInt(st.UserID, 10),
				ClassID:       strconv.FormatInt(st.ClassID, 10),
				InstructorID:  strconv.FormatInt(st.InstructorID, 10),
				TodayState:    st.TodayState,
				RoleID:        st.RoleID,
				TodayTime:     time.Unix(st.TodayTime, 0),
			}
			return token, student, nil
		}
	}
}

// SelectPageInformationNucleicAcid 查找核酸记录
func SelectPageInformationNucleicAcid(p *models.ParamPage) (int64, []*models.NucleicAcid, error) {
	total, re, err := mysql.SelectPageStudentInfo(p.Page, p.Limit, p.Uid)
	if err != nil {
		return 0, nil, err
	}
	result := make([]*models.NucleicAcid, 0, len(re))
	for _, v := range re {
		one := &models.NucleicAcid{
			CreatTime:     v.Model.CreatedAt,
			UpdateTime:    v.Model.UpdatedAt,
			Time:          strconv.FormatInt(v.Time, 10),
			Address:       v.Address,
			NucleicAcidID: strconv.FormatInt(v.NucleicAcidID, 10),
			UserID:        strconv.FormatInt(v.UserID, 10),
			SuperID:       strconv.FormatInt(v.SuperID, 10),
			State:         v.State,
		}
		result = append(result, one)
	}
	return total, result, nil
}

// SelectPageInformationStudent 查询学生
func SelectPageInformationStudent(p *models.ParamPage) (int64, []*models.Student, error) {
	total, re, err := mysql.SelectPageStudentClassInfo(p.Page, p.Limit, p.Uid)
	if err != nil {
		return 0, nil, err
	}
	result := make([]*models.Student, 0, len(re))
	for _, v := range re {
		one := &models.Student{
			CreatTime:     v.Model.CreatedAt,
			UpdateTime:    v.Model.UpdatedAt,
			Phone:         v.Phone,
			StudentName:   v.StudentName,
			College:       v.College,
			StudentNumber: v.StudentNumber,
			OpenID:        v.OpenID,
			UserID:        strconv.FormatInt(v.UserID, 10),
			ClassID:       strconv.FormatInt(v.ClassID, 10),
			InstructorID:  strconv.FormatInt(v.InstructorID, 10),
			TodayState:    v.TodayState,
			RoleID:        v.RoleID,
			TodayTime:     time.Unix(v.TodayTime, 0),
		}
		result = append(result, one)
	}
	return total, result, nil
}

// SelectOneByUidStudent
func SelectOneByUidStudent(p int64) (*models.Student, error) {
	v, err := mysql.SelectOneByUidStudent(p)
	if err != nil {
		return nil, err
	} else {
		one := &models.Student{
			CreatTime:     v.Model.CreatedAt,
			UpdateTime:    v.Model.UpdatedAt,
			Phone:         v.Phone,
			StudentName:   v.StudentName,
			College:       v.College,
			StudentNumber: v.StudentNumber,
			OpenID:        v.OpenID,
			UserID:        strconv.FormatInt(v.UserID, 10),
			ClassID:       strconv.FormatInt(v.ClassID, 10),
			InstructorID:  strconv.FormatInt(v.InstructorID, 10),
			TodayState:    v.TodayState,
			RoleID:        v.RoleID,
			TodayTime:     time.Unix(v.TodayTime, 0),
		}
		return one, nil
	}
}

// SelectAllClassInfo
func SelectAllClassInfo(p int64) (int64, []*models.Class, error) {
	i, re, err := mysql.SelectAllClassInfo(p)
	if err != nil {
		return 0, nil, err
	} else {
		result := make([]*models.Class, 0, len(re))
		for _, v := range re {
			one := &models.Class{
				CreatTime:    v.Model.CreatedAt,
				UpdateTime:   v.Model.UpdatedAt,
				InstructorID: strconv.FormatInt(v.InstructorID, 10),
				ClassID:      strconv.FormatInt(v.ClassID, 10),
				Number:       v.Number,
				StateNumber:  v.StateNumber,
				Grade:        v.Grade,
				College:      v.College,
				Name:         v.Name,
			}
			result = append(result, one)
		}
		return i, result, nil
	}
}

// GetAllInstructors
func GetAllInstructors() (int64, []*models.Instructor, error) {
	total, re, err := mysql.SelectAllInstructor()
	if err != nil {
		return 0, nil, err
	} else {
		result := make([]*models.Instructor, 0, len(re))
		for _, v := range re {
			one := &models.Instructor{
				CreatTime:    v.Model.CreatedAt,
				UpdateTime:   v.Model.UpdatedAt,
				InstructorID: strconv.FormatInt(v.InstructorID, 10),
				Username:     v.Username,
				Password:     v.Password,
				Name:         v.Name,
			}
			result = append(result, one)
		}
		return total, result, nil
	}
}
