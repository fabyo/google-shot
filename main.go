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
	// cria o contexto principal do chromedp
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// timeout pra não ficar preso pra sempre
	ctx, cancel = context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	url := "https://www.google.com.br"

	var screenshot []byte
	var pdfBytes []byte

	err := chromedp.Run(ctx,
		chromedp.Navigate(url),

		// dá um tempinho pra página carregar
		chromedp.Sleep(3*time.Second),

		// screenshot da página inteira (PNG)
		chromedp.FullScreenshot(&screenshot, 90),

		// gerar PDF via DevTools
		chromedp.ActionFunc(func(ctx context.Context) error {
			// repara aqui: 3 retornos
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

	// salva o PNG
	if err := os.WriteFile("google.png", screenshot, 0644); err != nil {
		log.Fatalf("erro ao salvar PNG: %v", err)
	}

	// salva o PDF
	if err := os.WriteFile("google.pdf", pdfBytes, 0644); err != nil {
		log.Fatalf("erro ao salvar PDF: %v", err)
	}

	log.Println("Arquivos gerados: google.png e google.pdf")
}
