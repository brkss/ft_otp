package domain



type KeyRepository interface {
	Load() (*Key, error);
	Save(key *Key) error;
}
