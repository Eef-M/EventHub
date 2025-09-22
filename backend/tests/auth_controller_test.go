package tests

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Eef-M/EventHub/backend/config"
	"github.com/Eef-M/EventHub/backend/controllers"
	"github.com/Eef-M/EventHub/backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupMockDB(t *testing.T) (sqlmock.Sqlmock, func()) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("failed to open sqlmock db: %v", err)
	}

	dialector := postgres.New(postgres.Config{
		Conn: db,
	})
	gdb, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to init gorm with sqlmock: %v", err)
	}

	config.DB = gdb

	cleanup := func() {
		db.Close()
	}
	return mock, cleanup
}

func setupTestEnvironment() {
	os.Setenv("JWT_SECRET", "test-secret-key")
	os.Setenv("ACCESS_TOKEN_EXPIRE_MINUTES", "15")

	utils.GenerateAccessTokenFunc = func(uid, role string) (string, error) {
		return "dummy-access-token", nil
	}
	utils.CreateRefreshTokenFunc = func(uid string) (string, error) {
		return "dummy-refresh-token", nil
	}
}

func TestLogin_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	setupTestEnvironment()

	mock, cleanup := setupMockDB(t)
	defer cleanup()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	body := controllers.LoginInput{
		Email:    "test@example.com",
		Password: "password123",
	}
	jsonBody, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	userID := uuid.New()
	hashedPass, _ := utils.HashPassword("password123")

	t.Logf("User ID: %s", userID)
	t.Logf("Hashed Password: %s", hashedPass)
	fixedTime := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)

	rows := sqlmock.NewRows([]string{
		"id", "username", "email", "password", "first_name", "last_name", "role", "avatar_url", "created_at", "updated_at",
	}).AddRow(
		userID, "testuser", "test@example.com", hashedPass, "Test", "User", "participant", "http://dummy/avatar.png",
		fixedTime, fixedTime,
	)

	expectedSQL := `SELECT * FROM "users" WHERE email = $1 ORDER BY "users"."id" LIMIT $2`
	mock.ExpectQuery(expectedSQL).
		WithArgs("test@example.com", 1).
		WillReturnRows(rows)

	controllers.Login(c)

	t.Logf("Response Code: %d", w.Code)
	t.Logf("Response Body: %s", w.Body.String())

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d but got %d", http.StatusOK, w.Code)

		var errorResponse map[string]interface{}
		if err := json.Unmarshal(w.Body.Bytes(), &errorResponse); err == nil {
			t.Logf("Error response: %+v", errorResponse)
		}
		return
	}

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	if message, exists := response["message"]; exists {
		assert.Equal(t, "Login success", message)
	} else {
		t.Errorf("Response does not contain 'message' field. Response: %+v", response)
	}
}

func TestLogin_InvalidPassword(t *testing.T) {
	gin.SetMode(gin.TestMode)
	setupTestEnvironment()

	mock, cleanup := setupMockDB(t)
	defer cleanup()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	body := controllers.LoginInput{
		Email:    "test@example.com",
		Password: "wrongpassword",
	}
	jsonBody, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	userID := uuid.New()
	hashedPass, _ := utils.HashPassword("password123")
	fixedTime := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)

	rows := sqlmock.NewRows([]string{
		"id", "username", "email", "password", "first_name", "last_name", "role", "avatar_url", "created_at", "updated_at",
	}).AddRow(
		userID, "testuser", "test@example.com", hashedPass, "Test", "User", "participant", "http://dummy/avatar.png",
		fixedTime, fixedTime,
	)

	expectedSQL := `SELECT * FROM "users" WHERE email = $1 ORDER BY "users"."id" LIMIT $2`
	mock.ExpectQuery(expectedSQL).
		WithArgs("test@example.com", 1).
		WillReturnRows(rows)

	controllers.Login(c)

	assert.Equal(t, http.StatusUnauthorized, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Invalid email or password", response["error"])
}

func TestLogin_UserNotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)
	setupTestEnvironment()

	mock, cleanup := setupMockDB(t)
	defer cleanup()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	body := controllers.LoginInput{
		Email:    "notfound@example.com",
		Password: "password123",
	}
	jsonBody, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	expectedSQL := `SELECT * FROM "users" WHERE email = $1 ORDER BY "users"."id" LIMIT $2`
	mock.ExpectQuery(expectedSQL).
		WithArgs("notfound@example.com", 1).
		WillReturnError(gorm.ErrRecordNotFound)

	controllers.Login(c)

	assert.Equal(t, http.StatusUnauthorized, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Invalid email or password", response["error"])
}

func TestLogin_InvalidJSON(t *testing.T) {
	gin.SetMode(gin.TestMode)
	setupTestEnvironment()

	_, cleanup := setupMockDB(t)
	defer cleanup()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer([]byte(`{"email": "invalid-email"}`)))
	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	controllers.Login(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Contains(t, response["error"], "required")
}

func TestLogin_TokenGenerationError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mock, cleanup := setupMockDB(t)
	defer cleanup()

	os.Setenv("JWT_SECRET", "test-secret-key")
	os.Setenv("ACCESS_TOKEN_EXPIRE_MINUTES", "15")

	utils.GenerateAccessTokenFunc = func(uid, role string) (string, error) {
		return "", assert.AnError
	}
	utils.CreateRefreshTokenFunc = func(uid string) (string, error) {
		return "dummy-refresh-token", nil
	}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	body := controllers.LoginInput{
		Email:    "test@example.com",
		Password: "password123",
	}
	jsonBody, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	userID := uuid.New()
	hashedPass, _ := utils.HashPassword("password123")
	fixedTime := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)

	rows := sqlmock.NewRows([]string{
		"id", "username", "email", "password", "first_name", "last_name", "role", "avatar_url", "created_at", "updated_at",
	}).AddRow(
		userID, "testuser", "test@example.com", hashedPass, "Test", "User", "participant", "http://dummy/avatar.png",
		fixedTime, fixedTime,
	)

	expectedSQL := `SELECT * FROM "users" WHERE email = $1 ORDER BY "users"."id" LIMIT $2`
	mock.ExpectQuery(expectedSQL).
		WithArgs("test@example.com", 1).
		WillReturnRows(rows)

	controllers.Login(c)

	assert.Equal(t, http.StatusInternalServerError, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Could not generate access token", response["error"])
}

func TestLogin_All(t *testing.T) {
	t.Run("Success", TestLogin_Success)
	t.Run("InvalidPassword", TestLogin_InvalidPassword)
	t.Run("UserNotFound", TestLogin_UserNotFound)
	t.Run("InvalidJSON", TestLogin_InvalidJSON)
	t.Run("TokenGenerationError", TestLogin_TokenGenerationError)
}

func TestRegister_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mock, cleanup := setupMockDB(t)
	defer cleanup()

	os.Setenv("JWT_SECRET", "test-secret-key")
	os.Setenv("BASE_URL", "http://localhost:8080")

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	body := controllers.RegisterInput{
		Username:  "newuser",
		FirstName: "New",
		LastName:  "User",
		Email:     "newuser@example.com",
		Role:      "participant",
		Password:  "password123",
	}
	jsonBody, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	mock.ExpectBegin()
	userID := uuid.New()

	expectedSQL := `INSERT INTO "users" ("username","email","password","first_name","last_name","role","avatar_url","created_at","updated_at","id") VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10) RETURNING "id"`
	mock.ExpectQuery(expectedSQL).
		WithArgs(
			"newuser",
			"newuser@example.com",
			sqlmock.AnyArg(),
			"New",
			"User",
			"participant",
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
		).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(userID))

	mock.ExpectCommit()

	controllers.Register(c)

	t.Logf("Response Code: %d", w.Code)
	t.Logf("Response Body: %s", w.Body.String())

	assert.Equal(t, http.StatusCreated, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "User registered successfully", response["message"])

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}

func TestRegister_DuplicateEmail(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mock, cleanup := setupMockDB(t)
	defer cleanup()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	body := controllers.RegisterInput{
		Username:  "existing",
		FirstName: "Existing",
		LastName:  "User",
		Email:     "existing@example.com",
		Role:      "participant",
		Password:  "password123",
	}
	jsonBody, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	mock.ExpectBegin()
	mock.ExpectQuery(`INSERT INTO "users"`).
		WillReturnError(gorm.ErrDuplicatedKey)
	mock.ExpectRollback()

	controllers.Register(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Contains(t, response["error"], "Failed to create user")
}

func TestRegister_DuplicateUsername(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mock, cleanup := setupMockDB(t)
	defer cleanup()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	body := controllers.RegisterInput{
		Username:  "existinguser",
		FirstName: "Existing",
		LastName:  "User",
		Email:     "different@example.com",
		Role:      "participant",
		Password:  "password123",
	}
	jsonBody, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	expectedSQL := `INSERT INTO "users" ("username","email","password","first_name","last_name","role","avatar_url","created_at","updated_at","id") VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10) RETURNING "id"`
	mock.ExpectBegin()
	mock.ExpectQuery(expectedSQL).
		WillReturnError(gorm.ErrDuplicatedKey)
	mock.ExpectRollback()

	controllers.Register(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Contains(t, response["error"], "Failed to create user")
}

func TestRegister_InvalidEmail(t *testing.T) {
	gin.SetMode(gin.TestMode)
	_, cleanup := setupMockDB(t)
	defer cleanup()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	body := controllers.RegisterInput{
		Username:  "newuser",
		FirstName: "New",
		LastName:  "User",
		Email:     "invalid-email-format",
		Role:      "participant",
		Password:  "password123",
	}
	jsonBody, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	controllers.Register(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Contains(t, response["error"], "email")
}

func TestRegister_InvalidRole(t *testing.T) {
	gin.SetMode(gin.TestMode)
	_, cleanup := setupMockDB(t)
	defer cleanup()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	body := controllers.RegisterInput{
		Username:  "newuser",
		FirstName: "New",
		LastName:  "User",
		Email:     "newuser@example.com",
		Role:      "invalid_role",
		Password:  "password123",
	}
	jsonBody, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	controllers.Register(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Contains(t, response["error"], "oneof")
}

func TestRegister_MissingFields(t *testing.T) {
	gin.SetMode(gin.TestMode)
	_, cleanup := setupMockDB(t)
	defer cleanup()

	testCases := []struct {
		name        string
		requestBody string
		expectError string
	}{
		{
			name:        "Missing Username",
			requestBody: `{"first_name":"New","last_name":"User","email":"test@example.com","role":"participant","password":"password123"}`,
			expectError: "required",
		},
		{
			name:        "Empty Password",
			requestBody: `{"username":"newuser","first_name":"New","last_name":"User","email":"test@example.com","role":"participant","password":""}`,
			expectError: "required",
		},
		{
			name:        "Missing First Name",
			requestBody: `{"username":"newuser","last_name":"User","email":"test@example.com","role":"participant","password":"password123"}`,
			expectError: "required",
		},
		{
			name:        "Missing Last Name",
			requestBody: `{"username":"newuser","first_name":"New","email":"test@example.com","role":"participant","password":"password123"}`,
			expectError: "required",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer([]byte(tc.requestBody)))
			req.Header.Set("Content-Type", "application/json")
			c.Request = req

			controllers.Register(c)

			assert.Equal(t, http.StatusBadRequest, w.Code)

			var response map[string]interface{}
			err := json.Unmarshal(w.Body.Bytes(), &response)
			assert.NoError(t, err)
			assert.Contains(t, response["error"], tc.expectError)
		})
	}
}

func TestRegister_InvalidJSON(t *testing.T) {
	gin.SetMode(gin.TestMode)
	_, cleanup := setupMockDB(t)
	defer cleanup()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer([]byte(`{"email":"bad@example.com"}`)))
	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	controllers.Register(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Contains(t, response["error"], "required")
}

func TestRegister_DatabaseError(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mock, cleanup := setupMockDB(t)
	defer cleanup()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	body := controllers.RegisterInput{
		Username:  "newuser",
		FirstName: "New",
		LastName:  "User",
		Email:     "newuser@example.com",
		Role:      "participant",
		Password:  "password123",
	}
	jsonBody, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	expectedSQL := `INSERT INTO "users" ("username","email","password","first_name","last_name","role","avatar_url","created_at","updated_at","id") VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10) RETURNING "id"`
	mock.ExpectBegin()
	mock.ExpectQuery(expectedSQL).
		WillReturnError(errors.New("database connection lost"))
	mock.ExpectRollback()

	controllers.Register(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Contains(t, response["error"], "Failed to create user")
}

func TestRegister_Success_Organizer(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mock, cleanup := setupMockDB(t)
	defer cleanup()

	os.Setenv("BASE_URL", "http://localhost:8080")

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	body := controllers.RegisterInput{
		Username:  "orguser",
		FirstName: "Org",
		LastName:  "User",
		Email:     "org@example.com",
		Role:      "organizer",
		Password:  "password123",
	}
	jsonBody, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	userID := uuid.New()
	expectedSQL := `INSERT INTO "users" ("username","email","password","first_name","last_name","role","avatar_url","created_at","updated_at","id") VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10) RETURNING "id"`

	mock.ExpectBegin()
	mock.ExpectQuery(expectedSQL).
		WithArgs(
			"orguser",
			"org@example.com",
			sqlmock.AnyArg(),
			"Org",
			"User",
			"organizer",
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
		).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(userID))
	mock.ExpectCommit()

	controllers.Register(c)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "User registered successfully", response["message"])
}

func TestRegister_All(t *testing.T) {
	t.Run("Success", TestRegister_Success)
	t.Run("SuccessOrganizer", TestRegister_Success_Organizer)
	t.Run("DuplicateEmail", TestRegister_DuplicateEmail)
	t.Run("DuplicateUsername", TestRegister_DuplicateUsername)
	t.Run("InvalidEmail", TestRegister_InvalidEmail)
	t.Run("InvalidRole", TestRegister_InvalidRole)
	t.Run("MissingFields", TestRegister_MissingFields)
	t.Run("InvalidJSON", TestRegister_InvalidJSON)
	t.Run("DatabaseError", TestRegister_DatabaseError)
}
