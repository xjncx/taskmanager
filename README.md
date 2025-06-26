
## Task Manager API

Микросервис для управления задачами с асинхронным выполнением (3-5 мин)

### Запуск

Склонируйте репозиторий:
```bash
git clone https://github.com/xjncx/taskmanager/
cd taskmanager
```

Создайте .env файл

```bash
echo "PORT=8080" > .env
```
Запустите:

```bash
make run
```
или 

```bash
go run ./cmd/taskmanager
```
Сервер будет доступен на http://localhost:8080

**🔧 Технические особенности**

 - Чистая архитектура: разделение на слои (API/Service/Manager/Repo)
 - Потокобезопасность: sync.RWMutex для всех shared-ресурсов
 - Graceful shutdown: корректная отмена работающих задач
 - Детальный error handling: 400/404/500 с понятными сообщениями
 - In-memory хранилище: с thread-safe реализацией

**📡 API Endpoints**
Создать задачу

```bash
POST /tasks
```
Ответ:

``` json
{"id": "550e8400-e29b-41d4-a716-446655440000"} 
```
Проверить статус

```bash
GET /tasks/{id}
```

Ответ (пример):

```json
{
  "id": "550e8400...",
  "status": "running", 
  "createdAt": "2023-01-01T15:04:05Z",
  "duration": 125000000000,
  "result": "unknown or in process"
}
```
Удалить задачу
```bash
DELETE /tasks/{id}
```

Ответ:

```json
{"message": "task deleted successfully"}
```
⏱️ **Состояния задачи**
pending → running → done/cancelled

Результат: success/cancelled
