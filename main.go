package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/joho/godotenv"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

func skype() {
	err1 := godotenv.Load()
	if err1 != nil {
		log.Fatal("Error loading .env file")
	}
	USER_EMAIL := os.Getenv("USER_EMAIL")
	USER_PASSWORD := os.Getenv("USER_PASSWORD")
	MESSAGE := os.Getenv("MESSAGE")
	SKYPE_CONTACT := os.Getenv("SKYPE_CONTACT")
	// Start a Selenium WebDriver server
	opts := []selenium.ServiceOption{}
	_, err := selenium.NewChromeDriverService("chromedriver", 9515, opts...)
	if err != nil {
		panic(err)
	}
	// defer service.Stop()

	// Set up Chrome options
	chromeCaps := chrome.Capabilities{
		Path: "",
		Args: []string{
			"--headless", // Run Chrome in headless mode // this command can't open a browser
			"--disable-gpu",
			"--no-sandbox",
			"--disable-dev-shm-usage",
		},
	}
	// ------------------this block of code is for open the browser-----------

	// chromeDriver := chrome.Capabilities{
	// 	Args: []string{
	// 		"--whitelisted-ips=''", // Disable IP whitelisting
	// 		"--disable-extensions",
	// 		"--disable-logging",
	// 		"--disable-notifications",
	// 		"--disable-popup-blocking",
	// 		"--disable-save-password-bubble",
	// 		"--disable-translate",
	// 		"--disable-web-security",
	// 		"--incognito", // Start Chrome in incognito mode
	// 		"--lang=en-US",
	// 	},
	// }

	// Connect to the WebDriver instance
	caps := selenium.Capabilities{"browserName": "chrome"}
	caps.AddChrome(chromeCaps)
	// caps.AddChrome(chromeDriver)

	webDriver, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", 9515))
	if err != nil {
		panic(err)
	}
	// defer webDriver.Quit()

	// Navigate to the Skype login page
	if err := webDriver.Get("https://web.skype.com"); err != nil {
		panic(err)
	}

	// Fill in the login form
	time.Sleep(5 * time.Second)
	// Find the email/phone input field and enter your login credentials
	emailField, err := webDriver.FindElement(selenium.ByCSSSelector, "input[name='loginfmt']")
	if err != nil {
		panic(err)
	}
	emailField.SendKeys(USER_EMAIL)

	// Click the "Next" button to proceed to the password input field
	time.Sleep(5 * time.Second)
	nextButton, err := webDriver.FindElement(selenium.ByCSSSelector, "input[type='submit'][value='Next']")
	if err != nil {
		panic(err)
	}
	nextButton.Click()

	passwordField, err := webDriver.FindElement(selenium.ByCSSSelector, "input[name='passwd']")
	if err != nil {
		panic(err)
	}
	passwordField.SendKeys(USER_PASSWORD)

	time.Sleep(5 * time.Second)
	signInButton, err := webDriver.FindElement(selenium.ByCSSSelector, "input[type='submit'][value='Sign in']")
	if err != nil {
		panic(err)
	}
	signInButton.Click()

	time.Sleep(5 * time.Second)

	stayfield, err := webDriver.FindElement(selenium.ByCSSSelector, "#idSIButton9")

	if err != nil {
		panic(err)
	}
	stayfield.Click()

	time.Sleep(5 * time.Second)

	gotitButton, err := webDriver.FindElement(selenium.ByXPATH, "/html/body/div[1]/div/div/div[1]/div[2]/div/div[2]/div/div[1]/div/div/div/div/div/div[3]/button/div")
	if err != nil {
		panic(err)
	}
	gotitButton.Click()

	time.Sleep(5 * time.Second)

	search_field, err := webDriver.FindElement(selenium.ByCSSSelector, ".r-16dba41")
	if err != nil {
		panic(err)
	}
	search_field.Click()
	time.Sleep(5 * time.Second)
	search_field1, err := webDriver.FindElement(selenium.ByCSSSelector, "input[placeholder='Search Skype']")
	if err != nil {
		panic(err)
	}
	search_field1.SendKeys(SKYPE_CONTACT)

	time.Sleep(10 * time.Second)

	get_user, err := webDriver.FindElement(selenium.ByXPATH, "/html/body/div[1]/div/div/div[1]/div[2]/div/div/div[1]/div/div[1]/div[2]/div[2]/div/div[2]/div/div/div[1]/div/div/div[1]/div[3]")
	if err != nil {
		panic(err)
	}
	get_user.Click()
	time.Sleep(5 * time.Second)

	message_field, err := webDriver.FindElement(selenium.ByCSSSelector, "div[class='public-DraftStyleDefault-block public-DraftStyleDefault-ltr']")
	if err != nil {
		panic(err)
	}
	message_field.SendKeys(MESSAGE)

	time.Sleep(5 * time.Second)

	Sendbtn, err := webDriver.FindElement(selenium.ByCSSSelector, "button[title='Send message']")
	if err != nil {
		panic(err)
	}
	Sendbtn.Click()

}

func RunCronJobs() {
	s := gocron.NewScheduler(time.Local)
	s.Every(1).Day().At("18:44").Do(func() {
		skype()
	})

	s.StartBlocking()
}
func main() {
	RunCronJobs()
}
