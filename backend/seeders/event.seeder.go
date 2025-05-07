package seeders

import (
	"time"

	"github.com/Eef-M/EventHub/backend/initializers"
	"github.com/Eef-M/EventHub/backend/models"
)

func SeedEvents(organizer []models.User) {
	if len(organizer) == 0 {
		return
	}

	events := []models.Event{
		{
			OrganizerID: organizer[0].ID,
			Title:       "Vue Conference 2025",
			Location:    "Jakarta",
			Date:        time.Date(2025, 6, 10, 0, 0, 0, 0, time.UTC),
			Category:    "Webinar",
			Description: "A conference about Vue.js and modern web development.",
			BannerURL:   "https://images.pexels.com/photos/3861969/pexels-photo-3861969.jpeg",
		},
		{
			OrganizerID: organizer[0].ID,
			Title:       "Tech Music Fest",
			Location:    "Bandung",
			Date:        time.Date(2025, 5, 12, 0, 0, 0, 0, time.UTC),
			Category:    "Concert",
			Description: "Music and tech come together in this spectacular event.",
			BannerURL:   "https://images.pexels.com/photos/167636/pexels-photo-167636.jpeg",
		},
		{
			OrganizerID: organizer[0].ID,
			Title:       "Startup Workshop",
			Location:    "Surabaya",
			Date:        time.Date(2025, 7, 1, 0, 0, 0, 0, time.UTC),
			Category:    "Workshop",
			Description: "A workshop for aspiring startup founders.",
			BannerURL:   "https://images.pexels.com/photos/3183197/pexels-photo-3183197.jpeg",
		},
		{
			OrganizerID: organizer[1].ID,
			Title:       "Data Science Summit",
			Location:    "Yogyakarta",
			Date:        time.Date(2025, 8, 15, 0, 0, 0, 0, time.UTC),
			Category:    "Conference",
			Description: "Exploring the latest trends in data science, AI, and machine learning.",
			BannerURL:   "https://images.pexels.com/photos/2004161/pexels-photo-2004161.jpeg",
		},
		{
			OrganizerID: organizer[1].ID,
			Title:       "Mobile App Development Bootcamp",
			Location:    "Bali",
			Date:        time.Date(2025, 9, 5, 0, 0, 0, 0, time.UTC),
			Category:    "Workshop",
			Description: "Intensive 3-day bootcamp for building modern mobile applications.",
			BannerURL:   "https://images.pexels.com/photos/1181244/pexels-photo-1181244.jpeg",
		},
		{
			OrganizerID: organizer[1].ID,
			Title:       "Gaming Industry Expo",
			Location:    "Jakarta",
			Date:        time.Date(2025, 10, 20, 0, 0, 0, 0, time.UTC),
			Category:    "Exhibition",
			Description: "Showcasing the future of gaming technology and innovation.",
			BannerURL:   "https://images.pexels.com/photos/159393/gamepad-video-game-controller-game-controller-controller-159393.jpeg",
		},
	}

	for _, event := range events {
		initializers.DB.Create(&event)
	}
}
