package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

type Conditions struct {
	Summary     string
	Temperature Temperature
	City        string
	Country     string
}

type OWMResponse struct {
	Weather []struct {
		Main string
	}
	Main struct {
		Temp Temperature
	}
	Name string
	Sys  struct {
		Country string
	}
}

type Client struct {
	APIKey     string
	BaseURL    string
	HTTPClient *http.Client
}

type Temperature float64

const Usage = `Usage: weather LOCATION

Example: weather London,UK`

func (t Temperature) Celsius() float64 {
	return float64(t) - 273.15
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

func Get(location, key string) (Conditions, error) {
	c := NewClient(key)
	conditions, err := c.GetWeather(location)
	if err != nil {
		return Conditions{}, err
	}
	return conditions, nil
}

func ParseResponse(data []byte) (Conditions, error) {
	var resp OWMResponse
	err := json.Unmarshal(data, &resp)
	if err != nil {
		return Conditions{}, fmt.Errorf("invalid API response %q: %w", data, err)
	}
	if len(resp.Weather) < 1 {
		return Conditions{}, fmt.Errorf("invalid API response %q: want at least one Weather element", data)
	}
	conditions := Conditions{
		Summary:     resp.Weather[0].Main,
		Temperature: resp.Main.Temp,
		City:        resp.Name,
		Country:     resp.Sys.Country,
	}
	return conditions, nil
}

func (c Client) FormatURL(location string) string {
	return fmt.Sprintf("%s/data/2.5/weather?q=%s&appid=%s", c.BaseURL, location, c.APIKey)
}

func (c *Client) GetWeather(location string) (Conditions, error) {
	URL := c.FormatURL(location)
	resp, err := c.HTTPClient.Get(URL)
	if err != nil {
		return Conditions{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return Conditions{}, fmt.Errorf("unexpected response status %q", resp.Status)
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

func Main() int {
	if len(os.Args) < 2 {
		fmt.Println(Usage)
		return 0
	}
	key := os.Getenv("OPENWEATHERMAP_API_KEY")
	if key == "" {
		fmt.Fprintln(os.Stderr, "Please set the environment variable OPENWEATHERMAP_API_KEY.")
		return 1
	}
	location := strings.Join(os.Args[1:], " ")
	conditions, err := Get(location, key)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	fmt.Printf("Weather for %s, %s: ", conditions.City, conditions.Country)
	fmt.Printf("%s %.1fºC\n", conditions.Summary, conditions.Temperature.Celsius())
	return 0
}
