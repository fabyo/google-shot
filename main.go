package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

func main() {
	// cria o contexto do chromedp
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// usando timeout pra não ficar preso
	ctx, cancel = context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	url := "https://www.google.com.br"

	var screenshot []byte
	var pdfBytes []byte

	err := chromedp.Run(ctx,
		chromedp.Navigate(url),

		chromedp.Sleep(3*time.Second),

		// screenshot da página
		chromedp.FullScreenshot(&screenshot, 90),

		// gera o PDF
		chromedp.ActionFunc(func(ctx context.Context) error {
			pdf, _, err := page.PrintToPDF().
				WithPrintBackground(true).
				Do(ctx)
			if err != nil {
				return err
			}
			pdfBytes = pdf
			return nil
		}),
	)
	if err != nil {
		log.Fatalf("erro ao capturar página: %v", err)
	}

	// salva PNG
	if err := os.WriteFile("google.png", screenshot, 0644); err != nil {
		log.Fatalf("erro ao salvar PNG: %v", err)
	}

	// salva PDF
	if err := os.WriteFile("google.pdf", pdfBytes, 0644); err != nil {
		log.Fatalf("erro ao salvar PDF: %v", err)
	}

	log.Println("Arquivos gerados: google.png e google.pdf")
}
