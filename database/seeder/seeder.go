package seeder

import "github.com/Final-Project-Kelompok-3/payments/database"


func Seed() {

	conn := database.GetConnection()

	paymentMethodTableSeeder(conn)
	// otherTableSeeder(conn)
}