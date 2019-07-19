package main

import (
	"bytes"
	"context"
	db "jiaojiao/database"
	user "jiaojiao/srv/user/proto"
	"jiaojiao/utils"

	"github.com/jinzhu/gorm"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
)

type srvUser struct{}
type srvAvatar struct{}

/**
 * @apiDefine DBServerDown
 * @apiError (Error 500) DBServerDown can't connect to database server
 */

/**
 * @api {rpc} /rpc user.User.Create
 * @apiVersion 1.0.0
 * @apiGroup Service
 * @apiName user.User.Create
 * @apiDescription Create new user.
 *
 * @apiParam {string} studentId student id.
 * @apiParam {string} studentName student name.
 * @apiSuccess {int32} status -1 for invalid param <br> 1 for success <br> 2 for exist user
 * @apiSuccess {int32} userId created or existed userid
 * @apiUse DBServerDown
 */
func (a *srvUser) Create(ctx context.Context, req *user.UserCreateRequest, rsp *user.UserCreateResponse) error {
	if req.StudentId == "" || req.StudentName == "" {
		rsp.Status = user.UserCreateResponse_INVALID_PARAM
	} else {
		var usr db.User
		err := db.Ormer.Where("student_id = ?", req.StudentId).First(&usr).Error
		if gorm.IsRecordNotFoundError(err) {
			usr = db.User{
				UserName:    req.StudentName,
				AvatarId:    utils.GetStringConfig("srv_config", "default_avatar"),
				StudentId:   req.StudentId,
				StudentName: req.StudentName,
			}
			utils.LogContinue(db.Ormer.Create(&usr).Error, utils.Warning)
			rsp.Status = user.UserCreateResponse_SUCCESS
		} else if utils.LogContinue(err, utils.Warning) {
			return err
		} else {
			rsp.Status = user.UserCreateResponse_USER_EXIST
		}
		rsp.User = new(user.UserInfo)
		parseUser(&usr, rsp.User)
	}
	return nil
}

/**
 * @api {rpc} /rpc user.User.Query
 * @apiVersion 1.0.0
 * @apiGroup Service
 * @apiName user.User.Query
 * @apiDescription Query user info.
 *
 * @apiParam {int32} userId user id
 * @apiSuccess {int32} userId user id
 * @apiSuccess {string} userName user name
 * @apiSuccess {string} avatarId user avatar id
 * @apiSuccess {string} telephone user telephone
 * @apiSuccess {string} studentId student id
 * @apiSuccess {string} studentName student name
 * @apiSuccess {int32} status user status, 1 for normal <br> 2 for frozen
 * @apiSuccess {int32} role user role, 1 for user <br> 2 for admin
 * @apiUse DBServerDown
 */
func (a *srvUser) Query(ctx context.Context, req *user.UserQueryRequest, rsp *user.UserInfo) error {
	if req.UserId == 0 {
		return nil
	}
	usr := db.User{
		ID: req.UserId,
	}
	err := db.Ormer.First(&usr).Error
	if gorm.IsRecordNotFoundError(err) {
		return nil
	} else if utils.LogContinue(err, utils.Warning) {
		return err
	}
	parseUser(&usr, rsp)
	return nil
}

/**
 * @api {rpc} /rpc user.User.Update
 * @apiVersion 1.0.0
 * @apiGroup Service
 * @apiName user.User.Update
 * @apiDescription Update user info, only update provided field. If clearEmpty=1 and param support allow clear, clear the field when not provided.
 *
 * @apiParam {int32} userId user id
 * @apiParam {string} [userName] user name
 * @apiParam {string} [avatarId] user avatar id
 * @apiParam {string} [telephone] user telephone, allow clear
 * @apiParam {string} [studentId] student id
 * @apiParam {string} [studentName] student name
 * @apiParam {int32} [status] user status, 1 for normal <br> 2 for frozen
 * @apiParam {int32} [role] user role, 1 for user <br> 2 for admin
 * @apiParam {bool} clearEmpty=0 clear the empty field
 * @apiSuccess {int32} status -1 for invalid param <br> 1 for success <br> 2 for user not found
 * @apiUse DBServerDown
 */
func (a *srvUser) Update(ctx context.Context, req *user.UserInfo, rsp *user.UserUpdateResponse) error {
	if req.UserId == 0 {
		rsp.Status = user.UserUpdateResponse_INVALID_PARAM
		return nil
	}

	usr := db.User{
		ID: req.UserId,
	}
	err := db.Ormer.First(&usr).Error
	if err == nil {
		utils.AssignNotEmpty(&req.UserName, &usr.UserName)
		utils.AssignNotEmpty(&req.AvatarId, &usr.AvatarId)
		if req.ClearEmpty {
			usr.Telephone = req.Telephone
		} else {
			utils.AssignNotEmpty(&req.Telephone, &usr.Telephone)
		}
		utils.AssignNotEmpty(&req.StudentId, &usr.StudentId)
		utils.AssignNotEmpty(&req.StudentName, &usr.StudentName)
		utils.AssignNotZero((*int32)(&req.Status), &usr.Status)
		utils.AssignNotZero((*int32)(&req.Role), &usr.Role)
		err := db.Ormer.Save(&usr).Error
		if utils.LogContinue(err, utils.Warning) {
			return err
		}
		rsp.Status = user.UserUpdateResponse_SUCCESS
	} else if gorm.IsRecordNotFoundError(err) {
		rsp.Status = user.UserUpdateResponse_NOT_FOUND
		return nil
	} else {
		utils.Warning(err)
		return err
	}
	return nil
}

/**
 * @api {rpc} /rpc user.User.Find
 * @apiVersion 1.0.0
 * @apiGroup Service
 * @apiName user.User.Find
 * @apiDescription Find user(fuzzy).
 *
 * @apiParam {string} [userName] username
 * @apiParam {uint32} limit=100 row limit
 * @apiParam {uint32} offset=0 row offset
 * @apiSuccess {list} user see [User Service](#api-Service-user_User_Query)
 * @apiUse DBServerDown
 */
func (a *srvUser) Find(ctx context.Context, req *user.UserFindRequest, rsp *user.UserFindResponse) error {
	if req.Limit == 0 {
		req.Limit = 100
	}

	var res []*db.User
	err := db.Ormer.Where("user_name LIKE ?", "%"+req.UserName+"%").Limit(req.Limit).Offset(req.Offset).Find(&res).Error
	if utils.LogContinue(err, utils.Warning) {
		return err
	}
	for i, v := range res {
		rsp.User = append(rsp.User, new(user.UserInfo))
		parseUser(v, rsp.User[i])
	}
	return nil
}

func parseUser(s *db.User, d *user.UserInfo) {
	d.UserId = int32(s.ID)
	d.UserName = s.UserName
	d.AvatarId = s.AvatarId
	d.Telephone = s.Telephone
	d.StudentId = s.StudentId
	d.StudentName = s.StudentName
	d.Status = user.UserInfo_Status(s.Status)
	d.Role = user.UserInfo_Role(s.Role)
}

/**
 * @api {rpc} /rpc user.Avatar.Create
 * @apiVersion 1.0.0
 * @apiGroup Service
 * @apiName user.Avatar.Create
 * @apiDescription create user avatar and return avatarId.
 *
 * @apiParam {int32} userId user id
 * @apiParam {bytes} content binary content
 * @apiSuccess {int32} status -1 for invalid param <br> 1 for success <br> 2 for not found
 * @apiSuccess {int32} avatarId new avatar id
 * @apiUse DBServerDown
 */
func (a *srvAvatar) Create(ctx context.Context, req *user.AvatarCreateRequest, rsp *user.AvatarCreateResponse) error {
	if bytes.Equal(req.Content, []byte{0}) || req.UserId == 0 {
		rsp.Status = user.AvatarCreateResponse_INVALID_PARAM
	} else {
		bucket, err := gridfs.NewBucket(db.MongoDatabase)
		if utils.LogContinue(err, utils.Warning) {
			return err
		}

		id, err := bucket.UploadFromStream("", bytes.NewReader(req.Content))
		if utils.LogContinue(err, utils.Warning) {
			return err
		}

		usr := db.User{
			ID: req.UserId,
		}
		err = db.Ormer.First(&usr).Error
		if utils.LogContinue(err, utils.Warning) {
			return err
		}
		err = db.Ormer.Model(&usr).Update("avatarId", id.Hex()).Error
		if utils.LogContinue(err, utils.Warning) {
			return err
		}

		rsp.AvatarId = id.Hex()
		rsp.Status = user.AvatarCreateResponse_SUCCESS
	}
	return nil
}

func main() {
	db.InitORM("userdb", new(db.User))
	defer db.CloseORM()
	service := utils.InitMicroService("user")
	utils.LogPanic(user.RegisterUserHandler(service.Server(), new(srvUser)))
	utils.RunMicroService(service)
}
