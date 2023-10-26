package db

const (
	DB_KEY_WORKSPACE = "workspace"
)

func GetCurrentWorkspace() (string, error) {
	db, err := OpenDB()
	if err != nil {
		return "", err
	}
	defer db.Close()

	err = CreateBucketIfNotExist(db, DB_BUCKET_CORE)
	if err != nil {
		return "", err
	}

	val, err := ViewValueInBucket(db, DB_BUCKET_CORE, DB_KEY_WORKSPACE)
	if err != nil {
		return "", err
	}

	return val, nil
}
