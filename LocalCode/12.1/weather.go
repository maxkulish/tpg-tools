package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type Temperature float64

func (t Temperature) Celsius() float64 {
	return float64(t) - 273.15
}

type Conditions struct {
	Summary     string      // `json:"weather[0].main"`
	Temperature Temperature // `json:"main.temp"`
}

type OWMResponse struct {
	Weather []struct {
		Main string `json:"main"`
	}
	Main struct {
		Temp float64 `json:"temp"`
	}
}

type Client struct {
	APIKey     string
	BaseURL    string
	HTTPClient *http.Client
}

func NewClient(apiKey string) *Client {
	return &Client{
		APIKey:  apiKey,
		BaseURL: "https://api.openweathermap.org",
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c Client) FormatURL(location string) string {
	return fmt.Sprintf("%s/data/2.5/weather?q=%s&appid=%s", c.BaseURL, location, c.APIKey)
}

func ParseResponse(data []byte) (Conditions, error) {
	var resp OWMResponse
	err := json.Unmarshal(data, &resp)
	if err != nil {
		return Conditions{}, fmt.Errorf("invalid API response %s: %w", data, err)
	}
	if len(resp.Weather) == 0 {
		return Conditions{}, fmt.Errorf("invalid API response %s: want at least one Weather element", data)
	}
	conditions := Conditions{
		Summary:     resp.Weather[0].Main,
		Temperature: Temperature(resp.Main.Temp),
	}

	return conditions, nil
}

func MakeAPIRequest(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected response status: %s", resp.Status)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c Client) GetWeather(location string) (Conditions, error) {
	url := c.FormatURL(location)
	resp, err := c.HTTPClient.Get(url)
	if err != nil {
		return Conditions{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Conditions{}, fmt.Errorf("unexpected response status: %q", resp.Status)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Conditions{}, err
	}
	conditions, err := ParseResponse(data)
	if err != nil {
		return Conditions{}, err
	}
	return conditions, nil
}

func Get(location, key string) (Conditions, error) {
	c := NewClient(key)
	conditions, err := c.GetWeather(location)
	if err != nil {
		return Conditions{}, err
	}
	return conditions, nil
}

func RunCLI() {
	key := os.Getenv("OPENWEATHERMAP_API_KEY")
	if key == "" {
		log.Fatal("OPENWEATHERMAP_API_KEY not set")
	}

	if len(os.Args) < 2 {
		log.Fatalf("\nUsage: %s LOCATION\n\nExample: %[1]s London,UK", os.Args[0])
	}
	location := os.Args[1]

	c := NewClient(key)

	conditions, err := c.GetWeather(location)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s %.1f°C\n", conditions.Summary, conditions.Temperature.Celsius())
}
