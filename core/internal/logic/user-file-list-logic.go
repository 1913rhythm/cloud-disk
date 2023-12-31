package logic

import (
	"cloud-disk/core/define"
	"cloud-disk/core/models"
	"context"
	"time"

	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileListLogic {
	return &UserFileListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileListLogic) UserFileList(req *types.UserFileListRequest, UserIdentity string) (resp *types.UserFileListReply, err error) {
	uf := make([]*types.UserFile, 0)
	var cnt int64
	resp = new(types.UserFileListReply)
	// 分页参数
	size := req.Size
	if size == 0 {
		size = define.PageSize
	}
	page := req.Page
	if page == 0 {
		page = 1
	}
	offset := (page - 1) * size

	// 查询用户文件列表
	// 查看执行的SQL语句
	//l.svcCtx.Engine.ShowSQL(true)
	err = l.svcCtx.Engine.Table("user_repository").Where("parent_id = ? AND user_identity = ?", req.Id, UserIdentity).
		Select("user_repository.id,user_repository.identity,user_repository.repository_identity,user_repository.ext,user_repository.name,repository_pool.path,repository_pool.size").
		Join("LEFT", "repository_pool", "user_repository.repository_identity = repository_pool.identity").
		Where("user_repository.deleted_at = ? OR user_repository.deleted_at IS NULL", time.Time{}.Format(define.DateTime)).
		Limit(size, offset).Find(&uf)
	if err != nil {
		return nil, err
	}
	// 查询用户文件总数
	cnt, err = l.svcCtx.Engine.Where("parent_id = ? AND user_identity = ?", req.Id, UserIdentity).Count(new(models.UserRepository))
	if err != nil {
		return nil, err
	}
	resp.List = uf
	resp.Count = cnt

	return
}
