package services

import (
	"fmt"
	"log"
	"regexp"

	"github.com/JscorpTech/jst-go/go-generate/internal/domain"
	"github.com/JscorpTech/jst-go/go-generate/internal/providers"
	"github.com/charmbracelet/huh"
)

type aicService struct {
	GitService *gitService
}

func NewAicService() domain.AicServicePort {
	return &aicService{
		GitService: NewGitService(),
	}
}

func (a *aicService) GenerateComment() {
	provider := providers.NewPizzaGptProvider()
	changes, err := a.GitService.Changes()
	if err != nil {
		log.Fatal(err.Error())
	}
	result := provider.Generate("Sen git o'zgarishlariga comment yozishing kerak. Javobing faqat quyidagi formatda bo'lishi shart: Hech qanday qo'shimcha so'z, belgi yoki izoh yozilmaydi.  Faqat comment, va u ham ```(triple backtick) orasida bo'ladi boshiga emoji qo'shiladi. Endi mana bu o'zgarishga comment yoz: " + string(changes))
	re := regexp.MustCompile("(?s)```(.*?)```")
	matches := re.FindStringSubmatch(result)
	var comment string
	if len(matches) >= 2 {
		comment = matches[1]
	} else {
		comment = result
	}
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewText().Title("Comment:").Value(&comment),
		),
	)
	err = form.Run()
	if err != nil {
		fmt.Print(err.Error())
	}
	cmd := a.GitService.Commit(comment)
	res, err := cmd.Output()
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Print(string(res))
}
