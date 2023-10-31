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

	val, err := ViewValueInBucket(db, DB_BUCKET_CORE, DB_KEY_WORKSPACE)
	if err != nil {
		return "", err
	}

	return val, nil
}

func SetCurrentWorkspace(value string) error {
	db, err := OpenDB()
	if err != nil {
		return err
	}
	defer db.Close()

	err = SetValueInBucket(db, []byte(DB_BUCKET_CORE), []byte(DB_KEY_WORKSPACE), []byte(value))
	if err != nil {
		return err
	}

	return nil
}
