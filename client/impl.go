package client

import "github.com/odysseia-greek/hippokrates/client/models"

type OdysseiaClient interface {
	Alexandros() Alexandros
	Dionysios() Dionysios
	Herodotos() Herodotos
	Sokrates() Sokrates
	Solon() Solon
}

type Odysseia struct {
	alexandros *AlexandrosImpl
	dionysios  *DionysiosImpl
	herodotos  *HerodotosImpl
	sokrates   *SokratesImpl
	solon      *SolonImpl
}

type Alexandros interface {
	Health() (*models.Health, error)
	QueryWord(word string) ([]models.Meros, error)
}

type Dionysios interface {
	Health() (*models.Health, error)
	CheckGrammar(word string) (*models.DeclensionTranslationResults, error)
}

type Herodotos interface {
	Health() (*models.Health, error)
	Authors() (*models.Authors, error)
	Books(author string) (*models.Books, error)
	CreateQuestion(author, book string) (*models.CreateSentenceResponse, error)
	CheckSentence(requestBody models.CheckSentenceRequest) (*models.CheckSentenceResponse, error)
}

type Sokrates interface {
	Health() (*models.Health, error)
	Methods() (*models.Methods, error)
	Categories(method string) (*models.Categories, error)
	LastChapter(method, category string) (*models.LastChapterResponse, error)
	CreateQuestion(method, category, chapter string) (models.QuizResponse, error)
	Answer(request models.CheckAnswerRequest) (*models.CheckAnswerResponse, error)
}

type Solon interface {
	Health() (*models.Health, error)
	OneTimeToken() (*models.TokenResponse, error)
	Register(requestBody models.SolonCreationRequest) (*models.SolonResponse, error)
}

type ClientConfig struct {
	Scheme        string
	AlexandrosUrl string
	DionysiosUrl  string
	HerodotosUrl  string
	SokratesUrl   string
	SolonUrl      string
}

func NewClient(config ClientConfig) (OdysseiaClient, error) {
	client := NewHttpClient()

	alexandrosImpl, err := NewAlexandrosImpl(config.Scheme, config.AlexandrosUrl, client)
	if err != nil {
		return nil, err
	}

	dionysiosImpl, err := NewDionysiosImpl(config.Scheme, config.DionysiosUrl, client)
	if err != nil {
		return nil, err
	}

	herodotosImpl, err := NewHerodotosImpl(config.Scheme, config.HerodotosUrl, client)
	if err != nil {
		return nil, err
	}

	sokratesImpl, err := NewSokratesImpl(config.Scheme, config.SokratesUrl, client)
	if err != nil {
		return nil, err
	}

	solonImpl, err := NewSolonImpl(config.Scheme, config.SolonUrl, client)
	if err != nil {
		return nil, err
	}

	return &Odysseia{
		alexandros: alexandrosImpl,
		dionysios:  dionysiosImpl,
		herodotos:  herodotosImpl,
		sokrates:   sokratesImpl,
		solon:      solonImpl,
	}, nil
}

func (o *Odysseia) Alexandros() Alexandros {
	if o == nil {
		return nil
	}
	return o.alexandros
}

func (o *Odysseia) Dionysios() Dionysios {
	if o == nil {
		return nil
	}
	return o.dionysios
}

func (o *Odysseia) Herodotos() Herodotos {
	if o == nil {
		return nil
	}
	return o.herodotos
}

func (o *Odysseia) Sokrates() Sokrates {
	if o == nil {
		return nil
	}
	return o.sokrates
}

func (o *Odysseia) Solon() Solon {
	if o == nil {
		return nil
	}
	return o.solon
}
