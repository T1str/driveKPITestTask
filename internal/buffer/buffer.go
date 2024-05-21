package buffer

import (
	"bytes"
	"errors"
	"mime/multipart"
	"net/http"
	"testtask/internal/config"
	"testtask/internal/entities"
)

type Buffer struct {
	Config config.Config
	Client *http.Client
}

func NewBuffer(cfg config.Config) *Buffer {
	return &Buffer{
		Config: cfg,
		Client: &http.Client{},
	}
}

func (b *Buffer) Retry(req entities.Request) error {
	for i := 0; i < 3; i++ {
		ok, err := b.HandleItem(req)
		if err != nil {
			return err
		}
		if !ok {
			continue
		} else {
			return nil
		}
	}
	return errors.New("failed to handle item")
}

func (b *Buffer) HandleItem(req entities.Request) (bool, error) {
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	fields := map[string]string{
		"period_start":            req.PeriodStart,
		"period_end":              req.PeriodEnd,
		"period_key":              req.PeriodKey,
		"indicator_to_mo_id":      req.IndicatorToMoID,
		"indicator_to_mo_fact_id": req.IndicatorToMoFactID,
		"value":                   req.Value,
		"fact_time":               req.FactTime,
		"is_plan":                 req.IsPlan,
		"auth_user_id":            req.AuthUserID,
		"comment":                 req.Comment,
	}

	for key, val := range fields {
		part, err := writer.CreateFormField(key)
		if err != nil {
			return false, err
		}
		_, err = part.Write([]byte(val))
		if err != nil {
			return false, err
		}
	}

	err := writer.Close()
	if err != nil {
		return false, err
	}

	request, err := http.NewRequest("POST", b.Config.URL, &buf)
	if err != nil {
		return false, err
	}

	request.Header.Set("Authorization", "Bearer "+b.Config.Token)
	request.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := b.Client.Do(request)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, errors.New("request failed")
	}

	return true, nil
}
