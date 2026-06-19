package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/472893749723489727432hjsdjkgf/ai-hack/configs"
	"github.com/472893749723489727432hjsdjkgf/ai-hack/internal/domain"
	"github.com/472893749723489727432hjsdjkgf/ai-hack/internal/repository/postgres"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	dbUrl := configs.GetDbUrl()

	fmt.Println(dbUrl)
	ctx,cannel := context.WithTimeout(context.Background(),5*time.Second)
	defer cannel()
	pool,err := pgxpool.New(ctx,dbUrl)
	if err != nil{
		fmt.Println("Не удалось подключиться к бд: %w",err)
	}
	defer pool.Close()
	if err := pool.Ping(ctx); err != nil{
		log.Fatalf("Ошибка при пинге: %v",err)
	}
	log.Println("Успешное подключение!")
	userRepo := postgres.NewPostgresUserRepository(pool)
	testUser := &domain.User {
		UserName: "Ivan",
		Email: "test@gmail.com",
		Password: "5667",
	}
	log.Println("Добавление данных...")
	err = userRepo.CreateNewUserDB(ctx,testUser)
	if err != nil{
		log.Fatalf("Ошибка при добавлении данных: %v",err)
	}
	log.Println("Данные добавлены...")
	testCreds := &domain.Credentials {
		UserName: "Ivan",
		Password: "5667",
	}
	invalidCreds := &domain.Credentials {
		UserName: "Error",
		Password: "1234",
	}
	trueExists,_ := userRepo.CheckExistsUserDB(ctx,testCreds)
	if trueExists{
		log.Println("Поиск пользователя успешна")
	}
	falseExists,_ := userRepo.CheckExistsUserDB(ctx,invalidCreds)
	if !falseExists{
		log.Fatalf("Ошибка %v",err)
	}


	
}
