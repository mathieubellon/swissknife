package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/joho/godotenv"
	"github.com/supabase-community/supabase-go"
	"github.com/urfave/cli/v2"
)

type Competitor struct {
	Name string
	URL  string
}

var competitors = []Competitor{
	{"gitguardian", "https://marketplace.visualstudio.com/items?itemName=gitguardian-secret-security.gitguardian"},
	{"cycode", "https://marketplace.visualstudio.com/items?itemName=cycode.cycode"},
	{"snyk", "https://marketplace.visualstudio.com/items?itemName=snyk-security.snyk-vulnerability-scanner-vs"},
	{"mend2019", "https://marketplace.visualstudio.com/items?itemName=Mend.mend-vs2019"},
	{"mend2022", "https://marketplace.visualstudio.com/items?itemName=Mend.mend-vs2022"},
	{"sap_credentialdigger", "https://marketplace.visualstudio.com/items?itemName=SAPOSS.vs-code-extension-for-project-credential-digger"},
}

func vscode(Ctx *cli.Context) error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	SUPABASE_API_URL := os.Getenv("SUPABASE_API_URL")
	SUPABASE_API_KEY := os.Getenv("SUPABASE_API_KEY")

	client, err := supabase.NewClient(SUPABASE_API_URL, SUPABASE_API_KEY, &supabase.ClientOptions{})
	if err != nil {
		fmt.Println("cannot initalize client", err)
	}

	counts := make(map[string]int)

	for _, competitor := range competitors {
		count, err := getInstallCount(competitor.URL)
		if err != nil {
			log.Printf("Error processing %s: %v", competitor.Name, err)
			continue
		}
		counts[competitor.Name] = count
	}

	resp, respcount, err := client.From("vscode_downloads").Insert(counts, false, "", "", "").Execute()
	if err != nil {
		log.Fatalf("Error inserting data into Supabase: %v", err)
		return err
	}
	if Ctx.Bool("verbose") {
		log.Printf("Response: %v, Count: %v", string(resp), respcount)
	}
	return nil
}

func getInstallCount(url string) (int, error) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, fmt.Errorf("failed to retrieve the page: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("failed to retrieve the page. Status code: %d", resp.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return 0, fmt.Errorf("failed to parse the page: %w", err)
	}

	installsText := doc.Find("span.installs-text").Text()
	if installsText == "" {
		return 0, fmt.Errorf("the specified element was not found")
	}

	countStr := strings.ReplaceAll(strings.Split(installsText, " ")[1], ",", "")
	count, err := strconv.Atoi(countStr)
	if err != nil {
		return 0, fmt.Errorf("failed to convert install count to integer: %w", err)
	}

	return count, nil
}
