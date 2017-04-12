package common

func StartUp() {
	// Initialize AppConfig variable
	initConfig()
	// Start a MongoDB session
	createDbSession()
	// Add indexes into MongoDB
	addIndexes()
}
