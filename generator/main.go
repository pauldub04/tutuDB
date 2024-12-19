package main

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"fmt"
	"log"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "tutu"
)

func createConnection() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected!")
	return db
}

//----------------------------------

func generateOperators(db *sql.DB, count int) {
	var operatorCount int
	err := db.QueryRow(`SELECT COUNT(*) FROM operators`).Scan(&operatorCount)
	if err != nil {
		log.Fatal(err)
	}

	for operatorCount < count {
		name := gofakeit.Company()
		contactInfo := gofakeit.Sentence(5)
		country := gofakeit.Country()

		_, err := db.Exec(`INSERT INTO operators (name, contact_info, country) VALUES ($1, $2, $3)`,
			name, contactInfo, country)
		if err != nil {
			log.Fatal(err)
		}

		err = db.QueryRow(`SELECT COUNT(*) FROM operators`).Scan(&operatorCount)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Operators: ", operatorCount)
}

func randomVehicleType() string {
	countries := []string{"Plane", "Train", "Bus"}
	return countries[rand.Intn(len(countries))]
}

func randomVehicleModel(vehicle string) string {
	planes := []string{"Airbus A320", "Boeing 737", "Sukhoi Superjet 100", "Airbus A380", "Boeing 747", "Boeing 777", "Airbus A350", "Boeing 787", "Embraer E-Jet", "Bombardier CRJ", "Antonov An-148", "Tupolev Tu-204", "Ilyushin Il-96", "McDonnell Douglas MD-80", "Airbus A330", "Boeing 767", "Airbus A340", "Boeing 757", "Airbus A310", "Boeing 727", "Boeing 737 MAX", "Boeing 717", "Boeing 737"}
	trains := []string{"Sapsan", "Swallow", "Electric train", "High-speed train", "Locomotive", "Express train"}
	buses := []string{"MAZ", "LiAZ", "PAZ", "GAZ", "NEFAZ", "KAVZ", "Bogdan", "Mercedes"}
	if vehicle == "Plane" {
		return planes[rand.Intn(len(planes))]
	} else if vehicle == "Train" {
		return trains[rand.Intn(len(trains))]
	} else {
		return buses[rand.Intn(len(buses))]
	}
}

func generateTransport(db *sql.DB, count int) {
	var operatorCount int
	err := db.QueryRow(`SELECT COUNT(*) FROM operators`).Scan(&operatorCount)
	if err != nil {
		log.Fatal(err)
	}

	var transportCount int
	err = db.QueryRow(`SELECT COUNT(*) FROM transport`).Scan(&transportCount)
	if err != nil {
		log.Fatal(err)
	}

	for transportCount < count {
		transportType := randomVehicleType()
		model := randomVehicleModel(transportType)
		capacity := gofakeit.Number(30, 300)
		operatorID := gofakeit.Number(1, operatorCount)
		_, err := db.Exec(`INSERT INTO transport (type, model, capacity, operator_id) VALUES ($1, $2, $3, $4)`,
			transportType, model, capacity, operatorID)
		if err != nil {
			log.Fatal(err)
		}

		err = db.QueryRow(`SELECT COUNT(*) FROM transport`).Scan(&transportCount)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Transport: ", transportCount)
}

func generateStations(db *sql.DB, count int) {
	var stationsCount int
	err := db.QueryRow(`SELECT COUNT(*) FROM stations`).Scan(&stationsCount)
	if err != nil {
		log.Fatal(err)
	}

	for stationsCount < count {
		city := gofakeit.City()
		adj := gofakeit.Adjective()
		adj = strings.ToUpper(adj[:1]) + strings.ToLower(adj[1:])
		name := city + " " + adj + " Station"
		stationType := randomVehicleType()
		latitude := fmt.Sprintf("%.6f", gofakeit.Latitude())
		longitude := fmt.Sprintf("%.6f", gofakeit.Longitude())

		_, err := db.Exec(`INSERT INTO stations (name, city, type, latitude, longitude) VALUES ($1, $2, $3, $4, $5)`,
			name, city, stationType, latitude, longitude)
		if err != nil {
			log.Fatal(err)
		}

		err = db.QueryRow(`SELECT COUNT(*) FROM stations`).Scan(&stationsCount)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Stations: ", stationsCount)
}

func generateRoutes(db *sql.DB, count int) {
	var routesCount int
	err := db.QueryRow(`SELECT COUNT(*) FROM routes`).Scan(&routesCount)
	if err != nil {
		log.Fatal(err)
	}

	var stationsCount int
	err = db.QueryRow(`SELECT COUNT(*) FROM stations`).Scan(&stationsCount)
	if err != nil {
		log.Fatal(err)
	}

	var operatorCount int
	err = db.QueryRow(`SELECT COUNT(*) FROM operators`).Scan(&operatorCount)
	if err != nil {
		log.Fatal(err)
	}

	for routesCount < count {
		startStationID := rand.Intn(stationsCount) + 1
		endStationID := rand.Intn(stationsCount) + 1
		transportType := randomVehicleType()

		if startStationID == endStationID {
			continue
		}

		var startStationType string
		err = db.QueryRow(`SELECT type FROM stations WHERE station_id = ` + strconv.Itoa(startStationID)).Scan(&startStationType)
		if err != nil {
			continue
		}
		var endStationType string
		err = db.QueryRow(`SELECT type FROM stations WHERE station_id = ` + strconv.Itoa(endStationID)).Scan(&endStationType)
		if err != nil {
			continue
		}

		if transportType != startStationType || startStationType != endStationType {
			continue
		}

		var operatorCount int
		err = db.QueryRow(`SELECT COUNT(*) FROM operators`).Scan(&operatorCount)
		if err != nil {
			log.Fatal(err)
		}

		duration := gofakeit.Number(30, 720)
		operatorID := rand.Intn(operatorCount) + 1

		_, err := db.Exec(`INSERT INTO routes (start_station_id, end_station_id, duration, transport_type, operator_id) VALUES ($1, $2, $3, $4, $5)`,
			startStationID, endStationID, duration, transportType, operatorID)
		if err != nil {
			log.Fatal(err)
		}

		err = db.QueryRow(`SELECT COUNT(*) FROM routes`).Scan(&routesCount)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Routes: ", routesCount)
}

func randomTime() time.Time {
	now := time.Now()
	randomDays := gofakeit.Number(0, 30)
	randomHours := gofakeit.Number(0, 23)
	randomMinutes := gofakeit.Number(0, 59)
	return now.AddDate(0, 0, randomDays).Add(time.Duration(randomHours)*time.Hour + time.Duration(randomMinutes)*time.Minute)
}

func indexOf(slice []string, item string) int {
	for i, v := range slice {
		if v == item {
			return i
		}
	}
	return -1
}

func randomDaysOfWeek() string {
	days := []string{"Пн", "Вт", "Ср", "Чт", "Пт", "Сб", "Вс"}
	count := gofakeit.Number(1, 7)
	rand.Shuffle(len(days), func(i, j int) {
		days[i], days[j] = days[j], days[i]
	})
	selectedDays := days[:count]
	sort.Slice(selectedDays, func(i, j int) bool {
		return indexOf(days, selectedDays[i]) < indexOf(days, selectedDays[j])
	})
	return strings.Join(selectedDays, ",")
}

func generateSchedules(db *sql.DB, count int) {
	var schedulesCount int
	err := db.QueryRow(`SELECT COUNT(*) FROM schedules`).Scan(&schedulesCount)
	if err != nil {
		log.Fatal(err)
	}

	var routesCount int
	err = db.QueryRow(`SELECT COUNT(*) FROM routes`).Scan(&routesCount)
	if err != nil {
		log.Fatal(err)
	}

	var transportCount int
	err = db.QueryRow(`SELECT COUNT(*) FROM transport`).Scan(&transportCount)
	if err != nil {
		log.Fatal(err)
	}

	for schedulesCount < count {
		routeID := rand.Intn(routesCount) + 1
		transportID := rand.Intn(transportCount) + 1

		departureTime := gofakeit.DateRange(time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC))
		duration := gofakeit.Number(30, 720)
		arrivalTime := departureTime.Add(time.Duration(duration) * time.Minute)
		daysOfWeek := randomDaysOfWeek()

		_, err := db.Exec(`INSERT INTO schedules (route_id, departure_time, arrival_time, days_of_week, transport_id) VALUES ($1, $2, $3, $4, $5)`,
			routeID, departureTime, arrivalTime, daysOfWeek, transportID)
		if err != nil {
			log.Fatal(err)
		}

		err = db.QueryRow(`SELECT COUNT(*) FROM schedules`).Scan(&schedulesCount)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Schedules: ", schedulesCount)
}

func generatePromotions(db *sql.DB, count int) {
	var promotionsCount int
	err := db.QueryRow(`SELECT COUNT(*) FROM promotions`).Scan(&promotionsCount)
	if err != nil {
		log.Fatal(err)
	}

	for promotionsCount < count {
		name := gofakeit.Company() + strings.ToLower(gofakeit.Sentence(2))
		discount := gofakeit.Float64Range(5.00, 50.00)
		startDate := gofakeit.Date()
		endDate := startDate.AddDate(0, 0, gofakeit.Number(7, 30))
		conditions := gofakeit.Paragraph(1, 2, 5, " ")

		_, err := db.Exec(`INSERT INTO promotions (name, discount, start_date, end_date, conditions) VALUES ($1, $2, $3, $4, $5)`,
			name, discount, startDate, endDate, conditions)
		if err != nil {
			log.Fatal(err)
		}

		err = db.QueryRow(`SELECT COUNT(*) FROM promotions`).Scan(&promotionsCount)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Promotions: ", promotionsCount)
}

func hashPassword(password string) string {
	hash := sha256.Sum256([]byte(password))
	return hex.EncodeToString(hash[:])
}

func generateUsers(db *sql.DB, count int) {
	var usersCount int
	err := db.QueryRow(`SELECT COUNT(*) FROM users`).Scan(&usersCount)
	if err != nil {
		log.Fatal(err)
	}

	for usersCount < count {
		name := gofakeit.Name()
		email := gofakeit.Email()
		password := gofakeit.Password(true, true, true, true, false, 12)

		phone := gofakeit.Phone()
		createdAt := gofakeit.DateRange(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC))
		hashedPassword := hashPassword(password)
		_, err = db.Exec(`INSERT INTO users (name, email, password, phone, created_at) VALUES ($1, $2, $3, $4, $5)`,
			name, email, hashedPassword, phone, createdAt)
		if err != nil {
			log.Fatal(err)
		}

		err = db.QueryRow(`SELECT COUNT(*) FROM users`).Scan(&usersCount)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Users: ", usersCount)
}

func randomDocumentType() string {
	documentTypes := []string{"Passport", "Driver's license", "International passport"}
	return documentTypes[rand.Intn(len(documentTypes))]
}

func generatePassengers(db *sql.DB, count int) {
	var passengersCount int
	err := db.QueryRow(`SELECT COUNT(*) FROM passengers`).Scan(&passengersCount)
	if err != nil {
		log.Fatal(err)
	}

	for passengersCount < count {
		firstName := gofakeit.FirstName()
		lastName := gofakeit.LastName()
		birthDate := gofakeit.DateRange(time.Date(1950, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC))
		documentType := randomDocumentType()
		documentNumber := gofakeit.Number(10000000000, 99999999999)
		if documentType == "Passport" {
			documentNumber /= 10000
		} else if documentType == "Driver's license" {
			documentNumber /= 100
		}

		_, err := db.Exec(`INSERT INTO passengers (first_name, last_name, birth_date, document_type, document_number) VALUES ($1, $2, $3, $4, $5)`,
			firstName, lastName, birthDate, documentType, documentNumber)
		if err != nil {
			log.Fatal(err)
		}

		err = db.QueryRow(`SELECT COUNT(*) FROM passengers`).Scan(&passengersCount)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Passengers: ", passengersCount)
}

func generateTicketsAndPayments(db *sql.DB, count int) {
	var ticketsCount int
	err := db.QueryRow(`SELECT COUNT(*) FROM tickets`).Scan(&ticketsCount)
	if err != nil {
		log.Fatal(err)
	}

	var usersCount int
	err = db.QueryRow(`SELECT COUNT(*) FROM users`).Scan(&usersCount)
	if err != nil {
		log.Fatal(err)
	}

	var schedulesCount int
	err = db.QueryRow(`SELECT COUNT(*) FROM schedules`).Scan(&schedulesCount)
	if err != nil {
		log.Fatal(err)
	}

	var promotionsCount int
	err = db.QueryRow(`SELECT COUNT(*) FROM promotions`).Scan(&promotionsCount)
	if err != nil {
		log.Fatal(err)
	}

	var passengersCount int
	err = db.QueryRow(`SELECT COUNT(*) FROM passengers`).Scan(&passengersCount)
	if err != nil {
		log.Fatal(err)
	}

	for ticketsCount < count {
		userID := rand.Intn(usersCount) + 1
		scheduleID := rand.Intn(schedulesCount) + 1
		price := gofakeit.Price(500, 10000)
		seatNumber := fmt.Sprintf("%d%s", gofakeit.Number(1, 50), string(rune(gofakeit.RandomInt([]int{65, 70}))))
		purchaseDate := gofakeit.Date()
		var promotionID sql.NullInt64
		if promotionsCount > 0 && gofakeit.Bool() {
			promotionID.Int64 = int64(rand.Intn(promotionsCount) + 1)
			promotionID.Valid = true
		} else {
			promotionID.Valid = false
		}
		passengerID := rand.Intn(passengersCount) + 1

		var ticketID int
		err := db.QueryRow(`
            INSERT INTO tickets (user_id, schedule_id, price, seat_number, purchase_date, promotion_id, passenger_id)
            VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING ticket_id`,
			userID, scheduleID, price, seatNumber, purchaseDate, promotionID, passengerID).Scan(&ticketID)
		if err != nil {
			log.Fatal(err)
		}

		amount := price
		paymentDate := purchaseDate.Add(time.Duration(gofakeit.Number(0, 2)) * time.Hour)
		paymentMethod := gofakeit.RandomString([]string{"Cash", "Card", "Online"})
		status := "Paid"
		if gofakeit.Int()%300 == 0 {
			status = "Refunded"
		} else if gofakeit.Int()%500 == 0 {
			status = "Pending"
		} else if gofakeit.Int()%1000 == 0 {
			status = "Canceled"
		}

		_, err = db.Exec(`
            INSERT INTO payments (ticket_id, amount, payment_date, payment_method, status)
            VALUES ($1, $2, $3, $4, $5)`,
			ticketID, amount, paymentDate, paymentMethod, status)
		if err != nil {
			log.Fatal(err)
		}

		err = db.QueryRow(`SELECT COUNT(*) FROM tickets`).Scan(&ticketsCount)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Tickets and Payments: ", ticketsCount)
}

func generateFeedback(db *sql.DB, count int) {
	var feedbackCount int
	err := db.QueryRow(`SELECT COUNT(*) FROM feedback`).Scan(&feedbackCount)
	if err != nil {
		log.Fatal(err)
	}

	var usersCount int
	err = db.QueryRow(`SELECT COUNT(*) FROM users`).Scan(&usersCount)
	if err != nil {
		log.Fatal(err)
	}

	var routesCount int
	err = db.QueryRow(`SELECT COUNT(*) FROM routes`).Scan(&routesCount)
	if err != nil {
		log.Fatal(err)
	}

	for feedbackCount < count {
		userID := rand.Intn(usersCount) + 1
		routeID := rand.Intn(routesCount) + 1
		rating := gofakeit.Number(1, 5)
		comment := gofakeit.Sentence(gofakeit.Number(10, 100))
		createdAt := gofakeit.DateRange(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC))

		_, err := db.Exec(`INSERT INTO feedback (user_id, route_id, rating, comment, created_at) VALUES ($1, $2, $3, $4, $5)`,
			userID, routeID, rating, comment, createdAt)
		if err != nil {
			log.Fatal(err)
		}

		err = db.QueryRow(`SELECT COUNT(*) FROM feedback`).Scan(&feedbackCount)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Feedback: ", feedbackCount)
}

func main() {
	db := createConnection()
	defer db.Close()

	generateOperators(db, 500)
	generateTransport(db, 2500)
	generateStations(db, 200)
	generateRoutes(db, 1000)
	generateSchedules(db, 5000)
	generatePromotions(db, 100)
	generateUsers(db, 5000)
	generatePassengers(db, 50000)
	generateTicketsAndPayments(db, 7000)
	generateFeedback(db, 1000)

}
