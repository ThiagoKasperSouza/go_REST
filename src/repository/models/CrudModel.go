package models

import r "newsRestFiber/src/repository"

type CrudModel struct{}

func (m *CrudModel) Create(client r.DbClient, key string, id string, data []byte) (int64, error) {
	return client.Create(key, id, data).Result()
}

func (m *CrudModel) GetItemById(client r.DbClient, key string, id string) (string, error) {
	return client.GetItemById(key, id).Result()
}

func (m *CrudModel) GetAll(client r.DbClient, key string, id string) (map[string]string, error) {
	return client.GetAll(key).Result()
}

func (m *CrudModel) Update(client r.DbClient, key string, id string, data []byte) (int64, error) {
	return client.Update(key, id, data).Result()
}

func (m *CrudModel) Delete(client r.DbClient, key string, id string) (int64, error) {
	return client.Delete(key, id).Result()
}
