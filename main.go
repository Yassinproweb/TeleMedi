package main

import (
	"database/sql"
	"log"
	"time"

	"github.com/Yassinproweb/TeleMedi/data"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/template/html/v2"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	data.InitDB()
	defer data.DB.Close()

	app := fiber.New(fiber.Config{
		Views: html.New("./templates", ".html"),
	})

	store := session.New(session.Config{
		Expiration:     24 * time.Hour,
		CookieHTTPOnly: true,
		CookieSecure:   false, // Set to true in production with HTTPS
	})

	app.Static("/static", "./static")

	// Authentication middleware
	isAuthenticated := func(c *fiber.Ctx) error {
		sess, err := store.Get(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Session error")
		}
		doctorID := sess.Get("doctor_id")
		if doctorID == nil {
			return c.Redirect("/signin")
		}
		return c.Next()
	}

	// Routes
	app.Get("/signup", func(c *fiber.Ctx) error {
		return c.Render("signup", nil)
	})

	app.Get("/signin", func(c *fiber.Ctx) error {
		return c.Render("signin", nil)
	})

	app.Post("/signup", func(c *fiber.Ctx) error {
		name := c.FormValue("name")
		email := c.FormValue("email")
		password := c.FormValue("password")
		skill := c.FormValue("skill")
		title := c.FormValue("title")

		if len(name) < 2 || len(email) < 3 || len(password) < 6 {
			return c.Status(fiber.StatusBadRequest).SendString("<p class='text-red-500 text-center'>Invalid input. Ensure name, email, and password meet requirements.</p>")
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("<p class='text-red-500 text-center'>Error hashing password</p>")
		}

		// Insert doctor into data
		result, err := data.DB.Exec("INSERT INTO doctors (name, email, password, title, skill) VALUES (?, ?, ?, ?, ?)", name, email, hashedPassword, title, skill)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("<p class='text-red-500 text-center'>Email already exists</p>")
		}

		// Get doctor ID
		doctorID, _ := result.LastInsertId()

		// Create session
		sess, err := store.Get(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("<p class='text-red-500 text-center'>Session error</p>")
		}
		sess.Set("doctor_id", doctorID)
		sess.Set("doctor_name", name)
		if err := sess.Save(); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("<p class='text-red-500 text-center'>Failed to save session</p>")
		}

		// Redirect for HTMX
		c.Set("HX-Redirect", "/dashboard")
		return c.SendStatus(fiber.StatusOK)
	})

	app.Post("/signin", func(c *fiber.Ctx) error {
		email := c.FormValue("email")
		password := c.FormValue("password")

		var storedPassword, name string
		var doctorID int
		err := data.DB.QueryRow("SELECT id, name, password FROM doctors(where email = ?", email).Scan(&doctorID, &name, &storedPassword)
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusUnauthorized).SendString("<p class='text-red-500 text-center'>Invalid email or password</p>")
		}
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("<p class='text-red-500 text-center'>data error</p>")
		}

		// Verify password
		if err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(password)); err != nil {
			return c.Status(fiber.StatusUnauthorized).SendString("<p class='text-red-500 text-center'>Invalid email or password</p>")
		}

		// Create session
		sess, err := store.Get(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("<p class='text-red-500 text-center'>Session error</p>")
		}
		sess.Set("doctor_id", doctorID)
		sess.Set("doctor_name", name)
		if err := sess.Save(); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("<p class='text-red-500 text-center'>Failed to save session</p>")
		}

		// Redirect for HTMX
		c.Set("HX-Redirect", "/dashboard")
		return c.SendStatus(fiber.StatusOK)
	})

	// Protected dashboard route
	app.Get("/dashboard", isAuthenticated, func(c *fiber.Ctx) error {
		sess, err := store.Get(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Session error")
		}
		name := sess.Get("doctor_name").(string)
		return c.Render("dashboard", fiber.Map{
			"Name": name,
		})
	})

	// Logout route
	app.Get("/logout", func(c *fiber.Ctx) error {
		sess, err := store.Get(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Session error")
		}
		if err := sess.Destroy(); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to logout")
		}
		return c.Redirect("/signin")
	})

	// Start server
	log.Fatal(app.Listen(":3000"))
}
