package file

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/suyuan32/simple-admin-file/ent/file"
	"github.com/suyuan32/simple-admin-file/ent/predicate"
	"go.uber.org/zap"
	"golang.org/x/exp/rand"
	"strconv"
	"time"

	"github.com/suyuan32/simple-admin-file/internal/svc"
	"github.com/suyuan32/simple-admin-file/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRandomImageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRandomImageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRandomImageLogic {
	return &GetRandomImageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *GetRandomImageLogic) GetRandomImage(req *types.RandomImageReq) (resp *types.RandomImageResp, err error) {
	// todo: add your logic here and delete this line
	defer func() {
		if err := recover(); err != nil {
			l.Logger.Error("GetRandomImage panic", zap.Any("err", err))
		}
	}()
	var predicates []predicate.File
	departmentIdStr, err := l.ctx.Value("deptId").(json.Number).Int64()
	departmentId := strconv.FormatInt(departmentIdStr, 10)

	predicates = append(predicates, file.CategoryID(int(req.CategoryId)))
	predicates = append(predicates, file.DepartmentIdEQ(departmentId))
	predicates = append(predicates, file.FileTypeEQ(2))
	count, err := l.svcCtx.DB.File.Query().Where(predicates...).Count(l.ctx)
	if err != nil {
		l.Logger.Error("GetRandomImage count error", zap.Error(err))
		return nil, errors.New("获取随机头像失败")
	}
	if count <= 0 {
		l.Logger.Error("GetRandomImage this category count empty", zap.Error(err))
		return nil, errors.New("该分类下无头像")
	}
	rand.Seed(uint64(time.Now().UnixNano()))
	randIndex := rand.Intn(count)
	first, err := l.svcCtx.DB.File.Query().Where(predicates...).Offset(randIndex).Limit(1).First(l.ctx)
	if err != nil {
		l.Logger.Error("GetRandomImage first error", zap.Error(err))
		return nil, errors.New("获取随机头像失败")
	}
	result := &types.RandomImageResp{}
	result.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	result.Data = types.RandomImageInfo{
		BaseUUIDInfo: types.BaseUUIDInfo{
			Id:        pointy.GetPointer(first.ID.String()),
			CreatedAt: pointy.GetPointer(first.CreatedAt.UnixMilli()),
			UpdatedAt: pointy.GetPointer(first.UpdatedAt.UnixMilli()),
		},
		UserUUID:   &first.UserID,
		Name:       &first.Name,
		PublicPath: pointy.GetPointer(l.svcCtx.Config.UploadConf.ServerURL + first.Path),
	}
	return result, nil
}
