package repository

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"godatabase"
	"godatabase/entity"
	"testing"
)

func TestCommentInsert(t *testing.T) {
	commentRepository := NewCommentRepository(godatabase.GetConnection())

	ctx := context.Background()
	comment := entity.Comment{
		Email:   "repository@test.com",
		Comment: "Test Repository",
	}
	insert, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}

	fmt.Println(insert)
}

func TestFindById(t *testing.T) {
	commentRepository := NewCommentRepository(godatabase.GetConnection())
	comment, err := commentRepository.FindById(context.Background(), 90)
	if err != nil {
		panic(err)
	}
	fmt.Println(comment)

}

func TestFindByAll(t *testing.T) {
	commentRepository := NewCommentRepository(godatabase.GetConnection())
	comments, err := commentRepository.FindAll(context.Background())
	if err != nil {
		panic(err)
	}
	for _, comments := range comments {
		fmt.Println(comments)
	}

}
