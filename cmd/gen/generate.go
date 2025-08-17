package main

import (
	"fmt"
	"oa/internal/config"
	"oa/internal/database"
	"oa/internal/model/common"
	errUtil "oa/internal/util/err"

	"gorm.io/gen"
)

// gen生成代码
func main() {
	// 加载配置
	if err := config.LoadConfig(); err != nil {
		errUtil.HandlePanicError(&common.ServerErr{
			Msg: err.Error(),
		})
	}

	DB, _ := database.NewDatabase()

	// 创建 gen 实例
	g := gen.NewGenerator(gen.Config{
		OutPath:           "./internal/dal/query", // 查询代码输出路径
		ModelPkgPath:      "/model",               // 模型代码输出路径
		Mode:              gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
		FieldNullable:     true,
		FieldCoverable:    false,
		FieldSignable:     false,
		FieldWithIndexTag: true,
		FieldWithTypeTag:  true,
	})

	// 使用数据库连接
	g.UseDB(DB)

	// 应用基本配置，生成所有表的模型
	g.ApplyBasic(
		g.GenerateModel("users"),
		g.GenerateModel("roles"),
	)

	// 执行生成
	g.Execute()

	fmt.Println("代码生成完成！")
}
