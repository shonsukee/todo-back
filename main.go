package main

import (
	"todo/controller"
	"todo/db"
	"todo/repository"
	"todo/router"
	"todo/usecase"
	"todo/validator"
)

func main() {
	db := db.NewDB()
	userValidator := validator.NewUserValidator()
	taskValidator := validator.NewTaskValidator()
	// 外側でインスタンス化したものを引数としている
	userRepository := repository.NewUserRepository(db)
	taskRepository := repository.NewTaskRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	taskUsecase := usecase.NewTaskUsecase(taskRepository, taskValidator)
	userController := controller.NewUserController(userUsecase)
	taskController := controller.NewTaskController(taskUsecase)
	e := router.NewRouter(userController, taskController)

	// echoのインスタンスを利用してサーバ起動
	// err発生時，ログを出力して強制終了
	e.Logger.Fatal(e.Start(":8090"))
}

// ログインでパスワード違う時のエラーメッセージ要変更
