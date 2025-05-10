package seeders

import (
	"math/rand"
	"time"

	"github.com/Eef-M/EventHub/backend/initializers"
	"github.com/Eef-M/EventHub/backend/models"
	"github.com/Eef-M/EventHub/backend/utils"
	"github.com/google/uuid"
)

func SeedEvents(user []models.User) []models.Event {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	startDate, _ := utils.ParseDate("2025-06-01")
	endDate, _ := utils.ParseDate("2025-10-31")

	events := []models.Event{
		{
			ID:          uuid.New(),
			OrganizerID: user[0].ID,
			Title:       "Vue Conference 2025",
			Location:    "Jakarta",
			Date:        utils.RandomDateBetween(startDate, endDate, r),
			Time:        utils.RandomTimeWIB(r),
			Category:    "Webinar",
			Description: "A conference about Vue.js and modern web development.",
			BannerURL:   "https://images.pexels.com/photos/3861969/pexels-photo-3861969.jpeg",
		},
		{
			ID:          uuid.New(),
			OrganizerID: user[0].ID,
			Title:       "Tech Music Fest",
			Location:    "Bandung",
			Date:        utils.RandomDateBetween(startDate, endDate, r),
			Time:        utils.RandomTimeWIB(r),
			Category:    "Concert",
			Description: "Music and tech come together in this spectacular event.",
			BannerURL:   "https://images.pexels.com/photos/167636/pexels-photo-167636.jpeg",
		},
		{
			ID:          uuid.New(),
			OrganizerID: user[0].ID,
			Title:       "Startup Workshop",
			Location:    "Surabaya",
			Date:        utils.RandomDateBetween(startDate, endDate, r),
			Time:        utils.RandomTimeWIB(r),
			Category:    "Workshop",
			Description: "A workshop for aspiring startup founders.",
			BannerURL:   "https://images.pexels.com/photos/3183197/pexels-photo-3183197.jpeg",
		},
		{
			ID:          uuid.New(),
			OrganizerID: user[1].ID,
			Title:       "Data Science Summit",
			Location:    "Yogyakarta",
			Date:        utils.RandomDateBetween(startDate, endDate, r),
			Time:        utils.RandomTimeWIB(r),
			Category:    "Conference",
			Description: "Exploring the latest trends in data science, AI, and machine learning.",
			BannerURL:   "https://images.pexels.com/photos/2004161/pexels-photo-2004161.jpeg",
		},
		{
			ID:          uuid.New(),
			OrganizerID: user[1].ID,
			Title:       "Mobile App Development Bootcamp",
			Location:    "Bali",
			Date:        utils.RandomDateBetween(startDate, endDate, r),
			Time:        utils.RandomTimeWIB(r),
			Category:    "Workshop",
			Description: "Intensive 3-day bootcamp for building modern mobile applications.",
			BannerURL:   "https://images.pexels.com/photos/1181244/pexels-photo-1181244.jpeg",
		},
		{
			ID:          uuid.New(),
			OrganizerID: user[1].ID,
			Title:       "Gaming Industry Expo",
			Location:    "Jakarta",
			Date:        utils.RandomDateBetween(startDate, endDate, r),
			Time:        utils.RandomTimeWIB(r),
			Category:    "Exhibition",
			Description: "Showcasing the future of gaming technology and innovation.",
			BannerURL:   "https://images.pexels.com/photos/159393/gamepad-video-game-controller-game-controller-controller-159393.jpeg",
		},
	}

	for _, event := range events {
		initializers.DB.Create(&event)
	}

	return events
}
